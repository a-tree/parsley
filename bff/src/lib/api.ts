import createClient from "openapi-fetch";
import type { paths } from "../types/api.d.ts";
import fs from "node:fs";
import { fileURLToPath } from "node:url";
import { parse } from "smol-toml";

// config.toml の読み込み
const configUrl = new URL("../../../config/config.toml", import.meta.url);
const configPath = fileURLToPath(configUrl);
const configFile = fs.readFileSync(configPath, "utf8");
const config = parse(configFile) as any;

// 環境変数による上書き
const host = process.env.API_HOST || config.api?.host || "localhost";
const port = process.env.API_PORT || config.api?.port || 8080;

const baseUrl = `http://${host}:${port}`;

// openapi-fetch クライアントの初期化
const client = createClient<paths>({ 
  baseUrl: baseUrl 
});

export default client;