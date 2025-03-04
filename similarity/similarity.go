package similarity

import (
	"math"
	"strings"
)

// CalculateSimilarity вычисляет процент сходства между двумя текстами с использованием косинусного сходства (cosine similarity)
func CalculateSimilarity(text1, text2 string) float64 {
	// Преобразуем тексты в нижний регистр и разбиваем на слова
	words1 := strings.Fields(strings.ToLower(text1))
	words2 := strings.Fields(strings.ToLower(text2))

	// Создаем словари для подсчета частоты слов
	freq1 := make(map[string]int)
	freq2 := make(map[string]int)

	// Заполняем частотные словари
	for _, word := range words1 {
		freq1[word]++
	}
	for _, word := range words2 {
		freq2[word]++
	}

	// Вычисляем скалярное произведение
	dotProduct := 0.0
	for word, count1 := range freq1 {
		if count2, exists := freq2[word]; exists {
			dotProduct += float64(count1 * count2)
		}
	}

	// Вычисляем длины векторов
	magnitude1 := 0.0
	for _, count := range freq1 {
		magnitude1 += float64(count * count)
	}
	magnitude1 = math.Sqrt(magnitude1)

	magnitude2 := 0.0
	for _, count := range freq2 {
		magnitude2 += float64(count * count)
	}
	magnitude2 = math.Sqrt(magnitude2)

	// Если один из векторов нулевой, возвращаем 0%
	if magnitude1 == 0 || magnitude2 == 0 {
		return 0
	}

	// Вычисляем косинусное сходство и переводим в проценты
	similarity := (dotProduct / (magnitude1 * magnitude2)) * 100
	return math.Round(similarity*100) / 100 // округляем до 2 знаков после запятой
}
