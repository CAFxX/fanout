package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 uint }
func genFn0(p0 bool, p1 bool, p2 uint) int32 {
    v7 := GenS1{F0: -330}
    if p1 {
        v8 := GenS2{F0: (^(p2))}
        fmt.Println(p2)
        v9 := GenS3{F0: 0, F1: ""}
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -2147483648
}

func main() {
    v1 := make(map[string]bool)
    var v2 int32 = -179
    v3 := make(map[string]bool)
    var v4 bool = ("hello" <= (("foo bar" + "\n\t") + ("foo bar" + "a")))
    var v5 int64 = 1
    var v6 int = 491
    v10 := GenS3{F0: 127, F1: ("hello" + ("" + "foo bar"))}
    v11 := []int{v6}
    fmt.Println(v4, v2, v6)
    v2 -= v2
    v12 := v1[""]
    if ((v5 > v5) && v12) {
        _ = len(v11)
        var v13 uint32 = 42
        if (!false) {
            var v14 uint32 = 734
            var v15 int = (-(v6))
            v16 := GenS1{F0: v5}
            _ = v14
            _ = v15
            _ = v16
        }
        _ = v13
    }
    if v12 {
        v2 = 132
        fmt.Println(v5)
    } else {
        var v17 uint8 = 1
        v17 -= 1
        _ = v17
    }
    _ = len(v3)
    v5 = v5
    if ((!(359 != 0)) && (false && (v12 && ((-128 < 1) && ((v6 ^ 42) < (v6 * v6)))))) {
        for fi18 := 0; fi18 < 18; fi18++ {
            var v19 float64 = 0.1
            for fi20 := 0; fi20 < 11; fi20++ {
                var v21 uint16 = 1
                fmt.Println(v2)
                v12 = (uint32(v21) < 918)
                _ = fi20
                _ = v21
            }
            fmt.Println(v6)
            _ = fi18
            _ = v19
        }
    }
    if v4 {
        for fi22 := 0; fi22 < 13; fi22++ {
            v12 = v12
            _ = fi22
        }
        var v23 int8 = 96
        v24 := GenS4{F0: 255}
        _ = v23
        _ = v24
    }
    _ = len(v11)
    var v25 int = (v6 << ((225) & 63))
    v2 += -2147483648
    if (372 > 4294967295) {
        fmt.Println(v2)
    } else {
        var v26 uint16 = 1
        fmt.Println(v6)
        _ = v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
