package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 string }
type GenS2 struct { F0 int16 }

func main() {
    var v1 uint16 = 764
    var v2 uint64 = 42
    var v3 uint16 = (^(v1))
    var v4 uint = 2
    if (!(!false)) {
        var v5 int8 = 100
        _ = v5
    } else {
        switch 128 {
        case 1:
            var v6 uint8 = 255
            _ = v6
        case 277:
            var v7 int32 = 42
            _ = v7
        case 2147483647:
            if false {
                var v8 float32 = 3.14159
                v9 := GenS1{F0: true, F1: (v4 >> (((v4 ^ v4)) & 63)), F2: "hello"}
                switch (("hello" + "hello") + "\n\t") {
                case "\n\t":
                    var v10 uint32 = 689
                    _ = v10
                }
                _ = v8
                _ = v9
            }
        }
        fmt.Println(v2, v4)
    }
    fmt.Println(v1, v1, v3)
    var v11 int16 = 1
    fmt.Println(v11, v11, v1)
    v12 := []int{-782, 9223372036854775807}
    switch int(v11) {
    case 1:
        v13 := GenS2{F0: (v11 % (((v11 + (-(v11)))) | 1))}
        _ = v13
    case 0:
        var v14 float32 = 1e10
        _ = v14
    }
    var v15 uint16 = (v3 ^ 1)
    v16 := GenS1{F0: (1.5 != ((1e10 * (0.0 + ((0.1 / (1e10 + (0.0 + 0.0))) - 0.0))) / 1.5)), F1: v4, F2: ""}
    fmt.Println(v1, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
