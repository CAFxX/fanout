package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 uint16 }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint64; F1 int32 }
type GenS6 struct { F0 uint8; F1 int64; F2 bool }
type GenS7 struct { F0 int64; F1 int32; F2 int64 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint64; F1 uint32 }
type GenS10 struct { F0 int; F1 string; F2 int32 }
func genFn0(p0 uint8, p1 int64) uint8 {
    var v5 int16 = -32768
    if (true || (42 <= 0)) {
        for fi6 := 0; fi6 < 2; fi6++ {
            var v7 int32 = -1
            v8 := GenS1{F0: 382, F1: "\n\t", F2: 255}
            v5 = v5
            _ = fi6
            _ = v7
            _ = v8
        }
        v9 := GenS1{F0: 1, F1: "hello", F2: 427}
        _ = v9
    }
    v10 := []int{2147483647, -747}
    if true {
        for fi11 := 0; fi11 < 17; fi11++ {
            v5 -= (v5 + 32767)
            if len(v10) > 0 {
                v10[0] = fi11
            }
            _ = fi11
        }
        for fi12 := 0; fi12 < 15; fi12++ {
            v13 := GenS2{F0: 975, F1: 2}
            v14 := GenS3{F0: "foo bar"}
            _ = fi12
            _ = v13
            _ = v14
        }
        var v15 uint16 = 42
        _ = v15
    }
    p1 += (p1 / ((p1) | 1))
    _ = p0
    _ = p1
    _ = v10
    _ = v5
    return 255
}

