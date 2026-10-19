package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 string }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16; F1 int; F2 uint8 }
type GenS4 struct { F0 int64; F1 int8; F2 int64 }
type GenS5 struct { F0 int32; F1 int; F2 string }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int; F1 int32 }
type GenS8 struct { F0 int; F1 int16 }
type GenS9 struct { F0 uint32; F1 int64 }
type GenS10 struct { F0 bool; F1 uint8 }
func genFn0(p0 int) uint {
    p0 += p0
    if false {
        p0 -= 32768
    }
    _ = p0
    return 963
}

func main() {
    v1 := []uint32{65535, 900, 568}
    v2 := []string{(("" + "foo bar") + ("\n\t" + "a")), "foo bar", ("" + "")}
    v3 := make(map[uint32]bool)
    var v4 uint64 = 1
    v5 := make(map[int]bool)
    var v6 string = ("hello" + ("foo bar" + "hello"))
    v3[2147483648] = true
    for fi7 := 0; fi7 < 7; fi7++ {
        if (true || ((1.5 + (0.0 + 0.0)) == float32(float64(1e10)))) {
            var v8 int16 = -95
            fmt.Println(fi7, v6, v6)
            var v9 uint64 = 444
            _ = v8
            _ = v9
        } else {
            v4 = 330
            var v10 uint64 = 763
            if (((127 != -128) && (!true)) && true) {
                v6 = (v6 + "\n\t")
                _ = len(v2)
            } else {
                switch (-(fi7)) {
                case 9223372036854775807:
                    if len(v1) > 0 {
                        v1[0] = 4294967295
                    }
                default:
                    fmt.Println(v4, v4)
                }
                if false {
                    v11 := GenS1{F0: -251, F1: 0, F2: v6}
                    _ = v11
                }
                v10 += v10
            }
            _ = v10
        }
        v6 = ""
        _ = fi7
    }
    switch v6 {
    case "\n\t":
        switch v6 {
        case "\n\t":
            var v12 int8 = -1
            _ = v12
        case "":
            if (true || (false || (false || true))) {
                var v13 int32 = -1
                v4 = 4294967295
                _ = v13
            } else {
                v6 = v6
                var v14 bool = (-1 > -1)
                _ = v14
            }
        default:
            v15 := v5[-547]
            _ = v15
        }
    default:
        switch -143 {
        case -149:
            var v16 float64 = 1.5
            _ = v16
        case -9223372036854775808:
            v6 = v6
        case -901:
            var v17 uint = 0
            _ = v17
        }
    }
    v4 = (v4 ^ (v4 << ((134) & 63)))
    for fi18 := 0; fi18 < 2; fi18++ {
        for fi19 := 0; fi19 < 1; fi19++ {
            switch "foo bar" {
            case "hello":
                fmt.Println(v6)
            case "a":
                fmt.Println(v6, fi18, v6)
            }
            v1 = append(v1, 0)
            if false {
                var v20 string = (v6 + "")
                _ = v20
            }
            switch (("\n\t" + "foo bar") + "a") {
            case "hello":
                v6 = (((v6 + "") + v6) + "\n\t")
            case "a":
                v4 -= (v4 & v4)
            default:
                switch (fi19 % ((628) | 1)) {
                case 1:
                    var v21 int64 = 42
                    _ = v21
                case 2147483647:
                    v6 = v6
                case 32767:
                    var v22 string = v6
                    _ = v22
                }
            }
            _ = fi19
        }
        fmt.Println(fi18, fi18)
        _ = fi18
    }
    v23 := GenS2{F0: -2147483648}
    var v24 uint32 = 1
    v24 += v24
    if (42 < 1) {
        fmt.Println(v24, v6)
        fmt.Println(v6, v6)
        if true {
            switch "foo bar" {
            case "\n\t":
                v5[-232] = false
            default:
                if (true || (false && (v4 <= 18446744073709551615))) {
                    var v25 int8 = 127
                    _ = v25
                }
            }
            var v26 int16 = 675
            var v27 bool = true
            _ = v26
            _ = v27
        }
    } else {
        v28 := GenS3{F0: 32767, F1: 32768, F2: 3}
        v4 += (v4 & (v4 * (v4 / ((v4) | 1))))
        _ = v28
    }
    if false {
        if (-70 != -9223372036854775808) {
            for fi29 := 0; fi29 < 16; fi29++ {
                if false {
                    for fi30 := 0; fi30 < 2; fi30++ {
                        _ = len(v5)
                        if (!(!((float32((3.14159 / ((((0.0 - 0.1) - float64(-2.25)) / (0.0 - 0.1)) - 0.1))) + 0.1) != 3.14159))) {
                            v24 = (v24 - 2147483648)
                            var v31 int = (fi30 * fi30)
                            v32 := GenS4{F0: -277, F1: 127, F2: 42}
                            var v33 int64 = -1
                            _ = v31
                            _ = v32
                            _ = v33
                        }
                        v34 := GenS5{F0: 584, F1: fi30, F2: v6}
                        var v35 int32 = 42
                        var v36 uint = 18446744073709551615
                        if (!(!true)) {
                            switch fi29 {
                            case -754:
                                switch "hello" {
                                case "hello":
                                    var v37 uint = v36
                                    _ = v37
                                case "\n\t":
                                    fmt.Println(v6)
                                default:
                                    _ = len(v1)
                                    v38 := GenS6{F0: (v4 << (((^(v36))) & 63))}
                                    _ = v38
                                }
                            case 0:
                                fmt.Println(v24, v24)
                            case 32767:
                                var v39 int64 = -69
                                _ = v39
                            }
                            var v40 string = v6
                            _ = v40
                        } else {
                            if false {
                                _ = len(v5)
                                var v41 int32 = (v35 % ((int32(fi30)) | 1))
                                if false {
                                    var v42 string = "a"
                                    _ = v42
                                } else {
                                    v35 = -317
                                    v41 -= (v41 >> ((v36) & 63))
                                    if false {
                                        if (1000 < 42) {
                                            var v43 int16 = -944
                                            _ = v43
                                        } else {
                                            v44 := GenS7{F0: fi29, F1: -1}
                                            v45 := GenS5{F0: (v41 >> (((^(v36))) & 63)), F1: fi29, F2: v6}
                                            fmt.Println(v36, v4, v24)
                                            _ = v44
                                            _ = v45
                                        }
                                        if len(v1) > 0 {
                                            v1[0] = v24
                                        }
                                    }
                                }
                                _ = v41
                            } else {
                                if (v35 < (v35 - (-(v35)))) {
                                    var v46 int16 = 1
                                    if true {
                                        v4 = v4
                                        v47 := GenS8{F0: (^(fi29)), F1: (^(v46))}
                                        _ = v47
                                    } else {
                                        if ((v35 % ((v35) | 1)) <= 606) {
                                            var v48 bool = false
                                            _ = v48
                                        }
                                    }
                                    v49 := v5[(fi30 >> ((1) & 63))]
                                    v2 = append(v2, v6)
                                    _ = v46
                                    _ = v49
                                } else {
                                    v50 := GenS9{F0: (v24 | 645), F1: -1}
                                    v24 -= (^(v24))
                                    var v51 int16 = 42
                                    _ = v50
                                    _ = v51
                                }
                            }
                            switch (fi30 / ((fi29) | 1)) {
                            case 1:
                                var v52 uint16 = 1
                                _ = v52
                            case -1:
                                if (734 <= genFn0((fi29 % ((32768) | 1)))) {
                                    v4 = (v4 - v4)
                                    var v53 uint32 = 1
                                    v54 := GenS2{F0: (fi29 | fi29)}
                                    _ = v53
                                    _ = v54
                                }
                            }
                            switch ((v6 + v6) + "") {
                            case "foo bar":
                                v4 += (v4 * (^(v4)))
                            default:
                                fmt.Println(fi29)
                            }
                        }
                        delete(v3, (^(v24)))
                        _ = fi30
                        _ = v34
                        _ = v35
                        _ = v36
                    }
                }
                _ = fi29
            }
            var v55 int8 = 42
            if (true && true) {
                v56 := GenS3{F0: 1000, F1: 645, F2: 1}
                for fi57 := 0; fi57 < 2; fi57++ {
                    switch fi57 {
                    case 2147483647:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    case -9223372036854775808:
                        v2 = append(v2, v6)
                    }
                    _ = fi57
                }
                _ = v56
            } else {
                var v58 uint = 859
                _ = v58
            }
            _ = v55
        } else {
            for fi59 := 0; fi59 < 8; fi59++ {
                delete(v5, 0)
                fmt.Println(v24, v4, v4)
                _ = fi59
            }
        }
        var v60 int8 = 42
        for fi61 := 0; fi61 < 15; fi61++ {
            v6 = ((("hello" + ((v6 + v6) + v6)) + ("hello" + (v6 + "foo bar"))) + "foo bar")
            for fi62 := 0; fi62 < 8; fi62++ {
                v63 := GenS10{F0: (false || true), F1: 0}
                if (0.1 <= (-2.25 * 3.14159)) {
                    var v64 float32 = 1e10
                    _ = v64
                } else {
                    var v65 int32 = 808
                    v6 = "\n\t"
                    v24 = (v24 % (((v24 & (v24 % (((v24 % ((v24) | 1))) | 1)))) | 1))
                    _ = v65
                }
                _ = fi62
                _ = v63
            }
            var v66 string = v6
            _ = fi61
            _ = v66
        }
        _ = v60
    } else {
        var v67 int32 = 724
        switch -1 {
        case 68:
            var v68 uint32 = (v24 & v24)
            _ = v68
        case 9223372036854775807:
            var v69 uint64 = v4
            _ = v69
        }
        _ = v67
    }
    var v70 uint8 = 180
    v71 := GenS2{F0: 2147483647}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v70
    _ = v71
}
