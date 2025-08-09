import { Book, emptyBook } from "@/lib/definitions/book"
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
}

export default function BookFormDialog({
  isDialogOpen,
  setIsDialogOpen,
  handleSubmit,
}: Props) {
  const form = useForm({
    defaultValues: emptyBook,
  })

  const onSubmit = async (book: Book) => {
    const error = await handleSubmit(book)
    if (!error) {
      setIsDialogOpen(false)
      form.reset()
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
  )
}
