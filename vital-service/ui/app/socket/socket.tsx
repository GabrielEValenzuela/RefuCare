import type { Server as NetServer } from "http"
import { Server as SocketIOServer } from "socket.io"
import { NextResponse } from "next/server"

// Store the Socket.IO server instance
let io: SocketIOServer | null = null

export async function GET() {
  // Return a simple response for health check
  return NextResponse.json({ status: "Socket.IO server is running" })
}

// This function initializes the Socket.IO server
function initSocketServer(req: Request) {
  if (!io) {
    // Get the server instance from the environment
    const res = new NextResponse()
    const server = (res as any).socket?.server as NetServer

    if (server) {
      // Create a new Socket.IO server
      io = new SocketIOServer(server, {
        cors: {
          origin: "*",
          methods: ["GET", "POST"],
        },
      })

      // Set up event handlers
      io.on("connection", (socket) => {
        console.log("Client connected:", socket.id)

        // Handle incoming messages
        socket.on("message", (message) => {
          // Broadcast the message to all connected clients
          io?.emit("message", message)
        })

        // Handle disconnection
        socket.on("disconnect", () => {
          console.log("Client disconnected:", socket.id)
        })
      })

      console.log("Socket.IO server initialized")
    }
  }

  return NextResponse.json({ status: "Socket.IO server initialized" })
}

export const POST = initSocketServer
