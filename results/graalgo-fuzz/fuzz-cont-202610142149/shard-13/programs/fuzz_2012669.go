package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 uint8 }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 int8; F1 uint64; F2 int8 }
type GenS11 struct { F0 uint8; F1 int8 }
func genFn0(p0 int8, p1 bool) uint8 {
    var v5 float64 = (1.5 - ((3.14159 - -2.25) / (1.5 + 0.0)))
    for fi6 := 0; fi6 < 16; fi6++ {
        fmt.Println(fi6, p1, p1)
        fmt.Println(p0, fi6)
        _ = fi6
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v5
    return 2
}

func main() {
    var v1 uint = 456
    v2 := []string{"foo bar", ("foo bar" + ("a" + "\n\t")), "a"}
    var v3 int32 = -1
    v4 := []int{32768, 9223372036854775807, -111}
    fmt.Println(v3)
    var v7 uint8 = 127
    v7 -= (v7 - v7)
    v8 := GenS1{F0: uint8(v1), F1: v1, F2: genFn0(1, true)}
    switch ("a" + "a") {
    case "\n\t":
        v3 = v3
    case "a":
        for fi9 := 0; fi9 < 9; fi9++ {
            v10 := GenS1{F0: v7, F1: (v1 >> (((^(v1))) & 63)), F2: genFn0(-128, true)}
            v7 = 127
            _ = fi9
            _ = v10
        }
    default:
        switch 9223372036854775807 {
        case 4294967295:
            v11 := GenS2{F0: "a", F1: 1}
            _ = v11
        }
    }
    if (true && true) {
        v3 = (v3 & (^(v3)))
    }
    v1 -= (^(v1))
    v1 = 65535
    switch 128 {
    case 44:
        for fi12 := 0; fi12 < 5; fi12++ {
            for fi13 := 0; fi13 < 12; fi13++ {
                var v14 uint16 = 729
                var v15 uint8 = genFn0(-21, (1e10 > 1e10))
                var v16 bool = false
                _ = fi13
                _ = v14
                _ = v15
                _ = v16
            }
            switch "a" {
            case "\n\t":
                if (!true) {
                    switch ("\n\t" + "a") {
                    case "foo bar":
                        var v17 int32 = v3
                        _ = v17
                    default:
                        v1 = (^(v1))
                    }
                    v1 = (^(v1))
                } else {
                    var v18 uint16 = 497
                    _ = v18
                }
            case "":
                v7 += 128
            }
            _ = fi12
        }
    }
    var v19 float32 = (float32(1.5) + 0.0)
    switch "\n\t" {
    case "hello":
        if (uint(v7) != v1) {
            var v20 int = 127
            v21 := GenS3{F0: (v20 >> ((v1) & 63))}
            _ = v20
            _ = v21
        } else {
            v22 := GenS4{F0: -794}
            for fi23 := 0; fi23 < 12; fi23++ {
                v24 := GenS5{F0: 65535}
                if (!(1.5 != (-2.25 + float32(1.5)))) {
                    switch (fi23 & fi23) {
                    case 9223372036854775807:
                        v2 = append(v2, "a")
                    case 0:
                        switch "a" {
                        case "":
                            v19 += v19
                        case "hello":
                            v7 += 237
                        case "\n\t":
                            v19 -= 3.14159
                        default:
                            var v25 bool = (!false)
                            _ = v25
                        }
                    }
                    v3 = -937
                } else {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    var v26 int = 128
                    if false {
                        fmt.Println(v19, v1, v19)
                    } else {
                        var v27 uint = v1
                        for fi28 := 0; fi28 < 3; fi28++ {
                            v1 = 255
                            _ = fi28
                        }
                        switch "foo bar" {
                        case "foo bar":
                            for fi29 := 0; fi29 < 14; fi29++ {
                                v7 -= (v7 + v7)
                                _ = fi29
                            }
                        case "":
                            fmt.Println(v27, v7)
                        case "\n\t":
                            if len(v4) > 0 {
                                v4[0] = (fi23 >> (((v27 - 255)) & 63))
                            }
                        }
                        _ = v27
                    }
                    _ = v26
                }
                switch "foo bar" {
                case "foo bar":
                    switch "a" {
                    case "":
                        if false {
                            var v30 int8 = -86
                            _ = len(v2)
                            var v31 uint64 = 0
                            _ = v30
                            _ = v31
                        }
                    case "a":
                        for fi32 := 0; fi32 < 4; fi32++ {
                            if true {
                                if len(v4) > 0 {
                                    v4[0] = 1
                                }
                                v7 -= 1
                            } else {
                                v19 = 1.5
                                var v33 string = "a"
                                _ = v33
                            }
                            if true {
                                if (42 < -1) {
                                    if true {
                                        switch fi23 {
                                        case 0:
                                            switch (-(fi32)) {
                                            case -656:
                                                var v34 int8 = 100
                                                _ = v34
                                            }
                                        case -9223372036854775808:
                                            if len(v2) > 0 {
                                                v2[0] = "foo bar"
                                            }
                                        case 946:
                                            if ((3.14159 - 0.1) == ((((3.14159 / 1.5) + 0.0) - -2.25) + (((3.14159 - float64(float32(0.0))) - 0.0) - -2.25))) {
                                                switch (fi32 + fi32) {
                                                case 22:
                                                    if true {
                                                        fmt.Println(v19)
                                                    } else {
                                                        v35 := GenS6{F0: 4294967295}
                                                        if false {
                                                            var v36 string = "\n\t"
                                                            v37 := GenS6{F0: 18446744073709551615}
                                                            _ = v36
                                                            _ = v37
                                                        }
                                                        var v38 uint64 = 42
                                                        _ = v35
                                                        _ = v38
                                                    }
                                                case -492:
                                                    if true {
                                                        v19 -= 0.1
                                                        if ("a" <= "\n\t") {
                                                            v39 := GenS7{F0: 65535}
                                                            v19 -= 1e10
                                                            _ = v39
                                                        }
                                                        if false {
                                                            var v40 uint = (v1 % (((v1 + v1)) | 1))
                                                            _ = v40
                                                        }
                                                    } else {
                                                        fmt.Println(v19)
                                                    }
                                                case 0:
                                                    fmt.Println(fi32, v3)
                                                }
                                                v7 += genFn0(-128, (990 <= 73))
                                            } else {
                                                v7 += v7
                                                var v41 uint16 = 255
                                                if false {
                                                    if false {
                                                        var v42 float64 = -2.25
                                                        v1 = 0
                                                        _ = v42
                                                    }
                                                    v43 := GenS4{F0: 1000}
                                                    switch "" {
                                                    case "":
                                                        if true {
                                                            v7 = 188
                                                        }
                                                    case "hello":
                                                        fmt.Println(fi32)
                                                    }
                                                    if (!false) {
                                                        var v44 int16 = -32768
                                                        var v45 uint8 = genFn0(-128, (639 != 0))
                                                        v46 := GenS8{F0: 2}
                                                        _ = v44
                                                        _ = v45
                                                        _ = v46
                                                    } else {
                                                        v47 := GenS3{F0: (-(fi32))}
                                                        v41 = 32767
                                                        _ = v47
                                                    }
                                                    _ = v43
                                                } else {
                                                    switch "a" {
                                                    case "\n\t":
                                                        switch (fi23 / (((fi23 / (((fi32 % (((fi23 >> ((4294967295) & 63))) | 1))) | 1))) | 1)) {
                                                        case 2147483647:
                                                            v7 = v7
                                                        case 32768:
                                                            fmt.Println(fi32, v41)
                                                        case 42:
                                                            fmt.Println(v7, v19, v19)
                                                        default:
                                                            v3 = v3
                                                        }
                                                    case "a":
                                                        v48 := GenS3{F0: (fi23 >> ((18446744073709551615) & 63))}
                                                        _ = v48
                                                    default:
                                                        _ = len(v4)
                                                    }
                                                    v49 := GenS9{F0: (-(fi32))}
                                                    var v50 int16 = 32767
                                                    _ = v49
                                                    _ = v50
                                                }
                                                _ = v41
                                            }
                                        default:
                                            v7 += 128
                                        }
                                        if false {
                                            var v51 int64 = -535
                                            var v52 uint8 = v7
                                            v4 = append(v4, (fi23 ^ (fi32 - (^(fi32)))))
                                            _ = v51
                                            _ = v52
                                        }
                                        switch "" {
                                        case "":
                                            v3 = (^(v3))
                                        case "foo bar":
                                            v4 = append(v4, fi23)
                                        case "hello":
                                            if true {
                                                fmt.Println(v7, fi32)
                                            }
                                        default:
                                            var v53 float64 = (1.5 / 1e10)
                                            _ = v53
                                        }
                                    } else {
                                        v1 = 998
                                        v4 = append(v4, fi23)
                                        v19 -= ((0.0 * float32(float64(v19))) * 3.14159)
                                    }
                                    if (true && (0 >= -322)) {
                                        v1 = v1
                                        v54 := GenS10{F0: 0, F1: 18446744073709551615, F2: -128}
                                        if false {
                                            switch (("" + "hello") + "\n\t") {
                                            case "a":
                                                switch "" {
                                                case "a":
                                                    var v55 bool = false
                                                    _ = v55
                                                default:
                                                    v56 := GenS3{F0: fi32}
                                                    _ = v56
                                                }
                                            case "hello":
                                                if true {
                                                    v7 = v7
                                                    var v57 uint64 = 879
                                                    _ = v57
                                                } else {
                                                    v3 = -314
                                                    if false {
                                                        v3 = 2147483647
                                                        v58 := GenS3{F0: fi32}
                                                        fmt.Println(v3, v19)
                                                        _ = v58
                                                    } else {
                                                        v59 := GenS10{F0: 111, F1: 2, F2: -128}
                                                        _ = v59
                                                    }
                                                }
                                            default:
                                                _ = len(v4)
                                            }
                                        } else {
                                            v1 = 507
                                            var v60 uint8 = 183
                                            var v61 uint8 = genFn0(-19, true)
                                            _ = v60
                                            _ = v61
                                        }
                                        switch fi23 {
                                        case 107:
                                            fmt.Println(v1)
                                        case 0:
                                            v3 = (v3 % (((v3 << (((^(v1))) & 63))) | 1))
                                        case 42:
                                            v62 := GenS5{F0: uint32(v1)}
                                            _ = v62
                                        default:
                                            v1 = (^(v1))
                                        }
                                        var v63 string = "a"
                                        _ = v54
                                        _ = v63
                                    }
                                } else {
                                    v3 = v3
                                    switch 760 {
                                    case -724:
                                        v2 = append(v2, "\n\t")
                                    case 453:
                                        fmt.Println(v3)
                                    default:
                                        switch "" {
                                        case "a":
                                            if len(v4) > 0 {
                                                v4[0] = 366
                                            }
                                        }
                                    }
                                }
                            } else {
                                var v64 int32 = (v3 << ((4294967295) & 63))
                                _ = v64
                            }
                            v7 = 127
                            var v65 uint64 = 2
                            _ = fi32
                            _ = v65
                        }
                    case "foo bar":
                        fmt.Println(v1)
                    }
                case "":
                    fmt.Println(v1, fi23, fi23)
                default:
                    v1 += v1
                }
                _ = fi23
                _ = v24
            }
            _ = v22
        }
    default:
        var v66 uint32 = 813
        _ = v66
    }
    switch "foo bar" {
    case "foo bar":
        for fi67 := 0; fi67 < 8; fi67++ {
            v19 -= v19
            _ = fi67
        }
    case "a":
        v68 := GenS11{F0: v7, F1: 60}
        _ = v68
    default:
        var v69 string = "\n\t"
        _ = v69
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
