import React from 'react'
import ReactDOM from 'react-dom/client'
import {createBrowserRouter, RouterProvider} from "react-router-dom";
import App from './App.jsx'
import * as cmp from './components/components.js'
import './index.css'

const apiUrlBase = 'http://localhost:8888/api'

const router = createBrowserRouter([
   {
      path: '/',
      element: <App />,
      errorElement: <cmp.ErrorPage />,
      children: [],
   },
])

ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>,
)

export default apiUrlBase