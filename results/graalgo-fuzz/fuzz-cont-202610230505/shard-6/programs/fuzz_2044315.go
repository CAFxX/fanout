package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string; F1 int }
type GenS4 struct { F0 uint16; F1 uint16; F2 uint16 }
type GenS5 struct { F0 int64; F1 uint16 }
func genFn0(p0 int, p1 int64) int32 {
    for fi7 := 0; fi7 < 16; fi7++ {
        var v8 int16 = -32768
        if (!((0 <= 0) || true)) {
            switch p0 {
            case 127:
                var v9 uint8 = 42
                _ = v9
            case -1:
                var v10 string = "\n\t"
                _ = v10
            case 2147483647:
                var v11 string = "hello"
                _ = v11
            }
        } else {
            v8 = (v8 & -32768)
        }
        _ = fi7
        _ = v8
    }
    v12 := []int{p0, (p0 + p0), (p0 >> ((18446744073709551615) & 63))}
    _ = p0
    _ = p1
    _ = v12
    return -919
}
func genFn1(p0 uint64, p1 int16, p2 uint64) int8 {
    var v13 int16 = p1
    var v14 int8 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    return (v14 - (v14 << ((65535) & 63)))
}

func main() {
    var v1 int8 = 2
    var v2 uint16 = 42
    var v3 int16 = -114
    var v4 int32 = 174
    v5 := []int{32767, 0, -9223372036854775808}
    v6 := make(map[uint32]bool)
    v3 = v3
    v1 = -128
    v15 := v6[2147483647]
    if v15 {
        v5 = append(v5, 42)
    }
    if ((0.0 - -2.25) <= 0.0) {
        if false {
            for fi16 := 0; fi16 < 7; fi16++ {
                var v17 int8 = (v1 % ((1) | 1))
                if len(v5) > 0 {
                    v5[0] = fi16
                }
                _ = fi16
                _ = v17
            }
            for fi18 := 0; fi18 < 20; fi18++ {
                if (!(!false)) {
                    v19 := GenS1{F0: int16(v4), F1: (^(v3))}
                    v20 := v6[0]
                    _ = v19
                    _ = v20
                } else {
                    for fi21 := 0; fi21 < 2; fi21++ {
                        switch "a" {
                        case "hello":
                            if v15 {
                                v15 = (0.1 <= (-2.25 * 0.1))
                                v22 := GenS2{F0: -737}
                                if v15 {
                                    v23 := GenS1{F0: v3, F1: (-(v3))}
                                    fmt.Println(v3)
                                    var v24 bool = v15
                                    v24 = v24
                                    v25 := GenS3{F0: "", F1: fi21}
                                    _ = v23
                                    _ = v24
                                    _ = v25
                                } else {
                                    v26 := GenS1{F0: v3, F1: (v3 - (v3 % (((v3 + v3)) | 1)))}
                                    switch ("hello" + ("hello" + ("" + "\n\t"))) {
                                    case "":
                                        v6[4294967295] = (!(((1.5 - 0.1) / 0.1) != (1.5 - ((((1e10 / (3.14159 + 1e10)) / (float64(1e10) + 0.1)) / (-2.25 + -2.25)) / -2.25))))
                                    case "\n\t":
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    default:
                                        if v15 {
                                            var v27 bool = (true && false)
                                            v1 += (-(v1))
                                            _ = v27
                                        } else {
                                            if (!(fi21 <= 601)) {
                                                if ((v15 || false) && (!false)) {
                                                    v28 := GenS3{F0: ("hello" + ""), F1: -1}
                                                    v2 += v2
                                                    switch "foo bar" {
                                                    case "foo bar":
                                                        fmt.Println(fi21)
                                                        var v29 int16 = 42
                                                        _ = v29
                                                    case "hello":
                                                        v3 += v3
                                                    default:
                                                        var v30 int32 = genFn0(fi18, -1)
                                                        _ = v30
                                                    }
                                                    _ = v28
                                                } else {
                                                    v2 = (v2 - 438)
                                                    var v31 int32 = v4
                                                    _ = len(v5)
                                                    _ = v31
                                                }
                                                v32 := GenS2{F0: 483}
                                                switch fi18 {
                                                case -1:
                                                    var v33 int8 = int8(fi21)
                                                    _ = v33
                                                }
                                                _ = v32
                                            }
                                            var v34 int32 = (v4 >> ((1) & 63))
                                            switch (fi18 % ((73) | 1)) {
                                            case 1:
                                                v3 = int16(v34)
                                            case 32767:
                                                v35 := GenS4{F0: (^(v2)), F1: v2, F2: v2}
                                                _ = v35
                                            default:
                                                v36 := GenS3{F0: "foo bar", F1: fi18}
                                                _ = v36
                                            }
                                            _ = v34
                                        }
                                    }
                                    var v37 int32 = (v4 / ((-2147483648) | 1))
                                    _ = v26
                                    _ = v37
                                }
                                _ = v22
                            }
                        case "\n\t":
                            var v38 float32 = (0.1 / 3.14159)
                            v39 := GenS1{F0: (v3 << ((689) & 63)), F1: -320}
                            var v40 uint32 = 952
                            delete(v6, v40)
                            v41 := v6[v40]
                            v3 += 0
                            fmt.Println(v40, v2)
                            var v42 int16 = 32767
                            _ = v38
                            _ = v39
                            _ = v40
                            _ = v41
                            _ = v42
                        case "foo bar":
                            v5 = append(v5, fi18)
                        }
                        _ = fi21
                    }
                }
                _ = fi18
            }
        } else {
            var v43 string = ("\n\t" + "a")
            v5 = append(v5, 42)
            _ = v43
        }
        fmt.Println(v4, v1)
        if (!v15) {
            fmt.Println(v15)
            v44 := v6[0]
            _ = v44
        } else {
            if ((2147483647 <= v4) && false) {
                if v15 {
                    v15 = (!((!(!((128 < 157) && ((v1 ^ (v1 >> ((42) & 63))) >= (-(v1)))))) || v15))
                }
            } else {
                v2 -= (v2 & (v2 % (((^(v2))) | 1)))
            }
        }
    } else {
        for fi45 := 0; fi45 < 2; fi45++ {
            for fi46 := 0; fi46 < 16; fi46++ {
                fmt.Println(fi46, v2)
                var v47 float64 = (1.5 / 1.5)
                _ = fi46
                _ = v47
            }
            var v48 int64 = 0
            _ = fi45
            _ = v48
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        fmt.Println(v3, v3)
    }
    v49 := GenS2{F0: 0}
    var v50 int = 32768
    v15 = (42 == 277)
    delete(v6, 0)
    v2 = (v2 % ((32767) | 1))
    var v51 string = ""
    if true {
        fmt.Println(v1, v15, v3)
        switch "\n\t" {
        case "hello":
            fmt.Println(v50, v50, v2)
        default:
            for fi52 := 0; fi52 < 15; fi52++ {
                v53 := GenS2{F0: 912}
                v54 := GenS5{F0: 290, F1: (v2 | (v2 * (v2 / ((700) | 1))))}
                _ = len(v5)
                _ = fi52
                _ = v53
                _ = v54
            }
        }
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v55 := GenS4{F0: (v2 << ((1) & 63)), F1: v2, F2: (v2 & 983)}
    if (((!v15) || ((-88 >= (v1 + genFn1(4294967295, v3, 916))) || false)) || false) {
        v4 = v4
        v5 = append(v5, v50)
        if ((v50 & (v50 >> ((4294967295) & 63))) < (v50 ^ v50)) {
            for fi56 := 0; fi56 < 11; fi56++ {
                v50 -= v50
                for fi57 := 0; fi57 < 15; fi57++ {
                    var v58 uint8 = 0
                    v59 := GenS5{F0: -1, F1: v2}
                    if ((3.14159 < 0.0) && (((!((v15 && (!true)) || (!(872 != 774)))) || (v58 <= 255)) || (true && ((v3 >= v3) && (875 > 65535))))) {
                        v60 := GenS5{F0: -9223372036854775808, F1: (v2 * 255)}
                        _ = v60
                    }
                    _ = fi57
                    _ = v58
                    _ = v59
                }
                _ = fi56
            }
            switch ("foo bar" + v51) {
            case "":
                for fi61 := 0; fi61 < 2; fi61++ {
                    v50 = int(v4)
                    v3 += (v3 / (((v3 & (v3 >> ((835) & 63)))) | 1))
                    _ = fi61
                }
            case "hello":
                for fi62 := 0; fi62 < 17; fi62++ {
                    var v63 int32 = v4
                    _ = fi62
                    _ = v63
                }
            case "foo bar":
                v3 = v3
            }
        }
    }
    for fi64 := 0; fi64 < 5; fi64++ {
        switch (v51 + v51) {
        case "":
            var v65 uint32 = 716
            _ = v65
        case "a":
            delete(v6, 2147483647)
        case "foo bar":
            switch "\n\t" {
            case "a":
                v1 += 1
            case "":
                if v15 {
                    switch "hello" {
                    case "foo bar":
                        if v15 {
                            _ = len(v6)
                        } else {
                            v66 := v6[65535]
                            _ = v66
                        }
                    case "":
                        v2 -= v2
                    }
                    var v67 int = (v50 >> ((255) & 63))
                    v2 -= (^(v2))
                    _ = v67
                } else {
                    var v68 float32 = 0.0
                    var v69 uint16 = v2
                    _ = v68
                    _ = v69
                }
            default:
                switch v51 {
                case "a":
                    v70 := v6[42]
                    _ = v70
                case "foo bar":
                    var v71 int8 = genFn1(713, v3, 42)
                    _ = v71
                default:
                    for fi72 := 0; fi72 < 4; fi72++ {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        _ = fi72
                    }
                }
            }
        }
        var v73 int16 = -799
        _ = fi64
        _ = v73
    }
    fmt.Println(v2, v50)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v55
    _ = v6
}
