package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 uint; F1 int16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 string; F1 uint }
type GenS5 struct { F0 int16; F1 int16 }
type GenS6 struct { F0 int16 }

func main() {
    var v1 bool = ((true || true) && (false && false))
    var v2 float32 = (0.1 * 1.5)
    var v3 float32 = v2
    v4 := make(map[int]uint32)
    switch -2147483648 {
    case 587:
        fmt.Println(v2)
    case -840:
        v5 := GenS1{F0: 42, F1: 4294967295}
        _ = v5
    }
    switch 32767 {
    case 32768:
        var v6 int32 = 479
        _ = v6
    case -635:
        switch -1 {
        case 0:
            for fi7 := 0; fi7 < 5; fi7++ {
                var v8 int8 = 79
                v9 := GenS1{F0: (v8 * v8), F1: 662}
                _ = fi7
                _ = v8
                _ = v9
            }
        default:
            var v10 uint32 = 0
            _ = v10
        }
    default:
        switch 2147483647 {
        case -150:
            v3 = 1e10
        case 558:
            for fi11 := 0; fi11 < 15; fi11++ {
                if true {
                    if v1 {
                        v1 = v1
                        for fi12 := 0; fi12 < 16; fi12++ {
                            var v13 uint64 = 789
                            switch (-(fi11)) {
                            case 32768:
                                v2 += ((0.0 + (3.14159 + 3.14159)) - 1e10)
                            case -784:
                                v14 := GenS2{F0: 0, F1: 1}
                                _ = v14
                            case -613:
                                switch "foo bar" {
                                case "hello":
                                    switch fi12 {
                                    case 32767:
                                        var v15 uint = 65535
                                        _ = v15
                                    case -9223372036854775808:
                                        v16 := GenS2{F0: 81, F1: 42}
                                        _ = v16
                                    case -1:
                                        v17 := GenS2{F0: 18446744073709551615, F1: 1}
                                        _ = v17
                                    default:
                                        if (688 < 2) {
                                            fmt.Println(fi11, fi12, v13)
                                        } else {
                                            v18 := []int{fi12, 533}
                                            v19 := GenS2{F0: 1, F1: 1000}
                                            v20 := GenS3{F0: -88}
                                            _ = v18
                                            _ = v19
                                            _ = v20
                                        }
                                    }
                                case "foo bar":
                                    var v21 uint64 = (^(v13))
                                    _ = v21
                                case "a":
                                    switch fi11 {
                                    case 127:
                                        if v1 {
                                            fmt.Println(fi12, v1, v3)
                                            v3 += (float32(1.5) / 1.5)
                                        } else {
                                            var v22 uint16 = 32767
                                            var v23 int32 = 1
                                            v22 += (^(v22))
                                            v13 = (v13 % ((0) | 1))
                                            if (116 > 42) {
                                                var v24 string = "hello"
                                                var v25 string = (v24 + ("a" + (("" + "foo bar") + v24)))
                                                if v1 {
                                                    if v1 {
                                                        fmt.Println(v24, v2)
                                                        var v26 string = ("a" + "")
                                                        v22 += (v22 * (v22 - 65535))
                                                        _ = v26
                                                    } else {
                                                        switch (fi11 << ((188) & 63)) {
                                                        case -2147483648:
                                                            fmt.Println(v1, fi11, v24)
                                                        }
                                                        fmt.Println(v25, v24, v13)
                                                    }
                                                    fmt.Println(v25)
                                                }
                                                v27 := GenS4{F0: (v25 + (v25 + (((("hello" + v24) + "a") + "") + (v25 + (v24 + v24))))), F1: 65535}
                                                switch (fi11 & int(v23)) {
                                                case 238:
                                                    var v28 int32 = v23
                                                    var v29 int = fi12
                                                    v25 = (v24 + v24)
                                                    switch (-(fi12)) {
                                                    case 128:
                                                        if v1 {
                                                            var v30 int32 = -2147483648
                                                            var v31 float32 = v3
                                                            var v32 float32 = v2
                                                            _ = v30
                                                            _ = v31
                                                            _ = v32
                                                        } else {
                                                            var v33 int64 = -892
                                                            v34 := GenS2{F0: 65535, F1: 71}
                                                            var v35 float64 = -2.25
                                                            _ = v33
                                                            _ = v34
                                                            _ = v35
                                                        }
                                                    case 136:
                                                        v36 := []int{v29, v29, v29}
                                                        _ = v36
                                                    }
                                                    var v37 int8 = 0
                                                    if true {
                                                        var v38 uint16 = (v22 ^ v22)
                                                        _ = v38
                                                    }
                                                    var v39 int8 = (v37 / (((v37 | -1)) | 1))
                                                    v1 = (32767 == -1)
                                                    var v40 uint8 = 0
                                                    v41 := GenS2{F0: 873, F1: 0}
                                                    var v42 int64 = -9223372036854775808
                                                    _ = v28
                                                    _ = v29
                                                    _ = v37
                                                    _ = v39
                                                    _ = v40
                                                    _ = v41
                                                    _ = v42
                                                case 42:
                                                    v4[int(v13)] = 1
                                                case -893:
                                                    var v43 string = "\n\t"
                                                    _ = v43
                                                default:
                                                    v44 := GenS5{F0: 42, F1: 0}
                                                    _ = v44
                                                }
                                                v3 = (float32(1e10) / 1.0)
                                                _ = v24
                                                _ = v25
                                                _ = v27
                                            }
                                            _ = v22
                                            _ = v23
                                        }
                                    case -659:
                                        if (float32(((3.14159 * float64(v2)) * ((0.0 / -2.25) - float64(1.5)))) < v2) {
                                            if (!v1) {
                                                v4[127] = 1
                                            }
                                            var v45 uint8 = 0
                                            v2 = v2
                                            _ = v45
                                        }
                                    case 4294967295:
                                        var v46 uint32 = 2147483647
                                        _ = v46
                                    default:
                                        v1 = false
                                    }
                                default:
                                    if v1 {
                                        v3 -= v3
                                        v47 := GenS1{F0: 64, F1: 283}
                                        v3 = (v2 * (v2 / (v2 - v2)))
                                        _ = v47
                                    }
                                }
                            default:
                                var v48 float32 = -2.25
                                _ = v48
                            }
                            if (216 == 128) {
                                var v49 int8 = 42
                                _ = v49
                            } else {
                                var v50 string = "a"
                                fmt.Println(fi12)
                                var v51 int16 = 1
                                _ = v50
                                _ = v51
                            }
                            _ = fi12
                            _ = v13
                        }
                    }
                }
                _ = fi11
            }
        case 2147483647:
            var v52 uint64 = 21
            _ = v52
        }
    }
    var v53 int = 483
    v4[42] = 1
    switch "foo bar" {
    case "\n\t":
        switch "" {
        case "\n\t":
            if v1 {
                if v1 {
                    v54 := GenS3{F0: -1}
                    _ = v54
                }
            }
        default:
            switch (v53 ^ (v53 & v53)) {
            case 32768:
                for fi55 := 0; fi55 < 5; fi55++ {
                    fmt.Println(fi55)
                    _ = fi55
                }
            default:
                _ = len(v4)
            }
        }
    default:
        for fi56 := 0; fi56 < 9; fi56++ {
            v57 := []int{v53, 9223372036854775807, v53, (-(fi56))}
            if ((fi56 & -2147483648) > -9223372036854775808) {
                var v58 float32 = (v2 - (v2 * 1.5))
                _ = v58
            }
            _ = fi56
            _ = v57
        }
    }
    for fi59 := 0; fi59 < 12; fi59++ {
        fmt.Println(v53, v1, v3)
        v3 -= 1e10
        _ = fi59
    }
    var v60 uint64 = 42
    v53 = v53
    v61 := GenS6{F0: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v53:", v53)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v53
    _ = v60
    _ = v61
}
