package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 uint8 }

func main() {
    var v1 uint8 = 120
    var v2 int8 = 42
    var v3 uint16 = 365
    var v4 uint32 = 857
    var v5 int32 = 858
    var v6 int = 978
    var v7 uint32 = (^(v4))
    switch ("hello" + (("hello" + ("hello" + "a")) + ("foo bar" + ""))) {
    case "\n\t":
        var v8 uint = 0
        _ = v8
    }
    v4 = v4
    v9 := []int{v6}
    for fi10 := 0; fi10 < 14; fi10++ {
        for fi11 := 0; fi11 < 7; fi11++ {
            if false {
                var v12 int16 = -32768
                _ = v12
            }
            fmt.Println(v3, v3)
            _ = fi11
        }
        _ = fi10
    }
    if (!((v3 - v3) <= 65535)) {
        v6 = (v6 | v6)
        v2 = 127
    } else {
        var v13 int8 = (v2 * 0)
        _ = len(v9)
        _ = v13
    }
    switch (v6 % (((v6 | (v6 * v6))) | 1)) {
    case 2147483647:
        v14 := GenS1{F0: (v5 >> ((65535) & 63)), F1: (v6 * 2147483647)}
        _ = v14
    case -2147483648:
        fmt.Println(v1, v4, v1)
    default:
        v3 += 0
    }
    if false {
        v2 = 1
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
    } else {
        v5 += 2147483647
        var v15 int8 = v2
        _ = v15
    }
    fmt.Println(v6, v1)
    if (false && true) {
        for fi16 := 0; fi16 < 20; fi16++ {
            var v17 uint32 = (v4 << ((4294967295) & 63))
            _ = fi16
            _ = v17
        }
        v18 := GenS2{F0: (v1 ^ 2)}
        _ = v18
    }
    v2 -= (v2 >> ((421) & 63))
    v9 = append(v9, 0)
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
