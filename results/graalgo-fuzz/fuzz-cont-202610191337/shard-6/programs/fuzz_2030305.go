package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 string; F1 int32 }
type GenS4 struct { F0 int16; F1 uint64; F2 uint64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint8; F1 bool; F2 int8 }
type GenS7 struct { F0 uint16; F1 uint16; F2 uint32 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 int8; F1 int16; F2 uint64 }
type GenS10 struct { F0 int16; F1 uint64 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]uint32)
    v3 := make(map[int]string)
    var v4 uint64 = 553
    v5 := make(map[int]bool)
    var v6 float64 = 1e10
    if (!("" > ("hello" + ("\n\t" + "foo bar")))) {
        var v7 uint8 = 240
        _ = v7
    } else {
        v8 := []int{32767, 127, -1, 9223372036854775807}
        switch -710 {
        case 797:
            v9 := GenS1{F0: -1}
            _ = v9
        case -9223372036854775808:
            for fi10 := 0; fi10 < 6; fi10++ {
                if ((0 != 4294967295) && true) {
                    for fi11 := 0; fi11 < 10; fi11++ {
                        v12 := GenS1{F0: -1}
                        _ = fi11
                        _ = v12
                    }
                }
                _ = fi10
            }
        default:
            var v13 bool = true
            _ = v13
        }
        var v14 uint32 = 781
        _ = v14
        _ = v8
    }
    var v15 int16 = 1000
    switch 577 {
    case 42:
        var v16 int = 9223372036854775807
        _ = v16
    }
    if true {
        switch "foo bar" {
        case "foo bar":
            var v17 uint8 = 105
            _ = v17
        case "":
            var v18 int32 = 0
            _ = v18
        }
        if false {
            fmt.Println(v6, v6, v4)
            v19 := []int{}
            var v20 int8 = 100
            _ = v19
            _ = v20
        } else {
            v21 := GenS1{F0: 83}
            if true {
                for fi22 := 0; fi22 < 8; fi22++ {
                    var v23 uint16 = 103
                    v15 -= v15
                    var v24 uint64 = uint64(v23)
                    _ = fi22
                    _ = v23
                    _ = v24
                }
            } else {
                v4 = 2
                v25 := v3[2147483647]
                fmt.Println(v15, v15, v15)
                _ = v25
            }
            v6 = (float64(1.5) * (v6 + v6))
            _ = v21
        }
        v6 = (v6 - v6)
    } else {
        v26 := GenS2{F0: 42, F1: 7}
        v27 := GenS2{F0: 425, F1: 176}
        fmt.Println(v15, v15)
        _ = v26
        _ = v27
    }
    v28 := GenS2{F0: 139, F1: 0}
    fmt.Println(v6, v15, v15)
    v29 := GenS2{F0: 958, F1: 2}
    v30 := []int{}
    v31 := GenS3{F0: "\n\t", F1: 58}
    switch -2147483648 {
    case 600:
        if len(v30) > 0 {
            v30[0] = 0
        }
    case 1:
        v15 = v15
    case 9223372036854775807:
        if false {
            v32 := GenS1{F0: -2147483648}
            v33 := GenS4{F0: (v15 + 42), F1: v4, F2: (^(v4))}
            if (false && ((v4 * 18446744073709551615) != 4294967295)) {
                v4 -= (v4 % (((v4 | (v4 ^ v4))) | 1))
                v34 := GenS5{F0: true}
                v6 = (0.0 * v6)
                _ = v34
            }
            _ = v32
            _ = v33
        } else {
            fmt.Println(v4, v6, v6)
            var v35 uint64 = (v4 | (v4 / ((v4) | 1)))
            _ = v35
        }
    default:
        fmt.Println(v4, v15, v4)
    }
    for fi36 := 0; fi36 < 19; fi36++ {
        for fi37 := 0; fi37 < 2; fi37++ {
            var v38 float32 = 0.0
            _ = len(v30)
            _ = fi37
            _ = v38
        }
        _ = fi36
    }
    for fi39 := 0; fi39 < 7; fi39++ {
        if (!(-2.25 <= ((0.0 / 3.14159) + float32((float64(1e10) / 1e10))))) {
            fmt.Println(v4, v15, fi39)
            v6 += (float64(0.0) / (((1.5 - -2.25) + (v6 + ((v6 * v6) / v6))) - (v6 - v6)))
            var v40 uint = 42
            _ = v40
        }
        _ = fi39
    }
    for fi41 := 0; fi41 < 13; fi41++ {
        var v42 uint8 = 1
        if ((-2.25 * -2.25) < (1.5 * 1.5)) {
            for fi43 := 0; fi43 < 9; fi43++ {
                var v44 int16 = int16(v4)
                v44 += (v15 << ((1) & 63))
                v2["a"] = 2147483648
                var v45 int64 = 42
                _ = fi43
                _ = v44
                _ = v45
            }
            fmt.Println(v4)
            v6 -= 0.1
        } else {
            for fi46 := 0; fi46 < 4; fi46++ {
                var v47 int16 = v15
                if (119 > int8(fi41)) {
                    fmt.Println(v47, v47)
                    v48 := GenS6{F0: v42, F1: (!false), F2: 100}
                    _ = len(v3)
                    var v49 int = (fi46 * (fi41 % (((fi46 >> ((255) & 63))) | 1)))
                    v4 = v4
                    _ = v48
                    _ = v49
                } else {
                    if len(v30) >= 2 {
                        v30 = v30[1:len(v30)]
                    }
                    var v50 string = (("a" + "") + ("" + "\n\t"))
                    if true {
                        v42 -= v42
                    } else {
                        var v51 string = (("\n\t" + "\n\t") + ("a" + (v50 + v50)))
                        _ = v51
                    }
                    if len(v30) > 0 {
                        v30[0] = fi41
                    }
                    _ = v50
                }
                fmt.Println(v47, v15, v42)
                switch fi41 {
                case 399:
                    v52 := GenS7{F0: 400, F1: 582, F2: 65535}
                    if true {
                        var v53 int16 = 42
                        v54 := GenS8{F0: (^(v4))}
                        _ = v53
                        _ = v54
                    } else {
                        fmt.Println(v6)
                        if (255 == 215) {
                            var v55 float64 = float64(-2.25)
                            _ = v55
                        } else {
                            v56 := GenS9{F0: 0, F1: 42, F2: (v4 - v4)}
                            switch (fi46 ^ 42) {
                            case 755:
                                v57 := GenS5{F0: (false && (!(!(0 < 0))))}
                                _ = v57
                            case 1:
                                v47 -= v47
                            case -1:
                                v47 = -37
                            default:
                                var v58 int64 = -1
                                _ = v58
                            }
                            v59 := GenS4{F0: (-(v47)), F1: uint64(v42), F2: (v4 * (v4 << ((18446744073709551615) & 63)))}
                            var v60 float64 = (v6 / (0.0 - (float64(1e10) + v6)))
                            if false {
                                var v61 int = (fi46 + 9223372036854775807)
                                v62 := v1[(v61 & 128)]
                                _ = v61
                                _ = v62
                            } else {
                                v63 := v5[(fi41 / (((fi41 << ((255) & 63))) | 1))]
                                _ = v63
                            }
                            v60 = (float64(0.0) / v6)
                            v60 += (((0.0 - float64(1e10)) - (v60 / 1e10)) / -2.25)
                            _ = len(v30)
                            _ = v56
                            _ = v59
                            _ = v60
                        }
                        fmt.Println(v4, fi46, fi41)
                    }
                    fmt.Println(fi41, v47, v6)
                    _ = v52
                default:
                    var v64 uint16 = 377
                    var v65 string = "\n\t"
                    v64 += v64
                    _ = v64
                    _ = v65
                }
                v66 := GenS7{F0: 65535, F1: 42, F2: 2147483647}
                v42 = 140
                _ = fi46
                _ = v47
                _ = v66
            }
            if len(v30) >= 2 {
                v30 = v30[1:len(v30)]
            }
            var v67 uint64 = (^(v4))
            _ = v67
        }
        if false {
            var v68 uint = 55
            _ = v68
        } else {
            fmt.Println(v15)
            v4 = (v4 - v4)
        }
        _ = fi41
        _ = v42
    }
    v69 := GenS10{F0: v15, F1: (v4 / (((^(v4))) | 1))}
    v6 = v6
    v4 = v4
    if true {
        var v70 uint16 = 32767
        _ = v70
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v69
}
