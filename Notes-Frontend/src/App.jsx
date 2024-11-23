import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import AddNote from './components/AddNote'
import ShowNote from './components/ShowNote'
import SplitScreen from './components/SplitScreen'
import MainScreen from './components/MainScreen'

function App() {
  

  return (
    <>
      {/* <AddNote/>
      <ShowNote/> */}

      <MainScreen/>
      {/* <SplitScreen/> */}
    </>
  )
}

export default App
