# Task 3 — AI blind test generation

> Homework — Task 3 (Lesson 6: File I/O, JSON and Testing).
> This task is **not auto-graded** by CI — it's reviewed by your mentor.
> The workflow only checks that you actually filled this file in
> (see the "Task 3 — AI report present" step in the Actions log).

## Instructions

1. Pick one function you wrote for Task 1 or Task 2 (e.g. `LoadTodos` or
   `ValidateEmail`).
2. Give an AI assistant **only the function signature** — no explanation
   of the logic, no existing test file. Ask it to generate a full
   table-driven test suite for that signature.
3. Run the AI-generated tests against your implementation.
4. Fill in the sections below.

---

## Function under test

func LoadTodos(path string) ([]Todo, error)

## Prompt you gave the AI

generate a full table-driven test suite for that signature. No questions.

## Edge cases the AI found that you had missed

Після того, як я спитав у ШІ, чи є в нього питання стосовно граничних випадків, які я можливо пропустив, він почав задавати питання стосовно мого проекту, тести для якого він написав, генеруючи припущення. 

По пунктах: 1) Шлях вказує на директорію, а не на файл:, 2) Відсутність прав доступу (Permission Denied):, 3) Відмінність між null, порожнім масивом [] та порожнім файлом:, 4) Невідповідність типів структури (JSON Object замість Array):, 5) Частково некоректні дані у структурі (Partial Fields):, 6) Символічні посилання (Symlinks) та біті шляхи:, 7) Файли з Byte Order Mark (UTF-8 BOM)

## Edge cases you had that the AI missed

Так як у ШІ не було контексту, він його додумав самостійно. Відповідно можна стверджувати, що він або відхилився від концепції завдання, або створив свою, хоч й дотичну.

## Cases where the AI's expected output was wrong

ШІ не використав бібліотеку time, хоча в оригінальних тестах вона задіяна. Проте натомість він застосував reflect, аргументуючи це тим, що в Go зрізи та масиви ([]Todo) еможливо порівняти напряму через стандартний оператор рівності (==). Використав reflect.DeepEqual. 

А Time потрібен для роботи з часом - затримки, таймаути і тд. "Функція LoadTodos(path string) виконує звичайну синхронну операцію зчитування файлу та розпаршування JSON. Тут немає асинхронного коду, контекстів з обмеженням за часом чи часових затримок". 

## What you'd change about your own test-writing process after this

Буду давати ще більше контексту ШІ, буду прописувати, щоб він задавав уточнюючі питання. Буду використовувати також бібліотеку reflect.