package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 uint }
type GenS2 struct { F0 int; F1 int; F2 int64 }
func genFn0() uint64 {
    if (!(0 != 65535)) {
        for fi6 := 0; fi6 < 9; fi6++ {
            for fi7 := 0; fi7 < 17; fi7++ {
                var v8 uint32 = 0
                v9 := []int{(-(fi7)), fi7}
                _ = fi7
                _ = v8
                _ = v9
            }
            v10 := []int{(-(fi6))}
            fmt.Println(fi6)
            _ = fi6
            _ = v10
        }
    } else {
        v11 := []int{-1}
        for fi12 := 0; fi12 < 18; fi12++ {
            var v13 uint32 = 1
            switch -687 {
            case -1:
                if len(v11) > 0 {
                    v11[0] = 128
                }
            default:
                fmt.Println(v13, v13)
            }
            fmt.Println(fi12, fi12)
            _ = fi12
            _ = v13
        }
        var v14 uint32 = 0
        _ = v11
        _ = v14
    }
    v15 := GenS1{F0: true, F1: 18446744073709551615, F2: 18446744073709551615}
    var v16 float64 = float64(0.1)
    _ = v15
    _ = v16
    return 4294967295
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]string)
    var v3 uint16 = 195
    v4 := []uint32{65535, 42, 99}
    v5 := []int{127, 128, 4294967295, 119}
    var v17 int16 = 32767
    v18 := GenS1{F0: ((false || (false && (0 != 42))) || true), F1: 0, F2: 4294967295}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if false {
        var v19 int = 9223372036854775807
        _ = v19
    } else {
        for fi20 := 0; fi20 < 2; fi20++ {
            var v21 int8 = 42
            var v22 float64 = -2.25
            var v23 uint8 = 165
            _ = fi20
            _ = v21
            _ = v22
            _ = v23
        }
    }
    for fi24 := 0; fi24 < 7; fi24++ {
        v17 = (v17 + (-(v17)))
        var v25 bool = (true && false)
        _ = fi24
        _ = v25
    }
    v3 += (^(v3))
    v26 := GenS2{F0: 2147483647, F1: -987, F2: 1}
    v27 := v2[42]
    v1[128] = 2147483648
    v17 += (v17 << ((916) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
