package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 int64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 bool; F1 uint8; F2 int8 }
type GenS5 struct { F0 int16; F1 uint16 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int; F1 uint8 }
func genFn0(p0 int32, p1 uint) string {
    p1 += p1
    var v7 uint8 = 128
    if true {
        v8 := GenS1{F0: 65535, F1: 255, F2: -33}
        if false {
            v9 := GenS2{F0: -987}
            v7 = v7
            v7 = v7
            _ = v9
        }
        _ = v8
    } else {
        if true {
            for fi10 := 0; fi10 < 7; fi10++ {
                fmt.Println(p0)
                for fi11 := 0; fi11 < 4; fi11++ {
                    var v12 float32 = 0.1
                    _ = fi11
                    _ = v12
                }
                _ = fi10
            }
        }
        v7 -= v7
    }
    _ = p0
    _ = p1
    _ = v7
    return "\n\t"
}

func main() {
    var v1 float64 = 3.14159
    var v2 uint8 = 255
    var v3 uint64 = 2
    v4 := make(map[int]string)
    var v5 uint16 = 1
    v6 := []int{-600}
    var v13 string = "hello"
    delete(v4, 0)
    var v14 int32 = 2147483647
    v15 := GenS3{F0: -1}
    if false {
        v2 = (^(v2))
        v5 += 0
    }
    fmt.Println(v5)
    var v16 int64 = 0
    if true {
        v17 := GenS1{F0: 65535, F1: (v2 >> ((65535) & 63)), F2: v16}
        for fi18 := 0; fi18 < 7; fi18++ {
            var v19 uint32 = 65535
            v5 = 2
            v20 := GenS4{F0: (1 == (v2 & (^(v2)))), F1: v2, F2: -1}
            _ = fi18
            _ = v19
            _ = v20
        }
        var v21 uint32 = 927
        _ = v17
        _ = v21
    }
    v13 = "\n\t"
    v22 := GenS5{F0: 68, F1: (^(v5))}
    v23 := GenS5{F0: 1, F1: 1}
    switch 9223372036854775807 {
    case 740:
        switch v13 {
        case "":
            v24 := v4[0]
            _ = v24
        }
    case 32768:
        if false {
            var v25 string = ("a" + "")
            switch v25 {
            case "":
                v4[4294967295] = (v13 + v25)
            case "foo bar":
                switch 0 {
                case -319:
                    fmt.Println(v2)
                case -113:
                    v16 -= 9223372036854775807
                default:
                    switch v13 {
                    case "hello":
                        if (-952 < 42) {
                            fmt.Println(v1, v13, v13)
                            var v26 int16 = 995
                            _ = v26
                        } else {
                            v27 := GenS2{F0: -1}
                            _ = v27
                        }
                    case "foo bar":
                        var v28 float32 = ((0.1 - ((-2.25 / (((1e10 + (0.1 - (float32(3.14159) - (3.14159 - -2.25)))) / 1.0) - (-2.25 * 0.0))) - 0.0)) + 0.0)
                        _ = v28
                    case "a":
                        for fi29 := 0; fi29 < 4; fi29++ {
                            fmt.Println(fi29, v14)
                            var v30 string = "foo bar"
                            if ((100 <= 100) || false) {
                                v6 = append(v6, -2147483648)
                            } else {
                                v13 = ("foo bar" + "foo bar")
                                switch fi29 {
                                case 9223372036854775807:
                                    if false {
                                        v31 := GenS6{F0: 1}
                                        if (float32(1.5) < (0.1 - -2.25)) {
                                            _ = len(v6)
                                            if false {
                                                for fi32 := 0; fi32 < 5; fi32++ {
                                                    var v33 uint32 = 0
                                                    var v34 float64 = 1.5
                                                    v2 = 54
                                                    fmt.Println(v2, v34, v13)
                                                    var v35 float64 = (v34 / (((v34 - v34) * v1) - 0.0))
                                                    _ = fi32
                                                    _ = v33
                                                    _ = v34
                                                    _ = v35
                                                }
                                                var v36 int16 = -32768
                                                _ = v36
                                            }
                                        } else {
                                            var v37 int = (fi29 + fi29)
                                            for fi38 := 0; fi38 < 11; fi38++ {
                                                fmt.Println(v30, v3)
                                                _ = fi38
                                            }
                                            _ = v37
                                        }
                                        _ = v31
                                    }
                                case 79:
                                    v39 := GenS6{F0: v5}
                                    _ = v39
                                case -9223372036854775808:
                                    v6 = append(v6, fi29)
                                default:
                                    var v40 int32 = (v14 << ((18446744073709551615) & 63))
                                    _ = v40
                                }
                                v30 = "foo bar"
                            }
                            _ = fi29
                            _ = v30
                        }
                    }
                }
            }
            for fi41 := 0; fi41 < 16; fi41++ {
                v16 -= int64(v14)
                _ = fi41
            }
            _ = v25
        } else {
            v42 := GenS2{F0: 42}
            _ = v42
        }
    }
    v43 := GenS7{F0: 42, F1: (v2 + (v2 << ((69) & 63)))}
    fmt.Println(v16, v16)
    if false {
        var v44 int32 = -2147483648
        fmt.Println(v44, v5, v14)
        _ = v44
    } else {
        var v45 uint64 = v3
        if (42 == (v5 | v5)) {
            v46 := GenS4{F0: false, F1: v2, F2: -1}
            v16 -= v16
            v47 := GenS7{F0: int(v2), F1: (v2 + (v2 >> ((877) & 63)))}
            _ = v46
            _ = v47
        }
        if false {
            var v48 int8 = 16
            if false {
                if (!false) {
                    for fi49 := 0; fi49 < 4; fi49++ {
                        v2 = (^(v2))
                        v50 := v4[-2147483648]
                        v51 := v4[int(v2)]
                        _ = fi49
                        _ = v50
                        _ = v51
                    }
                    v52 := v4[824]
                    for fi53 := 0; fi53 < 11; fi53++ {
                        v45 += (v45 | uint64(v5))
                        var v54 int8 = v48
                        fmt.Println(v54, v48, v5)
                        _ = fi53
                        _ = v54
                    }
                    _ = v52
                }
            } else {
                fmt.Println(v3, v2, v48)
            }
            for fi55 := 0; fi55 < 6; fi55++ {
                if true {
                    switch fi55 {
                    case 32768:
                        v5 -= (v5 + v5)
                    case 128:
                        delete(v4, fi55)
                    default:
                        v6 = append(v6, fi55)
                    }
                    fmt.Println(v16, v5)
                    fmt.Println(v2)
                }
                v1 += (0.1 / (v1 - -2.25))
                _ = fi55
            }
            _ = v48
        } else {
            v3 = v45
            if (42 <= 2) {
                for fi56 := 0; fi56 < 9; fi56++ {
                    switch int(v14) {
                    case 449:
                        var v57 uint64 = 1
                        _ = v57
                    default:
                        switch v13 {
                        case "hello":
                            v14 -= int32(fi56)
                        }
                    }
                    if false {
                        if false {
                            switch v13 {
                            case "hello":
                                v16 -= v16
                            }
                        }
                        v4[(fi56 ^ 30)] = v13
                        _ = len(v4)
                    } else {
                        switch v13 {
                        case "hello":
                            if (!((^(v14)) < (^(v14)))) {
                                var v58 uint16 = 32767
                                _ = v58
                            } else {
                                var v59 int64 = v16
                                var v60 uint64 = (v3 ^ v45)
                                var v61 uint16 = (v5 >> ((42) & 63))
                                _ = v59
                                _ = v60
                                _ = v61
                            }
                        default:
                            var v62 float64 = v1
                            _ = v62
                        }
                        v5 = v5
                    }
                    v4[int(v14)] = v13
                    _ = fi56
                }
            } else {
                v4[284] = (v13 + v13)
            }
        }
        _ = v45
    }
    v6 = append(v6, -240)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
}
