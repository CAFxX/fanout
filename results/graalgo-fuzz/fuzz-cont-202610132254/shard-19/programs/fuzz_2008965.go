package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64; F1 uint8 }
type GenS4 struct { F0 string; F1 int; F2 int16 }
type GenS5 struct { F0 int; F1 uint }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int64 }

func main() {
    var v1 string = "foo bar"
    var v2 uint = 4294967295
    v3 := []int{806, -9223372036854775808, 127}
    var v4 uint16 = 2
    v5 := make(map[uint32]uint32)
    for fi6 := 0; fi6 < 15; fi6++ {
        for fi7 := 0; fi7 < 4; fi7++ {
            var v8 int8 = 0
            if len(v3) > 0 {
                v3[0] = -41
            }
            _ = fi7
            _ = v8
        }
        var v9 uint64 = 42
        _ = fi6
        _ = v9
    }
    var v10 uint8 = 191
    var v11 string = ("\n\t" + "\n\t")
    var v12 int = -213
    for fi13 := 0; fi13 < 16; fi13++ {
        v1 = v11
        fmt.Println(v10)
        _ = fi13
    }
    v14 := v5[2]
    for fi15 := 0; fi15 < 10; fi15++ {
        if false {
            if (149 > (^(v10))) {
                v16 := GenS1{F0: 2147483647, F1: fi15, F2: -615}
                fmt.Println(v1, v14)
                fmt.Println(v10)
                _ = v16
            } else {
                if (uint8(v14) == (v10 * (^(v10)))) {
                    fmt.Println(v12, v14)
                    for fi17 := 0; fi17 < 7; fi17++ {
                        v18 := GenS1{F0: -1, F1: (-(fi15)), F2: -1}
                        v19 := GenS1{F0: 2147483647, F1: (v12 >> ((v2) & 63)), F2: 1000}
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                    v20 := GenS2{F0: 1}
                    _ = v20
                } else {
                    _ = len(v5)
                    if true {
                        var v21 uint32 = (^(v14))
                        if (!false) {
                            var v22 int16 = 590
                            _ = len(v3)
                            if ((!true) || (true || (42 != 1))) {
                                if (0 == 65535) {
                                    var v23 int32 = int32(v10)
                                    v24 := GenS3{F0: 446, F1: v10}
                                    _ = v23
                                    _ = v24
                                }
                                v4 += 800
                            }
                            _ = v22
                        }
                        _ = v21
                    } else {
                        if ((true || false) || true) {
                            v25 := GenS4{F0: "\n\t", F1: v12, F2: -272}
                            var v26 int = fi15
                            if false {
                                v27 := GenS5{F0: int(v4), F1: (v2 | (v2 - (v2 >> ((uint(v10)) & 63))))}
                                v4 -= uint16(v14)
                                var v28 int32 = int32(v14)
                                _ = v27
                                _ = v28
                            } else {
                                v1 = "hello"
                            }
                            _ = v25
                            _ = v26
                        } else {
                            var v29 uint = v2
                            var v30 string = v1
                            _ = v29
                            _ = v30
                        }
                    }
                    for fi31 := 0; fi31 < 3; fi31++ {
                        var v32 int64 = 1
                        v2 = 42
                        _ = fi31
                        _ = v32
                    }
                }
            }
            var v33 uint64 = 0
            _ = v33
        } else {
            if true {
                var v34 int16 = 775
                if (0.1 <= 0.0) {
                    if ((v34 >= (v34 % ((0) | 1))) && (!true)) {
                        if (false || ((v34 ^ 942) <= (v34 - 754))) {
                            var v35 uint8 = v10
                            _ = v35
                        } else {
                            switch ("a" + v1) {
                            case "":
                                var v36 uint8 = v10
                                _ = v36
                            case "a":
                                v5[4294967295] = v14
                            case "hello":
                                v37 := GenS6{F0: -2147483648}
                                _ = v37
                            }
                            var v38 bool = ((true && true) && true)
                            var v39 uint64 = 185
                            _ = v38
                            _ = v39
                        }
                    } else {
                        fmt.Println(v12, v2, fi15)
                    }
                    if (553 > 1) {
                        v40 := GenS4{F0: v11, F1: (v12 + (v12 / (((v12 % ((fi15) | 1))) | 1))), F2: v34}
                        _ = v40
                    } else {
                        for fi41 := 0; fi41 < 8; fi41++ {
                            fmt.Println(v14, v2)
                            var v42 int64 = 9223372036854775807
                            v43 := GenS7{F0: v42}
                            _ = fi41
                            _ = v42
                            _ = v43
                        }
                        fmt.Println(v11, v10)
                    }
                } else {
                    var v44 float32 = 1.5
                    if false {
                        v34 = (v34 % ((1) | 1))
                    } else {
                        v1 = v1
                        v4 += (v4 * v4)
                        v45 := v5[v14]
                        _ = v45
                    }
                    _ = v44
                }
                _ = v34
            }
        }
        var v46 uint32 = 1
        _ = fi15
        _ = v46
    }
    var v47 uint8 = (v10 % ((241) | 1))
    var v48 uint8 = v47
    if (!false) {
        if ((v10 != 42) || ((^(v12)) < v12)) {
            v1 = v1
            if len(v3) > 0 {
                v3[0] = v12
            }
            for fi49 := 0; fi49 < 20; fi49++ {
                if (894 != v2) {
                    v5[uint32(v2)] = (v14 >> (((v2 << ((v2) & 63))) & 63))
                    for fi50 := 0; fi50 < 10; fi50++ {
                        _ = len(v5)
                        _ = fi50
                    }
                    for fi51 := 0; fi51 < 10; fi51++ {
                        v52 := GenS5{F0: fi49, F1: (v2 - uint(v4))}
                        var v53 float32 = (1.5 / 1.0)
                        v48 += 0
                        _ = fi51
                        _ = v52
                        _ = v53
                    }
                } else {
                    for fi54 := 0; fi54 < 9; fi54++ {
                        if false {
                            switch "hello" {
                            case "a":
                                v1 = ("" + v1)
                            case "hello":
                                var v55 uint16 = v4
                                _ = v55
                            case "\n\t":
                                v5[v14] = v14
                            default:
                                var v56 int8 = 1
                                _ = v56
                            }
                            var v57 string = (v1 + "\n\t")
                            _ = v57
                        } else {
                            switch (v12 / ((32767) | 1)) {
                            case 32767:
                                var v58 int64 = 1
                                _ = v58
                            case -8:
                                v1 = v1
                            default:
                                if false {
                                    var v59 float64 = 1e10
                                    v60 := GenS5{F0: fi49, F1: (v2 % ((v2) | 1))}
                                    switch (v12 & fi54) {
                                    case -9223372036854775808:
                                        v10 += v47
                                    case 1:
                                        v59 -= 3.14159
                                    case 116:
                                        v61 := GenS1{F0: 0, F1: v12, F2: 1}
                                        _ = v61
                                    default:
                                        if ("\n\t" == "foo bar") {
                                            var v62 uint64 = 1
                                            switch v11 {
                                            case "hello":
                                                v47 -= v10
                                            case "":
                                                if false {
                                                    v5[289] = uint32(v2)
                                                    v5[v14] = (v14 & 2)
                                                    if false {
                                                        fmt.Println(v59)
                                                        v63 := GenS5{F0: 510, F1: (v2 ^ (^(v2)))}
                                                        v48 = (v10 + 42)
                                                        _ = v63
                                                    }
                                                    var v64 uint8 = 128
                                                    _ = v64
                                                } else {
                                                    if ((0.0 + 0.0) <= (0.1 - 3.14159)) {
                                                        var v65 uint32 = v14
                                                        v5[(^(v14))] = (^(v14))
                                                        _ = v65
                                                    }
                                                    switch fi54 {
                                                    case 32767:
                                                        if len(v3) > 0 {
                                                            v3[0] = (fi54 >> ((v2) & 63))
                                                        }
                                                    case 42:
                                                        var v66 int64 = 42
                                                        _ = v66
                                                    default:
                                                        fmt.Println(v10)
                                                    }
                                                }
                                            case "foo bar":
                                                var v67 uint64 = (v62 | (v62 >> (((v2 - 1)) & 63)))
                                                _ = v67
                                            default:
                                                fmt.Println(v11, v62)
                                            }
                                            _ = v62
                                        }
                                    }
                                    _ = v59
                                    _ = v60
                                } else {
                                    v68 := GenS4{F0: ("foo bar" + (("hello" + ("a" + v11)) + v11)), F1: (-(fi49)), F2: -329}
                                    switch (fi49 / ((fi49) | 1)) {
                                    case 828:
                                        var v69 float32 = 0.0
                                        _ = v69
                                    }
                                    var v70 uint8 = v48
                                    _ = v68
                                    _ = v70
                                }
                            }
                        }
                        v71 := v5[65535]
                        v71 = (v14 << ((v2) & 63))
                        _ = fi54
                        _ = v71
                    }
                    _ = len(v3)
                }
                v10 = (^(v48))
                v72 := GenS1{F0: 1, F1: (fi49 + 0), F2: -1}
                _ = fi49
                _ = v72
            }
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v4 = (v4 >> ((v2) & 63))
    } else {
        v12 = (v12 >> ((v2) & 63))
        if (true || false) {
            var v73 int64 = -1
            _ = v73
        }
        v3 = append(v3, (-(v12)))
    }
    v47 = (v48 | v47)
    _ = len(v3)
    if len(v3) > 0 {
        v3[0] = (v12 ^ v12)
    }
    delete(v5, 4294967295)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v5
}
