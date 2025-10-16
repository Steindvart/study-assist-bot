package telegram

// Predefined messages and responses for the Telegram bot.
var (
    WelcomeMessage = "Добро пожаловать в наш бот! Как я могу помочь вам сегодня?"
    HelpMessage    = "Вы можете задать вопросы по разделам и темам, а также пройти тесты."
    ErrorMessage   = "Произошла ошибка. Пожалуйста, попробуйте еще раз."
    SectionNotFoundMessage = "Раздел не найден. Пожалуйста, выберите другой раздел."
    TestCompletedMessage = "Тест завершен! Ваш результат: %d из %d."
    ExplanationMessage = "Объяснение: %s"
    LocalizationErrorMessage = "Ошибка локализации. Пожалуйста, проверьте настройки языка."
)