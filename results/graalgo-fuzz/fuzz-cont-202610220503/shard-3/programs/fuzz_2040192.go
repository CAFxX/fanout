package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 int64 }
type GenS2 struct { F0 int; F1 int8; F2 uint }
func genFn0(p0 int64, p1 uint16) int16 {
    v5 := []int{}
    p1 = (p1 << ((65535) & 63))
    v6 := GenS1{F0: "", F1: 979, F2: p0}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return -32768
}

func main() {
    var v1 float32 = -2.25
    var v2 int64 = 837
    v3 := make(map[uint32]uint32)
    v4 := make(map[int]bool)
    if false {
        v7 := GenS2{F0: 0, F1: 0, F2: 0}
        _ = v7
    } else {
        v8 := GenS2{F0: -727, F1: 1, F2: 451}
        v1 = (float32(float64((v1 - 0.1))) - v1)
        _ = v8
    }
    for fi9 := 0; fi9 < 2; fi9++ {
        v10 := GenS2{F0: fi9, F1: 127, F2: 4294967295}
        switch fi9 {
        case 127:
            switch (fi9 / (((fi9 | (fi9 ^ (fi9 - (fi9 << ((76) & 63)))))) | 1)) {
            case -254:
                v2 -= (v2 << ((255) & 63))
            case 9223372036854775807:
                delete(v4, fi9)
            case -2147483648:
                var v11 uint32 = 688
                _ = v11
            }
        case 0:
            fmt.Println(v2, v1)
        case 32767:
            switch fi9 {
            case 0:
                fmt.Println(fi9)
            default:
                v2 += v2
            }
        }
        v2 += (v2 - v2)
        _ = fi9
        _ = v10
    }
    switch (("hello" + ("hello" + ("\n\t" + "foo bar"))) + "") {
    case "\n\t":
        fmt.Println(v1, v1)
    case "foo bar":
        var v12 int32 = 2147483647
        _ = v12
    default:
        v13 := GenS1{F0: "a", F1: 4294967295, F2: 485}
        _ = v13
    }
    var v14 uint32 = 0
    v1 -= ((-2.25 * 0.1) + (v1 - (v1 * v1)))
    v15 := GenS1{F0: (((("\n\t" + ("hello" + ("a" + "foo bar"))) + "foo bar") + "foo bar") + ("hello" + "foo bar")), F1: v14, F2: 702}
    v14 -= 42
    var v16 int16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
