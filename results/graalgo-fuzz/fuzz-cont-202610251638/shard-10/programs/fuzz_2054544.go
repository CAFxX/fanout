package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool; F1 int32 }
type GenS4 struct { F0 uint8; F1 int8; F2 int16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int16; F1 int64 }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint32 = 814
    v3 := []int{-133, 905, -122}
    var v4 uint16 = 255
    v5 := make(map[string]bool)
    var v6 string = (("a" + "hello") + "")
    switch "hello" {
    case "":
        v2 = v2
    case "\n\t":
        v2 -= v2
    case "hello":
        if (2 == 1) {
            for fi7 := 0; fi7 < 2; fi7++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi7
            }
        } else {
            v3 = append(v3, 32768)
            var v8 float32 = ((float32(float64(float32(((3.14159 * (1e10 + -2.25)) * ((((1.5 / (0.0 + -2.25)) + 1e10) + -2.25) / -2.25))))) + 1.5) * ((0.1 * float32(1.5)) + (-2.25 + 0.1)))
            _ = v8
        }
    default:
        v9 := GenS1{F0: -32768, F1: 1}
        _ = v9
    }
    switch 42 {
    case -1:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case 303:
        v2 = v2
    default:
        switch 2147483647 {
        case -520:
            var v10 int64 = 1
            _ = v10
        default:
            if (928 < -2147483648) {
                if false {
                    var v11 float32 = (float32(3.14159) / 1.0)
                    _ = len(v1)
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = v11
                } else {
                    var v12 int32 = 0
                    var v13 uint = 511
                    if (v12 == (v12 >> (((v13 + (v13 / ((1) | 1)))) & 63))) {
                        v14 := GenS1{F0: -137, F1: (v2 + (v2 ^ v2))}
                        var v15 uint8 = 128
                        _ = v14
                        _ = v15
                    } else {
                        for fi16 := 0; fi16 < 11; fi16++ {
                            v17 := GenS2{F0: (false || (v2 != (v2 + (^(v2)))))}
                            for fi18 := 0; fi18 < 3; fi18++ {
                                var v19 int64 = 999
                                var v20 bool = (18446744073709551615 <= 60)
                                _ = fi18
                                _ = v19
                                _ = v20
                            }
                            _ = fi16
                            _ = v17
                        }
                        switch ((v6 + v6) + (((("" + v6) + v6) + ("" + "\n\t")) + ((v6 + "") + ((((("a" + "a") + v6) + "") + "\n\t") + "hello")))) {
                        case "":
                            switch 4294967295 {
                            case 2147483647:
                                fmt.Println(v12, v4)
                            case 4294967295:
                                for fi21 := 0; fi21 < 17; fi21++ {
                                    if true {
                                        v22 := GenS1{F0: 32767, F1: v2}
                                        _ = v22
                                    } else {
                                        fmt.Println(v4)
                                        v23 := GenS3{F0: false, F1: v12}
                                        _ = v23
                                    }
                                    v24 := GenS3{F0: (v13 <= 419), F1: 1}
                                    v6 = "foo bar"
                                    _ = fi21
                                    _ = v24
                                }
                            case 32768:
                                var v25 string = v6
                                _ = v25
                            }
                        case "\n\t":
                            fmt.Println(v13)
                        default:
                            var v26 float64 = (((0.1 - (((0.1 / (1e10 + (0.1 - ((1.5 + ((1e10 + 1e10) - -2.25)) / -2.25)))) - 0.0) * 3.14159)) + -2.25) - ((0.1 + (1e10 - 3.14159)) / 1.0))
                            _ = v26
                        }
                    }
                    _ = v12
                    _ = v13
                }
            }
        }
    }
    if ((!true) && false) {
        fmt.Println(v4, v6, v4)
        var v27 int = -572
        _ = v27
    }
    _ = len(v5)
    var v28 int8 = -128
    if (1 <= 2) {
        fmt.Println(v4, v4, v4)
        var v29 uint64 = 1
        if len(v3) > 0 {
            v3[0] = -9223372036854775808
        }
        _ = v29
    }
    v30 := GenS3{F0: true, F1: int32(v28)}
    v4 += 255
    for fi31 := 0; fi31 < 5; fi31++ {
        for fi32 := 0; fi32 < 9; fi32++ {
            switch 127 {
            case -2147483648:
                v33 := GenS4{F0: 127, F1: v28, F2: 32767}
                _ = v33
            case 127:
                if true {
                    v2 = v2
                    v28 = v28
                    v4 -= v4
                } else {
                    var v34 float32 = 0.1
                    _ = v34
                }
            default:
                switch (^(fi31)) {
                case -426:
                    if (v6 <= (v6 + "")) {
                        var v35 uint32 = v2
                        var v36 float32 = float32(((1e10 * float64(1e10)) / 1.0))
                        switch 32767 {
                        case 127:
                            v37 := GenS1{F0: 32767, F1: 270}
                            _ = v37
                        }
                        fmt.Println(v4)
                        _ = v35
                        _ = v36
                    } else {
                        v6 = v6
                        v38 := GenS2{F0: (!((!((v2 > v2) || true)) || (!(127 > 1))))}
                        v39 := GenS4{F0: 42, F1: v28, F2: 106}
                        _ = v38
                        _ = v39
                    }
                default:
                    if len(v3) > 0 {
                        v3[0] = (fi31 | int(v2))
                    }
                }
            }
            _ = fi32
        }
        _ = fi31
    }
    var v40 uint8 = 128
    fmt.Println(v40, v6, v28)
    v3 = append(v3, -2147483648)
    var v41 uint8 = (^(v40))
    var v42 uint16 = (v4 << ((724) & 63))
    for fi43 := 0; fi43 < 17; fi43++ {
        v6 = v6
        _ = fi43
    }
    if true {
        switch 583 {
        case 1:
            v41 = 255
        case 32768:
            v44 := GenS1{F0: 1, F1: 4294967295}
            _ = v44
        default:
            if true {
                switch 0 {
                case 42:
                    var v45 uint16 = 2
                    _ = v45
                case -9223372036854775808:
                    var v46 bool = false
                    _ = v46
                }
                v28 -= int8(v42)
                switch 128 {
                case 32768:
                    for fi47 := 0; fi47 < 6; fi47++ {
                        if false {
                            var v48 int32 = -2147483648
                            var v49 int8 = (v28 / (((v28 << ((65535) & 63))) | 1))
                            v6 = v6
                            _ = v48
                            _ = v49
                        } else {
                            v50 := GenS3{F0: (false || (!false)), F1: -1}
                            switch v6 {
                            case "hello":
                                if false {
                                    v51 := GenS5{F0: ""}
                                    if (true && true) {
                                        var v52 float64 = 0.0
                                        var v53 int16 = 1
                                        if false {
                                            if true {
                                                for fi54 := 0; fi54 < 11; fi54++ {
                                                    var v55 float32 = (float32((v52 / 1.5)) * -2.25)
                                                    v42 += 32767
                                                    _ = fi54
                                                    _ = v55
                                                }
                                            } else {
                                                v56 := GenS5{F0: v6}
                                                v41 = 42
                                                for fi57 := 0; fi57 < 6; fi57++ {
                                                    v52 = (3.14159 - 1.5)
                                                    v2 += 273
                                                    switch ("hello" + (v6 + v6)) {
                                                    case "foo bar":
                                                        var v58 int16 = v53
                                                        _ = v58
                                                    case "":
                                                        v6 = v6
                                                    default:
                                                        var v59 int32 = 227
                                                        _ = v59
                                                    }
                                                    _ = fi57
                                                }
                                                _ = v56
                                            }
                                        }
                                        _ = v52
                                        _ = v53
                                    } else {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        v3 = append(v3, fi47)
                                    }
                                    v60 := GenS5{F0: v6}
                                    _ = v51
                                    _ = v60
                                } else {
                                    var v61 uint16 = v4
                                    var v62 int8 = (v28 + 49)
                                    switch 42 {
                                    case 0:
                                        var v63 int16 = 0
                                        _ = v63
                                    case -961:
                                        v62 += (v28 & v28)
                                    }
                                    _ = v61
                                    _ = v62
                                }
                            }
                            var v64 int32 = 54
                            _ = v50
                            _ = v64
                        }
                        var v65 uint32 = (v2 << ((139) & 63))
                        if true {
                            v2 -= v2
                            for fi66 := 0; fi66 < 15; fi66++ {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                if true {
                                    fmt.Println(v42, v6)
                                    if (v6 < "hello") {
                                        var v67 uint16 = v42
                                        v68 := GenS5{F0: v6}
                                        v69 := GenS1{F0: 1, F1: v2}
                                        _ = v67
                                        _ = v68
                                        _ = v69
                                    }
                                    var v70 int8 = (v28 << ((65535) & 63))
                                    _ = v70
                                } else {
                                    v71 := GenS2{F0: (true || false)}
                                    v41 += v40
                                    var v72 uint32 = 0
                                    if len(v3) > 0 {
                                        v3[0] = fi47
                                    }
                                    _ = len(v5)
                                    if (!false) {
                                        switch "a" {
                                        case "":
                                            var v73 bool = false
                                            v41 += v41
                                            fmt.Println(v42)
                                            _ = v73
                                        case "\n\t":
                                            v40 -= 128
                                        default:
                                            v74 := GenS6{F0: 1}
                                            _ = v74
                                        }
                                    } else {
                                        switch (("hello" + v6) + "\n\t") {
                                        case "\n\t":
                                            var v75 float32 = 0.0
                                            _ = v75
                                        }
                                        v76 := v5[v6]
                                        v77 := GenS2{F0: v76}
                                        _ = v76
                                        _ = v77
                                    }
                                    if false {
                                        v28 -= int8(v4)
                                        if (65535 >= 456) {
                                            v4 -= (v4 - v42)
                                        }
                                        v42 = 2
                                    }
                                    _ = v71
                                    _ = v72
                                }
                                _ = fi66
                            }
                            v78 := GenS2{F0: false}
                            _ = v78
                        } else {
                            var v79 float32 = ((float32(0.1) + (0.0 / 3.14159)) + (0.0 - -2.25))
                            if len(v3) > 0 {
                                v3[0] = fi47
                            }
                            _ = v79
                        }
                        _ = fi47
                        _ = v65
                    }
                case 127:
                    v80 := GenS3{F0: false, F1: 1}
                    _ = v80
                case -867:
                    for fi81 := 0; fi81 < 4; fi81++ {
                        var v82 uint = 4294967295
                        _ = fi81
                        _ = v82
                    }
                }
            } else {
                for fi83 := 0; fi83 < 3; fi83++ {
                    fmt.Println(fi83, v28, v42)
                    switch 32767 {
                    case 4294967295:
                        for fi84 := 0; fi84 < 7; fi84++ {
                            fmt.Println(v42, v4, v2)
                            _ = len(v1)
                            _ = fi84
                        }
                    case 9223372036854775807:
                        for fi85 := 0; fi85 < 11; fi85++ {
                            if false {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                var v86 uint32 = v2
                                _ = v86
                            }
                            switch fi85 {
                            case -9223372036854775808:
                                fmt.Println(fi85)
                            }
                            if (0 > -1) {
                                v87 := GenS4{F0: v40, F1: v28, F2: 42}
                                fmt.Println(fi85, v28, v40)
                                _ = v87
                            }
                            _ = fi85
                        }
                    default:
                        v88 := GenS7{F0: 32767, F1: -9223372036854775808}
                        _ = v88
                    }
                    _ = fi83
                }
            }
        }
        v89 := GenS5{F0: v6}
        var v90 int32 = 339
        _ = v89
        _ = v90
    }
    var v91 uint = 4294967295
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v30
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
    _ = v6
    _ = v91
}
