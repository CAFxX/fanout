package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 uint32 }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 uint8; F1 uint; F2 bool }
type GenS4 struct { F0 bool; F1 int8 }
type GenS5 struct { F0 uint32; F1 uint8 }
type GenS6 struct { F0 uint16; F1 uint64 }
type GenS7 struct { F0 bool; F1 int32 }
type GenS8 struct { F0 string; F1 string }
type GenS9 struct { F0 uint8; F1 string }
func genFn0() bool {
    if true {
        var v5 uint16 = 1
        v5 -= (v5 ^ v5)
        var v6 bool = ((42 > 7) || (-2.25 <= 1e10))
        _ = v5
        _ = v6
    } else {
        var v7 float32 = 3.14159
        for fi8 := 0; fi8 < 15; fi8++ {
            v9 := []int{}
            fmt.Println(v7, v7)
            v7 = v7
            _ = fi8
            _ = v9
        }
        _ = v7
    }
    for fi10 := 0; fi10 < 4; fi10++ {
        if true {
            v11 := []int{fi10}
            _ = v11
        }
        var v12 uint8 = 127
        v12 = v12
        _ = fi10
        _ = v12
    }
    return false
}
func genFn1(p0 uint8, p1 uint64) uint16 {
    p1 += 305
    v13 := GenS1{F0: 4294967295, F1: 865, F2: uint32(p1)}
    var v14 float32 = -2.25
    v14 += v14
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    return 799
}

