# spam-bot
Spam bot for albion (any game) guilds

# 1. Сообщение
Пройдите в папку spam_message. Откройте файл message.txt, введите туда сообщение, которое бот будет использовать. Сохраните и закройте файл.

# 2. Запуск
Пройдите в папку main и запустите файл main.exe, затем перейдите в окно, в котором вы хотите спамить своим сообщением.

# 3. Настройки
Бот будет спамить сообщение в активном окне раз в 30 секунд. Чтобы поменять это время, поменяйте константу в main.go "timeSleep" на нужное кол-во секунд. Затем надо будет сделать go build, чтобы изменение подтянулось в exe'шник.

Если вы хотите поменять сообщение в процессе работы, надо будет закрыть бота, поменять сообщение и открыть снова. 
