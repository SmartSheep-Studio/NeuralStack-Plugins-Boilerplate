import { createApp } from "vue"
import { createPinia } from "pinia"
import { Quasar } from 'quasar'
import quasarIconSet from 'quasar/icon-set/mdi-v7'

import '@quasar/extras/roboto-font/roboto-font.css'
import '@quasar/extras/mdi-v7/mdi-v7.css'

import 'quasar/src/css/index.sass'

import main from "@/app.vue"

const app = createApp(main)

app.use(createPinia())
app.use(Quasar, {
  plugins: {},
  iconSet: quasarIconSet,
})

app.mount("#plugins-example")
