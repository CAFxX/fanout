package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32; F1 uint16; F2 uint16 }
type GenS3 struct { F0 int64; F1 uint }
type GenS4 struct { F0 uint8; F1 uint16 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int64; F1 bool }
type GenS7 struct { F0 uint64; F1 uint; F2 uint16 }
type GenS8 struct { F0 uint64 }
func genFn0(p0 uint64, p1 bool) int16 {
    var v6 int8 = 1
    v7 := GenS1{F0: 834}
    p1 = p1
    v8 := []int{2147483647, 32768, 2147483647, 2147483647}
    switch "" {
    case "\n\t":
        _ = len(v8)
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 uint, p1 uint64, p2 uint32) bool {
    var v9 int32 = 42
    switch ("" + (("a" + "foo bar") + "foo bar")) {
    case "hello":
        p0 -= 18446744073709551615
    case "\n\t":
        for fi10 := 0; fi10 < 19; fi10++ {
            switch ("" + "a") {
            case "hello":
                v9 = (v9 | (^(v9)))
            case "a":
                switch fi10 {
                case 707:
                    v11 := []int{127}
                    _ = v11
                }
            case "\n\t":
                var v12 int32 = (v9 % (((v9 + (v9 * v9))) | 1))
                _ = v12
            default:
                if true {
                    if (!(true && true)) {
                        var v13 uint16 = 0
                        _ = v13
                    } else {
                        var v14 uint = p0
                        _ = v14
                    }
                }
            }
            _ = fi10
        }
    case "a":
        var v15 int32 = (v9 % (((v9 << ((p0) & 63))) | 1))
        _ = v15
    default:
        v16 := []int{585, 747, 32768}
        _ = v16
    }
    v17 := []int{-306}
    v18 := GenS2{F0: (v9 - v9), F1: 65535, F2: 32767}
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v18
    _ = v9
    return (true && ((p0 - 65535) <= 538))
}

func main() {
    var v1 string = "\n\t"
    var v2 bool = false
    var v3 uint64 = 42
    var v4 uint32 = 1
    var v5 uint32 = (^(v4))
    var v19 int32 = -1
    v20 := []int{0, -9223372036854775808}
    v21 := GenS3{F0: 9223372036854775807, F1: 768}
    var v22 uint32 = v5
    fmt.Println(v22)
    v2 = true
    for fi23 := 0; fi23 < 18; fi23++ {
        if (!(false || (1e10 > 1e10))) {
            switch (fi23 - fi23) {
            case 0:
                if false {
                    var v24 int64 = -683
                    _ = v24
                } else {
                    switch v1 {
                    case "\n\t":
                        v2 = (4294967295 != 2)
                    }
                    fmt.Println(v5)
                    _ = len(v20)
                }
            case 45:
                v19 = v19
            case 127:
                switch fi23 {
                case 2147483647:
                    v25 := GenS1{F0: 3}
                    _ = v25
                case 127:
                    v26 := GenS3{F0: 0, F1: 0}
                    _ = v26
                default:
                    for fi27 := 0; fi27 < 8; fi27++ {
                        var v28 int16 = 1000
                        _ = fi27
                        _ = v28
                    }
                }
            default:
                var v29 uint8 = 178
                _ = v29
            }
        }
        fmt.Println(v19, v22, v22)
        var v30 int8 = 127
        _ = fi23
        _ = v30
    }
    for fi31 := 0; fi31 < 12; fi31++ {
        v4 = (v4 + 2)
        switch (^(fi31)) {
        case 32767:
            if len(v20) >= 2 {
                v20 = v20[1:len(v20)]
            }
        case 42:
            v5 -= 0
        case 4294967295:
            var v32 int64 = -1
            _ = v32
        default:
            fmt.Println(v19, v2, v19)
        }
        _ = fi31
    }
    if (184 <= 284) {
        switch -64 {
        case 517:
            var v33 int = 79
            _ = v33
        case -9223372036854775808:
            v2 = ((v3 <= (v3 << ((0) & 63))) && v2)
        case 9223372036854775807:
            v22 = v22
        default:
            v2 = (!v2)
        }
        for fi34 := 0; fi34 < 20; fi34++ {
            var v35 uint = 42
            switch ((v1 + "\n\t") + ("" + v1)) {
            case "\n\t":
                v5 = (^(v4))
            case "hello":
                fmt.Println(v35)
            default:
                switch ((("a" + "\n\t") + v1) + v1) {
                case "a":
                    if v2 {
                        if v2 {
                            var v36 string = ((v1 + "") + v1)
                            v2 = (v2 || v2)
                            if len(v20) >= 2 {
                                v20 = v20[1:len(v20)]
                            }
                            _ = v36
                        }
                        if (!false) {
                            _ = len(v20)
                            if ((-2.25 == (0.1 * (float32(1.5) + (1.5 * (1.5 - (-2.25 - (3.14159 + 3.14159))))))) || (1e10 != ((float32(1.5) * (1e10 * float32(0.0))) * 0.1))) {
                                if v2 {
                                    if v2 {
                                        var v37 uint8 = 164
                                        _ = v37
                                    }
                                    if ((((float32(3.14159) * ((0.0 + (0.0 + 0.0)) - -2.25)) / 1.5) / (float32(1.5) * 3.14159)) < (float32((0.0 + (0.1 + (float64(float32(float64(3.14159))) / (3.14159 / -2.25))))) - 0.0)) {
                                        if ((^(v3)) < v3) {
                                            for fi38 := 0; fi38 < 18; fi38++ {
                                                v35 = v35
                                                if (v2 && true) {
                                                    v1 = v1
                                                    var v39 float64 = 0.0
                                                    _ = v39
                                                } else {
                                                    v40 := GenS4{F0: 128, F1: 591}
                                                    switch -575 {
                                                    case 32768:
                                                        var v41 float32 = float32(-2.25)
                                                        _ = v41
                                                    case -374:
                                                        v42 := GenS4{F0: 255, F1: 42}
                                                        var v43 int64 = 9223372036854775807
                                                        _ = v42
                                                        _ = v43
                                                    default:
                                                        var v44 int8 = 0
                                                        _ = v44
                                                    }
                                                    _ = v40
                                                }
                                                var v45 uint16 = 254
                                                _ = fi38
                                                _ = v45
                                            }
                                        }
                                        var v46 float32 = 0.1
                                        v5 = (v5 & uint32(v35))
                                        _ = v46
                                    }
                                }
                                if len(v20) > 0 {
                                    v20[0] = (^(fi34))
                                }
                                fmt.Println(v19)
                            } else {
                                switch ("hello" + ("" + v1)) {
                                case "\n\t":
                                    var v47 uint64 = (v3 << ((v35) & 63))
                                    _ = v47
                                case "hello":
                                    fmt.Println(fi34, v35, v5)
                                case "a":
                                    switch fi34 {
                                    case 1:
                                        v20 = append(v20, fi34)
                                    case 128:
                                        var v48 float64 = 1.5
                                        _ = v48
                                    case 311:
                                        var v49 int8 = 100
                                        _ = v49
                                    default:
                                        fmt.Println(v4)
                                    }
                                }
                                if v2 {
                                    v50 := GenS5{F0: (^(v4))}
                                    v51 := GenS5{F0: v22}
                                    fmt.Println(v3, v19, v19)
                                    _ = v50
                                    _ = v51
                                }
                            }
                            if true {
                                if (242 <= (^(v35))) {
                                    if (v3 == v3) {
                                        if (42 < -128) {
                                            v22 = v22
                                            v20 = append(v20, fi34)
                                            fmt.Println(fi34, v5)
                                        } else {
                                            var v52 bool = (v19 < v19)
                                            _ = v52
                                        }
                                    }
                                    if len(v20) > 0 {
                                        v20[0] = (fi34 + fi34)
                                    }
                                    if len(v20) > 0 {
                                        v20[0] = 0
                                    }
                                }
                                v22 = (v4 & (v4 / ((176) | 1)))
                            }
                        } else {
                            for fi53 := 0; fi53 < 16; fi53++ {
                                v3 -= 18446744073709551615
                                v3 -= v3
                                switch (fi34 - fi34) {
                                case 143:
                                    if (v2 || (!(1000 <= -32768))) {
                                        v54 := GenS3{F0: -1, F1: v35}
                                        v3 -= (^(v3))
                                        var v55 int16 = 42
                                        var v56 float64 = (-2.25 * 0.0)
                                        _ = v54
                                        _ = v55
                                        _ = v56
                                    } else {
                                        if false {
                                            v19 -= (v19 - (v19 / (((^(v19))) | 1)))
                                            fmt.Println(v19, v5)
                                        }
                                        if len(v20) > 0 {
                                            v20[0] = int(v19)
                                        }
                                        fmt.Println(v22, fi34, v22)
                                    }
                                case -2147483648:
                                    v57 := GenS6{F0: 0, F1: v2}
                                    _ = v57
                                case 108:
                                    v1 = "\n\t"
                                default:
                                    v58 := GenS7{F0: v3, F1: v35, F2: 32767}
                                    _ = v58
                                }
                                _ = fi53
                            }
                            if len(v20) > 0 {
                                v20[0] = fi34
                            }
                        }
                    }
                case "":
                    v59 := GenS8{F0: (v3 >> (((v35 * v35)) & 63))}
                    _ = v59
                }
            }
            var v60 uint16 = 2
            _ = fi34
            _ = v35
            _ = v60
        }
        v20 = append(v20, -1)
    } else {
        for fi61 := 0; fi61 < 16; fi61++ {
            v62 := GenS6{F0: 1, F1: (-2.25 == 0.0)}
            _ = fi61
            _ = v62
        }
        var v63 string = "foo bar"
        _ = v63
    }
    var v64 int8 = -25
    fmt.Println(v19)
    v64 += (v64 | (-(v64)))
    var v65 float32 = float32(1.5)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v64
    _ = v65
}
