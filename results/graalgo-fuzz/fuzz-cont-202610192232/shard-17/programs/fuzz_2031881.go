package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 int }
type GenS2 struct { F0 int64; F1 int16 }
func genFn0(p0 int) int16 {
    for fi7 := 0; fi7 < 6; fi7++ {
        v8 := GenS1{F0: -282, F1: "foo bar", F2: (^(fi7))}
        v9 := GenS2{F0: 1, F1: -1}
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := []int{p0, 128}
    _ = p0
    _ = v10
    return 1000
}
func genFn1(p0 uint64, p1 uint64) int64 {
    switch -9223372036854775808 {
    case 368:
        v11 := GenS1{F0: 0, F1: ("hello" + "foo bar"), F2: -157}
        _ = v11
    case 0:
        p0 = 18446744073709551615
    case 128:
        var v12 string = "hello"
        _ = v12
    default:
        p1 = 42
    }
    v13 := []int{795, 178}
    _ = p0
    _ = p1
    _ = v13
    return -9223372036854775808
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]bool)
    var v3 int64 = 42
    var v4 int64 = (v3 | v3)
    v5 := make(map[int]string)
    var v6 int16 = 1000
    var v14 int8 = 100
    v15 := GenS2{F0: (-(v3)), F1: (v6 & genFn0(-113))}
    var v16 int64 = 42
    var v17 int8 = v14
    for fi18 := 0; fi18 < 11; fi18++ {
        var v19 float32 = (((-2.25 / ((-2.25 - float32((3.14159 + -2.25))) / 1.0)) / 1.5) + (0.0 - 0.1))
        _ = fi18
        _ = v19
    }
    v20 := GenS1{F0: 805, F1: "hello", F2: 9223372036854775807}
    var v21 int32 = -291
    var v22 float64 = float64(0.1)
    var v23 int32 = (v21 % (((^(v21))) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
