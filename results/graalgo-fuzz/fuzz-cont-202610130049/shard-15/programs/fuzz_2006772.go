package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint8; F1 uint16 }
type GenS4 struct { F0 int; F1 uint32; F2 int8 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int64; F1 int; F2 int32 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint8; F1 uint8; F2 int16 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := []uint32{}
    v3 := []string{"foo bar"}
    var v4 float64 = 3.14159
    v5 := v1[622]
    switch 0 {
    case 32768:
        v5 = true
    case -985:
        v6 := GenS1{F0: 0}
        _ = v6
    case -25:
        for fi7 := 0; fi7 < 6; fi7++ {
            var v8 uint8 = 255
            if false {
                for fi9 := 0; fi9 < 9; fi9++ {
                    if len(v2) > 0 {
                        v2[0] = 2147483648
                    }
                    if false {
                        v4 -= (v4 / 1.0)
                        v4 = (v4 - v4)
                    } else {
                        if (42 != 486) {
                            fmt.Println(v4, v5)
                            var v10 int8 = 0
                            _ = v10
                        } else {
                            var v11 int32 = 2147483647
                            v12 := GenS1{F0: v8}
                            _ = v11
                            _ = v12
                        }
                        v5 = (v5 && (0 < 199))
                        v8 -= 128
                    }
                    v8 += v8
                    _ = fi9
                }
                v13 := v1[329]
                _ = len(v1)
                _ = v13
            }
            fmt.Println(v4, v5, fi7)
            _ = fi7
            _ = v8
        }
    default:
        v5 = v5
    }
    v5 = v5
    if (v5 || v5) {
        var v14 int8 = -101
        _ = v14
    } else {
        if false {
            switch ("foo bar" + "a") {
            case "a":
                v4 += (v4 + v4)
            case "":
                switch 42 {
                case -745:
                    var v15 uint8 = 0
                    _ = v15
                case -923:
                    v5 = (!((0 > -642) && (!v5)))
                case -575:
                    for fi16 := 0; fi16 < 17; fi16++ {
                        fmt.Println(v5)
                        v17 := GenS2{F0: 876}
                        for fi18 := 0; fi18 < 5; fi18++ {
                            var v19 int16 = -1
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            v4 -= 1e10
                            v4 += v4
                            _ = fi18
                            _ = v19
                        }
                        _ = fi16
                        _ = v17
                    }
                }
            }
            v5 = (false || v5)
        } else {
            for fi20 := 0; fi20 < 14; fi20++ {
                v4 = v4
                _ = fi20
            }
        }
        var v21 string = ("hello" + (("a" + "\n\t") + "foo bar"))
        delete(v1, 2)
        _ = v21
    }
    v22 := GenS1{F0: 42}
    if v5 {
        v5 = v5
        v4 += 3.14159
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v23 := GenS1{F0: 128}
        v1[2] = ((v5 || (v5 && (33 != 0))) || v5)
        _ = v23
    }
    switch -836 {
    case -2147483648:
        v4 = (v4 - v4)
    case 615:
        if ("\n\t" == "\n\t") {
            v5 = v5
        } else {
            var v24 uint64 = 1
            if true {
                var v25 int16 = 204
                _ = v25
            } else {
                switch "" {
                case "foo bar":
                    if true {
                        var v26 float64 = ((1e10 * float64(float32((v4 + 3.14159)))) + v4)
                        _ = v26
                    }
                default:
                    v27 := GenS1{F0: 71}
                    _ = v27
                }
                var v28 uint32 = 2147483647
                if (float32(v4) == 0.1) {
                    var v29 int16 = 42
                    for fi30 := 0; fi30 < 13; fi30++ {
                        v31 := GenS1{F0: 195}
                        v32 := GenS2{F0: 446}
                        _ = fi30
                        _ = v31
                        _ = v32
                    }
                    v28 = 303
                    _ = v29
                } else {
                    if (-32768 < 1000) {
                        v33 := GenS3{F0: 128, F1: 65535}
                        _ = v33
                    } else {
                        var v34 float64 = v4
                        v35 := GenS2{F0: 0}
                        if true {
                            var v36 int32 = 0
                            for fi37 := 0; fi37 < 6; fi37++ {
                                fmt.Println(v34, v36)
                                var v38 int8 = 30
                                var v39 string = ("" + "a")
                                _ = fi37
                                _ = v38
                                _ = v39
                            }
                            _ = v36
                        }
                        _ = v34
                        _ = v35
                    }
                    v40 := v1[68]
                    v40 = v40
                    _ = v40
                }
                _ = v28
            }
            if (!(!true)) {
                var v41 float64 = 0.0
                if v5 {
                    switch "a" {
                    case "hello":
                        v42 := GenS4{F0: -676, F1: 2, F2: -89}
                        _ = v42
                    }
                    for fi43 := 0; fi43 < 19; fi43++ {
                        v44 := GenS5{F0: 32767}
                        if (!v5) {
                            for fi45 := 0; fi45 < 16; fi45++ {
                                fmt.Println(fi45, fi45, v4)
                                _ = len(v2)
                                _ = fi45
                            }
                            v46 := GenS6{F0: 727, F1: (fi43 / (((fi43 & fi43)) | 1)), F2: 1}
                            _ = v46
                        }
                        v24 -= v24
                        _ = fi43
                        _ = v44
                    }
                    fmt.Println(v41, v41)
                } else {
                    var v47 float32 = float32(v41)
                    var v48 uint16 = 2
                    var v49 int8 = -55
                    _ = v47
                    _ = v48
                    _ = v49
                }
                var v50 bool = (80 != 255)
                _ = v41
                _ = v50
            } else {
                v3 = append(v3, "a")
            }
            _ = v24
        }
    case 823:
        var v51 uint32 = 2
        _ = v51
    }
    for fi52 := 0; fi52 < 1; fi52++ {
        for fi53 := 0; fi53 < 19; fi53++ {
            var v54 float32 = 3.14159
            _ = fi53
            _ = v54
        }
        v55 := GenS3{F0: 0, F1: 32767}
        _ = fi52
        _ = v55
    }
    if len(v3) > 0 {
        v3[0] = "foo bar"
    }
    for fi56 := 0; fi56 < 10; fi56++ {
        var v57 int8 = 42
        v1[42] = v5
        _ = fi56
        _ = v57
    }
    v4 += float64(-2.25)
    var v58 int = 32768
    switch "" {
    case "\n\t":
        var v59 uint8 = 71
        _ = v59
    case "a":
        var v60 uint8 = 255
        _ = v60
    default:
        var v61 float64 = ((float64(((1e10 / 1.0) * (1e10 + (3.14159 - (float32(float64(1.5)) * 3.14159))))) - 0.1) - v4)
        _ = v61
    }
    if (!(-234 == -237)) {
        var v62 float64 = v4
        if (((905 <= 32767) || ((2147483647 <= -672) || true)) || ((!false) && (65535 < 668))) {
            for fi63 := 0; fi63 < 20; fi63++ {
                switch "a" {
                case "a":
                    switch -489 {
                    case -443:
                        v62 -= v4
                    case 2147483647:
                        v64 := GenS7{F0: 1}
                        _ = v64
                    case 127:
                        fmt.Println(fi63, v4)
                    }
                }
                v65 := GenS8{F0: 2, F1: 128, F2: -1}
                _ = fi63
                _ = v65
            }
        } else {
            switch (^(v58)) {
            case 2147483647:
                var v66 float64 = 0.0
                _ = v66
            case 973:
                v62 += (v4 + v62)
            case 9223372036854775807:
                v67 := GenS7{F0: 40}
                _ = v67
            }
            v4 -= -2.25
        }
        fmt.Println(v4, v5, v58)
        _ = v62
    } else {
        for fi68 := 0; fi68 < 10; fi68++ {
            v4 -= ((0.1 - v4) + 0.0)
            var v69 uint16 = 42
            _ = fi68
            _ = v69
        }
    }
    var v70 string = ""
    for fi71 := 0; fi71 < 12; fi71++ {
        v58 += (v58 << ((686) & 63))
        v5 = (65535 == 4294967295)
        var v72 int32 = 2147483647
        _ = fi71
        _ = v72
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v70
}
