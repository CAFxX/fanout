package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32; F1 int16; F2 uint8 }
type GenS4 struct { F0 uint32; F1 uint16 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint16; F1 int16 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 bool; F1 int }
type GenS9 struct { F0 bool; F1 bool; F2 uint32 }
type GenS10 struct { F0 int16; F1 uint }
func genFn0(p0 int8, p1 uint16) uint8 {
    var v5 int = -2147483648
    for fi6 := 0; fi6 < 16; fi6++ {
        switch "\n\t" {
        case "hello":
            v5 = -788
        case "foo bar":
            v7 := GenS1{F0: 285, F1: 42}
            _ = v7
        case "\n\t":
            v8 := GenS1{F0: int16(p1), F1: 0}
            _ = v8
        default:
            for fi9 := 0; fi9 < 3; fi9++ {
                v10 := []int{(fi6 / ((fi6) | 1)), fi6, fi9, (v5 % ((128) | 1))}
                v11 := GenS2{F0: 2}
                if (0 == 303) {
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                    p0 = -128
                }
                _ = fi9
                _ = v10
                _ = v11
            }
        }
        fmt.Println(fi6)
        _ = fi6
    }
    v12 := []int{v5}
    v13 := GenS1{F0: 1, F1: 42}
    v14 := GenS2{F0: 0}
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    _ = v14
    _ = v5
    return 128
}

