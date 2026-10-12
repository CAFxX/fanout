package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 uint64 }
type GenS2 struct { F0 uint16 }
func genFn0() bool {
    v4 := GenS1{F0: 990, F1: 642, F2: 4294967295}
    var v5 uint64 = 0
    _ = v4
    _ = v5
    return (3.14159 >= 1e10)
}

func main() {
    v1 := []string{"\n\t", ((("hello" + "a") + "") + "hello")}
    var v2 float64 = 1e10
    var v3 uint = 65535
    var v6 int32 = -947
    for fi7 := 0; fi7 < 14; fi7++ {
        v8 := GenS2{F0: 42}
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v2 = v2
        _ = fi7
        _ = v8
    }
    var v9 uint = (v3 * v3)
    var v10 float64 = v2
    fmt.Println(v9, v9, v10)
    var v11 uint8 = 193
    if true {
        var v12 uint64 = 0
        var v13 float32 = 3.14159
        _ = v12
        _ = v13
    }
    switch "" {
    case "foo bar":
        v1 = append(v1, "\n\t")
    default:
        if true {
            var v14 int16 = 441
            v14 = (v14 & (v14 * (v14 >> (((v9 / ((2) | 1))) & 63))))
            if false {
                var v15 string = "\n\t"
                v16 := GenS2{F0: 283}
                fmt.Println(v6)
                _ = v15
                _ = v16
            } else {
                var v17 bool = true
                _ = v17
            }
            _ = v14
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v6
    _ = v9
}
