package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 bool; F1 uint; F2 bool }

func main() {
    var v1 uint64 = 1
    var v2 uint8 = 65
    var v3 int8 = 42
    var v4 uint32 = 2147483648
    var v5 int32 = 42
    v2 = (v2 + 128)
    fmt.Println(v4, v3)
    v3 = 127
    if ((^(v2)) == (^(v2))) {
        switch 127 {
        case 32767:
            if false {
                var v6 int16 = 42
                var v7 int = -1
                var v8 int = -9223372036854775808
                _ = v6
                _ = v7
                _ = v8
            } else {
                if (!(77 > -32768)) {
                    v3 -= v3
                }
                if ((42 <= 0) && true) {
                    v9 := []int{-774, 42, 32768}
                    v9 = append(v9, 32767)
                    fmt.Println(v1, v3)
                    _ = v9
                }
                v3 += (v3 | v3)
            }
        }
        v10 := []int{-9223372036854775808, 806}
        _ = v10
    }
    var v11 float64 = 3.14159
    v4 += v4
    v12 := GenS1{F0: true, F1: ("foo bar" + "\n\t")}
    var v13 float64 = v11
    switch 4294967295 {
    case 0:
        v2 = v2
    case 4294967295:
        if (!(!true)) {
            var v14 int64 = 0
            _ = v14
        } else {
            switch -2147483648 {
            case 32767:
                v15 := []int{0, -571}
                _ = v15
            case -2147483648:
                fmt.Println(v5, v4)
            }
            var v16 string = "hello"
            v17 := GenS2{F0: ((v1 - v1) != v1), F1: 2, F2: (1.5 < (3.14159 * 0.0))}
            _ = v16
            _ = v17
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
