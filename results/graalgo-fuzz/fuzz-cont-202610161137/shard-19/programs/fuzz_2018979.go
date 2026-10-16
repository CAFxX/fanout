package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 uint16 }
type GenS2 struct { F0 uint32; F1 uint; F2 uint8 }
type GenS3 struct { F0 uint64; F1 uint }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int16 }
func genFn0(p0 string, p1 uint8) string {
    p1 = 1
    var v4 uint8 = 190
    if true {
        for fi5 := 0; fi5 < 15; fi5++ {
            for fi6 := 0; fi6 < 20; fi6++ {
                var v7 int32 = -708
                _ = fi6
                _ = v7
            }
            _ = fi5
        }
        var v8 float64 = (float64((0.0 / -2.25)) - 3.14159)
        _ = v8
    } else {
        v9 := GenS1{F0: 2147483648, F1: -32768, F2: 65535}
        for fi10 := 0; fi10 < 1; fi10++ {
            v4 = 128
            var v11 float64 = 1.5
            v12 := GenS1{F0: 65535, F1: 1, F2: 2}
            _ = fi10
            _ = v11
            _ = v12
        }
        switch 0 {
        case -2147483648:
            if (!true) {
                for fi13 := 0; fi13 < 14; fi13++ {
                    v14 := GenS1{F0: 1, F1: 32767, F2: 32767}
                    for fi15 := 0; fi15 < 8; fi15++ {
                        var v16 uint = 65535
                        _ = fi15
                        _ = v16
                    }
                    _ = fi13
                    _ = v14
                }
                var v17 uint64 = 1
                if true {
                    v18 := GenS1{F0: 0, F1: -1, F2: 257}
                    for fi19 := 0; fi19 < 13; fi19++ {
                        v20 := []int{(-(fi19)), fi19, fi19, -9223372036854775808}
                        if len(v20) >= 2 {
                            v20 = v20[1:len(v20)]
                        }
                        fmt.Println(p0, v17, fi19)
                        _ = fi19
                        _ = v20
                    }
                    for fi21 := 0; fi21 < 20; fi21++ {
                        for fi22 := 0; fi22 < 4; fi22++ {
                            var v23 uint = 65535
                            if (!((v4 >> (((v23 % (((^(v23))) | 1))) & 63)) <= 255)) {
                                v24 := GenS2{F0: 120, F1: (v23 / ((255) | 1)), F2: (p1 | p1)}
                                if false {
                                    p1 = 212
                                    var v25 int16 = -258
                                    if true {
                                        v26 := GenS2{F0: 2147483647, F1: (v23 ^ 1), F2: 0}
                                        fmt.Println(fi21)
                                        v25 -= 0
                                        _ = v26
                                    } else {
                                        p0 = p0
                                        if (int32(fi21) > -127) {
                                            switch (fi22 - fi21) {
                                            case 0:
                                                if false {
                                                    v17 -= v17
                                                }
                                            case -2147483648:
                                                var v27 int16 = 0
                                                _ = v27
                                            default:
                                                v28 := GenS3{F0: 522, F1: 4294967295}
                                                _ = v28
                                            }
                                            v29 := []int{(-(fi22)), 32768, fi21}
                                            v23 += v23
                                            v30 := GenS1{F0: uint32(v17), F1: -32768, F2: 684}
                                            _ = v29
                                            _ = v30
                                        }
                                        v31 := []int{(fi21 % ((fi21) | 1)), fi22, (fi21 - 32767)}
                                        _ = v31
                                    }
                                    v32 := []int{}
                                    var v33 uint64 = (v17 - (v17 % ((4294967295) | 1)))
                                    _ = v25
                                    _ = v32
                                    _ = v33
                                } else {
                                    fmt.Println(p0)
                                    v34 := GenS4{F0: 1}
                                    if ((267 != 1) && (88 <= 972)) {
                                        fmt.Println(p0, p1)
                                        v4 -= (v4 * p1)
                                        var v35 string = "foo bar"
                                        var v36 int = fi21
                                        _ = v35
                                        _ = v36
                                    }
                                    _ = v34
                                }
                                fmt.Println(v4, p1)
                                var v37 int16 = 42
                                _ = v24
                                _ = v37
                            } else {
                                fmt.Println(fi22)
                            }
                            _ = fi22
                            _ = v23
                        }
                        v17 -= v17
                        _ = fi21
                    }
                    _ = v18
                }
                _ = v17
            }
        case 0:
            var v38 int16 = 1000
            _ = v38
        case 127:
            v39 := []int{}
            _ = v39
        default:
            var v40 uint16 = 0
            _ = v40
        }
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v4
    return p0
}

func main() {
    var v1 float32 = 1e10
    var v2 uint = 367
    var v3 float32 = 1.5
    if true {
        fmt.Println(v3, v3)
        for fi41 := 0; fi41 < 13; fi41++ {
            v1 -= ((1e10 * v3) + 1.5)
            var v42 float64 = 3.14159
            _ = fi41
            _ = v42
        }
    }
    fmt.Println(v2, v3)
    v43 := GenS2{F0: 0, F1: (v2 << ((4294967295) & 63)), F2: 0}
    v44 := []int{489}
    v44 = append(v44, -814)
    v3 = v1
    if (!false) {
        if true {
            switch "foo bar" {
            case "hello":
                var v45 int32 = 662
                _ = v45
            case "\n\t":
                var v46 bool = false
                _ = v46
            }
            if (((("a" + "") + "\n\t") + ("hello" + "foo bar")) < "\n\t") {
                switch 128 {
                case -2147483648:
                    v2 -= (^(v2))
                case 32767:
                    switch 127 {
                    case -468:
                        var v47 uint16 = 65535
                        _ = v47
                    }
                }
            }
        }
        for fi48 := 0; fi48 < 19; fi48++ {
            if len(v44) > 0 {
                v44[0] = (fi48 * (fi48 % (((fi48 % ((fi48) | 1))) | 1)))
            }
            _ = fi48
        }
        v49 := GenS2{F0: 4294967295, F1: v2, F2: 44}
        _ = v49
    }
    v3 += (v1 * v3)
    v1 += 0.0
    v3 = 1.5
    var v50 uint = (v2 / ((v2) | 1))
    v51 := GenS2{F0: 250, F1: 1, F2: 107}
    var v52 uint16 = 0
    var v53 int16 = -32768
    fmt.Println(v52)
    if (2147483647 >= -228) {
        v52 = (v52 & 833)
    } else {
        v53 += (^(v53))
        switch 32768 {
        case 423:
            if true {
                fmt.Println(v53)
            } else {
                for fi54 := 0; fi54 < 12; fi54++ {
                    v52 = (v52 << ((1) & 63))
                    switch 128 {
                    case 4294967295:
                        for fi55 := 0; fi55 < 11; fi55++ {
                            v56 := GenS4{F0: (v52 ^ (v52 / ((255) | 1)))}
                            _ = fi55
                            _ = v56
                        }
                    case 32767:
                        var v57 int8 = -128
                        _ = v57
                    case 128:
                        for fi58 := 0; fi58 < 12; fi58++ {
                            v1 = float32(float64(v3))
                            v59 := GenS5{F0: int16(fi58)}
                            v52 = uint16(v2)
                            _ = fi58
                            _ = v59
                        }
                    default:
                        var v60 bool = true
                        _ = v60
                    }
                    _ = fi54
                }
                v1 -= v1
                if true {
                    var v61 uint32 = 1
                    _ = v61
                } else {
                    var v62 int = -978
                    _ = v62
                }
            }
        }
        v1 = 3.14159
    }
    if true {
        if true {
            switch -9223372036854775808 {
            case -2147483648:
                for fi63 := 0; fi63 < 11; fi63++ {
                    for fi64 := 0; fi64 < 20; fi64++ {
                        var v65 float32 = float32(((0.0 / 1.5) * (1e10 * (1.5 + 0.1))))
                        var v66 uint16 = v52
                        v67 := GenS5{F0: (v53 / ((-686) | 1))}
                        _ = fi64
                        _ = v65
                        _ = v66
                        _ = v67
                    }
                    var v68 bool = true
                    v69 := GenS5{F0: v53}
                    _ = fi63
                    _ = v68
                    _ = v69
                }
            case 32767:
                for fi70 := 0; fi70 < 4; fi70++ {
                    for fi71 := 0; fi71 < 1; fi71++ {
                        v50 += (v50 << ((806) & 63))
                        _ = fi71
                    }
                    var v72 string = "\n\t"
                    _ = fi70
                    _ = v72
                }
            case 495:
                for fi73 := 0; fi73 < 5; fi73++ {
                    if len(v44) > 0 {
                        v44[0] = fi73
                    }
                    _ = fi73
                }
            default:
                fmt.Println(v50, v3, v1)
            }
            fmt.Println(v3, v2, v2)
            var v74 uint16 = v52
            _ = v74
        }
    } else {
        var v75 int8 = -128
        var v76 uint16 = (v52 - 255)
        v77 := GenS4{F0: 0}
        _ = v75
        _ = v76
        _ = v77
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v44:", len(v44), v44)
    fmt.Println("v50:", v50)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v43
    _ = v44
    _ = v50
    _ = v51
    _ = v52
    _ = v53
}
