package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 uint16 }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int32; F1 uint64; F2 uint64 }
type GenS5 struct { F0 string; F1 uint }
type GenS6 struct { F0 int8 }
func genFn0(p0 int64) string {
    var v4 bool = (18446744073709551615 > 2)
    v5 := GenS1{F0: 88, F1: 9223372036854775807, F2: 32767}
    if (-264 > -654) {
        var v6 uint16 = 2
        v4 = v4
        _ = v6
    } else {
        switch 42 {
        case 741:
            if ("" < "hello") {
                for fi7 := 0; fi7 < 5; fi7++ {
                    var v8 int8 = 42
                    v8 = v8
                    _ = fi7
                    _ = v8
                }
                for fi9 := 0; fi9 < 20; fi9++ {
                    var v10 int32 = 42
                    _ = fi9
                    _ = v10
                }
            } else {
                var v11 int32 = 42
                _ = v11
            }
        case -947:
            v4 = (-2.25 == 0.0)
        case -1:
            v12 := []int{}
            _ = v12
        }
        v13 := GenS1{F0: -8, F1: 655, F2: 65535}
        if true {
            var v14 uint16 = 633
            _ = v14
        }
        _ = v13
    }
    _ = p0
    _ = v4
    _ = v5
    return ("hello" + "foo bar")
}
func genFn1(p0 string, p1 string) int8 {
    if ((float32(((3.14159 + 0.1) * float64(3.14159))) * 1.5) < 3.14159) {
        for fi15 := 0; fi15 < 3; fi15++ {
            fmt.Println(fi15)
            _ = fi15
        }
        p0 = ((((((("\n\t" + p0) + p0) + p1) + "foo bar") + p1) + p1) + ("" + p0))
        p1 = "a"
    } else {
        switch -1 {
        case 4294967295:
            var v16 int32 = -2147483648
            _ = v16
        case 497:
            var v17 string = p0
            _ = v17
        default:
            fmt.Println(p0, p0)
        }
        var v18 uint8 = 244
        v19 := []int{1}
        _ = v18
        _ = v19
    }
    if (433 != 4294967295) {
        for fi20 := 0; fi20 < 16; fi20++ {
            fmt.Println(p0, p0)
            var v21 uint8 = 161
            var v22 uint16 = 733
            _ = fi20
            _ = v21
            _ = v22
        }
    }
    var v23 string = p1
    var v24 uint16 = 32767
    v25 := []int{42, 880}
    _ = p0
    _ = p1
    _ = v23
    _ = v24
    _ = v25
    return 80
}

