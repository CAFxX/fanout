package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 int16; F2 int16 }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 bool; F1 uint32 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint8; F1 int32; F2 int8 }
func genFn0() int32 {
    v4 := GenS1{F0: false}
    var v5 int32 = 1
    for fi6 := 0; fi6 < 3; fi6++ {
        var v7 int64 = 0
        _ = fi6
        _ = v7
    }
    _ = v4
    _ = v5
    return v5
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 bool = (((float64((0.1 + 3.14159)) - 1.5) / 1e10) <= 0.0)
    var v3 uint64 = 18446744073709551615
    var v8 float64 = (-2.25 / 0.1)
    switch -860 {
    case 9223372036854775807:
        var v9 uint32 = 535
        _ = v9
    case 826:
        fmt.Println(v3, v2)
    case 951:
        var v10 uint8 = 1
        _ = v10
    default:
        for fi11 := 0; fi11 < 5; fi11++ {
            for fi12 := 0; fi12 < 17; fi12++ {
                var v13 uint8 = 2
                _ = fi12
                _ = v13
            }
            switch fi11 {
            case 128:
                for fi14 := 0; fi14 < 8; fi14++ {
                    var v15 int8 = 127
                    _ = fi14
                    _ = v15
                }
            case 2147483647:
                if (-66 > -308) {
                    for fi16 := 0; fi16 < 16; fi16++ {
                        var v17 int64 = 9223372036854775807
                        _ = fi16
                        _ = v17
                    }
                    var v18 int16 = -32768
                    v18 -= v18
                    _ = v18
                }
            default:
                switch fi11 {
                case 549:
                    v8 -= (float64(1e10) + ((((1.5 - ((float64(1.5) - 0.1) / 0.1)) / 1.0) * float64(1e10)) - v8))
                case 566:
                    var v19 int64 = 9223372036854775807
                    _ = v19
                case 1:
                    v3 -= (^(v3))
                default:
                    v20 := GenS1{F0: (!((669 <= 0) && (!(!(0 >= 228)))))}
                    _ = v20
                }
            }
            var v21 bool = (60 < 560)
            _ = fi11
            _ = v21
        }
    }
    switch -664 {
    case 2147483647:
        v8 = 0.1
    case 0:
        v22 := []int{32767}
        _ = v22
    case 4294967295:
        for fi23 := 0; fi23 < 6; fi23++ {
            v1[2] = (v2 || (!true))
            _ = fi23
        }
    }
    var v24 uint8 = 1
    v25 := GenS1{F0: v2}
    if (1 >= 928) {
        var v26 uint16 = 1
        if (42 <= 0) {
            v27 := GenS1{F0: (v2 || (0.1 == 0.1))}
            _ = v27
        } else {
            v8 += (v8 + v8)
            _ = len(v1)
        }
        _ = v26
    }
    switch -2147483648 {
    case -186:
        if false {
            v28 := []int{1, 0}
            _ = v28
        }
    default:
        v2 = true
    }
    v24 += v24
    for fi29 := 0; fi29 < 13; fi29++ {
        for fi30 := 0; fi30 < 14; fi30++ {
            if (1 <= 265) {
                v3 = v3
                switch 0 {
                case 9223372036854775807:
                    v31 := []int{fi29}
                    _ = v31
                case -588:
                    v32 := GenS1{F0: true}
                    _ = v32
                case 878:
                    v33 := GenS2{F0: 255, F1: 820, F2: -32768}
                    _ = v33
                }
                fmt.Println(v8)
                if v2 {
                    _ = len(v1)
                } else {
                    v34 := v1[2147483648]
                    v35 := GenS1{F0: v34}
                    _ = v34
                    _ = v35
                }
                v24 += 42
                var v36 int16 = 1
                var v37 int16 = (v36 | v36)
                var v38 uint = 1
                v39 := GenS3{F0: 0, F1: (v37 * v37)}
                if v2 {
                    v37 = -476
                    v37 -= (v36 >> (((v38 + (v38 << (((v38 ^ (v38 - v38))) & 63)))) & 63))
                } else {
                    v40 := GenS4{F0: (v2 && (!(v36 != v37))), F1: 2147483647}
                    fmt.Println(v3, v3, v24)
                    _ = v40
                }
                _ = v36
                _ = v37
                _ = v38
                _ = v39
            }
            switch "\n\t" {
            case "foo bar":
                v41 := GenS1{F0: v2}
                _ = v41
            case "\n\t":
                var v42 int32 = 42
                _ = v42
            default:
                var v43 float32 = 0.0
                _ = v43
            }
            var v44 uint32 = 65535
            var v45 int8 = -97
            var v46 uint64 = 4294967295
            _ = fi30
            _ = v44
            _ = v45
            _ = v46
        }
        for fi47 := 0; fi47 < 9; fi47++ {
            v48 := []int{fi29, (fi29 >> ((18446744073709551615) & 63)), (-(fi29))}
            _ = fi47
            _ = v48
        }
        _ = fi29
    }
    v49 := GenS3{F0: 2147483648, F1: 0}
    if (!v2) {
        v50 := GenS5{F0: 42}
        _ = v50
    }
    for fi51 := 0; fi51 < 6; fi51++ {
        fmt.Println(v2, v24, v24)
        var v52 bool = ("" <= ("" + ("hello" + ("a" + ""))))
        _ = fi51
        _ = v52
    }
    switch "a" {
    case "\n\t":
        var v53 string = ("" + (((("\n\t" + ("hello" + ("hello" + "hello"))) + "foo bar") + "a") + "foo bar"))
        _ = v53
    case "hello":
        var v54 uint16 = 255
        _ = v54
    }
    var v55 uint64 = (v3 % ((914) | 1))
    v56 := GenS2{F0: 65535, F1: -32768, F2: 1}
    switch 4294967295 {
    case -336:
        var v57 uint16 = 263
        _ = v57
    default:
        v58 := GenS6{F0: v24, F1: -516, F2: 42}
        _ = v58
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v55:", v55)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v49
    _ = v55
    _ = v56
    _ = v8
}
