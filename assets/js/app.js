import Alpine from "alpinejs";
import dayjs from "dayjs";
import { createIcons, Search, SquarePen, Users } from "lucide";

window.Alpine = Alpine
window.dayjs = dayjs

createIcons({
  icons: {
    SquarePen,
    Search,
    Users
  }
})

Alpine.start()

