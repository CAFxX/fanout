package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 string }
type GenS2 struct { F0 int16; F1 bool }

func main() {
    var v1 uint16 = 32767
    var v2 int32 = 2147483647
    var v3 int32 = (v2 - (-(v2)))
    var v4 int = 124
    var v5 uint32 = 1
    var v6 int = v4
    if (("" + "foo bar") > ("\n\t" + "hello")) {
        v6 -= 128
    } else {
        v6 = (v4 ^ (v6 >> ((1) & 63)))
    }
    var v7 uint = 65535
    fmt.Println(v6, v4)
    fmt.Println(v1)
    v8 := GenS1{F0: v5, F1: 75, F2: ""}
    v4 = v6
    var v9 int32 = v2
    v9 = v2
    fmt.Println(v7)
    v7 -= v7
    v6 = 227
    v10 := GenS2{F0: 1000, F1: ((-(v6)) != v6)}
    if (!(true || (((-(v2)) > int32(v1)) && (0 < -1)))) {
        v11 := []int{}
        if len(v11) > 0 {
            v11[0] = (v4 / ((v6) | 1))
        }
        _ = v11
    }
    fmt.Println(v6, v9, v1)
    var v12 string = "\n\t"
    v13 := []int{v6, v6, (v4 * 349)}
    for fi14 := 0; fi14 < 8; fi14++ {
        fmt.Println(v12, v3)
        _ = fi14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