func main() {
    v1 := []int{42, 84, 32768, -2147483648}
    var v2 uint = 547
    var v3 int = 42
    v2 -= v2
    if true {
        v3 = (v3 << (((^(v2))) & 63))
        v2 -= 443
        v26 := GenS1{F0: 42, F1: v3, F2: 255}
        _ = v26
    } else {
        switch "hello" {
        case "":
            v3 = v3
        case "\n\t":
            var v27 uint = 741
            _ = v27
        default:
            v3 -= v3
        }
        fmt.Println(v2, v3)
        fmt.Println(v3)
    }
    switch (v3 << ((v2) & 63)) {
    case 32767:
        var v28 float64 = 3.14159
        _ = v28
    default:
        if (2 > 42) {
            v29 := GenS1{F0: 0, F1: v3, F2: 42}
            _ = v29
        } else {
            for fi30 := 0; fi30 < 16; fi30++ {
                v31 := GenS1{F0: 36, F1: (fi30 * (fi30 / ((fi30) | 1))), F2: 0}
                var v32 int8 = -74
                _ = fi30
                _ = v31
                _ = v32
            }
            for fi33 := 0; fi33 < 18; fi33++ {
                v34 := GenS1{F0: int8(v3), F1: v3, F2: 386}
                _ = len(v1)
                v35 := GenS2{F0: 1, F1: "foo bar"}
                _ = fi33
                _ = v34
                _ = v35
            }
            v36 := GenS1{F0: -88, F1: (v3 | v3), F2: 472}
            _ = v36
        }
    }
    for fi37 := 0; fi37 < 6; fi37++ {
        var v38 float32 = 0.1
        v39 := GenS3{F0: 982}
        _ = fi37
        _ = v38
        _ = v39
    }
    var v40 uint32 = 661
    _ = len(v1)
    var v41 uint64 = 2
    switch (v3 / (((v3 & (v3 >> ((2) & 63)))) | 1)) {
    case -510:
        v41 -= v41
    }
    for fi42 := 0; fi42 < 9; fi42++ {
        switch v3 {
        case 128:
            var v43 float64 = -2.25
            _ = v43
        case -928:
            if ((^(v41)) == v41) {
                v2 = 1
            }
        }
        var v44 int64 = int64(v41)
        v41 -= 42
        _ = fi42
        _ = v44
    }
    var v45 float64 = -2.25
    if (true && true) {
        if (true || (!(!(!(true && false))))) {
            if len(v1) > 0 {
                v1[0] = 9223372036854775807
            }
            v46 := GenS2{F0: 1, F1: "a"}
            _ = v46
        }
    }
    if false {
        for fi47 := 0; fi47 < 8; fi47++ {
            v2 = 4294967295
            switch "\n\t" {
            case "hello":
                _ = len(v1)
            case "\n\t":
                for fi48 := 0; fi48 < 17; fi48++ {
                    if (v45 <= v45) {
                        v41 = 801
                        v3 = (fi48 << ((v2) & 63))
                    }
                    _ = fi48
                }
            case "foo bar":
                var v49 int64 = -9223372036854775808
                _ = v49
            default:
                v50 := GenS4{F0: -853, F1: (v41 | 1), F2: (^(v41))}
                _ = v50
            }
            _ = fi47
        }
        for fi51 := 0; fi51 < 18; fi51++ {
            v40 += v40
            v52 := GenS3{F0: 0}
            for fi53 := 0; fi53 < 19; fi53++ {
                v40 += (v40 << ((v2) & 63))
                v2 = 1
                switch ("hello" + ((("\n\t" + "hello") + "hello") + "a")) {
                case "\n\t":
                    switch "" {
                    case "":
                        switch ("\n\t" + "") {
                        case "foo bar":
                            var v54 float32 = -2.25
                            fmt.Println(v45, v3)
                            fmt.Println(v2, v54, v40)
                            _ = v54
                        case "a":
                            var v55 uint16 = 982
                            _ = v55
                        default:
                            v3 -= -198
                        }
                    default:
                        var v56 uint64 = (v41 % ((uint64(v2)) | 1))
                        _ = v56
                    }
                case "foo bar":
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                _ = fi53
            }
            _ = fi51
            _ = v52
        }
        var v57 int = 42
        _ = v57
    } else {
        var v58 uint8 = 180
        if ((255 == 255) && (42 != v40)) {
            var v59 uint8 = 32
            switch int(v40) {
            case 42:
                if ((v41 & v41) >= 42) {
                    v59 = v59
                    fmt.Println(v41, v2, v59)
                } else {
                    var v60 string = ("foo bar" + "")
                    switch "foo bar" {
                    case "":
                        v3 = v3
                    }
                    _ = v60
                }
            case 2147483647:
                switch "hello" {
                case "hello":
                    var v61 uint16 = uint16(v2)
                    _ = v61
                }
            default:
                for fi62 := 0; fi62 < 18; fi62++ {
                    var v63 int16 = 32767
                    fmt.Println(v2)
                    _ = fi62
                    _ = v63
                }
            }
            _ = len(v1)
            _ = v59
        }
        _ = v58
    }
    var v64 uint32 = 4294967295
    var v65 uint64 = (v41 ^ v41)
    if ((!true) || true) {
        if len(v1) > 0 {
            v1[0] = v3
        }
    } else {
        fmt.Println(v2)
        if len(v1) > 0 {
            v1[0] = 806
        }
        if false {
            v66 := GenS5{F0: ("" + (("\n\t" + "hello") + "hello")), F1: 0}
            _ = v66
        }
    }
    var v67 uint = v2
    if true {
        var v68 bool = (0 >= 78)
        if ("\n\t" >= "a") {
            var v69 int16 = -32768
            v40 = (v40 + v64)
            switch v3 {
            case -2147483648:
                var v70 uint64 = 42
                _ = v70
            }
            _ = v69
        }
        for fi71 := 0; fi71 < 2; fi71++ {
            switch v3 {
            case 127:
                for fi72 := 0; fi72 < 8; fi72++ {
                    switch 1 {
                    case 128:
                        if (true || v68) {
                            v73 := GenS5{F0: ((("hello" + ("foo bar" + "a")) + "foo bar") + "\n\t"), F1: v67}
                            if v68 {
                                var v74 uint8 = 1
                                _ = v74
                            }
                            _ = v73
                        } else {
                            fmt.Println(v40)
                            switch ("a" + "") {
                            case "\n\t":
                                switch (fi72 & (v3 / ((int(v64)) | 1))) {
                                case -263:
                                    fmt.Println(v45)
                                case 32768:
                                    v75 := GenS6{F0: 127}
                                    _ = v75
                                case -908:
                                    v40 -= v64
                                default:
                                    v41 = v41
                                }
                            case "hello":
                                var v76 uint32 = v40
                                _ = v76
                            default:
                                v2 += v2
                            }
                        }
                    case 32768:
                        v2 -= v2
                        _ = len(v1)
                        fmt.Println(fi71, fi72, v40)
                    }
                    var v77 float32 = 0.0
                    v68 = false
                    _ = fi72
                    _ = v77
                }
            default:
                if v68 {
                    var v78 int64 = 42
                    switch "" {
                    case "":
                        for fi79 := 0; fi79 < 12; fi79++ {
                            if len(v1) > 0 {
                                v1[0] = (fi71 << ((v67) & 63))
                            }
                            v67 += 18446744073709551615
                            _ = fi79
                        }
                    case "\n\t":
                        for fi80 := 0; fi80 < 8; fi80++ {
                            switch "" {
                            case "hello":
                                v81 := GenS3{F0: 0}
                                _ = v81
                            case "":
                                if true {
                                    fmt.Println(v3, v67)
                                    v1 = append(v1, (fi80 * (fi80 * (^(fi71)))))
                                } else {
                                    v67 = v67
                                    if len(v1) > 0 {
                                        v1[0] = fi71
                                    }
                                }
                            case "a":
                                fmt.Println(v2)
                            default:
                                var v82 int32 = -1
                                _ = v82
                            }
                            v83 := GenS1{F0: 53, F1: 2147483647, F2: 2}
                            v41 = 18446744073709551615
                            _ = fi80
                            _ = v83
                        }
                    }
                    _ = v78
                }
            }
            _ = fi71
        }
        _ = v68
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v45:", v45)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v40
    _ = v41
    _ = v45
    _ = v64
    _ = v65
    _ = v67
}
