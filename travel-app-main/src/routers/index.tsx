import { BrowserRouter, Route, Routes } from "react-router-dom";
import Layout from "../modules/Layout";
import InovationDigital from "../modules/Travel/Page/Travel";

const Routers = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route element={<Layout />}>
          <Route path="/" element={<InovationDigital />} />
        </Route>
      </Routes>
    </BrowserRouter>
  );
};

export default Routers;
