package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 uint8; F1 int8; F2 uint16 }
type GenS5 struct { F0 int32; F1 uint16; F2 uint32 }
func genFn0(p0 uint) int32 {
    v4 := []int{128, 0, -802, 32767}
    v5 := GenS1{F0: "hello", F1: 1}
    _ = p0
    _ = v4
    _ = v5
    return -254
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int = 42
    var v3 uint64 = 42
    if true {
        fmt.Println(v2, v3)
        for fi6 := 0; fi6 < 8; fi6++ {
            v2 = (fi6 + (v2 ^ (fi6 * (fi6 << ((18446744073709551615) & 63)))))
            switch 2147483647 {
            case -378:
                switch fi6 {
                case 2147483647:
                    fmt.Println(fi6, v2)
                }
            case 128:
                v7 := GenS2{F0: -255, F1: (!(3.14159 > (0.1 / 1.0)))}
                _ = v7
            case 42:
                if false {
                    v8 := []int{32768, (^(fi6)), -899}
                    fmt.Println(v3, fi6)
                    _ = v8
                }
            default:
                var v9 float32 = 1e10
                _ = v9
            }
            for fi10 := 0; fi10 < 12; fi10++ {
                switch (v2 / ((458) | 1)) {
                case 4294967295:
                    var v11 int64 = 9223372036854775807
                    v11 = 9223372036854775807
                    v12 := v1[v2]
                    _ = v11
                    _ = v12
                case 128:
                    if (65535 < 18446744073709551615) {
                        var v13 float32 = 0.1
                        _ = v13
                    } else {
                        var v14 uint8 = 210
                        _ = v14
                    }
                    var v15 float32 = 1.5
                    _ = v15
                case -2147483648:
                    v2 = 127
                }
                switch fi6 {
                case -569:
                    if false {
                        var v16 uint8 = 0
                        v17 := GenS2{F0: -44, F1: true}
                        _ = v16
                        _ = v17
                    }
                case 32767:
                    switch -65 {
                    case -2147483648:
                        var v18 uint16 = 538
                        _ = v18
                    case 612:
                        v19 := GenS3{F0: 2, F1: -673}
                        _ = v19
                    default:
                        v2 = (fi10 / (((fi6 + fi10)) | 1))
                    }
                default:
                    v20 := GenS4{F0: 22, F1: 1, F2: 2}
                    v21 := GenS3{F0: 4294967295, F1: 1000}
                    _ = v20
                    _ = v21
                }
                if (39 <= 127) {
                    v22 := []int{-9223372036854775808}
                    fmt.Println(fi10, fi10, v3)
                    _ = v22
                }
                _ = fi10
            }
            _ = fi6
        }
    } else {
        switch "a" {
        case "":
            fmt.Println(v2)
        case "a":
            var v23 uint64 = (v3 >> ((214) & 63))
            _ = v23
        default:
            v2 = v2
        }
        v24 := GenS4{F0: 156, F1: 63, F2: 881}
        _ = v24
    }
    if false {
        v3 -= (v3 << ((4294967295) & 63))
    }
    v25 := v1[(^(v2))]
    if true {
        var v26 uint = 4294967295
        _ = v26
    }
    v1[(^(v2))] = (^(v25))
    v27 := GenS5{F0: 599, F1: 2, F2: (^(v25))}
    if true {
        var v28 float64 = -2.25
        switch "\n\t" {
        case "a":
            var v29 float64 = v28
            _ = v29
        }
        v1[(v2 >> ((2) & 63))] = v25
        _ = v28
    }
    v30 := GenS4{F0: 42, F1: 1, F2: 0}
    v3 = (^(v3))
    var v31 bool = (v25 < (v25 / (((^(v25))) | 1)))
    for fi32 := 0; fi32 < 3; fi32++ {
        v33 := GenS2{F0: -1, F1: ((v31 && v31) || (v31 && (true && (int16(v25) > 1000))))}
        _ = fi32
        _ = v33
    }
    var v34 uint8 = 148
    v2 = v2
    v2 = (v2 % ((-209) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v25
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v34
}
