package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 uint32 }
func genFn0(p0 bool) int32 {
    var v7 int64 = -1
    switch "hello" {
    case "a":
        fmt.Println(v7, p0)
    case "\n\t":
        for fi8 := 0; fi8 < 19; fi8++ {
            var v9 int = (fi8 + (fi8 + -201))
            _ = fi8
            _ = v9
        }
    }
    if (!(42 > 1)) {
        v7 -= v7
        var v10 float64 = (1e10 * (1e10 + (1.5 * 3.14159)))
        _ = v10
    }
    _ = p0
    _ = v7
    return 598
}
func genFn1(p0 int8, p1 int32) uint {
    if (false && false) {
        for fi11 := 0; fi11 < 9; fi11++ {
            var v12 int32 = p1
            p1 += (p1 + p1)
            v13 := []int{(-(fi11)), (^(fi11)), fi11, (fi11 * (fi11 / ((int(p1)) | 1)))}
            _ = fi11
            _ = v12
            _ = v13
        }
        v14 := []int{9223372036854775807, 639, 32767}
        _ = v14
    } else {
        if (("a" > (("\n\t" + "a") + "foo bar")) || (-255 <= p1)) {
            var v15 int32 = (p1 % (((p1 * genFn0((("\n\t" + "foo bar") != "a")))) | 1))
            _ = v15
        }
    }
    v16 := []int{397, 714}
    p1 = genFn0(false)
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = v16
    return 4294967295
}

func main() {
    var v1 int64 = -1
    var v2 float32 = 0.1
    v3 := make(map[int]int)
    var v4 int64 = (-(v1))
    var v5 float32 = -2.25
    var v6 int = int(v1)
    v2 = 3.14159
    v17 := GenS1{F0: 1000, F1: v6}
    v2 = ((-2.25 + 1e10) * float32((0.1 * float64(1e10))))
    for fi18 := 0; fi18 < 16; fi18++ {
        var v19 uint32 = 2
        v20 := GenS1{F0: 32767, F1: fi18}
        v6 = fi18
        _ = fi18
        _ = v19
        _ = v20
    }
    if false {
        var v21 uint8 = 223
        _ = v21
    } else {
        var v22 float64 = 3.14159
        v6 = v6
        if (517 > 2147483648) {
            v23 := GenS2{F0: 2}
            _ = v23
        }
        _ = v22
    }
    v1 += 816
    switch (-(v6)) {
    case 662:
        v24 := GenS2{F0: 0}
        _ = v24
    case -610:
        v6 += (-(v6))
    default:
        delete(v3, v6)
    }
    switch (v6 & 128) {
    case 127:
        fmt.Println(v2)
    case 128:
        var v25 uint32 = 697
        _ = v25
    case -501:
        var v26 uint32 = 2
        _ = v26
    }
    v5 -= (v2 - 0.1)
    for fi27 := 0; fi27 < 9; fi27++ {
        v28 := GenS1{F0: 1000, F1: v6}
        _ = fi27
        _ = v28
    }
    v29 := GenS2{F0: 317}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
