package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint32; F1 uint16; F2 int64 }
type GenS5 struct { F0 bool; F1 uint; F2 uint32 }
type GenS6 struct { F0 int8; F1 int64 }
func genFn0(p0 int16, p1 bool, p2 uint8) int {
    var v6 int32 = 2147483647
    if false {
        for fi7 := 0; fi7 < 6; fi7++ {
            var v8 int = 520
            v9 := GenS1{F0: true, F1: 1}
            v10 := GenS1{F0: p1, F1: (v8 >> ((2) & 63))}
            _ = fi7
            _ = v10
            _ = v8
            _ = v9
        }
    } else {
        v6 = (v6 + v6)
    }
    v6 -= v6
    var v11 int8 = -39
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v6
    return -2147483648
}
func genFn1(p0 int8) uint32 {
    v12 := []int{424, 239}
    if (-9223372036854775808 < 9223372036854775807) {
        p0 += (p0 | p0)
    } else {
        for fi13 := 0; fi13 < 4; fi13++ {
            if (!false) {
                v14 := GenS1{F0: (-2.25 > 1.5), F1: (fi13 + fi13)}
                _ = v14
            }
            var v15 int32 = 2147483647
            var v16 int64 = -9223372036854775808
            _ = fi13
            _ = v15
            _ = v16
        }
        var v17 uint32 = 23
        _ = v17
    }
    if false {
        if true {
            v18 := GenS1{F0: (0 == 4294967295), F1: -829}
            var v19 bool = true
            _ = v18
            _ = v19
        } else {
            v20 := GenS1{F0: true, F1: 32768}
            if (1 > 613) {
                v21 := GenS2{F0: -1}
                p0 += 127
                _ = v21
            }
            v22 := GenS2{F0: (p0 & p0)}
            _ = v20
            _ = v22
        }
        if len(v12) > 0 {
            v12[0] = 4294967295
        }
        v23 := GenS2{F0: p0}
        _ = v23
    } else {
        if (int(p0) != -2147483648) {
            var v24 int16 = 1000
            if true {
                var v25 float64 = (((0.0 / ((1.5 + 1e10) + 1.5)) + (1e10 * (1.5 - 3.14159))) - 3.14159)
                v24 += v24
                var v26 int32 = -1
                _ = v25
                _ = v26
            } else {
                v12 = append(v12, genFn0(v24, true, 127))
                p0 = p0
            }
            for fi27 := 0; fi27 < 1; fi27++ {
                var v28 uint8 = 128
                _ = fi27
                _ = v28
            }
            _ = v24
        }
        if len(v12) > 0 {
            v12[0] = genFn0(0, (p0 == p0), 94)
        }
    }
    _ = p0
    _ = v12
    return 0
}

