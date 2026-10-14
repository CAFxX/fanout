package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint64; F1 int32 }
type GenS6 struct { F0 bool; F1 uint16; F2 uint }
func genFn0(p0 uint8, p1 bool, p2 uint16) int8 {
    switch "hello" {
    case "foo bar":
        p1 = true
    default:
        p1 = (-128 != 1)
    }
    p0 += (p0 | p0)
    p0 = (^(p0))
    switch 127 {
    case 130:
        if p1 {
            var v5 int8 = -1
            _ = v5
        } else {
            v6 := GenS1{F0: p1, F1: (p2 + (p2 - p2)), F2: -784}
            _ = v6
        }
    case -97:
        p1 = (!p1)
    default:
        for fi7 := 0; fi7 < 20; fi7++ {
            v8 := []int{(fi7 % ((98) | 1)), (fi7 & -1), fi7}
            _ = fi7
            _ = v8
        }
    }
    var v9 uint64 = 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return -51
}
func genFn1(p0 int, p1 uint16) int16 {
    var v10 string = ""
    v10 = v10
    if ((3.14159 == (1.5 + -2.25)) || (186 <= 127)) {
        switch 424 {
        case 1:
            p0 = (p0 >> ((65535) & 63))
        }
    } else {
        for fi11 := 0; fi11 < 16; fi11++ {
            var v12 int16 = -32768
            for fi13 := 0; fi13 < 12; fi13++ {
                v14 := []int{(fi13 | 4294967295)}
                v15 := GenS2{F0: 18446744073709551615}
                _ = fi13
                _ = v14
                _ = v15
            }
            _ = fi11
            _ = v12
        }
        switch ((("" + ("" + "\n\t")) + v10) + v10) {
        case "hello":
            for fi16 := 0; fi16 < 16; fi16++ {
                v10 = "foo bar"
                for fi17 := 0; fi17 < 20; fi17++ {
                    var v18 uint8 = 128
                    fmt.Println(fi17, v10, v18)
                    _ = fi17
                    _ = v18
                }
                v19 := GenS2{F0: 0}
                _ = fi16
                _ = v19
            }
        default:
            v10 = ((v10 + v10) + (((v10 + v10) + v10) + "hello"))
        }
        p0 = (p0 >> ((42) & 63))
    }
    var v20 bool = (2 != 114)
    var v21 uint64 = 2
    _ = p0
    _ = p1
    _ = v10
    _ = v20
    _ = v21
    return -32768
}

