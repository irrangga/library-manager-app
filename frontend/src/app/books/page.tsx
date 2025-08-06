import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { getBooks } from "@/lib/api/books"

export default async function Page() {
  const { data } = await getBooks()

  return (
    <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4 p-4">
      {data.map((book) => (
        <Card key={book.id}>
          <CardHeader>
            <CardTitle>{book.title}</CardTitle>
          </CardHeader>
          <CardContent className="text-sm space-y-1">
            <p>{book.author}</p>
            <p>{book.publisher}</p>
            <p>{book.year}</p>
          </CardContent>
        </Card>
      ))}
    </div>
  )
}
