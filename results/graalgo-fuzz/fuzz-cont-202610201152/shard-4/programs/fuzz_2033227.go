package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 uint16 }
func genFn0() int {
    v4 := []int{}
    var v5 uint = 1
    _ = v4
    _ = v5
    return 32768
}
func genFn1(p0 int64, p1 bool) int32 {
    for fi6 := 0; fi6 < 7; fi6++ {
        v7 := []int{genFn0(), -1, (fi6 | fi6), (fi6 - (fi6 & fi6))}
        switch ("foo bar" + "\n\t") {
        case "hello":
            var v8 int64 = -1
            _ = v8
        case "\n\t":
            for fi9 := 0; fi9 < 15; fi9++ {
                var v10 uint = 585
                switch 725 {
                case 1:
                    fmt.Println(p0, fi6)
                case -577:
                    fmt.Println(fi9, v10)
                case 702:
                    if len(v7) > 0 {
                        v7[0] = fi9
                    }
                default:
                    v10 -= (v10 * v10)
                }
                _ = fi9
                _ = v10
            }
        case "a":
            for fi11 := 0; fi11 < 7; fi11++ {
                var v12 int32 = 793
                v13 := GenS1{F0: ((418 > 32767) || p1), F1: -128, F2: 866}
                if len(v7) > 0 {
                    v7[0] = fi6
                }
                _ = fi11
                _ = v12
                _ = v13
            }
        }
        _ = fi6
        _ = v7
    }
    p1 = true
    p0 = p0
    _ = p0
    _ = p1
    return 0
}

func main() {
    var v1 int32 = 42
    var v2 uint64 = 42
    var v3 int64 = 0
    if (("foo bar" + "\n\t") == "") {
        var v14 uint32 = 42
        _ = v14
    }
    v1 = genFn1(v3, (!((3.14159 / 1e10) != -2.25)))
    var v15 float64 = ((3.14159 / 1.0) * 1.5)
    v1 -= 0
    var v16 float64 = 1.5
    v17 := []int{639, 9223372036854775807, 834}
    for fi18 := 0; fi18 < 14; fi18++ {
        if ((1e10 + -2.25) > (3.14159 * 3.14159)) {
            v3 = -1
        }
        fmt.Println(v2, fi18)
        var v19 uint64 = v2
        _ = fi18
        _ = v19
    }
    fmt.Println(v16, v15)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
