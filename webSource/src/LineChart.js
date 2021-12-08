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
			let colors = ['red','green','blue','yellow']
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
			arr.push({gt: 300, color: "red"})
		}
		

		// let cut = Math.ceil(list.length / 100)
		// for (let index = 0; index < cut; index++) {
		// 	if (index % 2 === 0) {
		// 		if (index === 0) {
		// 			arr.push({ lte: (index + 1) * 100, color: 'red' })
		// 		} else {
		// 			arr.push({ gt: index * 100, lte: (index + 1) * 100, color: 'red' })
		// 		}
		// 	} else {
		// 		if(index === cut-1){
		// 			arr.push({ gt: (index + 2) * 100, color: 'green' })
		// 		}else{
		// 			arr.push({ gt: index * 100, lte: (index + 1) * 100, color: 'green' })
		// 		}
		// 	}
		// }
		return arr
	}


	useEffect(() => {
		if(dataList.length===0) return
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
					smooth: true,
					data: dataList.map(item => item.totalScore)
				}
			]
		};
		echartsInstance?.setOption(option)
	}, [echartsInstance, dataList])
	return <div ref={chartRef} style={{ width: '100%', height: 400 }}></div>
};

export default LineChart;