// 创建WebSocket连接
const socket = new WebSocket('ws://47.93.236.205:30002/neup_ws');

// 学号和姓名（按照要求格式化）
const studentInfo = '20227291 李卓远\n';

// 监听连接成功事件
socket.addEventListener('open', (event) => {
  // 发送学号和姓名
  socket.send(studentInfo);
  console.log('已发送数据到服务器:', studentInfo);
});

// 监听消息事件
socket.addEventListener('message', (event) => {
  const serverResponse = event.data;
  console.log('服务器返回的信息:', serverResponse);
});

// 监听连接关闭事件
socket.addEventListener('close', (event) => {
  console.log('连接已关闭');
});

// 监听连接错误事件
socket.addEventListener('error', (event) => {
  console.error('发生错误:', event);
});
