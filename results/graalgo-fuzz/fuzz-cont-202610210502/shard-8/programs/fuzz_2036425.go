package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 uint16 }
type GenS2 struct { F0 string; F1 uint8; F2 bool }
func genFn0(p0 int64, p1 int8, p2 int64) bool {
    var v4 int16 = 0
    for fi5 := 0; fi5 < 2; fi5++ {
        var v6 float64 = 0.0
        v6 += ((0.1 * v6) / v6)
        v7 := []int{(fi5 ^ (fi5 & fi5)), -1}
        _ = fi5
        _ = v6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return false
}

func main() {
    var v1 int = 754
    var v2 string = ("\n\t" + ("foo bar" + ("hello" + "")))
    v3 := make(map[string]int)
    var v8 uint8 = 227
    var v9 int8 = 100
    v10 := GenS1{F0: 65535, F1: 686, F2: 1}
    v1 += v1
    fmt.Println(v2, v8)
    v1 += 32767
    delete(v3, v2)
    var v11 uint8 = v8
    var v12 int = v1
    switch v1 {
    case 1:
        v11 -= 8
    }
    switch (-(v12)) {
    case 32767:
        v13 := GenS1{F0: 0, F1: 0, F2: 42}
        _ = v13
    case 9223372036854775807:
        var v14 int = -646
        _ = v14
    default:
        fmt.Println(v2, v9)
    }
    v3[((v2 + "\n\t") + v2)] = v12
    if true {
        v2 = ("foo bar" + "a")
        v15 := GenS2{F0: (v2 + v2), F1: (v11 >> ((2) & 63)), F2: true}
        for fi16 := 0; fi16 < 14; fi16++ {
            var v17 bool = (!true)
            _ = fi16
            _ = v17
        }
        _ = v15
    } else {
        var v18 int8 = (v9 << ((65535) & 63))
        var v19 int32 = 33
        var v20 string = v2
        _ = v18
        _ = v19
        _ = v20
    }
    if ((589 < 42) && false) {
        var v21 uint32 = 653
        fmt.Println(v21, v8, v21)
        v22 := []int{v12, 127, v1}
        _ = v21
        _ = v22
    } else {
        if false {
            for fi23 := 0; fi23 < 11; fi23++ {
                v9 -= (^(v9))
                v24 := []int{-120, v1, (v12 / ((594) | 1))}
                _ = fi23
                _ = v24
            }
            for fi25 := 0; fi25 < 13; fi25++ {
                v2 = v2
                _ = fi25
            }
        } else {
            _ = len(v3)
            v26 := []int{}
            _ = v26
        }
        var v27 int = v1
        _ = v27
    }
    for fi28 := 0; fi28 < 4; fi28++ {
        v8 = v8
        var v29 int = v1
        _ = fi28
        _ = v29
    }
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
