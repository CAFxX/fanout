package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 uint32; F1 int16; F2 int32 }
type GenS4 struct { F0 uint; F1 uint64 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int32; F1 bool }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int32; F1 uint16; F2 uint }
type GenS9 struct { F0 uint64; F1 uint16; F2 int8 }
func genFn0(p0 int16, p1 int16) int8 {
    var v4 int16 = (-(p1))
    for fi5 := 0; fi5 < 1; fi5++ {
        p0 = (p0 << ((1) & 63))
        _ = fi5
    }
    _ = p0
    _ = p1
    _ = v4
    return 127
}
func genFn1(p0 bool, p1 uint8) uint32 {
    fmt.Println(p0, p0)
    for fi6 := 0; fi6 < 11; fi6++ {
        p0 = ((p0 || (!false)) || p0)
        switch 2147483647 {
        case 9223372036854775807:
            var v7 int8 = -128
            _ = v7
        case 32767:
            for fi8 := 0; fi8 < 13; fi8++ {
                v9 := GenS1{F0: p0, F1: (("" + "\n\t") + "hello")}
                var v10 int64 = 42
                v11 := []int{32767, fi8}
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
        case -244:
            v12 := []int{}
            _ = v12
        }
        _ = fi6
    }
    var v13 uint64 = 18446744073709551615
    var v14 int8 = -128
    v15 := GenS2{F0: 2, F1: v14}
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    _ = v15
    return 2147483648
}

func main() {
    v1 := []int{127}
    var v2 int8 = 100
    var v3 int64 = 625
    var v16 uint16 = 65535
    if false {
        v3 = -1
        var v17 int8 = 0
        _ = v17
    } else {
        v18 := GenS2{F0: 4294967295, F1: v2}
        if (false || true) {
            for fi19 := 0; fi19 < 19; fi19++ {
                for fi20 := 0; fi20 < 9; fi20++ {
                    v21 := GenS2{F0: 2147483648, F1: int8(v3)}
                    _ = fi20
                    _ = v21
                }
                _ = fi19
            }
            v3 = v3
        }
        var v22 string = "a"
        _ = v18
        _ = v22
    }
    var v23 uint = 790
    if false {
        v2 -= v2
        v2 -= genFn0(1000, 1)
    }
    var v24 string = "\n\t"
    v25 := GenS3{F0: 2147483648, F1: 42, F2: -2147483648}
    v23 -= (v23 | v23)
    for fi26 := 0; fi26 < 6; fi26++ {
        if false {
            if ((0.0 / float32(3.14159)) > (float32(-2.25) / -2.25)) {
                for fi27 := 0; fi27 < 5; fi27++ {
                    v23 += 1
                    v1 = append(v1, fi27)
                    v28 := GenS4{F0: uint(v16), F1: 2}
                    _ = fi27
                    _ = v28
                }
            }
        }
        _ = fi26
    }
    if ((!(-1 > -2147483648)) && (false || (-552 > 0))) {
        v29 := GenS4{F0: (v23 & uint(v16)), F1: 18446744073709551615}
        switch v24 {
        case "":
            switch 390 {
            case 847:
                v16 = 42
            case -9223372036854775808:
                _ = len(v1)
            }
        case "hello":
            var v30 uint64 = 426
            _ = v30
        default:
            var v31 uint8 = 2
            _ = v31
        }
        if (false || true) {
            var v32 uint = v23
            for fi33 := 0; fi33 < 9; fi33++ {
                switch v24 {
                case "hello":
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case "foo bar":
                    for fi34 := 0; fi34 < 2; fi34++ {
                        var v35 string = "a"
                        if (v35 != "a") {
                            var v36 uint32 = 510
                            _ = v36
                        }
                        if false {
                            v24 = v35
                            fmt.Println(v23, v16)
                            v35 = v35
                        }
                        _ = fi34
                        _ = v35
                    }
                }
                if false {
                    if true {
                        for fi37 := 0; fi37 < 1; fi37++ {
                            if (!((v16 >> ((2) & 63)) == 32767)) {
                                var v38 int16 = -1
                                if (134 == 18446744073709551615) {
                                    switch v24 {
                                    case "":
                                        var v39 int64 = v3
                                        _ = v39
                                    case "foo bar":
                                        switch -1 {
                                        case 32767:
                                            v1 = append(v1, fi37)
                                        case 395:
                                            if true {
                                                v40 := GenS5{F0: (v2 % ((genFn0(v38, (v38 & v38))) | 1))}
                                                if len(v1) > 0 {
                                                    v1[0] = fi37
                                                }
                                                _ = v40
                                            } else {
                                                if (((1.5 * (0.0 - float64(1e10))) - (-2.25 * (1e10 / 1.5))) != -2.25) {
                                                    if len(v1) > 0 {
                                                        v1[0] = fi37
                                                    }
                                                    v23 = v32
                                                }
                                                v32 -= v23
                                                v24 = v24
                                            }
                                        case -125:
                                            var v41 uint32 = 65535
                                            _ = v41
                                        }
                                    }
                                    var v42 int16 = (v38 ^ v38)
                                    var v43 string = v24
                                    _ = v42
                                    _ = v43
                                }
                                var v44 float32 = 1e10
                                _ = v38
                                _ = v44
                            } else {
                                if (!false) {
                                    if len(v1) > 0 {
                                        v1[0] = (fi33 % (((-(fi37))) | 1))
                                    }
                                    if true {
                                        if true {
                                            v23 -= 65535
                                        } else {
                                            switch (fi37 << ((842) & 63)) {
                                            case -1:
                                                if true {
                                                    var v45 uint16 = 330
                                                    var v46 uint8 = 2
                                                    var v47 float32 = float32(0.1)
                                                    var v48 uint8 = (v46 >> (((v23 & v23)) & 63))
                                                    var v49 int32 = -279
                                                    _ = v45
                                                    _ = v46
                                                    _ = v47
                                                    _ = v48
                                                    _ = v49
                                                } else {
                                                    v50 := GenS5{F0: v2}
                                                    v51 := GenS6{F0: -2147483648, F1: false}
                                                    _ = v50
                                                    _ = v51
                                                }
                                            case 214:
                                                switch v24 {
                                                case "":
                                                    v52 := GenS7{F0: v32}
                                                    _ = v52
                                                case "foo bar":
                                                    v32 -= v32
                                                }
                                            case -9223372036854775808:
                                                v32 = 496
                                            default:
                                                v24 = v24
                                            }
                                            fmt.Println(v3, v24)
                                            var v53 float32 = float32((1e10 / 1.0))
                                            _ = v53
                                        }
                                        v54 := GenS8{F0: -1, F1: v16, F2: (v23 | v23)}
                                        v24 = ((("hello" + ((v24 + "hello") + "foo bar")) + ("foo bar" + v24)) + "foo bar")
                                        _ = v54
                                    }
                                    fmt.Println(v32, v23)
                                }
                            }
                            _ = fi37
                        }
                        if len(v1) > 0 {
                            v1[0] = fi33
                        }
                    }
                } else {
                    v24 = "foo bar"
                    for fi55 := 0; fi55 < 8; fi55++ {
                        v1 = append(v1, 487)
                        v16 = v16
                        switch fi55 {
                        case 700:
                            v16 = (v16 << ((v32) & 63))
                        default:
                            var v56 uint64 = 2
                            _ = v56
                        }
                        _ = fi55
                    }
                }
                _ = fi33
            }
            _ = v32
        }
        _ = v29
    }
    if (((uint32(v16) >= 2147483648) || true) && (((v3 + (v3 / (((v3 * v3)) | 1))) <= (v3 - (v3 | v3))) || false)) {
        v57 := GenS3{F0: 2, F1: 442, F2: -216}
        _ = v57
    }
    var v58 int8 = (v2 * 100)
    var v59 int16 = 42
    if ((591 >= 476) || false) {
        v3 = v3
    }
    var v60 int64 = 0
    v58 -= v58
    v61 := GenS9{F0: 636, F1: 42, F2: 0}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v16
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v58
    _ = v59
    _ = v60
    _ = v61
}
