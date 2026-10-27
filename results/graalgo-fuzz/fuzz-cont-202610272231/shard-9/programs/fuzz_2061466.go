package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 bool; F1 uint8; F2 uint16 }
type GenS3 struct { F0 uint; F1 bool }
type GenS4 struct { F0 int16; F1 uint64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int64; F1 bool; F2 uint32 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int8; F1 int8; F2 int32 }
type GenS9 struct { F0 int64; F1 string; F2 int64 }
type GenS10 struct { F0 uint8; F1 int }
type GenS11 struct { F0 int; F1 string; F2 string }
type GenS12 struct { F0 uint16; F1 int }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]int)
    var v3 uint16 = 65535
    var v4 float64 = float64(0.0)
    v5 := make(map[string]uint32)
    if false {
        if false {
            var v6 uint16 = 42
            _ = v6
        }
    } else {
        v7 := GenS1{F0: 13, F1: -105}
        for fi8 := 0; fi8 < 17; fi8++ {
            for fi9 := 0; fi9 < 19; fi9++ {
                var v10 bool = (!(-821 != 332))
                v11 := GenS1{F0: 514, F1: 1}
                var v12 float32 = float32(v4)
                _ = fi9
                _ = v10
                _ = v11
                _ = v12
            }
            _ = len(v5)
            var v13 int32 = -33
            _ = fi8
            _ = v13
        }
        _ = v7
    }
    var v14 uint16 = 0
    _ = len(v2)
    var v15 uint64 = 4294967295
    v16 := []int{-1, 127, 32768, 490}
    if false {
        if len(v16) > 0 {
            v16[0] = 1
        }
        fmt.Println(v14, v3, v3)
    }
    if (true && (!true)) {
        v17 := GenS1{F0: 409, F1: 100}
        switch "" {
        case "foo bar":
            for fi18 := 0; fi18 < 11; fi18++ {
                var v19 uint64 = (v15 << ((1) & 63))
                for fi20 := 0; fi20 < 17; fi20++ {
                    switch (^(fi18)) {
                    case 2147483647:
                        if false {
                            v4 -= (-2.25 * (v4 + v4))
                        } else {
                            if true {
                                if true {
                                    v2["hello"] = -2147483648
                                } else {
                                    var v21 uint16 = (v3 + 32767)
                                    _ = v21
                                }
                            } else {
                                v19 -= v19
                                v22 := GenS1{F0: 364, F1: 42}
                                v15 -= (v15 >> ((1) & 63))
                                var v23 uint32 = 220
                                _ = v22
                                _ = v23
                            }
                            v24 := GenS2{F0: ((fi20 > -2147483648) || (!(int8(v3) >= -1))), F1: 85, F2: v14}
                            if (!((v3 * (v3 & v14)) > (^(v3)))) {
                                if (true && false) {
                                    v25 := GenS1{F0: uint32(v3), F1: -128}
                                    fmt.Println(v14, fi20, v15)
                                    fmt.Println(v15)
                                    switch fi20 {
                                    case -2147483648:
                                        var v26 float64 = 1e10
                                        _ = v26
                                    case 1:
                                        v19 += (v15 & 149)
                                    default:
                                        v27 := GenS3{F0: 4294967295, F1: false}
                                        _ = v27
                                    }
                                    if (!false) {
                                        v15 = (^(v19))
                                        v4 = v4
                                    }
                                    _ = v25
                                } else {
                                    v14 += (v14 % (((v14 * v3)) | 1))
                                    var v28 int64 = 9223372036854775807
                                    var v29 int32 = 0
                                    switch fi18 {
                                    case 4294967295:
                                        var v30 float64 = (1e10 - -2.25)
                                        _ = v30
                                    case -2147483648:
                                        if (596 != v19) {
                                            v31 := GenS2{F0: true, F1: 2, F2: v14}
                                            if (((0.0 + 0.1) >= (float32(1.5) - float32(v4))) && true) {
                                                if (255 != 42) {
                                                    v32 := GenS4{F0: -434, F1: v19}
                                                    _ = v32
                                                } else {
                                                    switch "foo bar" {
                                                    case "":
                                                        v29 += v29
                                                    case "\n\t":
                                                        fmt.Println(v19)
                                                    case "hello":
                                                        var v33 int16 = 42
                                                        _ = v33
                                                    default:
                                                        if len(v16) >= 2 {
                                                            v16 = v16[1:len(v16)]
                                                        }
                                                    }
                                                    _ = len(v5)
                                                    var v34 bool = true
                                                    var v35 int32 = (v29 << ((65535) & 63))
                                                    _ = v34
                                                    _ = v35
                                                }
                                            }
                                            fmt.Println(v14, v4)
                                            _ = v31
                                        }
                                    case 32768:
                                        v4 += (v4 / ((v4 - float64(0.1)) + 0.0))
                                    default:
                                        var v36 uint32 = 2147483648
                                        _ = v36
                                    }
                                    if len(v16) >= 2 {
                                        v16 = v16[1:len(v16)]
                                    }
                                    if false {
                                        switch (fi20 % ((fi20) | 1)) {
                                        case 543:
                                            v19 += v15
                                            var v37 float64 = -2.25
                                            _ = v37
                                        case -2147483648:
                                            v14 -= (v3 + 500)
                                        default:
                                            var v38 uint8 = 128
                                            _ = v38
                                        }
                                    } else {
                                        var v39 uint = 0
                                        switch (fi20 ^ (fi20 ^ 128)) {
                                        case 32767:
                                            v28 -= v28
                                        case -936:
                                            switch "\n\t" {
                                            case "\n\t":
                                                var v40 int64 = (-(v28))
                                                _ = v40
                                            case "a":
                                                var v41 bool = ((v14 - 0) >= v14)
                                                _ = v41
                                            }
                                        }
                                        _ = v39
                                    }
                                    if (0 < 128) {
                                        if len(v16) > 0 {
                                            v16[0] = fi18
                                        }
                                        v28 = -486
                                    } else {
                                        v3 += (v3 - v14)
                                    }
                                    if (100 != 100) {
                                        var v42 float32 = 0.1
                                        _ = v42
                                    } else {
                                        v29 -= v29
                                    }
                                    if true {
                                        v43 := GenS5{F0: (fi20 * (fi18 % ((-9223372036854775808) | 1)))}
                                        if len(v16) >= 2 {
                                            v16 = v16[1:len(v16)]
                                        }
                                        v19 -= 42
                                        _ = v43
                                    }
                                    _ = v28
                                    _ = v29
                                }
                                if ("foo bar" < "") {
                                    if ((1.5 - ((-2.25 - v4) + v4)) <= (0.1 - 1e10)) {
                                        var v44 uint16 = (v3 << ((712) & 63))
                                        if false {
                                            var v45 float32 = 0.0
                                            _ = v45
                                        }
                                        v15 = (v19 * uint64(v3))
                                        _ = v44
                                    } else {
                                        var v46 string = "hello"
                                        _ = v46
                                    }
                                    switch "" {
                                    case "\n\t":
                                        _ = len(v1)
                                        v19 = (v19 | v19)
                                        if false {
                                            var v47 int8 = 123
                                            v48 := GenS6{F0: 42, F1: true, F2: 65535}
                                            v49 := GenS7{F0: 180}
                                            _ = v47
                                            _ = v48
                                            _ = v49
                                        } else {
                                            delete(v2, ("" + "foo bar"))
                                            _ = len(v5)
                                            v16 = append(v16, (-(fi18)))
                                        }
                                    case "hello":
                                        v3 = (v3 % ((42) | 1))
                                    default:
                                        switch "" {
                                        case "hello":
                                            v50 := GenS5{F0: (fi20 << ((0) & 63))}
                                            _ = v50
                                        case "\n\t":
                                            var v51 string = "\n\t"
                                            _ = v51
                                        case "foo bar":
                                            if false {
                                                var v52 float64 = (-2.25 + v4)
                                                v15 = v15
                                                var v53 float32 = 0.0
                                                _ = v52
                                                _ = v53
                                            }
                                        }
                                    }
                                } else {
                                    var v54 int64 = -9223372036854775808
                                    fmt.Println(fi18, v15)
                                    _ = v54
                                }
                                var v55 uint = 42
                                _ = v55
                            } else {
                                if true {
                                    if true {
                                        v3 += v3
                                    } else {
                                        var v56 bool = (float64(0.0) < v4)
                                        _ = v56
                                    }
                                    v57 := GenS6{F0: 0, F1: (127 != -1), F2: uint32(v19)}
                                    _ = v57
                                } else {
                                    if ((!((^(fi18)) > -1)) || false) {
                                        var v58 int = 102
                                        v59 := v2[""]
                                        v5["\n\t"] = 446
                                        var v60 float64 = (v4 / v4)
                                        v60 -= v4
                                        _ = v58
                                        _ = v59
                                        _ = v60
                                    }
                                }
                                fmt.Println(v19, fi20, fi18)
                                v4 = (v4 / (v4 * (1e10 * v4)))
                            }
                            _ = v24
                        }
                    case -274:
                        switch "hello" {
                        case "hello":
                            v15 += (^(v19))
                        default:
                            switch fi18 {
                            case -1:
                                if false {
                                    v5[("a" + "foo bar")] = 444
                                    switch "hello" {
                                    case "foo bar":
                                        v61 := GenS8{F0: 78, F1: -128, F2: 42}
                                        _ = v61
                                    }
                                    _ = len(v16)
                                    v3 = (^(v14))
                                    delete(v2, "hello")
                                } else {
                                    delete(v2, "foo bar")
                                    var v62 uint64 = (^(v15))
                                    switch (fi20 << ((1) & 63)) {
                                    case 4294967295:
                                        if (-431 != 0) {
                                            var v63 int16 = -32768
                                            _ = v63
                                        }
                                    case -438:
                                        v4 -= 3.14159
                                    case 9223372036854775807:
                                        fmt.Println(v15, v19, fi20)
                                    default:
                                        var v64 string = (("\n\t" + "a") + "\n\t")
                                        _ = v64
                                    }
                                    _ = v62
                                }
                            case 32767:
                                if false {
                                    if false {
                                        v65 := GenS9{F0: 9223372036854775807, F1: ("hello" + "hello"), F2: int64(v14)}
                                        _ = v65
                                    }
                                    v66 := GenS4{F0: 489, F1: v19}
                                    var v67 uint64 = v15
                                    _ = v66
                                    _ = v67
                                }
                            case 2147483647:
                                v68 := GenS10{F0: 0, F1: fi18}
                                _ = v68
                            }
                        }
                    }
                    v3 = (v3 + 1)
                    var v69 uint32 = 4294967295
                    _ = fi20
                    _ = v69
                }
                v4 -= (0.1 + v4)
                _ = fi18
                _ = v19
            }
        case "\n\t":
            fmt.Println(v14, v15, v14)
        }
        _ = v17
    } else {
        var v70 bool = (false || true)
        _ = v70
    }
    delete(v2, "a")
    v3 -= 2
    for fi71 := 0; fi71 < 10; fi71++ {
        switch 128 {
        case 988:
            var v72 float32 = ((float32(float64((1.5 * 1e10))) / 1.5) - float32((v4 / float64(1.5))))
            _ = v72
        case -1:
            var v73 int16 = 42
            _ = v73
        case 2147483647:
            var v74 int16 = 1000
            _ = v74
        }
        var v75 uint64 = v15
        v4 = v4
        _ = fi71
        _ = v75
    }
    fmt.Println(v14, v15)
    v76 := GenS11{F0: 32768, F1: (("\n\t" + "\n\t") + "\n\t"), F2: ("hello" + ("foo bar" + "a"))}
    v15 += v15
    fmt.Println(v3, v14)
    for fi77 := 0; fi77 < 3; fi77++ {
        v78 := v2["hello"]
        v79 := GenS12{F0: (^(v14)), F1: fi77}
        for fi80 := 0; fi80 < 8; fi80++ {
            v4 = ((float64(1e10) - v4) / ((v4 / 1.5) + 0.0))
            v15 = v15
            _ = len(v2)
            _ = fi80
        }
        _ = fi77
        _ = v78
        _ = v79
    }
    if len(v16) > 0 {
        v16[0] = -2147483648
    }
    v3 -= v14
    v81 := GenS9{F0: 9223372036854775807, F1: "\n\t", F2: -9223372036854775808}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v76
    _ = v81
}
