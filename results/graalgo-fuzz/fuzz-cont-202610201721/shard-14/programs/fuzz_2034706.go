package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 int; F2 int64 }
func genFn0() string {
    for fi6 := 0; fi6 < 1; fi6++ {
        switch (fi6 | 32767) {
        case 947:
            var v7 int64 = 0
            _ = v7
        case -96:
            if false {
                var v8 uint8 = 114
                for fi9 := 0; fi9 < 7; fi9++ {
                    v8 += v8
                    if (-211 > -9223372036854775808) {
                        v10 := GenS1{F0: 743}
                        v8 = v8
                        _ = v10
                    }
                    var v11 int64 = 0
                    switch "hello" {
                    case "hello":
                        fmt.Println(fi6, fi6, v8)
                    case "":
                        switch -34 {
                        case 127:
                            v12 := GenS2{F0: 769, F1: fi6, F2: 42}
                            _ = v12
                        case -591:
                            var v13 float32 = (1.5 / (-2.25 - (((((float32((float64(1.5) - (1.5 * 0.0))) + 1e10) * (1.5 / 1.0)) / -2.25) + 0.1) + (0.1 - (-2.25 - 3.14159)))))
                            _ = v13
                        }
                    }
                    v8 -= (^(v8))
                    v8 = (v8 / ((v8) | 1))
                    _ = fi9
                    _ = v11
                }
                _ = v8
            }
        }
        if false {
            switch fi6 {
            case 32767:
                var v14 int = -9223372036854775808
                _ = v14
            case 1:
                var v15 uint16 = 32767
                _ = v15
            case 127:
                v16 := []int{fi6, -915, fi6}
                _ = v16
            }
            var v17 uint64 = 824
            if (2 < 255) {
                v17 -= (v17 * (v17 << ((693) & 63)))
                if true {
                    v18 := []int{fi6, fi6, (fi6 << ((2) & 63))}
                    v17 = v17
                    _ = v18
                } else {
                    var v19 float32 = (0.1 - 1.5)
                    v17 += 282
                    _ = v19
                }
            }
            _ = v17
        }
        fmt.Println(fi6)
        _ = fi6
    }
    if (!false) {
        if false {
            var v20 int8 = 100
            v20 = (v20 << ((4294967295) & 63))
            _ = v20
        } else {
            var v21 float64 = -2.25
            var v22 bool = ((0 <= 255) && ((!false) && false))
            _ = v21
            _ = v22
        }
    }
    var v23 int16 = 1000
    _ = v23
    return "foo bar"
}

func main() {
    var v1 uint8 = 42
    var v2 int64 = -183
    v3 := make(map[uint32]uint32)
    var v4 int16 = -986
    v5 := make(map[uint32]uint32)
    v24 := []int{0, -2147483648, -709}
    var v25 uint8 = v1
    fmt.Println(v2, v25)
    if len(v24) > 0 {
        v24[0] = 127
    }
    var v26 int32 = 989
    v27 := GenS2{F0: 381, F1: 9223372036854775807, F2: v2}
    var v28 uint32 = 4294967295
    _ = len(v5)
    delete(v3, 0)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
