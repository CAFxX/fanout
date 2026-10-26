package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint16; F1 int8; F2 bool }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint64; F1 bool }
type GenS7 struct { F0 int16; F1 uint16; F2 int16 }
type GenS8 struct { F0 int32 }
func genFn0(p0 int, p1 int64, p2 int64) uint32 {
    var v7 uint8 = 1
    p2 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return uint32(v7)
}
func genFn1() int {
    for fi8 := 0; fi8 < 18; fi8++ {
        fmt.Println(fi8)
        _ = fi8
    }
    v9 := GenS1{F0: -32768, F1: 1}
    if (2 > 65535) {
        for fi10 := 0; fi10 < 2; fi10++ {
            var v11 int16 = -32768
            if (2 > 2) {
                v12 := GenS1{F0: (v11 >> ((18446744073709551615) & 63)), F1: -649}
                _ = v12
            } else {
                if (!true) {
                    v13 := []int{0}
                    for fi14 := 0; fi14 < 17; fi14++ {
                        v11 += (v11 - v11)
                        if false {
                            var v15 uint64 = 823
                            _ = v15
                        } else {
                            switch fi14 {
                            case 9223372036854775807:
                                if ((255 == 292) || (-2.25 < (-2.25 * float64((0.1 + ((float32(float64(0.1)) + -2.25) * (-2.25 - (0.1 / (1e10 / 3.14159))))))))) {
                                    v16 := GenS1{F0: 1000, F1: -1}
                                    var v17 uint8 = 0
                                    if false {
                                        if len(v13) >= 2 {
                                            v13 = v13[1:len(v13)]
                                        }
                                        v18 := GenS1{F0: (v11 - v11), F1: 0}
                                        v19 := GenS2{F0: 2147483648, F1: -9223372036854775808}
                                        _ = v18
                                        _ = v19
                                    }
                                    _ = v16
                                    _ = v17
                                } else {
                                    if false {
                                        var v20 float32 = ((((3.14159 - float32(0.0)) + 1e10) - 0.0) / (float32(0.0) + (3.14159 / 3.14159)))
                                        _ = v20
                                    } else {
                                        var v21 string = ("hello" + (("foo bar" + "\n\t") + (("\n\t" + "foo bar") + (("a" + "foo bar") + ""))))
                                        switch ((v21 + v21) + v21) {
                                        case "\n\t":
                                            var v22 float64 = 0.0
                                            _ = v22
                                        case "foo bar":
                                            switch 32768 {
                                            case -2147483648:
                                                if (true && (0 != 404)) {
                                                    if len(v13) > 0 {
                                                        v13[0] = 0
                                                    }
                                                    var v23 bool = false
                                                    _ = v23
                                                } else {
                                                    if len(v13) >= 2 {
                                                        v13 = v13[1:len(v13)]
                                                    }
                                                    v11 += 0
                                                    v24 := GenS2{F0: 201, F1: (fi14 & fi10)}
                                                    _ = v24
                                                }
                                            case 127:
                                                v21 = v21
                                            case 1:
                                                var v25 uint64 = 95
                                                _ = v25
                                            }
                                        }
                                        _ = v21
                                    }
                                    v11 -= (v11 & (v11 & v11))
                                }
                                v11 += v11
                            default:
                                if true {
                                    if ((0.0 < 0.0) || ((50 < 63) && true)) {
                                        if len(v13) >= 2 {
                                            v13 = v13[1:len(v13)]
                                        }
                                        var v26 uint64 = 18446744073709551615
                                        v11 = 700
                                        _ = v26
                                    } else {
                                        if len(v13) > 0 {
                                            v13[0] = (^(fi14))
                                        }
                                    }
                                    if ((-2.25 < ((float32(1.5) - -2.25) - (1.5 * ((0.1 + 1.5) + 3.14159)))) && (!false)) {
                                        v27 := GenS2{F0: 0, F1: fi14}
                                        var v28 int8 = 100
                                        var v29 uint16 = 65535
                                        var v30 float64 = 0.1
                                        v30 = v30
                                        _ = v27
                                        _ = v28
                                        _ = v29
                                        _ = v30
                                    } else {
                                        v11 += v11
                                        switch fi14 {
                                        case 32768:
                                            v31 := GenS3{F0: 577}
                                            _ = v31
                                        default:
                                            fmt.Println(fi14, fi10, v11)
                                        }
                                        v11 += v11
                                    }
                                    fmt.Println(fi10)
                                } else {
                                    if (true || false) {
                                        fmt.Println(fi10, fi14, fi10)
                                    } else {
                                        _ = len(v13)
                                    }
                                    if (fi14 != int(v11)) {
                                        v11 += (^(v11))
                                        var v32 int8 = 100
                                        v32 = (v32 + (v32 | (^(v32))))
                                        v33 := GenS4{F0: 2, F1: (v32 * v32), F2: false}
                                        _ = v32
                                        _ = v33
                                    }
                                    v34 := GenS5{F0: true}
                                    _ = v34
                                }
                            }
                            var v35 uint64 = 0
                            if (!((v11 == v11) || false)) {
                                var v36 int8 = -13
                                var v37 float32 = 1e10
                                _ = v36
                                _ = v37
                            } else {
                                var v38 uint32 = 1
                                v11 = 1
                                v38 -= 4294967295
                                var v39 uint32 = (v38 - genFn0(-191, 414, 42))
                                _ = v38
                                _ = v39
                            }
                            _ = v35
                        }
                        if false {
                            if true {
                                var v40 string = ("foo bar" + ("a" + "hello"))
                                var v41 uint8 = 1
                                v40 = "\n\t"
                                _ = v40
                                _ = v41
                            } else {
                                v11 -= (v11 & v11)
                                fmt.Println(fi10)
                            }
                            fmt.Println(fi14)
                            if true {
                                var v42 int32 = -1
                                var v43 int8 = 42
                                v13 = append(v13, -429)
                                if false {
                                    var v44 int32 = 1
                                    var v45 uint = 128
                                    v13 = append(v13, fi10)
                                    v11 += v11
                                    _ = v44
                                    _ = v45
                                }
                                _ = v42
                                _ = v43
                            }
                        } else {
                            if len(v13) > 0 {
                                v13[0] = fi10
                            }
                        }
                        if true {
                            var v46 float32 = 1.5
                            if len(v13) > 0 {
                                v13[0] = fi14
                            }
                            v11 = v11
                            _ = v46
                        } else {
                            v11 -= v11
                            var v47 int64 = 1
                            _ = v47
                        }
                        fmt.Println(fi14, v11, fi10)
                        _ = fi14
                    }
                    v48 := GenS6{F0: 42, F1: (42 != 42)}
                    _ = v13
                    _ = v48
                }
                var v49 float32 = 1.5
                _ = v49
            }
            _ = fi10
            _ = v11
        }
        var v50 int32 = 2147483647
        _ = v50
    } else {
        if (!(!(false && ((6 > 694) || true)))) {
            var v51 uint = 1
            fmt.Println(v51)
            v51 = 255
            _ = v51
        }
        v52 := []int{32767, -2147483648, 128, 42}
        if true {
            var v53 float32 = (float32(-2.25) * (0.1 * (1e10 / 1e10)))
            v53 += v53
            _ = v53
        } else {
            var v54 int8 = 64
            fmt.Println(v54)
            fmt.Println(v54)
            _ = v54
        }
        _ = v52
    }
    _ = v9
    return 128
}

