package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint32; F1 int32; F2 int32 }
type GenS5 struct { F0 int; F1 int32; F2 int32 }
func genFn0(p0 uint32) int32 {
    var v4 uint32 = 844
    if (!(440 >= 42)) {
        v5 := GenS1{F0: 1000, F1: 42}
        p0 = v4
        _ = v5
    }
    switch (("foo bar" + "foo bar") + ("hello" + (("hello" + "\n\t") + "foo bar"))) {
    case "":
        v6 := GenS2{F0: -943, F1: (!(896 < 803))}
        _ = v6
    case "\n\t":
        p0 -= (p0 | (v4 * 1))
    case "hello":
        p0 -= (^(v4))
    }
    v7 := []int{128, 42, 2147483647}
    var v8 uint32 = (v4 * 42)
    _ = p0
    _ = v4
    _ = v7
    _ = v8
    return 2147483647
}

func main() {
    var v1 uint64 = 583
    v2 := make(map[int]int)
    v3 := make(map[string]uint32)
    for fi9 := 0; fi9 < 1; fi9++ {
        for fi10 := 0; fi10 < 4; fi10++ {
            var v11 int64 = 9223372036854775807
            v12 := []int{(fi10 | int(v1))}
            v11 = (v11 / (((v11 + (v11 ^ (v11 * (v11 << ((2) & 63)))))) | 1))
            _ = fi10
            _ = v11
            _ = v12
        }
        v13 := []int{}
        _ = fi9
        _ = v13
    }
    var v14 int = 32768
    var v15 int = 0
    var v16 uint64 = 3
    v1 += v16
    fmt.Println(v15, v16)
    if true {
        v15 = v15
        v17 := GenS3{F0: int32(v1)}
        _ = v17
    } else {
        v18 := GenS3{F0: 1}
        var v19 float64 = (0.1 - -2.25)
        v20 := GenS4{F0: 2147483648, F1: 2147483647, F2: 565}
        _ = v18
        _ = v19
        _ = v20
    }
    fmt.Println(v16, v14)
    v14 = 736
    switch (v15 * v15) {
    case 9223372036854775807:
        v16 += v1
    case -1:
        fmt.Println(v1)
    case 329:
        v21 := GenS5{F0: 9223372036854775807, F1: -217, F2: genFn0(442)}
        _ = v21
    default:
        v1 -= 244
    }
    v22 := GenS5{F0: (v14 * (^(v14))), F1: -1, F2: 334}
    if true {
        var v23 int = v14
        var v24 uint64 = v1
        v25 := []int{(v14 + v14)}
        _ = v23
        _ = v24
        _ = v25
    }
    v14 += v14
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v22
    _ = v3
}
