import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { getBookById } from "@/lib/api/books"
import { Label } from "@radix-ui/react-label"

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
          <Label className="font-semibold">Author:</Label>
          <p>{data.author || "-"}</p>
          <Label className="font-semibold">Publisher:</Label>
          <p>{data.publisher || "-"}</p>
          <Label className="font-semibold">Year:</Label>
          <p>{data.year || "-"}</p>
        </CardContent>
      </Card>
    </div>
  )
}
