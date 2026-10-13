package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 string; F2 uint32 }
type GenS3 struct { F0 int16; F1 uint64; F2 int64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int64; F1 uint64 }
func genFn0(p0 uint8, p1 int8, p2 bool) int8 {
    v6 := GenS1{F0: p1}
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return p1
}
func genFn1(p0 bool) uint16 {
    switch "a" {
    case "foo bar":
        var v7 int32 = 0
        _ = v7
    case "a":
        v8 := GenS2{F0: 32767, F1: "foo bar", F2: 411}
        _ = v8
    default:
        v9 := []int{-1}
        _ = v9
    }
    p0 = false
    p0 = (!true)
    v10 := GenS3{F0: 32767, F1: 0, F2: -1}
    _ = p0
    _ = v10
    return 2
}

func main() {
    var v1 int64 = -26
    var v2 float32 = -2.25
    v3 := make(map[uint32]int)
    var v4 int32 = -1
    var v5 uint8 = 255
    var v11 uint32 = 1
    var v12 uint32 = (v11 + (v11 % (((v11 >> ((uint(v5)) & 63))) | 1)))
    v13 := GenS4{F0: 42}
    switch 42 {
    case 695:
        var v14 uint8 = (v5 >> ((1) & 63))
        _ = v14
    default:
        v15 := v3[v11]
        _ = v15
    }
    var v16 uint16 = genFn1(true)
    if false {
        for fi17 := 0; fi17 < 1; fi17++ {
            var v18 int = fi17
            _ = fi17
            _ = v18
        }
        switch 9223372036854775807 {
        case -2147483648:
            var v19 uint32 = v11
            _ = v19
        }
        v1 += (-(v1))
    } else {
        var v20 int16 = 1000
        _ = v20
    }
    v1 -= v1
    v21 := GenS5{F0: v1, F1: 892}
    _ = len(v3)
    var v22 uint64 = 1
    v11 -= v12
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
