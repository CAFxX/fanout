package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 uint64; F1 uint8; F2 uint32 }
func genFn0(p0 int64, p1 bool) int {
    if p1 {
        switch (((("\n\t" + "a") + "a") + "") + "") {
        case "a":
            v6 := []int{-9223372036854775808, 767, 42}
            _ = v6
        case "foo bar":
            v7 := GenS1{F0: ("" + "a"), F1: ""}
            _ = v7
        case "\n\t":
            var v8 int64 = -1
            _ = v8
        default:
            fmt.Println(p0, p1)
        }
        var v9 int64 = 9223372036854775807
        _ = v9
    } else {
        var v10 int64 = -1
        _ = v10
    }
    for fi11 := 0; fi11 < 19; fi11++ {
        if (false || false) {
            v12 := []int{-338, fi11, fi11, -1}
            _ = v12
        } else {
            var v13 int64 = p0
            _ = v13
        }
        var v14 int = -982
        v14 = fi11
        _ = fi11
        _ = v14
    }
    _ = p0
    _ = p1
    return -1
}

func main() {
    var v1 uint16 = 0
    v2 := []string{"hello"}
    var v3 uint32 = 96
    var v4 uint32 = (v3 >> ((18446744073709551615) & 63))
    var v5 float64 = 3.14159
    if len(v2) > 0 {
        v2[0] = ("" + ("a" + ("\n\t" + "\n\t")))
    }
    fmt.Println(v1, v1)
    var v15 float64 = 1e10
    v15 += float64(1e10)
    v4 += (v3 + (v4 / (((v3 * v3)) | 1)))
    var v16 int32 = 1
    for fi17 := 0; fi17 < 20; fi17++ {
        switch "foo bar" {
        case "foo bar":
            v18 := GenS1{F0: "a", F1: "hello"}
            _ = v18
        case "\n\t":
            v4 += (v3 - v4)
        }
        _ = fi17
    }
    fmt.Println(v1)
    v19 := GenS1{F0: "\n\t", F1: "hello"}
    v4 -= (v4 * v3)
    for fi20 := 0; fi20 < 13; fi20++ {
        v21 := GenS2{F0: 42, F1: 255, F2: v4}
        _ = fi20
        _ = v21
    }
    for fi22 := 0; fi22 < 15; fi22++ {
        fmt.Println(v5)
        v2 = append(v2, ("hello" + ""))
        fmt.Println(v15)
        _ = fi22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
