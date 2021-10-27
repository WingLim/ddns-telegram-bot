import { VercelRequest, VercelResponse } from '@vercel/node'
import { InlineKeyboard, webhookCallback } from 'grammy'
import { bot } from '../index'

const { VERCEL_URL } = process.env

bot.command('start', async (ctx) => {
    await ctx.reply('Welcome to use DDNS Bot')
})

bot.command('gethook', async (ctx) => {
    const chanId = ctx.message.chat.id
    const hookUrl = `https://${VERCEL_URL}/api/hook/${chanId}`
    const links = new InlineKeyboard()
        .url('Usage', 'https://github.com/WingLim/ddns-telegram-bot/blob/main/README.md')
    await ctx.reply(`Your Webhook URL:\n ${hookUrl}`, {
        reply_markup: links
    })
})

export default async (req: VercelRequest, res: VercelResponse) => {
    webhookCallback(bot, 'http')(req, res)
}
