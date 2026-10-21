package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8; F1 int32; F2 int }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int8; F1 int16; F2 uint16 }
type GenS6 struct { F0 uint32; F1 int8 }
type GenS7 struct { F0 int32; F1 string; F2 uint8 }
type GenS8 struct { F0 int16 }
func genFn0(p0 uint16, p1 int64, p2 uint8) uint32 {
    var v6 int32 = -334
    p0 -= (p0 << ((255) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 505
}
func genFn1(p0 uint16, p1 int16, p2 string) int {
    if (false && true) {
        var v7 int64 = 1
        var v8 bool = (0.1 > (-2.25 / 1.0))
        v9 := GenS1{F0: 1, F1: (!v8)}
        _ = v7
        _ = v8
        _ = v9
    }
    v10 := []int{}
    p0 += 2
    var v11 int16 = 308
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return -1
}

func main() {
    v1 := []int{}
    var v2 string = ("hello" + "foo bar")
    var v3 int = 127
    v4 := make(map[string]int)
    var v5 float32 = (0.1 / 1e10)
    var v12 uint64 = 1
    var v13 bool = true
    if v13 {
        if (!(-9223372036854775808 <= 869)) {
            var v14 uint32 = 2
            _ = v14
        }
    } else {
        fmt.Println(v3, v13)
        var v15 uint32 = genFn0(0, -10, 2)
        _ = v15
    }
    v12 = 0
    if v13 {
        var v16 uint16 = 599
        v17 := GenS2{F0: 65535}
        for fi18 := 0; fi18 < 8; fi18++ {
            var v19 uint64 = v12
            v4["a"] = (fi18 & v3)
            _ = fi18
            _ = v19
        }
        _ = v16
        _ = v17
    }
    switch (v3 % (((^(v3))) | 1)) {
    case 128:
        var v20 string = (v2 + v2)
        _ = v20
    default:
        v13 = v13
    }
    switch v3 {
    case -9223372036854775808:
        for fi21 := 0; fi21 < 8; fi21++ {
            if v13 {
                v12 = (v12 * (v12 + (v12 >> ((2) & 63))))
                fmt.Println(fi21, v12)
            }
            v22 := GenS1{F0: -1, F1: (76 == 42)}
            var v23 int64 = 366
            _ = fi21
            _ = v22
            _ = v23
        }
    case -2147483648:
        fmt.Println(v5, v3)
    case 4294967295:
        v4[v2] = (v3 + (v3 * -105))
    }
    for fi24 := 0; fi24 < 18; fi24++ {
        for fi25 := 0; fi25 < 17; fi25++ {
            if (42 > 9223372036854775807) {
                switch v2 {
                case "\n\t":
                    v5 -= 1.5
                case "":
                    if v13 {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        switch v2 {
                        case "foo bar":
                            var v26 int8 = -127
                            _ = v26
                        case "a":
                            _ = len(v1)
                        case "":
                            v5 -= float32(1.5)
                        }
                        v2 = (v2 + "foo bar")
                        if true {
                            if v13 {
                                v1 = append(v1, genFn1(211, 42, ((v2 + (v2 + v2)) + ((v2 + v2) + v2))))
                                switch v2 {
                                case "foo bar":
                                    var v27 uint8 = 255
                                    _ = v27
                                case "\n\t":
                                    if (1 >= 2) {
                                        v2 = ""
                                    }
                                case "":
                                    if false {
                                        v13 = v13
                                    } else {
                                        var v28 int8 = 9
                                        switch "\n\t" {
                                        case "\n\t":
                                            var v29 int64 = 42
                                            _ = v29
                                        case "foo bar":
                                            fmt.Println(fi24, v12, v12)
                                        case "":
                                            switch (v2 + (("a" + v2) + v2)) {
                                            case "":
                                                v1 = append(v1, (^(v3)))
                                            default:
                                                v30 := GenS2{F0: 0}
                                                _ = v30
                                            }
                                            if (!v13) {
                                                v31 := GenS1{F0: 649, F1: false}
                                                if len(v1) > 0 {
                                                    v1[0] = (fi24 % (((v3 - fi24)) | 1))
                                                }
                                                _ = v31
                                            }
                                            fmt.Println(v13, v28)
                                            var v32 uint32 = 2
                                            var v33 float64 = 1e10
                                            _ = v32
                                            _ = v33
                                        default:
                                            if (4294967295 <= 42) {
                                                fmt.Println(fi24, v2)
                                            }
                                        }
                                        _ = v28
                                    }
                                }
                            } else {
                                if v13 {
                                    if len(v1) > 0 {
                                        v1[0] = (v3 % ((fi24) | 1))
                                    }
                                    _ = len(v4)
                                }
                                var v34 int16 = 32767
                                v35 := GenS3{F0: uint8(v12), F1: 720, F2: 42}
                                if v13 {
                                    var v36 float32 = 1e10
                                    _ = v36
                                } else {
                                    var v37 uint8 = 128
                                    switch "hello" {
                                    case "foo bar":
                                        v12 += 42
                                    }
                                    _ = v37
                                }
                                v1 = append(v1, (fi24 % ((2147483647) | 1)))
                                _ = len(v1)
                                switch v3 {
                                case 1:
                                    var v38 int16 = 519
                                    _ = v38
                                case 32767:
                                    v5 = ((v5 * float32(0.1)) * ((((v5 + v5) + float32(-2.25)) - v5) + v5))
                                }
                                var v39 bool = v13
                                v40 := GenS4{F0: 82}
                                var v41 bool = (0 >= 32767)
                                v2 = "a"
                                _ = v34
                                _ = v35
                                _ = v39
                                _ = v40
                                _ = v41
                            }
                        } else {
                            fmt.Println(v5, fi24)
                            var v42 bool = (true && (!v13))
                            if v13 {
                                var v43 uint = 2
                                _ = v43
                            } else {
                                fmt.Println(v12, v12, v2)
                                v44 := GenS5{F0: 100, F1: 1, F2: 533}
                                if ((1e10 - ((float32((0.0 * 1e10)) - (1.5 * (3.14159 + 3.14159))) - 1.5)) < ((1e10 / 1.5) * 1e10)) {
                                    var v45 int16 = 1000
                                    _ = v45
                                } else {
                                    v2 = v2
                                }
                                _ = v44
                            }
                            v46 := GenS1{F0: -100, F1: v42}
                            _ = v42
                            _ = v46
                        }
                        var v47 int8 = 42
                        var v48 int32 = int32(fi25)
                        v1 = append(v1, (v3 | (v3 ^ (^(fi24)))))
                        v47 += v47
                        v13 = (412 > fi25)
                        var v49 bool = false
                        var v50 int32 = v48
                        var v51 int32 = v50
                        _ = v47
                        _ = v48
                        _ = v49
                        _ = v50
                        _ = v51
                    }
                }
                if (9223372036854775807 == genFn1(42, 42, "\n\t")) {
                    v4[(v2 + v2)] = (v3 / ((v3) | 1))
                }
                v13 = (1 >= 2)
            } else {
                v2 = ((v2 + (v2 + v2)) + v2)
                v1 = append(v1, 173)
                fmt.Println(v3)
            }
            if v13 {
                var v52 uint64 = (^(v12))
                switch v2 {
                case "":
                    v53 := GenS6{F0: 65535, F1: 105}
                    _ = v53
                default:
                    v12 = (v52 / ((v12) | 1))
                }
                _ = v52
            } else {
                v54 := GenS2{F0: 815}
                _ = v54
            }
            var v55 int32 = 42
            if len(v1) > 0 {
                v1[0] = fi25
            }
            _ = fi25
            _ = v55
        }
        for fi56 := 0; fi56 < 7; fi56++ {
            if false {
                v57 := GenS5{F0: 127, F1: -1, F2: 197}
                if len(v1) > 0 {
                    v1[0] = fi24
                }
                var v58 uint = uint(v12)
                _ = v57
                _ = v58
            } else {
                v59 := GenS6{F0: 65535, F1: 127}
                _ = len(v1)
                _ = v59
            }
            if (v13 || v13) {
                v60 := GenS7{F0: 0, F1: v2, F2: 0}
                v61 := GenS1{F0: 0, F1: v13}
                if (190 >= 128) {
                    var v62 int16 = -1
                    switch v2 {
                    case "a":
                        var v63 int64 = 798
                        _ = v63
                    default:
                        if ((float64(v5) + 1.5) >= 3.14159) {
                            switch genFn1(793, 1000, ("" + (v2 + "a"))) {
                            case -1:
                                v64 := GenS8{F0: v62}
                                _ = v64
                            case -621:
                                v2 = ""
                            case 4294967295:
                                v65 := GenS3{F0: 42, F1: 2147483647, F2: (-(fi24))}
                                _ = v65
                            default:
                                if (0.0 < -2.25) {
                                    var v66 uint16 = 32767
                                    v2 = ("a" + ("" + v2))
                                    _ = v66
                                }
                                v3 -= v3
                                fmt.Println(v62, v2, v62)
                            }
                            var v67 int8 = 1
                            var v68 uint = 0
                            _ = v67
                            _ = v68
                        }
                    }
                    _ = len(v4)
                    _ = v62
                }
                _ = v60
                _ = v61
            }
            _ = fi56
        }
        if len(v1) > 0 {
            v1[0] = v3
        }
        _ = fi24
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
