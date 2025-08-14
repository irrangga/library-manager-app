import { Book, emptyBook } from "@/lib/definitions/book"
import { useEffect } from "react"
import { useForm } from "react-hook-form"
import { Button } from "../ui/button"
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "../ui/dialog"
import { Input } from "../ui/input"
import { Label } from "../ui/label"

type Props = {
  isDialogOpen: boolean
  setIsDialogOpen: (open: boolean) => void
  handleSubmit: (book: Book) => Promise<string | null | undefined>
  initialBook?: Book
}

export default function BookFormDialog({
  isDialogOpen,
  setIsDialogOpen,
  handleSubmit,
  initialBook,
}: Props) {
  const form = useForm({
    defaultValues: initialBook || emptyBook,
  })
  const { errors } = form.formState

  useEffect(() => {
    form.reset(initialBook)
  }, [initialBook])

  const onSubmit = async (book: Book) => {
    const error = await handleSubmit(book)
    if (!error) {
      setIsDialogOpen(false)
    }
  }

  return (
    <Dialog open={isDialogOpen} onOpenChange={setIsDialogOpen}>
      <DialogTrigger asChild>
        <Button
          variant="outline"
          onClick={() => {
            setIsDialogOpen(true)
            form.reset()
          }}
        >
          {initialBook ? "Edit book" : "+ Add book"}
        </Button>
      </DialogTrigger>
      <DialogContent className="sm:max-w-[425px]">
        <DialogHeader>
          <DialogTitle>
            {initialBook ? "Edit book" : "Add new book"}
          </DialogTitle>
          <DialogDescription>
            {initialBook
              ? "Edit the details of this book here."
              : "Enter the details of the book you want to add here."}
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
              {errors.title && (
                <span className="text-sm text-red-600">Title is required</span>
              )}
            </div>
            <div className="grid gap-3">
              <Label htmlFor="author">Author</Label>
              <Input
                id="author"
                {...form.register("author", { required: true })}
              />
              {errors.author && (
                <span className="text-sm text-red-600">Author is required</span>
              )}
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
            <div className="grid gap-3">
              <Label htmlFor="image-url">Image URL</Label>
              <Input id="image-url" {...form.register("image_url")} />
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
  )
}
