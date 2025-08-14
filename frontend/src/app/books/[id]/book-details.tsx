"use client"

import BookFormDialog from "@/components/shared/book-form-dialog"
import { Button } from "@/components/ui/button"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { Book } from "@/lib/definitions/book"
import { ArrowLeft } from "lucide-react"
import Link from "next/link"
import { usePathname, useRouter, useSearchParams } from "next/navigation"
import { useEffect, useState } from "react"
import { toast } from "sonner"

export default function BookDetails({ initialBook }: { initialBook: Book }) {
  const searchParams = useSearchParams()
  const router = useRouter()
  const pathname = usePathname()
  const [book, setBook] = useState(initialBook)
  const [isDialogOpen, setIsDialogOpen] = useState(false)

  useEffect(() => {
    if (searchParams.get("edit")) {
      setIsDialogOpen(true)
      router.replace(pathname)
    }
  }, [])

  const handleSubmit = async (book: Book) => {
    const { id, ...body } = book

    const response = await fetch(`/api/books/${book.id}`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(body),
    })

    const { error } = await response.json()

    if (!error) {
      setBook(book)
      toast.success("Book updated successfully")
    } else {
      toast.error("Failed to update a book")
    }
    return error
  }

  return (
    <div className="p-4 space-y-4">
      <div className="flex justify-between">
        <Link href="/books">
          <Button variant="outline">
            <ArrowLeft className="h-4 w-4" />
            Back to list
          </Button>
        </Link>
        <BookFormDialog
          isDialogOpen={isDialogOpen}
          setIsDialogOpen={setIsDialogOpen}
          handleSubmit={handleSubmit}
          initialBook={book}
        />
      </div>
      <Card>
        <CardContent className="flex flex-row">
          <img
            src={book.image_url}
            alt={book.title}
            className="h-72 w-48 object-cover"
          />
          <div className="flex-1 space-y-4">
            <CardHeader>
              <CardTitle>{book.title || "-"}</CardTitle>
            </CardHeader>
            <CardContent className="space-y-4">
              <div>
                <p className="font-medium">Author:</p>
                <p>{book.author || "-"}</p>
              </div>
              <div>
                <p className="font-medium">Publisher:</p>
                <p>{book.publisher || "-"}</p>
              </div>
              <div>
                <p className="font-medium">Year:</p>
                <p>{book.year || "-"}</p>
              </div>
            </CardContent>
          </div>
        </CardContent>
      </Card>
    </div>
  )
}
