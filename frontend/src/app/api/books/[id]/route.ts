import { deleteBookById, updateBookById } from "@/lib/api/books"
import { Book } from "@/lib/definitions/book"
import { NextResponse } from "next/server"

export async function PUT(
  request: Request,
  { params }: { params: Promise<{ id: string }> },
) {
  const { id } = await params
  const body: Omit<Book, "id"> = await request.json()

  const { data, error } = await updateBookById(id, body)
  return NextResponse.json({ data, error })
}

export async function DELETE(
  _request: Request,
  { params }: { params: Promise<{ id: string }> },
) {
  const { id } = await params
  const { error } = await deleteBookById(id)
  return NextResponse.json({ error })
}
