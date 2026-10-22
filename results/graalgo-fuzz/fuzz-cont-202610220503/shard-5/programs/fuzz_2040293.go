package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 int8 }
type GenS2 struct { F0 uint16; F1 uint16; F2 int }
type GenS3 struct { F0 int; F1 int8 }
type GenS4 struct { F0 uint; F1 uint16 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int; F1 int32 }
type GenS7 struct { F0 int32; F1 string }
type GenS8 struct { F0 bool; F1 uint64 }
type GenS9 struct { F0 int64; F1 uint8; F2 int }
type GenS10 struct { F0 uint32; F1 uint8 }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 string; F1 int8; F2 int32 }
type GenS13 struct { F0 int64; F1 int64; F2 int64 }
type GenS14 struct { F0 uint16 }
func genFn0(p0 uint64, p1 int32) int32 {
    if false {
        if true {
            var v5 float64 = (0.1 / (float64((1.5 + float32(1.5))) - 0.0))
            p0 += p0
            fmt.Println(p1)
            _ = v5
        }
    }
    switch "foo bar" {
    case "\n\t":
        var v6 int = 862
        _ = v6
    default:
        var v7 int64 = 9223372036854775807
        _ = v7
    }
    var v8 uint64 = p0
    v8 += 4294967295
    for fi9 := 0; fi9 < 9; fi9++ {
        v10 := GenS1{F0: 2, F1: -1, F2: 45}
        for fi11 := 0; fi11 < 1; fi11++ {
            if false {
                v12 := []int{}
                p1 -= (^(p1))
                v8 = (^(v8))
                fmt.Println(fi11, p0, fi11)
                v13 := GenS1{F0: 524, F1: -9223372036854775808, F2: -1}
                v14 := GenS1{F0: 1, F1: -209, F2: 42}
                _ = v12
                _ = v13
                _ = v14
            }
            v8 += v8
            _ = fi11
        }
        v15 := GenS1{F0: 18446744073709551615, F1: -1, F2: 46}
        _ = fi9
        _ = v10
        _ = v15
    }
    _ = p0
    _ = p1
    _ = v8
    return p1
}
func genFn1(p0 int16) uint {
    v16 := []int{}
    if len(v16) > 0 {
        v16[0] = 128
    }
    _ = p0
    _ = v16
    return 15
}

