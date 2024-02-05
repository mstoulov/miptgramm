import {useDispatch, useSelector} from "react-redux";
import {useEffect, useRef} from "react";
import {Message} from "../Message/Message";
import {markRead} from "../../serverApi/sending";
import {setPinned} from "../../redux/scrollPin";
import "./ChatBody.css"


export function ChatBody() {
  const messagesToShow = useSelector((state) => state.chatMessages.messages)
  const chatId = useSelector(state => state.chatId.chatId)
  const isPinned = useSelector(state => state.scrollPin.pinned)
  const elRef = useRef(null);
  const dispatch = useDispatch()

  markRead(messagesToShow.slice(-1)[0], chatId)

  const handleScroll = (event) => {
    if (event.currentTarget.scrollHeight - event.currentTarget.scrollTop - event.currentTarget.clientHeight <= 5) {
      dispatch(setPinned(true))
    } else {
      dispatch(setPinned(false))
    }
  }

  const executeScroll = () => {
    if (isPinned) {
      elRef.current.scrollIntoView({ behavior: "smooth" });
    }
  }
  useEffect(() => {
    executeScroll()
  }, [messagesToShow]);

  return (
      <ul className="ChatBody" onScroll={handleScroll}>
        {messagesToShow.map((message) => (
            <Message
                key={message.id}
                text={message.text}
                dateTime={message.dateTime}
                senderId={message.senderId}
                isRead={message.isRead}
            />
        ))}
        <div ref={elRef}></div>
      </ul>
  );
}