func main() {
    var v1 float64 = 1e10
    var v2 int = 245
    var v3 uint32 = 42
    var v4 uint = 42
    var v5 uint8 = uint8(v4)
    var v29 uint64 = 42
    if ((("" + ("a" + (("" + (("foo bar" + "foo bar") + "\n\t")) + ""))) + "") >= "foo bar") {
        v1 = ((3.14159 * float64((((float32(v1) / 1.5) / (1.5 / (-2.25 * (3.14159 / 3.14159)))) / 3.14159))) * v1)
        for fi30 := 0; fi30 < 5; fi30++ {
            switch ("" + "a") {
            case "foo bar":
                var v31 uint32 = (^(v3))
                _ = v31
            case "":
                switch fi30 {
                case 281:
                    var v32 int8 = 42
                    _ = v32
                default:
                    switch (v2 / ((fi30) | 1)) {
                    case -1:
                        for fi33 := 0; fi33 < 11; fi33++ {
                            v34 := GenS1{F0: (true && true), F1: v2}
                            v5 += (v5 | 128)
                            _ = fi33
                            _ = v34
                        }
                    case 32767:
                        v29 = 965
                    }
                }
            case "hello":
                fmt.Println(v3, fi30)
            default:
                v1 += (v1 / v1)
            }
            var v35 int16 = 226
            _ = fi30
            _ = v35
        }
        for fi36 := 0; fi36 < 6; fi36++ {
            for fi37 := 0; fi37 < 5; fi37++ {
                v1 = ((v1 - v1) / v1)
                v29 = uint64(v5)
                v38 := GenS1{F0: (-1 >= fi36), F1: fi36}
                _ = fi37
                _ = v38
            }
            if false {
                if (1 < (v29 / (((v29 / ((uint64(v5)) | 1))) | 1))) {
                    for fi39 := 0; fi39 < 9; fi39++ {
                        switch int(v4) {
                        case 2147483647:
                            var v40 uint64 = 1
                            _ = v40
                        case 287:
                            v3 += 126
                        }
                        switch (fi39 / ((-748) | 1)) {
                        case -672:
                            switch "hello" {
                            case "hello":
                                v29 = 18446744073709551615
                            case "foo bar":
                                v3 -= (v3 << ((uint(v5)) & 63))
                            case "\n\t":
                                var v41 int16 = 1000
                                _ = v41
                            default:
                                var v42 int32 = 2147483647
                                _ = v42
                            }
                        case 1:
                            v43 := []int{}
                            _ = v43
                        }
                        switch (v2 - genFn0(-1, false, v5)) {
                        case 42:
                            v1 += v1
                        default:
                            v44 := GenS2{F0: 126}
                            _ = v44
                        }
                        _ = fi39
                    }
                }
            }
            _ = fi36
        }
    } else {
        fmt.Println(v2, v4, v4)
    }
    v2 -= v2
    var v45 uint16 = 79
    v2 -= (v2 << ((120) & 63))
    var v46 uint8 = 42
    v47 := []int{(^(v2)), (v2 - genFn0(-1, true, (v5 | v5))), v2}
    v48 := GenS3{F0: 216}
    var v49 bool = (!(3.14159 != 0.1))
    for fi50 := 0; fi50 < 12; fi50++ {
        v51 := GenS4{F0: (v3 | (v3 + 42)), F1: (^(v45)), F2: 9223372036854775807}
        _ = fi50
        _ = v51
    }
    var v52 int = (^(v2))
    var v53 uint32 = (v3 % (((^(v3))) | 1))
    fmt.Println(v1, v1, v2)
    v54 := GenS2{F0: 100}
    var v55 uint16 = (v45 * v45)
    if v49 {
        if (v49 && (-9223372036854775808 > -709)) {
            switch (v2 / ((0) | 1)) {
            case 2147483647:
                switch ("foo bar" + "foo bar") {
                case "hello":
                    for fi56 := 0; fi56 < 10; fi56++ {
                        v57 := GenS4{F0: genFn1(-1), F1: v55, F2: 42}
                        var v58 int64 = 0
                        _ = fi56
                        _ = v57
                        _ = v58
                    }
                case "foo bar":
                    if ((3.14159 - -2.25) > 0.1) {
                        if false {
                            switch 770 {
                            case 308:
                                v47 = append(v47, v2)
                            case 317:
                                var v59 uint64 = uint64(v45)
                                _ = v59
                            default:
                                v45 = v45
                            }
                            switch (("foo bar" + ("" + "\n\t")) + ("\n\t" + ("" + "foo bar"))) {
                            case "a":
                                var v60 int8 = 21
                                _ = v60
                            }
                        } else {
                            v1 += (v1 / (1.5 * v1))
                            if false {
                                v2 = 127
                            }
                        }
                        var v61 uint8 = uint8(v3)
                        v62 := GenS3{F0: (v3 << ((uint(v29)) & 63))}
                        _ = v61
                        _ = v62
                    }
                default:
                    var v63 float32 = (float32(1.5) + 0.0)
                    _ = v63
                }
            case 32767:
                if len(v47) >= 2 {
                    v47 = v47[1:len(v47)]
                }
            case 940:
                if (v49 || v49) {
                    v64 := GenS5{F0: true, F1: v4, F2: v53}
                    _ = v64
                } else {
                    var v65 int = v2
                    if len(v47) >= 2 {
                        v47 = v47[1:len(v47)]
                    }
                    if v49 {
                        var v66 int32 = -1
                        var v67 int64 = -1
                        _ = v66
                        _ = v67
                    }
                    _ = v65
                }
            }
            switch (v2 >> (((v4 | (v4 - (v4 & v4)))) & 63)) {
            case 42:
                v53 = genFn1(-128)
            default:
                fmt.Println(v29)
            }
            switch v52 {
            case -1:
                switch ("foo bar" + ("a" + "hello")) {
                case "foo bar":
                    switch v2 {
                    case -2147483648:
                        var v68 uint8 = v46
                        _ = v68
                    case 127:
                        var v69 bool = ("\n\t" >= ("foo bar" + "\n\t"))
                        _ = v69
                    case -257:
                        v70 := GenS4{F0: v53, F1: v55, F2: 1}
                        _ = v70
                    default:
                        if v49 {
                            _ = len(v47)
                            var v71 string = ""
                            for fi72 := 0; fi72 < 15; fi72++ {
                                v52 -= (^(v52))
                                var v73 string = ((v71 + (v71 + "\n\t")) + "hello")
                                _ = fi72
                                _ = v73
                            }
                            _ = v71
                        } else {
                            v46 = (v46 >> ((4294967295) & 63))
                            switch ("foo bar" + "foo bar") {
                            case "":
                                var v74 int64 = 9223372036854775807
                                _ = v74
                            case "\n\t":
                                for fi75 := 0; fi75 < 13; fi75++ {
                                    for fi76 := 0; fi76 < 7; fi76++ {
                                        v47 = append(v47, v52)
                                        var v77 float64 = v1
                                        v47 = append(v47, 4294967295)
                                        _ = fi76
                                        _ = v77
                                    }
                                    var v78 uint8 = (^(v5))
                                    _ = fi75
                                    _ = v78
                                }
                            }
                        }
                    }
                case "":
                    var v79 int8 = 127
                    _ = v79
                case "a":
                    v52 += 32768
                }
            }
        } else {
            v80 := GenS6{F0: 42, F1: int64(v55)}
            _ = v80
        }
        v81 := GenS5{F0: (42 != -32768), F1: (v4 - v4), F2: v53}
        for fi82 := 0; fi82 < 10; fi82++ {
            v53 = (v53 % (((v53 / ((226) | 1))) | 1))
            var v83 float64 = v1
            _ = fi82
            _ = v83
        }
        _ = v81
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v47:", len(v47), v47)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v52
    _ = v53
    _ = v54
    _ = v55
}
