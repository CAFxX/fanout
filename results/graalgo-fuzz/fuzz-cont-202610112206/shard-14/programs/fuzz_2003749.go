package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 uint16 }
type GenS2 struct { F0 uint32; F1 uint64; F2 uint8 }
type GenS3 struct { F0 int32 }
func genFn0(p0 uint32, p1 bool, p2 uint8) int {
    v4 := []int{4294967295, 32768, -1, -496}
    fmt.Println(p2, p0, p2)
    if len(v4) > 0 {
        v4[0] = 128
    }
    if false {
        var v5 float64 = 0.1
        if true {
            var v6 string = (("foo bar" + "hello") + "hello")
            if p1 {
                v7 := GenS1{F0: uint(p2), F1: v6, F2: 990}
                var v8 float64 = (v5 - (v5 / 0.1))
                var v9 int16 = 0
                _ = v7
                _ = v8
                _ = v9
            }
            _ = v6
        }
        v10 := GenS1{F0: 1, F1: "a", F2: 2}
        _ = v10
        _ = v5
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 0
}

func main() {
    var v1 int32 = 0
    var v2 int = -454
    v3 := make(map[int]string)
    switch ("a" + "\n\t") {
    case "\n\t":
        var v11 int8 = -1
        _ = v11
    case "":
        v2 = v2
    default:
        for fi12 := 0; fi12 < 13; fi12++ {
            for fi13 := 0; fi13 < 11; fi13++ {
                v14 := GenS2{F0: 166, F1: 18446744073709551615, F2: 127}
                _ = fi13
                _ = v14
            }
            var v15 uint16 = 42
            _ = fi12
            _ = v15
        }
    }
    for fi16 := 0; fi16 < 12; fi16++ {
        v2 -= (v2 % ((v2) | 1))
        _ = fi16
    }
    var v17 uint32 = 2
    v18 := GenS2{F0: (v17 % ((v17) | 1)), F1: 974, F2: 11}
    v19 := GenS2{F0: (v17 / ((1) | 1)), F1: 4294967295, F2: 112}
    switch 32768 {
    case 780:
        var v20 uint16 = 65535
        _ = v20
    case -1:
        v21 := GenS2{F0: v17, F1: 137, F2: 0}
        _ = v21
    default:
        v22 := GenS1{F0: 149, F1: "\n\t", F2: 0}
        _ = v22
    }
    fmt.Println(v17)
    fmt.Println(v1)
    for fi23 := 0; fi23 < 18; fi23++ {
        v24 := GenS3{F0: (v1 * (v1 * (v1 / ((42) | 1))))}
        _ = fi23
        _ = v24
    }
    var v25 float32 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v25
    _ = v3
}
