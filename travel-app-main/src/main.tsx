import React from 'react'
import ReactDOM from 'react-dom/client'

import 'swiper/css/bundle';
import './assets/css/navbar.css'
import './assets/css/boxicons.min.css'
import "./assets/styles/index.scss";
import './index.css'

import App from './App.tsx'
import { Provider } from 'react-redux';
import { store } from './app/store.ts';

const root = ReactDOM.createRoot(
  document.getElementById("root") as HTMLElement
);

root.render(
  <React.StrictMode>
    <Provider store={store}>
      <App />
    </Provider>
  </React.StrictMode>
);

