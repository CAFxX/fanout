package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 uint8; F1 int16; F2 int64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint32; F1 int64; F2 int }
type GenS5 struct { F0 int32; F1 uint8; F2 string }
type GenS6 struct { F0 uint32; F1 string }
type GenS7 struct { F0 uint16; F1 int16; F2 uint32 }
type GenS8 struct { F0 int32; F1 int16 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 uint16; F1 int16; F2 bool }
type GenS13 struct { F0 uint32; F1 uint32; F2 uint64 }
type GenS14 struct { F0 int32 }
type GenS15 struct { F0 int8; F1 string; F2 uint16 }
type GenS16 struct { F0 bool; F1 bool; F2 bool }

func main() {
    v1 := []int{}
    var v2 int = 127
    var v3 float64 = 1e10
    var v4 uint64 = 4294967295
    var v5 uint16 = 42
    var v6 float64 = float64((1.5 / 1.0))
    if (true && (-389 > 362)) {
        v2 -= v2
    } else {
        v5 = 65535
        var v7 int32 = -1
        v5 += (v5 << ((190) & 63))
        _ = v7
    }
    var v8 float32 = (3.14159 * 1.5)
    v8 += 1e10
    v9 := GenS1{F0: 81, F1: 0}
    v2 -= 42
    if false {
        var v10 uint32 = 0
        if len(v1) > 0 {
            v1[0] = (v2 * (v2 - v2))
        }
        _ = v10
    }
    v11 := GenS1{F0: 2, F1: (^(v4))}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch 32767 {
    case 4294967295:
        v12 := GenS1{F0: (v5 % (((^(v5))) | 1)), F1: (^(v4))}
        _ = v12
    case -1:
        v3 += v3
    }
    for fi13 := 0; fi13 < 19; fi13++ {
        var v14 int16 = -1
        if (-54 != 42) {
            for fi15 := 0; fi15 < 17; fi15++ {
                var v16 int8 = 14
                switch (v2 | (fi13 - fi15)) {
                case 0:
                    fmt.Println(v3)
                case 127:
                    if ((false || (true && true)) && true) {
                        if (true || false) {
                            v16 -= v16
                            v17 := GenS2{F0: 42, F1: v14, F2: int64(v16)}
                            _ = len(v1)
                            _ = v17
                        }
                        if ((0 != -1) && (!("\n\t" >= "hello"))) {
                            v6 = 0.0
                            v18 := GenS3{F0: (v16 + v16)}
                            v2 -= -2147483648
                            _ = v18
                        } else {
                            if false {
                                v19 := GenS4{F0: 88, F1: -1, F2: fi15}
                                v5 = v5
                                if (uint32(v4) <= 545) {
                                    var v20 uint16 = uint16(v4)
                                    v14 = (v14 * 987)
                                    v21 := GenS4{F0: 0, F1: 42, F2: (fi13 | (fi15 + fi13))}
                                    _ = v20
                                    _ = v21
                                }
                                _ = v19
                            } else {
                                var v22 int = fi15
                                if true {
                                    switch (fi13 & fi13) {
                                    case 127:
                                        fmt.Println(v6, v4)
                                    case 32768:
                                        v22 -= fi15
                                    case 0:
                                        var v23 bool = true
                                        _ = v23
                                    default:
                                        v24 := GenS5{F0: 2147483647, F1: 187, F2: ""}
                                        _ = v24
                                    }
                                } else {
                                    var v25 uint32 = 4294967295
                                    v26 := GenS6{F0: v25, F1: "a"}
                                    v27 := GenS7{F0: (v5 * uint16(v25)), F1: v14, F2: (v25 / (((^(v25))) | 1))}
                                    if (!((^(fi13)) < fi13)) {
                                        fmt.Println(v3, fi13)
                                        var v28 int16 = (^(v14))
                                        if (uint(v25) >= 1) {
                                            var v29 int8 = 127
                                            _ = v29
                                        } else {
                                            v3 -= v3
                                            var v30 bool = false
                                            if v30 {
                                                var v31 int = v22
                                                switch v2 {
                                                case -716:
                                                    switch -9223372036854775808 {
                                                    case -9223372036854775808:
                                                        v1 = append(v1, -9223372036854775808)
                                                    }
                                                default:
                                                    if v30 {
                                                        switch "foo bar" {
                                                        case "":
                                                            v16 -= (v16 / ((v16) | 1))
                                                        default:
                                                            var v32 uint16 = v5
                                                            _ = v32
                                                        }
                                                        v14 = (v14 / ((v28) | 1))
                                                        switch 32767 {
                                                        case 32768:
                                                            v3 += v3
                                                            _ = len(v1)
                                                        case 758:
                                                            if v30 {
                                                                if (v30 || ((v16 | v16) >= 42)) {
                                                                    v6 = float64(v8)
                                                                    var v33 float64 = v6
                                                                    v6 += (float64(-2.25) - v33)
                                                                    _ = v33
                                                                } else {
                                                                    v6 += v6
                                                                }
                                                                v5 = v5
                                                            } else {
                                                                if (!v30) {
                                                                    fmt.Println(v16)
                                                                } else {
                                                                    v34 := GenS6{F0: v25, F1: "a"}
                                                                    _ = v34
                                                                }
                                                                if ((v25 | v25) >= v25) {
                                                                    var v35 uint64 = (v4 ^ v4)
                                                                    if len(v1) > 0 {
                                                                        v1[0] = v22
                                                                    }
                                                                    _ = v35
                                                                } else {
                                                                    if v30 {
                                                                        if false {
                                                                            var v36 uint = 255
                                                                            v37 := GenS7{F0: uint16(v36), F1: -1, F2: v25}
                                                                            if len(v1) > 0 {
                                                                                v1[0] = (v31 % ((v2) | 1))
                                                                            }
                                                                            _ = v36
                                                                            _ = v37
                                                                        }
                                                                        var v38 uint16 = uint16(v25)
                                                                        v22 += (v31 - fi15)
                                                                        v5 = v5
                                                                        v25 += (v25 >> ((65535) & 63))
                                                                        _ = v38
                                                                    }
                                                                    var v39 float64 = 0.1
                                                                    _ = v39
                                                                }
                                                                fmt.Println(v4)
                                                                fmt.Println(v31, v25)
                                                                v5 -= (v5 / ((v5) | 1))
                                                            }
                                                        case 32767:
                                                            v14 += v28
                                                        default:
                                                            v40 := GenS8{F0: 303, F1: (v28 + (v14 & (v14 * (^(v14)))))}
                                                            _ = v40
                                                        }
                                                        var v41 bool = v30
                                                        v42 := GenS9{F0: 128}
                                                        _ = v41
                                                        _ = v42
                                                    }
                                                }
                                                _ = v31
                                            }
                                            _ = v30
                                        }
                                        _ = v28
                                    }
                                    var v43 int8 = (v16 >> ((4294967295) & 63))
                                    _ = v25
                                    _ = v26
                                    _ = v27
                                    _ = v43
                                }
                                v44 := GenS10{F0: -718}
                                _ = v22
                                _ = v44
                            }
                            _ = len(v1)
                            var v45 int = 32767
                            if len(v1) > 0 {
                                v1[0] = (fi13 >> ((65535) & 63))
                            }
                            if false {
                                switch -909 {
                                case 128:
                                    var v46 int8 = v16
                                    _ = v46
                                }
                            }
                            _ = v45
                        }
                    } else {
                        if false {
                            v47 := GenS6{F0: 1, F1: (("" + "hello") + ("hello" + "hello"))}
                            var v48 uint64 = 42
                            _ = v47
                            _ = v48
                        } else {
                            var v49 uint64 = (v4 / ((v4) | 1))
                            var v50 float32 = v8
                            _ = v49
                            _ = v50
                        }
                        switch (fi13 + v2) {
                        case -1:
                            v51 := GenS3{F0: (-(v16))}
                            _ = v51
                        case 46:
                            v4 -= (v4 | uint64(v5))
                        case 127:
                            v5 -= v5
                        }
                    }
                case -172:
                    var v52 uint32 = 0
                    _ = v52
                }
                v3 += v3
                _ = fi15
                _ = v16
            }
            if (2147483648 < 1) {
                switch v2 {
                case 2147483647:
                    _ = len(v1)
                default:
                    if true {
                        if (!false) {
                            v3 = (-2.25 * v3)
                            switch (fi13 & 127) {
                            case -1:
                                if true {
                                    switch (-(fi13)) {
                                    case -9223372036854775808:
                                        v3 += (float64(-2.25) * ((float64(v8) * v3) + ((float64((v8 * (((float32(v3) + 1e10) * (v8 + v8)) / (v8 * ((((v8 * float32(v3)) * 3.14159) + (v8 * v8)) + v8))))) / float64(((((v8 * ((float32(((3.14159 - float64((v8 - v8))) + (v6 - ((v6 * 0.1) + 3.14159)))) + v8) * (0.0 / v8))) - v8) / (0.0 / v8)) * (v8 - 0.0)))) / v6)))
                                    default:
                                        for fi53 := 0; fi53 < 18; fi53++ {
                                            v2 -= (-(v2))
                                            _ = fi53
                                        }
                                    }
                                    fmt.Println(fi13)
                                } else {
                                    var v54 uint64 = v4
                                    if ((v8 - 0.0) >= v8) {
                                        v6 += 1.5
                                        v55 := GenS11{F0: 0}
                                        _ = v55
                                    } else {
                                        v5 = (v5 - v5)
                                        v56 := GenS2{F0: 76, F1: (v14 | -32768), F2: 379}
                                        _ = v56
                                    }
                                    var v57 int16 = (v14 * (-(v14)))
                                    _ = v54
                                    _ = v57
                                }
                            }
                            for fi58 := 0; fi58 < 4; fi58++ {
                                _ = len(v1)
                                v59 := GenS3{F0: -9}
                                _ = fi58
                                _ = v59
                            }
                        } else {
                            for fi60 := 0; fi60 < 15; fi60++ {
                                var v61 string = ("hello" + ("a" + ""))
                                if (2 > 2) {
                                    v62 := GenS12{F0: v5, F1: v14, F2: (((!true) || (487 == 536)) || true)}
                                    v3 = v3
                                    if false {
                                        v63 := GenS1{F0: v5, F1: (^(v4))}
                                        if false {
                                            if len(v1) > 0 {
                                                v1[0] = (fi60 % ((32767) | 1))
                                            }
                                        } else {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            fmt.Println(v3, fi13)
                                        }
                                        switch "a" {
                                        case "\n\t":
                                            _ = len(v1)
                                        }
                                        _ = v63
                                    }
                                    _ = v62
                                }
                                _ = fi60
                                _ = v61
                            }
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                        switch 4294967295 {
                        case 1:
                            v64 := GenS13{F0: 4294967295, F1: 955, F2: (v4 | 42)}
                            _ = v64
                        }
                        if (53 != 127) {
                            if len(v1) > 0 {
                                v1[0] = 2147483647
                            }
                            for fi65 := 0; fi65 < 16; fi65++ {
                                _ = len(v1)
                                var v66 uint32 = 42
                                v67 := GenS14{F0: 521}
                                _ = fi65
                                _ = v66
                                _ = v67
                            }
                            var v68 int8 = 100
                            _ = v68
                        }
                    }
                }
            }
        } else {
            v3 -= (((1e10 - float64((-2.25 - v8))) / 3.14159) / v3)
            for fi69 := 0; fi69 < 18; fi69++ {
                v3 = v3
                switch "foo bar" {
                case "hello":
                    v3 = (3.14159 * -2.25)
                case "foo bar":
                    switch fi13 {
                    case 895:
                        v4 += (v4 ^ 0)
                    case 1:
                        var v70 int = 128
                        _ = v70
                    case 4294967295:
                        switch (("hello" + ("\n\t" + "\n\t")) + "a") {
                        case "a":
                            if len(v1) > 0 {
                                v1[0] = fi69
                            }
                            if len(v1) > 0 {
                                v1[0] = 128
                            }
                        case "":
                            v71 := GenS15{F0: int8(v5), F1: "", F2: v5}
                            v4 += (v4 % ((387) | 1))
                            v2 = (fi13 / ((2147483647) | 1))
                            _ = v71
                        case "\n\t":
                            v72 := GenS15{F0: 0, F1: "\n\t", F2: v5}
                            _ = v72
                        }
                    }
                }
                _ = fi69
            }
        }
        v73 := GenS4{F0: 42, F1: 390, F2: (fi13 + v2)}
        _ = fi13
        _ = v14
        _ = v73
    }
    fmt.Println(v2, v6, v4)
    v74 := GenS16{F0: false, F1: false, F2: false}
    v4 = v4
    fmt.Println(v8)
    fmt.Println(v3, v3)
    if ((100 > 127) && (!((65535 < 255) && ((!true) && (v2 != v2))))) {
        v75 := GenS13{F0: 550, F1: 42, F2: (v4 % ((v4) | 1))}
        _ = v75
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v74
    _ = v8
    _ = v9
}
