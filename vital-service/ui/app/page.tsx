"use client";

import type React from "react";

import { useState, useEffect, useRef } from "react";
import { io, type Socket } from "socket.io-client";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Card, CardContent, CardFooter, CardHeader, CardTitle } from "@/components/ui/card";
import { Avatar, AvatarFallback } from "@/components/ui/avatar";
import { ScrollArea } from "@/components/ui/scroll-area";
import { Separator } from "@/components/ui/separator";
import { MessageSquare, Send } from "lucide-react";

// Message type definition
interface Message {
  id: string;
  user: string;
  text: string;
  timestamp: Date;
}

export default function ChatPage() {
  const [socket, setSocket] = useState<Socket | null>(null);
  const [messages, setMessages] = useState<Message[]>([]);
  const [inputMessage, setInputMessage] = useState("");
  const [username, setUsername] = useState("");
  const [isConnected, setIsConnected] = useState(false);
  const scrollAreaRef = useRef<HTMLDivElement>(null);

  // Generate a random username when the component mounts
  useEffect(() => {
    const randomUsername = `User${Math.floor(Math.random() * 10000)}`;
    setUsername(randomUsername);
  }, []);

  // Initialize Socket.IO connection
  useEffect(() => {
    // Connect to the Socket.IO server
    const newSocket = io("http://localhost:3001",{
      path: "/socket.io",
      transports: ["websocket"],
      withCredentials: false,
    });

    // Set up event listeners
    newSocket.on("connect", () => {
      console.log("Connected to Socket.IO server");
      setIsConnected(true);
    });

    newSocket.on("disconnect", () => {
      console.log("Disconnected from Socket.IO server");
      setIsConnected(false);
    });

    newSocket.on("message", (message: Message) => {
      setMessages((prevMessages) => [...prevMessages, message]);
    });

    // Save the socket instance
    setSocket(newSocket);

    // Clean up the socket connection when the component unmounts
    return () => {
      newSocket.disconnect();
    };
  }, []);

  // Auto-scroll to the bottom when new messages arrive
  useEffect(() => {
    if (scrollAreaRef.current) {
      scrollAreaRef.current.scrollTop = scrollAreaRef.current.scrollHeight;
    }
  }, [messages]);

  // Send a message
  const sendMessage = () => {
    if (inputMessage.trim() && socket) {
      const newMessage: Message = {
        id: Date.now().toString(),
        user: username,
        text: inputMessage,
        timestamp: new Date(),
      };

      // Emit the message to the server
      socket.emit("message", newMessage);

      // Add the message to the local state
      setMessages((prevMessages) => [...prevMessages, newMessage]);

      // Clear the input field
      setInputMessage("");
    }
  };

  // Handle Enter key press
  const handleKeyPress = (e: React.KeyboardEvent) => {
    if (e.key === "Enter") {
      sendMessage();
    }
  };

  // Format timestamp
  const formatTime = (date: Date) => {
    return new Date(date).toLocaleTimeString([], { hour: "2-digit", minute: "2-digit" });
  };

  return (
    <div className="flex items-center justify-center min-h-screen bg-gray-100 p-4">
      <Card className="w-full max-w-md shadow-lg">
        <CardHeader className="bg-white border-b">
          <div className="flex items-center space-x-2">
            <MessageSquare className="h-5 w-5 text-emerald-500" />
            <CardTitle>Real-Time Chat</CardTitle>
          </div>
          <div className="text-sm text-muted-foreground">
            {isConnected ? (
              <span className="text-emerald-500">Connected as {username}</span>
            ) : (
              <span className="text-red-500">Disconnected</span>
            )}
          </div>
        </CardHeader>
        <CardContent className="p-0">
          <ScrollArea className="h-[400px] p-4" ref={scrollAreaRef}>
            {messages.length === 0 ? (
              <div className="flex items-center justify-center h-full text-muted-foreground">No messages yet. Start the conversation!</div>
            ) : (
              <div className="space-y-4">
                {messages.map((message) => (
                  <div key={message.id} className={`flex ${message.user === username ? "justify-end" : "justify-start"}`}>
                    <div className={`flex max-w-[80%] ${message.user === username ? "flex-row-reverse" : "flex-row"}`}>
                      <Avatar className={`h-8 w-8 ${message.user === username ? "ml-2" : "mr-2"}`}>
                        <AvatarFallback className="bg-emerald-100 text-emerald-800">
                          {message.user.substring(0, 2).toUpperCase()}
                        </AvatarFallback>
                      </Avatar>
                      <div>
                        <div
                          className={`rounded-lg p-3 ${
                            message.user === username ? "bg-emerald-500 text-white" : "bg-gray-200 text-gray-800"
                          }`}
                        >
                          {message.text}
                        </div>
                        <div className={`text-xs mt-1 text-gray-500 ${message.user === username ? "text-right" : "text-left"}`}>
                          {message.user === username ? "You" : message.user} • {formatTime(message.timestamp)}
                        </div>
                      </div>
                    </div>
                  </div>
                ))}
              </div>
            )}
          </ScrollArea>
        </CardContent>
        <Separator />
        <CardFooter className="p-3">
          <div className="flex w-full items-center space-x-2">
            <Input
              type="text"
              placeholder="Type your message..."
              value={inputMessage}
              onChange={(e) => setInputMessage(e.target.value)}
              onKeyUp={handleKeyPress}
              disabled={!isConnected}
              className="flex-1"
            />
            <Button
              onClick={sendMessage}
              disabled={!isConnected || !inputMessage.trim()}
              size="icon"
              className="bg-emerald-500 hover:bg-emerald-600"
            >
              <Send className="h-4 w-4" />
            </Button>
          </div>
        </CardFooter>
      </Card>
    </div>
  );
}
