package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint16; F1 int8; F2 int8 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint; F1 int8 }
func genFn0() int64 {
    v5 := GenS1{F0: 4294967295, F1: 1}
    for fi6 := 0; fi6 < 15; fi6++ {
        var v7 float64 = 1e10
        var v8 uint32 = 921
        fmt.Println(fi6, v7, v8)
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 int64 = -1
    _ = v5
    _ = v9
    return v9
}

func main() {
    v1 := make(map[int]int)
    v2 := []string{("\n\t" + ("" + ("foo bar" + ""))), "\n\t"}
    var v3 int32 = 42
    v4 := make(map[uint32]bool)
    fmt.Println(v3)
    fmt.Println(v3)
    v3 -= v3
    v10 := GenS1{F0: 1, F1: -998}
    var v11 int16 = 32767
    for fi12 := 0; fi12 < 6; fi12++ {
        var v13 float32 = float32((1e10 - 1e10))
        for fi14 := 0; fi14 < 6; fi14++ {
            if true {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                var v15 string = (((("" + "a") + "foo bar") + ("\n\t" + "")) + "hello")
                v13 = v13
                switch (fi12 >> ((30) & 63)) {
                case 718:
                    var v16 float64 = ((float64(float32((0.1 * (0.0 * float64(v13))))) * float64(((float32(float64(0.0)) * (3.14159 / v13)) / 1.0))) / 1.5)
                    _ = v16
                }
                delete(v1, (fi14 << ((65535) & 63)))
                var v17 int8 = -1
                _ = v15
                _ = v17
            }
            _ = fi14
        }
        _ = fi12
        _ = v13
    }
    v18 := GenS1{F0: 255, F1: 0}
    v2 = append(v2, "hello")
    var v19 int32 = -2147483648
    if (!(!true)) {
        var v20 int32 = (v19 + -2147483648)
        switch 32768 {
        case -907:
            if true {
                v11 = (v11 << ((223) & 63))
            } else {
                var v21 int64 = -1
                v1[-363] = -797
                _ = v21
            }
        default:
            v22 := GenS2{F0: 0}
            _ = v22
        }
        v19 -= v3
        _ = v20
    } else {
        if (float64((3.14159 + (3.14159 + 1e10))) >= ((-2.25 - ((((1.5 * 0.1) * 0.0) - -2.25) / 0.1)) / (0.1 * 0.1))) {
            for fi23 := 0; fi23 < 5; fi23++ {
                delete(v4, 2147483648)
                var v24 uint = 18446744073709551615
                _ = fi23
                _ = v24
            }
            if ((226 == 42) || true) {
                if false {
                    v3 -= (v19 & v3)
                    if ((42 != 11) || (65535 != 569)) {
                        fmt.Println(v19, v11)
                        var v25 uint64 = 1
                        _ = v25
                    } else {
                        v26 := GenS1{F0: 11, F1: 269}
                        v19 -= 42
                        _ = v26
                    }
                } else {
                    var v27 uint8 = 1
                    _ = v27
                }
                for fi28 := 0; fi28 < 7; fi28++ {
                    if len(v2) > 0 {
                        v2[0] = ("\n\t" + ("foo bar" + "\n\t"))
                    }
                    _ = fi28
                }
            } else {
                if (4294967295 <= 2) {
                    delete(v4, 513)
                    fmt.Println(v19)
                } else {
                    var v29 uint32 = 0
                    if len(v2) > 0 {
                        v2[0] = "hello"
                    }
                    _ = v29
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                switch 127 {
                case 32767:
                    var v30 int16 = (v11 << ((42) & 63))
                    _ = v30
                case 2147483647:
                    var v31 uint16 = 42
                    _ = v31
                case 4294967295:
                    var v32 int64 = 1
                    _ = v32
                }
            }
            v19 = 0
        } else {
            if ((true || ((v3 >= (v3 / ((v3) | 1))) && true)) || ((23 < 0) && false)) {
                v11 -= v11
                for fi33 := 0; fi33 < 15; fi33++ {
                    if (!true) {
                        fmt.Println(v11, fi33, v3)
                    }
                    var v34 int16 = 986
                    switch "foo bar" {
                    case "foo bar":
                        v3 = 0
                    case "":
                        v11 = v11
                    case "a":
                        v35 := GenS2{F0: 766}
                        _ = v35
                    default:
                        switch "a" {
                        case "":
                            v36 := GenS1{F0: 763, F1: 9223372036854775807}
                            _ = v36
                        case "\n\t":
                            v19 -= -1
                        default:
                            switch "foo bar" {
                            case "foo bar":
                                v34 += v34
                            case "\n\t":
                                fmt.Println(v11)
                            default:
                                v19 -= 879
                            }
                        }
                    }
                    _ = fi33
                    _ = v34
                }
                v37 := GenS3{F0: 4294967295}
                _ = v37
            }
        }
        v38 := GenS2{F0: 2}
        _ = v38
    }
    for fi39 := 0; fi39 < 18; fi39++ {
        var v40 string = ""
        fmt.Println(fi39)
        switch ("" + (("foo bar" + (((v40 + (v40 + "hello")) + v40) + v40)) + v40)) {
        case "hello":
            if false {
                switch (((v40 + v40) + "foo bar") + (v40 + "hello")) {
                case "hello":
                    v40 = v40
                default:
                    v1[(fi39 / (((fi39 & (fi39 / ((fi39) | 1)))) | 1))] = -836
                }
                v40 = (("hello" + v40) + v40)
                switch v40 {
                case "foo bar":
                    var v41 string = "\n\t"
                    _ = v41
                case "hello":
                    for fi42 := 0; fi42 < 13; fi42++ {
                        if true {
                            switch "a" {
                            case "a":
                                var v43 int32 = v3
                                switch (fi42 & 32768) {
                                case 32767:
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                case 723:
                                    v3 -= v3
                                case 32768:
                                    switch v40 {
                                    case "\n\t":
                                        fmt.Println(v11, v19)
                                    case "":
                                        v1[(fi39 * fi42)] = 32768
                                    case "hello":
                                        fmt.Println(v40)
                                        if len(v2) > 0 {
                                            v2[0] = (v40 + "")
                                        }
                                        v44 := GenS4{F0: 255, F1: 42, F2: -22}
                                        _ = v44
                                    }
                                default:
                                    if true {
                                        var v45 uint = 255
                                        fmt.Println(v11)
                                        _ = v45
                                    } else {
                                        v46 := v4[1]
                                        var v47 int = fi39
                                        v48 := GenS5{F0: v46}
                                        _ = v46
                                        _ = v47
                                        _ = v48
                                    }
                                }
                                v19 += 969
                                _ = v43
                            case "":
                                v40 = v40
                            case "hello":
                                switch v40 {
                                case "\n\t":
                                    v11 -= v11
                                case "foo bar":
                                    v2 = append(v2, v40)
                                case "":
                                    if (0.1 < (-2.25 - 1.5)) {
                                        var v49 int64 = -9223372036854775808
                                        var v50 int64 = genFn0()
                                        _ = v49
                                        _ = v50
                                    } else {
                                        _ = len(v4)
                                    }
                                default:
                                    v11 += v11
                                }
                            default:
                                v51 := GenS2{F0: 65535}
                                _ = v51
                            }
                            var v52 int16 = v11
                            _ = v52
                        } else {
                            fmt.Println(fi39, fi39)
                            v11 += 32767
                        }
                        _ = fi42
                    }
                default:
                    if true {
                        if (!(fi39 != 32767)) {
                            var v53 int = (fi39 ^ -1)
                            _ = v53
                        }
                        v54 := GenS1{F0: 2, F1: 0}
                        _ = v54
                    } else {
                        v3 = 792
                        if true {
                            v40 = v40
                            for fi55 := 0; fi55 < 16; fi55++ {
                                if (178 != 0) {
                                    var v56 int = 4294967295
                                    fmt.Println(v40, v56)
                                    var v57 int64 = 1
                                    _ = v56
                                    _ = v57
                                }
                                var v58 int64 = 0
                                _ = fi55
                                _ = v58
                            }
                            var v59 int32 = v19
                            _ = v59
                        }
                    }
                }
            }
        default:
            for fi60 := 0; fi60 < 10; fi60++ {
                v61 := GenS6{F0: 2, F1: 105}
                if len(v2) > 0 {
                    v2[0] = v40
                }
                var v62 int32 = (v19 >> ((582) & 63))
                _ = fi60
                _ = v61
                _ = v62
            }
        }
        _ = fi39
        _ = v40
    }
    v63 := GenS5{F0: (2 > 255)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v63
}
