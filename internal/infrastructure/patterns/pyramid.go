package patterns

import "strings"

type PyramidPattern struct{}

func (PyramidPattern) Generate(size, down, right int) string {
	if size <= 0 || down <= 0 || right <= 0 {
		return ""
	}

	var result strings.Builder

	for d := 0; d < down; d++ {
		for row := 1; row <= size; row++ {
			for col := 0; col < right; col++ {
				// Spasi di sebelah kiri setiap piramida
				spaces := size - row
				result.WriteString(strings.Repeat(" ", spaces))

				// Jumlah bintang per baris: 1, 3, 5, 7, ...
				stars := 2*row - 1
				result.WriteString(strings.Repeat("*", stars))

				// Spasi di sebelah kanan piramida agar simetris
				result.WriteString(strings.Repeat(" ", spaces))

				// Jarak pemisah antar-piramida jika lebih dari 1 kolom
				if col < right-1 {
					result.WriteString("   ") // Ubah/atur jumlah spasi pemisah di sini jika ingin lebih renggang/rapat
				}
			}

			// Pindah baris untuk setiap row, kecuali baris terakhir dari keseluruhan pola
			if !(d == down-1 && row == size) {
				result.WriteString("\n")
			}
		}
	}

	return result.String()
}
