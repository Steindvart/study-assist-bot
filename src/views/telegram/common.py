from aiogram import Bot
from aiogram.types import BotCommand, InlineKeyboardMarkup, InlineKeyboardButton
from aiogram.utils.keyboard import InlineKeyboardBuilder

import utils

from config import config

res = config.resources

# ---------------------------------------------


def get_text_start_message() -> str:
  text = (f'{res['hello']}\n\n'
          f'{"\n".join(utils.get_formated_main_commands_desc(res['main_commands']))}'
  )

  return text


async def set_main_menu(bot: Bot):
  main_menu_commands = [
    BotCommand(
      command=command,
      description=description
    ) for command, description in res['main_commands'].items() if command != '_prev'
  ]

  await bot.set_my_commands(main_menu_commands)


def get_keyboard_start_interact() -> InlineKeyboardMarkup:
  keyboard = InlineKeyboardBuilder([
    [InlineKeyboardButton(text=res['sections']['list_sections'], callback_data='start:interact')]
  ])

  return keyboard.as_markup()

def get_body_start_interact() -> (str, InlineKeyboardMarkup):
  text = res['start_interact']
  keyboard = get_keyboard_start_interact()

  return (text, keyboard)