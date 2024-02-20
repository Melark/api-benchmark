const express = require("express");
const app = express();
// respond with "hello world" when a GET request is made to the homepage
app.get("/", (req, res) => {
  res.json({ Name: "John", Surname: "Doe" });
});

app.listen(5204);
