package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int8; F1 uint32 }
type GenS5 struct { F0 int32; F1 uint16 }

func main() {
    var v1 bool = false
    var v2 int32 = 429
    var v3 uint8 = 83
    for fi4 := 0; fi4 < 19; fi4++ {
        switch fi4 {
        case 958:
            v2 = 0
        case -2147483648:
            for fi5 := 0; fi5 < 19; fi5++ {
                v2 += 42
                _ = fi5
            }
        case 755:
            v2 = v2
        }
        if true {
            v3 = 0
        } else {
            v6 := []int{fi4, -664, (-(fi4)), (fi4 * fi4)}
            _ = v6
        }
        fmt.Println(v2, v1, fi4)
        _ = fi4
    }
    if (103 < 65535) {
        if (-1 > 4294967295) {
            if (((3.14159 - (((float64(0.0) + 3.14159) * 0.1) + 1.5)) <= 0.1) && (0 != -91)) {
                if (1 > 2147483647) {
                    v7 := []int{int(v3)}
                    v2 += v2
                    _ = v7
                } else {
                    if (!v1) {
                        if (v1 || v1) {
                            v1 = (!v1)
                        }
                        for fi8 := 0; fi8 < 16; fi8++ {
                            var v9 uint64 = 2
                            switch "a" {
                            case "":
                                v10 := []int{fi8, 9223372036854775807, (fi8 & fi8), fi8}
                                _ = v10
                            }
                            var v11 uint = 255
                            _ = fi8
                            _ = v11
                            _ = v9
                        }
                    }
                    switch -1 {
                    case 32768:
                        var v12 int64 = 9223372036854775807
                        _ = v12
                    case 9223372036854775807:
                        switch "" {
                        case "hello":
                            v13 := []int{220, -1, -1, 128}
                            _ = v13
                        }
                    }
                }
                if (!v1) {
                    fmt.Println(v2, v3, v2)
                }
                var v14 float32 = ((float32((((float64(1.5) * float64((1e10 * -2.25))) - (0.1 - (1.5 / ((0.0 - -2.25) * 0.1)))) / 1e10)) - 0.1) / 1.5)
                _ = v14
            }
        } else {
            if (1 >= 4294967295) {
                if v1 {
                    if (!(0.0 != 1e10)) {
                        var v15 int64 = 9223372036854775807
                        _ = v15
                    } else {
                        v16 := []int{32768, 2147483647, int(v2)}
                        switch 0 {
                        case -369:
                            if (-1 == 42) {
                                v17 := GenS1{F0: 200, F1: (427 <= 374), F2: 0}
                                v18 := GenS2{F0: -841}
                                _ = v17
                                _ = v18
                            } else {
                                switch 128 {
                                case 9223372036854775807:
                                    var v19 int16 = -32768
                                    _ = v19
                                default:
                                    var v20 int8 = 65
                                    _ = v20
                                }
                                v1 = v1
                            }
                        case 192:
                            var v21 uint = 1
                            _ = v21
                        default:
                            var v22 uint32 = 1
                            _ = v22
                        }
                        _ = v16
                    }
                    switch "a" {
                    case "\n\t":
                        v2 -= 0
                    case "a":
                        if (!(18446744073709551615 >= 677)) {
                            v1 = v1
                            if v1 {
                                var v23 int32 = (v2 + v2)
                                v24 := GenS3{F0: 1}
                                v25 := GenS3{F0: 2}
                                _ = v23
                                _ = v24
                                _ = v25
                            } else {
                                v26 := []int{127, 534, 128}
                                v27 := GenS3{F0: 2147483647}
                                switch 127 {
                                case 0:
                                    if len(v26) > 0 {
                                        v26[0] = 32768
                                    }
                                case 32767:
                                    var v28 uint64 = uint64(v3)
                                    _ = v28
                                default:
                                    var v29 int = 127
                                    _ = v29
                                }
                                _ = v26
                                _ = v27
                            }
                        } else {
                            switch int(v2) {
                            case 0:
                                var v30 uint16 = 42
                                _ = v30
                            case 2147483647:
                                switch "hello" {
                                case "\n\t":
                                    if v1 {
                                        v31 := GenS1{F0: 1, F1: v1, F2: 40}
                                        fmt.Println(v2, v3, v3)
                                        _ = v31
                                    } else {
                                        var v32 float32 = (0.0 / (0.1 - (1e10 * 0.0)))
                                        _ = v32
                                    }
                                case "foo bar":
                                    v2 -= (v2 << ((42) & 63))
                                default:
                                    v33 := GenS4{F0: 127, F1: 830}
                                    _ = v33
                                }
                            case 42:
                                v34 := GenS3{F0: 42}
                                _ = v34
                            }
                            v2 += v2
                            var v35 bool = (0 > 1)
                            _ = v35
                        }
                    }
                }
            }
            switch ("hello" + "") {
            case "foo bar":
                v36 := GenS3{F0: 2147483648}
                _ = v36
            }
            v1 = (550 == 1)
        }
    }
    for fi37 := 0; fi37 < 11; fi37++ {
        if v1 {
            switch (("foo bar" + ("\n\t" + "")) + "hello") {
            case "foo bar":
                switch fi37 {
                case 24:
                    v38 := []int{fi37, (fi37 / ((int(v2)) | 1)), fi37}
                    _ = v38
                case 127:
                    switch fi37 {
                    case 977:
                        var v39 uint = 255
                        _ = v39
                    }
                case 9223372036854775807:
                    var v40 uint8 = (v3 << ((2) & 63))
                    _ = v40
                default:
                    v41 := GenS1{F0: 2147483648, F1: v1, F2: 42}
                    _ = v41
                }
            case "":
                v3 += (v3 + 255)
            default:
                v42 := GenS4{F0: 76, F1: 1}
                _ = v42
            }
            var v43 uint32 = 2147483647
            _ = v43
        }
        var v44 int32 = (-(v2))
        v3 -= (v3 / ((v3) | 1))
        _ = fi37
        _ = v44
    }
    if v1 {
        v2 -= (v2 & 1)
        for fi45 := 0; fi45 < 7; fi45++ {
            v46 := []int{}
            switch "foo bar" {
            case "a":
                for fi47 := 0; fi47 < 11; fi47++ {
                    var v48 int32 = v2
                    switch fi45 {
                    case 412:
                        var v49 uint64 = 355
                        _ = v49
                    }
                    _ = fi47
                    _ = v48
                }
            case "hello":
                for fi50 := 0; fi50 < 8; fi50++ {
                    var v51 int = (fi50 * (fi50 & 0))
                    if false {
                        v52 := GenS5{F0: (v2 - 1), F1: 65535}
                        _ = v52
                    } else {
                        var v53 uint8 = v3
                        v1 = true
                        var v54 int = 4294967295
                        _ = v53
                        _ = v54
                    }
                    _ = fi50
                    _ = v51
                }
            default:
                v55 := GenS1{F0: 2147483647, F1: true, F2: -1}
                _ = v55
            }
            if (-119 > 127) {
                if v1 {
                    fmt.Println(v2, v3, fi45)
                }
                var v56 bool = true
                if false {
                    var v57 int8 = -1
                    _ = v57
                } else {
                    v1 = v1
                    var v58 float64 = (((((3.14159 * (0.0 + ((3.14159 + 1e10) * float64(1e10)))) / (3.14159 / (-2.25 - 0.1))) / 1e10) + ((float64(((0.0 / 3.14159) * ((0.1 + float32(0.0)) * (((1e10 + 0.1) / 0.1) + 0.0)))) - -2.25) - float64(3.14159))) + 0.0)
                    v1 = v56
                    _ = v58
                }
                _ = v56
            } else {
                v2 = (v2 / ((-926) | 1))
                if v1 {
                    if ((!v1) && v1) {
                        v59 := GenS1{F0: 1, F1: v1, F2: 1000}
                        _ = v59
                    }
                }
            }
            _ = fi45
            _ = v46
        }
        var v60 int32 = v2
        _ = v60
    }
    var v61 string = ("hello" + "a")
    var v62 float64 = 0.0
    v63 := []int{int(v2)}
    var v64 uint32 = 2147483648
    _ = len(v63)
    switch ((v61 + "a") + (("foo bar" + (v61 + v61)) + "")) {
    case "a":
        var v65 float32 = float32(float64(float32(v62)))
        _ = v65
    case "foo bar":
        v61 = "hello"
    }
    var v66 uint64 = 0
    var v67 string = (v61 + (v61 + (v61 + (v61 + (v61 + v61)))))
    var v68 uint16 = 844
    var v69 int8 = 127
    var v70 int8 = (v69 ^ -99)
    if true {
        switch ("" + v67) {
        case "foo bar":
            fmt.Println(v62, v64)
        }
        var v71 uint = 336
        _ = v71
    } else {
        var v72 float64 = v62
        var v73 float64 = -2.25
        _ = v72
        _ = v73
    }
    for fi74 := 0; fi74 < 2; fi74++ {
        if v1 {
            switch v61 {
            case "\n\t":
                switch v67 {
                case "":
                    v63 = append(v63, int(v68))
                case "a":
                    v61 = ""
                default:
                    if true {
                        var v75 string = v67
                        if true {
                            for fi76 := 0; fi76 < 15; fi76++ {
                                if false {
                                    var v77 uint32 = 497
                                    if v1 {
                                        if len(v63) >= 2 {
                                            v63 = v63[1:len(v63)]
                                        }
                                        var v78 bool = (1e10 >= float32(3.14159))
                                        if (v78 && v78) {
                                            switch v61 {
                                            case "":
                                                var v79 float64 = v62
                                                _ = v79
                                            }
                                            var v80 int8 = v70
                                            _ = v80
                                        }
                                        _ = v78
                                    }
                                    _ = v77
                                } else {
                                    fmt.Println(v61, fi74)
                                }
                                _ = fi76
                            }
                        }
                        _ = len(v63)
                        _ = v75
                    }
                }
            }
            fmt.Println(v62, fi74, v1)
        }
        v70 = 0
        _ = fi74
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    fmt.Println("slicelen v63:", len(v63), v63)
    fmt.Println("v64:", v64)
    fmt.Println("v66:", v66)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v3
    _ = v61
    _ = v62
    _ = v63
    _ = v64
    _ = v66
    _ = v67
    _ = v68
    _ = v69
    _ = v70
}
