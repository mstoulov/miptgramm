import {useState} from "react";
import {loginToServer, registerToServer} from "../../serverApi/sending";


export function LoginModal() {
  const [login, setLogin] = useState("")
  const [userName, setUserName] = useState("")
  const [password, setPassword] = useState("")
  const handleLoginChange = (event) => {
    setLogin(event.target.value)
  }
  const handlePasswordChange = (event) => {
    setPassword(event.target.value)
  }
  const handleUserNameChange = (event) => {
    setUserName(event.target.value)
  }
  const handleLoginPressed = (event) => {
    loginToServer(login, password)
  }
  const handleRegisterPressed = (event) => {
    registerToServer(login, userName, password)
  }

  return (
      <div className="modal" onClick={(event) => event.stopPropagation()}>
        <span className="FieldName">login</span>
        <input className="InputField" onChange={handleLoginChange} />
        <span className="FieldName">password</span>
        <input className="InputField" onChange={handlePasswordChange} />
        <span className="FieldName">user name</span>
        <input className="InputField" onChange={handleUserNameChange} />
        <div className="Buttons">
          <button className="Button" onClick={handleLoginPressed}>Login</button>
          <button className="Button" onClick={handleRegisterPressed}>Register</button>
        </div>
      </div>
  )
}