const { createServer } = require("http")
const { Server } = require("socket.io")
const next = require("next")

const dev = process.env.NODE_ENV !== "production"
const app = next({ dev })
const handle = app.getRequestHandler()

app.prepare().then(() => {
  const server = createServer((req, res) => {
    handle(req, res)
  })

  // Create a Socket.IO server
  const io = new Server(server, {
    path: "/socket.io",
    cors: {
      origin: "http://localhost:3001",
      methods: ["GET", "POST"],
    },
    withCredentials: false,
  });
  

  // Set up Socket.IO event handlers
  io.on("connection", (socket) => {
    console.log("Client connected:", socket.id)

    // Handle incoming messages
    socket.on("message", (message) => {
      // Broadcast the message to all connected clients
      io.emit("message", message)
    })

    // Handle disconnection
    socket.on("disconnect", () => {
      console.log("Client disconnected:", socket.id)
    })
  })

  // Start the server
  const PORT = process.env.PORT || 3001
  server.listen(PORT, () => {
    console.log(`> Server listening on port ${PORT}`)
  })
})
