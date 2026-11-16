package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16; F1 int8; F2 int8 }
type GenS3 struct { F0 int16; F1 uint8; F2 bool }
type GenS4 struct { F0 string; F1 uint }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint16; F1 bool }
type GenS7 struct { F0 string; F1 uint }
type GenS8 struct { F0 int16; F1 uint8; F2 int32 }
func genFn0() uint8 {
    v5 := GenS1{F0: 2}
    var v6 string = "\n\t"
    var v7 uint32 = 42
    var v8 bool = ((1 > 100) || ((!true) || (!true)))
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 127
}
func genFn1(p0 int16) uint32 {
    v9 := GenS1{F0: 42}
    if (!false) {
        fmt.Println(p0)
    }
    _ = p0
    _ = v9
    return 2147483647
}

func main() {
    v1 := []uint32{0, 2, 2}
    var v2 string = "a"
    v3 := make(map[uint32]bool)
    var v4 uint = 65535
    switch v2 {
    case "foo bar":
        var v10 string = v2
        _ = v10
    }
    var v11 float64 = (1.5 * 3.14159)
    fmt.Println(v11, v4, v11)
    var v12 uint16 = 1
    if false {
        v11 -= float64(-2.25)
        for fi13 := 0; fi13 < 14; fi13++ {
            v4 += v4
            if (false && false) {
                var v14 int32 = int32(v4)
                _ = v14
            }
            _ = fi13
        }
        switch v2 {
        case "a":
            for fi15 := 0; fi15 < 9; fi15++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                for fi16 := 0; fi16 < 4; fi16++ {
                    if true {
                        switch fi15 {
                        case 32767:
                            if true {
                                if (false || (2147483647 < 32768)) {
                                    v17 := v3[genFn1(-32768)]
                                    var v18 uint32 = 361
                                    _ = len(v3)
                                    _ = v17
                                    _ = v18
                                } else {
                                    var v19 int64 = 0
                                    v20 := GenS2{F0: -884, F1: -1, F2: -84}
                                    if true {
                                        var v21 bool = false
                                        switch (fi16 ^ (fi16 / ((4294967295) | 1))) {
                                        case -1:
                                            if true {
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                fmt.Println(v2)
                                            } else {
                                                delete(v3, 0)
                                                var v22 float64 = v11
                                                var v23 float32 = (((-2.25 * 1e10) * (0.0 + 3.14159)) + 0.1)
                                                _ = v22
                                                _ = v23
                                            }
                                        case 9223372036854775807:
                                            v24 := v3[65535]
                                            _ = v24
                                        case 383:
                                            switch (fi16 * (^(fi15))) {
                                            case -686:
                                                var v25 float64 = 3.14159
                                                _ = v25
                                            case 128:
                                                switch fi15 {
                                                case -1:
                                                    v26 := GenS2{F0: 42, F1: -1, F2: -1}
                                                    _ = v26
                                                case 9223372036854775807:
                                                    v1 = append(v1, 722)
                                                case 32767:
                                                    var v27 float32 = 0.1
                                                    _ = v27
                                                default:
                                                    var v28 uint16 = (v12 - (^(v12)))
                                                    _ = v28
                                                }
                                            case 127:
                                                var v29 int = fi16
                                                _ = v29
                                            default:
                                                v30 := GenS3{F0: 68, F1: 255, F2: (0 == -843)}
                                                _ = v30
                                            }
                                        default:
                                            v31 := GenS4{F0: (v2 + ("hello" + ((("\n\t" + v2) + "hello") + v2))), F1: (v4 * (v4 * (^(v4))))}
                                            _ = v31
                                        }
                                        v32 := GenS2{F0: 42, F1: 1, F2: -115}
                                        _ = v21
                                        _ = v32
                                    }
                                    _ = v19
                                    _ = v20
                                }
                                if true {
                                    v11 += (v11 - 3.14159)
                                    fmt.Println(fi16, v4)
                                }
                                if (1 <= -499) {
                                    _ = len(v1)
                                    var v33 uint8 = 2
                                    v33 = (^(v33))
                                    _ = v33
                                } else {
                                    switch v2 {
                                    case "foo bar":
                                        if len(v1) > 0 {
                                            v1[0] = 2147483648
                                        }
                                    }
                                    fmt.Println(v12, v4, v12)
                                    if len(v1) > 0 {
                                        v1[0] = 4294967295
                                    }
                                }
                            }
                        }
                        if (((v11 / 1.5) + 0.0) == v11) {
                            v4 += 2
                            if true {
                                if true {
                                    v11 -= (float64((1.5 * float32(v11))) + float64(-2.25))
                                } else {
                                    v2 = v2
                                    var v34 uint8 = 127
                                    _ = len(v3)
                                    _ = v34
                                }
                                if (true || (genFn0() > genFn0())) {
                                    fmt.Println(v12, v4, v4)
                                } else {
                                    var v35 string = v2
                                    if true {
                                        if true {
                                            v36 := GenS4{F0: v2, F1: v4}
                                            var v37 int64 = 1
                                            v2 = (v2 + (v2 + v2))
                                            switch fi16 {
                                            case -34:
                                                switch ("hello" + ("\n\t" + v35)) {
                                                case "hello":
                                                    var v38 string = v2
                                                    _ = v38
                                                }
                                            case 128:
                                                var v39 int16 = 1
                                                v2 = v35
                                                v39 = 42
                                                _ = len(v3)
                                                _ = v39
                                            case 32768:
                                                switch (fi15 % (((fi16 + fi15)) | 1)) {
                                                case 127:
                                                    var v40 int16 = -775
                                                    _ = v40
                                                default:
                                                    var v41 float64 = float64(0.0)
                                                    _ = v41
                                                }
                                            }
                                            _ = v36
                                            _ = v37
                                        } else {
                                            fmt.Println(v12, v2, fi15)
                                            v4 -= v4
                                        }
                                        v2 = ("foo bar" + "")
                                        v42 := GenS3{F0: -1, F1: 62, F2: (v2 < v2)}
                                        _ = v42
                                    }
                                    var v43 float32 = ((3.14159 / 1.0) / (1e10 / 0.1))
                                    _ = v35
                                    _ = v43
                                }
                                v4 = v4
                            }
                        }
                        if (v4 <= (v4 / (((v4 | v4)) | 1))) {
                            if false {
                                switch (v2 + ("a" + ("" + (v2 + ("foo bar" + "a"))))) {
                                case "\n\t":
                                    var v44 int16 = 993
                                    _ = v44
                                case "hello":
                                    if (false || ((v12 << ((18) & 63)) < (v12 ^ 201))) {
                                        fmt.Println(v12, fi16)
                                        v4 = v4
                                        fmt.Println(fi16, v11, v12)
                                        if ((fi15 <= (fi16 / ((-607) | 1))) || (1 > -128)) {
                                            v12 += (v12 | v12)
                                        } else {
                                            var v45 float64 = float64(1e10)
                                            v2 = "foo bar"
                                            _ = v45
                                        }
                                    }
                                default:
                                    var v46 int16 = 42
                                    _ = v46
                                }
                            } else {
                                v2 = "foo bar"
                            }
                        } else {
                            fmt.Println(v2)
                        }
                    } else {
                        v2 = v2
                        switch v2 {
                        case "foo bar":
                            v11 += v11
                        case "":
                            v2 = ((v2 + (v2 + "\n\t")) + (v2 + (v2 + v2)))
                            v11 = ((float64((0.1 / (1e10 - 0.0))) - float64(1.5)) / 1.0)
                            if (false && false) {
                                var v47 int16 = 0
                                var v48 int = fi15
                                delete(v3, genFn1(1))
                                _ = v47
                                _ = v48
                            } else {
                                if ((!(1 < 104)) && true) {
                                    var v49 uint8 = 1
                                    v12 = 65535
                                    v49 = (v49 | v49)
                                    _ = v49
                                } else {
                                    v11 = v11
                                }
                                switch v2 {
                                case "":
                                    switch v2 {
                                    case "foo bar":
                                        var v50 uint = 255
                                        _ = v50
                                    case "":
                                        v1 = append(v1, 1)
                                    }
                                case "hello":
                                    v12 = (v12 % (((v12 + v12)) | 1))
                                default:
                                    v51 := GenS5{F0: 42}
                                    _ = v51
                                }
                            }
                            v52 := GenS6{F0: v12, F1: true}
                            if ((v2 + (v2 + v2)) == ("a" + (v2 + v2))) {
                                if (-966 != 1) {
                                    var v53 uint16 = (v12 << ((65535) & 63))
                                    v11 -= 0.1
                                    _ = v53
                                }
                            } else {
                                v4 = v4
                                var v54 uint16 = (v12 ^ (v12 ^ v12))
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                _ = v54
                            }
                            _ = v52
                        default:
                            v55 := GenS7{F0: v2, F1: 255}
                            _ = v55
                        }
                        var v56 uint8 = 90
                        var v57 int16 = 0
                        fmt.Println(v4, v57)
                        _ = v56
                        _ = v57
                    }
                    var v58 uint16 = (v12 / ((v12) | 1))
                    v2 = v2
                    _ = fi16
                    _ = v58
                }
                _ = fi15
            }
        }
    } else {
        _ = len(v1)
    }
    v4 += v4
    if (127 < 127) {
        if (true && false) {
            var v59 int = 32768
            _ = v59
        } else {
            switch int(v4) {
            case 32767:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case -9223372036854775808:
                var v60 uint8 = 127
                _ = v60
            }
            v61 := GenS8{F0: 269, F1: 42, F2: 0}
            _ = v61
        }
        for fi62 := 0; fi62 < 4; fi62++ {
            v2 = (v2 + (v2 + "hello"))
            v2 = (((v2 + ((((v2 + (v2 + (v2 + ("a" + ((v2 + "\n\t") + ((v2 + (v2 + v2)) + ((v2 + "\n\t") + v2))))))) + v2) + ((v2 + ((v2 + ("hello" + (v2 + v2))) + "a")) + v2)) + v2)) + v2) + "")
            _ = fi62
        }
    } else {
        fmt.Println(v11, v12)
        v12 += (v12 << (((v4 * 42)) & 63))
        if ((false && true) && true) {
            fmt.Println(v11)
        }
    }
    if (!(genFn0() <= 127)) {
        switch v2 {
        case "":
            v12 -= 65535
        case "hello":
            var v63 float32 = (-2.25 + float32(v11))
            _ = v63
        case "\n\t":
            for fi64 := 0; fi64 < 14; fi64++ {
                switch 9223372036854775807 {
                case -984:
                    var v65 uint64 = 1
                    _ = v65
                default:
                    v4 += (v4 - v4)
                }
                _ = fi64
            }
        default:
            var v66 float64 = (-2.25 + v11)
            _ = v66
        }
        var v67 bool = true
        _ = v67
    } else {
        var v68 bool = (v4 > 42)
        _ = v68
    }
    var v69 uint64 = 2
    switch 2147483647 {
    case 32767:
        v2 = (v2 + v2)
    case 0:
        _ = len(v3)
    default:
        v4 = v4
    }
    var v70 float64 = v11
    var v71 bool = (-80 > int8(v69))
    var v72 uint16 = v12
    switch 42 {
    case 98:
        if (false && (uint32(v69) > 693)) {
            v73 := v3[4294967295]
            _ = v73
        }
    case -788:
        for fi74 := 0; fi74 < 19; fi74++ {
            v75 := GenS3{F0: 1000, F1: genFn0(), F2: v71}
            _ = fi74
            _ = v75
        }
    case 32767:
        if ((v69 + v69) < (v69 | (v69 << ((v4) & 63)))) {
            var v76 int16 = -631
            var v77 float64 = float64((0.0 + float32(3.14159)))
            _ = v76
            _ = v77
        } else {
            var v78 uint64 = (v69 % (((^(v69))) | 1))
            v79 := GenS5{F0: v69}
            _ = v78
            _ = v79
        }
    default:
        v2 = "hello"
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v69
    _ = v70
    _ = v71
    _ = v72
}
