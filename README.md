# similarity

Вычисляет процент сходства между двумя текстами с использованием косинусного сходства (cosine similarity)

Пример использования:

    import "github.com/Nemnon/go-similarity/similarity"

    func main() {  
        text1 := "Я люблю программировать на Go"  
        text2 := "Мне нравится писать код на Go"  
        similarity := CalculateSimilarity(text1, text2) 
        fmt.Printf("Сходство текстов: %.2f%%\n", similarity)
    }
