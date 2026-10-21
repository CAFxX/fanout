package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 uint64; F1 uint16 }
func genFn0(p0 string, p1 int16) int32 {
    v6 := GenS1{F0: "a", F1: -1}
    v7 := GenS1{F0: ("hello" + (((((p0 + p0) + (p0 + p0)) + p0) + p0) + (((p0 + p0) + p0) + p0))), F1: -1}
    v8 := []int{777, 128}
    if false {
        for fi9 := 0; fi9 < 8; fi9++ {
            var v10 int16 = p1
            _ = fi9
            _ = v10
        }
    } else {
        p1 = p1
        fmt.Println(p0, p0)
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return -2147483648
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]uint32)
    var v3 bool = (1 != 1)
    v4 := make(map[uint32]uint32)
    var v5 uint8 = 127
    v11 := v2[371]
    v3 = (uint64(v5) >= 42)
    for fi12 := 0; fi12 < 5; fi12++ {
        if (0 < 42) {
            v13 := []int{(fi12 / ((fi12) | 1))}
            _ = v13
        } else {
            if v3 {
                fmt.Println(v3)
                if false {
                    var v14 int64 = 9223372036854775807
                    fmt.Println(fi12, v5)
                    _ = v14
                }
            } else {
                v3 = (((42 != 1) || false) || v3)
                v15 := GenS1{F0: "foo bar", F1: -93}
                _ = len(v1)
                _ = v15
            }
        }
        if ((127 < -29) || (!(-1 != 35))) {
            v16 := GenS2{F0: 331, F1: 32767}
            switch fi12 {
            case 0:
                var v17 float32 = 1.5
                _ = v17
            default:
                var v18 uint16 = 2
                _ = v18
            }
            if (-9223372036854775808 >= 9223372036854775807) {
                var v19 uint8 = (^(v5))
                _ = v19
            }
            _ = v16
        } else {
            var v20 bool = (v5 == v5)
            v20 = false
            for fi21 := 0; fi21 < 8; fi21++ {
                v3 = (32767 > 2)
                _ = fi21
            }
            _ = v20
        }
        for fi22 := 0; fi22 < 15; fi22++ {
            if v3 {
                v5 -= 66
            }
            _ = fi22
        }
        _ = fi12
    }
    v3 = (100 != 80)
    v3 = v3
    v3 = (!(1 < 9223372036854775807))
    v11 += (^(v11))
    var v23 int32 = -1
    var v24 bool = (0 > -128)
    var v25 uint32 = v11
    fmt.Println(v11, v23, v3)
    v24 = (1 < 439)
    var v26 float64 = 3.14159
    v27 := []int{9223372036854775807}
    fmt.Println(v5)
    v11 = v11
    v28 := GenS1{F0: (("" + "hello") + "a"), F1: 42}
    for fi29 := 0; fi29 < 11; fi29++ {
        v26 -= v26
        v5 -= 255
        _ = fi29
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
