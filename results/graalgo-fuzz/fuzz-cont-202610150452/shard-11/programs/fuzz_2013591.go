package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 bool, p1 uint16, p2 string) uint8 {
    v7 := []int{319, 0}
    p0 = false
    if (32767 >= 32767) {
        v8 := GenS1{F0: 100}
        v9 := GenS1{F0: 64}
        p2 = p2
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 255
}

func main() {
    var v1 string = "a"
    v2 := make(map[int]int)
    v3 := []string{(v1 + v1), v1, "hello", v1}
    var v4 uint32 = 42
    var v5 int = 1
    v6 := make(map[string]bool)
    for fi10 := 0; fi10 < 13; fi10++ {
        v4 = (v4 % ((684) | 1))
        fmt.Println(v4)
        _ = fi10
    }
    v1 = ""
    var v11 uint16 = 821
    switch v5 {
    case 42:
        v4 += (v4 % ((v4) | 1))
    case 127:
        switch v1 {
        case "":
            v1 = (v1 + v1)
        case "\n\t":
            var v12 uint = 237
            _ = v12
        }
    case 128:
        v2[v5] = v5
    }
    if ((v1 < "foo bar") && false) {
        fmt.Println(v1, v11, v5)
        if (0 < 143) {
            v6["a"] = true
            v2[v5] = int(v11)
            fmt.Println(v1, v5, v4)
        }
    }
    var v13 int = v5
    v2[939] = v5
    var v14 uint8 = 42
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
