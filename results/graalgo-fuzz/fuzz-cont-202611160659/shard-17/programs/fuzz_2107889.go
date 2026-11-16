package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 uint }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 uint16; F1 uint32; F2 int16 }
type GenS4 struct { F0 uint64 }

func main() {
    var v1 string = "foo bar"
    var v2 string = v1
    var v3 int64 = 42
    v4 := make(map[int]string)
    v5 := make(map[int]int)
    v6 := make(map[uint32]string)
    var v7 int8 = -128
    v8 := GenS1{F0: 127, F1: 174, F2: 0}
    for fi9 := 0; fi9 < 8; fi9++ {
        var v10 uint32 = 2147483647
        for fi11 := 0; fi11 < 8; fi11++ {
            fmt.Println(v10, v10)
            _ = fi11
        }
        _ = fi9
        _ = v10
    }
    switch "a" {
    case "a":
        for fi12 := 0; fi12 < 7; fi12++ {
            for fi13 := 0; fi13 < 11; fi13++ {
                v1 = ("foo bar" + (v2 + v1))
                switch (v2 + ("foo bar" + v2)) {
                case "foo bar":
                    v1 = (v2 + (((v2 + "foo bar") + ((v1 + "a") + ("" + ((v2 + v1) + v2)))) + v1))
                case "a":
                    v7 += (v7 * 81)
                }
                _ = fi13
            }
            var v14 float64 = 3.14159
            v15 := GenS1{F0: 42, F1: 4294967295, F2: 2}
            _ = fi12
            _ = v14
            _ = v15
        }
    default:
        for fi16 := 0; fi16 < 5; fi16++ {
            fmt.Println(v1)
            for fi17 := 0; fi17 < 19; fi17++ {
                v3 = (v3 << ((370) & 63))
                _ = fi17
            }
            var v18 float64 = (-2.25 * 3.14159)
            _ = fi16
            _ = v18
        }
    }
    v19 := GenS1{F0: 42, F1: 18446744073709551615, F2: 213}
    var v20 int = -545
    if ((3.14159 < -2.25) && (true && ((!(((float64(((3.14159 * (1e10 / 0.1)) * (0.1 / 1.0))) * 1.5) != 0.1) && (255 != 42))) || true))) {
        v21 := GenS2{F0: -244, F1: -1}
        _ = v21
    }
    switch v20 {
    case 128:
        var v22 uint8 = 127
        _ = v22
    case -1:
        for fi23 := 0; fi23 < 14; fi23++ {
            switch (-(v20)) {
            case -9223372036854775808:
                for fi24 := 0; fi24 < 12; fi24++ {
                    var v25 int16 = -32768
                    _ = fi24
                    _ = v25
                }
            }
            var v26 int8 = (v7 % ((-97) | 1))
            v1 = ""
            _ = fi23
            _ = v26
        }
    default:
        v27 := GenS1{F0: 1, F1: 65535, F2: 1}
        _ = v27
    }
    v28 := []int{v20, (v20 - (v20 * v20)), v20}
    for fi29 := 0; fi29 < 10; fi29++ {
        v30 := GenS1{F0: 1, F1: 1, F2: 42}
        for fi31 := 0; fi31 < 7; fi31++ {
            var v32 float32 = (1.5 / 3.14159)
            if true {
                if (42 == 1) {
                    v2 = "foo bar"
                    if ((1.5 + 3.14159) >= float64(0.0)) {
                        switch (v1 + ("" + v2)) {
                        case "foo bar":
                            fmt.Println(v20)
                            var v33 uint64 = 18446744073709551615
                            if len(v28) >= 2 {
                                v28 = v28[1:len(v28)]
                            }
                            _ = v33
                        case "":
                            var v34 uint8 = 127
                            _ = v34
                        }
                        v35 := GenS2{F0: -1, F1: 2147483647}
                        v3 = 9223372036854775807
                        _ = v35
                    }
                    if (float64(v32) < 1.5) {
                        if false {
                            var v36 float64 = 0.1
                            if (v3 >= v3) {
                                if len(v28) > 0 {
                                    v28[0] = (fi29 % ((32768) | 1))
                                }
                                switch fi31 {
                                case 32768:
                                    if (771 <= 255) {
                                        v37 := GenS1{F0: 1, F1: 255, F2: 4294967295}
                                        _ = v37
                                    }
                                    var v38 uint64 = 42
                                    var v39 int = (fi31 >> ((255) & 63))
                                    _ = v38
                                    _ = v39
                                }
                            } else {
                                var v40 int32 = -2147483648
                                var v41 int8 = v7
                                _ = v40
                                _ = v41
                            }
                            v5[(-(fi31))] = (v20 + fi29)
                            _ = v36
                        }
                        var v42 uint64 = 1
                        v32 -= v32
                        _ = v42
                    }
                } else {
                    if (true || false) {
                        _ = len(v4)
                        v32 += v32
                        v43 := v5[(fi31 * -51)]
                        _ = v43
                    } else {
                        switch (v20 * 1) {
                        case 2147483647:
                            switch "foo bar" {
                            case "foo bar":
                                if true {
                                    switch (-(fi31)) {
                                    case 0:
                                        if (!true) {
                                            var v44 float64 = -2.25
                                            var v45 uint32 = 1
                                            var v46 uint64 = 205
                                            _ = v44
                                            _ = v45
                                            _ = v46
                                        }
                                    case 314:
                                        v47 := GenS1{F0: 128, F1: 255, F2: 113}
                                        _ = v47
                                    }
                                    fmt.Println(v20)
                                }
                            case "hello":
                                fmt.Println(v3)
                            case "":
                                v1 = v2
                            default:
                                v32 -= v32
                            }
                        case 945:
                            if len(v28) > 0 {
                                v28[0] = (fi31 | fi31)
                            }
                        case -30:
                            v20 -= 277
                        default:
                            switch v20 {
                            case 0:
                                if (false || (true || false)) {
                                    v48 := GenS1{F0: 0, F1: 65535, F2: 182}
                                    v49 := GenS2{F0: -569, F1: 42}
                                    fmt.Println(fi29, fi29, v20)
                                    v20 = (fi31 << ((42) & 63))
                                    _ = v48
                                    _ = v49
                                } else {
                                    var v50 float64 = float64((v32 - (-2.25 / v32)))
                                    v51 := GenS1{F0: 127, F1: 42, F2: 27}
                                    switch v20 {
                                    case 0:
                                        var v52 bool = (v7 != (v7 << ((65535) & 63)))
                                        _ = v52
                                    }
                                    v2 = v1
                                    switch -2147483648 {
                                    case -9223372036854775808:
                                        var v53 int = v20
                                        _ = v53
                                    case 127:
                                        if true {
                                            v3 += (v3 % ((v3) | 1))
                                            var v54 uint32 = 719
                                            _ = v54
                                        }
                                    default:
                                        if ((-(fi31)) != int(v7)) {
                                            var v55 uint64 = 1
                                            v56 := GenS3{F0: 1, F1: 1, F2: int16(v55)}
                                            if ((1e10 == 0.1) && true) {
                                                var v57 int32 = 2147483647
                                                v55 += v55
                                                _ = v57
                                            }
                                            var v58 uint8 = 127
                                            _ = v55
                                            _ = v56
                                            _ = v58
                                        } else {
                                            if false {
                                                var v59 bool = (false && true)
                                                v50 = (v50 + ((float64(v32) / ((1e10 + -2.25) + -2.25)) + v50))
                                                _ = v59
                                            }
                                        }
                                    }
                                    _ = v50
                                    _ = v51
                                }
                            case -759:
                                switch (v20 / ((v20) | 1)) {
                                case 32768:
                                    if false {
                                        v3 = 1
                                        v60 := GenS3{F0: 2, F1: 961, F2: 1000}
                                        if (572 < 894) {
                                            v1 = (v2 + ("hello" + (v1 + "foo bar")))
                                            _ = len(v28)
                                            if true {
                                                fmt.Println(v1, v7, v2)
                                            } else {
                                                v7 += v7
                                                switch v2 {
                                                case "hello":
                                                    var v61 uint32 = 65535
                                                    _ = v61
                                                case "a":
                                                    v62 := GenS3{F0: 21, F1: 69, F2: -32768}
                                                    _ = v62
                                                }
                                            }
                                            var v63 float32 = (v32 - 1.5)
                                            _ = v63
                                        }
                                        var v64 uint32 = 2147483647
                                        _ = v60
                                        _ = v64
                                    } else {
                                        v32 = 1.5
                                        var v65 int64 = -1
                                        v7 -= v7
                                        var v66 bool = (((0.0 * (0.0 + ((0.1 - -2.25) * (float64(((0.0 / ((0.0 - (v32 + (float32(1e10) + (((3.14159 * (1e10 / v32)) - (v32 * (3.14159 / float32(-2.25)))) / v32)))) - v32)) / 1.5)) / float64((float32(float64((v32 / v32))) / v32)))))) + (1e10 - 1.5)) != (((float64(float32(3.14159)) / -2.25) / 0.1) * (1e10 + (-2.25 * 3.14159))))
                                        fmt.Println(v3)
                                        _ = v65
                                        _ = v66
                                    }
                                }
                            case 32767:
                                v3 -= (v3 - v3)
                            }
                        }
                        if true {
                            v2 = ("" + (v2 + "\n\t"))
                            delete(v4, -2147483648)
                        }
                        if false {
                            if (!((255 <= 471) && (true && (2147483647 > 15)))) {
                                v7 = 0
                                fmt.Println(v20, v3)
                            } else {
                                v3 += v3
                                var v67 uint16 = 42
                                _ = v67
                            }
                            switch v20 {
                            case -9223372036854775808:
                                _ = len(v5)
                            case -581:
                                var v68 int64 = v3
                                _ = v68
                            }
                        }
                    }
                    if (0 == 0) {
                        var v69 uint16 = 42
                        fmt.Println(v7)
                        var v70 uint = 65535
                        fmt.Println(fi29, v3, v20)
                        var v71 uint = (v70 - 836)
                        _ = v69
                        _ = v70
                        _ = v71
                    }
                    _ = len(v5)
                }
                var v72 int32 = -2147483648
                _ = v72
            }
            _ = fi31
            _ = v32
        }
        _ = fi29
        _ = v30
    }
    v73 := GenS4{F0: 42}
    v7 = -128
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v73
    _ = v8
}
