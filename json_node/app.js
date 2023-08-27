const express = require("express");
const bodyParser = require("body-parser");

app = express();
app.use(express.json());
app.use(bodyParser.urlencoded({ extended: true }));

app.get("/", (_, res) => {
  res.json({
    message: "Hello from json_node server",
  });
});

app.post("/post", (req, res) => {
  let customJson = req.body;

  res.status(200).json(customJson);
});

app.post("/postForm", (req, res) => {
  console.log(req.body);
  res.status(200).json(JSON.stringify(req.body));
});

app.listen(3000, () => {
  console.log("App is listening on port 3000");
});
