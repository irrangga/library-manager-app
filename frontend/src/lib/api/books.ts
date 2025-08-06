import { env } from "@/lib/config/env"
import { Book } from "@/lib/definitions/book"
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
