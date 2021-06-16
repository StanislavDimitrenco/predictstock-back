package messages

func Tutorial() string {
	return "Для использования бота напишите тикер *акции* сообщением\\. \n\n*Например:* \nЕсли хотите увидеть прогноз по акции *Тесла*\\, нужно ввести *TSLA*\\." +
		"\n\n*Consensus \\(консенсус\\)* – одна из пяти возможных рекомендаций:\n\n" +
		"STRONG BUY – уверенно покупать\n" +
		"BUY – покупать\n" +
		"HOLD – держать\n" +
		"SELL – продавать\n" +
		"STRONG SELL – уверенно продавать\n\n" +
		"*Консенсус* – это общее мнение наиболее авторитетных по данной акции аналитиков на период от одного до трех месяцев\\.\n\n" +
		"*Status \\(статус\\)* может быть UPGRADED \\(повышен\\) или DOWNGRADED \\(понижен\\)\\.\nОн показывает\\, сколько дней назад изменился консенсус по данной бумаге и в какую сторону\\. \n\n" +
		"*Например:* \n\n" +
		"при запросе тикера АА выдается консенсус BUY\\, статус Downgraded 4 days\\. Это означает\\, что консенсус по акции был «уверенно покупать» \\(STRONG BUY\\)\\, а 4 дня назад он понизился до «покупать» \\(BUY\\)\\."

}
