import React, { useEffect, useRef, useState } from 'react';
import moment from 'moment';
import * as echarts from 'echarts';


const LineChart = ({ dataList, gameId }) => {
	const chartRef = useRef(null)
	const [echartsInstance, setEchartsInstance] = useState(null)
	useEffect(() => {
		if (chartRef.current) {
			setEchartsInstance(echarts.init(chartRef.current))
		}
	}, [])
	const showLineColor = (list) => {
		let arr = []
		if(list.length != 0) {
			let colors = ['red','green','blue','black']
			var start = 0
			for (let index = 0; index < list.length ; index++) {
				if(index === list.length-1  || list[index].quarter != list[start].quarter) {
					arr.push({ gte: start, lt: index, color: colors[arr.length]})
					start = index
				}
			}
		} else {
			arr.push({lte: 100, color: "red"})
			arr.push({gt: 100, lte: 200, color: "green"})
			arr.push({gt: 200, lte: 300, color: "red"})
			arr.push({gt: 300, color: "black"})
		}
		return arr
	}

	const getMinMax = (list) =>{
		var max = 0
		var min = 500
		for (let index = 0; index < list.length ; index++) {
			if(list[index].totalScore > max) {
				max = list[index].totalScore;
			} 
			if(list[index].totalScore < min) {
				min = list[index].totalScore;
			}
		}
		return {min:Math.floor(min-(max-min)/5),max:Math.ceil(max+(max -min)/5)}
	}



	useEffect(() => {
		if(dataList.length===0) return
		let minMax = getMinMax(dataList)
		const option = {
			tooltip: {
				trigger: 'axis',
				axisPointer: {
					type: 'cross'
				}
			},
			xAxis: {
				type: 'category',
				boundaryGap: false,
				// prettier-ignore
				data: dataList.map(item => moment(new Date(item.time)).format('h:mm:ss')),
			
			},
			yAxis: {
				type: 'value',
				min: minMax.min,
				max: minMax.max,
				position: 'left'
			},
			visualMap: {
				show: false,
				dimension: 0,
				pieces: showLineColor(dataList)
			},
			series: [
				{
					name: '总得分',
					type: 'line',
					data: dataList.map(item => item.totalScore)
				}
			]
		};
		echartsInstance?.setOption(option)
	}, [echartsInstance, dataList])
	return <div ref={chartRef} style={{ width: '100%', height: 400 }}></div>
};

export default LineChart;