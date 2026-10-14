package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 uint64 }

func main() {
    var v1 uint16 = 42
    var v2 uint8 = 1
    var v3 uint32 = 2
    var v4 bool = false
    var v5 bool = (255 > v1)
    v3 += 0
    switch 2147483647 {
    case -944:
        v6 := []int{32767, 127, 4294967295, 845}
        _ = v6
    case 4294967295:
        v3 = 778
    case -2147483648:
        v1 -= v1
    }
    fmt.Println(v4)
    var v7 int = -1
    v5 = true
    v8 := []int{}
    if (42 > 581) {
        if (!v4) {
            if (("a" + "") > "hello") {
                fmt.Println(v5, v2, v7)
                v7 -= -208
                var v9 uint16 = (v1 ^ 2)
                _ = v9
            }
            if ("\n\t" != ("foo bar" + ("\n\t" + "\n\t"))) {
                v4 = v5
            } else {
                var v10 int16 = -420
                v10 -= v10
                _ = v10
            }
        }
        fmt.Println(v1, v7)
    }
    switch v7 {
    case 9223372036854775807:
        var v11 uint32 = (v3 | (v3 & v3))
        _ = v11
    default:
        if v5 {
            v12 := GenS1{F0: 127, F1: (^(v3)), F2: 42}
            var v13 bool = false
            if false {
                v8 = append(v8, v7)
                _ = len(v8)
            }
            _ = v12
            _ = v13
        }
    }
    var v14 int32 = -607
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
