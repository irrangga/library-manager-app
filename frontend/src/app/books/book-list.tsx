"use client"

import { Button } from "@/components/ui/button"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { useBookContext } from "@/lib/context/book"
import { Book, emptyBook } from "@/lib/definitions/book"
import { useEffect, useState } from "react"
import { useForm } from "react-hook-form"
import { toast } from "sonner"

export default function BookList({ initialBooks }: { initialBooks: Book[] }) {
  const { books, setInitialBooks, addBook } = useBookContext()
  const [isDialogOpen, setIsDialogOpen] = useState(false)

  useEffect(() => {
    setInitialBooks(initialBooks)
  }, [initialBooks])

  const form = useForm({
    defaultValues: emptyBook,
  })

  const onSubmit = async (book: Book) => {
    const { error } = await addBook(book)

    if (!error) {
      setIsDialogOpen(false)
      form.reset()

      toast.success("Book added successfully")
    } else {
      toast.error("Failed to add a book")
    }
  }

  return (
    <div className="p-4 space-y-4">
      <div className="flex justify-end">
        <Dialog open={isDialogOpen} onOpenChange={setIsDialogOpen}>
          <DialogTrigger asChild>
            <Button
              variant="outline"
              onClick={() => {
                setIsDialogOpen(true)
                form.reset()
              }}
            >
              + Add Book
            </Button>
          </DialogTrigger>
          <DialogContent className="sm:max-w-[425px]">
            <DialogHeader>
              <DialogTitle>Add new book</DialogTitle>
              <DialogDescription>
                Add a new book to your library here.
              </DialogDescription>
            </DialogHeader>
            <form onSubmit={form.handleSubmit(onSubmit)}>
              <div className="grid gap-4">
                <div className="grid gap-3">
                  <Label htmlFor="title">Title</Label>
                  <Input
                    id="title"
                    {...form.register("title", { required: true })}
                  />
                </div>
                <div className="grid gap-3">
                  <Label htmlFor="author">Author</Label>
                  <Input
                    id="author"
                    {...form.register("author", { required: true })}
                  />
                </div>
                <div className="grid gap-3">
                  <Label htmlFor="publisher">Publisher</Label>
                  <Input id="publisher" {...form.register("publisher")} />
                </div>
                <div className="grid gap-3">
                  <Label htmlFor="year">Year</Label>
                  <Input
                    id="year"
                    type="number"
                    {...form.register("year", { valueAsNumber: true })}
                  />
                </div>
                <DialogFooter>
                  <DialogClose asChild>
                    <Button variant="outline">Cancel</Button>
                  </DialogClose>
                  <Button type="submit">Save changes</Button>
                </DialogFooter>
              </div>
            </form>
          </DialogContent>
        </Dialog>
      </div>
      <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
        {books.map((book) => (
          <Card key={book.id}>
            <CardHeader>
              <CardTitle>{book.title || "-"}</CardTitle>
            </CardHeader>
            <CardContent className="text-sm space-y-1">
              <p>{book.author || "-"}</p>
              <p>{book.publisher || "-"}</p>
              <p>{book.year || "-"}</p>
            </CardContent>
          </Card>
        ))}
      </div>
    </div>
  )
}
