package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int8; F1 uint16; F2 int8 }
type GenS5 struct { F0 uint; F1 uint32; F2 uint8 }
type GenS6 struct { F0 uint8; F1 int }
type GenS7 struct { F0 uint32; F1 int }
type GenS8 struct { F0 uint64; F1 uint; F2 string }

func main() {
    var v1 int16 = -1
    var v2 uint64 = 1
    var v3 int = 4294967295
    v4 := []string{"\n\t", "a", "hello", "hello"}
    var v5 uint64 = v2
    v6 := make(map[string]uint32)
    var v7 int32 = 2147483647
    if false {
        v6["a"] = 65535
        if true {
            v8 := GenS1{F0: false}
            if true {
                var v9 string = ("\n\t" + ("hello" + "hello"))
                if (-513 < 88) {
                    for fi10 := 0; fi10 < 1; fi10++ {
                        _ = len(v4)
                        if false {
                            fmt.Println(v5, fi10)
                            for fi11 := 0; fi11 < 6; fi11++ {
                                v9 = (v9 + v9)
                                var v12 int64 = 42
                                switch -371 {
                                case 1:
                                    fmt.Println(v1)
                                default:
                                    v1 = int16(v7)
                                }
                                _ = fi11
                                _ = v12
                            }
                            if (v2 == (^(v2))) {
                                v1 -= (v1 | (-(v1)))
                                v3 -= 175
                                var v13 uint8 = 226
                                _ = v13
                            }
                        } else {
                            v7 = (v7 / (((v7 - (v7 * (v7 / (((v7 & -1)) | 1))))) | 1))
                            v7 = -2147483648
                        }
                        v14 := GenS2{F0: 501, F1: false}
                        _ = fi10
                        _ = v14
                    }
                    var v15 string = (v9 + v9)
                    _ = v15
                } else {
                    v3 = v3
                }
                for fi16 := 0; fi16 < 7; fi16++ {
                    v7 -= int32(v1)
                    _ = fi16
                }
                _ = v9
            } else {
                v7 = (v7 >> ((0) & 63))
                v17 := GenS2{F0: 2, F1: ((v7 | (v7 ^ (v7 + v7))) != v7)}
                if ((v7 + v7) > v7) {
                    fmt.Println(v7)
                    v2 -= (^(v5))
                    v2 += (v5 + (v5 * v2))
                } else {
                    switch "" {
                    case "":
                        v3 -= -1
                    case "\n\t":
                        var v18 int16 = -32768
                        _ = v18
                    default:
                        if (!((1.5 <= 1.5) || false)) {
                            if true {
                                v1 += (v1 << ((2) & 63))
                                v7 = (v7 % ((42) | 1))
                                switch (v3 - (-(v3))) {
                                case 128:
                                    v5 += v2
                                case -1:
                                    v19 := GenS3{F0: 0}
                                    _ = v19
                                case 1:
                                    v6["foo bar"] = 2147483647
                                default:
                                    v20 := GenS4{F0: -76, F1: 1, F2: 127}
                                    _ = v20
                                }
                            }
                            fmt.Println(v1, v2, v7)
                            var v21 bool = (255 > 42)
                            _ = v21
                        }
                    }
                    v3 += -1
                }
                _ = v17
            }
            _ = v8
        }
    } else {
        var v22 int16 = v1
        for fi23 := 0; fi23 < 10; fi23++ {
            v24 := GenS5{F0: 2, F1: 764, F2: 255}
            _ = fi23
            _ = v24
        }
        if true {
            if false {
                for fi25 := 0; fi25 < 14; fi25++ {
                    v3 = (fi25 & (-(v3)))
                    if true {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                    _ = fi25
                }
                v3 = v3
                switch 127 {
                case 1:
                    if ((("\n\t" < "foo bar") && true) && false) {
                        v3 = (v3 >> ((255) & 63))
                    } else {
                        switch v3 {
                        case 32767:
                            for fi26 := 0; fi26 < 15; fi26++ {
                                v27 := GenS5{F0: 2, F1: 367, F2: 241}
                                var v28 int = (-(fi26))
                                _ = fi26
                                _ = v27
                                _ = v28
                            }
                        default:
                            var v29 int16 = (v22 + -1)
                            _ = v29
                        }
                        fmt.Println(v7)
                    }
                }
            } else {
                var v30 int64 = 151
                for fi31 := 0; fi31 < 12; fi31++ {
                    var v32 int16 = (v22 >> ((255) & 63))
                    delete(v6, "")
                    v33 := GenS6{F0: 128, F1: (fi31 * (-(v3)))}
                    _ = fi31
                    _ = v32
                    _ = v33
                }
                var v34 uint32 = 0
                _ = v30
                _ = v34
            }
        }
        _ = v22
    }
    if (false || false) {
        v2 += 919
    } else {
        v1 = 42
        var v35 float64 = (0.1 - (-2.25 / 3.14159))
        _ = v35
    }
    var v36 bool = true
    var v37 uint64 = v5
    var v38 uint32 = 4294967295
    var v39 string = "\n\t"
    if (65535 != 65535) {
        if ((-2.25 * 0.0) != -2.25) {
            var v40 int8 = 0
            _ = v40
        }
    }
    for fi41 := 0; fi41 < 13; fi41++ {
        for fi42 := 0; fi42 < 18; fi42++ {
            v43 := GenS6{F0: 255, F1: (fi41 | (fi41 + v3))}
            if len(v4) > 0 {
                v4[0] = (v39 + "")
            }
            _ = fi42
            _ = v43
        }
        for fi44 := 0; fi44 < 9; fi44++ {
            switch fi44 {
            case 1:
                var v45 int = v3
                _ = v45
            case 128:
                fmt.Println(v38, v2, v38)
            case 0:
                switch fi44 {
                case 32768:
                    v36 = false
                case 955:
                    if v36 {
                        v7 = v7
                        var v46 int8 = 42
                        v47 := GenS7{F0: (v38 / (((v38 & v38)) | 1)), F1: (fi41 % ((fi41) | 1))}
                        _ = v46
                        _ = v47
                    }
                case 32767:
                    if (!v36) {
                        var v48 int = v3
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        var v49 string = (v39 + v39)
                        v7 += 42
                        v1 = v1
                        _ = v48
                        _ = v49
                    } else {
                        v4 = append(v4, "\n\t")
                        var v50 int = v3
                        var v51 uint = 775
                        fmt.Println(fi44)
                        var v52 bool = (1 <= int64(v2))
                        _ = v50
                        _ = v51
                        _ = v52
                    }
                }
                v5 -= (v5 - 0)
            }
            var v53 float64 = 0.1
            switch v39 {
            case "":
                v36 = ((v7 * v7) != int32(v1))
            }
            _ = fi44
            _ = v53
        }
        if (v39 < (v39 + v39)) {
            for fi54 := 0; fi54 < 9; fi54++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                if true {
                    v6[("foo bar" + v39)] = 1
                    if (127 == uint8(v38)) {
                        var v55 uint = 42
                        var v56 int16 = (v1 | 529)
                        _ = v55
                        _ = v56
                    }
                } else {
                    var v57 uint8 = 0
                    if v36 {
                        if false {
                            fmt.Println(v37)
                            v39 = v39
                        }
                    }
                    if v36 {
                        fmt.Println(v7, v7)
                    } else {
                        v1 = -458
                    }
                    _ = v57
                }
                fmt.Println(v38)
                _ = fi54
            }
            v58 := GenS4{F0: int8(v5), F1: 42, F2: -1}
            var v59 uint32 = (v38 | 2)
            _ = v58
            _ = v59
        }
        _ = fi41
    }
    v60 := GenS8{F0: v37, F1: uint(v2), F2: (v39 + "hello")}
    fmt.Println(v36, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v60
    _ = v7
}
