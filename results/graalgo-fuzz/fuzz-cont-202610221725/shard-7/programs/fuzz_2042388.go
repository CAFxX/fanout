package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 int16 }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 int32; F1 int; F2 string }
type GenS4 struct { F0 uint; F1 uint16; F2 string }
type GenS5 struct { F0 string; F1 int32 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint8; F1 int16; F2 string }
type GenS8 struct { F0 int32 }
func genFn0(p0 int16) int64 {
    p0 = 1000
    fmt.Println(p0)
    p0 = p0
    _ = p0
    return 66
}
func genFn1(p0 int8, p1 int32) uint64 {
    var v7 uint16 = 32767
    if (!(420 > 529)) {
        for fi8 := 0; fi8 < 3; fi8++ {
            var v9 uint8 = 184
            var v10 bool = (true && false)
            v11 := GenS1{F0: p1, F1: p0, F2: 1000}
            _ = fi8
            _ = v10
            _ = v11
            _ = v9
        }
        for fi12 := 0; fi12 < 19; fi12++ {
            var v13 uint = 600
            for fi14 := 0; fi14 < 11; fi14++ {
                var v15 int = fi12
                switch 0 {
                case -999:
                    var v16 float64 = float64((3.14159 * (0.1 + (1e10 + 3.14159))))
                    _ = v16
                case -1:
                    v15 = (fi14 | fi14)
                case 670:
                    fmt.Println(v15)
                    v17 := GenS2{F0: v7, F1: 249}
                    _ = v17
                default:
                    fmt.Println(fi14, v15, v15)
                }
                var v18 int16 = 1000
                _ = fi14
                _ = v15
                _ = v18
            }
            _ = fi12
            _ = v13
        }
        v19 := GenS3{F0: 42, F1: 32767, F2: ""}
        _ = v19
    } else {
        var v20 float32 = 1.5
        for fi21 := 0; fi21 < 2; fi21++ {
            p0 = p0
            if (2147483647 > fi21) {
                var v22 int64 = -323
                fmt.Println(fi21, p0, v20)
                _ = v22
            }
            v23 := GenS4{F0: 873, F1: (v7 & 522), F2: ("\n\t" + "foo bar")}
            _ = fi21
            _ = v23
        }
        if true {
            var v24 string = "foo bar"
            _ = v24
        }
        _ = v20
    }
    for fi25 := 0; fi25 < 6; fi25++ {
        var v26 string = ""
        fmt.Println(p0)
        _ = fi25
        _ = v26
    }
    _ = p0
    _ = p1
    _ = v7
    return 18446744073709551615
}

