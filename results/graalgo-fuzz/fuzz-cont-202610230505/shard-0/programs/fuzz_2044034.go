package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 int16; F1 uint16; F2 bool }

func main() {
    var v1 int8 = -27
    var v2 uint16 = 0
    var v3 uint = 787
    var v4 int32 = 42
    var v5 uint = v3
    if true {
        fmt.Println(v1)
        var v6 uint = (^(v3))
        v7 := GenS1{F0: (v1 >> ((v3) & 63)), F1: (-13 <= v1)}
        _ = v6
        _ = v7
    } else {
        var v8 string = ""
        v8 = ""
        _ = v8
    }
    var v9 string = "\n\t"
    if false {
        if (true || (1 == 4294967295)) {
            fmt.Println(v9)
        } else {
            v9 = "a"
            switch ("foo bar" + v9) {
            case "hello":
                v9 = v9
            }
        }
        v1 += 72
    } else {
        v10 := GenS1{F0: (v1 & v1), F1: false}
        _ = v10
    }
    v11 := GenS2{F0: -358, F1: (v2 >> (((v3 / (((^(v3))) | 1))) & 63)), F2: true}
    v12 := []int{-1, 128, int(v2)}
    v4 += 1
    fmt.Println(v4)
    v5 = (v5 + 0)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
