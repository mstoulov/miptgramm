import {checkGotData, checkGotError, checkGotUserName, checkLogin, checkRegister} from "./receiving";
import {setInfoModal, setLoginModal} from "../redux/modal";
import {store} from "../store";

export let conn = new WebSocket("ws://" + "158.160.140.98:8080" + "/ws");

export function setupConnection() {
  conn.onopen = function(evt) {
    console.log("OPEN");
    store.dispatch(setLoginModal())
  }
  conn.onclose = function(evt) {
    console.log("CLOSE");
    conn = null;
  }
  conn.onerror = function(evt) {
    console.log("ERROR: " + evt.data);
  }

  conn.onmessage = (event) => {
    const msg = event.data
    console.log("RESPONSE: " + msg)
    const handled =
        checkGotData(msg) ||
        checkGotUserName(msg) ||
        checkGotError(msg) ||
        checkLogin(msg) ||
        checkRegister(msg)
    if (!handled) {
      store.dispatch(setInfoModal("error: unknown request"))
    }
  }
}
