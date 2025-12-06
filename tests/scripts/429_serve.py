from http.server import BaseHTTPRequestHandler, HTTPServer

class RateLimitHandler(BaseHTTPRequestHandler):
    def do_GET(self):
        self.send_response(429, "Too Many Requests")
        self.send_header("Content-Type", "text/plain")
        self.end_headers()
        self.wfile.write(b"429 Too Many Requests")

    # Optional: handle POST, etc.
    def do_POST(self):
        self.do_GET()

def run(server_class=HTTPServer, handler_class=RateLimitHandler, port=8080):
    server = server_class(("", port), handler_class)
    print(f"Serving 429 server on port {port}")
    server.serve_forever()

if __name__ == "__main__":
    run()
