package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int8; F1 bool; F2 uint16 }
type GenS4 struct { F0 int32; F1 int8 }
type GenS5 struct { F0 int32; F1 uint8 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int8; F1 uint8; F2 int64 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 int32; F1 int16 }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 uint8; F1 bool }
type GenS14 struct { F0 uint8 }

func main() {
    v1 := []string{}
    var v2 int16 = -1
    v3 := []int{1}
    v4 := GenS1{F0: 927, F1: -9223372036854775808, F2: 65535}
    v2 += 1000
    if len(v1) > 0 {
        v1[0] = ("hello" + ("a" + ((("hello" + "") + "a") + ("a" + "a"))))
    }
    if true {
        if true {
            switch "foo bar" {
            case "":
                if (("\n\t" >= "\n\t") && (-2147483648 <= 1)) {
                    if ((!true) && (0 > 32767)) {
                        for fi5 := 0; fi5 < 7; fi5++ {
                            v6 := GenS2{F0: 267}
                            _ = fi5
                            _ = v6
                        }
                        for fi7 := 0; fi7 < 16; fi7++ {
                            for fi8 := 0; fi8 < 19; fi8++ {
                                v3 = append(v3, fi8)
                                var v9 uint = 255
                                fmt.Println(fi7, v9)
                                fmt.Println(fi8, fi8)
                                var v10 uint8 = 255
                                if (true || true) {
                                    v9 = (^(v9))
                                } else {
                                    v3 = append(v3, fi8)
                                }
                                if false {
                                    if (2147483647 <= -1) {
                                        v9 += v9
                                    } else {
                                        fmt.Println(v10)
                                    }
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                } else {
                                    v2 += 998
                                    if true {
                                        v9 += 885
                                        if ((!(!true)) || false) {
                                            v11 := GenS1{F0: 34, F1: (fi7 - fi8), F2: 0}
                                            _ = v11
                                        } else {
                                            fmt.Println(fi8)
                                        }
                                        if (-1 >= -886) {
                                            if len(v1) > 0 {
                                                v1[0] = "foo bar"
                                            }
                                            v3 = append(v3, (fi7 >> ((v9) & 63)))
                                        }
                                    }
                                }
                                fmt.Println(v9)
                                _ = fi8
                                _ = v10
                                _ = v9
                            }
                            v2 += (v2 & v2)
                            _ = fi7
                        }
                    }
                } else {
                    if (!true) {
                        v12 := GenS1{F0: 599, F1: 327, F2: 255}
                        for fi13 := 0; fi13 < 9; fi13++ {
                            var v14 float32 = 1.5
                            _ = fi13
                            _ = v14
                        }
                        fmt.Println(v2)
                        _ = v12
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    fmt.Println(v2)
                }
            case "\n\t":
                v15 := GenS3{F0: -128, F1: true, F2: 42}
                _ = v15
            default:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            if len(v3) > 0 {
                v3[0] = 870
            }
        } else {
            for fi16 := 0; fi16 < 5; fi16++ {
                v17 := GenS3{F0: 0, F1: (1 <= 42), F2: 32767}
                _ = fi16
                _ = v17
            }
            switch "hello" {
            case "":
                if true {
                    if ((42 > 2147483647) && true) {
                        for fi18 := 0; fi18 < 17; fi18++ {
                            if true {
                                if len(v1) > 0 {
                                    v1[0] = ""
                                }
                            } else {
                                v19 := GenS3{F0: -85, F1: (!(!(!((0 >= 54) || (4294967295 == 65535))))), F2: 642}
                                _ = v19
                            }
                            _ = fi18
                        }
                    } else {
                        var v20 string = "foo bar"
                        for fi21 := 0; fi21 < 6; fi21++ {
                            var v22 bool = (152 > 4294967295)
                            var v23 int32 = 42
                            if (v20 >= v20) {
                                fmt.Println(fi21, fi21, v20)
                                for fi24 := 0; fi24 < 16; fi24++ {
                                    switch 0 {
                                    case -650:
                                        switch (fi21 & -115) {
                                        case 128:
                                            switch (fi24 * 42) {
                                            case 42:
                                                if (151 > 4294967295) {
                                                    if (v2 == v2) {
                                                        if v22 {
                                                            var v25 float32 = float32((3.14159 * 0.1))
                                                            v20 = ("a" + "hello")
                                                            _ = v25
                                                        }
                                                        if (213 >= 0) {
                                                            v22 = v22
                                                            v20 = v20
                                                            v26 := GenS3{F0: 127, F1: v22, F2: 32767}
                                                            switch v20 {
                                                            case "foo bar":
                                                                v23 -= (v23 % (((v23 | 42)) | 1))
                                                                if (-2.25 == ((3.14159 + 3.14159) - (1e10 * 0.1))) {
                                                                    v23 -= (v23 << ((255) & 63))
                                                                    _ = len(v3)
                                                                    v3 = append(v3, (fi21 | fi21))
                                                                } else {
                                                                    var v27 bool = ((v2 % (((v2 % (((v2 & 1)) | 1))) | 1)) <= v2)
                                                                    var v28 uint8 = 1
                                                                    _ = v27
                                                                    _ = v28
                                                                }
                                                            case "":
                                                                v29 := GenS2{F0: 1}
                                                                _ = v29
                                                            case "hello":
                                                                var v30 float64 = (3.14159 + (3.14159 * -2.25))
                                                                _ = v30
                                                            }
                                                            _ = v26
                                                        }
                                                    }
                                                    v31 := GenS1{F0: (v2 * (^(v2))), F1: (fi24 + 396), F2: 324}
                                                    v20 = v20
                                                    v3 = append(v3, fi21)
                                                    var v32 int = (fi21 | (-(fi24)))
                                                    _ = v31
                                                    _ = v32
                                                } else {
                                                    if true {
                                                        _ = len(v3)
                                                        if false {
                                                            v33 := GenS4{F0: 2147483647, F1: 100}
                                                            _ = v33
                                                        }
                                                        v34 := GenS5{F0: (^(v23)), F1: 0}
                                                        _ = v34
                                                    } else {
                                                        var v35 uint32 = 2147483648
                                                        switch (fi21 | fi24) {
                                                        case 9223372036854775807:
                                                            if len(v1) > 0 {
                                                                v1[0] = v20
                                                            }
                                                        case 0:
                                                            fmt.Println(v35, v22, v23)
                                                        case -441:
                                                            var v36 uint = 2
                                                            _ = v36
                                                        }
                                                        var v37 int64 = 9223372036854775807
                                                        _ = v35
                                                        _ = v37
                                                    }
                                                    v38 := GenS6{F0: 858}
                                                    var v39 string = v20
                                                    _ = v38
                                                    _ = v39
                                                }
                                            case 4294967295:
                                                switch fi21 {
                                                case 4294967295:
                                                    switch 699 {
                                                    case 9223372036854775807:
                                                        v40 := GenS7{F0: 71}
                                                        _ = v40
                                                    case 1:
                                                        switch 2147483647 {
                                                        case 4294967295:
                                                            fmt.Println(v22)
                                                            if (1 > 42) {
                                                                v22 = (!false)
                                                            } else {
                                                                var v41 uint16 = 0
                                                                v42 := GenS7{F0: 32767}
                                                                _ = v41
                                                                _ = v42
                                                            }
                                                        }
                                                    default:
                                                        var v43 uint16 = 735
                                                        var v44 uint8 = 255
                                                        var v45 string = (v20 + (v20 + v20))
                                                        v20 = v20
                                                        var v46 float32 = (((1.5 - 3.14159) + (3.14159 * (float32(float64(0.1)) + -2.25))) / float32(0.1))
                                                        _ = v43
                                                        _ = v44
                                                        _ = v45
                                                        _ = v46
                                                    }
                                                case 9223372036854775807:
                                                    v47 := GenS1{F0: (v2 + -1), F1: (fi24 - int(v2)), F2: 42}
                                                    _ = v47
                                                case -1:
                                                    var v48 string = ""
                                                    _ = v48
                                                }
                                            }
                                        default:
                                            v3 = append(v3, fi24)
                                        }
                                    default:
                                        if (1 <= 42) {
                                            if len(v1) > 0 {
                                                v1[0] = (v20 + "foo bar")
                                            }
                                            v20 = (((v20 + v20) + "hello") + "hello")
                                            if v22 {
                                                var v49 uint32 = 4294967295
                                                var v50 int = (fi24 ^ (fi24 / ((fi21) | 1)))
                                                _ = v49
                                                _ = v50
                                            }
                                        }
                                    }
                                    var v51 int = (fi24 ^ fi24)
                                    v2 = 32767
                                    _ = fi24
                                    _ = v51
                                }
                                var v52 int8 = 127
                                _ = v52
                            } else {
                                switch "\n\t" {
                                case "":
                                    v53 := GenS8{F0: 127, F1: 0, F2: -800}
                                    _ = v53
                                case "a":
                                    v54 := GenS9{F0: int16(fi21)}
                                    _ = v54
                                case "hello":
                                    var v55 int64 = -9223372036854775808
                                    _ = v55
                                default:
                                    var v56 uint64 = 134
                                    _ = v56
                                }
                                var v57 string = ((((v20 + "hello") + ("a" + ((v20 + v20) + ((v20 + "a") + v20)))) + v20) + v20)
                                _ = v57
                            }
                            _ = fi21
                            _ = v22
                            _ = v23
                        }
                        _ = v20
                    }
                    fmt.Println(v2)
                    v2 = v2
                } else {
                    v58 := GenS10{F0: 545}
                    _ = v58
                }
            case "hello":
                v2 = (v2 % ((v2) | 1))
            case "a":
                for fi59 := 0; fi59 < 5; fi59++ {
                    _ = len(v3)
                    v2 = (v2 / ((-1) | 1))
                    _ = fi59
                }
            }
        }
        if (255 == 0) {
            v2 -= 42
            v3 = append(v3, 1)
        } else {
            v60 := GenS11{F0: -2147483648, F1: v2}
            var v61 uint = 129
            _ = v60
            _ = v61
        }
        var v62 bool = (255 >= 1)
        _ = v62
    } else {
        v63 := GenS12{F0: 1}
        _ = v63
    }
    if false {
        var v64 uint32 = 2
        for fi65 := 0; fi65 < 7; fi65++ {
            var v66 int64 = int64(v2)
            _ = fi65
            _ = v66
        }
        _ = v64
    }
    v1 = append(v1, "hello")
    if false {
        v2 = (v2 + (v2 / ((v2) | 1)))
        v2 -= v2
    } else {
        v67 := GenS11{F0: 2147483647, F1: v2}
        switch "foo bar" {
        case "foo bar":
            v68 := GenS13{F0: 0, F1: ("\n\t" > "hello")}
            _ = v68
        case "\n\t":
            v2 -= -584
        case "hello":
            fmt.Println(v2)
        }
        _ = v67
    }
    if len(v1) > 0 {
        v1[0] = "a"
    }
    for fi69 := 0; fi69 < 2; fi69++ {
        v2 -= -1
        if len(v3) > 0 {
            v3[0] = fi69
        }
        _ = fi69
    }
    v2 = (v2 | (-(v2)))
    v2 += (v2 | v2)
    v2 += 1000
    if len(v3) > 0 {
        v3[0] = -2147483648
    }
    for fi70 := 0; fi70 < 10; fi70++ {
        for fi71 := 0; fi71 < 15; fi71++ {
            v3 = append(v3, fi70)
            v72 := GenS12{F0: 2}
            _ = fi71
            _ = v72
        }
        if false {
            var v73 int32 = -1
            v74 := GenS14{F0: 135}
            if true {
                v2 = 32767
            }
            _ = v73
            _ = v74
        } else {
            for fi75 := 0; fi75 < 7; fi75++ {
                fmt.Println(v2)
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi75
            }
        }
        _ = fi70
    }
    v3 = append(v3, 191)
    switch "a" {
    case "":
        var v76 int8 = -128
        _ = v76
    case "\n\t":
        var v77 float32 = (-2.25 + float32(0.0))
        _ = v77
    case "hello":
        if ((!true) && (!true)) {
            for fi78 := 0; fi78 < 6; fi78++ {
                fmt.Println(v2, fi78)
                _ = fi78
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v2 += 1
        } else {
            switch "foo bar" {
            case "a":
                var v79 uint16 = 32767
                _ = v79
            case "hello":
                for fi80 := 0; fi80 < 17; fi80++ {
                    if true {
                        v2 -= 32767
                        v2 -= (v2 | v2)
                    }
                    v2 = v2
                    _ = fi80
                }
            default:
                v2 = v2
            }
            fmt.Println(v2)
        }
    default:
        var v81 int = 128
        _ = v81
    }
    v2 = v2
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
