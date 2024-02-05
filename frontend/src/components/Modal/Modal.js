import {useDispatch, useSelector} from "react-redux";
import {setModalClose} from "../../redux/modal";
import {LoginModal} from "./LoginModal";
import "./Modal.css"


export function Modal() {
  const modalType = useSelector((state) => state.modal.type)
  const data = useSelector((state) => state.modal.data)
  const dispatch = useDispatch();

  if (modalType === "") {
    return null
  }

  return (
      <div className="modalContainer" onClick={() => dispatch(setModalClose())}>
        {modalType === "info" ?
            <div className="modal" onClick={(event) => event.stopPropagation()}>
              {data}
            </div> :
            <LoginModal />
        }

      </div>
  );
}
