package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16; F1 uint64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int32; F1 uint16 }
type GenS6 struct { F0 bool; F1 uint8; F2 bool }
type GenS7 struct { F0 uint; F1 int32 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 uint; F1 uint16; F2 string }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 string; F1 uint16; F2 int8 }

func main() {
    var v1 int16 = 1000
    var v2 string = "a"
    var v3 uint8 = 153
    var v4 uint16 = 65535
    v5 := make(map[string]string)
    delete(v5, v2)
    v4 = (v4 % (((v4 * 42)) | 1))
    fmt.Println(v2, v2)
    v6 := []int{-2147483648, 2147483647, 42, -114}
    v4 = (v4 ^ v4)
    fmt.Println(v4, v2, v1)
    var v7 float64 = (3.14159 / ((((3.14159 + 3.14159) * ((0.1 / 3.14159) * 0.1)) - 0.0) + float64((float32(1.5) * (3.14159 * (((0.1 + 3.14159) * (1e10 / (0.1 * 1e10))) - 1.5))))))
    for fi8 := 0; fi8 < 16; fi8++ {
        if false {
            fmt.Println(v1, v3, v2)
            switch "a" {
            case "\n\t":
                if false {
                    v3 = (^(v3))
                    v3 = (v3 * (v3 & v3))
                    var v9 int64 = 42
                    _ = v9
                } else {
                    for fi10 := 0; fi10 < 14; fi10++ {
                        if (v2 != (v2 + "foo bar")) {
                            if len(v6) > 0 {
                                v6[0] = 140
                            }
                            var v11 int64 = -1
                            _ = v11
                        } else {
                            if true {
                                v2 = "a"
                                if true {
                                    var v12 uint8 = 255
                                    v1 = (v1 / (((v1 % ((v1) | 1))) | 1))
                                    fmt.Println(fi10)
                                    _ = v12
                                }
                                var v13 uint16 = v4
                                _ = v13
                            } else {
                                switch v2 {
                                case "\n\t":
                                    var v14 float32 = 0.1
                                    _ = v14
                                }
                                switch v2 {
                                case "hello":
                                    if true {
                                        if (1 == 1) {
                                            var v15 bool = false
                                            v16 := GenS1{F0: -211, F1: int64(fi8)}
                                            _ = v15
                                            _ = v16
                                        } else {
                                            v7 += v7
                                            var v17 uint64 = 1
                                            _ = v17
                                        }
                                    } else {
                                        v18 := GenS1{F0: 1, F1: 740}
                                        if (846 <= 0) {
                                            if false {
                                                var v19 int8 = 1
                                                var v20 int64 = 42
                                                v7 = (v7 + 0.0)
                                                v7 = (v7 * float64(-2.25))
                                                v1 = (v1 / ((32767) | 1))
                                                v2 = v2
                                                _ = v19
                                                _ = v20
                                            }
                                            v6 = append(v6, fi8)
                                        } else {
                                            switch fi10 {
                                            case 127:
                                                v6 = append(v6, (-(fi10)))
                                            }
                                            var v21 int8 = 100
                                            _ = v21
                                        }
                                        _ = v18
                                    }
                                    var v22 uint32 = 249
                                    _ = v22
                                case "\n\t":
                                    if (false || true) {
                                        var v23 int32 = 350
                                        if false {
                                            var v24 int16 = (v1 + (^(v1)))
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                            v2 = ("a" + "foo bar")
                                            _ = v24
                                        }
                                        _ = v23
                                    }
                                case "a":
                                    if (-2147483648 == (fi10 * fi8)) {
                                        var v25 uint = 311
                                        var v26 string = "a"
                                        _ = v25
                                        _ = v26
                                    } else {
                                        v27 := GenS2{F0: 1}
                                        _ = v27
                                    }
                                }
                                switch v2 {
                                case "\n\t":
                                    switch (fi10 * fi10) {
                                    case 127:
                                        var v28 uint8 = v3
                                        _ = v28
                                    case 1:
                                        fmt.Println(v2, v2)
                                    case 128:
                                        var v29 uint8 = v3
                                        _ = v29
                                    }
                                default:
                                    switch (fi8 % (((fi10 & fi8)) | 1)) {
                                    case -2147483648:
                                        fmt.Println(v1, fi10, v3)
                                    case -733:
                                        v30 := GenS3{F0: v1, F1: 448}
                                        _ = v30
                                    default:
                                        if ((!(false || false)) || (0 < 42)) {
                                            var v31 uint8 = 0
                                            v1 -= v1
                                            fmt.Println(v31, v31, v2)
                                            _ = v31
                                        } else {
                                            if len(v6) > 0 {
                                                v6[0] = (-(fi8))
                                            }
                                        }
                                        if false {
                                            if true {
                                                var v32 uint16 = v4
                                                var v33 float64 = v7
                                                switch v2 {
                                                case "hello":
                                                    v34 := GenS2{F0: -915}
                                                    _ = v34
                                                case "\n\t":
                                                    switch "foo bar" {
                                                    case "a":
                                                        if false {
                                                            switch (fi10 >> ((121) & 63)) {
                                                            case 0:
                                                                v2 = "foo bar"
                                                            case 461:
                                                                v7 -= 3.14159
                                                                var v35 int = (fi10 >> ((65535) & 63))
                                                                _ = v35
                                                            case 32767:
                                                                var v36 float64 = (v33 - 3.14159)
                                                                _ = v36
                                                            default:
                                                                v2 = ("hello" + v2)
                                                            }
                                                            var v37 int32 = int32(v3)
                                                            _ = v37
                                                        }
                                                    case "foo bar":
                                                        v1 -= (v1 % (((^(v1))) | 1))
                                                    case "hello":
                                                        if false {
                                                            fmt.Println(v32, v32, v7)
                                                            if true {
                                                                switch (fi10 << ((742) & 63)) {
                                                                case 9223372036854775807:
                                                                    fmt.Println(v1, v4, v7)
                                                                default:
                                                                    var v38 int = fi8
                                                                    _ = v38
                                                                }
                                                                switch "" {
                                                                case "a":
                                                                    if (-399 >= 1) {
                                                                        v1 = v1
                                                                        delete(v5, "hello")
                                                                        v39 := GenS4{F0: v2}
                                                                        _ = v39
                                                                    } else {
                                                                        v40 := GenS5{F0: -425, F1: 1}
                                                                        _ = len(v5)
                                                                        if len(v6) >= 2 {
                                                                            v6 = v6[1:len(v6)]
                                                                        }
                                                                        _ = v40
                                                                    }
                                                                    if (false && true) {
                                                                        v41 := GenS5{F0: -2147483648, F1: (^(v4))}
                                                                        fmt.Println(fi10, v4, v7)
                                                                        _ = v41
                                                                    }
                                                                default:
                                                                    if true {
                                                                        fmt.Println(v2)
                                                                        var v42 float64 = v33
                                                                        _ = v42
                                                                    }
                                                                }
                                                                if len(v6) >= 2 {
                                                                    v6 = v6[1:len(v6)]
                                                                }
                                                            } else {
                                                                if true {
                                                                    if (true || ((v2 + (((v2 + ("hello" + "foo bar")) + "") + ((((v2 + ((v2 + "a") + "\n\t")) + v2) + (((("foo bar" + (v2 + v2)) + (v2 + (("foo bar" + v2) + "foo bar"))) + v2) + v2)) + ((v2 + v2) + v2)))) < v2)) {
                                                                        if false {
                                                                            v4 -= (^(v32))
                                                                            v1 = (v1 % ((395) | 1))
                                                                        } else {
                                                                            if false {
                                                                                if (!(true || (4294967295 > 4294967295))) {
                                                                                    if true {
                                                                                        if (4294967295 == 2147483648) {
                                                                                            var v43 uint8 = (v3 - v3)
                                                                                            v6 = append(v6, (fi8 - fi10))
                                                                                            _ = v43
                                                                                        } else {
                                                                                            v3 = v3
                                                                                        }
                                                                                        v44 := GenS6{F0: false, F1: (v3 - v3), F2: false}
                                                                                        var v45 int32 = -1
                                                                                        _ = v44
                                                                                        _ = v45
                                                                                    } else {
                                                                                        if len(v6) >= 2 {
                                                                                            v6 = v6[1:len(v6)]
                                                                                        }
                                                                                        fmt.Println(fi10, v7)
                                                                                    }
                                                                                    v3 -= 255
                                                                                    v2 = v2
                                                                                }
                                                                                var v46 uint = 2
                                                                                _ = v46
                                                                            }
                                                                            v47 := GenS7{F0: 2, F1: 2147483647}
                                                                            _ = v47
                                                                        }
                                                                        if len(v6) > 0 {
                                                                            v6[0] = (^(fi10))
                                                                        }
                                                                    }
                                                                } else {
                                                                    var v48 uint32 = 2147483647
                                                                    v33 += -2.25
                                                                    _ = v48
                                                                }
                                                                _ = len(v5)
                                                                _ = len(v6)
                                                                fmt.Println(v1, fi8, v3)
                                                                fmt.Println(v4, v2, v1)
                                                            }
                                                        } else {
                                                            var v49 int = -729
                                                            _ = v49
                                                        }
                                                    default:
                                                        var v50 float32 = ((-2.25 - -2.25) / 3.14159)
                                                        _ = v50
                                                    }
                                                case "":
                                                    if true {
                                                        if len(v6) > 0 {
                                                            v6[0] = (fi10 << ((255) & 63))
                                                        }
                                                    }
                                                }
                                                _ = v32
                                                _ = v33
                                            } else {
                                                v51 := GenS6{F0: (!false), F1: v3, F2: true}
                                                _ = v51
                                            }
                                        }
                                        var v52 string = v2
                                        _ = v52
                                    }
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    switch "\n\t" {
                                    case "\n\t":
                                        switch (("a" + (v2 + v2)) + v2) {
                                        case "foo bar":
                                            if len(v6) > 0 {
                                                v6[0] = fi8
                                            }
                                        case "a":
                                            if true {
                                                switch (fi8 & int(v1)) {
                                                case -2147483648:
                                                    var v53 bool = false
                                                    _ = v53
                                                case 42:
                                                    if (2147483647 < -2147483648) {
                                                        if (0.1 != 0.0) {
                                                            v54 := GenS1{F0: -1, F1: int64(fi10)}
                                                            _ = v54
                                                        } else {
                                                            v55 := GenS8{F0: "hello"}
                                                            _ = v55
                                                        }
                                                        if ((!true) && false) {
                                                            if (!((v1 == (v1 - -1)) && false)) {
                                                                if ((((1.5 - -2.25) - 0.1) - (1e10 + 3.14159)) <= float32((v7 - ((v7 / (1.5 * v7)) * v7)))) {
                                                                    v56 := GenS6{F0: (1 > -668), F1: 42, F2: (true && false)}
                                                                    var v57 int32 = -1
                                                                    fmt.Println(v57, v1)
                                                                    var v58 float64 = float64(((0.0 - 1e10) + 1.5))
                                                                    v4 += v4
                                                                    _ = v56
                                                                    _ = v57
                                                                    _ = v58
                                                                }
                                                            }
                                                            var v59 uint64 = 42
                                                            _ = v59
                                                        }
                                                    }
                                                default:
                                                    v60 := GenS3{F0: v1, F1: 4294967295}
                                                    _ = v60
                                                }
                                                v7 = v7
                                            }
                                        }
                                    case "foo bar":
                                        fmt.Println(v2, v7, v4)
                                    case "":
                                        switch (("foo bar" + "\n\t") + v2) {
                                        case "hello":
                                            if (-59 != 46) {
                                                var v61 uint32 = 4294967295
                                                v62 := GenS7{F0: 860, F1: -547}
                                                var v63 uint32 = (v61 & v61)
                                                v64 := GenS9{F0: 164, F1: v4, F2: ("hello" + (v2 + v2))}
                                                v1 -= v1
                                                fmt.Println(v2, v1)
                                                _ = v61
                                                _ = v62
                                                _ = v63
                                                _ = v64
                                            } else {
                                                v65 := GenS6{F0: (1 <= 4294967295), F1: v3, F2: (v4 != (^(v4)))}
                                                fmt.Println(v4)
                                                if ((v3 | v3) >= v3) {
                                                    if (863 > 1) {
                                                        switch fi8 {
                                                        case 32768:
                                                            delete(v5, (("\n\t" + v2) + v2))
                                                        case 32767:
                                                            fmt.Println(v4, v3, v1)
                                                        }
                                                        var v66 string = (((v2 + (v2 + (v2 + v2))) + v2) + ("foo bar" + v2))
                                                        v66 = ("" + (v2 + v2))
                                                        _ = v66
                                                    }
                                                    fmt.Println(v4, v1, fi10)
                                                } else {
                                                    switch "foo bar" {
                                                    case "a":
                                                        v67 := GenS3{F0: v1, F1: 4294967295}
                                                        _ = v67
                                                    case "\n\t":
                                                        switch fi10 {
                                                        case 32768:
                                                            v68 := GenS10{F0: v3}
                                                            _ = v68
                                                        }
                                                    case "":
                                                        v3 = v3
                                                    }
                                                    v7 = (v7 * (float64(1e10) + float64(0.0)))
                                                    v2 = v2
                                                }
                                                _ = v65
                                            }
                                        case "foo bar":
                                            v2 = (v2 + v2)
                                        default:
                                            fmt.Println(fi8, fi10, v7)
                                        }
                                        _ = len(v5)
                                        v3 = v3
                                    default:
                                        switch v2 {
                                        case "hello":
                                            v2 = (("\n\t" + v2) + v2)
                                        }
                                    }
                                }
                            }
                            var v69 float32 = 3.14159
                            v2 = ((v2 + v2) + (((v2 + "") + ("foo bar" + (v2 + v2))) + v2))
                            _ = v69
                        }
                        v70 := GenS7{F0: uint(v3), F1: 427}
                        v71 := GenS11{F0: "foo bar", F1: 1, F2: -51}
                        _ = fi10
                        _ = v70
                        _ = v71
                    }
                    v72 := GenS3{F0: (v1 & (-(v1))), F1: 815}
                    var v73 int = 128
                    _ = v72
                    _ = v73
                }
            case "":
                for fi74 := 0; fi74 < 6; fi74++ {
                    v5[((v2 + v2) + v2)] = ""
                    var v75 int16 = -1
                    var v76 float32 = (-2.25 + 0.1)
                    var v77 int = fi8
                    _ = fi74
                    _ = v75
                    _ = v76
                    _ = v77
                }
            }
            v2 = v2
        }
        v2 = v2
        _ = fi8
    }
    var v78 int8 = 100
    switch v2 {
    case "":
        v79 := GenS4{F0: (v2 + "\n\t")}
        _ = v79
    }
    switch v2 {
    case "hello":
        fmt.Println(v1)
    case "\n\t":
        fmt.Println(v1, v2)
    case "":
        switch 1 {
        case -1:
            var v80 uint = 506
            _ = v80
        case 128:
            var v81 bool = ((v3 % (((v3 ^ v3)) | 1)) != 35)
            _ = v81
        case 966:
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        default:
            v78 -= v78
        }
    }
    v5[v2] = v2
    v2 = "a"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v78
}
