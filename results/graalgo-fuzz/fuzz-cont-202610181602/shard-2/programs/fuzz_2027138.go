package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16; F1 uint }
type GenS4 struct { F0 bool }
func genFn0() uint {
    for fi4 := 0; fi4 < 9; fi4++ {
        fmt.Println(fi4)
        _ = fi4
    }
    for fi5 := 0; fi5 < 6; fi5++ {
        if false {
            for fi6 := 0; fi6 < 8; fi6++ {
                v7 := []int{fi5, (fi5 % (((fi5 ^ (fi6 + (-(fi5))))) | 1)), (fi6 & fi5)}
                var v8 uint64 = 346
                switch "" {
                case "\n\t":
                    if (42 < 9223372036854775807) {
                        v8 += v8
                        if len(v7) > 0 {
                            v7[0] = 42
                        }
                    }
                case "":
                    switch fi6 {
                    case 507:
                        var v9 float64 = 0.1
                        _ = v9
                    case 0:
                        v8 -= (^(v8))
                    case 4294967295:
                        fmt.Println(v8)
                        if (!true) {
                            var v10 int32 = 2147483647
                            _ = v10
                        }
                    }
                case "hello":
                    fmt.Println(v8)
                default:
                    fmt.Println(fi5, v8)
                }
                _ = fi6
                _ = v7
                _ = v8
            }
            var v11 uint = 42
            if (false || true) {
                v11 -= 2
            }
            _ = v11
        }
        var v12 int32 = 2147483647
        v13 := []int{128}
        _ = fi5
        _ = v12
        _ = v13
    }
    switch -280 {
    case 4294967295:
        v14 := GenS1{F0: 42, F1: 0}
        _ = v14
    default:
        var v15 int8 = -40
        _ = v15
    }
    return 476
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int32 = -6
    var v3 int = -9223372036854775808
    var v16 int8 = 1
    v17 := GenS2{F0: 364}
    if (-25 < 684) {
        var v18 uint8 = 255
        _ = v18
    } else {
        switch "foo bar" {
        case "foo bar":
            _ = len(v1)
        case "":
            v19 := GenS3{F0: 0, F1: genFn0()}
            _ = v19
        default:
            fmt.Println(v2, v2)
        }
        v1[65535] = 42
        var v20 int8 = v16
        _ = v20
    }
    var v21 int16 = 1000
    var v22 uint8 = 127
    v23 := []int{(v3 >> ((316) & 63))}
    switch ("" + (("" + "foo bar") + ("hello" + ""))) {
    case "":
        for fi24 := 0; fi24 < 18; fi24++ {
            for fi25 := 0; fi25 < 11; fi25++ {
                var v26 string = "\n\t"
                _ = fi25
                _ = v26
            }
            switch fi24 {
            case -815:
                v27 := GenS3{F0: v21, F1: 408}
                _ = v27
            case 0:
                var v28 int32 = (-(v2))
                _ = v28
            case 2147483647:
                delete(v1, 286)
            default:
                if true {
                    fmt.Println(v16, v16, v16)
                }
            }
            if true {
                if true {
                    v21 += -610
                    var v29 int = (fi24 % (((-(fi24))) | 1))
                    v30 := GenS1{F0: 2147483648, F1: (v2 | 2147483647)}
                    _ = v29
                    _ = v30
                }
                var v31 uint64 = 4294967295
                _ = v31
            }
            _ = fi24
        }
    case "hello":
        if true {
            var v32 uint = 4294967295
            var v33 float64 = 0.0
            _ = v32
            _ = v33
        } else {
            v16 = int8(v3)
        }
    default:
        v21 += v21
    }
    var v34 float32 = -2.25
    if true {
        v35 := v1[2147483647]
        for fi36 := 0; fi36 < 8; fi36++ {
            var v37 uint64 = 4294967295
            if (((true || (!true)) || (true || ((v21 / (((v21 / ((-1) | 1))) | 1)) <= v21))) || false) {
                var v38 int64 = 0
                _ = v38
            }
            _ = fi36
            _ = v37
        }
        v23 = append(v23, (-(v3)))
        _ = v35
    } else {
        delete(v1, 327)
        var v39 float64 = 1.5
        var v40 uint8 = v22
        _ = v39
        _ = v40
    }
    switch v3 {
    case 0:
        var v41 uint16 = 2
        _ = v41
    case -930:
        var v42 uint = 42
        _ = v42
    default:
        v43 := GenS4{F0: false}
        _ = v43
    }
    v3 = (v3 | 9223372036854775807)
    v34 = 0.0
    var v44 int = v3
    v21 -= 1000
    _ = len(v1)
    v22 = v22
    _ = len(v23)
    v16 = (v16 * v16)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v34
    _ = v44
}
