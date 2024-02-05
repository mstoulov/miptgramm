import {useSelector} from "react-redux";
import {DateTimeParser} from "../../utility/dateTimeParser";
import "./Message.css"


export function Message(props) {
  const {text, dateTime, senderId, isRead} = props;
  const myId = useSelector((state) => state.myId.myId)

  return (
      <div className={myId === senderId ? "RightMessage" : "LeftMessage"}>
        <div className="Message">
          <div className="MessageText">{text}</div>
          <div className="MessageMeta">
            <div className="MessageTime">{DateTimeParser(dateTime)}</div>
            <div className={isRead ? "" : "UnreadMessageMarker"}/>
          </div>
        </div>
      </div>
  );
}
