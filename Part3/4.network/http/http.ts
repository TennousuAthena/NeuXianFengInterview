import axios from 'axios';

async function sendData() {
  const url = 'http://47.93.236.205:30002/neup';

  // 学号和姓名
  const data = {
    name: '李卓远',
    number: 20227291,
  };

  try {
    const response = await axios.post(url, data);
    console.log('服务器响应数据:', response.data);
  } catch (error) {
    console.error('发生错误:', (error as Error).message);
  }
}

sendData();
