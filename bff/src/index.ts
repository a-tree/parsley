import express from "express";
import userRouter from "./routes/users.js";

const app = express();

app.use(express.json());
app.use(express.urlencoded({ extended: true }));

// /users へのリクエストを userRouter に委譲
app.use("/users", userRouter);

app.listen(3000, () => {
  console.log("BFF started on http://localhost:3000");
});