func main() {
    var v1 uint64 = 846
    var v2 int32 = -120
    v3 := make(map[int]int)
    v4 := []int{-9223372036854775808, 0, 32767, 2147483647}
    var v5 uint = 2
    var v6 int32 = (v2 - v2)
    var v55 int8 = -128
    _ = len(v3)
    v56 := GenS2{F0: 199, F1: genFn1()}
    v57 := GenS2{F0: 2, F1: 945}
    if (int64(v2) >= 9223372036854775807) {
        v2 = (v2 & (v6 ^ 2147483647))
        for fi58 := 0; fi58 < 13; fi58++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi58
        }
    } else {
        var v59 float64 = (float64((float32((3.14159 * (0.1 * 3.14159))) * 1.5)) / 3.14159)
        switch genFn1() {
        case -2147483648:
            var v60 float32 = 0.1
            _ = v60
        case 4294967295:
            for fi61 := 0; fi61 < 9; fi61++ {
                v62 := GenS7{F0: 32767, F1: 700, F2: 612}
                switch genFn1() {
                case 1:
                    fmt.Println(v6)
                case 32767:
                    for fi63 := 0; fi63 < 18; fi63++ {
                        _ = len(v4)
                        _ = fi63
                    }
                default:
                    for fi64 := 0; fi64 < 10; fi64++ {
                        v5 -= v5
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = fi64
                    }
                }
                if (false || (1 != 0)) {
                    fmt.Println(v55, v1, v55)
                    if (!false) {
                        var v65 uint32 = 347
                        switch fi61 {
                        case 680:
                            var v66 uint8 = 187
                            _ = v66
                        default:
                            v4 = append(v4, -566)
                        }
                        _ = v65
                    } else {
                        fmt.Println(v5)
                    }
                } else {
                    var v67 float32 = -2.25
                    _ = v67
                }
                _ = fi61
                _ = v62
            }
        }
        _ = v59
    }
    if len(v4) > 0 {
        v4[0] = int(v1)
    }
    var v68 int = 32767
    v6 -= v6
    for fi69 := 0; fi69 < 6; fi69++ {
        v2 -= v2
        _ = fi69
    }
    for fi70 := 0; fi70 < 4; fi70++ {
        var v71 int32 = v2
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi70
        _ = v71
    }
    if false {
        for fi72 := 0; fi72 < 17; fi72++ {
            v73 := GenS8{F0: (^(v6))}
            for fi74 := 0; fi74 < 1; fi74++ {
                var v75 int64 = 888
                var v76 uint8 = 255
                fmt.Println(v5, v6, fi74)
                _ = fi74
                _ = v75
                _ = v76
            }
            for fi77 := 0; fi77 < 3; fi77++ {
                v78 := v3[genFn1()]
                v4 = append(v4, genFn1())
                _ = fi77
                _ = v78
            }
            _ = fi72
            _ = v73
        }
        var v79 float32 = 1.5
        _ = v79
    } else {
        var v80 string = "hello"
        var v81 float32 = 3.14159
        _ = v80
        _ = v81
    }
    v82 := GenS3{F0: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("v6:", v6)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v56
    _ = v57
    _ = v6
    _ = v68
    _ = v82
}
