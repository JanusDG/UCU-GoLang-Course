type Retriable interface {
	GetMaxAttempts() uint
	Retry(action func() bool)
}

// Реалізуйте інтерфейс у двох варіантах DefaultRetriable, RetriableWithDelay
// В обох структурах Retry має викликати передану функцію не більше ніж повертає GetMaxAttempts.
// Якщо action() поверне true - зупиняєте виконання циклу
// RetriableWithDelay структура має мати поле DelayInSec. У випадку якщо action() повернув false - програма має почекати вказану кількість секунд (time.Sleep)
// Якщо action() поверне false - вивести в консоль повідомлення в якому вказати номер неуспішної спроби
// приклад використання
// ret := RetriableWithDelay{
// 	DelayInSec: 1,
// }
// ret.Retry(func() bool {
// 	return false
// })