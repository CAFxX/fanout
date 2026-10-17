package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 uint32; F1 uint8 }
type GenS4 struct { F0 int8; F1 string; F2 int8 }
type GenS5 struct { F0 string }
func genFn0(p0 uint, p1 int8, p2 uint) uint {
    p1 += (p1 - p1)
    v5 := []int{32768, -9223372036854775808, -664, 434}
    v5 = append(v5, 0)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return (p0 << ((p0) & 63))
}

func main() {
    var v1 uint8 = 193
    var v2 uint8 = 47
    var v3 uint32 = 2
    v4 := make(map[uint32]string)
    v6 := []int{-923, 127}
    for fi7 := 0; fi7 < 11; fi7++ {
        v2 = 185
        if false {
            v1 += v1
        } else {
            v3 = (v3 | 859)
            var v8 uint16 = 2
            var v9 string = "\n\t"
            _ = v8
            _ = v9
        }
        v3 = v3
        _ = fi7
    }
    v10 := v4[v3]
    for fi11 := 0; fi11 < 8; fi11++ {
        for fi12 := 0; fi12 < 8; fi12++ {
            v13 := GenS1{F0: 627, F1: 47}
            _ = fi12
            _ = v13
        }
        v14 := GenS1{F0: 42, F1: 52}
        v10 = ("a" + ((v10 + v10) + (v10 + v10)))
        _ = fi11
        _ = v14
    }
    v2 += (v1 ^ v2)
    v6 = append(v6, 9223372036854775807)
    v15 := GenS1{F0: 1, F1: 0}
    v10 = ""
    v1 -= (v2 * v2)
    for fi16 := 0; fi16 < 8; fi16++ {
        var v17 int32 = 1
        v1 += v1
        var v18 bool = ("a" > v10)
        _ = fi16
        _ = v17
        _ = v18
    }
    v3 += (v3 + (v3 % ((0) | 1)))
    for fi19 := 0; fi19 < 3; fi19++ {
        if (!(1 > 9223372036854775807)) {
            if (true && ((42 == 404) && true)) {
                if (3.14159 < 0.0) {
                    if (((float64((1e10 * 0.0)) <= 0.0) && true) && ((v10 + ((v10 + "") + v10)) == "a")) {
                        var v20 int = fi19
                        _ = v20
                    }
                    v1 -= v1
                    var v21 uint32 = (^(v3))
                    _ = v21
                } else {
                    _ = len(v6)
                }
            }
        } else {
            v1 -= v2
            v22 := GenS2{F0: v2, F1: -96}
            _ = v22
        }
        for fi23 := 0; fi23 < 5; fi23++ {
            if false {
                v24 := v4[(v3 * (v3 >> ((0) & 63)))]
                var v25 uint64 = 0
                _ = v24
                _ = v25
            }
            v4[v3] = v10
            var v26 int16 = -422
            v27 := GenS2{F0: v2, F1: 127}
            v10 = v10
            var v28 int8 = int8(v2)
            switch int(v28) {
            case 1:
                v26 = 32767
                var v29 uint8 = (v1 % ((0) | 1))
                _ = v29
            }
            v10 = v10
            _ = fi23
            _ = v26
            _ = v27
            _ = v28
        }
        _ = fi19
    }
    var v30 int8 = 11
    v1 -= (v2 - v2)
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v31 := GenS1{F0: uint64(v3), F1: -128}
    if false {
        fmt.Println(v30, v30, v2)
        _ = len(v4)
    } else {
        for fi32 := 0; fi32 < 4; fi32++ {
            if (float32(-2.25) <= (0.0 - (0.0 + 1e10))) {
                for fi33 := 0; fi33 < 20; fi33++ {
                    if true {
                        if true {
                            _ = len(v6)
                            var v34 int = (fi33 >> ((genFn0(18446744073709551615, v30, 452)) & 63))
                            _ = v34
                        }
                        fmt.Println(v10, v3, v2)
                        var v35 int64 = 42
                        _ = v35
                    }
                    var v36 int32 = 483
                    switch 128 {
                    case -9223372036854775808:
                        v30 = (v30 | v30)
                    case 32768:
                        var v37 string = (v10 + v10)
                        _ = v37
                    default:
                        switch ((v10 + "\n\t") + "") {
                        case "foo bar":
                            if true {
                                v4[(v3 - (v3 + v3))] = "foo bar"
                                v36 -= -536
                                if len(v6) > 0 {
                                    v6[0] = int(v1)
                                }
                            } else {
                                if true {
                                    delete(v4, v3)
                                    if true {
                                        v1 += (v2 % ((v1) | 1))
                                    }
                                } else {
                                    var v38 bool = true
                                    var v39 uint32 = v3
                                    fmt.Println(v10)
                                    _ = v38
                                    _ = v39
                                }
                                if false {
                                    v3 = v3
                                    v40 := GenS3{F0: (^(v3)), F1: (v1 + v1)}
                                    v30 -= (v30 % (((-(v30))) | 1))
                                    _ = v40
                                }
                                v30 -= int8(v36)
                                v36 = 42
                            }
                        default:
                            v41 := GenS2{F0: v2, F1: v30}
                            _ = v41
                        }
                    }
                    _ = fi33
                    _ = v36
                }
                v42 := GenS1{F0: 18446744073709551615, F1: int8(v2)}
                if false {
                    _ = len(v4)
                    if true {
                        v30 = v30
                        if ((v30 | 1) >= -71) {
                            fmt.Println(v1, v10, v2)
                            switch v10 {
                            case "a":
                                v43 := GenS4{F0: v30, F1: (v10 + "\n\t"), F2: v30}
                                _ = v43
                            case "\n\t":
                                fmt.Println(v10, fi32, v30)
                            case "foo bar":
                                for fi44 := 0; fi44 < 12; fi44++ {
                                    if false {
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                    }
                                    if len(v6) > 0 {
                                        v6[0] = (fi44 | (fi44 & (fi44 | fi44)))
                                    }
                                    v45 := GenS1{F0: 1, F1: 127}
                                    _ = fi44
                                    _ = v45
                                }
                            default:
                                for fi46 := 0; fi46 < 11; fi46++ {
                                    v47 := GenS2{F0: v1, F1: (v30 - v30)}
                                    var v48 uint64 = 4294967295
                                    _ = fi46
                                    _ = v47
                                    _ = v48
                                }
                            }
                            switch (fi32 ^ (fi32 << ((42) & 63))) {
                            case -793:
                                switch fi32 {
                                case 127:
                                    for fi49 := 0; fi49 < 8; fi49++ {
                                        fmt.Println(v30)
                                        switch v10 {
                                        case "hello":
                                            switch ("\n\t" + "hello") {
                                            case "foo bar":
                                                fmt.Println(v1, v1)
                                            case "hello":
                                                fmt.Println(v3, v30)
                                            case "\n\t":
                                                var v50 uint64 = 0
                                                _ = v50
                                            }
                                        case "a":
                                            _ = len(v6)
                                        case "\n\t":
                                            if true {
                                                if (((((3.14159 + -2.25) * ((-2.25 / 1e10) + ((0.1 * 1.5) - 3.14159))) - 0.1) - 1e10) >= 1.5) {
                                                    v51 := v4[2147483647]
                                                    v30 = (v30 / ((127) | 1))
                                                    fmt.Println(v2)
                                                    var v52 uint8 = v2
                                                    _ = v51
                                                    _ = v52
                                                }
                                            } else {
                                                v3 = 2
                                                v53 := GenS4{F0: (v30 ^ 0), F1: "", F2: v30}
                                                _ = v53
                                            }
                                        }
                                        switch "\n\t" {
                                        case "":
                                            var v54 int = 32767
                                            _ = v54
                                        case "\n\t":
                                            fmt.Println(v3, v2)
                                        case "a":
                                            v10 = ((v10 + v10) + v10)
                                        }
                                        _ = fi49
                                    }
                                case 81:
                                    for fi55 := 0; fi55 < 1; fi55++ {
                                        _ = len(v6)
                                        v3 -= v3
                                        switch v10 {
                                        case "\n\t":
                                            v3 = v3
                                        case "":
                                            if ((v3 | (v3 & (^(v3)))) >= 65535) {
                                                v1 += 128
                                                fmt.Println(fi32)
                                                v1 += 0
                                            }
                                        case "a":
                                            var v56 int8 = (^(v30))
                                            _ = v56
                                        default:
                                            v57 := GenS5{F0: ("hello" + v10)}
                                            _ = v57
                                        }
                                        _ = fi55
                                    }
                                }
                            case 32768:
                                var v58 int64 = -555
                                _ = v58
                            default:
                                for fi59 := 0; fi59 < 16; fi59++ {
                                    fmt.Println(v3, v10)
                                    _ = fi59
                                }
                            }
                        } else {
                            v4[(v3 | v3)] = ""
                        }
                    } else {
                        for fi60 := 0; fi60 < 15; fi60++ {
                            var v61 uint = 42
                            _ = fi60
                            _ = v61
                        }
                        v10 = v10
                    }
                    fmt.Println(v2)
                }
                _ = v42
            }
            var v62 string = v10
            _ = fi32
            _ = v62
        }
        var v63 bool = true
        fmt.Println(v63, v2)
        _ = v63
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v6
}
