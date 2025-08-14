export type Book = {
  id: number
  title: string
  author: string
  publisher: string
  year: number | undefined
  image_url: string
}

export const emptyBook: Book = {
  id: 0,
  title: "",
  author: "",
  publisher: "",
  year: undefined,
  image_url: "",
}
