package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64; F1 int32 }

func main() {
    var v1 string = "a"
    var v2 int16 = 1000
    v3 := make(map[uint32]bool)
    v2 = v2
    delete(v3, 4294967295)
    if (127 > 127) {
        if ((0.1 / (0.0 + 3.14159)) >= 0.0) {
            var v4 bool = (((!false) && false) || (false || (true && false)))
            fmt.Println(v2, v2)
            _ = v4
        }
    }
    fmt.Println(v2, v1)
    v5 := GenS1{F0: ""}
    v1 = v1
    var v6 bool = false
    v6 = false
    if ((1e10 < (((1e10 - (0.1 / float32(((3.14159 - float64(-2.25)) * (0.1 - 3.14159))))) / 1.5) * 1e10)) || v6) {
        for fi7 := 0; fi7 < 12; fi7++ {
            v1 = v1
            switch "a" {
            case "\n\t":
                v2 += -1
            case "hello":
                switch v1 {
                case "foo bar":
                    v2 = (v2 ^ (v2 + (v2 % ((32767) | 1))))
                default:
                    v8 := GenS1{F0: (v1 + "\n\t")}
                    _ = v8
                }
            case "foo bar":
                var v9 float64 = 0.1
                _ = v9
            }
            v2 = (v2 * (^(v2)))
            _ = fi7
        }
        fmt.Println(v1)
    } else {
        fmt.Println(v2, v6)
    }
    var v10 int16 = v2
    v11 := []int{-9223372036854775808, -9223372036854775808}
    for fi12 := 0; fi12 < 7; fi12++ {
        for fi13 := 0; fi13 < 4; fi13++ {
            v14 := GenS2{F0: -154, F1: 8}
            _ = fi13
            _ = v14
        }
        _ = fi12
    }
    var v15 int8 = 1
    v16 := GenS2{F0: 42, F1: 2147483647}
    var v17 int = 4294967295
    var v18 uint8 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v5
    _ = v6
}
