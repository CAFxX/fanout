package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 string; F2 uint }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint; F1 uint64; F2 uint64 }

func main() {
    var v1 uint16 = 748
    var v2 uint = 0
    var v3 int32 = 2147483647
    var v4 int16 = 32767
    var v5 int32 = v3
    var v6 float32 = 0.1
    if false {
        v3 += v5
        switch "foo bar" {
        case "":
            fmt.Println(v5)
        case "\n\t":
            for fi7 := 0; fi7 < 13; fi7++ {
                v8 := []int{fi7}
                _ = fi7
                _ = v8
            }
        case "foo bar":
            var v9 string = "hello"
            _ = v9
        default:
            v10 := []int{542, 0}
            _ = v10
        }
    }
    switch "a" {
    case "hello":
        v3 = -207
    case "\n\t":
        v11 := GenS1{F0: -9223372036854775808}
        _ = v11
    case "":
        v12 := GenS2{F0: (^(v1)), F1: "a", F2: (v2 ^ 710)}
        _ = v12
    }
    for fi13 := 0; fi13 < 20; fi13++ {
        fmt.Println(v2)
        _ = fi13
    }
    for fi14 := 0; fi14 < 18; fi14++ {
        v1 -= (v1 << ((v2) & 63))
        _ = fi14
    }
    if (!false) {
        var v15 int = 4294967295
        _ = v15
    } else {
        var v16 float64 = (0.0 + (1e10 + -2.25))
        switch "foo bar" {
        case "foo bar":
            if true {
                var v17 uint16 = 42
                switch "a" {
                case "foo bar":
                    if true {
                        var v18 int32 = 0
                        var v19 float32 = 0.0
                        v20 := []int{}
                        _ = v18
                        _ = v19
                        _ = v20
                    }
                case "hello":
                    v5 = 2147483647
                }
                _ = v17
            }
        case "a":
            var v21 int16 = (v4 ^ (v4 / ((int16(v1)) | 1)))
            _ = v21
        }
        if (42 < -1) {
            var v22 uint = (v2 % ((v2) | 1))
            _ = v22
        } else {
            if (!true) {
                fmt.Println(v2, v4)
                var v23 uint16 = v1
                _ = v23
            } else {
                switch -317 {
                case 731:
                    v6 = ((v6 * (0.1 + (((v6 + (v6 + ((v6 / (float32((v16 - (1e10 - 0.1))) / (-2.25 / ((v6 + (v6 / 3.14159)) + v6)))) * (v6 * v6)))) + v6) / v6))) + v6)
                case -9223372036854775808:
                    if true {
                        v24 := GenS2{F0: (v1 * (v1 % (((^(v1))) | 1))), F1: "a", F2: (v2 + (v2 + v2))}
                        v1 = (v1 * uint16(v2))
                        v4 = v4
                        _ = v24
                    } else {
                        switch "hello" {
                        case "a":
                            var v25 uint = (v2 / ((v2) | 1))
                            _ = v25
                        default:
                            if true {
                                v26 := []int{4294967295, -841}
                                _ = v26
                            } else {
                                fmt.Println(v6, v3)
                            }
                        }
                        for fi27 := 0; fi27 < 7; fi27++ {
                            var v28 int16 = (v4 + (v4 / ((1000) | 1)))
                            _ = fi27
                            _ = v28
                        }
                        if (556 >= 1) {
                            switch "a" {
                            case "":
                                v29 := GenS3{F0: true}
                                _ = v29
                            case "a":
                                var v30 int64 = 149
                                _ = v30
                            default:
                                v31 := GenS3{F0: (4294967295 <= 4294967295)}
                                _ = v31
                            }
                            if false {
                                var v32 uint32 = uint32(v2)
                                fmt.Println(v16, v32)
                                if ((!(!true)) || (!false)) {
                                    v4 = v4
                                }
                                _ = v32
                            }
                            v33 := GenS4{F0: v2, F1: 501, F2: 42}
                            _ = v33
                        } else {
                            v34 := []int{9223372036854775807, 42, -403}
                            _ = v34
                        }
                    }
                case 116:
                    v35 := []int{-691}
                    _ = v35
                default:
                    for fi36 := 0; fi36 < 17; fi36++ {
                        if true {
                            switch fi36 {
                            case -2147483648:
                                var v37 float64 = -2.25
                                _ = v37
                            case 1:
                                for fi38 := 0; fi38 < 9; fi38++ {
                                    var v39 string = ""
                                    if (!true) {
                                        var v40 int64 = 728
                                        if (true && (!(2 == 471))) {
                                            switch fi38 {
                                            case -9223372036854775808:
                                                v41 := []int{fi36, fi38, (fi36 >> ((v2) & 63))}
                                                _ = v41
                                            case -169:
                                                if false {
                                                    v42 := []int{}
                                                    _ = v42
                                                } else {
                                                    fmt.Println(fi36)
                                                    if (((v2 % (((v2 & v2)) | 1)) != v2) || true) {
                                                        if true {
                                                            v16 = (v16 + (float64(float32(v16)) * float64(v6)))
                                                            var v43 float32 = 3.14159
                                                            v40 = 42
                                                            var v44 float32 = v6
                                                            var v45 uint16 = 284
                                                            v46 := GenS4{F0: (v2 << ((2) & 63)), F1: 16, F2: 18446744073709551615}
                                                            switch v39 {
                                                            case "":
                                                                fmt.Println(v40, v16, v6)
                                                            case "a":
                                                                var v47 float64 = 1e10
                                                                _ = v47
                                                            case "hello":
                                                                v48 := GenS1{F0: (v40 - (v40 >> ((v2) & 63)))}
                                                                _ = v48
                                                            }
                                                            v49 := GenS3{F0: ((v44 + v6) != float32(((v16 / 1.0) + (v16 - (((0.0 * -2.25) - (v16 / 1e10)) + v16)))))}
                                                            var v50 bool = false
                                                            _ = v43
                                                            _ = v44
                                                            _ = v45
                                                            _ = v46
                                                            _ = v49
                                                            _ = v50
                                                        } else {
                                                            v2 -= 2
                                                        }
                                                        v51 := []int{(^(fi38)), (fi38 % (((^(fi36))) | 1))}
                                                        v40 = v40
                                                        _ = v51
                                                    }
                                                }
                                            case 1:
                                                switch -1 {
                                                case 0:
                                                    fmt.Println(v5)
                                                }
                                            }
                                            var v52 uint16 = (v1 >> (((v2 | v2)) & 63))
                                            var v53 bool = ((-(v4)) <= -32768)
                                            _ = v52
                                            _ = v53
                                        } else {
                                            var v54 int64 = -9223372036854775808
                                            var v55 int32 = 854
                                            var v56 uint8 = uint8(v1)
                                            _ = v54
                                            _ = v55
                                            _ = v56
                                        }
                                        v3 -= -462
                                        _ = v40
                                    }
                                    _ = fi38
                                    _ = v39
                                }
                            case 2147483647:
                                var v57 int8 = 33
                                _ = v57
                            default:
                                if false {
                                    switch "\n\t" {
                                    case "a":
                                        for fi58 := 0; fi58 < 15; fi58++ {
                                            v6 += (v6 / 3.14159)
                                            v4 += v4
                                            _ = fi58
                                        }
                                    case "":
                                        fmt.Println(v5, v6, fi36)
                                    default:
                                        if false {
                                            fmt.Println(v1)
                                        }
                                    }
                                    var v59 int32 = (v5 & v3)
                                    var v60 int64 = 42
                                    _ = v59
                                    _ = v60
                                }
                            }
                        }
                        _ = fi36
                    }
                }
                fmt.Println(v4)
            }
        }
        _ = v16
    }
    if true {
        var v61 uint16 = (v1 - (v1 + v1))
        v1 = v1
        _ = v61
    }
    if (!(42 > 569)) {
        var v62 float64 = ((1.5 + ((3.14159 + 1.5) / (1e10 + 3.14159))) + 1.5)
        _ = v62
    } else {
        v3 = (v5 & 988)
        for fi63 := 0; fi63 < 10; fi63++ {
            fmt.Println(fi63)
            var v64 uint64 = 4294967295
            switch (-(fi63)) {
            case -2147483648:
                v2 = v2
            default:
                v65 := []int{32767}
                _ = v65
            }
            _ = fi63
            _ = v64
        }
    }
    v6 = v6
    for fi66 := 0; fi66 < 18; fi66++ {
        if true {
            v67 := []int{-2147483648, (^(fi66)), (fi66 >> ((v2) & 63)), (fi66 | -9223372036854775808)}
            fmt.Println(v6)
            _ = v67
        }
        var v68 int = (fi66 >> ((v2) & 63))
        _ = fi66
        _ = v68
    }
    var v69 int32 = v5
    for fi70 := 0; fi70 < 8; fi70++ {
        var v71 string = ("" + "")
        _ = fi70
        _ = v71
    }
    var v72 float32 = (float32((3.14159 * ((3.14159 * 3.14159) * 0.1))) + v6)
    v73 := []int{int(v4), 42, 1}
    if len(v73) > 0 {
        v73[0] = 32767
    }
    if false {
        if true {
            if (9223372036854775807 > 4294967295) {
                v73 = append(v73, 2147483647)
            } else {
                var v74 int = 127
                _ = v74
            }
        }
        var v75 uint64 = 0
        _ = v75
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v69:", v69)
    fmt.Println("v72:", v72)
    fmt.Println("slicelen v73:", len(v73), v73)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v69
    _ = v72
    _ = v73
}
