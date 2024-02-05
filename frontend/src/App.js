import {BrowserRouter, Routes, Route, Navigate} from "react-router-dom";
import {Main} from "./components/Main/Main";
import "./App.css";

export default function App() {
  return (
      <div className="App">
        <BrowserRouter>
          <Routes>
            <Route path="/" element={<Navigate to="/chat/" replace/>} />
            <Route path="/chat/" element={<Main/>} />
            <Route path="/chat/:id" element={<Main/>} />
          </Routes>
        </BrowserRouter>
      </div>
  );
}
