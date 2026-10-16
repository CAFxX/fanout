package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }

func main() {
    var v1 int8 = -128
    var v2 uint64 = 944
    var v3 float32 = (3.14159 - (3.14159 - 1e10))
    v4 := make(map[string]string)
    var v5 int8 = v1
    var v6 uint16 = 65535
    v5 += (^(v1))
    v7 := GenS1{F0: v6, F1: uint(v2)}
    var v8 uint8 = 34
    for fi9 := 0; fi9 < 15; fi9++ {
        v5 -= v5
        v10 := GenS1{F0: v6, F1: uint(v8)}
        fmt.Println(v8)
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 8; fi11++ {
        if (!true) {
            v3 = 1e10
            switch 32768 {
            case -90:
                for fi12 := 0; fi12 < 13; fi12++ {
                    switch "\n\t" {
                    case "foo bar":
                        var v13 float64 = -2.25
                        _ = v13
                    case "a":
                        v4[""] = ("\n\t" + "foo bar")
                    }
                    v14 := GenS1{F0: v6, F1: 1}
                    _ = fi12
                    _ = v14
                }
            case -372:
                var v15 uint32 = 2
                _ = v15
            }
            if (1 <= 1) {
                var v16 uint32 = 1
                v17 := []int{(fi11 - (fi11 >> ((65535) & 63))), -9223372036854775808, 988, (fi11 | (fi11 % ((0) | 1)))}
                _ = v16
                _ = v17
            }
        } else {
            fmt.Println(v5, v1, v2)
        }
        var v18 int = fi11
        _ = fi11
        _ = v18
    }
    v19 := []int{258, 127, -9223372036854775808, int(v1)}
    _ = len(v4)
    v20 := GenS1{F0: v6, F1: 790}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
