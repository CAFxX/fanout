package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 uint32 }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 int64; F1 int8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int8; F1 uint8; F2 uint32 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int }

func main() {
    v1 := make(map[uint32]string)
    var v2 int64 = -1
    v3 := make(map[string]string)
    var v4 uint16 = 0
    v5 := []int{-2147483648, -9223372036854775808, -985}
    switch -2147483648 {
    case 2147483647:
        var v6 int8 = 0
        _ = v6
    case 32767:
        var v7 uint32 = 185
        _ = v7
    }
    v8 := GenS1{F0: -9223372036854775808, F1: 0, F2: 65535}
    var v9 uint64 = 42
    for fi10 := 0; fi10 < 1; fi10++ {
        v4 -= 89
        for fi11 := 0; fi11 < 18; fi11++ {
            fmt.Println(v4, v2, fi11)
            v9 -= v9
            _ = fi11
        }
        _ = fi10
    }
    var v12 uint16 = 2
    if true {
        for fi13 := 0; fi13 < 2; fi13++ {
            if true {
                var v14 uint16 = v12
                switch fi13 {
                case 128:
                    var v15 uint32 = 622
                    _ = v15
                }
                var v16 uint32 = 1
                _ = v14
                _ = v16
            } else {
                v12 -= (v12 - v4)
                v9 += (v9 | (v9 & 893))
                v17 := GenS2{F0: 2147483648, F1: uint64(v4)}
                _ = v17
            }
            for fi18 := 0; fi18 < 12; fi18++ {
                if false {
                    v19 := GenS2{F0: 2147483648, F1: 2}
                    if ((v9 >= v9) || (v2 == v2)) {
                        v5 = append(v5, -1)
                        switch "a" {
                        case "a":
                            v9 += (^(v9))
                        case "hello":
                            var v20 float64 = -2.25
                            _ = v20
                        default:
                            v2 += v2
                        }
                        if (!(1e10 == 0.0)) {
                            v21 := GenS2{F0: 42, F1: 335}
                            _ = len(v3)
                            _ = v21
                        }
                        if ((!true) || true) {
                            var v22 int32 = int32(v9)
                            v22 = (v22 << ((1) & 63))
                            var v23 float64 = float64(0.1)
                            if (65535 > 240) {
                                v4 = v12
                            }
                            fmt.Println(v23)
                            _ = v22
                            _ = v23
                        }
                    } else {
                        var v24 int = (fi13 - (fi13 & fi18))
                        v4 += (v4 | 0)
                        if (!false) {
                            if (true || true) {
                                v4 = 0
                                fmt.Println(fi18)
                            }
                            switch "foo bar" {
                            case "hello":
                                var v25 uint32 = 520
                                _ = v25
                            case "foo bar":
                                switch fi18 {
                                case 9223372036854775807:
                                    v24 = (v24 % ((fi18) | 1))
                                    v26 := GenS2{F0: 42, F1: (v9 ^ (v9 >> ((439) & 63)))}
                                    v27 := GenS2{F0: 65535, F1: (v9 + 173)}
                                    fmt.Println(fi13)
                                    if (!false) {
                                        v28 := GenS1{F0: v2, F1: -128, F2: 477}
                                        _ = v28
                                    } else {
                                        switch "foo bar" {
                                        case "\n\t":
                                            switch (v24 % ((0) | 1)) {
                                            case 1:
                                                var v29 int32 = 276
                                                _ = v29
                                            case 4294967295:
                                                if (2 >= 42) {
                                                    v4 = 255
                                                } else {
                                                    v30 := GenS2{F0: 42, F1: v9}
                                                    if true {
                                                        v31 := GenS2{F0: 2147483648, F1: (^(v9))}
                                                        _ = v31
                                                    }
                                                    v4 = v4
                                                    _ = v30
                                                }
                                            }
                                        }
                                    }
                                    _ = v26
                                    _ = v27
                                default:
                                    var v32 string = ""
                                    _ = v32
                                }
                            default:
                                var v33 uint = 42
                                _ = v33
                            }
                            switch "" {
                            case "hello":
                                v24 += 32767
                            }
                        } else {
                            var v34 float64 = 0.0
                            v34 += v34
                            switch ((("\n\t" + (("\n\t" + ("hello" + "")) + "hello")) + ("foo bar" + "")) + ("" + "foo bar")) {
                            case "":
                                v3["foo bar"] = "a"
                            case "foo bar":
                                var v35 uint32 = 42
                                _ = v35
                            default:
                                fmt.Println(v24, v9)
                            }
                            _ = v34
                        }
                        v24 -= v24
                        switch 0 {
                        case 9223372036854775807:
                            v5 = append(v5, 42)
                        }
                        if (65535 < 2) {
                            v36 := GenS3{F0: (v2 % (((v2 - v2)) | 1)), F1: -128}
                            var v37 uint8 = 128
                            _ = v36
                            _ = v37
                        } else {
                            v24 -= (fi13 & 42)
                            switch "a" {
                            case "\n\t":
                                v12 += 42
                                delete(v3, "a")
                                if ((v2 >> ((65535) & 63)) != (v2 / (((v2 | v2)) | 1))) {
                                    _ = len(v5)
                                    var v38 uint16 = (^(v12))
                                    _ = v38
                                }
                            case "hello":
                                delete(v1, 255)
                            default:
                                var v39 string = "hello"
                                v9 += 229
                                v2 = -64
                                _ = v39
                            }
                            var v40 int16 = -466
                            fmt.Println(fi13, v2)
                            switch 0 {
                            case 227:
                                if len(v5) > 0 {
                                    v5[0] = (-(v24))
                                }
                            case 1:
                                v41 := GenS4{F0: 0}
                                fmt.Println(fi18, fi18)
                                v3[""] = ("a" + "hello")
                                _ = v41
                            }
                            _ = v40
                        }
                        _ = v24
                    }
                    var v42 float32 = float32(0.0)
                    v2 += (^(v2))
                    switch "foo bar" {
                    case "a":
                        v2 = (v2 ^ v2)
                    case "hello":
                        switch "a" {
                        case "a":
                            var v43 int = (fi18 % ((-98) | 1))
                            _ = v43
                        case "":
                            v42 += float32(0.0)
                        default:
                            if false {
                                switch "foo bar" {
                                case "foo bar":
                                    v44 := v1[1]
                                    _ = v44
                                case "a":
                                    if (v12 == v4) {
                                        v9 += 468
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        v9 = (v9 ^ v9)
                                        fmt.Println(v2, fi13, fi18)
                                        v4 -= 870
                                    }
                                }
                            }
                        }
                        v3["a"] = "a"
                        if false {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            v45 := GenS4{F0: -268}
                            v46 := GenS1{F0: 184, F1: 1, F2: 576}
                            switch "" {
                            case "\n\t":
                                if (v9 > (v9 ^ v9)) {
                                    v4 -= 255
                                    fmt.Println(v2)
                                    var v47 int32 = 42
                                    _ = v47
                                } else {
                                    var v48 int16 = 32767
                                    if false {
                                        if false {
                                            if true {
                                                if true {
                                                    if true {
                                                        switch 9223372036854775807 {
                                                        case -2147483648:
                                                            fmt.Println(fi13)
                                                        case -9223372036854775808:
                                                            fmt.Println(v9, fi18)
                                                        case 128:
                                                            fmt.Println(fi18, fi18, fi13)
                                                        }
                                                        if len(v5) > 0 {
                                                            v5[0] = (fi18 | (-(fi13)))
                                                        }
                                                        var v49 int8 = -65
                                                        _ = v49
                                                    } else {
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                    }
                                                }
                                            } else {
                                                v50 := v1[2]
                                                _ = v50
                                            }
                                            delete(v3, "foo bar")
                                        }
                                        v51 := GenS2{F0: 28, F1: v9}
                                        v52 := v3["hello"]
                                        _ = v51
                                        _ = v52
                                    }
                                    _ = v48
                                }
                            case "":
                                var v53 bool = (true || false)
                                _ = v53
                            case "hello":
                                var v54 uint32 = 2
                                _ = v54
                            }
                            _ = v45
                            _ = v46
                        } else {
                            switch fi18 {
                            case 127:
                                fmt.Println(v42)
                                v9 -= (v9 + (v9 >> ((2) & 63)))
                                if false {
                                    switch "\n\t" {
                                    case "":
                                        switch "a" {
                                        case "":
                                            if true {
                                                fmt.Println(v4)
                                            }
                                            if (!(!false)) {
                                                v9 = (^(v9))
                                                v42 += ((0.0 - v42) + 0.0)
                                                v9 = v9
                                            } else {
                                                var v55 uint8 = 42
                                                _ = v55
                                            }
                                            fmt.Println(v4, v9)
                                        case "a":
                                            v56 := GenS1{F0: -1, F1: 0, F2: 273}
                                            _ = v56
                                        case "\n\t":
                                            v2 = v2
                                        default:
                                            v9 = (v9 % ((173) | 1))
                                        }
                                    case "\n\t":
                                        v3[("" + "")] = "\n\t"
                                    }
                                }
                                var v57 uint8 = 28
                                _ = v57
                            default:
                                switch "foo bar" {
                                case "foo bar":
                                    if len(v5) > 0 {
                                        v5[0] = fi13
                                    }
                                case "\n\t":
                                    var v58 uint = uint(v9)
                                    _ = v58
                                case "hello":
                                    var v59 string = ("foo bar" + "\n\t")
                                    _ = v59
                                }
                            }
                            fmt.Println(v42, v42)
                        }
                        v42 = (0.1 * v42)
                    case "\n\t":
                        v60 := GenS5{F0: -81, F1: 0, F2: 111}
                        _ = v60
                    default:
                        v2 -= v2
                    }
                    _ = v19
                    _ = v42
                }
                v61 := v1[377]
                var v62 int8 = 84
                _ = fi18
                _ = v61
                _ = v62
            }
            _ = fi13
        }
        if false {
            if true {
                var v63 uint64 = 944
                var v64 string = "foo bar"
                _ = v63
                _ = v64
            } else {
                for fi65 := 0; fi65 < 12; fi65++ {
                    v66 := GenS2{F0: 711, F1: 93}
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    v12 = (v4 >> ((968) & 63))
                    _ = fi65
                    _ = v66
                }
            }
            v2 -= v2
        } else {
            var v67 int32 = 42
            v9 = 1
            _ = v67
        }
    } else {
        v9 = v9
        v12 -= v12
        var v68 uint16 = 2
        _ = v68
    }
    v69 := GenS6{F0: 42}
    var v70 int64 = -878
    v71 := GenS3{F0: 0, F1: 78}
    var v72 bool = ((true && true) || true)
    switch ("" + ((("foo bar" + "hello") + (("\n\t" + "a") + ("hello" + (("\n\t" + "a") + ("foo bar" + (("\n\t" + "\n\t") + "foo bar")))))) + "foo bar")) {
    case "\n\t":
        if v72 {
            v73 := GenS7{F0: 265}
            _ = v73
        }
    }
    if (v72 && v72) {
        v74 := GenS7{F0: 9223372036854775807}
        _ = len(v5)
        _ = v74
    }
    v5 = append(v5, 1)
    for fi75 := 0; fi75 < 13; fi75++ {
        v5 = append(v5, (fi75 << ((4294967295) & 63)))
        for fi76 := 0; fi76 < 11; fi76++ {
            var v77 uint = 255
            _ = fi76
            _ = v77
        }
        v78 := GenS3{F0: (v70 - (^(v70))), F1: 0}
        _ = fi75
        _ = v78
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v70:", v70)
    fmt.Println("v72:", v72)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v69
    _ = v70
    _ = v71
    _ = v72
    _ = v8
    _ = v9
}
