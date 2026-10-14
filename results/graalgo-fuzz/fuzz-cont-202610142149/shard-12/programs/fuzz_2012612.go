package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 string; F1 uint; F2 int32 }
type GenS4 struct { F0 string; F1 uint32; F2 uint }
type GenS5 struct { F0 int64; F1 bool }
type GenS6 struct { F0 int16; F1 uint16; F2 uint64 }

func main() {
    var v1 float32 = 1e10
    var v2 string = (("foo bar" + "a") + ("hello" + "foo bar"))
    var v3 uint16 = 42
    var v4 int = -9223372036854775808
    v5 := make(map[uint32]string)
    v6 := make(map[uint32]uint32)
    var v7 uint16 = v3
    var v8 uint = 65535
    var v9 uint = (^(v8))
    delete(v6, uint32(v7))
    fmt.Println(v9)
    if true {
        var v10 int8 = 107
        _ = v10
    }
    for fi11 := 0; fi11 < 13; fi11++ {
        v12 := []int{fi11, 42, fi11, int(v9)}
        if (127 >= 2) {
            if (42 <= 32767) {
                switch (v4 % ((fi11) | 1)) {
                case 4294967295:
                    if (3.14159 != (1.5 - ((((0.1 * ((1.5 + ((0.0 - (3.14159 - float64((((v1 - v1) + v1) - v1)))) / float64(v1))) / ((1e10 / (3.14159 * 1e10)) - 0.1))) / 1.0) / 3.14159) / -2.25))) {
                        var v13 uint8 = 50
                        _ = v13
                    }
                case 0:
                    v14 := GenS1{F0: 255}
                    _ = v14
                default:
                    var v15 uint8 = 2
                    _ = v15
                }
                switch (v4 % ((fi11) | 1)) {
                case 2147483647:
                    v12 = append(v12, (^(v4)))
                }
                v7 -= (v3 - (v7 % ((v7) | 1)))
            }
            var v16 int = (v4 | v4)
            _ = v16
        } else {
            switch v2 {
            case "a":
                fmt.Println(v3)
            default:
                if true {
                    v8 = (^(v8))
                    fmt.Println(v3, v1)
                }
            }
            var v17 int64 = 0
            switch "\n\t" {
            case "a":
                for fi18 := 0; fi18 < 2; fi18++ {
                    var v19 bool = (!(-128 >= 127))
                    if v19 {
                        fmt.Println(v9, fi11, fi11)
                        var v20 float32 = -2.25
                        if ((v8 ^ v9) > 4294967295) {
                            switch (fi18 / ((fi11) | 1)) {
                            case -845:
                                v21 := GenS1{F0: 128}
                                _ = v21
                            case -2147483648:
                                var v22 uint32 = 883
                                _ = v22
                            case 127:
                                _ = len(v12)
                            default:
                                v23 := GenS1{F0: 206}
                                _ = v23
                            }
                            v24 := GenS1{F0: 27}
                            switch int(v3) {
                            case 127:
                                if len(v12) >= 2 {
                                    v12 = v12[1:len(v12)]
                                }
                                v6[240] = 918
                                var v25 uint = 0
                                v12 = append(v12, 127)
                                _ = v25
                            case -560:
                                var v26 uint64 = 2
                                _ = v26
                            case -2147483648:
                                v27 := GenS2{F0: -2147483648, F1: v4}
                                _ = v27
                            default:
                                if (false && (1 <= 0)) {
                                    var v28 uint16 = v3
                                    if ((v17 / ((v17) | 1)) < -1) {
                                        var v29 int = (^(fi18))
                                        var v30 int16 = 990
                                        _ = v29
                                        _ = v30
                                    }
                                    var v31 int64 = v17
                                    _ = v28
                                    _ = v31
                                } else {
                                    var v32 bool = v19
                                    if ((1 >= 1) || (255 <= 208)) {
                                        var v33 int16 = 1000
                                        var v34 float64 = 1.5
                                        var v35 int16 = -32768
                                        _ = v33
                                        _ = v34
                                        _ = v35
                                    }
                                    _ = v32
                                }
                            }
                            _ = v24
                        }
                        _ = v20
                    } else {
                        if v19 {
                            switch -948 {
                            case 127:
                                switch (^(fi18)) {
                                case -2147483648:
                                    v8 = (v9 << ((v8) & 63))
                                default:
                                    switch (v2 + v2) {
                                    case "hello":
                                        if (379 >= 2) {
                                            v3 = v3
                                        }
                                    case "a":
                                        fmt.Println(v4)
                                    case "\n\t":
                                        fmt.Println(v2, v7)
                                    default:
                                        v7 += v3
                                    }
                                }
                            default:
                                var v36 uint64 = 750
                                _ = v36
                            }
                            fmt.Println(fi18, fi18)
                            if (!v19) {
                                v37 := GenS2{F0: -2147483648, F1: v4}
                                _ = v37
                            } else {
                                switch (v4 | (fi18 | int(v7))) {
                                case -291:
                                    v38 := GenS1{F0: 183}
                                    _ = v38
                                case 4294967295:
                                    v39 := GenS1{F0: 2}
                                    _ = v39
                                case 1:
                                    switch v4 {
                                    case 845:
                                        v17 = v17
                                    case -565:
                                        v40 := v6[2147483648]
                                        _ = v40
                                    default:
                                        if len(v12) >= 2 {
                                            v12 = v12[1:len(v12)]
                                        }
                                    }
                                default:
                                    var v41 uint64 = 181
                                    _ = v41
                                }
                            }
                            var v42 string = "foo bar"
                            v3 += 781
                            _ = v42
                        } else {
                            delete(v5, uint32(v9))
                        }
                        v19 = (v19 || (v19 || v19))
                    }
                    if true {
                        fmt.Println(v8)
                        var v43 uint16 = (v3 / ((v7) | 1))
                        v1 = (((0.1 * 3.14159) / v1) / (v1 * (1e10 * (((0.0 / -2.25) * float32(3.14159)) * ((1.5 + v1) / (v1 + (v1 - v1)))))))
                        var v44 int8 = 100
                        _ = v43
                        _ = v44
                    } else {
                        var v45 string = (v2 + v2)
                        v46 := GenS1{F0: 255}
                        v47 := GenS2{F0: 2147483647, F1: -177}
                        _ = v45
                        _ = v46
                        _ = v47
                    }
                    fmt.Println(v3, v9)
                    fmt.Println(v8, v1, v7)
                    _ = fi18
                    _ = v19
                }
            case "foo bar":
                v2 = "\n\t"
            case "\n\t":
                for fi48 := 0; fi48 < 11; fi48++ {
                    v9 = (^(v8))
                    _ = fi48
                }
            }
            _ = v17
        }
        for fi49 := 0; fi49 < 12; fi49++ {
            fmt.Println(v7)
            fmt.Println(v8, fi11, fi49)
            _ = fi49
        }
        _ = fi11
        _ = v12
    }
    v50 := []int{v4, v4, -947}
    fmt.Println(v8, v1, v8)
    var v51 float64 = 0.0
    switch (v4 ^ (v4 * (^(v4)))) {
    case 32767:
        for fi52 := 0; fi52 < 1; fi52++ {
            fmt.Println(v3)
            fmt.Println(fi52)
            _ = fi52
        }
    }
    var v53 int64 = 0
    if (0 >= 1) {
        switch "" {
        case "a":
            if true {
                if (1 >= 4294967295) {
                    v8 += (^(v8))
                    if false {
                        fmt.Println(v7, v53, v3)
                        var v54 float32 = v1
                        if (0 >= 0) {
                            if false {
                                if true {
                                    v55 := GenS3{F0: v2, F1: 4294967295, F2: 0}
                                    _ = len(v5)
                                    _ = v55
                                }
                                _ = len(v50)
                            }
                            var v56 int16 = 1
                            _ = v56
                        } else {
                            var v57 uint64 = 42
                            v58 := GenS3{F0: v2, F1: v8, F2: 2147483647}
                            _ = v57
                            _ = v58
                        }
                        _ = v54
                    }
                } else {
                    v59 := GenS4{F0: (((((((v2 + (v2 + (v2 + (((v2 + v2) + v2) + ("foo bar" + ((("foo bar" + "foo bar") + "\n\t") + (v2 + (("foo bar" + v2) + "a")))))))) + v2) + v2) + ("\n\t" + v2)) + "hello") + v2) + v2), F1: 2147483648, F2: v8}
                    _ = v59
                }
                for fi60 := 0; fi60 < 6; fi60++ {
                    v51 = (1e10 + float64(v1))
                    if ((v3 << ((v9) & 63)) > v7) {
                        if true {
                            switch (v4 << ((v8) & 63)) {
                            case -2147483648:
                                for fi61 := 0; fi61 < 7; fi61++ {
                                    switch v2 {
                                    case "foo bar":
                                        v62 := GenS3{F0: ((("hello" + "hello") + v2) + v2), F1: (v9 | (v8 | v8)), F2: 0}
                                        _ = v62
                                    case "a":
                                        fmt.Println(v1, v53, fi61)
                                        v8 = v8
                                    }
                                    v63 := GenS4{F0: ((v2 + v2) + "foo bar"), F1: 65535, F2: (v8 >> ((v8) & 63))}
                                    _ = fi61
                                    _ = v63
                                }
                            }
                        }
                        var v64 int8 = -128
                        v9 = 2
                        _ = v64
                    }
                    v51 = 3.14159
                    _ = fi60
                }
                fmt.Println(v8, v2)
            } else {
                if (!(false && (-1 >= 0))) {
                    v65 := GenS5{F0: 1, F1: ("" <= "a")}
                    v66 := GenS6{F0: 0, F1: (v3 | (v3 | 719)), F2: 921}
                    for fi67 := 0; fi67 < 11; fi67++ {
                        fmt.Println(v9, v4, fi67)
                        fmt.Println(v3, v8)
                        var v68 float64 = v51
                        _ = fi67
                        _ = v68
                    }
                    _ = v65
                    _ = v66
                }
            }
        case "":
            fmt.Println(v2, v3, v1)
        }
        if false {
            v51 -= v51
        }
        switch v4 {
        case 9223372036854775807:
            v9 -= 65535
        case 4294967295:
            if ((!(-2147483648 <= 42)) && (int32(v53) <= -780)) {
                var v69 bool = ((^(v9)) != uint(v3))
                if false {
                    _ = len(v50)
                    for fi70 := 0; fi70 < 10; fi70++ {
                        var v71 bool = (v69 || true)
                        _ = fi70
                        _ = v71
                    }
                    _ = len(v5)
                } else {
                    v9 += (v9 % (((v9 ^ 1)) | 1))
                }
                _ = v69
            } else {
                if (!(42 == v4)) {
                    v72 := GenS4{F0: ((("hello" + "hello") + (v2 + v2)) + "\n\t"), F1: 479, F2: 255}
                    if (42 != 255) {
                        v8 = v9
                        v3 += uint16(v8)
                    } else {
                        var v73 int = (v4 >> ((v8) & 63))
                        _ = v73
                    }
                    if true {
                        for fi74 := 0; fi74 < 4; fi74++ {
                            for fi75 := 0; fi75 < 9; fi75++ {
                                var v76 uint32 = 438
                                v76 -= (^(v76))
                                _ = fi75
                                _ = v76
                            }
                            for fi77 := 0; fi77 < 13; fi77++ {
                                v51 = v51
                                if len(v50) >= 2 {
                                    v50 = v50[1:len(v50)]
                                }
                                _ = fi77
                            }
                            _ = fi74
                        }
                    } else {
                        v7 -= (v3 | (v3 & v7))
                        switch (v2 + v2) {
                        case "":
                            for fi78 := 0; fi78 < 7; fi78++ {
                                fmt.Println(v3)
                                v6[0] = 751
                                _ = fi78
                            }
                        case "a":
                            var v79 int = (v4 ^ v4)
                            _ = v79
                        default:
                            var v80 uint64 = 15
                            _ = v80
                        }
                    }
                    _ = v72
                }
            }
        case 1:
            _ = len(v6)
        default:
            v81 := GenS1{F0: 2}
            _ = v81
        }
    }
    v8 -= v8
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v50:", len(v50), v50)
    fmt.Println("v51:", v51)
    fmt.Println("v53:", v53)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v53
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
