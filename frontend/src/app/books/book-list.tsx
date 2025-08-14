"use client"

import BookFormDialog from "@/components/shared/book-form-dialog"
import { Button } from "@/components/ui/button"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { useBookContext } from "@/lib/context/book"
import { Book } from "@/lib/definitions/book"
import Link from "next/link"
import { useEffect, useState } from "react"
import { toast } from "sonner"

export default function BookList({ initialBooks }: { initialBooks: Book[] }) {
  const { books, setInitialBooks, addBook, deleteBookById } = useBookContext()
  const [isDialogOpen, setIsDialogOpen] = useState(false)

  useEffect(() => {
    setInitialBooks(initialBooks)
  }, [initialBooks])

  const handleSubmit = async (book: Book) => {
    const { error } = await addBook(book)
    if (!error) {
      toast.success("Book added successfully")
    } else {
      toast.error("Failed to add a book")
    }
    return error
  }

  const handleDelete = async (id: number) => {
    const { error } = await deleteBookById(id)
    if (!error) {
      toast.success("Book deleted successfully")
    } else {
      toast.error("Failed to delete a book")
    }
  }

  return (
    <div className="p-4 space-y-4">
      <div className="flex justify-end">
        <BookFormDialog
          isDialogOpen={isDialogOpen}
          setIsDialogOpen={setIsDialogOpen}
          handleSubmit={handleSubmit}
        />
      </div>
      <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
        {books.map((book) => (
          <Card
            key={book.id}
            asChild
            className="group relative cursor-pointer gap-4"
          >
            <Link href={`/books/${book.id}`}>
              <CardHeader className="flex flex-col items-center gap-4">
                <img
                  src={book.image_url}
                  alt={book.title}
                  className="h-72 w-auto object-cover"
                />
                <CardTitle className="self-start">
                  {book.title || "-"}
                </CardTitle>
              </CardHeader>
              <CardContent className="text-sm space-y-1">
                <p>{book.author || "-"}</p>
                <p>{book.publisher || "-"}</p>
                <p>{book.year || "-"}</p>
              </CardContent>
              <div className="opacity-0 group-hover:opacity-100 transition-opacity flex justify-center gap-4">
                <Link href={`/books/${book.id}?edit=true`}>
                  <Button size="sm">Edit</Button>
                </Link>
                <Button
                  size="sm"
                  variant="destructive"
                  onClick={(e) => {
                    e.preventDefault()
                    handleDelete(book.id)
                  }}
                >
                  Delete
                </Button>
              </div>
            </Link>
          </Card>
        ))}
      </div>
    </div>
  )
}
