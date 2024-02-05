import {useSelector} from "react-redux";
import "./RightTopPart.css"


export function RightTopPart() {
  const userId = useSelector(state => state.chatId.chatId);
  const userName = useSelector((state) => state.userNames.userNames[userId])

  return (
      <div className="RightTopPart">
        <span className="UserName">{userName}</span>
        <span className="UserId">{"@" + userId}</span>
      </div>
  )
}
