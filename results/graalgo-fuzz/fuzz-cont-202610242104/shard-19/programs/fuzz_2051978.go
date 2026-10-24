package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int; F1 int8; F2 int64 }
type GenS5 struct { F0 bool; F1 int32; F2 int16 }
type GenS6 struct { F0 uint16; F1 string; F2 uint32 }
type GenS7 struct { F0 int; F1 int16; F2 uint }
type GenS8 struct { F0 uint64 }
func genFn0(p0 bool, p1 string) uint8 {
    if (1 < 1) {
        var v5 string = "foo bar"
        _ = v5
    } else {
        v6 := GenS1{F0: -17, F1: p0}
        _ = v6
    }
    if (!true) {
        var v7 uint32 = 2147483647
        _ = v7
    }
    _ = p0
    _ = p1
    return 111
}
func genFn1(p0 int32, p1 int32, p2 int8) uint16 {
    switch 128 {
    case 128:
        var v8 uint64 = 18446744073709551615
        _ = v8
    }
    p0 += p0
    v9 := GenS1{F0: 0, F1: true}
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 65535
}

func main() {
    var v1 float32 = 0.0
    var v2 uint = 42
    var v3 uint8 = 0
    var v4 string = ("foo bar" + "a")
    v3 = (v3 + 255)
    v10 := []int{}
    fmt.Println(v2)
    v3 = (^(v3))
    var v11 uint32 = 65535
    var v12 uint64 = 18446744073709551615
    for fi13 := 0; fi13 < 17; fi13++ {
        var v14 int32 = 2147483647
        if ((-32768 == 1000) || false) {
            var v15 int8 = 100
            _ = v15
        } else {
            v12 = v12
            v1 += (3.14159 / v1)
        }
        var v16 float32 = (v1 - float32(((float64((1.5 + (1.5 - 3.14159))) / 0.1) / (1e10 * 3.14159))))
        _ = fi13
        _ = v14
        _ = v16
    }
    v17 := GenS2{F0: uint(v12)}
    if true {
        if true {
            fmt.Println(v4, v2, v1)
            for fi18 := 0; fi18 < 5; fi18++ {
                if (false && true) {
                    if (2 <= 1) {
                        switch v4 {
                        case "foo bar":
                            for fi19 := 0; fi19 < 15; fi19++ {
                                var v20 int32 = 2147483647
                                _ = fi19
                                _ = v20
                            }
                        case "\n\t":
                            v21 := GenS3{F0: -55}
                            _ = v21
                        default:
                            var v22 int32 = 922
                            _ = v22
                        }
                    }
                } else {
                    var v23 uint64 = v12
                    _ = v23
                }
                switch fi18 {
                case -672:
                    for fi24 := 0; fi24 < 17; fi24++ {
                        fmt.Println(v2, v11, v2)
                        if true {
                            v25 := GenS4{F0: fi24, F1: 100, F2: 42}
                            if (false || false) {
                                v26 := GenS4{F0: -514, F1: 23, F2: -9223372036854775808}
                                var v27 uint64 = 0
                                switch 128 {
                                case 682:
                                    var v28 int = (fi24 - (^(fi18)))
                                    _ = v28
                                case 4294967295:
                                    var v29 uint64 = (v12 ^ v27)
                                    _ = v29
                                }
                                _ = v26
                                _ = v27
                            } else {
                                v30 := GenS4{F0: fi24, F1: -125, F2: 9223372036854775807}
                                if len(v10) > 0 {
                                    v10[0] = -370
                                }
                                v1 -= (1.5 * (((1.5 + (float32(1e10) * float32(1e10))) * (v1 + v1)) / 3.14159))
                                _ = v30
                            }
                            var v31 int64 = 649
                            fmt.Println(v3)
                            if (!(v11 != (v11 * uint32(v12)))) {
                                v4 = v4
                                var v32 uint8 = 255
                                _ = v32
                            }
                            _ = v25
                            _ = v31
                        }
                        _ = fi24
                    }
                case 111:
                    v2 = 18446744073709551615
                }
                _ = fi18
            }
            if true {
                if (true || false) {
                    _ = len(v10)
                    if (-128 != 36) {
                        var v33 uint64 = v12
                        v34 := GenS3{F0: -2147483648}
                        if false {
                            v2 = v2
                            v35 := GenS4{F0: 127, F1: -1, F2: 354}
                            _ = v35
                        }
                        _ = v33
                        _ = v34
                    } else {
                        var v36 uint = 371
                        switch v4 {
                        case "\n\t":
                            if false {
                                if len(v10) >= 2 {
                                    v10 = v10[1:len(v10)]
                                }
                                fmt.Println(v2, v12, v11)
                                if false {
                                    var v37 int64 = 1
                                    v38 := GenS5{F0: false, F1: -1, F2: 32767}
                                    fmt.Println(v4, v1)
                                    _ = v37
                                    _ = v38
                                }
                            } else {
                                if len(v10) >= 2 {
                                    v10 = v10[1:len(v10)]
                                }
                            }
                        case "":
                            v12 -= (v12 >> ((v36) & 63))
                        case "foo bar":
                            v1 = v1
                        }
                        var v39 int = -143
                        _ = v36
                        _ = v39
                    }
                    v12 = 42
                }
                if false {
                    v40 := GenS5{F0: ((((^(v12)) == v12) || false) || false), F1: -1, F2: -32768}
                    switch "\n\t" {
                    case "hello":
                        v10 = append(v10, -106)
                    case "foo bar":
                        if (true || false) {
                            for fi41 := 0; fi41 < 19; fi41++ {
                                switch v4 {
                                case "foo bar":
                                    var v42 int64 = 1
                                    _ = v42
                                case "a":
                                    for fi43 := 0; fi43 < 12; fi43++ {
                                        var v44 int64 = 227
                                        fmt.Println(v11)
                                        _ = fi43
                                        _ = v44
                                    }
                                }
                                _ = fi41
                            }
                            var v45 uint64 = (v12 * v12)
                            switch 490 {
                            case 42:
                                for fi46 := 0; fi46 < 18; fi46++ {
                                    if (true && (!(!((1 > 2147483647) || false)))) {
                                        v1 = v1
                                    }
                                    var v47 float32 = v1
                                    _ = fi46
                                    _ = v47
                                }
                            case 128:
                                fmt.Println(v11)
                            case -919:
                                switch 201 {
                                case 1:
                                    var v48 uint64 = (v45 & v12)
                                    _ = v48
                                }
                            }
                            _ = v45
                        } else {
                            fmt.Println(v12, v12)
                        }
                    }
                    var v49 int = int(v3)
                    _ = v40
                    _ = v49
                }
            } else {
                v12 -= (^(v12))
                switch 2147483647 {
                case 431:
                    if (-9223372036854775808 >= -1) {
                        v3 += v3
                        v1 -= 1e10
                    } else {
                        fmt.Println(v1, v2, v2)
                        if false {
                            v10 = append(v10, -9223372036854775808)
                            if len(v10) >= 2 {
                                v10 = v10[1:len(v10)]
                            }
                            if ((v1 - 1e10) <= 3.14159) {
                                if true {
                                    v10 = append(v10, -541)
                                } else {
                                    if true {
                                        var v50 uint16 = genFn1(42, 2147483647, -25)
                                        if (v1 > ((3.14159 * v1) / (v1 + float32(float64((0.0 * (3.14159 - ((((-2.25 * float32(-2.25)) + float32(0.1)) * float32((3.14159 / 1e10))) - ((v1 / v1) / 1.0))))))))) {
                                            var v51 int8 = 1
                                            var v52 float32 = (v1 / (v1 - 0.0))
                                            v52 -= v52
                                            _ = v51
                                            _ = v52
                                        } else {
                                            var v53 int32 = 787
                                            _ = v53
                                        }
                                        _ = v50
                                    } else {
                                        if true {
                                            switch -1 {
                                            case 2147483647:
                                                v11 = (v11 ^ (v11 << ((29) & 63)))
                                            case -1:
                                                fmt.Println(v3)
                                            default:
                                                if len(v10) >= 2 {
                                                    v10 = v10[1:len(v10)]
                                                }
                                            }
                                        }
                                    }
                                    v2 = v2
                                    v2 += 18446744073709551615
                                }
                            } else {
                                var v54 int32 = -2147483648
                                v12 += v12
                                if false {
                                    fmt.Println(v11, v4)
                                    v54 = 2147483647
                                    for fi55 := 0; fi55 < 8; fi55++ {
                                        for fi56 := 0; fi56 < 6; fi56++ {
                                            var v57 uint8 = (^(v3))
                                            var v58 uint = 255
                                            v59 := GenS6{F0: genFn1((v54 ^ (v54 << (((v2 % (((v58 * v2)) | 1))) & 63))), (v54 * v54), 1), F1: v4, F2: (v11 % ((42) | 1))}
                                            _ = fi56
                                            _ = v57
                                            _ = v58
                                            _ = v59
                                        }
                                        v1 += (float32((-2.25 - (float64(float32(1.5)) - 0.1))) * v1)
                                        v60 := GenS7{F0: (fi55 & (fi55 & (^(fi55)))), F1: -32768, F2: (v2 << (((v2 & v2)) & 63))}
                                        _ = fi55
                                        _ = v60
                                    }
                                } else {
                                    if true {
                                        var v61 int64 = 167
                                        for fi62 := 0; fi62 < 11; fi62++ {
                                            v63 := GenS2{F0: 42}
                                            if false {
                                                if ("foo bar" > "\n\t") {
                                                    fmt.Println(fi62, v2)
                                                    fmt.Println(v4)
                                                    if len(v10) >= 2 {
                                                        v10 = v10[1:len(v10)]
                                                    }
                                                }
                                                switch fi62 {
                                                case -241:
                                                    _ = len(v10)
                                                }
                                                if len(v10) > 0 {
                                                    v10[0] = (fi62 + fi62)
                                                }
                                            }
                                            v12 = (v12 % ((18446744073709551615) | 1))
                                            _ = fi62
                                            _ = v63
                                        }
                                        _ = v61
                                    }
                                    var v64 uint64 = (v12 ^ (v12 - 416))
                                    _ = v64
                                }
                                _ = v54
                            }
                        }
                        switch "foo bar" {
                        case "a":
                            if true {
                                for fi65 := 0; fi65 < 6; fi65++ {
                                    v11 = uint32(v2)
                                    _ = fi65
                                }
                                var v66 string = v4
                                for fi67 := 0; fi67 < 16; fi67++ {
                                    v68 := GenS7{F0: fi67, F1: 1000, F2: v2}
                                    v1 = 1e10
                                    _ = fi67
                                    _ = v68
                                }
                                _ = v66
                            }
                        case "":
                            v69 := GenS5{F0: true, F1: -289, F2: -4}
                            _ = v69
                        }
                    }
                case 0:
                    v1 = -2.25
                }
                for fi70 := 0; fi70 < 19; fi70++ {
                    v11 = 1
                    _ = fi70
                }
            }
        } else {
            v4 = v4
            for fi71 := 0; fi71 < 13; fi71++ {
                if true {
                    var v72 uint16 = 2
                    _ = v72
                }
                v73 := GenS4{F0: int(v12), F1: int8(v12), F2: -1}
                v3 = v3
                _ = fi71
                _ = v73
            }
            var v74 int64 = 9223372036854775807
            _ = v74
        }
        v75 := GenS8{F0: v12}
        var v76 int = 504
        _ = v75
        _ = v76
    }
    v10 = append(v10, 2147483647)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
