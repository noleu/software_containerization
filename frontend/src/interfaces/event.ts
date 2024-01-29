export interface ReturnEvent {
  id: number
  title: string
  date_time: Date
  location: string
  description: string
}

export interface CreateEvent {
  title: string
  date_time: Date
  location: string
  description: string
}
