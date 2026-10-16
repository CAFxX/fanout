package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 bool; F1 uint; F2 uint64 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int16; F1 int64; F2 int }

func main() {
    var v1 uint16 = 255
    v2 := make(map[uint32]bool)
    v3 := []uint32{2147483647, 2147483648, 2147483648}
    var v4 uint64 = 4294967295
    var v5 uint8 = 77
    var v6 uint = 524
    if false {
        switch 128 {
        case 1:
            v7 := GenS1{F0: 2147483647}
            _ = v7
        default:
            v8 := GenS1{F0: int32(v1)}
            _ = v8
        }
    } else {
        for fi9 := 0; fi9 < 7; fi9++ {
            for fi10 := 0; fi10 < 15; fi10++ {
                switch (("hello" + "\n\t") + ("" + "foo bar")) {
                case "":
                    switch (fi9 << ((1) & 63)) {
                    case -2147483648:
                        if true {
                            if false {
                                if (true || false) {
                                    var v11 uint8 = v5
                                    _ = v11
                                } else {
                                    delete(v2, uint32(v4))
                                }
                                v12 := GenS1{F0: 42}
                                _ = v12
                            } else {
                                v6 = v6
                            }
                            var v13 uint = v6
                            v1 = (v1 + v1)
                            _ = v13
                        } else {
                            if false {
                                if (((false || true) && (1.5 > -2.25)) || (42 >= 65535)) {
                                    var v14 uint64 = v4
                                    _ = v14
                                } else {
                                    v15 := GenS1{F0: -267}
                                    fmt.Println(v6)
                                    v16 := v2[1]
                                    _ = v15
                                    _ = v16
                                }
                                if ((9223372036854775807 == 0) || (-440 != 42)) {
                                    var v17 uint64 = (v4 + v4)
                                    fmt.Println(fi9, fi9, v17)
                                    v6 = v6
                                    _ = v17
                                } else {
                                    if (33 >= 1) {
                                        var v18 int = -740
                                        v1 = 0
                                        var v19 int16 = -1
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        if false {
                                            fmt.Println(v5, v4)
                                        }
                                        _ = v18
                                        _ = v19
                                    } else {
                                        var v20 uint16 = 65535
                                        fmt.Println(fi10, v5)
                                        fmt.Println(v1)
                                        _ = v20
                                    }
                                    v4 = uint64(v1)
                                }
                            }
                            v6 -= v6
                            if false {
                                fmt.Println(v5, v5, fi9)
                                var v21 int16 = -79
                                v4 = 18446744073709551615
                                _ = v21
                            }
                        }
                    default:
                        switch 9223372036854775807 {
                        case 32767:
                            v22 := GenS1{F0: -568}
                            _ = v22
                        case 344:
                            var v23 uint8 = uint8(v6)
                            _ = v23
                        case 278:
                            v24 := GenS2{F0: v4, F1: v5}
                            _ = v24
                        }
                    }
                case "\n\t":
                    var v25 bool = false
                    v26 := GenS2{F0: (v4 * v4), F1: v5}
                    if true {
                        switch fi10 {
                        case 127:
                            v27 := GenS2{F0: (^(v4)), F1: 42}
                            _ = v27
                        case 9223372036854775807:
                            v25 = false
                        case 0:
                            delete(v2, 1)
                        default:
                            delete(v2, 42)
                        }
                        v28 := GenS3{F0: true, F1: v6, F2: (v4 | (^(v4)))}
                        _ = v28
                    } else {
                        if (v25 && (0.0 <= 0.1)) {
                            if false {
                                if v25 {
                                    switch 32767 {
                                    case 127:
                                        fmt.Println(v25, fi9, v25)
                                    }
                                    var v29 int = fi9
                                    _ = v29
                                }
                                v4 -= (v4 & (v4 ^ 2))
                                if len(v3) > 0 {
                                    v3[0] = 684
                                }
                            }
                        } else {
                            fmt.Println(v6, v6)
                            var v30 uint8 = (v5 >> (((v6 | (v6 * 770))) & 63))
                            if ((false || v25) && true) {
                                var v31 bool = (!(!("foo bar" != ("" + ("\n\t" + "")))))
                                v31 = v31
                                fmt.Println(v5, v1, fi9)
                                var v32 uint = uint(v5)
                                v6 = uint(v4)
                                var v33 uint32 = 42
                                if v31 {
                                    if (v33 > v33) {
                                        fmt.Println(v4)
                                        if (531 > 611) {
                                            v31 = (-873 != (^(fi10)))
                                            v34 := GenS3{F0: v31, F1: (v6 | (v32 - v6)), F2: v4}
                                            fmt.Println(fi10, v6, v31)
                                            _ = v34
                                        }
                                        var v35 uint64 = v4
                                        _ = v35
                                    } else {
                                        v2[(v33 >> ((v32) & 63))] = v25
                                        v4 -= 2
                                        switch "foo bar" {
                                        case "":
                                            var v36 float32 = (-2.25 * 0.1)
                                            _ = v36
                                        case "foo bar":
                                            fmt.Println(fi9, v30)
                                        case "a":
                                            if v31 {
                                                v37 := GenS3{F0: v31, F1: 255, F2: uint64(v1)}
                                                if len(v3) > 0 {
                                                    v3[0] = (v33 & (v33 << (((v6 ^ v32)) & 63)))
                                                }
                                                var v38 int64 = 0
                                                _ = v37
                                                _ = v38
                                            } else {
                                                var v39 int16 = -1
                                                var v40 float32 = 3.14159
                                                _ = v39
                                                _ = v40
                                            }
                                        default:
                                            if true {
                                                var v41 uint16 = v1
                                                var v42 uint64 = v4
                                                fmt.Println(fi9, v31)
                                                _ = v41
                                                _ = v42
                                            }
                                        }
                                    }
                                    if false {
                                        v1 = v1
                                        switch (^(fi10)) {
                                        case 0:
                                            v6 = v32
                                        case 9223372036854775807:
                                            v1 -= 255
                                        default:
                                            var v43 uint32 = (^(v33))
                                            _ = v43
                                        }
                                    }
                                } else {
                                    v4 -= v4
                                }
                                _ = v31
                                _ = v32
                                _ = v33
                            } else {
                                _ = len(v3)
                                fmt.Println(v1, v4)
                            }
                            _ = v30
                        }
                    }
                    _ = v25
                    _ = v26
                }
                v44 := GenS2{F0: (v4 - v4), F1: v5}
                _ = fi10
                _ = v44
            }
            v45 := GenS2{F0: 901, F1: 1}
            var v46 uint32 = 42
            _ = fi9
            _ = v45
            _ = v46
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi47 := 0; fi47 < 5; fi47++ {
        for fi48 := 0; fi48 < 8; fi48++ {
            switch "\n\t" {
            case "foo bar":
                v49 := GenS4{F0: -1}
                _ = v49
            case "a":
                v50 := GenS5{F0: 67, F1: -9223372036854775808, F2: fi47}
                _ = v50
            default:
                v6 = v6
            }
            _ = fi48
        }
        for fi51 := 0; fi51 < 14; fi51++ {
            v52 := GenS2{F0: 1, F1: (^(v5))}
            var v53 int = (fi51 / (((-(fi51))) | 1))
            _ = fi51
            _ = v52
            _ = v53
        }
        _ = fi47
    }
    v5 -= v5
    for fi54 := 0; fi54 < 5; fi54++ {
        switch -2147483648 {
        case 1:
            if len(v3) > 0 {
                v3[0] = 42
            }
        case -2147483648:
            var v55 int32 = -1
            _ = v55
        case -821:
            for fi56 := 0; fi56 < 12; fi56++ {
                v5 = v5
                _ = fi56
            }
        default:
            switch (fi54 % ((fi54) | 1)) {
            case 614:
                for fi57 := 0; fi57 < 9; fi57++ {
                    fmt.Println(v5, v5)
                    fmt.Println(fi57, v6, v6)
                    v5 = (v5 >> (((v6 + uint(v4))) & 63))
                    _ = fi57
                }
            case -402:
                switch (fi54 / (((fi54 ^ (-(fi54)))) | 1)) {
                case 1:
                    for fi58 := 0; fi58 < 16; fi58++ {
                        fmt.Println(fi58)
                        v5 += v5
                        v4 = (v4 + (v4 * v4))
                        _ = fi58
                    }
                case 948:
                    if false {
                        switch ("\n\t" + "a") {
                        case "foo bar":
                            v59 := GenS3{F0: true, F1: 1, F2: (v4 ^ (^(v4)))}
                            _ = v59
                        }
                        for fi60 := 0; fi60 < 3; fi60++ {
                            var v61 uint16 = 2
                            v62 := GenS5{F0: -32768, F1: 0, F2: (fi60 * fi54)}
                            v63 := GenS5{F0: 1000, F1: -9223372036854775808, F2: 268}
                            _ = fi60
                            _ = v61
                            _ = v62
                            _ = v63
                        }
                    }
                default:
                    var v64 uint32 = 65535
                    _ = v64
                }
            }
        }
        v1 = v1
        _ = fi54
    }
    v4 += v4
    var v65 uint8 = (v5 >> ((uint(v1)) & 63))
    for fi66 := 0; fi66 < 13; fi66++ {
        var v67 string = "hello"
        _ = fi66
        _ = v67
    }
    switch "\n\t" {
    case "foo bar":
        switch 617 {
        case 2147483647:
            for fi68 := 0; fi68 < 18; fi68++ {
                if (!(!(-278 >= 42))) {
                    for fi69 := 0; fi69 < 12; fi69++ {
                        var v70 bool = false
                        _ = len(v2)
                        if len(v3) > 0 {
                            v3[0] = 277
                        }
                        v4 = (v4 | v4)
                        _ = fi69
                        _ = v70
                    }
                } else {
                    _ = len(v2)
                    fmt.Println(v65, v4, v6)
                }
                _ = fi68
            }
        case -1:
            var v71 int = -1
            _ = v71
        }
    case "":
        v72 := GenS5{F0: -1, F1: 0, F2: -847}
        _ = v72
    case "hello":
        var v73 int = 127
        _ = v73
    }
    var v74 int8 = -28
    if (v6 <= (v6 ^ v6)) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v74 = v74
    }
    for fi75 := 0; fi75 < 19; fi75++ {
        for fi76 := 0; fi76 < 7; fi76++ {
            v74 -= v74
            _ = fi76
        }
        v65 += (v65 ^ v5)
        fmt.Println(v74, v4, v6)
        _ = fi75
    }
    if (false && ((v5 / (((v5 * (v65 - (v65 - v5)))) | 1)) < v5)) {
        if false {
            var v77 int = -2147483648
            _ = v77
        } else {
            var v78 int64 = 42
            var v79 int16 = 66
            switch "" {
            case "hello":
                var v80 uint = (^(v6))
                _ = v80
            case "\n\t":
                var v81 uint = v6
                _ = v81
            }
            _ = v78
            _ = v79
        }
        delete(v2, 991)
    } else {
        v82 := v2[2]
        _ = v82
    }
    for fi83 := 0; fi83 < 9; fi83++ {
        v74 = v74
        _ = fi83
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v65:", v65)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v65
    _ = v74
}
