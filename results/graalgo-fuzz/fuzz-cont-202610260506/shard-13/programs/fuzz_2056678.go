package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 int8 }
func genFn0(p0 int8) int8 {
    p0 = (-(p0))
    v6 := []int{226, 2147483647, 878}
    v7 := GenS1{F0: 100, F1: 1000}
    for fi8 := 0; fi8 < 13; fi8++ {
        v9 := GenS1{F0: 255, F1: -1}
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 15; fi10++ {
        var v11 uint32 = 42
        var v12 uint32 = (v11 - 2)
        var v13 int32 = -920
        _ = fi10
        _ = v11
        _ = v12
        _ = v13
    }
    _ = p0
    _ = v6
    _ = v7
    return p0
}
func genFn1(p0 bool, p1 uint64, p2 int16) string {
    switch "" {
    case "\n\t":
        switch "hello" {
        case "hello":
            v14 := []int{-9223372036854775808, 42, int(p2)}
            _ = v14
        case "a":
            v15 := GenS1{F0: 113, F1: p2}
            _ = v15
        }
    case "hello":
        p0 = (0 <= 1)
    case "a":
        p2 += p2
    default:
        switch 32767 {
        case -9223372036854775808:
            v16 := []int{9223372036854775807, -165, 4294967295}
            _ = v16
        case 436:
            var v17 uint8 = 127
            _ = v17
        case 127:
            v18 := GenS1{F0: 196, F1: (p2 << ((1) & 63))}
            _ = v18
        default:
            var v19 string = ("foo bar" + ("a" + ("foo bar" + "hello")))
            _ = v19
        }
    }
    var v20 string = (("hello" + "foo bar") + ("" + ((("hello" + "\n\t") + ("hello" + "hello")) + ("\n\t" + ("\n\t" + "a")))))
    var v21 uint64 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    _ = v21
    return (v20 + ("foo bar" + ("a" + v20)))
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{("a" + "\n\t"), (("foo bar" + (((("foo bar" + "hello") + ("a" + ((("hello" + "hello") + "hello") + "\n\t"))) + "foo bar") + "hello")) + "a"), (("hello" + "\n\t") + ""), "foo bar"}
    v3 := []uint32{2147483648, 48, 4294967295, 42}
    v4 := make(map[uint32]int)
    var v5 string = "a"
    var v22 uint64 = 42
    if len(v2) > 0 {
        v2[0] = "foo bar"
    }
    _ = len(v1)
    fmt.Println(v22)
    v5 = v5
    v23 := GenS1{F0: 2, F1: 32767}
    v5 = "hello"
    for fi24 := 0; fi24 < 11; fi24++ {
        v25 := GenS2{F0: 1}
        fmt.Println(v22)
        _ = fi24
        _ = v25
    }
    fmt.Println(v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
