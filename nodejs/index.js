const express = require('express');

const app = express();
app.disable('etag');
app.disable('x-powered-by');
app.use(express.json());

app.post('/', (req, res) => {
  const { a, b } = req.body;
  const sum = a + b;
  res.status(200).json({
    result: sum,
  });
});

app.listen(3000);