func main() {
    v1 := make(map[string]uint32)
    var v2 string = ((("" + ("foo bar" + ("foo bar" + "foo bar"))) + "foo bar") + (("a" + ("\n\t" + "a")) + ("hello" + (("a" + "\n\t") + ""))))
    var v3 float64 = (0.1 - (float64((3.14159 / 1.5)) * (float64(((float32(-2.25) / (1e10 + 3.14159)) / 0.1)) - 0.1)))
    var v4 uint32 = 42
    v15 := GenS1{F0: 1, F1: 0, F2: (^(v4))}
    fmt.Println(v2, v4, v3)
    v2 = v2
    if false {
        var v16 bool = (0 < 255)
        switch 42 {
        case 0:
            v16 = true
        case 128:
            switch -57 {
            case 9223372036854775807:
                if (!(255 > 1)) {
                    switch 1 {
                    case 274:
                        v17 := []int{-1}
                        _ = v17
                    case -2147483648:
                        if (("a" + "") != (v2 + (v2 + v2))) {
                            v18 := []int{128, 42}
                            for fi19 := 0; fi19 < 1; fi19++ {
                                if v16 {
                                    if (!(42 == int8(fi19))) {
                                        v16 = v16
                                    } else {
                                        fmt.Println(fi19, v3)
                                        if (127 == uint8(v4)) {
                                            if (v16 && v16) {
                                                var v20 float32 = -2.25
                                                var v21 uint = 42
                                                _ = v20
                                                _ = v21
                                            } else {
                                                v4 = 65535
                                            }
                                            var v22 bool = (((0.1 * 0.0) - v3) < (((float64(1e10) + 3.14159) / ((v3 / 1.5) - v3)) / v3))
                                            _ = v22
                                        }
                                    }
                                    switch (fi19 / ((fi19) | 1)) {
                                    case 128:
                                        v23 := GenS1{F0: 4294967295, F1: 42, F2: 42}
                                        _ = v23
                                    }
                                }
                                for fi24 := 0; fi24 < 17; fi24++ {
                                    v25 := GenS1{F0: 65535, F1: 38, F2: 0}
                                    var v26 int32 = 42
                                    _ = fi24
                                    _ = v25
                                    _ = v26
                                }
                                _ = fi19
                            }
                            v18 = append(v18, -68)
                            _ = v18
                        }
                    case -9223372036854775808:
                        var v27 int64 = 0
                        _ = v27
                    default:
                        var v28 float32 = (3.14159 * 1e10)
                        _ = v28
                    }
                    v29 := GenS1{F0: 0, F1: 0, F2: 42}
                    switch 32768 {
                    case 4294967295:
                        v4 = (v4 << ((731) & 63))
                    }
                    _ = v29
                }
            case 42:
                var v30 int16 = 1
                _ = v30
            case -2147483648:
                var v31 uint64 = 18446744073709551615
                _ = v31
            default:
                switch v2 {
                case "hello":
                    var v32 int = 185
                    _ = v32
                case "foo bar":
                    var v33 int16 = 32767
                    _ = v33
                default:
                    _ = len(v1)
                }
            }
        default:
            fmt.Println(v2, v16)
        }
        if v16 {
            v16 = true
            if v16 {
                v3 += v3
                v3 += float64(1.5)
                for fi34 := 0; fi34 < 5; fi34++ {
                    if v16 {
                        v16 = v16
                        var v35 int64 = 0
                        v3 += v3
                        _ = v35
                    } else {
                        for fi36 := 0; fi36 < 16; fi36++ {
                            v37 := GenS2{F0: 1, F1: 1}
                            _ = fi36
                            _ = v37
                        }
                    }
                    for fi38 := 0; fi38 < 19; fi38++ {
                        v4 -= (v4 | v4)
                        _ = len(v1)
                        if v16 {
                            var v39 int32 = 0
                            _ = v39
                        } else {
                            _ = len(v1)
                            v3 += -2.25
                            fmt.Println(fi38)
                        }
                        _ = fi38
                    }
                    _ = fi34
                }
            } else {
                fmt.Println(v16, v3)
            }
            switch 1 {
            case -9223372036854775808:
                var v40 int = 446
                _ = v40
            default:
                v4 -= (v4 | 131)
            }
        }
        _ = v16
    } else {
        v41 := GenS2{F0: 1, F1: 1}
        _ = v41
    }
    v42 := GenS2{F0: -490, F1: 0}
    fmt.Println(v3, v3)
    for fi43 := 0; fi43 < 9; fi43++ {
        switch fi43 {
        case -243:
            if true {
                switch "" {
                case "\n\t":
                    var v44 int16 = 32767
                    _ = v44
                case "foo bar":
                    v4 = (^(v4))
                case "a":
                    var v45 int8 = 100
                    _ = v45
                }
                v2 = ""
                if false {
                    v46 := []int{fi43, 32768, (fi43 ^ (-(fi43)))}
                    v2 = v2
                    _ = v46
                }
            } else {
                v47 := []int{fi43, 32768, 2147483647, fi43}
                for fi48 := 0; fi48 < 7; fi48++ {
                    var v49 string = ""
                    v3 += ((v3 + v3) - 0.0)
                    v50 := GenS3{F0: 77, F1: 668, F2: true}
                    v51 := GenS4{F0: true, F1: -10}
                    _ = fi48
                    _ = v49
                    _ = v50
                    _ = v51
                }
                v4 = (v4 >> ((638) & 63))
                _ = v47
            }
        case 1:
            if (4294967295 >= 0) {
                var v52 uint32 = (v4 << ((255) & 63))
                v53 := GenS5{F0: 0, F1: 126}
                v54 := GenS5{F0: (v52 - 1), F1: 2}
                _ = v52
                _ = v53
                _ = v54
            } else {
                for fi55 := 0; fi55 < 11; fi55++ {
                    v56 := []int{}
                    _ = fi55
                    _ = v56
                }
                v2 = (v2 + "hello")
            }
        default:
            v57 := GenS3{F0: 2, F1: 42, F2: true}
            _ = v57
        }
        _ = fi43
    }
    v2 = "hello"
    for fi58 := 0; fi58 < 7; fi58++ {
        v4 -= 4294967295
        _ = fi58
    }
    var v59 float64 = ((1.5 - v3) * 3.14159)
    if (true || (-128 > 67)) {
        switch 42 {
        case 9223372036854775807:
            v60 := GenS1{F0: 0, F1: 4294967295, F2: 758}
            _ = v60
        case 171:
            v61 := GenS4{F0: (true || false), F1: 42}
            _ = v61
        }
    } else {
        for fi62 := 0; fi62 < 9; fi62++ {
            var v63 uint = 1
            v64 := GenS6{F0: 920, F1: 2}
            if (42 <= -76) {
                switch v2 {
                case "":
                    v63 += v63
                }
                if (-676 >= -528) {
                    if (int32(v63) > -2147483648) {
                        var v65 int8 = int8(v4)
                        for fi66 := 0; fi66 < 19; fi66++ {
                            v59 -= 3.14159
                            v67 := GenS7{F0: true, F1: -1}
                            v68 := GenS8{F0: "foo bar", F1: v2}
                            v69 := v1["\n\t"]
                            _ = fi66
                            _ = v67
                            _ = v68
                            _ = v69
                        }
                        v65 += v65
                        _ = v65
                    } else {
                        if (!true) {
                            fmt.Println(v59, v2)
                        } else {
                            if true {
                                var v70 int = fi62
                                for fi71 := 0; fi71 < 3; fi71++ {
                                    var v72 int32 = -262
                                    if true {
                                        if ((1.5 >= 0.1) && true) {
                                            fmt.Println(fi71, v3, v63)
                                            var v73 bool = (true && true)
                                            _ = v73
                                        }
                                        _ = len(v1)
                                    } else {
                                        v72 += -2147483648
                                        v74 := []int{(fi62 ^ (-(fi71))), fi62, 32768}
                                        v75 := GenS9{F0: 169, F1: v2}
                                        _ = v74
                                        _ = v75
                                    }
                                    _ = fi71
                                    _ = v72
                                }
                                if true {
                                    v70 -= (v70 * (fi62 + (v70 + (v70 | (-(fi62))))))
                                }
                                _ = v70
                            }
                        }
                    }
                } else {
                    var v76 int8 = -100
                    var v77 string = (v2 + v2)
                    if (-2147483648 != 1) {
                        v78 := v1[v2]
                        v63 = v63
                        _ = v78
                    }
                    _ = v76
                    _ = v77
                }
                var v79 uint = v63
                _ = v79
            } else {
                fmt.Println(v4, v2, v63)
                fmt.Println(v63, v2)
            }
            _ = fi62
            _ = v63
            _ = v64
        }
    }
    var v80 int = -318
    v80 += -207
    fmt.Println(v4, v80, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v59:", v59)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v59
    _ = v80
}
