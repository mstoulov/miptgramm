import {Link} from "react-router-dom";
import {useDispatch, useSelector} from "react-redux";
import {DateTimeParser} from "../../utility/dateTimeParser";
import {setChatId} from "../../redux/chatId";
import {fetchData} from "../../serverApi/sending";
import "./ChatPreview.css"


export function ChatPreview(props) {
  const { lastMessage } = props
  const myId = useSelector((state) => state.myId.myId)
  const chatId = lastMessage.senderId === myId ? lastMessage.receiverId : lastMessage.senderId;
  const currentSelectedChat = useSelector(state => state.chatId.chatId)
  const userName = useSelector((state) => state.userNames.userNames[chatId])
  const dispatch = useDispatch()


  const handlerLink = () => {
    dispatch(setChatId(chatId))
    fetchData(myId, chatId)
  }

  return (
      <Link to={"/chat/" + chatId} onClick={handlerLink} className="ChatPreview">
        <div className={chatId === currentSelectedChat ? "SelectedChatPreview" : ""}>
          <div className="ChatPreviewInfo">
            <p className="ChatName">{userName}</p>
            <div className="ChatPreviewMeta">
              <div className="LastMessageTime">{DateTimeParser(lastMessage.dateTime)} </div>
              <div className={lastMessage.senderId === chatId && !lastMessage.isRead ? "UnreadChatMarker" : ""} />
            </div>
          </div>
          <div className="LastMessageText">{lastMessage.text}</div>
        </div>
      </Link>
  )
}
