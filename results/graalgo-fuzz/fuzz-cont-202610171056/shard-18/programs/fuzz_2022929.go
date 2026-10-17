package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 uint32 }

func main() {
    var v1 int8 = -128
    var v2 bool = true
    var v3 string = ("\n\t" + "a")
    var v4 uint8 = 2
    for fi5 := 0; fi5 < 13; fi5++ {
        v6 := []int{(-(fi5)), (^(fi5)), fi5, fi5}
        _ = fi5
        _ = v6
    }
    fmt.Println(v1)
    switch 32768 {
    case -902:
        if ("hello" != v3) {
            v1 = 0
            v2 = (2147483648 > 2147483647)
        }
    }
    if v2 {
        var v7 int = -119
        v4 += v4
        v3 = ""
        _ = v7
    }
    for fi8 := 0; fi8 < 18; fi8++ {
        var v9 int16 = -1
        _ = fi8
        _ = v9
    }
    fmt.Println(v1, v3, v3)
    var v10 int16 = -32768
    if (42 != 0) {
        v11 := GenS1{F0: 2147483647, F1: v2, F2: 2147483648}
        v12 := GenS1{F0: int32(v1), F1: (v2 && v2), F2: 65535}
        _ = v11
        _ = v12
    }
    fmt.Println(v4, v1)
    var v13 int16 = -1
    var v14 int64 = -54
    v15 := []int{-2147483648}
    fmt.Println(v4)
    v4 = (v4 - (v4 >> ((68) & 63)))
    if v2 {
        v14 -= (v14 % (((v14 << ((18446744073709551615) & 63))) | 1))
    }
    v1 += 42
    v3 = (v3 + "foo bar")
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