func main() {
    var v1 uint8 = 2
    v2 := []int{}
    var v3 int32 = 2147483647
    var v4 uint16 = 0
    for fi17 := 0; fi17 < 19; fi17++ {
        if (false || false) {
            v18 := GenS1{F0: 104, F1: 337, F2: 1}
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            if ((v3 - genFn0(2, v3)) == v3) {
                if (255 >= (^(v4))) {
                    switch ("" + "a") {
                    case "hello":
                        if false {
                            switch fi17 {
                            case 127:
                                _ = len(v2)
                            default:
                                v19 := GenS2{F0: 65535, F1: v4, F2: fi17}
                                _ = v19
                            }
                            v3 -= v3
                        }
                    case "\n\t":
                        v1 -= v1
                    default:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    v20 := GenS2{F0: (v4 - 1), F1: v4, F2: fi17}
                    _ = v20
                } else {
                    v1 = (v1 * 211)
                    var v21 uint8 = (v1 ^ (v1 + 255))
                    var v22 uint16 = 215
                    _ = v21
                    _ = v22
                }
                fmt.Println(v1)
                for fi23 := 0; fi23 < 10; fi23++ {
                    v3 = -2147483648
                    _ = len(v2)
                    v24 := GenS3{F0: fi23, F1: 100}
                    _ = fi23
                    _ = v24
                }
            } else {
                var v25 int32 = 42
                _ = v25
            }
            _ = v18
        } else {
            var v26 int64 = -9223372036854775808
            var v27 float32 = (3.14159 * float32(-2.25))
            _ = v26
            _ = v27
        }
        for fi28 := 0; fi28 < 16; fi28++ {
            if len(v2) > 0 {
                v2[0] = (fi28 / ((4294967295) | 1))
            }
            _ = fi28
        }
        _ = fi17
    }
    v29 := GenS3{F0: -365, F1: 100}
    v3 += (v3 << ((893) & 63))
    if (255 == 1) {
        if (true || ("foo bar" < "foo bar")) {
            v30 := GenS2{F0: v4, F1: v4, F2: 32767}
            for fi31 := 0; fi31 < 9; fi31++ {
                var v32 int64 = 9223372036854775807
                v33 := GenS4{F0: 1, F1: 32767}
                _ = fi31
                _ = v32
                _ = v33
            }
            _ = v30
        }
    }
    v4 += 65535
    for fi34 := 0; fi34 < 8; fi34++ {
        var v35 uint8 = (^(v1))
        var v36 string = "\n\t"
        v1 = (^(v35))
        _ = fi34
        _ = v35
        _ = v36
    }
    v2 = append(v2, 956)
    var v37 float64 = 0.1
    v4 = 255
    for fi38 := 0; fi38 < 1; fi38++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi38
    }
    switch 128 {
    case 42:
        v4 = 2
    case 2147483647:
        var v39 uint = 255
        _ = v39
    }
    v40 := GenS5{F0: 720}
    v4 += v4
    v41 := GenS6{F0: 417, F1: (^(v3))}
    var v42 float32 = float32(v37)
    v37 += (v37 / ((0.0 - v37) - (v37 / 1.0)))
    if true {
        if true {
            for fi43 := 0; fi43 < 6; fi43++ {
                if (true && true) {
                    switch (fi43 / ((int(v1)) | 1)) {
                    case -218:
                        if false {
                            v44 := GenS4{F0: 65535, F1: v4}
                            _ = v44
                        } else {
                            if true {
                                if true {
                                    if (!(1 == 100)) {
                                        for fi45 := 0; fi45 < 6; fi45++ {
                                            fmt.Println(v37, v42, v37)
                                            var v46 bool = (-529 < 42)
                                            _ = fi45
                                            _ = v46
                                        }
                                        for fi47 := 0; fi47 < 3; fi47++ {
                                            var v48 uint = 286
                                            switch "" {
                                            case "":
                                                v49 := GenS7{F0: 897, F1: "\n\t"}
                                                _ = v49
                                            case "a":
                                                switch "foo bar" {
                                                case "foo bar":
                                                    var v50 int16 = 32767
                                                    _ = v50
                                                }
                                            case "foo bar":
                                                var v51 float64 = v37
                                                _ = v51
                                            default:
                                                if (!true) {
                                                    v52 := GenS8{F0: false, F1: 2}
                                                    v53 := GenS9{F0: 42, F1: (^(v1)), F2: (fi43 >> (((v48 | genFn1(774))) & 63))}
                                                    _ = v52
                                                    _ = v53
                                                }
                                            }
                                            v42 -= v42
                                            _ = len(v2)
                                            if true {
                                                v1 += v1
                                                fmt.Println(v48, fi47)
                                            } else {
                                                if true {
                                                    v54 := GenS10{F0: uint32(v4), F1: 200}
                                                    if false {
                                                        var v55 uint32 = 42
                                                        switch (^(fi43)) {
                                                        case -893:
                                                            v56 := GenS11{F0: false}
                                                            _ = v56
                                                        case 4294967295:
                                                            if ((fi47 + fi43) == fi47) {
                                                                v48 += uint(v55)
                                                                if true {
                                                                    v4 += (v4 | 32767)
                                                                    v4 = (^(v4))
                                                                    var v57 uint8 = 128
                                                                    _ = v57
                                                                }
                                                                var v58 float64 = 0.1
                                                                _ = v58
                                                            }
                                                        }
                                                        var v59 int16 = -434
                                                        switch fi47 {
                                                        case 0:
                                                            switch "\n\t" {
                                                            case "a":
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                            case "foo bar":
                                                                fmt.Println(fi43)
                                                            case "":
                                                                var v60 uint32 = 486
                                                                _ = v60
                                                            default:
                                                                fmt.Println(v48, v37, v55)
                                                            }
                                                        case 1:
                                                            v61 := GenS9{F0: -201, F1: 1, F2: 0}
                                                            _ = v61
                                                        case 46:
                                                            fmt.Println(v48)
                                                        }
                                                        _ = v55
                                                        _ = v59
                                                    }
                                                    _ = v54
                                                }
                                                var v62 uint32 = 65535
                                                _ = v62
                                            }
                                            _ = fi47
                                            _ = v48
                                        }
                                    } else {
                                        v63 := GenS12{F0: "", F1: 60, F2: (v3 + -76)}
                                        _ = v63
                                    }
                                }
                            } else {
                                v3 += 409
                                if len(v2) > 0 {
                                    v2[0] = fi43
                                }
                                if (!true) {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    if false {
                                        v1 += (v1 % ((1) | 1))
                                        fmt.Println(v42)
                                    } else {
                                        if false {
                                            if true {
                                                var v64 float64 = 1.5
                                                _ = v64
                                            }
                                        } else {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            var v65 float64 = v37
                                            fmt.Println(v37, v4, v3)
                                            _ = v65
                                        }
                                        if (-605 <= -898) {
                                            for fi66 := 0; fi66 < 19; fi66++ {
                                                if (v1 == v1) {
                                                    var v67 int = fi66
                                                    v68 := GenS4{F0: 818, F1: v4}
                                                    fmt.Println(v42)
                                                    _ = v67
                                                    _ = v68
                                                }
                                                _ = fi66
                                            }
                                            for fi69 := 0; fi69 < 14; fi69++ {
                                                switch fi69 {
                                                case -960:
                                                    if (1 < 86) {
                                                        v3 += genFn0(18446744073709551615, v3)
                                                        v42 = -2.25
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                    }
                                                }
                                                _ = fi69
                                            }
                                        }
                                        var v70 int = (fi43 >> ((65535) & 63))
                                        _ = v70
                                    }
                                }
                            }
                            var v71 string = ("a" + (((("\n\t" + ("" + "\n\t")) + ("hello" + "")) + "\n\t") + "foo bar"))
                            v72 := GenS1{F0: 789, F1: -9223372036854775808, F2: 127}
                            _ = v71
                            _ = v72
                        }
                    case -1:
                        v73 := GenS13{F0: 777, F1: 0, F2: 42}
                        _ = v73
                    }
                    var v74 int32 = v3
                    _ = v74
                }
                if false {
                    v1 -= v1
                    v42 -= 3.14159
                } else {
                    var v75 int32 = v3
                    _ = len(v2)
                    var v76 bool = false
                    _ = v75
                    _ = v76
                }
                switch (fi43 * fi43) {
                case 32768:
                    if (100 > 60) {
                        fmt.Println(v1, v1)
                        for fi77 := 0; fi77 < 19; fi77++ {
                            if len(v2) > 0 {
                                v2[0] = fi43
                            }
                            _ = fi77
                        }
                    }
                }
                _ = fi43
            }
            if false {
                if true {
                    if false {
                        v78 := GenS12{F0: "a", F1: 127, F2: genFn0(2, v3)}
                        _ = v78
                    } else {
                        v42 = ((float32((v37 + 1e10)) / (v42 * (v42 * ((v42 * v42) - float32(((((v37 * ((v37 * v37) + v37)) * v37) + float64(v42)) * v37)))))) * -2.25)
                        switch ("" + "foo bar") {
                        case "\n\t":
                            if (880 != genFn1(-938)) {
                                if len(v2) > 0 {
                                    v2[0] = 32768
                                }
                                switch 32767 {
                                case 1:
                                    for fi79 := 0; fi79 < 16; fi79++ {
                                        _ = len(v2)
                                        _ = fi79
                                    }
                                case 774:
                                    for fi80 := 0; fi80 < 18; fi80++ {
                                        var v81 int8 = -128
                                        var v82 int64 = 1
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        _ = fi80
                                        _ = v81
                                        _ = v82
                                    }
                                case 128:
                                    if true {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    }
                                default:
                                    v83 := GenS11{F0: (!(v1 < v1))}
                                    _ = v83
                                }
                            } else {
                                var v84 string = "hello"
                                _ = v84
                            }
                        default:
                            var v85 int = 0
                            _ = v85
                        }
                    }
                    if ((((2147483647 == 424) && true) && false) || (1 >= -9223372036854775808)) {
                        switch 787 {
                        case 32768:
                            v86 := GenS2{F0: 729, F1: (v4 << ((352) & 63)), F2: 42}
                            _ = v86
                        }
                        v42 += 1.5
                    }
                    v87 := GenS14{F0: 42}
                    _ = v87
                }
            } else {
                if true {
                    switch "foo bar" {
                    case "\n\t":
                        v88 := GenS10{F0: 0, F1: (v1 ^ v1)}
                        _ = v88
                    }
                } else {
                    fmt.Println(v3)
                }
            }
            if false {
                var v89 uint16 = v4
                _ = v89
            } else {
                v37 = (v37 - (-2.25 + ((-2.25 / (0.1 - (3.14159 * 1e10))) / v37)))
            }
        } else {
            for fi90 := 0; fi90 < 5; fi90++ {
                v3 += -1
                if true {
                    for fi91 := 0; fi91 < 8; fi91++ {
                        switch "hello" {
                        case "\n\t":
                            _ = len(v2)
                        case "":
                            var v92 uint8 = (v1 << ((uint(v1)) & 63))
                            _ = v92
                        }
                        var v93 uint8 = (^(v1))
                        var v94 int16 = 32767
                        _ = fi91
                        _ = v93
                        _ = v94
                    }
                } else {
                    v42 = v42
                }
                _ = fi90
            }
            v1 += v1
            fmt.Println(v37, v42)
        }
    } else {
        var v95 float64 = (-2.25 * float64(v42))
        _ = v95
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v37
    _ = v4
    _ = v40
    _ = v41
    _ = v42
}
