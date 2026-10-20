package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 bool }
type GenS2 struct { F0 bool; F1 uint64 }
type GenS3 struct { F0 uint32; F1 int16; F2 uint32 }

func main() {
    v1 := make(map[string]uint32)
    var v2 float64 = 3.14159
    v3 := make(map[uint32]string)
    var v4 int16 = 42
    if false {
        if (!(!true)) {
            switch -2147483648 {
            case 42:
                if (1 <= 18446744073709551615) {
                    v5 := []int{127, 452, -9223372036854775808}
                    var v6 int8 = 127
                    v7 := GenS1{F0: -528, F1: 1, F2: true}
                    _ = v5
                    _ = v6
                    _ = v7
                } else {
                    var v8 float32 = 3.14159
                    _ = v8
                }
            case -233:
                v9 := GenS1{F0: 42, F1: 2, F2: ((v2 / -2.25) > float64(1e10))}
                _ = v9
            default:
                v4 = v4
            }
            v10 := GenS2{F0: (!true), F1: 18446744073709551615}
            if false {
                if true {
                    v2 = (v2 + float64((0.1 * 1e10)))
                    for fi11 := 0; fi11 < 20; fi11++ {
                        v12 := GenS3{F0: 2147483648, F1: (v4 * -1), F2: 0}
                        var v13 float64 = v2
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                }
                for fi14 := 0; fi14 < 3; fi14++ {
                    for fi15 := 0; fi15 < 19; fi15++ {
                        fmt.Println(v4)
                        v4 = v4
                        v16 := []int{fi14, (fi14 / ((fi14) | 1))}
                        _ = fi15
                        _ = v16
                    }
                    _ = fi14
                }
                for fi17 := 0; fi17 < 9; fi17++ {
                    if false {
                        v18 := GenS1{F0: 2147483647, F1: 114, F2: false}
                        v19 := GenS3{F0: 2147483647, F1: v4, F2: 2}
                        _ = v18
                        _ = v19
                    } else {
                        v2 -= 0.1
                    }
                    _ = fi17
                }
            }
            _ = v10
        }
        var v20 int64 = -332
        _ = v20
    }
    var v21 float64 = 1.5
    var v22 uint64 = 0
    v23 := []int{32767}
    fmt.Println(v2, v21, v2)
    var v24 int16 = (v4 & (v4 / ((v4) | 1)))
    v25 := GenS3{F0: 4294967295, F1: v24, F2: 4294967295}
    fmt.Println(v22, v24, v4)
    var v26 int32 = 476
    if false {
        if (true && false) {
            var v27 uint64 = (v22 % (((^(v22))) | 1))
            if len(v23) > 0 {
                v23[0] = 0
            }
            _ = v27
        }
        v28 := v3[306]
        _ = v28
    }
    delete(v3, 4294967295)
    var v29 int64 = 0
    fmt.Println(v22, v21, v29)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v4
}
