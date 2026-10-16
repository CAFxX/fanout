package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32; F1 int8; F2 int }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int32; F1 int32; F2 uint32 }
type GenS7 struct { F0 int64; F1 int }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint8 = 128
    v3 := []uint32{}
    v2 += (v2 * (v2 * (v2 + 128)))
    v2 = 0
    if false {
        for fi4 := 0; fi4 < 2; fi4++ {
            v5 := GenS1{F0: 32767, F1: 1, F2: fi4}
            v1[435] = false
            _ = fi4
            _ = v5
        }
        var v6 float64 = (-2.25 - float64(3.14159))
        v7 := GenS2{F0: 42}
        _ = v6
        _ = v7
    } else {
        if false {
            v2 = (v2 % (((v2 / (((v2 >> ((65535) & 63))) | 1))) | 1))
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
    }
    var v8 uint = 1
    v8 -= (^(v8))
    switch "\n\t" {
    case "\n\t":
        v8 = (^(v8))
    case "hello":
        switch 4294967295 {
        case -9223372036854775808:
            fmt.Println(v2)
        case 128:
            var v9 uint8 = (v2 / ((v2) | 1))
            _ = v9
        case -296:
            _ = len(v1)
        }
    case "":
        var v10 int64 = 9223372036854775807
        _ = v10
    }
    var v11 uint8 = v2
    switch "foo bar" {
    case "a":
        v11 += (v11 & v11)
    case "":
        for fi12 := 0; fi12 < 7; fi12++ {
            fmt.Println(fi12, v2, v2)
            var v13 int64 = -1
            for fi14 := 0; fi14 < 19; fi14++ {
                v15 := GenS1{F0: 322, F1: 533, F2: 127}
                if len(v3) > 0 {
                    v3[0] = 0
                }
                if false {
                    if false {
                        var v16 int64 = v13
                        _ = v16
                    } else {
                        fmt.Println(v2, v13)
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        var v17 float64 = (1.5 / 1.0)
                        _ = v17
                    }
                } else {
                    if len(v3) > 0 {
                        v3[0] = 65535
                    }
                    v11 = 176
                }
                _ = fi14
                _ = v15
            }
            _ = fi12
            _ = v13
        }
    }
    switch "foo bar" {
    case "\n\t":
        v18 := GenS2{F0: -1}
        _ = v18
    case "a":
        fmt.Println(v11, v2)
    case "foo bar":
        v3 = append(v3, 2147483648)
    default:
        for fi19 := 0; fi19 < 8; fi19++ {
            var v20 uint32 = 544
            var v21 bool = false
            var v22 uint = (^(v8))
            _ = fi19
            _ = v20
            _ = v21
            _ = v22
        }
    }
    v8 -= (v8 << ((4294967295) & 63))
    if false {
        for fi23 := 0; fi23 < 7; fi23++ {
            fmt.Println(v11, v11, fi23)
            delete(v1, 87)
            _ = fi23
        }
        v24 := v1[1]
        for fi25 := 0; fi25 < 11; fi25++ {
            for fi26 := 0; fi26 < 14; fi26++ {
                if true {
                    if v24 {
                        if v24 {
                            fmt.Println(fi26, v11)
                            var v27 int = (fi25 | fi25)
                            _ = v27
                        } else {
                            v11 = (v2 % ((v2) | 1))
                            fmt.Println(fi26, v24, fi26)
                            fmt.Println(fi25)
                        }
                        v28 := GenS3{F0: 42, F1: -128, F2: fi25}
                        v29 := GenS1{F0: 32767, F1: 65535, F2: -526}
                        _ = v28
                        _ = v29
                    } else {
                        if v24 {
                            v8 -= (^(v8))
                            v11 += v2
                        } else {
                            var v30 int64 = 0
                            v8 += v8
                            _ = v30
                        }
                        var v31 int64 = 1
                        _ = v31
                    }
                    switch (fi25 % (((^(fi26))) | 1)) {
                    case 32768:
                        if ((^(v11)) <= 155) {
                            v32 := GenS4{F0: v24}
                            _ = v32
                        }
                    }
                    if (v2 >= v2) {
                        if v24 {
                            fmt.Println(fi26)
                            var v33 float32 = 1.5
                            _ = v33
                        } else {
                            if true {
                                v8 -= 255
                                v8 -= v8
                                if false {
                                    if v24 {
                                        switch (fi26 + (fi25 << (((^(v8))) & 63))) {
                                        case -9223372036854775808:
                                            var v34 int = -393
                                            _ = v34
                                        case 341:
                                            v2 += (v2 << ((0) & 63))
                                        }
                                        var v35 int32 = -1
                                        _ = v35
                                    } else {
                                        v36 := GenS5{F0: 2}
                                        _ = v36
                                    }
                                    var v37 uint32 = 2147483648
                                    _ = v37
                                } else {
                                    switch "foo bar" {
                                    case "a":
                                        switch fi25 {
                                        case 0:
                                            fmt.Println(v11)
                                        }
                                    case "hello":
                                        v24 = (428 == 42)
                                    default:
                                        if ((fi25 | fi26) == -9223372036854775808) {
                                            var v38 float32 = (1.5 / float32((float64((0.1 - 1.5)) / -2.25)))
                                            if (-2147483648 < (^(fi26))) {
                                                if v24 {
                                                    fmt.Println(v24, v2)
                                                    v39 := GenS1{F0: 768, F1: 2147483647, F2: -391}
                                                    v40 := GenS6{F0: 209, F1: -2147483648, F2: 65535}
                                                    v2 -= (v2 >> ((v8) & 63))
                                                    _ = v39
                                                    _ = v40
                                                } else {
                                                    var v41 uint32 = 2147483647
                                                    if v24 {
                                                        var v42 uint = uint(v41)
                                                        _ = v42
                                                    }
                                                    _ = v41
                                                }
                                                v11 = v11
                                                var v43 uint32 = 2147483647
                                                _ = v43
                                            }
                                            _ = len(v3)
                                            _ = v38
                                        }
                                    }
                                }
                                v8 = (v8 - (v8 << ((v8) & 63)))
                            }
                            v44 := GenS1{F0: 255, F1: 2147483648, F2: fi26}
                            switch fi26 {
                            case 0:
                                var v45 int = 32768
                                _ = v45
                            case 32768:
                                v24 = false
                                v3 = append(v3, 2147483647)
                                var v46 int64 = -28
                                _ = v46
                            }
                            _ = v44
                        }
                    }
                    switch -9223372036854775808 {
                    case -9223372036854775808:
                        v8 = (v8 >> (((v8 + (v8 & v8))) & 63))
                    case 981:
                        v3 = append(v3, 840)
                    }
                    var v47 int16 = 42
                    v48 := GenS6{F0: -1, F1: 42, F2: 4294967295}
                    _ = v47
                    _ = v48
                }
                _ = fi26
            }
            if v24 {
                switch (fi25 & 127) {
                case -1:
                    var v49 int64 = -1
                    _ = v49
                }
                if (true && (false || v24)) {
                    var v50 string = "\n\t"
                    _ = v50
                } else {
                    v11 += v11
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                v51 := GenS6{F0: -1, F1: int32(fi25), F2: 65535}
                _ = v51
            }
            _ = fi25
        }
        _ = v24
    }
    if (42 >= 2147483647) {
        for fi52 := 0; fi52 < 18; fi52++ {
            for fi53 := 0; fi53 < 15; fi53++ {
                if (-759 < 0) {
                    switch fi53 {
                    case 313:
                        if (0 == 77) {
                            var v54 float32 = 0.1
                            if (v11 >= 255) {
                                v2 = v11
                                fmt.Println(v54, fi53)
                            }
                            switch ("\n\t" + "a") {
                            case "a":
                                var v55 uint = (v8 / ((v8) | 1))
                                _ = v55
                            default:
                                v56 := GenS2{F0: -49}
                                _ = v56
                            }
                            _ = v54
                        }
                    case 2147483647:
                        if ((596 == 32767) && false) {
                            _ = len(v1)
                        }
                    case 127:
                        v2 = v11
                    default:
                        var v57 int8 = 42
                        _ = v57
                    }
                    var v58 uint32 = 74
                    _ = v58
                } else {
                    switch fi53 {
                    case 42:
                        v11 += 127
                    case 32767:
                        v2 += 90
                    default:
                        v11 += (v11 / ((v11) | 1))
                    }
                    if (float32(float64(1e10)) == float32(1e10)) {
                        if true {
                            var v59 float32 = -2.25
                            v11 = (v11 - v11)
                            var v60 int16 = -732
                            _ = v59
                            _ = v60
                        }
                    } else {
                        v61 := GenS3{F0: 0, F1: -1, F2: (-(fi52))}
                        var v62 uint16 = 0
                        var v63 float64 = float64(1e10)
                        fmt.Println(v2)
                        _ = v61
                        _ = v62
                        _ = v63
                    }
                }
                _ = fi53
            }
            fmt.Println(v11, v2)
            _ = fi52
        }
    } else {
        v64 := GenS6{F0: -1, F1: -2147483648, F2: 527}
        v11 += v2
        switch "foo bar" {
        case "foo bar":
            switch 1 {
            case 9223372036854775807:
                v2 += 251
            case -9223372036854775808:
                v2 = (v11 << (((^(v8))) & 63))
            case 196:
                for fi65 := 0; fi65 < 1; fi65++ {
                    if true {
                        fmt.Println(v8)
                    } else {
                        v2 = (v2 >> (((^(v8))) & 63))
                        var v66 uint16 = 704
                        _ = v66
                    }
                    v2 = v11
                    _ = fi65
                }
            }
        case "\n\t":
            delete(v1, 2)
        case "":
            if false {
                fmt.Println(v2, v8)
            } else {
                v1[744] = false
                var v67 uint8 = (v2 | 1)
                _ = v67
            }
        default:
            if ((2147483647 < 1) && (!(-6 > 356))) {
                _ = len(v1)
                v2 -= (v11 + (v2 << ((v8) & 63)))
                switch (("\n\t" + "foo bar") + ("a" + "foo bar")) {
                case "\n\t":
                    v68 := GenS5{F0: 0}
                    _ = v68
                }
            } else {
                v11 = v11
                switch "foo bar" {
                case "foo bar":
                    if (v8 < v8) {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        for fi69 := 0; fi69 < 11; fi69++ {
                            v70 := GenS2{F0: 0}
                            switch "\n\t" {
                            case "\n\t":
                                v71 := GenS3{F0: -2147483648, F1: 127, F2: (fi69 - fi69)}
                                _ = v71
                            }
                            _ = fi69
                            _ = v70
                        }
                        var v72 int32 = -84
                        _ = v72
                    }
                case "a":
                    if true {
                        if (!false) {
                            v8 = (v8 & v8)
                            v73 := GenS7{F0: 195, F1: 276}
                            fmt.Println(v2, v2, v11)
                            _ = v73
                        } else {
                            fmt.Println(v8)
                        }
                        v74 := GenS4{F0: ((42 != -589) && ((0.0 < 0.0) && (!false)))}
                        _ = v74
                    }
                }
                if (!(-128 > 127)) {
                    var v75 string = ("foo bar" + "\n\t")
                    var v76 uint16 = 65
                    v3 = append(v3, 688)
                    _ = v75
                    _ = v76
                }
            }
        }
        _ = v64
    }
    var v77 uint32 = 2147483647
    v2 += 42
    var v78 uint16 = 32767
    fmt.Println(v8)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v77
    _ = v78
    _ = v8
}
