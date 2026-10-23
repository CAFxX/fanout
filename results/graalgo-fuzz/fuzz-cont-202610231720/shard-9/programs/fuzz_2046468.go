package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 bool; F1 bool; F2 int }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint32; F1 uint8; F2 int64 }
type GenS7 struct { F0 int16; F1 int8 }
type GenS8 struct { F0 uint8; F1 string }
type GenS9 struct { F0 uint; F1 uint32; F2 uint16 }

func main() {
    var v1 bool = (false || (2 <= 890))
    var v2 uint32 = 202
    v3 := []int{}
    var v4 int32 = 0
    v5 := []string{(("a" + "foo bar") + "hello"), "\n\t", ("a" + "hello")}
    var v6 uint64 = 478
    for fi7 := 0; fi7 < 3; fi7++ {
        v8 := GenS1{F0: (^(v2))}
        _ = fi7
        _ = v8
    }
    var v9 uint64 = (^(v6))
    v1 = (!(42 == 9223372036854775807))
    v4 = (v4 | v4)
    var v10 int8 = 127
    if v1 {
        var v11 uint32 = (v2 ^ (v2 * 2147483647))
        _ = v11
    } else {
        if (42 > -1) {
            v12 := GenS2{F0: -60}
            var v13 float64 = (-2.25 + 1e10)
            var v14 float32 = 1.5
            _ = v12
            _ = v13
            _ = v14
        } else {
            if (-32768 < 1) {
                var v15 int = -1
                _ = v15
            }
        }
        v1 = v1
        v3 = append(v3, int(v2))
    }
    v9 = (^(v9))
    v2 = (v2 / ((v2) | 1))
    fmt.Println(v1, v6)
    if len(v5) > 0 {
        v5[0] = ""
    }
    var v16 bool = v1
    var v17 uint64 = (v9 ^ v9)
    v5 = append(v5, ("hello" + "hello"))
    var v18 string = ("\n\t" + "foo bar")
    switch 404 {
    case -453:
        switch 0 {
        case -1:
            v5 = append(v5, ("a" + v18))
        case 128:
            var v19 float32 = (float32(0.1) - 1e10)
            _ = v19
        default:
            if v1 {
                _ = len(v3)
                fmt.Println(v17, v10, v16)
                fmt.Println(v9, v6)
            }
        }
    case 2147483647:
        if ((v4 & (v4 ^ v4)) != (v4 / ((-589) | 1))) {
            if len(v3) > 0 {
                v3[0] = 1
            }
            switch -800 {
            case -80:
                v4 -= v4
            case 839:
                v18 = v18
            case 42:
                _ = len(v5)
            default:
                var v20 int8 = v10
                _ = v20
            }
            if len(v5) > 0 {
                v5[0] = v18
            }
        }
    case -2147483648:
        if v16 {
            switch ("" + ((v18 + "foo bar") + v18)) {
            case "":
                for fi21 := 0; fi21 < 14; fi21++ {
                    for fi22 := 0; fi22 < 4; fi22++ {
                        v4 -= 808
                        if v1 {
                            _ = len(v5)
                        }
                        var v23 uint8 = 1
                        _ = fi22
                        _ = v23
                    }
                    v18 = v18
                    _ = fi21
                }
            case "hello":
                var v24 int = 1
                _ = v24
            }
            var v25 string = v18
            _ = v25
        } else {
            switch 32767 {
            case -2147483648:
                v18 = "a"
            default:
                var v26 uint64 = (v9 ^ v9)
                _ = v26
            }
            if v16 {
                switch v18 {
                case "":
                    if ((v1 && (int16(v9) <= -173)) && (!(v16 && v1))) {
                        switch "" {
                        case "":
                            if len(v3) > 0 {
                                v3[0] = 474
                            }
                        default:
                            v27 := GenS3{F0: v16}
                            _ = v27
                        }
                        for fi28 := 0; fi28 < 1; fi28++ {
                            v6 = (v9 / ((v9) | 1))
                            if (1 < 65535) {
                                fmt.Println(v17, v18)
                                for fi29 := 0; fi29 < 3; fi29++ {
                                    if ((v16 && v1) && (!v1)) {
                                        var v30 uint64 = v17
                                        if v1 {
                                            var v31 uint32 = v2
                                            _ = v31
                                        } else {
                                            switch "a" {
                                            case "":
                                                if v1 {
                                                    var v32 uint32 = v2
                                                    v3 = append(v3, (fi28 >> ((1) & 63)))
                                                    _ = v32
                                                }
                                            default:
                                                v5 = append(v5, v18)
                                            }
                                            _ = len(v5)
                                        }
                                        v30 += v17
                                        _ = v30
                                    } else {
                                        v6 = 1
                                    }
                                    v3 = append(v3, fi29)
                                    _ = fi29
                                }
                            }
                            _ = fi28
                        }
                    } else {
                        if true {
                            v16 = ((v2 >= (v2 | v2)) && true)
                            var v33 int16 = -1
                            _ = v33
                        } else {
                            v16 = v1
                            v34 := GenS2{F0: 528}
                            _ = v34
                        }
                        if ((((0.0 <= 1e10) && (0.0 >= (1.5 / 1.0))) && (255 <= 65535)) || v16) {
                            v35 := GenS2{F0: 32768}
                            fmt.Println(v17, v4, v18)
                            if len(v5) > 0 {
                                v5[0] = "\n\t"
                            }
                            _ = v35
                        }
                        v36 := GenS4{F0: v1, F1: false, F2: int(v2)}
                        _ = v36
                    }
                case "a":
                    if (0 <= 1) {
                        if false {
                            var v37 int16 = 32767
                            if v1 {
                                if (!true) {
                                    if (1.5 > (1.5 * 0.0)) {
                                        v2 -= v2
                                        fmt.Println(v10)
                                        var v38 float32 = ((1e10 * 0.1) / 1.0)
                                        _ = v38
                                    }
                                    v39 := GenS3{F0: v1}
                                    _ = v39
                                } else {
                                    if len(v3) > 0 {
                                        v3[0] = 0
                                    }
                                    v17 = (v6 - v6)
                                }
                                fmt.Println(v37, v6)
                            } else {
                                fmt.Println(v17)
                                var v40 uint16 = 980
                                _ = v40
                            }
                            _ = v37
                        } else {
                            for fi41 := 0; fi41 < 6; fi41++ {
                                if v1 {
                                    if false {
                                        v5 = append(v5, (v18 + (("\n\t" + "hello") + v18)))
                                        var v42 uint32 = v2
                                        v16 = true
                                        _ = v42
                                    }
                                } else {
                                    v2 = (v2 * v2)
                                    var v43 int = fi41
                                    _ = v43
                                }
                                var v44 float64 = (0.0 + 0.1)
                                _ = fi41
                                _ = v44
                            }
                        }
                    }
                }
                v5 = append(v5, v18)
            }
        }
    }
    if v1 {
        v6 = v6
    } else {
        switch "hello" {
        case "foo bar":
            v45 := GenS2{F0: -695}
            _ = v45
        case "hello":
            var v46 uint = 2
            _ = v46
        case "a":
            if (169 >= 255) {
                v47 := GenS3{F0: (-1 >= (^(v4)))}
                var v48 uint16 = 32767
                _ = v47
                _ = v48
            } else {
                var v49 uint16 = 295
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = v49
            }
        }
    }
    if (!(1e10 <= float64((((0.1 / float32(3.14159)) / 3.14159) / 1.0)))) {
        switch 417 {
        case -589:
            v10 = (v10 / (((v10 & (v10 >> ((255) & 63)))) | 1))
        }
        for fi50 := 0; fi50 < 8; fi50++ {
            if true {
                for fi51 := 0; fi51 < 17; fi51++ {
                    switch 129 {
                    case 128:
                        var v52 uint = 205
                        if ((true || (false && false)) && v16) {
                            var v53 int16 = 0
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            fmt.Println(v17)
                            v54 := GenS5{F0: (v52 >> ((v52) & 63))}
                            _ = v53
                            _ = v54
                        } else {
                            var v55 uint8 = 128
                            _ = v55
                        }
                        if v1 {
                            v1 = (((fi51 * int(v9)) < int(v4)) || true)
                            fmt.Println(v9)
                            if ((v2 == v2) && v16) {
                                if len(v3) > 0 {
                                    v3[0] = 2147483647
                                }
                                v56 := GenS6{F0: v2, F1: 42, F2: 0}
                                switch (fi51 / ((fi51) | 1)) {
                                case 1:
                                    var v57 int = -331
                                    _ = v57
                                case 32768:
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                case 0:
                                    switch v18 {
                                    case "\n\t":
                                        var v58 int32 = (-(v4))
                                        _ = v58
                                    }
                                }
                                _ = v56
                            }
                        } else {
                            if ((v52 + 1) <= 18446744073709551615) {
                                v10 += v10
                                if len(v3) > 0 {
                                    v3[0] = 1
                                }
                                var v59 string = v18
                                var v60 int64 = 1
                                _ = v59
                                _ = v60
                            } else {
                                var v61 int64 = 1
                                _ = v61
                            }
                            fmt.Println(v9)
                            switch ("" + ((((v18 + "\n\t") + "hello") + "") + v18)) {
                            case "a":
                                if (0 > -641) {
                                    switch "\n\t" {
                                    case "\n\t":
                                        if (1 <= int64(v2)) {
                                            var v62 float64 = float64((((0.0 + -2.25) - 1.5) + (3.14159 / 1.0)))
                                            _ = v62
                                        }
                                    }
                                    v52 -= v52
                                }
                            case "\n\t":
                                var v63 uint32 = (v2 ^ (^(v2)))
                                var v64 string = v18
                                _ = len(v3)
                                _ = v63
                                _ = v64
                            case "foo bar":
                                if len(v5) > 0 {
                                    v5[0] = v18
                                }
                            }
                        }
                        v65 := GenS4{F0: true, F1: ((1.5 - float64(float32(-2.25))) >= ((float64(1e10) + 3.14159) + ((0.1 + (((-2.25 - 0.1) / 1e10) / 0.1)) - (1.5 / (-2.25 - 0.0))))), F2: (fi51 / ((127) | 1))}
                        v6 = v6
                        var v66 int64 = int64(v6)
                        v18 = v18
                        v16 = true
                        _ = v52
                        _ = v65
                        _ = v66
                    case 2147483647:
                        var v67 uint16 = 166
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        _ = len(v5)
                        _ = v67
                    }
                    if (true || false) {
                        fmt.Println(v4, v2, v18)
                        fmt.Println(v17, v1)
                        if v1 {
                            v68 := GenS4{F0: (-2.25 != 3.14159), F1: v16, F2: fi50}
                            _ = v68
                        } else {
                            v69 := GenS5{F0: 2}
                            var v70 uint8 = 202
                            _ = v69
                            _ = v70
                        }
                    }
                    var v71 int32 = (v4 + (v4 * (v4 / ((v4) | 1))))
                    _ = fi51
                    _ = v71
                }
                if v1 {
                    switch (fi50 | (fi50 ^ 128)) {
                    case 0:
                        var v72 uint64 = 18446744073709551615
                        _ = v72
                    case 32768:
                        v73 := GenS7{F0: 32767, F1: -1}
                        _ = v73
                    case -2147483648:
                        var v74 uint32 = (^(v2))
                        _ = v74
                    default:
                        if ((0.1 - -2.25) == 1.5) {
                            v3 = append(v3, fi50)
                        }
                    }
                }
                if v1 {
                    if len(v3) > 0 {
                        v3[0] = (fi50 >> ((18446744073709551615) & 63))
                    }
                    var v75 bool = v16
                    _ = v75
                } else {
                    v76 := GenS6{F0: v2, F1: 0, F2: 0}
                    v77 := GenS8{F0: 42, F1: "hello"}
                    v9 = v6
                    _ = v76
                    _ = v77
                }
            } else {
                v9 += 4294967295
            }
            v78 := GenS9{F0: 42, F1: (v2 % (((v2 | (v2 / (((v2 * v2)) | 1)))) | 1)), F2: 2}
            _ = fi50
            _ = v78
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
