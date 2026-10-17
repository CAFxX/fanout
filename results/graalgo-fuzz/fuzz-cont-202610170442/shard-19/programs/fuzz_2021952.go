package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 uint16 }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 uint8; F1 string; F2 int32 }
type GenS4 struct { F0 uint16; F1 string; F2 bool }
type GenS5 struct { F0 int16; F1 int8 }
type GenS6 struct { F0 int64; F1 string; F2 int32 }
func genFn0(p0 int16) uint64 {
    var v7 int64 = -1
    for fi8 := 0; fi8 < 8; fi8++ {
        var v9 float32 = ((((0.0 + ((((-2.25 - 1.5) * 1e10) + ((3.14159 / 3.14159) + 1.5)) - 0.0)) / ((3.14159 / 1.5) - 3.14159)) - float32(-2.25)) + 3.14159)
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: 124, F1: 96, F2: 42}
    fmt.Println(v7)
    _ = p0
    _ = v10
    _ = v7
    return 4294967295
}
func genFn1() int64 {
    switch "" {
    case "a":
        var v11 int8 = 1
        _ = v11
    default:
        var v12 int = 0
        _ = v12
    }
    var v13 uint64 = 18446744073709551615
    if (int(v13) > 1) {
        v14 := []int{32768, 32767, 42, 2147483647}
        if (true || true) {
            v13 = v13
            fmt.Println(v13)
        }
        switch (("a" + "hello") + "foo bar") {
        case "\n\t":
            v13 -= 18446744073709551615
        }
        _ = v14
    }
    _ = v13
    return -678
}

