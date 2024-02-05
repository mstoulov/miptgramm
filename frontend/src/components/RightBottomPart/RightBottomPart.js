import {useState} from "react";
import {useDispatch, useSelector} from "react-redux";
import {sendMessage} from "../../serverApi/sending";
import {setPinned} from "../../redux/scrollPin";
import "./RightBottomPart.css"


export function RightBottomPart() {
  const chatId = useSelector(state => state.chatId.chatId)
  const myId = useSelector((state) => state.myId.myId)
  const dispatch = useDispatch()
  const [typedText, setTypedText] = useState("");

  const handleChangeTypedText = (event) => {
    setTypedText(event.target.value);
  }

  const handleSendMessage = (event) => {
    event.preventDefault()
    dispatch(setPinned(true))
    sendMessage(myId, chatId, typedText)
    setTypedText("")
  }

  const handleKeyDown = (event) => {
    if (event.key === 'Enter') {
      handleSendMessage(event)
    }
  }

  return (
      <div className="RightBottomPart">
        <input className="MessageInput" value={typedText} onChange={handleChangeTypedText} onKeyDown={handleKeyDown}/>
        <button className="SendMessageButton" onClick={handleSendMessage}>Send</button>
      </div>
  );
}
