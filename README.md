# Software_analyzing_traffic_situation
Приложение предназначено для вывода рекомендаций и предупреждений при управлении транспортным средством. Может применяться водителями и обучающимися, занимающимися управлением или обучением управлению транспортного средством. 

# Установка
Для развертывания приложения необходимо:
  1) Прописать путь в Golang сервере (sppr/main.go) к example.jar файлу - это главный "обработчик программы"
  2) Для запуска сервера необходимо установить Golang с официального сайта
  3) Запустить Golang-сервер командой можно написав команду «go run main.go» в консоль или среду разработки находясь в папке, в которой лежит файл main.go;
  5) Необходимо разахривировать SPPR_EXE.rar - в нём находится unity-клиент, необходимый для визуального отображения программы
  6) Запустить SPPR.exe

# Горячие клавиши
  * Space - Смена дорожного знака
  * Enter - Обработка знака, анализ ситуации и вывод информации на экран
  * Enter - Если обработка знака только что была, то эта клавиша приведет к возвращению к виду сверху   
