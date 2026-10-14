package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 int; F1 uint16; F2 int16 }
type GenS4 struct { F0 uint64; F1 int16 }
type GenS5 struct { F0 uint32; F1 bool; F2 uint16 }
func genFn0(p0 int8, p1 uint8, p2 uint64) int8 {
    var v7 float32 = 0.0
    v8 := []int{32767, 2147483647, -1}
    v9 := GenS1{F0: (p0 ^ -8)}
    v10 := GenS1{F0: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return (p0 << ((2) & 63))
}
func genFn1(p0 uint64, p1 int32, p2 int32) bool {
    switch 32767 {
    case 1:
        for fi11 := 0; fi11 < 16; fi11++ {
            switch "hello" {
            case "\n\t":
                v12 := GenS2{F0: (!("foo bar" < ((("foo bar" + "foo bar") + "hello") + ""))), F1: 1}
                _ = v12
            default:
                p0 += 268
            }
            v13 := GenS3{F0: (^(fi11)), F1: 478, F2: 0}
            var v14 int16 = 42
            _ = fi11
            _ = v13
            _ = v14
        }
    case 847:
        p0 -= (p0 << ((0) & 63))
    }
    p1 = 42
    _ = p0
    _ = p1
    _ = p2
    return (0 != 433)
}

func main() {
    var v1 bool = true
    var v2 uint16 = 255
    v3 := make(map[int]int)
    v4 := []string{"foo bar", ("\n\t" + "\n\t")}
    var v5 uint16 = v2
    var v6 string = ("a" + ("a" + ("\n\t" + "\n\t")))
    var v15 int64 = 1
    fmt.Println(v15, v6)
    if false {
        _ = len(v4)
    }
    v16 := GenS1{F0: 127}
    v17 := GenS3{F0: 715, F1: v2, F2: 245}
    var v18 int64 = 761
    for fi19 := 0; fi19 < 7; fi19++ {
        v2 += 363
        if v1 {
            v1 = (1 < 2)
        } else {
            fmt.Println(v1)
            fmt.Println(fi19)
            switch (v6 + "a") {
            case "":
                for fi20 := 0; fi20 < 2; fi20++ {
                    v5 -= 32767
                    v6 = ""
                    if v1 {
                        v1 = v1
                        v3[(fi19 % ((fi20) | 1))] = (^(fi20))
                        v21 := GenS4{F0: 18446744073709551615, F1: -84}
                        _ = v21
                    }
                    _ = fi20
                }
            case "hello":
                v2 = 42
            case "\n\t":
                if v1 {
                    v5 += (v5 << ((1) & 63))
                    switch (^(fi19)) {
                    case 127:
                        var v22 int16 = 32767
                        _ = v22
                    case 32767:
                        v4 = append(v4, v6)
                    case 9223372036854775807:
                        if v1 {
                            if (0.1 <= (1.5 + (1e10 + 0.1))) {
                                switch fi19 {
                                case 42:
                                    for fi23 := 0; fi23 < 13; fi23++ {
                                        if true {
                                            if len(v4) > 0 {
                                                v4[0] = ("foo bar" + v6)
                                            }
                                            v24 := v3[(fi23 << ((42) & 63))]
                                            if v1 {
                                                v25 := GenS2{F0: v1, F1: 877}
                                                switch 648 {
                                                case 76:
                                                    v3[v24] = (-(v24))
                                                case 2147483647:
                                                    fmt.Println(v6, v5)
                                                case 9223372036854775807:
                                                    var v26 float32 = (1e10 - 0.1)
                                                    _ = v26
                                                default:
                                                    switch v24 {
                                                    case 4294967295:
                                                        v15 -= 225
                                                    }
                                                }
                                                v3[(fi23 >> ((65535) & 63))] = (fi19 % ((fi23) | 1))
                                                _ = v25
                                            } else {
                                                var v27 int16 = int16(v24)
                                                _ = v27
                                            }
                                            _ = v24
                                        }
                                        _ = fi23
                                    }
                                case 32767:
                                    var v28 int64 = v15
                                    _ = v28
                                case -2147483648:
                                    for fi29 := 0; fi29 < 17; fi29++ {
                                        var v30 int32 = 1
                                        if len(v4) > 0 {
                                            v4[0] = (v6 + v6)
                                        }
                                        _ = fi29
                                        _ = v30
                                    }
                                default:
                                    for fi31 := 0; fi31 < 13; fi31++ {
                                        var v32 float32 = (float32(1e10) - (3.14159 + -2.25))
                                        var v33 int = fi19
                                        _ = fi31
                                        _ = v32
                                        _ = v33
                                    }
                                }
                                _ = len(v3)
                            }
                            switch 4294967295 {
                            case 127:
                                for fi34 := 0; fi34 < 10; fi34++ {
                                    var v35 float32 = (0.0 - (float32(1e10) - 1.5))
                                    var v36 bool = true
                                    if (-223 == 180) {
                                        if (-2147483648 == 1) {
                                            var v37 int8 = genFn0(70, 255, 1)
                                            switch (fi34 & 639) {
                                            case 127:
                                                var v38 uint8 = 63
                                                _ = v38
                                            }
                                            v39 := GenS2{F0: v1, F1: -32768}
                                            var v40 int16 = 295
                                            v41 := GenS4{F0: 42, F1: v40}
                                            _ = v37
                                            _ = v39
                                            _ = v40
                                            _ = v41
                                        }
                                        var v42 float64 = 1.5
                                        if (genFn0(100, 2, 996) <= -43) {
                                            fmt.Println(v1)
                                            v6 = v6
                                        }
                                        _ = v42
                                    }
                                    _ = fi34
                                    _ = v35
                                    _ = v36
                                }
                            case 892:
                                if v1 {
                                    switch 2147483647 {
                                    case 883:
                                        var v43 float64 = (1e10 / 3.14159)
                                        _ = v43
                                    }
                                    var v44 uint32 = 1
                                    if (!(4294967295 <= 0)) {
                                        v2 = (v2 << ((42) & 63))
                                        var v45 int32 = -537
                                        v5 = v5
                                        _ = v45
                                    } else {
                                        if len(v4) > 0 {
                                            v4[0] = "hello"
                                        }
                                        var v46 uint = 65535
                                        for fi47 := 0; fi47 < 2; fi47++ {
                                            v1 = v1
                                            v48 := v3[fi47]
                                            if true {
                                                var v49 uint32 = 2147483648
                                                v6 = v6
                                                switch (v6 + v6) {
                                                case "\n\t":
                                                    switch 94 {
                                                    case 581:
                                                        v50 := GenS4{F0: 1, F1: 1}
                                                        _ = v50
                                                    default:
                                                        var v51 uint32 = (v44 & 2)
                                                        _ = v51
                                                    }
                                                    var v52 float32 = (((1e10 / -2.25) + -2.25) - ((0.1 * (-2.25 * 1.5)) + float32(float64(float32(0.1)))))
                                                    _ = v52
                                                case "hello":
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                case "foo bar":
                                                    var v53 uint32 = 2147483647
                                                    _ = v53
                                                default:
                                                    v46 = 680
                                                }
                                                _ = v49
                                            } else {
                                                v54 := GenS4{F0: 18446744073709551615, F1: 1}
                                                v55 := GenS5{F0: (^(v44)), F1: ((v2 + (v2 << ((v46) & 63))) >= v2), F2: 1}
                                                var v56 float32 = ((((0.1 + (((float32(0.1) + 0.1) - -2.25) - (float32(((1.5 - 0.1) / (0.1 + (0.0 * (((float64(1e10) / 1.5) * 0.1) / 1e10))))) + 3.14159))) - 0.1) - (float32(1e10) - (float32(0.1) + (float32(0.1) * -2.25)))) - 3.14159)
                                                _ = v54
                                                _ = v55
                                                _ = v56
                                            }
                                            _ = fi47
                                            _ = v48
                                        }
                                        _ = v46
                                    }
                                    _ = v44
                                }
                            case 32767:
                                v57 := v3[fi19]
                                _ = v57
                            default:
                                v2 = (v2 & v5)
                            }
                        }
                    }
                } else {
                    var v58 uint64 = 287
                    v59 := GenS3{F0: fi19, F1: (v5 | 2), F2: -59}
                    _ = v58
                    _ = v59
                }
            default:
                switch (v6 + (v6 + "hello")) {
                case "foo bar":
                    var v60 int64 = (v15 ^ -200)
                    _ = v60
                }
            }
        }
        _ = fi19
    }
    v61 := GenS4{F0: 1, F1: 1000}
    var v62 uint16 = v5
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v61
    _ = v62
}