func main() {
    var v1 int8 = -128
    var v2 string = "a"
    v3 := make(map[uint32]uint32)
    v4 := make(map[uint32]uint32)
    v5 := []uint32{42}
    var v6 int16 = -635
    switch (v2 + v2) {
    case "hello":
        if true {
            v4[65535] = 65535
        } else {
            v6 -= (v6 & 1)
        }
    default:
        v27 := GenS2{F0: 65535, F1: 65535}
        _ = v27
    }
    var v28 int32 = -2147483648
    fmt.Println(v2, v28)
    var v29 uint8 = 1
    if false {
        v6 = (^(v6))
    }
    var v30 int = 762
    switch 0 {
    case 408:
        var v31 int64 = genFn0(v6)
        _ = v31
    case 334:
        v32 := GenS5{F0: v2, F1: 42}
        _ = v32
    }
    var v33 int8 = v1
    switch v2 {
    case "":
        v6 -= v6
    default:
        if true {
            if (float64(3.14159) < ((1e10 * 1.5) - -2.25)) {
                v34 := GenS6{F0: (v28 / (((v28 | -2147483648)) | 1))}
                fmt.Println(v30)
                _ = v34
            } else {
                var v35 int = (v30 * v30)
                delete(v3, 4294967295)
                v36 := GenS3{F0: -2147483648, F1: int(v28), F2: ("foo bar" + (v2 + v2))}
                _ = v35
                _ = v36
            }
            if true {
                v37 := GenS4{F0: 0, F1: 42, F2: v2}
                fmt.Println(v1)
                _ = v37
            }
        } else {
            v28 += v28
        }
    }
    for fi38 := 0; fi38 < 5; fi38++ {
        if false {
            for fi39 := 0; fi39 < 14; fi39++ {
                var v40 uint32 = 1
                if (!(350 > 662)) {
                    if true {
                        if false {
                            var v41 uint16 = 0
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            var v42 uint8 = (^(v29))
                            _ = v41
                            _ = v42
                        }
                        switch fi39 {
                        case -1:
                            if false {
                                var v43 uint = 530
                                if true {
                                    _ = len(v5)
                                    v30 -= 276
                                    delete(v4, (v40 & (v40 ^ 2147483647)))
                                }
                                _ = v43
                            } else {
                                switch ((("" + v2) + v2) + v2) {
                                case "hello":
                                    switch fi38 {
                                    case 32768:
                                        v44 := GenS4{F0: 255, F1: 2, F2: v2}
                                        _ = v44
                                    }
                                case "a":
                                    v45 := GenS7{F0: (v29 & (v29 + (^(v29)))), F1: v6, F2: "\n\t"}
                                    _ = v45
                                case "":
                                    v33 = v1
                                default:
                                    var v46 uint32 = 2147483648
                                    _ = v46
                                }
                                v6 = v6
                            }
                            var v47 float32 = -2.25
                            v48 := GenS7{F0: (^(v29)), F1: v6, F2: (v2 + v2)}
                            _ = v47
                            _ = v48
                        }
                    }
                }
                var v49 float32 = (0.1 + (-2.25 * 0.1))
                _ = fi39
                _ = v40
                _ = v49
            }
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v1 = v1
        _ = fi38
    }
    switch v2 {
    case "\n\t":
        v29 = (v29 % ((v29) | 1))
    default:
        v50 := v3[2147483647]
        _ = v50
    }
    v6 += -32768
    switch ((v2 + "\n\t") + v2) {
    case "":
        var v51 uint16 = 32767
        _ = v51
    case "foo bar":
        var v52 uint = 1
        _ = v52
    case "hello":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    v28 += v28
    for fi53 := 0; fi53 < 20; fi53++ {
        _ = len(v4)
        _ = fi53
    }
    for fi54 := 0; fi54 < 16; fi54++ {
        var v55 int32 = v28
        _ = fi54
        _ = v55
    }
    if false {
        v28 = (v28 * v28)
        if true {
            for fi56 := 0; fi56 < 14; fi56++ {
                for fi57 := 0; fi57 < 4; fi57++ {
                    delete(v4, 65535)
                    var v58 uint = 0
                    switch 42 {
                    case -621:
                        v6 += (v6 | v6)
                    case -519:
                        v59 := GenS6{F0: 1}
                        _ = v59
                    case 979:
                        switch v2 {
                        case "\n\t":
                            var v60 int = fi56
                            _ = v60
                        case "foo bar":
                            v61 := GenS4{F0: v58, F1: 32767, F2: "\n\t"}
                            _ = v61
                        }
                    default:
                        switch ("foo bar" + (v2 + ("hello" + ((((v2 + (v2 + (v2 + ("\n\t" + "")))) + (v2 + v2)) + "foo bar") + v2)))) {
                        case "a":
                            if false {
                                var v62 int16 = (v6 - v6)
                                _ = v62
                            } else {
                                v63 := GenS4{F0: (v58 | v58), F1: 0, F2: ((v2 + v2) + v2)}
                                v33 -= (v1 % ((v1) | 1))
                                var v64 int = fi57
                                v28 = (v28 >> ((v58) & 63))
                                _ = v63
                                _ = v64
                            }
                        case "foo bar":
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        default:
                            v65 := GenS1{F0: (^(v28)), F1: (v1 >> (((v58 % ((v58) | 1))) & 63)), F2: 387}
                            delete(v4, 2147483648)
                            _ = v65
                        }
                    }
                    _ = fi57
                    _ = v58
                }
                _ = fi56
            }
            v66 := GenS8{F0: int32(v29)}
            _ = v66
        }
    } else {
        var v67 int32 = v28
        v1 = (^(v33))
        _ = v67
    }
    if (v6 == -798) {
        v33 = v33
        var v68 uint16 = 255
        fmt.Println(v2)
        _ = v68
    } else {
        v2 = (v2 + "")
        v69 := GenS8{F0: (v28 << ((255) & 63))}
        switch v30 {
        case 2147483647:
            v70 := GenS5{F0: v2, F1: (v28 % ((v28) | 1))}
            _ = v70
        case 1:
            fmt.Println(v1, v29)
        case 32767:
            v29 -= v29
        default:
            for fi71 := 0; fi71 < 2; fi71++ {
                v2 = v2
                v29 = (v29 >> ((255) & 63))
                _ = fi71
            }
        }
        _ = v69
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
