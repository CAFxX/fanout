package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 uint64 }
func genFn0() int {
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 uint = 294
        switch 4294967295 {
        case -9223372036854775808:
            var v9 bool = (((float64((1.5 + float32((1.5 / 1.5)))) + float64((1.5 * 1.5))) - ((1e10 - ((((((float64((0.1 * 0.0)) / 1e10) - 0.1) - 0.1) + float64((float32(0.0) - 0.0))) * 1.5) + (float64(1.5) * (0.1 / 1e10)))) + float64(-2.25))) >= (1.5 + (1e10 + (1.5 / ((1e10 + 1e10) * (3.14159 / -2.25))))))
            _ = v9
        case 9223372036854775807:
            for fi10 := 0; fi10 < 7; fi10++ {
                switch fi10 {
                case 467:
                    var v11 int32 = -329
                    _ = v11
                default:
                    v8 = (v8 >> (((v8 / (((v8 & 1)) | 1))) & 63))
                }
                v8 += 0
                _ = fi10
            }
        case -259:
            v12 := []int{(fi7 - fi7), (fi7 % (((-(fi7))) | 1))}
            _ = v12
        }
        _ = fi7
        _ = v8
    }
    var v13 uint8 = 0
    _ = v13
    return -9223372036854775808
}
func genFn1() uint32 {
    for fi14 := 0; fi14 < 20; fi14++ {
        fmt.Println(fi14)
        if false {
            var v15 bool = ((fi14 * (fi14 ^ fi14)) > 32768)
            var v16 string = "foo bar"
            v17 := []int{(-(fi14))}
            _ = v15
            _ = v16
            _ = v17
        }
        v18 := []int{(fi14 % ((genFn0()) | 1)), (fi14 ^ genFn0()), fi14, (fi14 << ((255) & 63))}
        _ = fi14
        _ = v18
    }
    for fi19 := 0; fi19 < 11; fi19++ {
        var v20 int32 = -354
        v21 := []int{}
        v20 += v20
        _ = fi19
        _ = v20
        _ = v21
    }
    v22 := []int{-9223372036854775808, 0, -1}
    switch 128 {
    case 32768:
        var v23 bool = false
        _ = v23
    case -876:
        if (!(!(1 <= 209))) {
            var v24 uint64 = 42
            var v25 int16 = 1000
            v25 = 42
            _ = v24
            _ = v25
        }
    case 4294967295:
        for fi26 := 0; fi26 < 11; fi26++ {
            switch (fi26 ^ fi26) {
            case 0:
                var v27 float64 = 3.14159
                _ = v27
            case 32768:
                if false {
                    var v28 string = "hello"
                    if len(v22) >= 2 {
                        v22 = v22[1:len(v22)]
                    }
                    _ = v28
                }
            case -380:
                for fi29 := 0; fi29 < 19; fi29++ {
                    fmt.Println(fi29, fi29)
                    switch 9223372036854775807 {
                    case -2147483648:
                        var v30 uint64 = 243
                        _ = v30
                    case 4294967295:
                        var v31 int8 = -87
                        _ = v31
                    default:
                        _ = len(v22)
                    }
                    _ = fi29
                }
            default:
                v32 := GenS1{F0: 42, F1: 42, F2: 1}
                _ = v32
            }
            for fi33 := 0; fi33 < 17; fi33++ {
                if len(v22) > 0 {
                    v22[0] = 128
                }
                if len(v22) > 0 {
                    v22[0] = fi33
                }
                v34 := GenS1{F0: -1, F1: 1, F2: 574}
                _ = fi33
                _ = v34
            }
            var v35 uint8 = 172
            _ = fi26
            _ = v35
        }
    default:
        var v36 int = -9223372036854775808
        _ = v36
    }
    _ = v22
    return 65535
}

func main() {
    var v1 int32 = -858
    v2 := []uint32{0, 2147483648}
    var v3 int32 = v1
    var v4 float32 = 1.5
    v5 := make(map[int]string)
    v6 := make(map[uint32]bool)
    fmt.Println(v1)
    v1 = -1
    var v37 int64 = -9223372036854775808
    var v38 int64 = (v37 | v37)
    for fi39 := 0; fi39 < 14; fi39++ {
        var v40 uint32 = 0
        v41 := GenS1{F0: -128, F1: genFn1(), F2: 911}
        _ = fi39
        _ = v40
        _ = v41
    }
    var v42 int64 = v37
    v42 -= (v42 & (v38 / ((1) | 1)))
    if (0 == 750) {
        delete(v6, 65535)
        v38 -= v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
