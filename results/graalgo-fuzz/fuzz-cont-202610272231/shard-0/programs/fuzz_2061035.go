package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 int8; F1 uint64 }

func main() {
    var v1 int16 = 1000
    var v2 int32 = 42
    var v3 int8 = 1
    var v4 float32 = ((3.14159 + float32((((3.14159 - (1e10 * (0.0 + 3.14159))) - 3.14159) / float64(1e10)))) * 0.1)
    var v5 uint64 = 1
    var v6 int64 = 125
    if ((!(0.1 < v4)) && (-524 > v6)) {
        if ("foo bar" != "foo bar") {
            switch "hello" {
            case "a":
                v7 := []int{2147483647}
                _ = v7
            case "\n\t":
                v5 -= (v5 & (^(v5)))
            case "":
                v3 += (v3 | (v3 + 83))
            }
        }
        v4 = v4
    } else {
        var v8 int8 = 0
        _ = v8
    }
    if ((^(v3)) < (v3 >> ((1) & 63))) {
        v9 := GenS1{F0: (v1 ^ v1), F1: (v5 - 0)}
        if (127 < 72) {
            v10 := []int{2147483647}
            _ = v10
        } else {
            for fi11 := 0; fi11 < 1; fi11++ {
                fmt.Println(v2, v4)
                v2 = 1
                switch (("" + "\n\t") + "") {
                case "a":
                    v4 = 1.5
                case "":
                    var v12 uint = 18446744073709551615
                    _ = v12
                }
                _ = fi11
            }
            var v13 int64 = v6
            var v14 uint32 = 741
            _ = v13
            _ = v14
        }
        v15 := []int{}
        _ = v15
        _ = v9
    }
    fmt.Println(v6, v6)
    v16 := GenS1{F0: (v1 * -32768), F1: 0}
    fmt.Println(v2)
    var v17 int = -2147483648
    if (!true) {
        var v18 bool = (!(32767 != 65535))
        switch ("foo bar" + ("foo bar" + "")) {
        case "\n\t":
            v19 := GenS2{F0: (v3 + v3), F1: 2}
            _ = v19
        default:
            v20 := []int{}
            _ = v20
        }
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
