import { getBooks } from "@/lib/api/books"
import BookList from "./book-list"

export default async function Page() {
  const { data } = await getBooks()
  return <BookList initialBooks={data} />
}
