import createClient from "openapi-fetch";
import type { paths } from "../types/api.d.ts";
import fs from "node:fs";
import path from "node:path";
import { parse } from "smol-toml";

// 1. config.toml の読み込み
const configPath = path.resolve(__dirname, "../../../config/config.toml");
const configFile = fs.readFileSync(configPath, "utf8");
const config = parse(configFile) as any;

// 2. 環境変数による上書き（優先）
const host = process.env.API_HOST || config.api?.host || "localhost";
const port = process.env.API_PORT || config.api?.port || 8080;

const baseUrl = `http://${host}:${port}`;

// openapi-fetch クライアントの初期化 [3, 4]
const client = createClient<paths>({ 
  baseUrl: baseUrl 
});

export default client;