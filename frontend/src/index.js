import {createRoot} from "react-dom/client";
import {Provider} from "react-redux";
import App from "./App";
import {store} from "./store";
import {setupConnection} from "./serverApi/socket";


const root = createRoot(document.getElementById('root'));
setupConnection()

root.render(
    <Provider store={store}>
      <App/>
    </Provider>
);
