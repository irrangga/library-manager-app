export type Book = {
  id: number
  title: string
  author: string
  publisher: string
  year: number | undefined
}

export const emptyBook: Book = {
  id: 0,
  title: "",
  author: "",
  publisher: "",
  year: undefined,
}
