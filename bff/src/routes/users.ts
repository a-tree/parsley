import { Router } from "express";
import client from "../lib/api.js"; // openapi-fetch クライアント

const router = Router();

// 全ユーザー取得: GET /users
router.get("/", async (req, res) => {
  const { data, error } = await client.GET("/users");
  if (error) {
    console.log("users.ts Error at client.GET ", error) 
    return res.status(500).send("Error");
  }
  
  const html = data.map(u => `<li>${u.name}</li>`).join("");
  res.send(`<ul>${html}</ul>`);
});

// 新規登録: POST /users
router.post("/", async (req, res) => {
  console.log("name: ", req.body.Name, " email: ", req.body.Email)
  const { data, error } = await client.POST("/users", {
    body: { name: req.body.Name, email: req.body.Email }
  });
  if (error) {
    console.log("users.ts Error at client.POST " + error) 
    return res.status(400).send("Failed");
  }
  
  res.send(`<li>${data.name} (Registered!)</li>`);
});

export default router;