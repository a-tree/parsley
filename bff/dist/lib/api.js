import createClient from "openapi-fetch";
import fs from "node:fs";
import { fileURLToPath } from "node:url";
import { parse } from "smol-toml";
// 1. config.toml の読み込み
const configUrl = new URL("../../../config/config.toml", import.meta.url);
const configPath = fileURLToPath(configUrl);
const configFile = fs.readFileSync(configPath, "utf8");
const config = parse(configFile);
// 2. 環境変数による上書き（優先）
const host = process.env.API_HOST || config.api?.host || "localhost";
const port = process.env.API_PORT || config.api?.port || 8080;
const baseUrl = `http://${host}:${port}`;
// openapi-fetch クライアントの初期化 [3, 4]
const client = createClient({
    baseUrl: baseUrl
});
export default client;
//# sourceMappingURL=api.js.map