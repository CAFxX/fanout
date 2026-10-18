package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8; F1 uint16; F2 uint16 }
type GenS4 struct { F0 int8; F1 int16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 bool; F1 string; F2 string }
type GenS7 struct { F0 int64; F1 uint64; F2 int64 }
func genFn0(p0 uint64, p1 uint32) int64 {
    v4 := []int{2147483647}
    for fi5 := 0; fi5 < 19; fi5++ {
        var v6 uint = 310
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = p1
    _ = v4
    return 920
}
func genFn1() uint8 {
    v7 := []int{}
    var v8 int8 = 42
    v9 := GenS1{F0: 29, F1: 128, F2: v8}
    _ = v7
    _ = v8
    _ = v9
    return 160
}

func main() {
    var v1 uint8 = 128
    v2 := []string{("\n\t" + ((("\n\t" + ("\n\t" + "\n\t")) + "\n\t") + "a")), "a", "\n\t"}
    var v3 uint = 42
    if true {
        var v10 bool = (((true || (1 <= 4294967295)) && false) || true)
        var v11 uint16 = 65535
        _ = v10
        _ = v11
    } else {
        var v12 int32 = 2147483647
        if len(v2) > 0 {
            v2[0] = (("a" + "a") + "hello")
        }
        _ = v12
    }
    var v13 uint16 = 32767
    for fi14 := 0; fi14 < 6; fi14++ {
        fmt.Println(v13, fi14, v3)
        _ = fi14
    }
    v15 := GenS1{F0: -1, F1: 697, F2: 0}
    v16 := GenS1{F0: 1000, F1: 816, F2: 1}
    var v17 uint32 = 42
    v18 := GenS1{F0: 1, F1: int(v3), F2: -1}
    if (-145 <= -247) {
        if (v17 <= v17) {
            if false {
                if (!true) {
                    var v19 bool = ((32767 <= 2147483647) || (!true))
                    var v20 uint64 = 600
                    _ = v19
                    _ = v20
                }
            } else {
                v1 = 1
                v21 := GenS2{F0: 1000}
                _ = v21
            }
            v1 += (v1 >> ((2) & 63))
        }
    }
    v22 := GenS2{F0: -339}
    v1 = genFn1()
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v23 string = "foo bar"
    var v24 uint32 = (v17 + (v17 / ((42) | 1)))
    for fi25 := 0; fi25 < 10; fi25++ {
        v24 -= (^(v17))
        var v26 bool = true
        if v26 {
            v17 -= (^(v17))
            if false {
                if (!(!v26)) {
                    for fi27 := 0; fi27 < 2; fi27++ {
                        v26 = v26
                        v3 = v3
                        _ = fi27
                    }
                } else {
                    v28 := GenS3{F0: (v1 | (v1 % (((v1 >> ((255) & 63))) | 1))), F1: v13, F2: (v13 << (((v3 % (((v3 + uint(v13))) | 1))) & 63))}
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = v28
                }
                for fi29 := 0; fi29 < 7; fi29++ {
                    v30 := GenS1{F0: 42, F1: (fi25 % ((-596) | 1)), F2: 0}
                    v1 = (v1 * v1)
                    v13 += (v13 + v13)
                    _ = fi29
                    _ = v30
                }
            } else {
                switch v23 {
                case "\n\t":
                    var v31 int16 = 1
                    _ = v31
                case "foo bar":
                    if (0.1 >= ((0.1 / -2.25) / 1.5)) {
                        v26 = (42 == 1)
                    } else {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        var v32 int16 = -84
                        _ = v32
                    }
                default:
                    v13 -= v13
                }
                if len(v2) > 0 {
                    v2[0] = "a"
                }
                switch fi25 {
                case 42:
                    switch ("foo bar" + (v23 + v23)) {
                    case "foo bar":
                        v33 := GenS3{F0: (v1 % (((v1 | 255)) | 1)), F1: (v13 % ((161) | 1)), F2: (^(v13))}
                        _ = v33
                    default:
                        v2 = append(v2, (v23 + v23))
                    }
                case 2147483647:
                    v24 = 65535
                default:
                    switch -180 {
                    case -584:
                        v3 = uint(v13)
                    case 32768:
                        v26 = ((v1 * genFn1()) <= (v1 << ((2) & 63)))
                    case 79:
                        v34 := GenS4{F0: 51, F1: -563}
                        _ = v34
                    default:
                        _ = len(v2)
                    }
                }
            }
            if v26 {
                v13 += v13
                switch v23 {
                case "":
                    var v35 int64 = genFn0(42, v24)
                    _ = v35
                case "\n\t":
                    var v36 uint64 = 18446744073709551615
                    _ = v36
                default:
                    if v26 {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        v24 -= (^(v17))
                        switch v23 {
                        case "a":
                            var v37 float32 = -2.25
                            _ = v37
                        case "\n\t":
                            var v38 int64 = 1
                            _ = v38
                        default:
                            for fi39 := 0; fi39 < 18; fi39++ {
                                v40 := GenS5{F0: ""}
                                v41 := GenS6{F0: v26, F1: ((v23 + "a") + "foo bar"), F2: (("\n\t" + "") + ((("\n\t" + v23) + "") + ((v23 + (v23 + v23)) + ("" + (v23 + ("\n\t" + "\n\t"))))))}
                                _ = fi39
                                _ = v40
                                _ = v41
                            }
                        }
                    } else {
                        if (100 != 42) {
                            for fi42 := 0; fi42 < 12; fi42++ {
                                v26 = v26
                                fmt.Println(v26, v26, v26)
                                _ = fi42
                            }
                            for fi43 := 0; fi43 < 8; fi43++ {
                                if len(v2) > 0 {
                                    v2[0] = v23
                                }
                                _ = fi43
                            }
                            v17 += v17
                        } else {
                            var v44 int8 = -128
                            _ = v44
                        }
                        v45 := GenS2{F0: -77}
                        _ = v45
                    }
                }
                v46 := GenS7{F0: 42, F1: 0, F2: -342}
                _ = v46
            }
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v17 = uint32(v1)
            for fi47 := 0; fi47 < 8; fi47++ {
                var v48 int = fi25
                var v49 uint16 = v13
                v23 = (("a" + "foo bar") + v23)
                _ = fi47
                _ = v48
                _ = v49
            }
        }
        _ = fi25
        _ = v26
    }
    for fi50 := 0; fi50 < 4; fi50++ {
        var v51 string = (("foo bar" + (v23 + v23)) + v23)
        v51 = (v51 + v51)
        _ = len(v2)
        _ = fi50
        _ = v51
    }
    for fi52 := 0; fi52 < 11; fi52++ {
        v13 += v13
        fmt.Println(v3, v23, v17)
        if ("" > "") {
            var v53 uint64 = 2
            var v54 uint8 = v1
            for fi55 := 0; fi55 < 18; fi55++ {
                var v56 uint16 = v13
                _ = fi55
                _ = v56
            }
            _ = v53
            _ = v54
        }
        _ = fi52
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
}
