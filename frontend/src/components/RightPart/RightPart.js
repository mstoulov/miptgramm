import { ChatBody } from "../ChatBody/ChatBody";
import { RightTopPart } from "../RightTopPart/RightTopPart";
import { RightBottomPart } from "../RightBottomPart/RightBottomPart";
import "./RightPart.css"


export function RightPart() {
  return (
      <div className="RightPart">
        <RightTopPart userName="ivan ivanych" userId="ivanych228"/>
        <ChatBody />
        <RightBottomPart />
      </div>
  )
}
