package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 uint }
type GenS2 struct { F0 int; F1 uint16; F2 uint8 }

func main() {
    var v1 uint16 = 65535
    var v2 uint = 65535
    v3 := make(map[string]bool)
    switch 42 {
    case 1:
        fmt.Println(v2, v1)
    case -455:
        var v4 uint = v2
        _ = v4
    default:
        v2 -= 0
    }
    switch ("hello" + (("hello" + "hello") + "hello")) {
    case "a":
        delete(v3, "foo bar")
    }
    if true {
        if true {
            switch "" {
            case "foo bar":
                var v5 uint32 = uint32(v1)
                _ = v5
            case "\n\t":
                v6 := GenS1{F0: 255, F1: 42, F2: v2}
                _ = v6
            case "":
                v7 := GenS1{F0: (v2 / ((v2) | 1)), F1: 128, F2: 65535}
                _ = v7
            }
            v2 -= 65535
            var v8 int16 = -663
            _ = v8
        }
    }
    var v9 int8 = 82
    v3["\n\t"] = false
    var v10 bool = (true && ((!true) && (!true)))
    var v11 uint32 = 852
    v10 = (((0 == 280) || ((float64(1e10) > ((((3.14159 - (1e10 - 3.14159)) / (1e10 - 0.0)) - 0.0) + 1e10)) && v10)) && (-1 != 1))
    v1 = (v1 - uint16(v2))
    if true {
        fmt.Println(v1)
        if (215 == 42) {
            if (133 == 449) {
                v10 = v10
                var v12 uint16 = v1
                for fi13 := 0; fi13 < 19; fi13++ {
                    v14 := GenS2{F0: -9223372036854775808, F1: v12, F2: 0}
                    var v15 string = "\n\t"
                    fmt.Println(v2, v9, v11)
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                _ = v12
            } else {
                v10 = (v10 || false)
                switch "\n\t" {
                case "":
                    v16 := []int{-9223372036854775808, -346, 9223372036854775807, 32768}
                    _ = v16
                case "a":
                    v1 += v1
                default:
                    fmt.Println(v2, v11, v2)
                }
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v9
}
