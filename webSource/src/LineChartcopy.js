import React from 'react';
import { Line } from 'react-chartjs-2';
import moment from 'moment';

const genData = dataList => ({
  labels: dataList.map(item => moment(new Date(item.time)).format('h:mm')),
  datasets: [
    {
      label: '总得分',
      data: dataList.map(item => item.totalScore),
      fill: false,
      // backgroundColor: 'rgb(255, 99, 132)',
      // borderColor: '#000',
      // borderColor: 'rgba(255, 99, 132, 0.2)',
      // borderWidth:1,
      // // borderDash:1,
      // pointBorderWidth:1,
      // pointRadius:1.5,
      // showLine:true
      backgroundColor: [
        '#ff6384',
        '#36a2eb',
        '#cc65fe',
        '#ffce56'
      ]
    },
  ],
});

const options = {
  animations: false,
  visualMap: {
    show: false,
    dimension: 0,
    pieces: [
      {
        lte: 1,
        color: 'green'
      },
      {
        gt: 10,
        lte: 16,
        color: 'red'
      },
      {
        gt: 16,
        color: 'green'
      }
    ]
  },
  scales: {
    yAxes: [
      {
        ticks: {
          beginAtZero: true,
        },

      },
    ],
  },
};

const LineChart = ({ dataList }) => (
  <>
    <Line data={genData(dataList)} options={options} />
  </>
);

export default LineChart;