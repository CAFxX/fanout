package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 uint32; F2 bool }
type GenS3 struct { F0 uint; F1 uint64; F2 uint32 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int64; F1 int64 }
type GenS6 struct { F0 int16; F1 int64; F2 bool }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int32; F1 uint16; F2 uint16 }
type GenS9 struct { F0 uint8; F1 uint32; F2 int }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 int; F1 int16 }
type GenS12 struct { F0 string; F1 string }
type GenS13 struct { F0 int; F1 int32 }
func genFn0(p0 uint64, p1 int8, p2 int16) string {
    p1 = 1
    var v4 float64 = 3.14159
    var v5 float64 = v4
    if (!(!true)) {
        var v6 uint8 = 0
        v6 += v6
        _ = v6
    } else {
        v5 = float64(0.0)
        switch "" {
        case "\n\t":
            v5 = (v5 * 1e10)
        }
        v7 := []int{1, 127, -1, 32768}
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return "hello"
}
func genFn1() string {
    var v8 int64 = 9223372036854775807
    v9 := GenS1{F0: 2}
    _ = v8
    _ = v9
    return ""
}

func main() {
    var v1 uint32 = 2
    v2 := []uint32{}
    var v3 int32 = -2147483648
    v10 := GenS1{F0: v1}
    switch 127 {
    case -4:
        if true {
            v11 := GenS1{F0: (^(v1))}
            v12 := GenS1{F0: 2}
            if true {
                var v13 uint32 = v1
                if (32767 >= 191) {
                    if len(v2) > 0 {
                        v2[0] = v13
                    }
                    if len(v2) > 0 {
                        v2[0] = v13
                    }
                    var v14 int32 = (v3 >> ((666) & 63))
                    _ = v14
                }
                var v15 string = "foo bar"
                _ = v13
                _ = v15
            } else {
                switch ("\n\t" + (("hello" + "") + (("\n\t" + "") + "\n\t"))) {
                case "a":
                    var v16 float64 = ((0.1 - ((3.14159 / float64(((1e10 + 0.1) + float32(0.0)))) - 1e10)) - float64((-2.25 / 1.0)))
                    _ = v16
                case "\n\t":
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                var v17 uint32 = v1
                v3 += v3
                _ = v17
            }
            _ = v11
            _ = v12
        } else {
            var v18 int32 = v3
            v18 = v18
            var v19 uint = 816
            _ = v18
            _ = v19
        }
    default:
        fmt.Println(v3, v3)
    }
    var v20 uint32 = v1
    fmt.Println(v20, v20)
    for fi21 := 0; fi21 < 7; fi21++ {
        v3 = 0
        _ = fi21
    }
    for fi22 := 0; fi22 < 16; fi22++ {
        fmt.Println(v20)
        var v23 int16 = -32768
        _ = fi22
        _ = v23
    }
    if ("foo bar" <= "hello") {
        v3 = (v3 << ((90) & 63))
        if ((!true) && false) {
            _ = len(v2)
        }
    }
    fmt.Println(v20, v20, v20)
    v24 := GenS1{F0: v1}
    if (568 > 42) {
        v25 := GenS1{F0: (v20 & v20)}
        _ = v25
    }
    fmt.Println(v20, v1, v1)
    switch ("hello" + "foo bar") {
    case "a":
        var v26 uint = 18446744073709551615
        _ = v26
    case "":
        if false {
            v3 += v3
            var v27 float64 = 0.0
            var v28 int8 = -1
            _ = v27
            _ = v28
        }
    case "foo bar":
        var v29 int32 = v3
        _ = v29
    default:
        v3 = 2147483647
    }
    if len(v2) > 0 {
        v2[0] = 2
    }
    for fi30 := 0; fi30 < 4; fi30++ {
        for fi31 := 0; fi31 < 18; fi31++ {
            if ((v3 / (((v3 & (v3 - (^(v3))))) | 1)) > -1) {
                var v32 int16 = 1
                _ = v32
            } else {
                var v33 uint32 = 358
                v34 := GenS2{F0: 311, F1: v20, F2: (!false)}
                fmt.Println(fi31, v1, v1)
                _ = v33
                _ = v34
            }
            var v35 int16 = 42
            v36 := GenS3{F0: 255, F1: 2, F2: v20}
            _ = fi31
            _ = v35
            _ = v36
        }
        if true {
            v20 = v20
            switch (fi30 * (fi30 + fi30)) {
            case 128:
                v37 := GenS4{F0: -1}
                _ = v37
            case 42:
                for fi38 := 0; fi38 < 9; fi38++ {
                    if len(v2) > 0 {
                        v2[0] = v1
                    }
                    _ = fi38
                }
            case 1:
                for fi39 := 0; fi39 < 13; fi39++ {
                    switch (fi30 * (fi30 * (fi30 - 1))) {
                    case -549:
                        var v40 float64 = 3.14159
                        _ = v40
                    case 127:
                        v2 = append(v2, v1)
                    case 32767:
                        var v41 float32 = 1.5
                        _ = v41
                    default:
                        v3 -= v3
                        var v42 uint64 = 2
                        fmt.Println(fi39)
                        _ = v42
                    }
                    _ = fi39
                }
            default:
                v1 = (^(v20))
            }
        } else {
            if (true || (32767 < 32767)) {
                for fi43 := 0; fi43 < 3; fi43++ {
                    fmt.Println(v1, v1, v1)
                    v3 = -430
                    _ = fi43
                }
            } else {
                v44 := GenS5{F0: 0, F1: 9223372036854775807}
                v45 := GenS3{F0: 822, F1: 215, F2: (v1 * v1)}
                _ = v44
                _ = v45
            }
        }
        if (36 == -1) {
            v3 -= (v3 >> ((0) & 63))
            var v46 int32 = 0
            var v47 int32 = (^(v46))
            _ = v46
            _ = v47
        } else {
            var v48 uint16 = 32767
            _ = v48
        }
        _ = fi30
    }
    v3 += v3
    v49 := GenS3{F0: 65535, F1: 1, F2: 211}
    v3 = (v3 / ((v3) | 1))
    for fi50 := 0; fi50 < 7; fi50++ {
        v20 = v1
        if (float32(1e10) != float32(1.5)) {
            for fi51 := 0; fi51 < 2; fi51++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi51
            }
        } else {
            if true {
                for fi52 := 0; fi52 < 10; fi52++ {
                    switch "foo bar" {
                    case "\n\t":
                        switch "" {
                        case "\n\t":
                            if false {
                                v3 = (v3 - v3)
                            }
                        }
                    case "":
                        v3 = v3
                        fmt.Println(fi50, fi50, v1)
                        var v53 uint16 = 42
                        fmt.Println(v20)
                        _ = v53
                    default:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    _ = fi52
                }
                switch ("hello" + "foo bar") {
                case "a":
                    v3 += (v3 + 262)
                case "hello":
                    if false {
                        _ = len(v2)
                        switch "foo bar" {
                        case "\n\t":
                            switch (fi50 * (-(fi50))) {
                            case -2147483648:
                                v3 = 0
                            case 32768:
                                v1 -= v1
                            case -9223372036854775808:
                                if (-833 != 668) {
                                    switch ("\n\t" + "hello") {
                                    case "foo bar":
                                        v20 += v20
                                    }
                                } else {
                                    var v54 int64 = 42
                                    _ = v54
                                }
                            }
                        default:
                            v55 := GenS6{F0: 42, F1: 1, F2: (v3 <= (v3 / (((v3 + 1)) | 1)))}
                            _ = v55
                        }
                    }
                case "\n\t":
                    for fi56 := 0; fi56 < 16; fi56++ {
                        switch fi50 {
                        case 0:
                            v1 -= v1
                        case -9223372036854775808:
                            switch 214 {
                            case -13:
                                v20 -= 517
                            }
                        }
                        if (true || (!false)) {
                            v57 := GenS3{F0: 65535, F1: 18446744073709551615, F2: (v1 + 2147483648)}
                            _ = v57
                        } else {
                            v3 += (v3 / ((v3) | 1))
                            var v58 float64 = (0.1 / (0.0 - 0.1))
                            if ((((-(v3)) != 1) && (-32768 != 1)) || true) {
                                if false {
                                    fmt.Println(v58)
                                } else {
                                    v59 := GenS7{F0: 441}
                                    if (-2.25 == 3.14159) {
                                        if len(v2) > 0 {
                                            v2[0] = 16
                                        }
                                        fmt.Println(fi56, v1, fi50)
                                    } else {
                                        if len(v2) > 0 {
                                            v2[0] = 4294967295
                                        }
                                        switch -1 {
                                        case 10:
                                            v1 = v20
                                        case 783:
                                            v60 := GenS8{F0: v3, F1: 255, F2: 32767}
                                            _ = v60
                                        }
                                        if (true || true) {
                                            var v61 int = fi56
                                            switch "foo bar" {
                                            case "foo bar":
                                                var v62 int16 = -1
                                                _ = v62
                                            case "hello":
                                                v63 := GenS2{F0: 1, F1: v20, F2: ((!(1 > -731)) && (!false))}
                                                _ = v63
                                            case "\n\t":
                                                v64 := GenS9{F0: 128, F1: 2, F2: (fi56 * (v61 + (v61 << ((0) & 63))))}
                                                var v65 int64 = 0
                                                _ = v64
                                                _ = v65
                                            default:
                                                v2 = append(v2, 2147483648)
                                            }
                                            v66 := GenS10{F0: 42}
                                            _ = v61
                                            _ = v66
                                        } else {
                                            v67 := GenS4{F0: -569}
                                            v68 := GenS11{F0: (-(fi50)), F1: 1}
                                            v3 = v3
                                            _ = v67
                                            _ = v68
                                        }
                                    }
                                    if true {
                                        v2 = append(v2, (v1 << ((65535) & 63)))
                                        v69 := GenS7{F0: 0}
                                        _ = v69
                                    }
                                    _ = v59
                                }
                            }
                            _ = v58
                        }
                        _ = fi56
                    }
                }
                if true {
                    v20 = 2147483648
                    v3 = (v3 - 42)
                    v70 := GenS12{F0: ("" + ("" + ((("foo bar" + "") + "hello") + "a"))), F1: "hello"}
                    _ = v70
                } else {
                    fmt.Println(v3, v1)
                    var v71 float32 = ((0.0 * (float32((0.1 * 3.14159)) - 1e10)) / 1.5)
                    if (2 <= 255) {
                        v72 := GenS13{F0: (fi50 - -1), F1: 0}
                        v20 -= v20
                        var v73 int16 = -1
                        _ = v72
                        _ = v73
                    } else {
                        switch (fi50 / ((fi50) | 1)) {
                        case 4294967295:
                            v74 := GenS8{F0: v3, F1: 556, F2: 0}
                            _ = v74
                        case -947:
                            if false {
                                if (255 > 255) {
                                    fmt.Println(fi50)
                                    var v75 uint16 = 614
                                    v2 = append(v2, 65535)
                                    _ = v75
                                }
                                v71 = v71
                            }
                        }
                    }
                    _ = v71
                }
            } else {
                v3 = (^(v3))
            }
        }
        _ = fi50
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v49
}
