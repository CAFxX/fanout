package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 uint }
func genFn0(p0 bool) int {
    var v5 uint8 = 1
    var v6 float32 = 1e10
    v7 := GenS1{F0: 42, F1: 100, F2: 0}
    v5 = (^(v5))
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 128
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint32 = 961
    var v3 float32 = ((1.5 * 0.0) + float32(0.1))
    v4 := make(map[int]uint32)
    v1["hello"] = v2
    v3 -= (-2.25 - v3)
    delete(v1, "")
    v2 = 2147483648
    v3 += v3
    var v8 uint = 688
    v9 := GenS1{F0: int64(v8), F1: 1, F2: (v8 / (((v8 % (((v8 + (v8 + uint(v2)))) | 1))) | 1))}
    v8 += v8
    fmt.Println(v8, v3)
    fmt.Println(v3, v8, v2)
    switch ("foo bar" + "hello") {
    case "foo bar":
        fmt.Println(v3, v3, v3)
    }
    switch 2147483647 {
    case 508:
        for fi10 := 0; fi10 < 1; fi10++ {
            var v11 int8 = 1
            _ = fi10
            _ = v11
        }
    default:
        fmt.Println(v2, v8)
    }
    v8 = (v8 / ((v8) | 1))
    for fi12 := 0; fi12 < 8; fi12++ {
        if (127 >= 127) {
            if true {
                var v13 uint32 = v2
                fmt.Println(v8, v2)
                fmt.Println(fi12)
                _ = v13
            }
            v14 := GenS1{F0: 871, F1: -128, F2: v8}
            _ = v14
        }
        _ = fi12
    }
    fmt.Println(v8)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