func main() {
    var v1 uint8 = 1
    var v2 string = "foo bar"
    var v3 uint16 = 1
    v4 := []uint32{2147483647, 2147483647, 65535}
    if (-9223372036854775808 <= 9223372036854775807) {
        fmt.Println(v2)
    } else {
        if (!true) {
            if true {
                switch 4294967295 {
                case 9223372036854775807:
                    for fi16 := 0; fi16 < 5; fi16++ {
                        if len(v4) > 0 {
                            v4[0] = 65535
                        }
                        var v17 int8 = 100
                        v1 -= v1
                        _ = fi16
                        _ = v17
                    }
                case 2147483647:
                    v18 := GenS3{F0: v2}
                    _ = v18
                case 127:
                    _ = len(v4)
                default:
                    if false {
                        v19 := GenS4{F0: -1}
                        for fi20 := 0; fi20 < 6; fi20++ {
                            v1 = (v1 >> ((588) & 63))
                            var v21 string = v2
                            _ = fi20
                            _ = v21
                        }
                        _ = v19
                    } else {
                        var v22 int8 = -78
                        _ = v22
                    }
                }
                for fi23 := 0; fi23 < 14; fi23++ {
                    for fi24 := 0; fi24 < 3; fi24++ {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        v1 += (v1 % (((^(v1))) | 1))
                        if ((^(fi23)) != (-(fi24))) {
                            var v25 int = fi23
                            _ = v25
                        } else {
                            if (v1 == (v1 | (^(v1)))) {
                                var v26 string = ""
                                v1 -= (v1 ^ v1)
                                _ = v26
                            }
                            v27 := GenS2{F0: 4294967295, F1: (^(v3))}
                            v2 = v2
                            var v28 float32 = (((3.14159 - 1e10) - ((float32(1e10) * (0.1 / float32(0.1))) / ((0.0 + 0.0) + 1e10))) - 1.5)
                            var v29 uint = 255
                            _ = v27
                            _ = v28
                            _ = v29
                        }
                        _ = fi24
                    }
                    v1 = (v1 >> ((776) & 63))
                    if true {
                        var v30 int8 = 127
                        var v31 int = 9223372036854775807
                        _ = v30
                        _ = v31
                    }
                    _ = fi23
                }
            } else {
                for fi32 := 0; fi32 < 10; fi32++ {
                    if false {
                        var v33 int = fi32
                        if (-478 <= 1000) {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            v2 = "foo bar"
                            if len(v4) > 0 {
                                v4[0] = 2147483647
                            }
                        } else {
                            fmt.Println(v33)
                            var v34 float32 = 0.0
                            _ = v34
                        }
                        _ = v33
                    }
                    var v35 uint32 = 772
                    _ = fi32
                    _ = v35
                }
                v36 := GenS2{F0: 2147483648, F1: (v3 << ((545) & 63))}
                _ = v36
            }
            if (false && (false || (1 < 969))) {
                var v37 uint16 = v3
                _ = v37
            } else {
                switch -2147483648 {
                case 42:
                    v2 = "\n\t"
                case 9223372036854775807:
                    v3 = v3
                case -832:
                    if false {
                        var v38 int64 = 1
                        v39 := GenS5{F0: 2, F1: 1}
                        _ = len(v4)
                        _ = v38
                        _ = v39
                    } else {
                        for fi40 := 0; fi40 < 19; fi40++ {
                            v41 := GenS1{F0: 1, F1: v2, F2: v3}
                            v42 := GenS6{F0: 2, F1: 1, F2: (32767 > v3)}
                            _ = fi40
                            _ = v41
                            _ = v42
                        }
                        for fi43 := 0; fi43 < 3; fi43++ {
                            for fi44 := 0; fi44 < 19; fi44++ {
                                v45 := GenS4{F0: -284}
                                _ = fi44
                                _ = v45
                            }
                            _ = fi43
                        }
                        v2 = v2
                    }
                }
                v46 := GenS2{F0: 2147483648, F1: v3}
                _ = v46
            }
        } else {
            v47 := GenS4{F0: 0}
            _ = v47
        }
        fmt.Println(v3, v3, v2)
        fmt.Println(v2, v3, v1)
    }
    if (false || (3.14159 >= -2.25)) {
        if (float64((1e10 * (3.14159 + (float32((0.0 / 0.1)) - 3.14159)))) < -2.25) {
            if (!(-769 <= 2147483647)) {
                switch 128 {
                case 9223372036854775807:
                    var v48 uint32 = 2147483647
                    _ = v48
                case 32768:
                    var v49 int64 = -506
                    _ = v49
                case 0:
                    switch (v2 + v2) {
                    case "":
                        v2 = (v2 + v2)
                    case "\n\t":
                        switch (v2 + v2) {
                        case "hello":
                            var v50 float64 = ((0.0 + 1e10) * float64(0.1))
                            _ = v50
                        case "\n\t":
                            var v51 uint8 = v1
                            _ = v51
                        }
                    default:
                        for fi52 := 0; fi52 < 20; fi52++ {
                            if len(v4) > 0 {
                                v4[0] = 1
                            }
                            v3 += (^(v3))
                            _ = fi52
                        }
                    }
                }
                var v53 int64 = 867
                v54 := GenS7{F0: -864, F1: 103, F2: v53}
                _ = v53
                _ = v54
            }
            switch (v2 + v2) {
            case "hello":
                fmt.Println(v3)
            case "\n\t":
                if ((2147483648 < 4294967295) || (0 >= 42)) {
                    var v55 int8 = 100
                    if len(v4) > 0 {
                        v4[0] = 4294967295
                    }
                    var v56 string = "hello"
                    _ = v55
                    _ = v56
                }
            }
            switch 32767 {
            case 128:
                for fi57 := 0; fi57 < 5; fi57++ {
                    v2 = "hello"
                    v58 := GenS8{F0: 111}
                    _ = fi57
                    _ = v58
                }
            }
        } else {
            for fi59 := 0; fi59 < 13; fi59++ {
                var v60 uint8 = v1
                var v61 int = (fi59 % (((fi59 * -334)) | 1))
                _ = fi59
                _ = v60
                _ = v61
            }
            fmt.Println(v1)
            v1 = (^(v1))
        }
    }
    fmt.Println(v2, v1, v2)
    for fi62 := 0; fi62 < 12; fi62++ {
        var v63 uint8 = 255
        _ = fi62
        _ = v63
    }
    for fi64 := 0; fi64 < 20; fi64++ {
        fmt.Println(fi64)
        if true {
            switch v2 {
            case "hello":
                if true {
                    fmt.Println(v3, fi64)
                    fmt.Println(v3)
                    v2 = v2
                } else {
                    v1 -= (^(v1))
                    _ = len(v4)
                }
            case "foo bar":
                if (!false) {
                    fmt.Println(fi64, fi64)
                }
            case "":
                v65 := GenS7{F0: 104, F1: -659, F2: 1}
                _ = v65
            }
            var v66 uint64 = 2
            _ = v66
        }
        _ = fi64
    }
    if (!(42 < -9223372036854775808)) {
        fmt.Println(v1, v3)
    }
    v2 = "hello"
    var v67 bool = (true || true)
    if false {
        var v68 uint64 = 1
        if len(v4) > 0 {
            v4[0] = 65535
        }
        _ = v68
    }
    if (!(0 > 206)) {
        fmt.Println(v3, v67)
        v67 = (v67 || v67)
        v1 -= v1
    } else {
        fmt.Println(v67, v1, v2)
        var v69 float64 = -2.25
        _ = v69
    }
    if false {
        for fi70 := 0; fi70 < 1; fi70++ {
            for fi71 := 0; fi71 < 19; fi71++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi71
            }
            _ = fi70
        }
        var v72 int = 1
        for fi73 := 0; fi73 < 14; fi73++ {
            var v74 int = (fi73 + v72)
            if (!(!((v1 << ((2) & 63)) > 0))) {
                if len(v4) > 0 {
                    v4[0] = 42
                }
                for fi75 := 0; fi75 < 4; fi75++ {
                    var v76 float64 = 1e10
                    _ = fi75
                    _ = v76
                }
            }
            _ = fi73
            _ = v74
        }
        _ = v72
    } else {
        fmt.Println(v3, v67, v67)
    }
    var v77 uint16 = v3
    if (v67 || v67) {
        v78 := GenS9{F0: 0, F1: 42}
        v77 += 42
        _ = v78
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    v3 = (^(v77))
    v67 = (((float32(-2.25) * 1.5) + ((((float32((1e10 - (1.5 / 0.1))) / (-2.25 / ((-2.25 + (0.1 - (-2.25 * 3.14159))) + 1.5))) - -2.25) / ((3.14159 / (1.5 * (0.1 - 1.5))) - (0.0 - 0.1))) + (0.0 - 1.5))) > 1e10)
    v79 := GenS10{F0: -793, F1: "\n\t", F2: -2147483648}
    fmt.Println(v2, v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v67:", v67)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v67
    _ = v77
    _ = v79
}
