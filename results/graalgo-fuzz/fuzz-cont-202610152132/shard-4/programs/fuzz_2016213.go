package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 bool; F1 int; F2 int16 }
type GenS4 struct { F0 uint; F1 uint }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int8; F1 uint32 }
func genFn0(p0 int32) int64 {
    var v7 uint8 = 175
    v7 = 205
    v8 := GenS1{F0: false, F1: 4294967295}
    _ = p0
    _ = v7
    _ = v8
    return 354
}

func main() {
    var v1 int32 = 2147483647
    v2 := []int{-2147483648}
    var v3 bool = (2 == 265)
    v4 := make(map[string]int)
    v5 := make(map[uint32]string)
    var v6 uint = 2
    var v9 int64 = 576
    for fi10 := 0; fi10 < 13; fi10++ {
        v11 := GenS1{F0: (817 != 495), F1: 2}
        _ = fi10
        _ = v11
    }
    switch ("a" + "\n\t") {
    case "a":
        v12 := GenS1{F0: v3, F1: 1}
        _ = v12
    default:
        if (!(false && v3)) {
            fmt.Println(v6, v1, v6)
            if v3 {
                for fi13 := 0; fi13 < 2; fi13++ {
                    v3 = (42 > 1000)
                    var v14 uint16 = 42
                    _ = len(v2)
                    _ = fi13
                    _ = v14
                }
                for fi15 := 0; fi15 < 20; fi15++ {
                    for fi16 := 0; fi16 < 18; fi16++ {
                        var v17 uint = v6
                        var v18 int32 = v1
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                    v19 := GenS1{F0: v3, F1: 893}
                    _ = fi15
                    _ = v19
                }
                switch "\n\t" {
                case "hello":
                    switch "hello" {
                    case "hello":
                        v20 := v4["hello"]
                        _ = v20
                    default:
                        var v21 int32 = v1
                        _ = v21
                    }
                }
            }
            for fi22 := 0; fi22 < 11; fi22++ {
                switch (fi22 | int(v1)) {
                case 32768:
                    var v23 uint16 = 32767
                    _ = v23
                default:
                    v24 := GenS1{F0: true, F1: 502}
                    _ = v24
                }
                v5[0] = ("\n\t" + ("foo bar" + (("" + "hello") + "hello")))
                _ = fi22
            }
        } else {
            v1 += v1
            if false {
                v4["hello"] = 963
            }
            if ((^(v6)) > 0) {
                for fi25 := 0; fi25 < 17; fi25++ {
                    v26 := GenS1{F0: (100 <= -55), F1: 0}
                    switch (fi25 / ((fi25) | 1)) {
                    case 42:
                        v3 = v3
                    case 9223372036854775807:
                        var v27 string = "a"
                        _ = v27
                    case 128:
                        var v28 int16 = int16(fi25)
                        _ = v28
                    default:
                        switch "a" {
                        case "\n\t":
                            var v29 uint = (v6 / ((97) | 1))
                            _ = v29
                        default:
                            v6 = v6
                        }
                    }
                    _ = fi25
                    _ = v26
                }
                if v3 {
                    v9 = v9
                    v1 += -1
                }
                for fi30 := 0; fi30 < 8; fi30++ {
                    if (("\n\t" + (("a" + "foo bar") + "")) >= ((("foo bar" + "") + ("foo bar" + "")) + (("hello" + "a") + ""))) {
                        var v31 int64 = 9223372036854775807
                        for fi32 := 0; fi32 < 8; fi32++ {
                            v33 := GenS1{F0: true, F1: 104}
                            var v34 string = (("\n\t" + "\n\t") + "foo bar")
                            _ = fi32
                            _ = v33
                            _ = v34
                        }
                        _ = v31
                    } else {
                        fmt.Println(v9, v3, v9)
                        if v3 {
                            var v35 bool = (v9 > (v9 << ((65535) & 63)))
                            for fi36 := 0; fi36 < 15; fi36++ {
                                if (fi30 <= 9223372036854775807) {
                                    v37 := GenS2{F0: v1, F1: 248}
                                    v38 := GenS2{F0: (v1 >> ((36) & 63)), F1: 798}
                                    fmt.Println(v35)
                                    _ = v37
                                    _ = v38
                                }
                                switch fi30 {
                                case 486:
                                    fmt.Println(v3, v3, v9)
                                case 127:
                                    if (-128 <= 16) {
                                        v39 := GenS3{F0: v35, F1: (fi30 + (fi36 * (fi36 ^ (fi36 >> ((v6) & 63))))), F2: -1}
                                        _ = v39
                                    } else {
                                        v40 := GenS1{F0: true, F1: 1}
                                        _ = v40
                                    }
                                }
                                v2 = append(v2, int(v1))
                                _ = fi36
                            }
                            _ = v35
                        } else {
                            v41 := GenS4{F0: (^(v6)), F1: 60}
                            v42 := GenS5{F0: 311}
                            _ = v41
                            _ = v42
                        }
                    }
                    _ = fi30
                }
            }
        }
    }
    v9 -= (v9 << ((0) & 63))
    var v43 int = -645
    if (32767 < -32768) {
        for fi44 := 0; fi44 < 16; fi44++ {
            v2 = append(v2, v43)
            var v45 bool = v3
            v46 := GenS5{F0: 1}
            _ = fi44
            _ = v45
            _ = v46
        }
        if false {
            switch (v43 >> ((v6) & 63)) {
            case -29:
                v43 = 176
            case -107:
                var v47 uint16 = 0
                _ = v47
            default:
                v2 = append(v2, (v43 % (((^(v43))) | 1)))
            }
        }
        fmt.Println(v43, v1, v3)
    } else {
        if v3 {
            if v3 {
                v9 += (v9 | v9)
                v48 := GenS1{F0: v3, F1: 0}
                _ = v48
            } else {
                _ = len(v4)
                v3 = (v6 <= (v6 & (v6 ^ (v6 / ((42) | 1)))))
                switch v43 {
                case 32768:
                    switch ("" + "\n\t") {
                    case "a":
                        v4[""] = (v43 / ((-119) | 1))
                    case "\n\t":
                        v49 := GenS6{F0: 65535}
                        _ = v49
                    }
                case 2147483647:
                    if v3 {
                        v43 += -1
                    } else {
                        v3 = true
                        switch "a" {
                        case "foo bar":
                            var v50 uint16 = 65535
                            _ = v50
                        }
                        switch -516 {
                        case 4294967295:
                            v6 -= v6
                        case 1:
                            v51 := GenS2{F0: (v1 / ((int32(v6)) | 1)), F1: 4294967295}
                            _ = v51
                        case -9223372036854775808:
                            if v3 {
                                var v52 int64 = v9
                                v1 = -2147483648
                                _ = v52
                            } else {
                                v6 = 42
                                v53 := GenS3{F0: (42 > 2), F1: (v43 | v43), F2: 1}
                                _ = v53
                            }
                        default:
                            delete(v5, 0)
                        }
                    }
                default:
                    var v54 bool = true
                    _ = v54
                }
            }
            v55 := GenS1{F0: v3, F1: 4294967295}
            _ = v55
        } else {
            var v56 float32 = 3.14159
            var v57 uint64 = 0
            _ = v56
            _ = v57
        }
        for fi58 := 0; fi58 < 19; fi58++ {
            v59 := GenS6{F0: 2}
            v43 += (v43 << ((0) & 63))
            _ = fi58
            _ = v59
        }
        var v60 int8 = 1
        _ = v60
    }
    v3 = (0 >= 2147483647)
    switch ("" + "hello") {
    case "a":
        v43 -= (v43 | v43)
    case "foo bar":
        if true {
            var v61 uint64 = 18446744073709551615
            if v3 {
                v62 := GenS4{F0: (v6 - 18446744073709551615), F1: (v6 + 65535)}
                _ = v62
            }
            if false {
                v63 := GenS5{F0: uint32(v6)}
                _ = v63
            } else {
                fmt.Println(v3, v9)
                for fi64 := 0; fi64 < 11; fi64++ {
                    for fi65 := 0; fi65 < 5; fi65++ {
                        v66 := GenS1{F0: v3, F1: 4294967295}
                        v1 = (v1 | (v1 + v1))
                        v6 = v6
                        _ = fi65
                        _ = v66
                    }
                    _ = len(v4)
                    v2 = append(v2, v43)
                    _ = fi64
                }
            }
            _ = v61
        } else {
            switch (((("" + "foo bar") + "hello") + ("hello" + "hello")) + "a") {
            case "":
                fmt.Println(v43, v6, v3)
            case "a":
                v6 = (v6 >> ((v6) & 63))
            }
            var v67 uint32 = 2
            _ = v67
        }
    case "":
        var v68 bool = true
        _ = v68
    default:
        if v3 {
            switch "\n\t" {
            case "hello":
                v3 = (!v3)
            default:
                switch "hello" {
                case "hello":
                    v3 = v3
                case "a":
                    v69 := GenS7{F0: genFn0((v1 >> ((v6) & 63)))}
                    _ = v69
                case "":
                    switch (v43 + v43) {
                    case 4294967295:
                        v1 = (v1 % (((v1 & v1)) | 1))
                    case 32767:
                        v9 += int64(v6)
                    case 0:
                        var v70 bool = (v3 && v3)
                        _ = v70
                    default:
                        for fi71 := 0; fi71 < 12; fi71++ {
                            if (((3.14159 <= 0.1) || (fi71 != 9223372036854775807)) || true) {
                                var v72 int8 = 72
                                for fi73 := 0; fi73 < 16; fi73++ {
                                    switch "hello" {
                                    case "":
                                        var v74 int16 = 1000
                                        _ = v74
                                    }
                                    v75 := GenS6{F0: 2147483647}
                                    delete(v4, "\n\t")
                                    switch 0 {
                                    case 42:
                                        v9 = (v9 | v9)
                                    }
                                    _ = fi73
                                    _ = v75
                                }
                                _ = v72
                            }
                            v1 = (-(v1))
                            _ = fi71
                        }
                    }
                }
            }
            if false {
                v76 := GenS2{F0: v1, F1: 932}
                _ = v76
            } else {
                v77 := GenS3{F0: v3, F1: 0, F2: 746}
                for fi78 := 0; fi78 < 15; fi78++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    fmt.Println(v3, v43, v1)
                    v3 = v3
                    _ = fi78
                }
                _ = v77
            }
            for fi79 := 0; fi79 < 12; fi79++ {
                switch ("hello" + "\n\t") {
                case "hello":
                    var v80 uint16 = 42
                    _ = v80
                case "\n\t":
                    _ = len(v5)
                default:
                    v43 = fi79
                }
                for fi81 := 0; fi81 < 5; fi81++ {
                    if v3 {
                        v82 := GenS5{F0: 2}
                        _ = v82
                    }
                    v83 := GenS8{F0: 42, F1: 2147483647}
                    fmt.Println(v43, v6, fi79)
                    _ = fi81
                    _ = v83
                }
                _ = fi79
            }
        } else {
            var v84 float32 = (0.1 * ((1.5 / float32(1.5)) / 1.0))
            _ = v84
        }
    }
    v85 := GenS5{F0: 42}
    var v86 float64 = 1e10
    var v87 int16 = 1000
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v86:", v86)
    fmt.Println("v87:", v87)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
    _ = v85
    _ = v86
    _ = v87
    _ = v9
}
