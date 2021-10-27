import { Bot } from 'grammy'

const { BOT_TOKEN, BOT_WEBHOOK } = process.env

export const bot = new Bot(BOT_TOKEN)

bot.api.setWebhook(BOT_WEBHOOK)