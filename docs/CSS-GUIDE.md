# Руководство по стилям Taskify

## Общая структура CSS
### Основные глобальные стили
```css
* {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
}

body {
    position: absolute;
    top: 0;
    right: 0;
    bottom: 0;
    left: 0;
    background-color: #E2E1DD;
}
```
### Цветовая палитра
**Основные цвета проекта:**

| Цвет               | HEX      | Использование                      |
|--------------------|----------|------------------------------------|
| Основной фон       | `#E2E1DD` | Фон страницы                       |
| Акцентный синий    | `#374DCB` | Кнопки регистрации                 |
| Серый фон          | `#D3D0D0` | Блоки контента                     |
| Текст              | `#545454` | Основной цвет текста               |
| Ховер-эффект       | `#A09E9E` | Состояние при наведении            | 

### Типографика
**Шрифты**

```css
font-family: "Bellota Text";  /* Основной текст */
font-family: "Italian";       /* Логотип */ 
```
**Размеры шрифтов**  

Заголовок логотипа: 40px

Основной текст: 20-24px

Кнопки: 23-36px

### Компоненты
**Шапка (Header)**
```css
header {
    background-color: #E2E1DD;
    padding: 20px 0;
    border-bottom: 1px solid #1077EC;
}

.navbar {
    display: flex;
    justify-content: space-between;
    align-items: center;
}
```
**Основной контент**
```css
.main-content {
    display: flex;
    justify-content: space-between;
    align-items: flex-start;
    position: relative;
}

.text-content {
    background-color: #D3D0D0;
    border-radius: 15px;
    box-shadow: 0px 5px 10px 0px rgba(0, 0, 0, 0.5);
}
```
**Кнопки**
```css 
    .cta-button {
    background-color: #C2C0C0;
    color: #6A7DEA;
    padding: 12px 25px;
    border-radius: 15px;
    font-size: 36px;
    transition: background-color 0.3s;
}

.cta-button:hover {
    background-color: #939292;
}
```

**Секция преимуществ**
```css
.features {
    background-color: #B0B2D3;
    border-radius: 10px;
    box-shadow: 0px 5px 10px 0px rgba(0, 0, 0, 0.5);
}

.features-list li {
    display: flex;
    align-items: center;
    padding: 15px 20px;
    border-bottom: 1px solid rgba(0, 0, 0, 0.1);
}
```
## Особенности реализации

**Позиционирование:**

Абсолютное позиционирование для точного размещения элементов  

z-index для управления слоями  

**Тени:**

Все блоки используют тени для глубины: box-shadow: 0px 5px 10px 0px rgba(0, 0, 0, 0.5)  

**Анимации:**

Плавные переходы для интерактивных элементов: transition: background-color 0.3s  