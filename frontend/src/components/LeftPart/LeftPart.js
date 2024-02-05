import { LeftTopPart } from "../LeftTopPart/LeftTopPart";
import { ChatPreviewList } from "../ChatPreviewList/ChatPreviewList";
import "./LeftPart.css"


export function LeftPart() {
  return (
      <div className="LeftPart">
        <LeftTopPart />
        <ChatPreviewList />
      </div>
  )
}
