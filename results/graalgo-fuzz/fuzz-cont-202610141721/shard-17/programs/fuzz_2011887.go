package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 int64 }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 uint16; F1 int64 }
type GenS4 struct { F0 int; F1 uint }
type GenS5 struct { F0 int32; F1 uint64 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint32; F1 uint8 }
type GenS8 struct { F0 int; F1 int16 }

func main() {
    var v1 float64 = (0.1 * 0.1)
    v2 := make(map[uint32]uint32)
    var v3 string = "foo bar"
    var v4 float32 = ((1e10 + 0.0) / (1e10 / float32(v1)))
    var v5 float32 = v4
    v6 := GenS1{F0: 0, F1: 64, F2: 0}
    var v7 float64 = ((v1 + v1) / (v1 * (float64(v4) / v1)))
    var v8 int16 = 0
    v9 := []int{-566}
    v10 := v2[1]
    var v11 bool = true
    v12 := GenS2{F0: 156, F1: -653}
    fmt.Println(v8, v4, v5)
    fmt.Println(v8, v3, v3)
    fmt.Println(v11)
    v5 += (v4 + 1e10)
    switch -2147483648 {
    case -10:
        switch 42 {
        case 2147483647:
            switch 744 {
            case 1:
                fmt.Println(v8, v8, v7)
            case 128:
                switch 737 {
                case 1:
                    _ = len(v9)
                case 816:
                    var v13 int8 = 1
                    _ = v13
                default:
                    for fi14 := 0; fi14 < 17; fi14++ {
                        _ = len(v9)
                        fmt.Println(v1, v5, v10)
                        _ = fi14
                    }
                }
            default:
                var v15 int16 = v8
                _ = v15
            }
        default:
            if (!true) {
                switch -1 {
                case -1:
                    v10 -= (^(v10))
                case 1:
                    v8 -= (v8 & v8)
                case 4294967295:
                    if (0 != -1) {
                        var v16 int16 = (^(v8))
                        _ = v16
                    }
                }
            } else {
                v2[2147483647] = v10
                _ = len(v9)
                fmt.Println(v7)
            }
        }
    }
    v17 := GenS1{F0: 42, F1: 2147483647, F2: -1}
    switch v3 {
    case "\n\t":
        var v18 bool = false
        _ = v18
    case "a":
        v5 += (-2.25 * (-2.25 * (v5 + v4)))
    case "":
        var v19 uint16 = 0
        _ = v19
    default:
        v20 := GenS3{F0: 302, F1: 9223372036854775807}
        _ = v20
    }
    for fi21 := 0; fi21 < 18; fi21++ {
        if (v8 != (v8 * v8)) {
            for fi22 := 0; fi22 < 5; fi22++ {
                var v23 int8 = -52
                switch "foo bar" {
                case "":
                    if true {
                        if (v8 <= (v8 / (((v8 | (v8 * (-(v8))))) | 1))) {
                            var v24 int = fi21
                            _ = v24
                        } else {
                            var v25 string = (v3 + v3)
                            var v26 int64 = -1
                            v27 := GenS4{F0: -715, F1: 452}
                            _ = v25
                            _ = v26
                            _ = v27
                        }
                        if (true && false) {
                            var v28 int8 = 0
                            var v29 float32 = 0.0
                            _ = v28
                            _ = v29
                        }
                    }
                }
                _ = fi22
                _ = v23
            }
        } else {
            if ((float64((v5 + v4)) * v1) >= (1e10 + v7)) {
                if v11 {
                    switch (fi21 - 898) {
                    case 127:
                        switch fi21 {
                        case -1:
                            var v30 uint64 = 0
                            _ = v30
                        default:
                            for fi31 := 0; fi31 < 1; fi31++ {
                                v11 = (1 < 2)
                                var v32 int32 = 1
                                var v33 bool = v11
                                var v34 int16 = v8
                                _ = fi31
                                _ = v32
                                _ = v33
                                _ = v34
                            }
                        }
                    default:
                        fmt.Println(v5)
                    }
                    v2[4294967295] = (^(v10))
                }
                var v35 int = fi21
                for fi36 := 0; fi36 < 3; fi36++ {
                    switch v3 {
                    case "foo bar":
                        if (v11 || (((float32(float64(0.0)) - v5) - ((v5 + 3.14159) * (0.1 + v5))) <= v5)) {
                            v3 = (v3 + v3)
                        } else {
                            delete(v2, (v10 % ((v10) | 1)))
                            fmt.Println(fi36, fi21)
                            var v37 float32 = v5
                            _ = v37
                        }
                    case "\n\t":
                        v8 -= v8
                    case "a":
                        switch "\n\t" {
                        case "hello":
                            v7 = v7
                            if false {
                                if (v11 || (v11 && v11)) {
                                    v5 += v5
                                    if (0.1 >= 3.14159) {
                                        fmt.Println(fi21, v8)
                                    }
                                }
                                if (0 > 2) {
                                    var v38 bool = (v11 && false)
                                    var v39 int32 = -2147483648
                                    if (255 == 0) {
                                        switch ((((v3 + (v3 + (v3 + v3))) + "") + "hello") + "foo bar") {
                                        case "":
                                            v40 := GenS1{F0: v39, F1: v39, F2: -669}
                                            _ = v40
                                        }
                                        v9 = append(v9, (fi36 % (((^(fi21))) | 1)))
                                        v8 += (v8 / ((int16(v10)) | 1))
                                    } else {
                                        if true {
                                            switch (-(fi21)) {
                                            case 128:
                                                fmt.Println(v5, v8, v3)
                                            }
                                            var v41 int16 = (v8 * v8)
                                            _ = v41
                                        } else {
                                            var v42 int16 = 1
                                            var v43 int32 = (v39 << ((4294967295) & 63))
                                            var v44 int8 = 127
                                            _ = v42
                                            _ = v43
                                            _ = v44
                                        }
                                        var v45 uint64 = 2
                                        fmt.Println(v38, v10, v10)
                                        if v11 {
                                            if false {
                                                var v46 uint8 = 255
                                                if (1 != (v10 >> ((42) & 63))) {
                                                    v47 := v2[2]
                                                    v5 = 0.0
                                                    var v48 uint16 = 450
                                                    _ = v47
                                                    _ = v48
                                                } else {
                                                    if v38 {
                                                        _ = len(v9)
                                                        if ((321 > 213) || ((v3 + (v3 + ("a" + v3))) <= "a")) {
                                                            if v11 {
                                                                fmt.Println(v1)
                                                                if len(v9) >= 2 {
                                                                    v9 = v9[1:len(v9)]
                                                                }
                                                                fmt.Println(v5)
                                                            } else {
                                                                if (938 == 32767) {
                                                                    var v49 string = v3
                                                                    switch ((v49 + "a") + "foo bar") {
                                                                    case "a":
                                                                        v11 = v38
                                                                    case "hello":
                                                                        var v50 int16 = (v8 ^ (v8 ^ v8))
                                                                        _ = v50
                                                                    }
                                                                    _ = v49
                                                                } else {
                                                                    _ = len(v9)
                                                                    var v51 uint8 = v46
                                                                    _ = v51
                                                                }
                                                                v52 := GenS1{F0: -1, F1: -599, F2: 0}
                                                                var v53 uint64 = (v45 << ((65535) & 63))
                                                                _ = v52
                                                                _ = v53
                                                            }
                                                            v54 := v2[(^(v10))]
                                                            _ = v54
                                                        }
                                                    }
                                                    var v55 float32 = float32(v7)
                                                    v38 = (2 >= 65535)
                                                    _ = v55
                                                }
                                                var v56 int = int(v8)
                                                var v57 uint8 = 2
                                                v58 := GenS5{F0: v39, F1: v45}
                                                if (v3 <= v3) {
                                                    var v59 uint = 722
                                                    _ = v59
                                                } else {
                                                    var v60 uint = 255
                                                    v61 := GenS6{F0: (v10 % ((v10) | 1))}
                                                    var v62 float64 = (-2.25 * v1)
                                                    fmt.Println(fi21, fi36)
                                                    _ = v60
                                                    _ = v61
                                                    _ = v62
                                                }
                                                _ = v46
                                                _ = v56
                                                _ = v57
                                                _ = v58
                                            }
                                        }
                                        if len(v9) >= 2 {
                                            v9 = v9[1:len(v9)]
                                        }
                                        _ = v45
                                    }
                                    fmt.Println(v38, fi36)
                                    switch 599 {
                                    case 42:
                                        v63 := GenS2{F0: -1, F1: (^(v39))}
                                        _ = v63
                                    }
                                    _ = v38
                                    _ = v39
                                }
                                delete(v2, 895)
                            }
                            v64 := GenS7{F0: (^(v10)), F1: 2}
                            _ = v64
                        case "a":
                            fmt.Println(fi21, fi36, v11)
                        }
                    }
                    fmt.Println(v11, v8)
                    var v65 uint = 1
                    _ = fi36
                    _ = v65
                }
                _ = v35
            } else {
                if len(v9) > 0 {
                    v9[0] = -137
                }
            }
            _ = len(v2)
            v66 := GenS8{F0: (fi21 & fi21), F1: 367}
            _ = v66
        }
        if true {
            v1 = (1e10 / 1e10)
        }
        v2[(v10 * 2)] = (v10 / (((v10 * (v10 + (v10 / ((89) | 1))))) | 1))
        _ = fi21
    }
    for fi67 := 0; fi67 < 4; fi67++ {
        v68 := GenS1{F0: 42, F1: 0, F2: -132}
        v11 = false
        _ = fi67
        _ = v68
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
