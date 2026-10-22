package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 uint64; F1 uint }
func genFn0(p0 int16, p1 uint64) uint8 {
    var v6 uint = 4294967295
    for fi7 := 0; fi7 < 9; fi7++ {
        if (42 <= -1) {
            var v8 int16 = -288
            v9 := GenS1{F0: 747, F1: 1}
            _ = v8
            _ = v9
        } else {
            p1 = 4294967295
        }
        _ = fi7
    }
    if (143 != 2) {
        p1 += 0
        var v10 int16 = (p0 ^ -400)
        _ = v10
    } else {
        switch int(v6) {
        case -730:
            fmt.Println(p1)
        default:
            var v11 string = "foo bar"
            _ = v11
        }
        v12 := GenS1{F0: 32767, F1: 255}
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v6
    return 1
}
func genFn1(p0 string, p1 int64, p2 int8) bool {
    var v13 int16 = 32767
    var v14 uint32 = 111
    fmt.Println(v13, v13, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    return true
}

func main() {
    var v1 int16 = 0
    var v2 uint64 = 4294967295
    var v3 int = -696
    v4 := make(map[string]bool)
    var v5 float32 = 0.1
    var v15 int8 = 42
    if true {
        v5 += 0.0
    }
    fmt.Println(v5)
    v2 = (^(v2))
    v2 = (v2 >> ((905) & 63))
    _ = len(v4)
    if true {
        switch ("" + "") {
        case "hello":
            fmt.Println(v5, v1)
        case "\n\t":
            v16 := GenS1{F0: 255, F1: 2}
            _ = v16
        }
        v17 := []int{v3, (v3 + (^(v3)))}
        v2 += (^(v2))
        _ = v17
    }
    var v18 int32 = 624
    if ((v5 / ((v5 * v5) * v5)) < (v5 - (v5 + v5))) {
        var v19 float64 = 1.5
        _ = v19
    }
    v20 := GenS2{F0: v2, F1: 741}
    for fi21 := 0; fi21 < 19; fi21++ {
        fmt.Println(v18, v2, fi21)
        _ = fi21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
