package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 string; F1 string; F2 uint8 }

func main() {
    var v1 int8 = 127
    var v2 uint16 = 32767
    var v3 float64 = 0.1
    var v4 uint8 = 117
    var v5 float32 = (1.5 - 1.5)
    fmt.Println(v3, v1, v3)
    v6 := []int{128, 42, 1}
    v7 := GenS1{F0: 319, F1: 100}
    switch "" {
    case "a":
        v1 = 42
    case "":
        v2 = 1
    case "\n\t":
        v2 = 1
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v8 uint = 0
    v1 = (-(v1))
    if (true && false) {
        var v9 uint16 = 1
        v10 := GenS2{F0: "hello", F1: (("foo bar" + "a") + "\n\t"), F2: v4}
        _ = v10
        _ = v9
    }
    var v11 int64 = -1
    v5 = 3.14159
    switch "" {
    case "a":
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    case "\n\t":
        if (true && false) {
            var v12 int8 = (-(v1))
            fmt.Println(v1, v12, v1)
            _ = v12
        } else {
            if (v8 < (v8 * v8)) {
                var v13 uint32 = 42
                var v14 string = ""
                v6 = append(v6, 562)
                _ = v13
                _ = v14
            } else {
                var v15 uint16 = (v2 / ((32767) | 1))
                _ = v15
            }
            v16 := GenS1{F0: 93, F1: v1}
            _ = v16
        }
    }
    v5 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
