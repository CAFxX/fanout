package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 uint16 }

func main() {
    var v1 string = ("foo bar" + ((("\n\t" + (("a" + ("" + ("hello" + ((("" + "foo bar") + "") + "\n\t")))) + ("hello" + ""))) + "a") + ("\n\t" + "hello")))
    var v2 uint32 = 999
    v3 := make(map[string]int)
    var v4 int32 = 116
    var v5 bool = (v4 <= v4)
    v6 := make(map[int]string)
    var v7 uint64 = 2
    if v5 {
        v2 = (v2 / (((v2 / ((v2) | 1))) | 1))
        var v8 uint16 = 1
        v4 = (v4 * v4)
        _ = v8
    }
    v9 := GenS1{F0: v2, F1: int16(v7), F2: 2}
    var v10 uint32 = 2147483647
    v10 -= v2
    var v11 float64 = -2.25
    var v12 bool = (-2.25 != 0.0)
    v6[-9223372036854775808] = v1
    v13 := GenS1{F0: (v2 << ((uint(v7)) & 63)), F1: -497, F2: 1}
    if true {
        if v5 {
            switch -1 {
            case -57:
                v2 = v10
            default:
                fmt.Println(v7)
            }
        } else {
            var v14 uint32 = 0
            var v15 float64 = float64(0.1)
            _ = v14
            _ = v15
        }
        if (!false) {
            if true {
                var v16 bool = false
                _ = v16
            }
            v10 = (v10 >> ((255) & 63))
        } else {
            v4 = (v4 & (^(v4)))
            v17 := v3[v1]
            v7 = (v7 * v7)
            _ = v17
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
