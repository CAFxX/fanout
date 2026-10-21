package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 int32; F1 int16 }
func genFn0(p0 string, p1 uint16, p2 uint32) uint32 {
    switch ("foo bar" + p0) {
    case "\n\t":
        p2 = p2
    }
    for fi5 := 0; fi5 < 3; fi5++ {
        v6 := []int{}
        if false {
            v7 := GenS1{F0: -128, F1: int16(fi5)}
            _ = v7
        }
        var v8 uint8 = 233
        _ = fi5
        _ = v6
        _ = v8
    }
    if (!("" > "a")) {
        p1 = (p1 + (^(p1)))
        p2 = (p2 | (p2 | (^(p2))))
        v9 := GenS2{F0: -2147483648, F1: 1}
        _ = v9
    }
    switch -9223372036854775808 {
    case -9223372036854775808:
        if true {
            v10 := GenS1{F0: int8(p2), F1: -1}
            _ = v10
        }
    case -392:
        var v11 string = ((p0 + ((((("foo bar" + p0) + ((p0 + ("foo bar" + p0)) + "")) + p0) + ("\n\t" + "\n\t")) + p0)) + p0)
        _ = v11
    case 4294967295:
        if (100 > 42) {
            switch 1 {
            case 827:
                if false {
                    p2 += (p2 + (p2 >> ((255) & 63)))
                    p0 = (p0 + (p0 + ("\n\t" + "")))
                    fmt.Println(p2, p1, p0)
                }
            case -1:
                fmt.Println(p0, p1)
            default:
                v12 := []int{-15}
                _ = v12
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 353
}

func main() {
    v1 := []uint32{}
    var v2 int16 = -287
    v3 := make(map[string]int)
    v4 := make(map[int]bool)
    _ = len(v4)
    v13 := v4[127]
    v14 := GenS2{F0: -1, F1: (v2 | v2)}
    var v15 bool = (!(!v13))
    if (1 != 197) {
        v13 = true
        switch "a" {
        case "\n\t":
            var v16 uint = 42
            _ = v16
        }
    }
    v2 = v2
    v2 = 1
    v17 := GenS2{F0: -1, F1: (v2 % ((-5) | 1))}
    v13 = false
    var v18 uint8 = 64
    switch "a" {
    case "\n\t":
        _ = len(v3)
    case "foo bar":
        if (0 != 42) {
            var v19 uint8 = v18
            _ = v19
        }
    case "a":
        var v20 int16 = (v2 << ((18446744073709551615) & 63))
        _ = v20
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
