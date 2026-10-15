package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 int }
type GenS2 struct { F0 uint64; F1 uint32; F2 int16 }
type GenS3 struct { F0 int; F1 uint64 }
func genFn0() int64 {
    v4 := []int{1, -2147483648}
    switch -146 {
    case 1:
        var v5 uint8 = 0
        _ = v5
    case -2147483648:
        for fi6 := 0; fi6 < 3; fi6++ {
            for fi7 := 0; fi7 < 7; fi7++ {
                v8 := GenS1{F0: 2, F1: 42, F2: fi7}
                v4 = append(v4, (-(fi6)))
                switch "\n\t" {
                case "foo bar":
                    v4 = append(v4, -2147483648)
                case "hello":
                    var v9 int = -245
                    _ = v9
                case "a":
                    switch fi6 {
                    case 42:
                        _ = len(v4)
                        switch 289 {
                        case -289:
                            v4 = append(v4, 1)
                        case 4294967295:
                            if len(v4) > 0 {
                                v4[0] = fi6
                            }
                        case 11:
                            var v10 int64 = 9223372036854775807
                            _ = v10
                        }
                    default:
                        var v11 uint32 = 2
                        _ = v11
                    }
                default:
                    _ = len(v4)
                }
                var v12 uint64 = 2
                _ = fi7
                _ = v12
                _ = v8
            }
            var v13 bool = true
            _ = fi6
            _ = v13
        }
    case 400:
        v14 := GenS1{F0: 65535, F1: 969, F2: -311}
        _ = v14
    default:
        var v15 string = "a"
        _ = v15
    }
    var v16 float32 = 0.1
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    _ = v16
    _ = v4
    return 42
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int32 = 689
    v3 := []uint32{1, 0, 42, 690}
    if ((v2 & (v2 & 1)) > v2) {
        v2 = v2
        v2 += v2
    }
    if ((2 == 18446744073709551615) || (-32768 == -1)) {
        for fi17 := 0; fi17 < 4; fi17++ {
            var v18 int32 = v2
            v19 := GenS2{F0: 42, F1: 42, F2: -1}
            var v20 uint8 = 1
            _ = fi17
            _ = v18
            _ = v19
            _ = v20
        }
        switch 0 {
        case -522:
            switch 128 {
            case 9223372036854775807:
                for fi21 := 0; fi21 < 10; fi21++ {
                    v1[127] = ((("\n\t" + (("" + "foo bar") + "foo bar")) + "foo bar") + ("" + "foo bar"))
                    _ = fi21
                }
            case -204:
                for fi22 := 0; fi22 < 2; fi22++ {
                    switch 4294967295 {
                    case 1:
                        var v23 int64 = 9223372036854775807
                        _ = v23
                    }
                    v2 -= (^(v2))
                    _ = fi22
                }
            case -9223372036854775808:
                fmt.Println(v2)
            }
        case 226:
            fmt.Println(v2)
        default:
            if true {
                v2 -= -2147483648
            } else {
                if len(v3) > 0 {
                    v3[0] = 65535
                }
            }
        }
    }
    var v24 bool = true
    if (!(2 > 129)) {
        fmt.Println(v2, v2)
        var v25 int32 = 1
        v2 = (v2 % ((v25) | 1))
        _ = v25
    }
    switch -872 {
    case 42:
        var v26 int8 = -128
        _ = v26
    case 853:
        _ = len(v1)
    case 127:
        switch "hello" {
        case "":
            if false {
                var v27 uint16 = 255
                _ = v27
            }
        case "a":
            fmt.Println(v24)
        default:
            v1[65535] = "\n\t"
        }
    }
    for fi28 := 0; fi28 < 10; fi28++ {
        var v29 int16 = 1
        if (!(1.5 >= ((1e10 / 1.0) - ((1.5 / (float64((float32((((float64((1.5 - 0.0)) / float64(-2.25)) - (0.0 / 1.5)) * 0.1)) * (0.0 - float32(((1e10 * 0.0) + 1.5))))) * (1.5 * (float64(1e10) * (1e10 + float64((0.1 + 1e10))))))) - float64((float32((0.1 - ((0.0 - 0.1) + 0.0))) + (float32(((0.1 + ((1e10 - 0.0) * -2.25)) / 0.1)) * 0.0))))))) {
            var v30 uint16 = 255
            var v31 int64 = -888
            var v32 uint = 729
            _ = v30
            _ = v31
            _ = v32
        } else {
            v33 := GenS2{F0: 308, F1: 825, F2: v29}
            v34 := GenS1{F0: 18446744073709551615, F1: -803, F2: (fi28 + fi28)}
            _ = v33
            _ = v34
        }
        _ = fi28
        _ = v29
    }
    for fi35 := 0; fi35 < 16; fi35++ {
        v2 += (v2 & v2)
        _ = fi35
    }
    var v36 int16 = -956
    v37 := GenS3{F0: -853, F1: 42}
    var v38 uint8 = 128
    v39 := GenS3{F0: 1, F1: 18446744073709551615}
    v36 = v36
    fmt.Println(v38)
    switch "\n\t" {
    case "hello":
        v40 := GenS2{F0: 1, F1: 2147483647, F2: (v36 % ((v36) | 1))}
        _ = v40
    case "a":
        v41 := v1[42]
        _ = v41
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v42 := GenS2{F0: 0, F1: 2147483647, F2: 32767}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v42
}