func main() {
    var v1 uint = 687
    v2 := []string{"hello"}
    var v3 string = (("hello" + (("hello" + "") + "a")) + (("\n\t" + "hello") + ("hello" + (("" + ("a" + "")) + "hello"))))
    var v4 float64 = (float64(((float32(3.14159) + (1e10 - ((1e10 / (1.5 * 1e10)) * ((1e10 + float32(0.1)) - 3.14159)))) / (1e10 * -2.25))) / 1.0)
    if false {
        for fi15 := 0; fi15 < 16; fi15++ {
            for fi16 := 0; fi16 < 2; fi16++ {
                v2 = append(v2, v3)
                fmt.Println(v1, v4, fi15)
                v17 := GenS1{F0: -198, F1: 128}
                _ = fi16
                _ = v17
            }
            if len(v2) > 0 {
                v2[0] = v3
            }
            for fi18 := 0; fi18 < 12; fi18++ {
                v19 := GenS2{F0: 221}
                var v20 float32 = -2.25
                if ((!(float64(0.1) != float64(v20))) || false) {
                    if (true || true) {
                        v3 = v3
                        if len(v2) > 0 {
                            v2[0] = v3
                        }
                        var v21 uint64 = 906
                        v22 := GenS3{F0: 65535, F1: 1, F2: 60}
                        v23 := GenS4{F0: 2, F1: 65535}
                        _ = v21
                        _ = v22
                        _ = v23
                    } else {
                        v2 = append(v2, v3)
                        v1 = 18446744073709551615
                        v3 = "a"
                    }
                } else {
                    var v24 int32 = -1
                    _ = v24
                }
                _ = fi18
                _ = v19
                _ = v20
            }
            _ = fi15
        }
    } else {
        for fi25 := 0; fi25 < 4; fi25++ {
            v26 := GenS3{F0: 71, F1: 42, F2: 95}
            _ = fi25
            _ = v26
        }
        v4 -= -2.25
    }
    var v27 uint64 = 785
    v28 := GenS4{F0: 2147483648, F1: 32767}
    for fi29 := 0; fi29 < 18; fi29++ {
        v4 = v4
        v4 += 0.0
        _ = fi29
    }
    if (!(-1 <= -1)) {
        for fi30 := 0; fi30 < 1; fi30++ {
            if false {
                switch "\n\t" {
                case "a":
                    if true {
                        if true {
                            var v31 bool = false
                            _ = v31
                        }
                    } else {
                        v32 := GenS1{F0: 0, F1: 128}
                        _ = v32
                    }
                }
            } else {
                v4 = ((float64(0.1) * v4) - ((float64(1e10) * float64((float32((float64(0.1) + v4)) + 1e10))) / v4))
            }
            var v33 int8 = 100
            _ = fi30
            _ = v33
        }
        switch "\n\t" {
        case "a":
            v4 = v4
        default:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
    } else {
        for fi34 := 0; fi34 < 20; fi34++ {
            for fi35 := 0; fi35 < 15; fi35++ {
                switch -590 {
                case -372:
                    switch (v3 + v3) {
                    case "foo bar":
                        switch (fi35 - (^(fi34))) {
                        case 32767:
                            fmt.Println(fi34)
                        default:
                            if (!(((fi35 + fi34) <= fi34) && false)) {
                                v2 = append(v2, "hello")
                            }
                        }
                    case "a":
                        var v36 bool = false
                        _ = v36
                    case "\n\t":
                        v4 = v4
                    default:
                        fmt.Println(v27, fi35)
                    }
                }
                _ = fi35
            }
            _ = fi34
        }
    }
    v3 = ("\n\t" + v3)
    if true {
        v37 := GenS5{F0: 2147483647}
        for fi38 := 0; fi38 < 11; fi38++ {
            var v39 uint = 1
            v1 = (^(v1))
            _ = fi38
            _ = v39
        }
        _ = v37
    } else {
        v1 += v1
        if (2 == 1) {
            var v40 uint64 = (v27 >> ((78) & 63))
            _ = v40
        }
        for fi41 := 0; fi41 < 14; fi41++ {
            v4 += -2.25
            v3 = v3
            switch (v3 + v3) {
            case "a":
                var v42 int = -517
                _ = v42
            case "foo bar":
                var v43 uint8 = genFn0(0, 32767)
                _ = v43
            case "hello":
                switch fi41 {
                case 42:
                    v44 := GenS5{F0: (fi41 ^ fi41)}
                    _ = v44
                case 32767:
                    if (!(true || false)) {
                        var v45 float64 = v4
                        _ = v45
                    }
                default:
                    for fi46 := 0; fi46 < 19; fi46++ {
                        switch v3 {
                        case "":
                            if (43 < 42) {
                                v47 := GenS6{F0: 65535, F1: -888}
                                _ = v47
                            } else {
                                var v48 string = v3
                                _ = v48
                            }
                        case "a":
                            switch fi46 {
                            case -851:
                                v27 -= v27
                            case 127:
                                v27 = v27
                            }
                        case "hello":
                            v1 -= v1
                        default:
                            v49 := GenS7{F0: 1}
                            _ = v49
                        }
                        _ = fi46
                    }
                }
            default:
                if (v1 > 0) {
                    v1 -= v1
                }
            }
            _ = fi41
        }
    }
    var v50 int8 = 42
    fmt.Println(v50)
    if false {
        if ((v1 ^ 255) >= v1) {
            var v51 uint8 = 228
            _ = v51
        } else {
            switch -1 {
            case 0:
                v27 = v27
            }
            var v52 int64 = 233
            _ = v52
        }
    }
    var v53 int64 = 9223372036854775807
    if (!(1 < 574)) {
        if true {
            var v54 int8 = (-(v50))
            _ = v54
        } else {
            if false {
                if true {
                    fmt.Println(v3, v27, v27)
                    switch 127 {
                    case -1:
                        for fi55 := 0; fi55 < 6; fi55++ {
                            v56 := GenS7{F0: -1}
                            fmt.Println(v1, v4)
                            fmt.Println(v50)
                            _ = fi55
                            _ = v56
                        }
                    case 4294967295:
                        if ((3.14159 >= 0.0) || true) {
                            v1 = v1
                        } else {
                            v50 += int8(v27)
                        }
                    case 9223372036854775807:
                        v57 := GenS4{F0: 1, F1: 42}
                        _ = v57
                    }
                }
            } else {
                v1 = v1
            }
            var v58 float64 = (float64(-2.25) + (((v4 * (0.0 - v4)) / 3.14159) - (v4 * (1e10 + 3.14159))))
            _ = v58
        }
        switch 238 {
        case -2147483648:
            switch 2147483647 {
            case 0:
                v59 := GenS8{F0: false, F1: -444}
                _ = v59
            default:
                v53 -= -976
            }
        }
        v60 := GenS6{F0: 135, F1: 1}
        _ = v60
    } else {
        v27 = v27
    }
    if true {
        var v61 int16 = 42
        _ = v61
    } else {
        if true {
            if ((v1 >> ((v1) & 63)) < v1) {
                fmt.Println(v1, v4, v53)
                _ = len(v2)
            } else {
                var v62 int = -991
                fmt.Println(v50, v53)
                var v63 uint = v1
                _ = v62
                _ = v63
            }
        }
        if true {
            var v64 bool = true
            _ = v64
        } else {
            v65 := GenS9{F0: false, F1: ((v50 - v50) >= (^(v50))), F2: 2147483647}
            _ = v65
        }
        if (!(!(((v1 != (v1 >> ((v1) & 63))) || false) && (!false)))) {
            var v66 float64 = float64((3.14159 * -2.25))
            v66 = (-2.25 * v66)
            _ = v66
        } else {
            v67 := GenS10{F0: 0, F1: (v1 & (v1 | 42))}
            var v68 float64 = v4
            _ = v67
            _ = v68
        }
    }
    var v69 int8 = v50
    if false {
        if false {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        var v70 int16 = 663
        switch "foo bar" {
        case "\n\t":
            if len(v2) > 0 {
                v2[0] = v3
            }
        case "":
            var v71 int16 = 32767
            _ = v71
        case "a":
            var v72 uint8 = 7
            _ = v72
        }
        _ = v70
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v50:", v50)
    fmt.Println("v53:", v53)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v50
    _ = v53
    _ = v69
}
