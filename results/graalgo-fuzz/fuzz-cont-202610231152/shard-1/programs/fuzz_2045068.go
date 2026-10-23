package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 uint }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 uint32; F1 bool; F2 uint64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int16 }
func genFn0() int {
    v7 := []int{42, -532}
    var v8 int32 = 254
    v8 = (v8 & v8)
    _ = v7
    _ = v8
    return 32767
}
func genFn1(p0 int32, p1 uint) bool {
    var v9 float32 = 0.0
    p1 -= p1
    _ = p0
    _ = p1
    _ = v9
    return (86 <= -1)
}

func main() {
    v1 := make(map[int]bool)
    v2 := []int{2147483647, 128}
    v3 := []uint32{42, 2147483648, 772, 2147483647}
    var v4 float32 = ((0.0 + 0.1) * ((3.14159 / (1.5 + 1.5)) * 1.5))
    var v5 uint64 = 4294967295
    var v6 int8 = 42
    switch -1 {
    case -716:
        var v10 int16 = -1
        _ = v10
    case -219:
        var v11 float64 = ((3.14159 / (0.1 + (3.14159 + (3.14159 * 0.1)))) / (1.5 / float64(v4)))
        _ = v11
    case 162:
        var v12 uint16 = 255
        _ = v12
    }
    var v13 uint16 = 1
    v14 := GenS1{F0: 246, F1: (true && false), F2: 18446744073709551615}
    var v15 bool = false
    fmt.Println(v15)
    v6 -= -27
    v16 := GenS1{F0: -685, F1: ((3.14159 * 0.1) >= float64(v4)), F2: 18446744073709551615}
    v15 = true
    var v17 int = 9223372036854775807
    if false {
        v18 := GenS1{F0: -1, F1: v15, F2: 18446744073709551615}
        if v15 {
            v19 := GenS1{F0: int16(v5), F1: (v13 <= v13), F2: 0}
            if v15 {
                v6 = 0
            }
            _ = v19
        } else {
            if v15 {
                for fi20 := 0; fi20 < 10; fi20++ {
                    if v15 {
                        v17 += fi20
                    } else {
                        switch "hello" {
                        case "foo bar":
                            if (v15 || (!v15)) {
                                var v21 uint64 = (v5 - (^(v5)))
                                var v22 uint = 42
                                _ = v21
                                _ = v22
                            } else {
                                v23 := GenS1{F0: -1, F1: true, F2: 209}
                                var v24 int = (v17 / ((-9223372036854775808) | 1))
                                if v15 {
                                    if true {
                                        v25 := GenS2{F0: v5, F1: v5}
                                        for fi26 := 0; fi26 < 13; fi26++ {
                                            if ((1 <= -1) || v15) {
                                                v15 = (v15 || (!true))
                                                v24 = fi26
                                                v27 := GenS3{F0: uint32(v13), F1: (float64(float32(0.0)) > float64(v4)), F2: (v5 << ((65535) & 63))}
                                                _ = v27
                                            } else {
                                                fmt.Println(v13, v5, v24)
                                                v28 := GenS3{F0: 0, F1: v15, F2: (v5 & (v5 / ((2) | 1)))}
                                                fmt.Println(v17, v13, v17)
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                _ = v28
                                            }
                                            v15 = (v15 || true)
                                            switch (v17 >> ((2) & 63)) {
                                            case 127:
                                                if v15 {
                                                    var v29 float32 = v4
                                                    v2 = append(v2, fi20)
                                                    _ = v29
                                                }
                                            default:
                                                var v30 uint = 317
                                                _ = v30
                                            }
                                            _ = fi26
                                        }
                                        _ = v25
                                    }
                                }
                                _ = v23
                                _ = v24
                            }
                        case "hello":
                            v5 -= 4294967295
                        }
                        if ((v5 <= v5) || v15) {
                            if ((-(fi20)) < genFn0()) {
                                v13 -= 42
                            }
                            for fi31 := 0; fi31 < 5; fi31++ {
                                v5 = (v5 - (v5 * (^(v5))))
                                v32 := GenS4{F0: (v5 / ((18446744073709551615) | 1))}
                                _ = fi31
                                _ = v32
                            }
                        }
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    _ = fi20
                }
            } else {
                switch v17 {
                case 9223372036854775807:
                    v17 += (v17 - v17)
                case -2147483648:
                    v33 := v1[(-(v17))]
                    _ = v33
                case 32768:
                    fmt.Println(v4)
                default:
                    var v34 uint8 = 128
                    _ = v34
                }
                for fi35 := 0; fi35 < 2; fi35++ {
                    var v36 string = (((("foo bar" + "") + ("foo bar" + ("" + "foo bar"))) + "hello") + ("a" + "foo bar"))
                    var v37 uint64 = v5
                    _ = fi35
                    _ = v36
                    _ = v37
                }
                var v38 int8 = 42
                _ = v38
            }
            fmt.Println(v5, v15)
        }
        switch (((("\n\t" + "a") + "") + ("a" + "foo bar")) + ("" + "a")) {
        case "a":
            fmt.Println(v6, v13)
        case "foo bar":
            for fi39 := 0; fi39 < 18; fi39++ {
                fmt.Println(v15, v15)
                _ = fi39
            }
        case "":
            v40 := GenS3{F0: 558, F1: (0 < 1), F2: (v5 ^ 140)}
            _ = v40
        default:
            if true {
                var v41 uint16 = (v13 >> ((874) & 63))
                if (float32((1e10 / ((1e10 + (0.1 + 0.1)) / ((((-2.25 / 3.14159) * 0.0) + 0.0) - (1.5 + (0.0 + 1.5)))))) < (v4 * (v4 + (v4 / 1.0)))) {
                    var v42 int = v17
                    for fi43 := 0; fi43 < 1; fi43++ {
                        var v44 int = -1
                        _ = fi43
                        _ = v44
                    }
                    _ = v42
                } else {
                    var v45 uint = 780
                    var v46 uint8 = 7
                    v15 = ((("" + "foo bar") + "") <= "a")
                    _ = v45
                    _ = v46
                }
                fmt.Println(v6, v5)
                _ = v41
            } else {
                if v15 {
                    for fi47 := 0; fi47 < 14; fi47++ {
                        var v48 string = "hello"
                        var v49 int64 = -707
                        _ = fi47
                        _ = v48
                        _ = v49
                    }
                    v15 = (v15 && false)
                    v50 := GenS2{F0: v5, F1: (v5 ^ (v5 - 219))}
                    _ = v50
                }
                v15 = (0 < 32767)
            }
        }
        _ = v18
    } else {
        var v51 int64 = 438
        _ = v51
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v52 float32 = 0.1
    v13 += 2
    switch (v17 + -624) {
    case 9223372036854775807:
        v2 = append(v2, v17)
    case 64:
        for fi53 := 0; fi53 < 14; fi53++ {
            v54 := GenS5{F0: -1}
            var v55 int32 = 389
            for fi56 := 0; fi56 < 13; fi56++ {
                if (v15 || (!((v55 <= -799) && (!(float32(((float64(float32(1e10)) + 3.14159) / -2.25)) <= v4))))) {
                    var v57 float64 = 3.14159
                    var v58 bool = v15
                    v59 := v1[fi56]
                    v60 := GenS1{F0: 0, F1: ((true && v59) && v15), F2: 2}
                    _ = v57
                    _ = v58
                    _ = v59
                    _ = v60
                } else {
                    switch "foo bar" {
                    case "hello":
                        var v61 int16 = 1000
                        _ = v61
                    }
                    v62 := GenS2{F0: (v5 | (v5 << ((494) & 63))), F1: v5}
                    _ = v62
                }
                v3 = append(v3, 42)
                fmt.Println(v6)
                _ = fi56
            }
            _ = fi53
            _ = v54
            _ = v55
        }
    case 4294967295:
        if (-2147483648 < -590) {
            v15 = v15
            v15 = (836 > 2147483647)
        }
    }
    for fi63 := 0; fi63 < 1; fi63++ {
        if len(v3) > 0 {
            v3[0] = 308
        }
        fmt.Println(v52, fi63)
        var v64 int16 = 776
        _ = fi63
        _ = v64
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v6
}
