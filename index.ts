import { Bot } from 'grammy'

const { BOT_TOKEN, BOT_WEBHOOK } = process.env

const bot = new Bot(BOT_TOKEN)

console.log('Setting Webhook')
let res = await bot.api.setWebhook(BOT_WEBHOOK)
console.log(res)