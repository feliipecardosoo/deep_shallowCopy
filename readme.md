
# Go: Shallow Copy, Deep Copy e Arquitetura com Interfaces (Exemplo da TV)

## 🧠 Conceitos abordados

- Shallow Copy (Cópia Rasa)
- Deep Copy (Cópia Profunda)
- Interface-Based Design (Programação orientada a interfaces)
- Boas práticas de arquitetura em Go
- Separação por arquivos (interface, lógica, entrada `main.go`)
- Princípio da Interface Segregation (ISP - SOLID)

---

## 📌 O que é Shallow Copy?

Uma **shallow copy** (ou cópia rasa) copia apenas **os endereços de memória** dos dados referenciados. Ou seja, se você modificar a cópia, o original também muda.

### Exemplo com `slice`:

```go
sliceOriginal := []int{1, 2, 3}
sliceCopia := sliceOriginal

sliceCopia[0] = 100
fmt.Println(sliceOriginal) // [100 2 3] → original também foi alterado
```

---

## 📌 O que é Deep Copy?

Uma **deep copy** (ou cópia profunda) copia **todos os dados de forma independente**, criando novas referências na memória.

### Exemplo com `slice`:

```go
sliceOriginal := []int{1, 2, 3}
sliceCopia := make([]int, len(sliceOriginal))
copy(sliceCopia, sliceOriginal)

sliceCopia[0] = 100
fmt.Println(sliceOriginal) // [1 2 3] → original permanece intacto
```
---

## 📚 Aprendizado

Esse projeto simples mostra como conceitos como **shallow copy, deep copy, interfaces e boas práticas de arquitetura** se aplicam até em algo tão cotidiano.
