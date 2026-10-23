package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 int16 }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 int32; F1 int16; F2 bool }

func main() {
    var v1 int = 32768
    v2 := []uint32{818}
    var v3 uint = 65535
    var v4 float32 = -2.25
    v5 := []string{("foo bar" + "foo bar"), ("" + "hello"), ("foo bar" + ""), ("a" + (("hello" + "\n\t") + "foo bar"))}
    var v6 int = v1
    var v7 int = v1
    v7 += v7
    var v8 uint64 = 994
    if (2 != 2147483648) {
        v4 += ((v4 + (-2.25 / v4)) + 0.1)
    }
    var v9 uint16 = 2
    for fi10 := 0; fi10 < 9; fi10++ {
        if (v4 < ((v4 / float32((0.1 / 3.14159))) + (v4 / v4))) {
            v11 := GenS1{F0: 0, F1: v8, F2: -1}
            var v12 uint8 = 127
            if true {
                v13 := GenS1{F0: 9223372036854775807, F1: v8, F2: 94}
                _ = v13
            }
            _ = v11
            _ = v12
        } else {
            _ = len(v2)
            var v14 int8 = 127
            _ = v14
        }
        _ = fi10
    }
    var v15 bool = (0.0 >= 1e10)
    v16 := GenS1{F0: 0, F1: 231, F2: 32767}
    fmt.Println(v9)
    switch (v6 - v6) {
    case 128:
        switch v6 {
        case -1:
            switch (^(v1)) {
            case 4294967295:
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            case -9223372036854775808:
                v9 += uint16(v8)
            case 516:
                v7 -= -65
            }
        case 458:
            for fi17 := 0; fi17 < 15; fi17++ {
                var v18 float32 = (v4 + 0.0)
                if false {
                    switch "hello" {
                    case "":
                        var v19 int = (fi17 + 127)
                        _ = v19
                    case "foo bar":
                        v7 = fi17
                    }
                    var v20 bool = (!((v4 + v18) >= ((1e10 / 0.1) / v18)))
                    _ = v20
                } else {
                    fmt.Println(v4, fi17, v3)
                    if (v8 > v8) {
                        for fi21 := 0; fi21 < 20; fi21++ {
                            v9 = v9
                            if true {
                                switch ("a" + "hello") {
                                case "foo bar":
                                    v22 := GenS1{F0: -1, F1: (v8 * v8), F2: 1}
                                    _ = v22
                                case "":
                                    fmt.Println(fi21, v4, v18)
                                }
                            } else {
                                v15 = v15
                                v3 = v3
                            }
                            var v23 uint8 = 255
                            _ = fi21
                            _ = v23
                        }
                        v15 = v15
                        v6 = v1
                    }
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
                _ = fi17
                _ = v18
            }
        }
    case 32768:
        v24 := GenS1{F0: -9223372036854775808, F1: (v8 & (v8 % (((^(v8))) | 1))), F2: 1}
        _ = v24
    case -78:
        for fi25 := 0; fi25 < 14; fi25++ {
            v26 := GenS2{F0: 2147483647, F1: 255}
            var v27 float64 = 1e10
            _ = fi25
            _ = v26
            _ = v27
        }
    default:
        v3 -= v3
    }
    if (!v15) {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v1 += 32768
    } else {
        v1 = (v1 + v6)
        for fi28 := 0; fi28 < 6; fi28++ {
            var v29 uint16 = 32767
            v4 -= (3.14159 / float32((1.5 * (-2.25 - (1.5 - -2.25)))))
            _ = fi28
            _ = v29
        }
        v4 -= v4
    }
    var v30 uint32 = 95
    for fi31 := 0; fi31 < 20; fi31++ {
        var v32 uint16 = v9
        _ = fi31
        _ = v32
    }
    if v15 {
        fmt.Println(v6)
    } else {
        if v15 {
            var v33 uint8 = 255
            if v15 {
                if false {
                    switch (v6 >> ((42) & 63)) {
                    case 390:
                        for fi34 := 0; fi34 < 20; fi34++ {
                            v7 = (v6 >> ((v3) & 63))
                            v9 -= (v9 + (^(v9)))
                            v1 -= -157
                            _ = fi34
                        }
                    case -9223372036854775808:
                        var v35 uint16 = (v9 / ((32767) | 1))
                        _ = v35
                    }
                    v15 = v15
                    var v36 string = "\n\t"
                    _ = v36
                } else {
                    var v37 int16 = -1
                    v37 -= 948
                    _ = v37
                }
                var v38 float32 = v4
                _ = v38
            } else {
                v15 = v15
                for fi39 := 0; fi39 < 19; fi39++ {
                    var v40 bool = v15
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = fi39
                    _ = v40
                }
            }
            _ = v33
        } else {
            var v41 uint32 = (^(v30))
            if false {
                var v42 int64 = -1
                for fi43 := 0; fi43 < 12; fi43++ {
                    fmt.Println(v41, v42, v7)
                    if len(v2) > 0 {
                        v2[0] = v30
                    }
                    _ = fi43
                }
                _ = v42
            } else {
                var v44 float64 = float64(v4)
                var v45 float32 = 1e10
                v46 := GenS2{F0: -677, F1: 128}
                _ = v44
                _ = v45
                _ = v46
            }
            _ = v41
        }
        if v15 {
            fmt.Println(v3, v3, v4)
            var v47 bool = ((!(!false)) && v15)
            _ = v47
        } else {
            switch ("a" + "") {
            case "a":
                v1 = v1
            default:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            v15 = v15
        }
        var v48 uint = 0
        _ = v48
    }
    v49 := GenS2{F0: 42, F1: 116}
    switch (^(v6)) {
    case -911:
        switch (^(v7)) {
        case -518:
            if true {
                fmt.Println(v30, v15)
                if v15 {
                    for fi50 := 0; fi50 < 13; fi50++ {
                        var v51 float32 = (((float32(((0.0 * 0.0) / ((3.14159 + (0.1 - -2.25)) / -2.25))) * (0.1 / 1e10)) / 0.1) / 1e10)
                        for fi52 := 0; fi52 < 9; fi52++ {
                            v4 = v4
                            v9 = (v9 | 0)
                            _ = fi52
                        }
                        _ = fi50
                        _ = v51
                    }
                    v53 := GenS2{F0: -1, F1: 255}
                    _ = v53
                } else {
                    switch v7 {
                    case 32767:
                        var v54 bool = v15
                        _ = v54
                    case 2147483647:
                        if v15 {
                            var v55 int = (v7 * v6)
                            for fi56 := 0; fi56 < 8; fi56++ {
                                var v57 float32 = v4
                                _ = fi56
                                _ = v57
                            }
                            for fi58 := 0; fi58 < 19; fi58++ {
                                if v15 {
                                    v59 := GenS2{F0: 42, F1: 185}
                                    for fi60 := 0; fi60 < 4; fi60++ {
                                        if len(v2) > 0 {
                                            v2[0] = v30
                                        }
                                        _ = fi60
                                    }
                                    fmt.Println(v4, fi58)
                                    _ = v59
                                }
                                _ = fi58
                            }
                            _ = v55
                        } else {
                            for fi61 := 0; fi61 < 4; fi61++ {
                                v15 = v15
                                var v62 bool = (!false)
                                _ = fi61
                                _ = v62
                            }
                        }
                    case 0:
                        var v63 float64 = (-2.25 / float64(v4))
                        _ = v63
                    }
                    v5 = append(v5, "")
                    v64 := GenS2{F0: 1, F1: 127}
                    _ = v64
                }
                var v65 uint = v3
                _ = v65
            }
        case -2147483648:
            switch v6 {
            case 691:
                var v66 uint64 = (v8 / ((uint64(v30)) | 1))
                _ = v66
            case -641:
                for fi67 := 0; fi67 < 7; fi67++ {
                    if v15 {
                        var v68 float32 = float32(1.5)
                        _ = v68
                    } else {
                        var v69 int8 = -128
                        var v70 int32 = -2147483648
                        for fi71 := 0; fi71 < 10; fi71++ {
                            if len(v2) > 0 {
                                v2[0] = uint32(v9)
                            }
                            v72 := GenS3{F0: v70, F1: 835, F2: ((v9 + (^(v9))) <= v9)}
                            _ = fi71
                            _ = v72
                        }
                        _ = v69
                        _ = v70
                    }
                    fmt.Println(v4, v4, v7)
                    v2 = append(v2, (v30 / ((v30) | 1)))
                    _ = fi67
                }
            case 2147483647:
                if (650 == int64(v9)) {
                    if len(v2) > 0 {
                        v2[0] = (v30 / (((v30 ^ (v30 - v30))) | 1))
                    }
                    switch ("foo bar" + "a") {
                    case "hello":
                        v73 := GenS3{F0: -453, F1: 463, F2: false}
                        _ = v73
                    default:
                        var v74 uint16 = 32767
                        _ = v74
                    }
                } else {
                    fmt.Println(v30, v6)
                    if v15 {
                        fmt.Println(v9, v15)
                    }
                }
            default:
                if v15 {
                    v4 = ((v4 / v4) * -2.25)
                    v1 = (v1 << (((^(v3))) & 63))
                    var v75 string = "hello"
                    _ = v75
                } else {
                    var v76 int16 = 32767
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    for fi77 := 0; fi77 < 3; fi77++ {
                        switch "hello" {
                        case "":
                            for fi78 := 0; fi78 < 5; fi78++ {
                                var v79 uint8 = uint8(v30)
                                v76 = 327
                                v8 = v8
                                _ = fi78
                                _ = v79
                            }
                        case "\n\t":
                            switch v6 {
                            case 634:
                                v7 = -632
                            default:
                                var v80 int64 = -1
                                _ = v80
                            }
                        }
                        _ = fi77
                    }
                    _ = v76
                }
            }
        }
    case 127:
        v3 = v3
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v49
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
