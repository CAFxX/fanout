package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 uint64 }
type GenS2 struct { F0 uint }
func genFn0(p0 uint, p1 string) string {
    var v5 uint = (^(p0))
    if ((((0.0 * 1e10) + 0.1) + -2.25) != 0.0) {
        v6 := GenS1{F0: true, F1: ("" + p1), F2: 193}
        _ = v6
    } else {
        v5 = p0
    }
    fmt.Println(p0, v5, p1)
    var v7 uint8 = 0
    _ = p0
    _ = p1
    _ = v5
    _ = v7
    return "\n\t"
}
func genFn1(p0 int8, p1 bool, p2 int32) uint8 {
    if false {
        var v8 string = ""
        p1 = (1000 != 42)
        _ = v8
    } else {
        for fi9 := 0; fi9 < 1; fi9++ {
            if false {
                p2 += 1
                v10 := GenS2{F0: 147}
                _ = v10
            }
            for fi11 := 0; fi11 < 9; fi11++ {
                v12 := []int{}
                _ = fi11
                _ = v12
            }
            _ = fi9
        }
        v13 := []int{-384, -695, 808, -494}
        _ = v13
    }
    var v14 float32 = 0.0
    v14 += 1e10
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return 128
}

func main() {
    var v1 bool = false
    v2 := make(map[uint32]string)
    var v3 int = -629
    var v4 int = 1
    var v15 float64 = 0.0
    v3 = v3
    v1 = (v1 || (9223372036854775807 == -333))
    var v16 uint16 = 2
    switch "" {
    case "":
        v17 := GenS1{F0: (!true), F1: "", F2: 601}
        _ = v17
    case "\n\t":
        v1 = v1
    case "a":
        v18 := []int{v4, (-(v4)), 128}
        _ = v18
    default:
        v19 := []int{4294967295, (v4 << ((0) & 63))}
        _ = v19
    }
    v16 += 32767
    for fi20 := 0; fi20 < 20; fi20++ {
        for fi21 := 0; fi21 < 11; fi21++ {
            switch ("" + (("hello" + ("a" + "\n\t")) + ((("a" + "foo bar") + ("a" + "\n\t")) + "a"))) {
            case "hello":
                v22 := []int{}
                _ = v22
            default:
                v23 := []int{(fi21 + v4), (fi21 << ((18446744073709551615) & 63)), v4, -789}
                _ = v23
            }
            if (0 > -9223372036854775808) {
                v15 = (v15 + v15)
                _ = len(v2)
                v15 = (v15 + v15)
            } else {
                fmt.Println(v15, fi21)
                var v24 string = "hello"
                if ((((!v1) || v1) && v1) || (42 != -32768)) {
                    v25 := GenS2{F0: 2}
                    v4 = v4
                    v16 += (v16 & v16)
                    v24 = v24
                    _ = v25
                } else {
                    fmt.Println(fi20, fi21, v4)
                }
                _ = v24
            }
            fmt.Println(v1, v4)
            _ = fi21
        }
        _ = fi20
    }
    v15 += v15
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
