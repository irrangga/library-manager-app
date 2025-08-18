"use client"

import { createContext, ReactNode, useContext, useState } from "react"
import { Book } from "../definitions/book"

type BookContextType = {
  books: Book[]
  setBooks: (initialBooks: Book[]) => void
  addBook: (book: Book) => Promise<{ data?: Book; error?: string | null }>
  deleteBookById: (id: number) => Promise<{ error?: string | null }>
}

const BookContext = createContext<BookContextType | undefined>(undefined)

export function BookProvider({ children }: { children: ReactNode }) {
  const [books, setBooks] = useState<Book[]>([])

  const addBook = async (book: Book) => {
    const { id, ...body } = book
    void id

    const response = await fetch(`/api/books`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(body),
    })

    const { data, error } = await response.json()
    if (error) {
      return { error }
    }

    setBooks((prev) => [...prev, data])
    return { data }
  }

  const deleteBookById = async (id: number) => {
    const response = await fetch(`/api/books/${id}`, {
      method: "DELETE",
    })
    const { error } = await response.json()
    if (error) {
      return { error }
    }
    setBooks((prev) => prev.filter((book) => book.id !== id))
    return { error }
  }

  return (
    <BookContext.Provider
      value={{
        books,
        setBooks,
        addBook,
        deleteBookById,
      }}
    >
      {children}
    </BookContext.Provider>
  )
}

export function useBookContext() {
  const context = useContext(BookContext)
  if (!context)
    throw new Error("useBookContext must be used within a BookProvider")
  return context
}
