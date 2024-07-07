import React from 'react'
import ReactDOM from 'react-dom/client'
import {createBrowserRouter, RouterProvider} from "react-router-dom";
import App from './App.jsx'
import * as cmp from './components/components.js'
import './index.css'

const apiBase = 'http://localhost:8888'

const router = createBrowserRouter([
   {
      path: '/',
      element: <App />,
      errorElement: <cmp.ErrorPage />,
      children: [
         {index: true, element: <cmp.Dashboard />},
         {path: "/register", element: <cmp.Register />},
         {path: "/login", element: <cmp.Login />},
      ],
   },
])

ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>,
)

export default apiBase