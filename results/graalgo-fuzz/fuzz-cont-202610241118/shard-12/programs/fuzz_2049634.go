package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16; F1 int16; F2 int }
func genFn0(p0 uint, p1 string) string {
    if false {
        fmt.Println(p0)
        v4 := []int{42, 920, 0}
        _ = v4
    }
    var v5 uint64 = 208
    _ = p0
    _ = p1
    _ = v5
    return p1
}

func main() {
    v1 := []int{0, 806}
    var v2 int8 = -128
    var v3 string = ("" + "hello")
    var v6 float32 = -2.25
    v1 = append(v1, 127)
    fmt.Println(v6, v3, v2)
    var v7 uint16 = 333
    v8 := GenS1{F0: v3}
    var v9 float32 = ((v6 * 3.14159) / -2.25)
    var v10 string = (v3 + "")
    switch ((("" + v10) + v3) + "") {
    case "foo bar":
        v3 = v10
    default:
        if (!(!false)) {
            v10 = v10
        }
    }
    var v11 int32 = -57
    if (!((-(v2)) > (v2 & v2))) {
        _ = len(v1)
        switch "\n\t" {
        case "foo bar":
            var v12 uint16 = 32767
            _ = v12
        case "\n\t":
            v3 = v10
        case "":
            v7 -= (v7 | v7)
        default:
            v9 += v6
        }
        v13 := GenS2{F0: (v7 & (v7 % ((v7) | 1))), F1: 703, F2: 1}
        _ = v13
    }
    if len(v1) > 0 {
        v1[0] = 9223372036854775807
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
