import { createClient } from '@supabase/supabase-js'

const supabaseAnonKey = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYW5vbiIsImlhdCI6MTYzNjU2ODc3NSwiZXhwIjoxOTUyMTQ0Nzc1fQ.CwUnJhI0_fMS5O77DxLKkhHf9ULziq5eCvHwEY9z5RI"
const supabaseUrl = "https://npthxfsjlkgtythiccdg.supabase.co"

export const supabase = createClient(supabaseUrl, supabaseAnonKey)