func main() {
    var v1 int = -2147483648
    v2 := make(map[int]int)
    var v3 bool = ((!true) && (3.14159 >= float32((1.5 - (-2.25 + -2.25)))))
    var v4 uint8 = 255
    v5 := make(map[uint32]string)
    var v6 float32 = ((1.5 / ((0.1 / float32(-2.25)) + 0.0)) - 0.1)
    v15 := []int{(v1 | (v1 >> ((18446744073709551615) & 63)))}
    var v16 int = (v1 / ((2147483647) | 1))
    v4 = v4
    var v17 int = v1
    for fi18 := 0; fi18 < 7; fi18++ {
        for fi19 := 0; fi19 < 8; fi19++ {
            v20 := GenS2{F0: 128, F1: 91}
            var v21 int64 = -9223372036854775808
            var v22 string = ""
            v3 = (0 != 2)
            if ((v21 + -9223372036854775808) < -1) {
                fmt.Println(v22)
                v15 = append(v15, -895)
            }
            if v3 {
                if (v22 > v22) {
                    if ((false && v3) && (!(!false))) {
                        switch v22 {
                        case "foo bar":
                            var v23 uint64 = 42
                            _ = v23
                        }
                    } else {
                        _ = len(v15)
                        var v24 int16 = 42
                        switch -274 {
                        case 0:
                            if true {
                                fmt.Println(v1, v21)
                                var v25 uint8 = v4
                                var v26 float64 = (0.1 - 0.1)
                                _ = v25
                                _ = v26
                            }
                        case -876:
                            var v27 int8 = -1
                            _ = v27
                        default:
                            if false {
                                switch v22 {
                                case "foo bar":
                                    v6 = (1.5 / (((1e10 * (float32((3.14159 - -2.25)) / 1.5)) - 0.0) - 0.1))
                                }
                                _ = len(v5)
                                v4 += (v4 | v4)
                                if v3 {
                                    fmt.Println(v16)
                                    var v28 uint64 = 2
                                    _ = v28
                                }
                                switch v22 {
                                case "":
                                    v29 := GenS2{F0: v4, F1: v24}
                                    _ = v29
                                default:
                                    v30 := GenS3{F0: v4, F1: "hello", F2: -106}
                                    _ = v30
                                }
                                v31 := GenS3{F0: (^(v4)), F1: (v22 + v22), F2: -1}
                                var v32 uint8 = v4
                                v33 := GenS4{F0: 2, F1: ("a" + v22), F2: ((0.1 / float64(((v6 - -2.25) - float32(1e10)))) < (1.5 + 0.0))}
                                if len(v15) >= 2 {
                                    v15 = v15[1:len(v15)]
                                }
                                _ = v31
                                _ = v32
                                _ = v33
                            } else {
                                var v34 int16 = v24
                                _ = v34
                            }
                        }
                        _ = v24
                    }
                    var v35 bool = true
                    _ = v35
                }
                var v36 bool = false
                switch v22 {
                case "\n\t":
                    var v37 float32 = v6
                    _ = v37
                }
                _ = v36
            }
            _ = fi19
            _ = v20
            _ = v21
            _ = v22
        }
        fmt.Println(v17, v6, v16)
        _ = fi18
    }
    switch v17 {
    case 4294967295:
        var v38 uint = 42
        _ = v38
    case -9223372036854775808:
        for fi39 := 0; fi39 < 15; fi39++ {
            if (0 <= 127) {
                var v40 int16 = -311
                for fi41 := 0; fi41 < 5; fi41++ {
                    v42 := GenS1{F0: v4, F1: (^(v4)), F2: 32767}
                    _ = fi41
                    _ = v42
                }
                _ = v40
            }
            v43 := GenS4{F0: uint16(v4), F1: "foo bar", F2: v3}
            _ = fi39
            _ = v43
        }
    }
    fmt.Println(v4, v6, v1)
    var v44 uint16 = 32767
    for fi45 := 0; fi45 < 11; fi45++ {
        fmt.Println(v17, fi45)
        v44 -= v44
        if v3 {
            v46 := GenS1{F0: 28, F1: (v4 >> ((2) & 63)), F2: v44}
            _ = v46
        } else {
            if (!true) {
                if (("" + "\n\t") >= "\n\t") {
                    if v3 {
                        v16 += 127
                        switch "" {
                        case "hello":
                            v6 = float32(1e10)
                        case "a":
                            v17 -= 304
                        case "":
                            for fi47 := 0; fi47 < 15; fi47++ {
                                v44 += (^(v44))
                                switch "foo bar" {
                                case "hello":
                                    v4 += v4
                                case "\n\t":
                                    var v48 uint8 = v4
                                    _ = v48
                                default:
                                    fmt.Println(v17)
                                    v49 := GenS2{F0: 134, F1: -592}
                                    _ = len(v15)
                                    if ((v3 && false) && (genFn1() >= 1)) {
                                        fmt.Println(fi47, v1, v16)
                                    }
                                    var v50 int8 = 1
                                    _ = v49
                                    _ = v50
                                }
                                var v51 uint = 255
                                delete(v5, 763)
                                v15 = append(v15, v1)
                                _ = fi47
                                _ = v51
                            }
                        default:
                            if false {
                                fmt.Println(v1)
                                if (!(!(790 < 424))) {
                                    var v52 int32 = -887
                                    fmt.Println(v52, v1, v44)
                                    if v3 {
                                        for fi53 := 0; fi53 < 16; fi53++ {
                                            delete(v2, 2147483647)
                                            var v54 uint32 = 944
                                            _ = fi53
                                            _ = v54
                                        }
                                        v6 = (v6 / ((v6 * 1e10) - ((0.0 / v6) + 0.1)))
                                    }
                                    _ = v52
                                }
                            }
                        }
                        var v55 bool = false
                        _ = v55
                    }
                    for fi56 := 0; fi56 < 7; fi56++ {
                        switch ("a" + "hello") {
                        case "":
                            v44 = v44
                        case "\n\t":
                            var v57 float64 = -2.25
                            _ = v57
                        default:
                            v4 = (v4 >> ((0) & 63))
                        }
                        if false {
                            v3 = (!v3)
                            v58 := GenS5{F0: 0, F1: 42}
                            _ = v58
                        } else {
                            fmt.Println(v6, fi56, v4)
                        }
                        var v59 int32 = 2147483647
                        _ = fi56
                        _ = v59
                    }
                    v16 = (v17 | v17)
                }
            }
            var v60 string = "foo bar"
            v60 = "a"
            _ = v60
        }
        _ = fi45
    }
    v17 = 32768
    if (!(-110 >= 1)) {
        for fi61 := 0; fi61 < 9; fi61++ {
            v1 += v17
            var v62 int16 = 42
            var v63 uint = 259
            _ = fi61
            _ = v62
            _ = v63
        }
    }
    for fi64 := 0; fi64 < 15; fi64++ {
        for fi65 := 0; fi65 < 2; fi65++ {
            v66 := GenS6{F0: genFn1(), F1: "a", F2: -2147483648}
            var v67 uint16 = (v44 >> ((519) & 63))
            _ = fi65
            _ = v66
            _ = v67
        }
        _ = fi64
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v5
    _ = v6
}
