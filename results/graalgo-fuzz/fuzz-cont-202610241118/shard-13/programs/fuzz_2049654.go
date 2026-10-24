package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 int }
type GenS2 struct { F0 bool; F1 string; F2 uint16 }
func genFn0(p0 uint, p1 uint64) string {
    if (-1 >= 0) {
        if (1 <= 1) {
            v6 := GenS1{F0: 417, F1: -1, F2: 848}
            _ = v6
        }
    } else {
        var v7 int8 = -13
        switch 0 {
        case 128:
            var v8 int16 = 42
            _ = v8
        case 2147483647:
            v9 := GenS1{F0: 9223372036854775807, F1: 55, F2: -19}
            _ = v9
        default:
            p1 = p1
        }
        _ = v7
    }
    for fi10 := 0; fi10 < 6; fi10++ {
        var v11 bool = true
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 9; fi12++ {
        p1 = p1
        _ = fi12
    }
    for fi13 := 0; fi13 < 16; fi13++ {
        fmt.Println(fi13)
        _ = fi13
    }
    p1 = p1
    _ = p0
    _ = p1
    return "foo bar"
}

func main() {
    var v1 int64 = 9223372036854775807
    v2 := []uint32{42}
    v3 := []string{(("" + ("\n\t" + "hello")) + ""), "foo bar", (("foo bar" + "a") + ""), ("" + "a")}
    v4 := make(map[int]int)
    var v5 uint = 255
    v1 = 0
    var v14 int16 = 436
    if true {
        var v15 uint16 = 553
        _ = v15
    }
    fmt.Println(v5)
    var v16 uint64 = 527
    _ = len(v2)
    if ((v14 + v14) >= (v14 & (^(v14)))) {
        v14 = (v14 / ((v14) | 1))
        if len(v3) > 0 {
            v3[0] = ("a" + "\n\t")
        }
    } else {
        var v17 uint8 = 128
        _ = v17
    }
    v18 := GenS2{F0: (true || (!true)), F1: "a", F2: 257}
    for fi19 := 0; fi19 < 2; fi19++ {
        v20 := GenS2{F0: true, F1: "", F2: uint16(v16)}
        var v21 int16 = -982
        v22 := GenS2{F0: true, F1: "hello", F2: 42}
        _ = fi19
        _ = v20
        _ = v21
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
