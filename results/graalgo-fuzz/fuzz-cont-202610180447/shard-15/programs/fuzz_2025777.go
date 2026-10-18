package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
func genFn0(p0 int, p1 int32) int16 {
    var v6 uint = 4294967295
    for fi7 := 0; fi7 < 8; fi7++ {
        var v8 uint64 = 377
        switch (fi7 - 127) {
        case 760:
            var v9 uint32 = 2147483648
            _ = v9
        case 1:
            v6 -= 548
        default:
            var v10 int32 = p1
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    fmt.Println(p0, v6, v6)
    _ = p0
    _ = p1
    _ = v6
    return 0
}
func genFn1(p0 uint, p1 int8, p2 uint32) int32 {
    var v11 uint16 = 562
    switch -9223372036854775808 {
    case 2147483647:
        v12 := []int{}
        _ = v12
    default:
        v11 += (^(v11))
    }
    fmt.Println(p0, p2)
    for fi13 := 0; fi13 < 2; fi13++ {
        var v14 int32 = 361
        _ = fi13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return -1
}

func main() {
    var v1 uint16 = 1
    var v2 uint8 = 61
    var v3 uint16 = (v1 - (v1 * (v1 / (((^(v1))) | 1))))
    var v4 uint16 = (v1 >> ((18446744073709551615) & 63))
    v5 := make(map[string]uint32)
    fmt.Println(v2, v4, v2)
    v15 := []int{-646, 42}
    switch ("hello" + "\n\t") {
    case "\n\t":
        var v16 string = "foo bar"
        _ = v16
    case "a":
        fmt.Println(v3, v1)
    }
    switch (("\n\t" + "") + ("\n\t" + "foo bar")) {
    case "hello":
        v2 = (v2 ^ v2)
    }
    v17 := v5[((("a" + "foo bar") + "\n\t") + "")]
    v1 += (v3 ^ v4)
    v4 = (^(v3))
    for fi18 := 0; fi18 < 18; fi18++ {
        v19 := GenS1{F0: genFn0(fi18, -1), F1: 791}
        _ = fi18
        _ = v19
    }
    var v20 uint = uint(v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