func main() {
    v1 := []int{}
    v2 := make(map[int]uint32)
    var v3 float32 = 1.5
    var v4 int32 = 592
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v22 int64 = 9223372036854775807
    if (true || false) {
        fmt.Println(v3, v22, v3)
    }
    _ = len(v1)
    v23 := v2[42]
    var v24 uint64 = 2
    if true {
        v2[-2147483648] = v23
        switch -1 {
        case 32768:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    } else {
        delete(v2, 2147483647)
        v25 := GenS3{F0: -2147483648}
        if false {
            for fi26 := 0; fi26 < 4; fi26++ {
                for fi27 := 0; fi27 < 6; fi27++ {
                    v23 = (^(v23))
                    _ = len(v1)
                    _ = fi27
                }
                _ = fi26
            }
        } else {
            var v28 uint64 = v24
            var v29 uint32 = (v23 >> ((42) & 63))
            v30 := GenS3{F0: v4}
            _ = v28
            _ = v29
            _ = v30
        }
        _ = v25
    }
    v3 += float32(((1e10 * 1e10) / 3.14159))
    for fi31 := 0; fi31 < 15; fi31++ {
        if (0 != 255) {
            fmt.Println(v4, fi31)
            v32 := GenS4{F0: 0}
            v3 = (float32(1e10) * ((1.5 / float32((-2.25 + 0.1))) * (v3 + v3)))
            _ = v32
        } else {
            v33 := GenS1{F0: true, F1: uint16(v24), F2: (fi31 >> ((4294967295) & 63))}
            _ = v33
        }
        for fi34 := 0; fi34 < 10; fi34++ {
            v1 = append(v1, fi31)
            if true {
                v4 -= (v4 / ((-268) | 1))
                v35 := GenS5{F0: 42, F1: int32(v23)}
                v23 += (v23 * v23)
                v36 := GenS3{F0: (v4 | v4)}
                switch 0 {
                case 2147483647:
                    switch ("\n\t" + ("a" + "hello")) {
                    case "\n\t":
                        fmt.Println(v24, v22, v22)
                    case "a":
                        var v37 float64 = 0.0
                        _ = v37
                    default:
                        v4 -= v4
                    }
                    v38 := GenS3{F0: (v4 ^ (^(v4)))}
                    v3 -= v3
                    _ = v38
                }
                if len(v1) > 0 {
                    v1[0] = fi31
                }
                var v39 int = fi31
                _ = v35
                _ = v36
                _ = v39
            } else {
                var v40 int64 = int64(v24)
                var v41 int = (fi34 & fi34)
                if len(v1) > 0 {
                    v1[0] = (^(fi31))
                }
                v42 := GenS1{F0: false, F1: 969, F2: (fi34 >> ((2) & 63))}
                _ = v40
                _ = v41
                _ = v42
            }
            _ = fi34
        }
        delete(v2, fi31)
        _ = fi31
    }
    v3 = 0.0
    v43 := GenS5{F0: v24, F1: -677}
    switch "hello" {
    case "":
        var v44 bool = (((v3 <= (-2.25 - v3)) || true) && true)
        _ = v44
    case "\n\t":
        if true {
            fmt.Println(v23, v23, v4)
            var v45 int = -620
            if ((1000 != 0) && (160 >= 18446744073709551615)) {
                if false {
                    for fi46 := 0; fi46 < 19; fi46++ {
                        var v47 uint8 = 128
                        _ = fi46
                        _ = v47
                    }
                } else {
                    var v48 bool = (!((genFn1(v45, 951) >= -32768) && (true && ((0.0 * v3) > 1.5))))
                    _ = v48
                }
            } else {
                for fi49 := 0; fi49 < 10; fi49++ {
                    v50 := GenS3{F0: v4}
                    _ = fi49
                    _ = v50
                }
                switch "foo bar" {
                case "":
                    switch "\n\t" {
                    case "a":
                        fmt.Println(v22)
                    default:
                        _ = len(v2)
                    }
                case "\n\t":
                    v4 = v4
                default:
                    v24 -= v24
                }
            }
            _ = v45
        } else {
            if true {
                if true {
                    var v51 float64 = (((((1.5 / (1e10 * 3.14159)) * (1e10 + 0.1)) + 0.0) / ((((1.5 + (0.1 + (1.5 + ((0.0 * (-2.25 + ((1.5 / 1.0) / 0.1))) / 0.1)))) + 1e10) * float64((v3 - 1.5))) / 1.0)) + (0.0 * (float64(((3.14159 * (1.5 - v3)) * float32(3.14159))) * (0.0 - 1e10))))
                    _ = v51
                } else {
                    for fi52 := 0; fi52 < 12; fi52++ {
                        v53 := GenS5{F0: (v24 % ((v24) | 1)), F1: (^(v4))}
                        var v54 float32 = (v3 - v3)
                        _ = fi52
                        _ = v53
                        _ = v54
                    }
                    delete(v2, -9223372036854775808)
                    var v55 int16 = 1000
                    _ = v55
                }
                fmt.Println(v4, v4, v4)
            } else {
                v1 = append(v1, 42)
                v2[225] = (v23 & (v23 ^ v23))
                var v56 bool = true
                _ = v56
            }
        }
    default:
        if (-32768 < -754) {
            v3 = 0.0
            v57 := v2[-409]
            fmt.Println(v3)
            _ = v57
        } else {
            for fi58 := 0; fi58 < 6; fi58++ {
                v24 += 42
                if true {
                    v24 -= (v24 % (((v24 - 2)) | 1))
                    fmt.Println(v23, v23, fi58)
                }
                _ = fi58
            }
        }
    }
    switch -287 {
    case -590:
        if true {
            v59 := GenS2{F0: 42}
            _ = v59
        } else {
            for fi60 := 0; fi60 < 12; fi60++ {
                var v61 float32 = v3
                switch (^(fi60)) {
                case 256:
                    for fi62 := 0; fi62 < 18; fi62++ {
                        v63 := GenS2{F0: 255}
                        v64 := GenS2{F0: 1}
                        var v65 int8 = 34
                        v3 = (v61 + (((v3 + v61) / v61) * v3))
                        _ = fi62
                        _ = v63
                        _ = v64
                        _ = v65
                    }
                case -334:
                    v66 := GenS2{F0: 42}
                    _ = v66
                case 2147483647:
                    if false {
                        if true {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        } else {
                            v67 := GenS6{F0: (1e10 != (v61 - 1e10)), F1: 1, F2: 18446744073709551615}
                            switch (("a" + ("foo bar" + ("a" + "foo bar"))) + ("\n\t" + "")) {
                            case "":
                                v4 -= (v4 >> ((255) & 63))
                            default:
                                fmt.Println(v4, v3, v61)
                            }
                            switch ("hello" + "") {
                            case "\n\t":
                                var v68 uint64 = (v24 - (^(v24)))
                                _ = v68
                            }
                            _ = v67
                        }
                        v61 = v3
                        var v69 uint = 18446744073709551615
                        _ = v69
                    }
                }
                _ = fi60
                _ = v61
            }
            v70 := GenS4{F0: v23}
            if false {
                v71 := GenS5{F0: v24, F1: 0}
                _ = v71
            }
            _ = v70
        }
    case 1:
        var v72 string = ""
        _ = v72
    default:
        v22 += -282
    }
    if ((101 > 0) && (128 > 128)) {
        var v73 int64 = 42
        var v74 int16 = -32768
        _ = v73
        _ = v74
    } else {
        v4 = v4
    }
    var v75 int64 = (v22 << ((717) & 63))
    var v76 uint = 803
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v75:", v75)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v43
    _ = v75
    _ = v76
}
