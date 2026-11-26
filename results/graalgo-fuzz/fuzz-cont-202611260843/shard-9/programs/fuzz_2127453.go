package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 bool; F2 int16 }
type GenS3 struct { F0 int8; F1 string; F2 string }
func genFn0(p0 uint32, p1 int, p2 string) int {
    p2 = ("a" + ((((("" + p2) + p2) + (("a" + "\n\t") + ((p2 + ("hello" + ("\n\t" + (p2 + p2)))) + p2))) + p2) + p2))
    var v4 int32 = 42
    fmt.Println(p0)
    v5 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return (p1 ^ -542)
}
func genFn1(p0 uint64, p1 uint16, p2 bool) uint8 {
    v6 := GenS1{F0: 0}
    p1 = (^(p1))
    var v7 bool = (283 <= p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 1
}

func main() {
    var v1 int64 = 90
    v2 := []uint32{2147483647, 4294967295, 1}
    var v3 uint8 = 1
    _ = len(v2)
    for fi8 := 0; fi8 < 13; fi8++ {
        v3 += (v3 + (^(v3)))
        var v9 int16 = -1
        v10 := GenS2{F0: true, F1: (false || (!(!true))), F2: (v9 * v9)}
        _ = fi8
        _ = v10
        _ = v9
    }
    if false {
        v1 += v1
        v11 := GenS2{F0: (genFn0(4294967295, -444, ("" + ("hello" + ("a" + "")))) <= -1), F1: true, F2: 42}
        fmt.Println(v3, v3)
        _ = v11
    }
    v3 -= (v3 >> ((4294967295) & 63))
    v1 += 42
    v3 = v3
    for fi12 := 0; fi12 < 6; fi12++ {
        switch "" {
        case "hello":
            fmt.Println(v3, fi12, v3)
        case "\n\t":
            var v13 int64 = (v1 + 42)
            _ = v13
        default:
            var v14 float32 = 1.5
            _ = v14
        }
        _ = len(v2)
        _ = fi12
    }
    if true {
        var v15 bool = true
        var v16 uint64 = 0
        var v17 int = 1
        _ = v15
        _ = v16
        _ = v17
    }
    v1 -= v1
    v18 := GenS3{F0: -70, F1: "a", F2: "\n\t"}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
}
