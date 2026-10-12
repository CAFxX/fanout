package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 bool; F1 uint32; F2 uint16 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 string; F1 int64; F2 uint32 }
type GenS7 struct { F0 uint64; F1 int8 }

func main() {
    v1 := []uint32{}
    v2 := make(map[int]bool)
    var v3 bool = (-32768 < -1)
    var v4 int64 = -9223372036854775808
    var v5 bool = (2147483647 == 2147483647)
    v3 = v5
    v4 += -9223372036854775808
    for fi6 := 0; fi6 < 18; fi6++ {
        if v3 {
            v7 := GenS1{F0: 144}
            _ = v7
        } else {
            v3 = v3
            if v3 {
                v8 := v2[(fi6 + fi6)]
                switch "" {
                case "\n\t":
                    if v5 {
                        v1 = append(v1, 631)
                    } else {
                        if true {
                            var v9 int8 = 127
                            if (false || ((!v8) && ((v4 << ((1) & 63)) == (v4 + (v4 << ((18446744073709551615) & 63)))))) {
                                if v8 {
                                    var v10 string = "a"
                                    _ = v10
                                } else {
                                    var v11 float64 = float64(0.1)
                                    v5 = (((v9 + v9) != v9) && v8)
                                    _ = v11
                                }
                                v12 := GenS1{F0: 255}
                                for fi13 := 0; fi13 < 17; fi13++ {
                                    fmt.Println(v5, v5)
                                    v5 = (v8 && v3)
                                    var v14 string = "hello"
                                    _ = fi13
                                    _ = v14
                                }
                                _ = v12
                            }
                            v5 = true
                            _ = v9
                        } else {
                            var v15 uint64 = 160
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = v15
                        }
                        v4 -= (v4 << ((552) & 63))
                        var v16 float32 = -2.25
                        _ = v16
                    }
                case "a":
                    if (255 <= 42) {
                        if true {
                            if false {
                                var v17 int64 = (v4 & (^(v4)))
                                _ = v17
                            } else {
                                fmt.Println(v8, v5)
                            }
                            if v8 {
                                v18 := GenS2{F0: ""}
                                _ = v18
                            } else {
                                if (356 < 65535) {
                                    switch 2147483647 {
                                    case -25:
                                        var v19 bool = v3
                                        _ = v19
                                    case 127:
                                        switch ("foo bar" + ("\n\t" + "hello")) {
                                        case "foo bar":
                                            for fi20 := 0; fi20 < 8; fi20++ {
                                                v8 = v5
                                                if false {
                                                    v21 := GenS3{F0: (v4 - (v4 * (v4 * v4)))}
                                                    var v22 float32 = -2.25
                                                    _ = v21
                                                    _ = v22
                                                }
                                                var v23 uint16 = 65535
                                                _ = fi20
                                                _ = v23
                                            }
                                        case "a":
                                            delete(v2, 988)
                                        case "\n\t":
                                            v24 := GenS1{F0: 1}
                                            _ = v24
                                        default:
                                            for fi25 := 0; fi25 < 14; fi25++ {
                                                var v26 float64 = (float64((float32(float64(((((1.5 * (0.0 / 3.14159)) * 0.0) * (1.5 * 0.0)) - -2.25))) + float32(float64((1.5 + (0.1 + (0.0 + (1.5 + (float32((0.0 + 1e10)) * -2.25))))))))) - (3.14159 - 1.5))
                                                v2[(fi6 | (fi6 ^ (fi6 << ((363) & 63))))] = v3
                                                _ = fi25
                                                _ = v26
                                            }
                                        }
                                    default:
                                        _ = len(v2)
                                    }
                                }
                                v27 := v2[(fi6 + (fi6 << ((255) & 63)))]
                                v4 += (v4 - (v4 << ((823) & 63)))
                                _ = v27
                            }
                            for fi28 := 0; fi28 < 5; fi28++ {
                                if ((!((v4 + v4) >= (-(v4)))) && ((true && v5) && (!("a" <= ("" + "hello"))))) {
                                    var v29 int32 = -743
                                    var v30 int32 = (v29 >> ((42) & 63))
                                    _ = v29
                                    _ = v30
                                } else {
                                    if len(v1) > 0 {
                                        v1[0] = 2147483648
                                    }
                                    if (4294967295 == 1) {
                                        var v31 float64 = 1.5
                                        _ = v31
                                    }
                                }
                                v3 = v3
                                v32 := GenS3{F0: 9223372036854775807}
                                _ = fi28
                                _ = v32
                            }
                        }
                        for fi33 := 0; fi33 < 11; fi33++ {
                            _ = len(v2)
                            _ = fi33
                        }
                        var v34 int32 = 888
                        _ = v34
                    }
                case "foo bar":
                    if (!true) {
                        v35 := GenS4{F0: v3, F1: 773, F2: 32767}
                        var v36 uint32 = 65535
                        _ = v35
                        _ = v36
                    } else {
                        var v37 int16 = 0
                        var v38 uint64 = 0
                        _ = v37
                        _ = v38
                    }
                }
                _ = v8
            } else {
                switch fi6 {
                case 9223372036854775807:
                    v5 = v3
                case 1:
                    var v39 bool = v3
                    _ = v39
                default:
                    v40 := GenS5{F0: (v4 & (v4 * -9223372036854775808))}
                    _ = v40
                }
            }
            var v41 uint16 = 65535
            _ = v41
        }
        switch fi6 {
        case 4294967295:
            v42 := v2[-9223372036854775808]
            _ = v42
        case 32768:
            fmt.Println(fi6, v3)
        case 743:
            switch (fi6 ^ fi6) {
            case 9223372036854775807:
                var v43 bool = v5
                _ = v43
            case -230:
                for fi44 := 0; fi44 < 18; fi44++ {
                    v45 := GenS3{F0: v4}
                    _ = fi44
                    _ = v45
                }
            }
        default:
            var v46 uint32 = 179
            _ = v46
        }
        fmt.Println(fi6, fi6, v5)
        _ = fi6
    }
    v5 = (0 == 42)
    v3 = true
    delete(v2, -809)
    switch -328 {
    case 128:
        switch "hello" {
        case "\n\t":
            if (v3 || (v3 || ((!((0.1 + 0.0) >= 1.5)) && (false || ("" == ""))))) {
                fmt.Println(v3)
                var v47 uint = 221
                var v48 uint16 = 2
                _ = v47
                _ = v48
            }
        case "":
            var v49 uint8 = 128
            _ = v49
        case "a":
            v5 = (!(!(127 == 127)))
        default:
            v50 := GenS3{F0: (v4 ^ (v4 - (v4 ^ (v4 | (^(v4))))))}
            _ = v50
        }
    case 0:
        _ = len(v1)
    case -1:
        v51 := GenS6{F0: "hello", F1: 1, F2: 42}
        _ = v51
    default:
        v52 := GenS2{F0: "foo bar"}
        _ = v52
    }
    v5 = v5
    for fi53 := 0; fi53 < 12; fi53++ {
        for fi54 := 0; fi54 < 12; fi54++ {
            var v55 string = "hello"
            v3 = ((4294967295 == 2) && (!v5))
            _ = fi54
            _ = v55
        }
        _ = len(v2)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi53
    }
    var v56 string = "foo bar"
    v5 = (-1 > 127)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v57 uint32 = 1
    v3 = ((-(v4)) < 42)
    var v58 uint32 = (v57 / (((v57 / ((v57) | 1))) | 1))
    if ((((0.1 - float32((0.0 * float64(0.1)))) * ((float32((0.1 * ((3.14159 / 1.0) * -2.25))) + -2.25) / float32(1e10))) / (3.14159 / -2.25)) <= -2.25) {
        v59 := GenS7{F0: 850, F1: 1}
        for fi60 := 0; fi60 < 5; fi60++ {
            v5 = (!v5)
            fmt.Println(v58, fi60, fi60)
            var v61 float32 = float32(-2.25)
            _ = fi60
            _ = v61
        }
        _ = v59
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v57
    _ = v58
}
