package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 uint8 }
type GenS2 struct { F0 int64; F1 bool; F2 string }
type GenS3 struct { F0 int8; F1 int; F2 int64 }
type GenS4 struct { F0 int16; F1 int64 }
type GenS5 struct { F0 uint8; F1 string }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int64; F1 bool }
type GenS8 struct { F0 int32; F1 uint8 }
func genFn0(p0 uint, p1 int64) uint16 {
    p0 = (p0 / (((p0 >> (((p0 ^ 1)) & 63))) | 1))
    switch 744 {
    case 786:
        v6 := []int{}
        _ = v6
    case 128:
        var v7 int32 = 2147483647
        _ = v7
    case -874:
        p1 = (p1 ^ (p1 + (p1 & p1)))
    }
    _ = p0
    _ = p1
    return 0
}
func genFn1(p0 uint, p1 int16, p2 string) int8 {
    v8 := []int{0, -707}
    v9 := GenS1{F0: 943, F1: p0, F2: 2}
    p0 += (p0 & (^(p0)))
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    v10 := GenS1{F0: 411, F1: 1, F2: 121}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v8
    _ = v9
    return 1
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]bool)
    var v3 uint8 = 2
    var v4 bool = ((!true) && ((307 > 74) && true))
    var v5 int = 0
    if (!v4) {
        if (!false) {
            v5 = v5
            var v11 string = "\n\t"
            v3 += 103
            _ = v11
        } else {
            if v4 {
                if v4 {
                    var v12 bool = (2 != 65535)
                    v12 = (v12 && ((905 != 536) || false))
                    _ = v12
                } else {
                    for fi13 := 0; fi13 < 16; fi13++ {
                        for fi14 := 0; fi14 < 15; fi14++ {
                            var v15 int32 = 42
                            _ = fi14
                            _ = v15
                        }
                        v16 := GenS2{F0: -1, F1: true, F2: ("" + "a")}
                        _ = fi13
                        _ = v16
                    }
                    if (!v4) {
                        v3 = (^(v3))
                        if true {
                            var v17 float32 = (1e10 / float32(float64(1e10)))
                            _ = v17
                        }
                        for fi18 := 0; fi18 < 3; fi18++ {
                            if (!v4) {
                                var v19 uint8 = (v3 & (^(v3)))
                                _ = v19
                            }
                            var v20 uint8 = 128
                            if ((-9223372036854775808 != 42) || true) {
                                v5 = fi18
                                fmt.Println(v3)
                                v21 := []int{(fi18 - (^(fi18)))}
                                _ = v21
                            } else {
                                if (v4 || ((3.14159 * 0.0) >= ((float32((float64(0.1) - (0.1 / (-2.25 * (((3.14159 + float64(0.1)) / 3.14159) / 0.1))))) * 1.5) + 1e10))) {
                                    fmt.Println(v3)
                                    fmt.Println(fi18, v5, v4)
                                } else {
                                    for fi22 := 0; fi22 < 11; fi22++ {
                                        if v4 {
                                            var v23 int32 = -759
                                            fmt.Println(v3)
                                            var v24 int16 = 42
                                            _ = v23
                                            _ = v24
                                        }
                                        _ = fi22
                                    }
                                }
                                v25 := []int{(fi18 / (((v5 * fi18)) | 1)), fi18}
                                _ = v25
                            }
                            _ = fi18
                            _ = v20
                        }
                    } else {
                        var v26 int16 = 1000
                        for fi27 := 0; fi27 < 1; fi27++ {
                            fmt.Println(v5)
                            if v4 {
                                v4 = false
                                v28 := GenS2{F0: 42, F1: v4, F2: "hello"}
                                if (!false) {
                                    var v29 bool = v4
                                    switch (("" + "") + "") {
                                    case "\n\t":
                                        var v30 int16 = (v26 - (v26 & (v26 / ((v26) | 1))))
                                        _ = v30
                                    case "a":
                                        var v31 string = "foo bar"
                                        _ = v31
                                    case "hello":
                                        v32 := []int{-811, (fi27 & 42), v5}
                                        _ = v32
                                    default:
                                        v3 += 1
                                    }
                                    v33 := []int{0, v5, v5}
                                    _ = v29
                                    _ = v33
                                } else {
                                    var v34 float64 = -2.25
                                    v4 = (genFn1(1, (v26 & (v26 + (v26 | (v26 * 941)))), "hello") > genFn1(18446744073709551615, (^(v26)), ("foo bar" + "foo bar")))
                                    _ = v34
                                }
                                _ = v28
                            } else {
                                v35 := GenS3{F0: 8, F1: v5, F2: 0}
                                switch (fi27 ^ v5) {
                                case -1:
                                    v26 += (v26 >> ((18446744073709551615) & 63))
                                }
                                _ = v35
                            }
                            _ = fi27
                        }
                        _ = v26
                    }
                    var v36 string = ("a" + "")
                    _ = v36
                }
            }
        }
        v37 := []int{(v5 ^ 32767), v5, 577}
        for fi38 := 0; fi38 < 8; fi38++ {
            if len(v37) >= 2 {
                v37 = v37[1:len(v37)]
            }
            if v4 {
                var v39 uint16 = 42
                if v4 {
                    var v40 bool = (0 == -1)
                    if len(v37) > 0 {
                        v37[0] = v5
                    }
                    _ = v40
                }
                var v41 uint8 = 126
                _ = v39
                _ = v41
            } else {
                fmt.Println(v5, fi38)
                v42 := GenS3{F0: 42, F1: (v5 / ((-9223372036854775808) | 1)), F2: 900}
                for fi43 := 0; fi43 < 16; fi43++ {
                    if len(v37) >= 2 {
                        v37 = v37[1:len(v37)]
                    }
                    _ = fi43
                }
                _ = v42
            }
            _ = fi38
        }
        _ = v37
    } else {
        var v44 int8 = genFn1(259, 260, "a")
        v45 := v1[(^(v5))]
        var v46 int8 = v44
        _ = v44
        _ = v45
        _ = v46
    }
    v4 = v4
    for fi47 := 0; fi47 < 5; fi47++ {
        switch "\n\t" {
        case "foo bar":
            var v48 uint8 = (v3 ^ v3)
            _ = v48
        case "":
            v3 = (v3 & v3)
        case "hello":
            v4 = v4
        }
        v5 = (v5 << ((0) & 63))
        _ = fi47
    }
    switch "foo bar" {
    case "foo bar":
        fmt.Println(v4, v5, v5)
    case "":
        switch "foo bar" {
        case "":
            v5 = 928
        case "hello":
            for fi49 := 0; fi49 < 3; fi49++ {
                var v50 uint32 = 65535
                _ = fi49
                _ = v50
            }
        case "foo bar":
            v5 = v5
        default:
            v3 = (^(v3))
        }
    }
    v5 = 2147483647
    fmt.Println(v3)
    v3 = (^(v3))
    v5 += (v5 / (((v5 % (((v5 * v5)) | 1))) | 1))
    for fi51 := 0; fi51 < 13; fi51++ {
        v52 := GenS4{F0: 32767, F1: 0}
        var v53 int8 = genFn1(255, -1, ("a" + ("hello" + "\n\t")))
        if ((v53 % (((v53 % (((-(v53))) | 1))) | 1)) <= (-(v53))) {
            var v54 uint8 = v3
            _ = v54
        } else {
            if v4 {
                if (753 > 1) {
                    v55 := GenS5{F0: 255, F1: ("" + (((("a" + (("" + "hello") + "foo bar")) + "") + (("" + (("\n\t" + ("" + ("\n\t" + (("a" + "\n\t") + "")))) + "foo bar")) + "\n\t")) + ("a" + "\n\t")))}
                    _ = v55
                }
                var v56 int16 = 42
                for fi57 := 0; fi57 < 18; fi57++ {
                    switch (^(v5)) {
                    case -9223372036854775808:
                        v56 += v56
                    case 127:
                        v58 := []int{(fi51 % (((fi51 ^ (fi51 - fi51))) | 1)), (^(v5)), v5}
                        _ = v58
                    case 965:
                        fmt.Println(fi51)
                    }
                    _ = fi57
                }
                _ = v56
            } else {
                v59 := GenS6{F0: 1}
                fmt.Println(v4)
                _ = v59
            }
            if ("\n\t" > "foo bar") {
                for fi60 := 0; fi60 < 11; fi60++ {
                    delete(v2, ("a" + "foo bar"))
                    _ = len(v2)
                    _ = fi60
                }
            } else {
                if (!v4) {
                    v53 -= v53
                    if true {
                        switch fi51 {
                        case -1:
                            for fi61 := 0; fi61 < 10; fi61++ {
                                if false {
                                    if v4 {
                                        if v4 {
                                            if (0 < 32767) {
                                                v62 := []int{v5}
                                                _ = v62
                                            }
                                            v53 -= (v53 ^ (v53 - 42))
                                            v1[(^(fi51))] = 184
                                            v63 := GenS4{F0: 32767, F1: 1}
                                            _ = v63
                                        }
                                        var v64 int = (fi61 & v5)
                                        v65 := GenS2{F0: -9223372036854775808, F1: v4, F2: ("\n\t" + "")}
                                        _ = v64
                                        _ = v65
                                    }
                                }
                                v66 := []int{(fi51 % (((v5 - int(v53))) | 1)), (fi51 / ((fi61) | 1)), (fi61 % ((553) | 1))}
                                var v67 int16 = 42
                                var v68 uint64 = 606
                                _ = fi61
                                _ = v66
                                _ = v67
                                _ = v68
                            }
                        case 0:
                            var v69 uint32 = 42
                            _ = v69
                        default:
                            for fi70 := 0; fi70 < 14; fi70++ {
                                if v4 {
                                    v71 := []int{32768, (fi70 + (fi70 >> ((0) & 63))), fi70}
                                    var v72 uint32 = 750
                                    var v73 uint64 = 227
                                    _ = v71
                                    _ = v72
                                    _ = v73
                                }
                                v74 := v1[fi70]
                                if v4 {
                                    v75 := GenS7{F0: 26, F1: (65535 >= 1)}
                                    if (!v4) {
                                        v3 -= 127
                                        var v76 int = fi70
                                        var v77 string = ("hello" + "hello")
                                        v4 = v4
                                        _ = v76
                                        _ = v77
                                    }
                                    v78 := GenS8{F0: 42, F1: v3}
                                    _ = v75
                                    _ = v78
                                } else {
                                    var v79 uint64 = 2
                                    _ = v79
                                }
                                _ = fi70
                                _ = v74
                            }
                        }
                    }
                }
            }
            switch -379 {
            case 2147483647:
                for fi80 := 0; fi80 < 1; fi80++ {
                    if ((fi80 >> ((706) & 63)) != 128) {
                        var v81 uint64 = 2
                        _ = v81
                    }
                    var v82 int8 = int8(v3)
                    _ = fi80
                    _ = v82
                }
            case -258:
                var v83 uint = 255
                _ = v83
            default:
                fmt.Println(v4, fi51, v5)
            }
        }
        _ = fi51
        _ = v52
        _ = v53
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
