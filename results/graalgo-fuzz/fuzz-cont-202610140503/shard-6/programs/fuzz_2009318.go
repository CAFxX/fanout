package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 uint8; F1 uint }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]string)
    var v4 uint16 = 894
    if (65535 > 2147483647) {
        var v5 uint = 1
        if false {
            v6 := GenS1{F0: (v5 ^ 42), F1: v4}
            var v7 string = "hello"
            _ = v6
            _ = v7
        }
        _ = v5
    } else {
        v8 := v2[1]
        v4 = 131
        _ = v8
    }
    v9 := GenS2{F0: 2, F1: 42}
    fmt.Println(v4)
    var v10 int8 = 100
    var v11 uint8 = 255
    switch ("\n\t" + "foo bar") {
    case "\n\t":
        var v12 uint64 = 18446744073709551615
        _ = v12
    case "a":
        var v13 float64 = 3.14159
        _ = v13
    case "":
        if true {
            v14 := GenS1{F0: 2, F1: (v4 + (v4 % (((v4 << ((560) & 63))) | 1)))}
            v11 = v11
            v15 := GenS1{F0: 255, F1: v4}
            _ = v14
            _ = v15
        } else {
            fmt.Println(v11, v10, v10)
            switch "foo bar" {
            case "\n\t":
                v16 := GenS1{F0: 42, F1: v4}
                _ = v16
            case "":
                var v17 uint16 = (v4 + v4)
                _ = v17
            case "hello":
                v10 += 1
            }
        }
    default:
        v11 = (v11 / (((v11 << ((1) & 63))) | 1))
    }
    var v18 uint32 = 1
    var v19 uint8 = 0
    v3["a"] = "\n\t"
    v10 += v10
    v1[32768] = -301
    fmt.Println(v4, v4)
    var v20 float32 = (1e10 / (1.5 + ((0.0 + 0.1) / (3.14159 / ((1e10 / 1.0) + (0.0 * -2.25))))))
    var v21 int16 = 1000
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v9
}
