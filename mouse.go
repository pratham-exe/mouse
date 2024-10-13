package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
)

func tokenize(input_str string) ([]string) {
    var res []string
    input := strings.Split(input_str, " ")
    for _, ch := range input {
        if ch != "" {
            res = append(res, ch)
        }
    }
    return res
}

func push_stack(stack []string, char string) ([]string) {
    stack = append(stack, char)
    return stack
}

func pop_stack(stack *[]string) (string) {
    last_ele := (*stack)[len(*stack) - 1]
    *stack = (*stack)[:len(*stack) - 1]
    return last_ele
}

func interactive()  {
    loop := true
    var stack []string
    alp_num := map[string]string{
		"A": "0", "B": "1", "C": "2", "D": "3", "E": "4", "F": "5", "G": "6", "H": "7", "I": "8", "J": "9", "K": "10", "L": "11", "M": "12", "N": "13", "O": "14", "P": "15", "Q": "16", "R": "17", "S": "18", "T": "19", "U": "20", "V": "21", "W": "22", "X": "23", "Y": "24", "Z": "25",
	}
    num_alp := map[string]string{
		"0":  "A", "1":  "B", "2":  "C", "3":  "D", "4":  "E", "5":  "F", "6":  "G", "7":  "H", "8":  "I", "9":  "J", "10": "K", "11": "L", "12": "M", "13": "N", "14": "O", "15": "P", "16": "Q", "17": "R", "18": "S", "19": "T", "20": "U", "21": "V", "22": "W", "23": "X", "24": "Y", "25": "Z",
	}
    var new_alp_num map[string]string = make(map[string]string)
    is_if := true
    is_str := false
    for loop {
        fmt.Print(">> ")
        scan := bufio.NewScanner(os.Stdin)
        var input string
        if scan.Scan() {
            input = scan.Text()
        }
        res := tokenize(input)
        for _, token := range res {
            if (token[0] == '"') || (token[len(token) - 1] == '"') || is_str {
                if (token[0] == '"' && token[len(token) - 1] == '"') {
                    for ch := 1; ch < len(token) - 1; ch++ {
                        if token[ch] == '!' {
                            fmt.Println()
                            continue
                        }
                        fmt.Print(string(token[ch]))
                    }
                    fmt.Println()
                    is_str = false
                    continue
                } else if token[0] == '"' {
                    for ch := 1; ch < len(token); ch++ {
                        if token[ch] == '!' {
                            fmt.Println()
                            continue
                        }
                        fmt.Print(string(token[ch]))
                    }
                    fmt.Print(" ")
                    is_str = true
                    continue
                } else if token[len(token) - 1] == '"' {
                    for ch := 0; ch < len(token) - 1; ch++ {
                        if token[ch] == '!' {
                            fmt.Println()
                            continue
                        }
                        fmt.Print(string(token[ch]))
                    }
                    fmt.Println()
                    is_str = false
                    continue
                } else {
                    for ch := 0; ch < len(token); ch++ {
                        if token[ch] == '!' {
                            fmt.Println()
                            continue
                        }
                        fmt.Print(string(token[ch]))
                    }
                    fmt.Print(" ")
                    is_str = true
                    continue
                }
            }
            if token == "]" {
                is_if = true
            }
            if !is_if {
                continue
            }
            switch token {
            case "[":
                top_ele, err := strconv.Atoi(pop_stack(&stack))
                if err != nil {
                    fmt.Println("Error converting string to int")
                }
                if top_ele <= 0 {
                    is_if = false
                }
            case "]":
                continue
            case "+", "-", "*", "/":
                second_ele, err1 := strconv.Atoi(pop_stack(&stack))
                first_ele, err2 := strconv.Atoi(pop_stack(&stack))
                if err1 != nil || err2 != nil {
                    fmt.Println("Error converting string to int")
                }
                var res string
                if token == "+" {
                    res = strconv.Itoa(second_ele + first_ele)
                } else if token == "-" {
                    res = strconv.Itoa(second_ele - first_ele)
                } else if token == "*" {
                    res = strconv.Itoa(second_ele * first_ele)
                } else {
                    res = strconv.Itoa(second_ele / first_ele)
                }
                stack = push_stack(stack, res)
            case "?":
                var num string
                some, err := fmt.Scanln(&num)
                if err != nil && some != 1 {
                    fmt.Println("Error taking input")
                }
                stack = push_stack(stack, num)
            case "!":
                top_ele := pop_stack(&stack)
                fmt.Println(top_ele)
            case "$$":
                os.Exit(0)
            case "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z":
                val := alp_num[token]
                stack = push_stack(stack, val)
            case "=":
                top_num := pop_stack(&stack)
                addr := pop_stack(&stack)
                top_alp := num_alp[addr]
                new_alp_num[top_alp] = top_num
            case ".":
                addr := pop_stack(&stack)
                var new_num string
                top_alp := num_alp[addr]
                new_num = new_alp_num[top_alp]
                if new_num == "" {
                    new_num = "0"
                }
                stack = push_stack(stack, new_num)
            default:
                stack = push_stack(stack, token)
            }
        }
        fmt.Println("Contents of the stack:")
        for i := len(stack) - 1; i >= 0; i-- {
            fmt.Println(stack[i])
        }
    }
}

func main()  {
    interactive()
}
