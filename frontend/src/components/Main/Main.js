import { LeftPart } from "../LeftPart/LeftPart";
import { RightPart } from "../RightPart/RightPart";
import {Modal} from "../Modal/Modal";
import "./Main.css"


export function Main() {
  return (
      <div className="Main">
        <LeftPart />
        <RightPart />
        <Modal />
      </div>
  )
}
