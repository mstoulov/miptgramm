import {useSelector} from "react-redux";
import {ChatPreview} from "../ChatPreview/ChatPreview";
import "./ChatPreviewList.css"


export function ChatPreviewList() {
  const previewList = useSelector((state) => state.chatPreviews.lastMessages)
  return (
      <ul className="ChatPreviewList">
        {
          previewList.map((lastMessage) => <ChatPreview key={lastMessage.id} lastMessage={lastMessage}/>)
        }
      </ul>
  )
}
