package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32; F1 uint8; F2 uint }
func genFn0(p0 uint16, p1 uint, p2 int16) uint8 {
    fmt.Println(p2, p2, p2)
    switch ("\n\t" + "a") {
    case "a":
        p0 = (p0 >> ((p1) & 63))
    case "\n\t":
        v4 := []int{32768, 127}
        _ = v4
    case "foo bar":
        var v5 uint16 = p0
        _ = v5
    }
    v6 := GenS1{F0: p0}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 66
}

func main() {
    var v1 uint64 = 2
    var v2 int16 = 1
    var v3 int16 = (v2 + (v2 + (v2 << ((65535) & 63))))
    switch -9223372036854775808 {
    case 42:
        var v7 int32 = 2147483647
        _ = v7
    case -49:
        var v8 int64 = -654
        _ = v8
    case 4294967295:
        v3 = v2
    default:
        var v9 int8 = 100
        _ = v9
    }
    for fi10 := 0; fi10 < 9; fi10++ {
        var v11 string = "foo bar"
        var v12 bool = (612 != 2147483648)
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 string = "foo bar"
    v3 = (v3 & (v2 | v2))
    switch v13 {
    case "a":
        var v14 uint = 4294967295
        _ = v14
    case "hello":
        fmt.Println(v13)
    default:
        v15 := GenS2{F0: -2147483648, F1: 42, F2: 18446744073709551615}
        _ = v15
    }
    if true {
        v2 -= v2
    }
    v16 := GenS1{F0: 460}
    v17 := []int{}
    var v18 uint8 = 0
    var v19 int64 = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
