package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint32; F1 uint; F2 uint64 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint64; F1 int32; F2 uint }
type GenS8 struct { F0 uint8; F1 uint8; F2 uint64 }
type GenS9 struct { F0 int; F1 uint8 }
func genFn0(p0 int8, p1 int) int {
    var v6 uint32 = 2147483648
    p1 -= (p1 % (((-(p1))) | 1))
    v7 := GenS1{F0: true}
    v6 = (^(v6))
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 42
}

func main() {
    v1 := []uint32{65535, 2, 317}
    var v2 uint8 = 255
    v3 := []uint32{858}
    var v4 string = ("hello" + "")
    var v5 int = 32767
    switch v4 {
    case "\n\t":
        v2 += v2
    case "":
        var v8 uint = 940
        _ = v8
    case "foo bar":
        for fi9 := 0; fi9 < 18; fi9++ {
            fmt.Println(v5)
            switch v4 {
            case "hello":
                var v10 int64 = int64(v2)
                _ = v10
            default:
                fmt.Println(v4, v5)
            }
            for fi11 := 0; fi11 < 4; fi11++ {
                v12 := GenS1{F0: false}
                _ = fi11
                _ = v12
            }
            _ = fi9
        }
    }
    for fi13 := 0; fi13 < 4; fi13++ {
        var v14 string = "\n\t"
        _ = fi13
        _ = v14
    }
    var v15 uint8 = 127
    v16 := GenS1{F0: (true && false)}
    fmt.Println(v2, v15)
    if false {
        for fi17 := 0; fi17 < 16; fi17++ {
            for fi18 := 0; fi18 < 6; fi18++ {
                fmt.Println(fi18)
                if true {
                    var v19 float32 = -2.25
                    v2 += (v15 >> ((4294967295) & 63))
                    switch fi17 {
                    case 127:
                        if (-1 == -88) {
                            if true {
                                if false {
                                    switch v4 {
                                    case "foo bar":
                                        v2 = v15
                                    case "a":
                                        fmt.Println(v4, v5, v2)
                                    case "":
                                        v20 := GenS2{F0: 2}
                                        _ = v20
                                    default:
                                        if (9223372036854775807 == (fi17 / ((32767) | 1))) {
                                            v5 = fi18
                                        }
                                    }
                                } else {
                                    v21 := GenS2{F0: 2}
                                    switch v5 {
                                    case -598:
                                        var v22 int = fi17
                                        _ = v22
                                    case 388:
                                        var v23 uint32 = 2
                                        _ = v23
                                    case 0:
                                        fmt.Println(v2, v5, v5)
                                    default:
                                        var v24 float64 = (1e10 - (0.1 - 1e10))
                                        _ = v24
                                    }
                                    v1 = append(v1, 2147483647)
                                    var v25 float64 = 3.14159
                                    _ = v21
                                    _ = v25
                                }
                                v19 = (v19 * (v19 + float32(float64((1.5 * v19)))))
                            } else {
                                v2 += 0
                            }
                        } else {
                            v26 := GenS2{F0: 0}
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = v26
                        }
                    case 128:
                        var v27 float64 = 3.14159
                        _ = v27
                    default:
                        v28 := GenS3{F0: -32768}
                        _ = v28
                    }
                    _ = v19
                } else {
                    v29 := GenS1{F0: (true || false)}
                    switch ((("a" + (v4 + v4)) + "foo bar") + (v4 + v4)) {
                    case "\n\t":
                        v5 = (fi17 / ((442) | 1))
                    case "a":
                        var v30 int8 = -35
                        _ = v30
                    default:
                        if (!true) {
                            _ = len(v3)
                            switch ((v4 + ((v4 + "foo bar") + v4)) + "foo bar") {
                            case "foo bar":
                                if (!true) {
                                    var v31 float32 = (0.1 / 3.14159)
                                    if false {
                                        v32 := GenS2{F0: 32767}
                                        var v33 int8 = 100
                                        v15 -= 0
                                        _ = v32
                                        _ = v33
                                    }
                                    _ = v31
                                }
                                v2 = v2
                            case "a":
                                switch (fi18 << ((619) & 63)) {
                                case -731:
                                    switch (fi17 / (((v5 & (v5 + (fi18 % ((-595) | 1))))) | 1)) {
                                    case 0:
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    }
                                case 0:
                                    switch -1 {
                                    case 578:
                                        var v34 uint8 = 1
                                        _ = v34
                                    case -1:
                                        fmt.Println(fi17)
                                    default:
                                        var v35 uint32 = 65535
                                        _ = v35
                                    }
                                }
                            case "hello":
                                switch (fi18 | genFn0(87, (fi17 - genFn0(-127, (fi18 * v5))))) {
                                case 540:
                                    v15 = (v2 << ((65535) & 63))
                                case -2147483648:
                                    v2 -= v15
                                default:
                                    v5 -= fi17
                                }
                            default:
                                var v36 float64 = (0.1 * 0.0)
                                _ = v36
                            }
                            switch fi18 {
                            case 2147483647:
                                v37 := GenS2{F0: 255}
                                _ = v37
                            case 422:
                                if false {
                                    var v38 uint16 = 53
                                    _ = v38
                                } else {
                                    v39 := GenS4{F0: 1}
                                    v1 = append(v1, 706)
                                    if (65535 == 89) {
                                        if (289 == -1) {
                                            _ = len(v1)
                                            switch v4 {
                                            case "\n\t":
                                                var v40 float32 = float32(3.14159)
                                                _ = v40
                                            case "hello":
                                                var v41 int32 = -176
                                                _ = v41
                                            case "":
                                                v42 := GenS5{F0: 42, F1: 0, F2: 2}
                                                _ = v42
                                            }
                                        }
                                        var v43 int16 = 974
                                        var v44 float64 = ((1.5 * (1e10 - 3.14159)) - 1.5)
                                        v4 = ""
                                        fmt.Println(v4, v15, fi18)
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        _ = v43
                                        _ = v44
                                    }
                                    _ = v39
                                }
                            default:
                                v45 := GenS1{F0: true}
                                var v46 int8 = 1
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                _ = v45
                                _ = v46
                            }
                        }
                    }
                    _ = len(v1)
                    _ = v29
                }
                _ = fi18
            }
            for fi47 := 0; fi47 < 10; fi47++ {
                var v48 int32 = -771
                if (255 >= 65535) {
                    v3 = append(v3, 296)
                } else {
                    v49 := GenS4{F0: 130}
                    _ = v49
                }
                _ = fi47
                _ = v48
            }
            _ = fi17
        }
    } else {
        if false {
            v5 -= int(v15)
            fmt.Println(v4, v15)
        }
        v15 = v2
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v4, v2, v4)
    for fi50 := 0; fi50 < 10; fi50++ {
        switch (v5 >> ((4294967295) & 63)) {
        case -244:
            v2 = (v15 % ((1) | 1))
        case 128:
            v51 := GenS6{F0: 0}
            _ = v51
        }
        v52 := GenS6{F0: 0}
        v2 += (v2 * 255)
        _ = fi50
        _ = v52
    }
    if (false || (0 <= -32768)) {
        if len(v1) > 0 {
            v1[0] = 0
        }
        for fi53 := 0; fi53 < 4; fi53++ {
            switch (v4 + v4) {
            case "foo bar":
                switch "a" {
                case "foo bar":
                    v54 := GenS7{F0: 42, F1: -1, F2: 4294967295}
                    _ = v54
                }
            default:
                switch v4 {
                case "hello":
                    if len(v3) > 0 {
                        v3[0] = 0
                    }
                case "foo bar":
                    v15 = v15
                case "\n\t":
                    for fi55 := 0; fi55 < 9; fi55++ {
                        var v56 float32 = 0.1
                        _ = fi55
                        _ = v56
                    }
                default:
                    v15 -= (v2 - v2)
                }
            }
            _ = fi53
        }
        fmt.Println(v2, v5, v4)
    } else {
        var v57 float32 = (1e10 / 1.0)
        v2 = (^(v15))
        _ = len(v3)
        _ = v57
    }
    switch v5 {
    case -704:
        v58 := GenS8{F0: v2, F1: 0, F2: 2}
        _ = v58
    }
    for fi59 := 0; fi59 < 6; fi59++ {
        v5 += (fi59 << ((0) & 63))
        _ = fi59
    }
    var v60 int = (v5 | (v5 >> ((4294967295) & 63)))
    switch v5 {
    case 128:
        for fi61 := 0; fi61 < 17; fi61++ {
            for fi62 := 0; fi62 < 11; fi62++ {
                var v63 float32 = 1.5
                if (9223372036854775807 <= 721) {
                    var v64 float32 = v63
                    _ = v64
                } else {
                    if (0 > 715) {
                        v65 := GenS1{F0: false}
                        switch 42 {
                        case 871:
                            fmt.Println(fi62, v15, v2)
                        case 127:
                            fmt.Println(v5, v2, v2)
                        case -676:
                            if len(v3) > 0 {
                                v3[0] = 0
                            }
                        }
                        v2 += v2
                        _ = v65
                    } else {
                        v66 := GenS8{F0: 127, F1: (v15 ^ (v15 % (((v15 / ((250) | 1))) | 1))), F2: 921}
                        var v67 uint32 = uint32(v2)
                        _ = len(v1)
                        _ = v66
                        _ = v67
                    }
                    var v68 uint16 = 1
                    _ = v68
                }
                _ = fi62
                _ = v63
            }
            fmt.Println(v60, v5)
            _ = fi61
        }
    }
    switch ("a" + v4) {
    case "foo bar":
        if (65535 <= 0) {
            switch 2147483647 {
            case 32767:
                if false {
                    var v69 int = (v60 + v5)
                    var v70 bool = (-595 != 42)
                    if v70 {
                        v71 := GenS9{F0: v60, F1: v15}
                        _ = v71
                    }
                    _ = v69
                    _ = v70
                }
            }
            if false {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                var v72 string = "foo bar"
                _ = v72
            }
        } else {
            var v73 int16 = -1
            if (((v4 + "") + ("hello" + v4)) == "hello") {
                if len(v3) > 0 {
                    v3[0] = 2147483647
                }
                var v74 uint = 1
                _ = v74
            }
            switch (((v4 + ("foo bar" + (v4 + v4))) + v4) + (v4 + v4)) {
            case "a":
                var v75 int32 = -2147483648
                _ = v75
            case "\n\t":
                var v76 float32 = (1e10 * -2.25)
                _ = v76
            default:
                var v77 int32 = 2147483647
                _ = v77
            }
            _ = v73
        }
    case "":
        switch v4 {
        case "\n\t":
            for fi78 := 0; fi78 < 12; fi78++ {
                var v79 float32 = (float32(0.1) + (0.1 / float32(1.5)))
                var v80 int16 = 374
                _ = fi78
                _ = v79
                _ = v80
            }
        default:
            _ = len(v1)
        }
    case "hello":
        var v81 uint8 = (v15 ^ (^(v15)))
        _ = v81
    default:
        if (1 < 881) {
            var v82 uint64 = 42
            if true {
                var v83 float64 = ((((-2.25 * ((1.5 - 1e10) * 3.14159)) + ((((-2.25 - 0.0) * 1.5) + float64(1e10)) / 3.14159)) * 0.1) - float64(0.1))
                v4 = (v4 + (v4 + v4))
                _ = v83
            } else {
                var v84 uint16 = 255
                _ = v84
            }
            _ = v82
        } else {
            v60 = (-(v5))
        }
    }
    _ = len(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v60
}
