package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 int8 }
type GenS2 struct { F0 bool; F1 int64 }
type GenS3 struct { F0 uint8; F1 uint16; F2 uint16 }
type GenS4 struct { F0 string; F1 uint16; F2 bool }
type GenS5 struct { F0 bool; F1 int8; F2 int64 }
func genFn0(p0 string, p1 int) uint8 {
    p1 += -9223372036854775808
    var v6 string = p0
    p0 = p0
    _ = p0
    _ = p1
    _ = v6
    return 255
}

func main() {
    var v1 bool = false
    v2 := []string{}
    v3 := make(map[int]uint32)
    var v4 string = ("\n\t" + "foo bar")
    var v5 int32 = -1
    for fi7 := 0; fi7 < 11; fi7++ {
        _ = len(v2)
        _ = fi7
    }
    var v8 string = ((v4 + v4) + v4)
    for fi9 := 0; fi9 < 6; fi9++ {
        v1 = (102 != 127)
        _ = len(v2)
        v10 := GenS1{F0: -32768, F1: 42, F2: 15}
        _ = fi9
        _ = v10
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if (65535 == 2) {
        v11 := GenS2{F0: (!v1), F1: -1}
        _ = v11
    } else {
        for fi12 := 0; fi12 < 4; fi12++ {
            if (255 == 237) {
                var v13 uint64 = 0
                v5 -= 0
                for fi14 := 0; fi14 < 1; fi14++ {
                    if v1 {
                        _ = len(v3)
                        v15 := GenS2{F0: false, F1: 42}
                        _ = v15
                    } else {
                        fmt.Println(v13)
                        fmt.Println(fi14)
                        if false {
                            if v1 {
                                v1 = (255 < 65535)
                                var v16 uint64 = 677
                                switch v8 {
                                case "\n\t":
                                    switch (fi14 * (fi14 << ((729) & 63))) {
                                    case 213:
                                        switch 30 {
                                        case -444:
                                            v5 -= int32(v13)
                                        }
                                    case 0:
                                        if false {
                                            v17 := GenS1{F0: 42, F1: -24, F2: 1}
                                            _ = v17
                                        }
                                    case -381:
                                        v18 := v3[(^(fi12))]
                                        _ = v18
                                    default:
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    }
                                case "foo bar":
                                    var v19 uint64 = v16
                                    _ = v19
                                case "":
                                    v1 = (!v1)
                                }
                                _ = v16
                            }
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            v5 -= v5
                            var v20 int8 = -1
                            delete(v3, fi14)
                            if ((("\n\t" + "foo bar") + "hello") > "hello") {
                                v5 = (v5 >> ((65535) & 63))
                            }
                            _ = v20
                        } else {
                            switch "foo bar" {
                            case "a":
                                v1 = false
                            case "hello":
                                v3[(fi12 % ((fi14) | 1))] = 65535
                            }
                            var v21 uint8 = 42
                            v3[fi12] = 123
                            _ = v21
                        }
                    }
                    fmt.Println(fi14, v1, v13)
                    if v1 {
                        if (35 > 65535) {
                            v1 = (32767 <= 255)
                            switch 42 {
                            case 42:
                                if v1 {
                                    v22 := GenS3{F0: 188, F1: 363, F2: 1}
                                    if (-1 > 42) {
                                        v4 = (v4 + (((v4 + v4) + "\n\t") + v4))
                                        if (!v1) {
                                            var v23 uint64 = (v13 >> ((346) & 63))
                                            v1 = v1
                                            v24 := v3[-441]
                                            _ = v23
                                            _ = v24
                                        }
                                        if v1 {
                                            v1 = true
                                            var v25 uint64 = (v13 ^ 1)
                                            _ = v25
                                        }
                                    }
                                    _ = v22
                                } else {
                                    v3[fi12] = 0
                                    if true {
                                        if v1 {
                                            v4 = ""
                                        } else {
                                            if false {
                                                var v26 uint = 1
                                                if (v8 == "\n\t") {
                                                    v26 += 4294967295
                                                    v4 = ((v4 + (("foo bar" + ("\n\t" + ("\n\t" + ((v4 + v4) + v4)))) + ((v4 + v4) + (v8 + v4)))) + ("a" + "\n\t"))
                                                    v5 = 0
                                                }
                                                _ = v26
                                            }
                                        }
                                        fmt.Println(v1, v13, v1)
                                    }
                                }
                            default:
                                v5 -= (v5 ^ 2147483647)
                            }
                            var v27 int8 = 79
                            var v28 int8 = v27
                            var v29 uint8 = 128
                            _ = v27
                            _ = v28
                            _ = v29
                        } else {
                            var v30 int32 = (v5 & (v5 - v5))
                            v8 = v8
                            _ = v30
                        }
                        if ((74 < 127) && (((("\n\t" + v8) + (v8 + "")) == v4) && (!v1))) {
                            v31 := GenS4{F0: v8, F1: 65535, F2: (!(!v1))}
                            _ = v31
                        } else {
                            delete(v3, (fi14 ^ fi12))
                            if (-86 < -128) {
                                var v32 int8 = 100
                                v4 = v4
                                v2 = append(v2, "")
                                _ = v32
                            } else {
                                var v33 int8 = 42
                                v13 -= 42
                                fmt.Println(fi14)
                                _ = v33
                            }
                        }
                    }
                    if v1 {
                        v5 -= v5
                        switch v4 {
                        case "":
                            var v34 uint16 = 32767
                            _ = v34
                        case "\n\t":
                            fmt.Println(fi14, v1)
                        default:
                            var v35 uint8 = 128
                            _ = v35
                        }
                    } else {
                        var v36 bool = ((100 != int8(fi14)) || v1)
                        _ = v36
                    }
                    _ = fi14
                }
                _ = v13
            }
            v5 = -640
            delete(v3, fi12)
            _ = fi12
        }
    }
    v37 := GenS1{F0: -32768, F1: -1, F2: 100}
    var v38 int = 128
    var v39 float64 = 3.14159
    v40 := GenS2{F0: v1, F1: 0}
    switch v38 {
    case -684:
        if (genFn0(v8, (v38 * (v38 % ((950) | 1)))) == 127) {
            v41 := GenS4{F0: "hello", F1: 0, F2: false}
            var v42 bool = (0 == -479)
            if (v42 || v1) {
                switch v38 {
                case -9223372036854775808:
                    v1 = (-527 >= -1)
                case 10:
                    v39 = v39
                default:
                    switch ("\n\t" + v4) {
                    case "a":
                        v43 := GenS1{F0: 429, F1: -1, F2: 83}
                        _ = v43
                    default:
                        var v44 uint8 = genFn0(v8, (v38 + 128))
                        _ = v44
                    }
                }
            }
            _ = v41
            _ = v42
        }
    case 910:
        var v45 uint = 47
        _ = v45
    case -897:
        fmt.Println(v39)
    }
    for fi46 := 0; fi46 < 15; fi46++ {
        v39 += float64(0.0)
        switch fi46 {
        case 528:
            switch (v38 << ((18446744073709551615) & 63)) {
            case -250:
                v47 := GenS5{F0: v1, F1: -128, F2: -9223372036854775808}
                _ = v47
            case 32768:
                var v48 float64 = (0.0 / v39)
                _ = v48
            case 2147483647:
                var v49 float32 = (-2.25 + 1e10)
                _ = v49
            default:
                switch v8 {
                case "foo bar":
                    if false {
                        v8 = "\n\t"
                        v50 := GenS4{F0: "a", F1: 0, F2: v1}
                        _ = v50
                    }
                }
            }
        }
        v51 := v3[-1]
        _ = fi46
        _ = v51
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v8
}
