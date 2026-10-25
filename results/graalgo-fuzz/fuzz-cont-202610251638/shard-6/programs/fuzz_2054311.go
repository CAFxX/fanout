package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 int }
type GenS2 struct { F0 int8; F1 bool; F2 uint64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint; F1 uint32 }
type GenS7 struct { F0 int8; F1 int64 }
func genFn0() uint {
    for fi4 := 0; fi4 < 3; fi4++ {
        var v5 bool = (0 != 127)
        var v6 bool = false
        v7 := []int{fi4, 32767, fi4}
        _ = fi4
        _ = v5
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 18; fi8++ {
        var v9 uint8 = 206
        v10 := GenS1{F0: 2, F1: 42, F2: (fi8 << ((42) & 63))}
        if (!(1 <= 65535)) {
            v9 -= v9
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    return 255
}
func genFn1(p0 bool) int64 {
    p0 = (255 > 33)
    var v11 int8 = 100
    v11 = (-(v11))
    var v12 uint64 = 1
    _ = p0
    _ = v11
    _ = v12
    return 42
}

func main() {
    v1 := []int{32767, 128, 364, -438}
    v2 := make(map[int]uint32)
    v3 := []string{"a"}
    if false {
        v3 = append(v3, "foo bar")
        var v13 int32 = 1
        delete(v2, 841)
        _ = v13
    }
    var v14 int16 = 32767
    v14 = v14
    v15 := GenS2{F0: -37, F1: true, F2: 1}
    v14 -= v14
    v14 = (v14 ^ (v14 ^ (^(v14))))
    v16 := GenS2{F0: 127, F1: false, F2: 0}
    var v17 int16 = 42
    if false {
        v18 := GenS2{F0: 42, F1: false, F2: 0}
        _ = v18
    } else {
        switch 42 {
        case -9223372036854775808:
            fmt.Println(v17, v14)
        case 32767:
            v19 := GenS1{F0: 1, F1: 986, F2: 32768}
            _ = v19
        case 793:
            v17 += -32768
        default:
            v20 := GenS3{F0: 810}
            _ = v20
        }
    }
    if (4294967295 < 2) {
        v14 = (v14 + 969)
        if ((65535 > 189) && false) {
            if (true && true) {
                for fi21 := 0; fi21 < 10; fi21++ {
                    if false {
                        if len(v1) > 0 {
                            v1[0] = fi21
                        }
                        var v22 int32 = 42
                        for fi23 := 0; fi23 < 4; fi23++ {
                            v14 -= 1
                            _ = fi23
                        }
                        _ = v22
                    }
                    for fi24 := 0; fi24 < 9; fi24++ {
                        fmt.Println(v17, fi24, fi21)
                        var v25 bool = (1 > 78)
                        v17 = (^(v17))
                        var v26 uint64 = 69
                        if false {
                            var v27 int = -2147483648
                            switch "a" {
                            case "\n\t":
                                switch "" {
                                case "foo bar":
                                    v14 -= v14
                                    v28 := GenS4{F0: (4294967295 != 2147483647)}
                                    var v29 uint = genFn0()
                                    var v30 float64 = 0.0
                                    _ = v28
                                    _ = v29
                                    _ = v30
                                case "hello":
                                    v1 = append(v1, (fi24 << ((1) & 63)))
                                case "a":
                                    if len(v1) > 0 {
                                        v1[0] = (fi21 % ((-2147483648) | 1))
                                    }
                                    v31 := v2[v27]
                                    _ = v31
                                default:
                                    v14 = v17
                                }
                            default:
                                var v32 uint = genFn0()
                                _ = v32
                            }
                            _ = v27
                        }
                        _ = fi24
                        _ = v25
                        _ = v26
                    }
                    _ = fi21
                }
            }
            var v33 string = ""
            var v34 uint = 65535
            _ = v33
            _ = v34
        }
    } else {
        for fi35 := 0; fi35 < 5; fi35++ {
            switch fi35 {
            case -9223372036854775808:
                var v36 bool = false
                _ = v36
            case 924:
                switch "\n\t" {
                case "hello":
                    var v37 uint16 = 65535
                    _ = v37
                case "":
                    v2[840] = 710
                case "\n\t":
                    v38 := v2[(fi35 + fi35)]
                    _ = v38
                default:
                    if (-2.25 > 0.0) {
                        var v39 int64 = -1
                        _ = v39
                    } else {
                        v2[751] = 896
                    }
                }
            default:
                v40 := GenS5{F0: v14}
                _ = v40
            }
            v41 := GenS6{F0: 505, F1: 2}
            _ = fi35
            _ = v41
        }
    }
    v42 := v2[-2147483648]
    for fi43 := 0; fi43 < 13; fi43++ {
        v44 := GenS6{F0: 803, F1: v42}
        if (1 <= 54) {
            for fi45 := 0; fi45 < 11; fi45++ {
                switch "" {
                case "foo bar":
                    v46 := GenS6{F0: 4294967295, F1: (v42 & 652)}
                    _ = v46
                case "":
                    v47 := v2[fi43]
                    _ = v47
                default:
                    _ = len(v2)
                }
                var v48 float32 = 0.1
                _ = fi45
                _ = v48
            }
            v3 = append(v3, "")
            v17 = -32768
        } else {
            v49 := GenS5{F0: -32768}
            _ = v49
        }
        switch (fi43 ^ fi43) {
        case 32768:
            var v50 int16 = (v14 & (v14 ^ v14))
            _ = v50
        case 9223372036854775807:
            v51 := GenS4{F0: (false && false)}
            _ = v51
        case 42:
            switch "a" {
            case "a":
                _ = len(v3)
            case "foo bar":
                if (false && true) {
                    switch "\n\t" {
                    case "foo bar":
                        fmt.Println(v17, fi43, fi43)
                    case "hello":
                        var v52 int8 = 127
                        _ = v52
                    }
                } else {
                    var v53 int64 = 469
                    switch "foo bar" {
                    case "":
                        v42 = v42
                    case "foo bar":
                        v14 += 1
                    default:
                        fmt.Println(fi43, v17)
                    }
                    v54 := GenS2{F0: 127, F1: true, F2: 933}
                    _ = v53
                    _ = v54
                }
            case "hello":
                v55 := GenS7{F0: 100, F1: genFn1(((v14 * v14) != 1000))}
                _ = v55
            default:
                switch (fi43 & 32767) {
                case 32767:
                    switch fi43 {
                    case -2147483648:
                        for fi56 := 0; fi56 < 11; fi56++ {
                            if true {
                                switch (((((("hello" + ("\n\t" + (("" + "") + "foo bar"))) + ("foo bar" + "a")) + ((("\n\t" + "hello") + "") + "a")) + "hello") + "a") + "a") {
                                case "hello":
                                    v42 = (v42 << ((1) & 63))
                                    fmt.Println(fi43, fi56)
                                    v57 := GenS6{F0: 118, F1: (v42 / (((v42 << ((42) & 63))) | 1))}
                                    switch "foo bar" {
                                    case "\n\t":
                                        var v58 string = ("a" + ("hello" + ""))
                                        _ = v58
                                    case "hello":
                                        v42 += (v42 * (^(v42)))
                                    }
                                    _ = v57
                                default:
                                    v2[-9223372036854775808] = 704
                                }
                                switch (-(fi43)) {
                                case 863:
                                    v2[fi43] = v42
                                case 42:
                                    var v59 bool = ((false && false) || false)
                                    _ = v59
                                case 127:
                                    v60 := v2[386]
                                    _ = v60
                                }
                            }
                            fmt.Println(v14, v17, v17)
                            v42 = v42
                            _ = fi56
                        }
                    }
                case 128:
                    v42 += (v42 >> ((65535) & 63))
                }
            }
        default:
            v17 = 703
        }
        _ = fi43
        _ = v44
    }
    v14 += -1
    if false {
        if (1.5 >= (1.5 - 3.14159)) {
            v14 -= 42
            for fi61 := 0; fi61 < 1; fi61++ {
                if false {
                    v62 := GenS5{F0: -646}
                    fmt.Println(fi61, v14, v14)
                    _ = v62
                } else {
                    v2[(fi61 >> ((700) & 63))] = v42
                    v42 = 481
                }
                var v63 int16 = v17
                v64 := GenS2{F0: 100, F1: true, F2: 42}
                _ = fi61
                _ = v63
                _ = v64
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v42
}
