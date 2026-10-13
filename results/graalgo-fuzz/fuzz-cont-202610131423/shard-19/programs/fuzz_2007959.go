package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 bool; F1 int; F2 string }
type GenS4 struct { F0 int; F1 int8 }
func genFn0() int32 {
    var v4 uint8 = 1
    if false {
        for fi5 := 0; fi5 < 1; fi5++ {
            v6 := GenS1{F0: (^(v4)), F1: -396, F2: fi5}
            var v7 uint64 = 0
            v8 := []int{-9223372036854775808, (^(fi5)), fi5}
            _ = fi5
            _ = v6
            _ = v7
            _ = v8
        }
    }
    for fi9 := 0; fi9 < 10; fi9++ {
        v10 := GenS2{F0: 9223372036854775807}
        v11 := GenS1{F0: (v4 >> ((42) & 63)), F1: 537, F2: 32767}
        for fi12 := 0; fi12 < 13; fi12++ {
            v4 = 0
            _ = fi12
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    v4 -= 106
    _ = v4
    return 1
}

func main() {
    var v1 uint = 65535
    var v2 int64 = 0
    v3 := make(map[int]string)
    v13 := GenS1{F0: 199, F1: -1, F2: 9223372036854775807}
    for fi14 := 0; fi14 < 4; fi14++ {
        v15 := GenS1{F0: 79, F1: int16(v2), F2: (fi14 / ((fi14) | 1))}
        var v16 int32 = -2147483648
        _ = fi14
        _ = v15
        _ = v16
    }
    fmt.Println(v2, v2)
    fmt.Println(v1, v2)
    v2 += 0
    v17 := []int{1, -409, -1}
    v2 += v2
    var v18 uint8 = 0
    v18 = v18
    var v19 uint16 = 65535
    for fi20 := 0; fi20 < 14; fi20++ {
        fmt.Println(v19, v2, v19)
        if false {
            for fi21 := 0; fi21 < 12; fi21++ {
                var v22 uint64 = 2
                var v23 uint = (^(v1))
                if ((fi21 + fi20) <= fi21) {
                    v24 := GenS2{F0: 0}
                    _ = v24
                }
                _ = fi21
                _ = v22
                _ = v23
            }
            v25 := GenS1{F0: (v18 & v18), F1: -676, F2: fi20}
            _ = v25
        } else {
            var v26 bool = true
            var v27 string = ""
            delete(v3, (fi20 - (fi20 << ((193) & 63))))
            _ = v26
            _ = v27
        }
        v1 -= (v1 ^ (v1 / (((^(v1))) | 1)))
        _ = fi20
    }
    v28 := GenS3{F0: (168 > 41), F1: 383, F2: ""}
    for fi29 := 0; fi29 < 2; fi29++ {
        v18 += (v18 * (v18 << (((v1 - (v1 - (v1 & 65535)))) & 63)))
        if len(v17) > 0 {
            v17[0] = -694
        }
        v2 += v2
        _ = fi29
    }
    for fi30 := 0; fi30 < 10; fi30++ {
        for fi31 := 0; fi31 < 15; fi31++ {
            v18 -= v18
            if len(v17) > 0 {
                v17[0] = 879
            }
            _ = fi31
        }
        v32 := GenS4{F0: -998, F1: 0}
        _ = fi30
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v13
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v28
    _ = v3
}
