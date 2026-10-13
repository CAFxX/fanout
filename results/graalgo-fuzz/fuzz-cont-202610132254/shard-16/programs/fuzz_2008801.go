package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint64; F1 string; F2 string }
func genFn0(p0 uint, p1 int64) uint32 {
    p0 = p0
    switch -9223372036854775808 {
    case 4294967295:
        v4 := []int{32768, 4294967295}
        _ = v4
    case -2147483648:
        p0 += p0
    case 130:
        v5 := []int{0}
        _ = v5
    }
    v6 := []int{128, -9223372036854775808}
    for fi7 := 0; fi7 < 16; fi7++ {
        p1 = (p1 + (^(p1)))
        for fi8 := 0; fi8 < 17; fi8++ {
            p0 = p0
            v9 := GenS1{F0: 18446744073709551615}
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    v10 := GenS1{F0: 18446744073709551615}
    _ = p0
    _ = p1
    _ = v10
    _ = v6
    return 984
}
func genFn1(p0 int16, p1 uint, p2 int16) int {
    switch "a" {
    case "a":
        var v11 int32 = 2147483647
        _ = v11
    }
    var v12 int64 = 9223372036854775807
    switch 32768 {
    case 1:
        var v13 int16 = (^(p2))
        _ = v13
    case -691:
        v14 := []int{127}
        _ = v14
    default:
        v15 := GenS2{F0: 128, F1: 1}
        _ = v15
    }
    p0 -= p2
    p2 -= (p2 & 1000)
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 32768
}

func main() {
    var v1 int64 = 0
    v2 := []uint32{2147483648, 324}
    v3 := make(map[string]int)
    switch "foo bar" {
    case "\n\t":
        v1 += -648
    }
    var v16 float64 = (1.5 / 1.5)
    delete(v3, ("a" + "\n\t"))
    v1 = v1
    v17 := GenS3{F0: -2147483648}
    v1 += v1
    var v18 uint16 = 32767
    switch "\n\t" {
    case "foo bar":
        if (!false) {
            var v19 float64 = 3.14159
            v20 := GenS4{F0: 1, F1: (("foo bar" + "") + ("a" + ("\n\t" + "hello"))), F2: ("foo bar" + "foo bar")}
            _ = v19
            _ = v20
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
