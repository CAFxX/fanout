package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 uint16; F1 string; F2 int16 }
type GenS4 struct { F0 string; F1 uint8; F2 int64 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int }
func genFn0(p0 int64, p1 bool, p2 int) int8 {
    if p1 {
        v7 := GenS1{F0: 1, F1: -108}
        _ = v7
    } else {
        var v8 bool = p1
        _ = v8
    }
    if (1 < 65535) {
        fmt.Println(p2)
    }
    var v9 int8 = 1
    v9 = (v9 << ((18446744073709551615) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return (v9 << ((255) & 63))
}
func genFn1(p0 uint32, p1 uint8, p2 uint) uint32 {
    v10 := []int{595, -847}
    var v11 int32 = -422
    v11 = 2147483647
    var v12 int = 128
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v12
    return 952
}

func main() {
    var v1 bool = true
    v2 := make(map[uint32]int)
    var v3 float32 = 1.5
    var v4 uint64 = 215
    v5 := []int{}
    v6 := make(map[uint32]string)
    for fi13 := 0; fi13 < 6; fi13++ {
        v14 := GenS1{F0: 9, F1: 32767}
        v4 -= v4
        _ = fi13
        _ = v14
    }
    v15 := GenS1{F0: genFn0(9223372036854775807, ((1 >= genFn1(1, 1, 349)) && (v1 || false)), 32768), F1: 498}
    if len(v5) > 0 {
        v5[0] = 9223372036854775807
    }
    if v1 {
        var v16 uint8 = 255
        fmt.Println(v1)
        _ = v16
    }
    v17 := GenS1{F0: 1, F1: 0}
    var v18 uint64 = v4
    if v1 {
        var v19 uint = 255
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        for fi20 := 0; fi20 < 6; fi20++ {
            v21 := GenS2{F0: 0, F1: 802}
            _ = fi20
            _ = v21
        }
        _ = v19
    } else {
        if (!(v1 || v1)) {
            var v22 string = ("hello" + "hello")
            v4 -= (v4 + 946)
            switch "a" {
            case "":
                var v23 float32 = v3
                _ = v23
            case "hello":
                switch (v22 + v22) {
                case "\n\t":
                    if (116 < 0) {
                        if len(v5) > 0 {
                            v5[0] = 9223372036854775807
                        }
                    } else {
                        if (255 >= 127) {
                            for fi24 := 0; fi24 < 2; fi24++ {
                                switch fi24 {
                                case -709:
                                    var v25 float32 = ((v3 * -2.25) - v3)
                                    _ = v25
                                default:
                                    var v26 bool = false
                                    _ = v26
                                }
                                _ = fi24
                            }
                            switch (v22 + (v22 + v22)) {
                            case "a":
                                delete(v2, 366)
                            default:
                                if v1 {
                                    for fi27 := 0; fi27 < 5; fi27++ {
                                        switch 22 {
                                        case 32767:
                                            v1 = v1
                                        case 2147483647:
                                            v28 := GenS3{F0: 1, F1: "\n\t", F2: 32767}
                                            _ = v28
                                        case 0:
                                            v1 = v1
                                        }
                                        var v29 float64 = float64((((((-2.25 - (v3 - v3)) + (v3 * 1e10)) * v3) + v3) - float32(1.5)))
                                        _ = fi27
                                        _ = v29
                                    }
                                    var v30 int16 = 42
                                    _ = v30
                                } else {
                                    var v31 int = 42
                                    _ = v31
                                }
                            }
                        }
                        var v32 int32 = -2147483648
                        v33 := GenS4{F0: v22, F1: 0, F2: 42}
                        _ = v32
                        _ = v33
                    }
                case "":
                    if v1 {
                        switch (v22 + ((v22 + (("a" + v22) + ("hello" + v22))) + "a")) {
                        case "foo bar":
                            v34 := GenS5{F0: -928}
                            _ = v34
                        case "hello":
                            for fi35 := 0; fi35 < 4; fi35++ {
                                if true {
                                    for fi36 := 0; fi36 < 16; fi36++ {
                                        var v37 string = v22
                                        v38 := GenS6{F0: v1}
                                        var v39 uint16 = 58
                                        _ = fi36
                                        _ = v37
                                        _ = v38
                                        _ = v39
                                    }
                                    var v40 uint = 15
                                    _ = v40
                                } else {
                                    if true {
                                        for fi41 := 0; fi41 < 1; fi41++ {
                                            var v42 int = (fi35 * (fi35 | (fi41 - int(v4))))
                                            _ = fi41
                                            _ = v42
                                        }
                                        v18 = 1
                                        var v43 bool = ((872 == 0) && v1)
                                        _ = v43
                                    } else {
                                        v5 = append(v5, (fi35 & fi35))
                                    }
                                    for fi44 := 0; fi44 < 5; fi44++ {
                                        v22 = (v22 + v22)
                                        if v1 {
                                            v18 = v4
                                            var v45 int16 = 30
                                            _ = v45
                                        } else {
                                            v22 = v22
                                            v18 = (v4 & 924)
                                            fmt.Println(fi35, v4)
                                        }
                                        _ = len(v5)
                                        _ = fi44
                                    }
                                    fmt.Println(v1, fi35, v4)
                                }
                                for fi46 := 0; fi46 < 10; fi46++ {
                                    if ((!(((-657 >= -1) && true) && (!(4294967295 <= 1)))) && (1 < 766)) {
                                        v6[0] = (v22 + (v22 + "hello"))
                                        if (((v4 % ((v18) | 1)) >= 597) || false) {
                                            var v47 uint16 = 32767
                                            if (fi35 < fi46) {
                                                fmt.Println(fi46, v22)
                                            } else {
                                                var v48 int16 = -417
                                                _ = v48
                                            }
                                            _ = v47
                                        } else {
                                            fmt.Println(v22, v18)
                                            var v49 uint64 = v18
                                            _ = v49
                                        }
                                        v4 -= v18
                                    }
                                    v22 = v22
                                    _ = fi46
                                }
                                _ = fi35
                            }
                        }
                        v50 := v2[397]
                        if (101 == 2) {
                            for fi51 := 0; fi51 < 20; fi51++ {
                                var v52 uint8 = 0
                                if v1 {
                                    fmt.Println(v3, v22)
                                }
                                _ = fi51
                                _ = v52
                            }
                            _ = len(v5)
                            var v53 string = v22
                            _ = v53
                        }
                        _ = v50
                    }
                default:
                    for fi54 := 0; fi54 < 3; fi54++ {
                        switch v22 {
                        case "hello":
                            v5 = append(v5, 127)
                        }
                        _ = fi54
                    }
                }
            default:
                switch 4294967295 {
                case 9223372036854775807:
                    var v55 int = 1
                    _ = v55
                case -9223372036854775808:
                    for fi56 := 0; fi56 < 19; fi56++ {
                        v57 := GenS6{F0: (true || (!v1))}
                        _ = fi56
                        _ = v57
                    }
                case -304:
                    for fi58 := 0; fi58 < 7; fi58++ {
                        var v59 int32 = -504
                        var v60 int8 = -111
                        fmt.Println(v59, fi58)
                        _ = fi58
                        _ = v59
                        _ = v60
                    }
                default:
                    var v61 int = 4294967295
                    _ = v61
                }
            }
            _ = v22
        } else {
            if false {
                if (!v1) {
                    for fi62 := 0; fi62 < 2; fi62++ {
                        fmt.Println(v1, v3)
                        if (!(285 <= -32768)) {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            for fi63 := 0; fi63 < 6; fi63++ {
                                v3 = (-2.25 - (((v3 / v3) / v3) - (1e10 * (((float32((((3.14159 - 1.5) + ((0.0 - float64(float32(float64(v3)))) / 1.0)) + float64(-2.25))) * 0.1) * (v3 * 1e10)) * ((((float32(0.0) + v3) - ((((v3 * (((float32(0.1) / v3) + float32(0.0)) + (3.14159 - 1.5))) / 3.14159) + v3) + (3.14159 - v3))) + v3) + ((float32(1e10) / v3) * v3))))))
                                v3 = (((v3 + v3) + 0.1) / ((float32((((float64((-2.25 - v3)) + 1.5) - (-2.25 / (1.5 * (1.5 + 0.0)))) - -2.25)) + -2.25) + v3))
                                v64 := GenS7{F0: fi62}
                                _ = fi63
                                _ = v64
                            }
                        }
                        _ = fi62
                    }
                } else {
                    var v65 int = 127
                    fmt.Println(v4, v4)
                    v65 += -527
                    _ = v65
                }
            } else {
                switch -9223372036854775808 {
                case 127:
                    v1 = v1
                case 128:
                    for fi66 := 0; fi66 < 19; fi66++ {
                        var v67 uint32 = 1
                        for fi68 := 0; fi68 < 9; fi68++ {
                            if ("\n\t" != ("hello" + "")) {
                                fmt.Println(v1, v18)
                            } else {
                                var v69 float64 = (0.0 + 0.1)
                                _ = v69
                            }
                            _ = len(v2)
                            v67 -= v67
                            _ = fi68
                        }
                        _ = fi66
                        _ = v67
                    }
                }
                v70 := v6[2147483647]
                _ = v70
            }
        }
        fmt.Println(v4)
        var v71 uint64 = (v18 >> ((1) & 63))
        _ = v71
    }
    v18 += v4
    for fi72 := 0; fi72 < 18; fi72++ {
        if v1 {
            for fi73 := 0; fi73 < 20; fi73++ {
                var v74 float64 = 3.14159
                _ = fi73
                _ = v74
            }
            for fi75 := 0; fi75 < 16; fi75++ {
                v1 = false
                var v76 int16 = 0
                _ = fi75
                _ = v76
            }
        }
        var v77 int16 = -1
        var v78 int16 = v77
        _ = fi72
        _ = v77
        _ = v78
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
