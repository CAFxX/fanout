package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int64; F1 uint8 }
type GenS6 struct { F0 bool; F1 int16 }
func genFn0() int16 {
    v5 := GenS1{F0: (true && ((false && false) || false)), F1: -707}
    switch "\n\t" {
    case "\n\t":
        var v6 uint8 = 128
        _ = v6
    case "a":
        switch 1 {
        case 128:
            var v7 uint = 0
            _ = v7
        }
    case "":
        for fi8 := 0; fi8 < 3; fi8++ {
            if false {
                var v9 float64 = ((-2.25 / (((1e10 / float64(1.5)) - 0.1) / 1.0)) + 3.14159)
                _ = v9
            } else {
                fmt.Println(fi8)
                for fi10 := 0; fi10 < 8; fi10++ {
                    switch fi10 {
                    case 1:
                        fmt.Println(fi10, fi10)
                    case 127:
                        v11 := []int{(-(fi8)), (fi10 * (^(fi10))), (^(fi10)), (fi10 + (fi10 >> ((658) & 63)))}
                        _ = v11
                    default:
                        if (false || (1 >= 4294967295)) {
                            switch fi8 {
                            case 9223372036854775807:
                                var v12 uint = 870
                                _ = v12
                            default:
                                var v13 float64 = 0.1
                                _ = v13
                            }
                        }
                        var v14 uint32 = 4294967295
                        var v15 int8 = 1
                        _ = v14
                        _ = v15
                    }
                    _ = fi10
                }
                var v16 uint16 = 32767
                _ = v16
            }
            var v17 uint8 = 7
            for fi18 := 0; fi18 < 9; fi18++ {
                v19 := GenS1{F0: true, F1: 0}
                _ = fi18
                _ = v19
            }
            _ = fi8
            _ = v17
        }
    }
    var v20 uint64 = 4294967295
    v20 = v20
    _ = v20
    _ = v5
    return 1000
}

func main() {
    var v1 float32 = float32(((1e10 + ((0.1 - (1e10 * (1.5 * float64(float32((((((1.5 * 0.1) - (1e10 + 1.5)) / 1e10) / 1.0) - 0.0)))))) - (3.14159 + (3.14159 - 1e10)))) * (0.0 - (3.14159 + (1.5 - 3.14159)))))
    v2 := make(map[int]string)
    var v3 int = 68
    v4 := make(map[string]string)
    v21 := GenS2{F0: 42, F1: 4294967295}
    var v22 int = -2147483648
    if false {
        v23 := GenS1{F0: (((1e10 - 0.0) - 1e10) <= ((1e10 / 1.5) * 1.5)), F1: -2147483648}
        switch -1 {
        case 32767:
            var v24 int64 = -634
            _ = v24
        case -2147483648:
            v1 = (float32((((0.0 + 0.1) + (((-2.25 + 3.14159) - (1.5 + ((0.0 / 1.5) / (1.5 + -2.25)))) - ((3.14159 / float64(1e10)) / (-2.25 - 1e10)))) + 3.14159)) * v1)
        case 0:
            var v25 float64 = 0.1
            _ = v25
        }
        _ = v23
    }
    fmt.Println(v3)
    fmt.Println(v3)
    v22 = v3
    v26 := GenS3{F0: false}
    if false {
        fmt.Println(v22, v22, v1)
        v27 := GenS4{F0: "a"}
        _ = v27
    } else {
        v3 = -588
        if (false || true) {
            var v28 uint8 = 155
            _ = v28
        }
        switch "\n\t" {
        case "hello":
            for fi29 := 0; fi29 < 17; fi29++ {
                for fi30 := 0; fi30 < 6; fi30++ {
                    if false {
                        var v31 float32 = float32((0.0 - 0.0))
                        _ = v31
                    }
                    v32 := GenS1{F0: true, F1: -1}
                    v33 := GenS5{F0: 0, F1: 208}
                    _ = fi30
                    _ = v32
                    _ = v33
                }
                v2[32767] = "a"
                if true {
                    var v34 int16 = 138
                    v35 := GenS6{F0: true, F1: v34}
                    v1 -= -2.25
                    _ = v34
                    _ = v35
                } else {
                    v22 = (v3 + -354)
                    for fi36 := 0; fi36 < 3; fi36++ {
                        v37 := GenS5{F0: -1, F1: 42}
                        switch fi29 {
                        case -720:
                            v38 := GenS5{F0: -73, F1: 107}
                            _ = v38
                        case -1:
                            switch 9223372036854775807 {
                            case 522:
                                var v39 uint8 = 128
                                _ = v39
                            case 738:
                                var v40 bool = false
                                _ = v40
                            case 812:
                                v41 := GenS5{F0: 1, F1: 2}
                                _ = v41
                            }
                        default:
                            v1 -= -2.25
                        }
                        var v42 uint64 = 4294967295
                        _ = fi36
                        _ = v37
                        _ = v42
                    }
                    fmt.Println(v1)
                }
                _ = fi29
            }
        case "a":
            if (!true) {
                v1 += v1
                var v43 uint16 = 745
                _ = v43
            }
        case "foo bar":
            v1 += (v1 * 0.1)
        }
    }
    v3 -= v22
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v26
    _ = v3
    _ = v4
}
