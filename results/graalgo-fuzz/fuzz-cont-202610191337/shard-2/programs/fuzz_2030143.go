package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 int }

func main() {
    v1 := make(map[int]string)
    v2 := []string{(((("a" + "a") + "a") + "\n\t") + "foo bar"), ("hello" + ("foo bar" + "\n\t")), "foo bar", "a"}
    var v3 int8 = -1
    var v4 uint64 = 752
    v5 := []string{"hello", "\n\t"}
    var v6 float32 = 0.1
    switch -9223372036854775808 {
    case 531:
        v6 = (v6 - ((1e10 * v6) * v6))
    case -100:
        switch int(v3) {
        case 127:
            switch "" {
            case "a":
                v7 := GenS1{F0: 128, F1: 603, F2: -9223372036854775808}
                _ = v7
            default:
                fmt.Println(v6, v6, v3)
            }
        case -2147483648:
            if true {
                fmt.Println(v4, v3, v6)
                var v8 int64 = 42
                v9 := GenS1{F0: 1, F1: 30, F2: 32768}
                _ = v8
                _ = v9
            }
        }
    }
    var v10 string = "a"
    var v11 float64 = 3.14159
    fmt.Println(v3)
    v4 -= v4
    v11 = (v11 / v11)
    v1[32767] = v10
    if (v3 != (v3 & -109)) {
        fmt.Println(v6, v3)
    }
    v11 = float64((v6 - v6))
    var v12 int64 = -1
    v13 := v1[127]
    var v14 uint16 = 532
    v11 = (0.0 + v11)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
