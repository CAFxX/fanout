package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 bool; F1 uint32; F2 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int8; F1 int32 }
type GenS7 struct { F0 bool }
func genFn0(p0 uint8, p1 uint16) string {
    if false {
        if (!(p0 == p0)) {
            fmt.Println(p1)
            if (true && false) {
                if (1 == 77) {
                    for fi4 := 0; fi4 < 15; fi4++ {
                        var v5 uint64 = uint64(p0)
                        _ = fi4
                        _ = v5
                    }
                }
            }
            p1 = 338
        }
        v6 := GenS1{F0: 4294967295, F1: (p0 >> ((18446744073709551615) & 63))}
        _ = v6
    }
    var v7 uint8 = p0
    _ = p0
    _ = p1
    _ = v7
    return ("hello" + "a")
}
func genFn1(p0 uint16, p1 uint32) int {
    for fi8 := 0; fi8 < 20; fi8++ {
        v9 := []int{(fi8 >> ((uint(p0)) & 63))}
        _ = fi8
        _ = v9
    }
    p0 = p0
    _ = p0
    _ = p1
    return -570
}

func main() {
    v1 := []uint32{1, 2147483647, 2147483647, 4294967295}
    var v2 bool = false
    var v3 uint8 = 2
    var v10 uint64 = 2
    var v11 string = ("\n\t" + ("hello" + ""))
    switch -176 {
    case 32767:
        var v12 uint8 = v3
        _ = v12
    case 803:
        v13 := GenS1{F0: 32768, F1: v3}
        _ = v13
    }
    for fi14 := 0; fi14 < 16; fi14++ {
        var v15 int32 = 2147483647
        v15 -= -882
        if len(v1) > 0 {
            v1[0] = 2147483647
        }
        _ = fi14
        _ = v15
    }
    v16 := GenS2{F0: v2, F1: 2147483647, F2: 28}
    var v17 string = v11
    switch (v11 + v17) {
    case "foo bar":
        for fi18 := 0; fi18 < 18; fi18++ {
            fmt.Println(v17)
            v19 := GenS2{F0: (v2 && v2), F1: 831, F2: 32767}
            _ = fi18
            _ = v19
        }
    default:
        _ = len(v1)
    }
    var v20 uint8 = 151
    switch 2147483647 {
    case 864:
        if false {
            v3 = v3
            var v21 uint16 = 0
            for fi22 := 0; fi22 < 16; fi22++ {
                var v23 int = (fi22 | 32767)
                _ = fi22
                _ = v23
            }
            _ = v21
        } else {
            for fi24 := 0; fi24 < 20; fi24++ {
                v17 = v11
                for fi25 := 0; fi25 < 5; fi25++ {
                    var v26 float64 = (1.5 / ((3.14159 / 3.14159) / 3.14159))
                    if len(v1) > 0 {
                        v1[0] = 129
                    }
                    _ = len(v1)
                    _ = fi25
                    _ = v26
                }
                _ = fi24
            }
            switch "hello" {
            case "hello":
                v17 = (v11 + v17)
            case "":
                var v27 bool = v2
                _ = v27
            default:
                switch genFn1(42, 1) {
                case 2147483647:
                    if (-926 >= 2147483647) {
                        var v28 uint8 = v3
                        v29 := GenS3{F0: -1}
                        _ = v28
                        _ = v29
                    } else {
                        var v30 int64 = -1
                        _ = v30
                    }
                case 23:
                    _ = len(v1)
                case 42:
                    switch genFn1(2, 2147483648) {
                    case 0:
                        for fi31 := 0; fi31 < 8; fi31++ {
                            fmt.Println(v10, fi31, v10)
                            var v32 float32 = 0.0
                            fmt.Println(v3)
                            _ = fi31
                            _ = v32
                        }
                    }
                }
            }
            var v33 uint16 = 255
            _ = v33
        }
    }
    for fi34 := 0; fi34 < 14; fi34++ {
        v11 = v11
        for fi35 := 0; fi35 < 18; fi35++ {
            var v36 int = genFn1(174, 2147483647)
            if v2 {
                var v37 uint32 = 65535
                v38 := GenS2{F0: (float32((0.0 - (1.5 * 0.1))) <= 0.1), F1: (v37 ^ v37), F2: 0}
                _ = v37
                _ = v38
            } else {
                if (-544 == 42) {
                    v1 = append(v1, 518)
                    _ = len(v1)
                    var v39 string = "foo bar"
                    _ = v39
                } else {
                    if (1 > 9223372036854775807) {
                        v2 = (-128 > 100)
                        v36 -= v36
                        v40 := GenS4{F0: v10}
                        if len(v1) > 0 {
                            v1[0] = 0
                        }
                        var v41 int16 = -676
                        _ = v40
                        _ = v41
                    } else {
                        v1 = append(v1, 0)
                    }
                    var v42 float64 = 1.5
                    _ = v42
                }
                switch (v11 + (v11 + v17)) {
                case "\n\t":
                    v11 = ("" + "foo bar")
                case "":
                    v43 := GenS5{F0: -394}
                    _ = v43
                default:
                    fmt.Println(v20)
                }
            }
            v17 = (("hello" + "hello") + v11)
            _ = fi35
            _ = v36
        }
        v44 := GenS1{F0: (fi34 % (((fi34 & -2147483648)) | 1)), F1: v20}
        _ = fi34
        _ = v44
    }
    for fi45 := 0; fi45 < 8; fi45++ {
        if v2 {
            var v46 uint = 18446744073709551615
            v3 += 42
            _ = v46
        } else {
            var v47 uint = 255
            _ = v47
        }
        v17 = v11
        _ = fi45
    }
    for fi48 := 0; fi48 < 10; fi48++ {
        var v49 int32 = -1
        _ = fi48
        _ = v49
    }
    for fi50 := 0; fi50 < 13; fi50++ {
        var v51 bool = v2
        var v52 uint32 = 0
        _ = fi50
        _ = v51
        _ = v52
    }
    var v53 int64 = -455
    v10 -= (v10 % ((v10) | 1))
    v54 := GenS3{F0: 42}
    if true {
        switch 2147483647 {
        case -282:
            v1 = append(v1, 2147483647)
        case 42:
            var v55 int32 = 42
            _ = v55
        default:
            if len(v1) > 0 {
                v1[0] = 2147483648
            }
        }
        switch -2147483648 {
        case -1:
            if (1000 != 42) {
                v20 = (v3 ^ (^(v20)))
                var v56 uint64 = (v10 - (^(v10)))
                _ = v56
            } else {
                v17 = "foo bar"
            }
        case 281:
            var v57 int = 32767
            _ = v57
        }
        for fi58 := 0; fi58 < 12; fi58++ {
            v20 += 67
            var v59 float64 = 0.0
            _ = fi58
            _ = v59
        }
    } else {
        switch "foo bar" {
        case "\n\t":
            for fi60 := 0; fi60 < 10; fi60++ {
                for fi61 := 0; fi61 < 12; fi61++ {
                    if (v10 != v10) {
                        if (42 != -1) {
                            v62 := GenS3{F0: -752}
                            var v63 bool = (!(v2 && (!(v53 > 0))))
                            _ = len(v1)
                            _ = v62
                            _ = v63
                        }
                        var v64 uint = 255
                        switch v17 {
                        case "a":
                            if v2 {
                                var v65 bool = (((v53 == -90) && (v11 < (v11 + ("a" + v11)))) && (v2 && v2))
                                _ = v65
                            } else {
                                v66 := GenS2{F0: (v2 && (v64 >= (^(v64)))), F1: 2147483647, F2: 32767}
                                v64 = (^(v64))
                                _ = v66
                            }
                        }
                        _ = v64
                    } else {
                        v10 += 18446744073709551615
                        fmt.Println(v2, v3, v17)
                        var v67 bool = (255 <= 253)
                        _ = v67
                    }
                    _ = fi61
                }
                if (!false) {
                    switch fi60 {
                    case 9223372036854775807:
                        v1 = append(v1, 42)
                    case 127:
                        fmt.Println(v53, v11, fi60)
                    case 999:
                        v10 = 1
                    }
                    v2 = ((fi60 >> ((65535) & 63)) != fi60)
                    switch (-(fi60)) {
                    case 978:
                        switch ((v11 + v17) + v11) {
                        case "foo bar":
                            switch (v11 + v11) {
                            case "foo bar":
                                switch (^(fi60)) {
                                case 32767:
                                    v68 := GenS2{F0: (v2 || v2), F1: 468, F2: 42}
                                    _ = v68
                                case -890:
                                    v3 -= v20
                                case -146:
                                    var v69 uint = 1
                                    _ = v69
                                default:
                                    var v70 uint16 = 0
                                    _ = v70
                                }
                            case "hello":
                                if v2 {
                                    if v2 {
                                        fmt.Println(v3, v10, v53)
                                        if v2 {
                                            var v71 int = fi60
                                            _ = v71
                                        }
                                    }
                                }
                            }
                        }
                    case 2147483647:
                        var v72 uint64 = (v10 * v10)
                        _ = v72
                    case 42:
                        switch genFn1(2, uint32(v20)) {
                        case 32767:
                            fmt.Println(v11, v3)
                        default:
                            if v2 {
                                v73 := GenS6{F0: -128, F1: -835}
                                v17 = ""
                                var v74 uint16 = 2
                                _ = v73
                                _ = v74
                            }
                        }
                    }
                } else {
                    if v2 {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    } else {
                        if len(v1) > 0 {
                            v1[0] = 200
                        }
                    }
                    if (!((v2 && (32767 < 32767)) || (!v2))) {
                        v75 := GenS3{F0: int16(v53)}
                        fmt.Println(fi60, v17, v10)
                        _ = v75
                    }
                    var v76 float32 = (1.5 / float32((1e10 + 0.1)))
                    _ = v76
                }
                v77 := GenS4{F0: (v10 | (v10 * 0))}
                _ = fi60
                _ = v77
            }
        case "hello":
            for fi78 := 0; fi78 < 7; fi78++ {
                var v79 int16 = -32768
                switch v17 {
                case "":
                    for fi80 := 0; fi80 < 4; fi80++ {
                        if len(v1) > 0 {
                            v1[0] = 4294967295
                        }
                        if ((fi80 << ((960) & 63)) != (-(fi78))) {
                            v11 = v11
                        } else {
                            var v81 uint32 = 49
                            _ = v81
                        }
                        if ((127 < 36) && (2 > 42)) {
                            switch ("foo bar" + ("hello" + v17)) {
                            case "":
                                v20 -= 128
                            default:
                                var v82 int = int(v3)
                                _ = v82
                            }
                            if (180 > 616) {
                                v53 = 9223372036854775807
                            } else {
                                v20 = v3
                                v1 = append(v1, 672)
                            }
                            v83 := GenS2{F0: v2, F1: 65535, F2: 255}
                            _ = v83
                        }
                        _ = fi80
                    }
                case "foo bar":
                    v79 += (v79 << ((18446744073709551615) & 63))
                default:
                    var v84 int8 = -1
                    _ = v84
                }
                _ = fi78
                _ = v79
            }
        case "":
            if (-1 != 851) {
                var v85 uint16 = 334
                var v86 int16 = 490
                _ = v85
                _ = v86
            } else {
                v87 := GenS7{F0: ((v53 - (v53 + v53)) < 67)}
                var v88 uint8 = (v3 << ((318) & 63))
                v89 := GenS1{F0: -647, F1: 0}
                _ = v87
                _ = v88
                _ = v89
            }
        default:
            var v90 string = ""
            _ = v90
        }
        switch -1 {
        case 4294967295:
            v53 += v53
        }
        for fi91 := 0; fi91 < 7; fi91++ {
            v17 = ((("hello" + (("" + v17) + "hello")) + "\n\t") + "foo bar")
            switch v11 {
            case "hello":
                fmt.Println(v11)
            case "":
                v92 := GenS1{F0: fi91, F1: uint8(v10)}
                _ = v92
            }
            _ = fi91
        }
    }
    fmt.Println(v2, v11, v17)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v53
    _ = v54
}
