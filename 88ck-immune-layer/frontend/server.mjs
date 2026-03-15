import { createServer } from "node:http";
import { readFile } from "node:fs/promises";
import { extname, join } from "node:path";

const root = "/app/dist";
const mime = new Map([
  [".html", "text/html; charset=utf-8"],
  [".js", "text/javascript; charset=utf-8"],
  [".css", "text/css; charset=utf-8"],
  [".svg", "image/svg+xml"],
  [".json", "application/json; charset=utf-8"]
]);

createServer(async (req, res) => {
  try {
    const path = req.url === "/" ? "/index.html" : req.url || "/index.html";
    const file = await readFile(join(root, path));
    res.writeHead(200, { "Content-Type": mime.get(extname(path)) || "application/octet-stream" });
    res.end(file);
  } catch {
    const file = await readFile(join(root, "index.html"));
    res.writeHead(200, { "Content-Type": "text/html; charset=utf-8" });
    res.end(file);
  }
}).listen(4173, "0.0.0.0");
