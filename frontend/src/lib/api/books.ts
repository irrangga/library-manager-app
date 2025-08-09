import { env } from "@/lib/config/env"
import { Book, emptyBook } from "@/lib/definitions/book"
import axios from "axios"

export async function getBooks(): Promise<{
  data: Book[]
  error: Error | null
}> {
  try {
    const response = await axios.get<{ data: Book[] }>(`${env.API_URL}/books`)
    return { data: response.data.data, error: null }
  } catch (error: any) {
    return {
      data: [],
      error: error.response?.data?.error || "Failed to fetch books",
    }
  }
}

export async function addBook(
  book: Omit<Book, "id">,
): Promise<{ data: Book; error: Error | null }> {
  try {
    const response = await axios.post<{ data: Book }>(
      `${env.API_URL}/books`,
      book,
    )
    return { data: response.data.data, error: null }
  } catch (error: any) {
    return {
      data: emptyBook,
      error: error.response?.data?.error || "Failed to add a book",
    }
  }
}

export async function getBookById(
  id: string,
): Promise<{ data: Book; error: Error | null }> {
  try {
    const response = await axios.get<{ data: Book }>(
      `${env.API_URL}/books/${id}`,
    )
    return { data: response.data.data, error: null }
  } catch (error: any) {
    return {
      data: emptyBook,
      error: error.response?.data?.error || "Failed to fetch a book",
    }
  }
}

export async function updateBookById(
  id: string,
  book: Omit<Book, "id">,
): Promise<{ data: Book; error: Error | null }> {
  try {
    const response = await axios.put<{ data: Book }>(
      `${env.API_URL}/books/${id}`,
      book,
    )
    return { data: response.data.data, error: null }
  } catch (error: any) {
    return {
      data: emptyBook,
      error: error.response?.data?.error || "Failed to update a book",
    }
  }
}
