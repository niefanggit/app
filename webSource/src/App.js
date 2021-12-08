import logo from './logo.svg';
import './App.css';
import 'antd/dist/antd.css'
import LineChart from './LineChart';
import React, { useState, useEffect } from 'react';
import { notification, Row, Col } from 'antd';

const HOST = '182.92.223.34' // "47.93.48.15" //
// const HOST = document.location.host.split(":")[0] // "47.93.48.15" //

function App() {
  let [gameList, setGameList] = useState([])

  let [gameData, setGameData] = useState({})
  useEffect(() => {
    setInterval(async () => {
      try {
        let res = await fetch(`http://${HOST}:8080/gameList`)
        let result = await res.json()
        setGameList(result["info"]) 
      } catch (err) {
        // notification.open({
        //   message: '服务端错误',
        //   description: err,
        // });
      }
    }, 2000)
  }, [])

  useEffect(async () => {
    let curData = {}
    for (let game of gameList) {
      try {
        let gameId = game.gameId
        let res = await fetch(`http://${HOST}:8080/game/${gameId}`)
        let result = await res.json()
        var filter = []
        result.forEach((element,i)=> {
            if(i === 0) {
              filter.push(element)
            } else if(filter[filter.length-1].totalScore != element.totalScore || filter[filter.length-1].quarter != element.quarter || new Date(element.time).getTime() -  new Date(filter[filter.length-1].time).getTime() >= 1000*59+500) {
              filter.push(element)
            }
        });
        curData[gameId] = filter
      } catch(err) {
        // notification.open({
        //   message: '服务端错误',
        //   description: err,
        // });
      }
    }
    setGameData(curData)
  }, [gameList])

  return (
    <div className="App">
      <Row gutter={16} margin={0}>
        {gameList.map(gameInfo => (
            <Col className="gutter-row" style={{marginBottom: 10}} xs={{span: 24}} sm={{span: 12}} xl={{span: 8}}>
              <header>
                {gameInfo.name}
              </header>
              <header>
                <span style={{marginRight: 10, color: "#0066cc"}}>{gameInfo.homeName}</span>
                VS
                <span style={{marginLeft: 10, color: "#cc0000"}}>{gameInfo.awayName}</span>
              </header>
              <LineChart gameId={gameInfo.gameId} dataList={gameData[gameInfo.gameId] || []}/>
            </Col>
        ))}
      </Row>
      
    </div>
  );
}

export default App;
