import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { Label } from "@/components/ui/label"
import { getBookById } from "@/lib/api/books"

export default async function Page({
  params,
}: {
  params: Promise<{ id: string }>
}) {
  const { id } = await params
  const { data } = await getBookById(id)

  return (
    <div className="p-4">
      <Card>
        <CardHeader className="flex items-center justify-between">
          <CardTitle>{data.title || "-"}</CardTitle>
        </CardHeader>
        <CardContent className="space-y-4">
          <Label>Author:</Label>
          <p>{data.author || "-"}</p>
          <Label>Publisher:</Label>
          <p>{data.publisher || "-"}</p>
          <Label>Year:</Label>
          <p>{data.year || "-"}</p>
        </CardContent>
      </Card>
    </div>
  )
}
