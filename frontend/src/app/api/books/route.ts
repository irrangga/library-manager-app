import { addBook } from "@/lib/api/books"
import { Book } from "@/lib/definitions/book"
import { NextResponse } from "next/server"

export async function POST(request: Request) {
  const body: Omit<Book, "id"> = await request.json()

  const { data, error } = await addBook(body)
  return NextResponse.json({ data, error })
}
