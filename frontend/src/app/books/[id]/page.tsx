import { getBookById } from "@/lib/api/books"
import BookDetails from "./book-details"

export default async function Page({
  params,
}: {
  params: Promise<{ id: string }>
}) {
  const { id } = await params
  const { data } = await getBookById(id)

  return <BookDetails initialBook={data} />
}
