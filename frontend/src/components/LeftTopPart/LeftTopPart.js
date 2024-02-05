import {useNavigate} from "react-router-dom";
import {useState} from "react";
import {useDispatch, useSelector} from "react-redux";
import {fetchData, fetchUserName} from "../../serverApi/sending";
import {setChatId} from "../../redux/chatId";
import "./LeftTopPart.css"


export function LeftTopPart() {
  const userNames = useSelector((state) => state.userNames.userNames)
  const myId = useSelector(state => state.myId.myId)
  const dispatch = useDispatch()
  const navigate = useNavigate();
  const [typedText, setTypedText] = useState("");

  const handleChangeTypedText = (event) => {
    setTypedText(event.target.value);
  }

  const handleKeyDown = (event) => {
    if (event.key === "Enter") {
      const userId = typedText
      if (userId in userNames) {
        dispatch(setChatId(userId))
        fetchData(myId, userId)
        navigate("/chat/" + userId)
      } else {
        fetchUserName(userId)
      }
    }
  }

  return (
      <div className="LeftTopPart">
        <input className="SearchDialog" value={typedText} onChange={handleChangeTypedText} onKeyDown={handleKeyDown}/>
      </div>
  )
}
