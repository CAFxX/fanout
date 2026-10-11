package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool; F1 int8 }
type GenS4 struct { F0 int; F1 uint8 }
type GenS5 struct { F0 uint32; F1 uint; F2 int64 }
type GenS6 struct { F0 int; F1 uint8 }
type GenS7 struct { F0 int16 }
func genFn0(p0 uint, p1 string, p2 int) int8 {
    if false {
        var v7 int64 = int64(p0)
        if false {
            var v8 int64 = v7
            _ = v8
        } else {
            fmt.Println(p1)
            v9 := GenS1{F0: 495, F1: 42}
            _ = v9
        }
        p1 = p1
        _ = v7
    } else {
        p0 = (p0 ^ (p0 >> (((p0 / (((p0 >> ((p0) & 63))) | 1))) & 63)))
    }
    v10 := GenS1{F0: 255, F1: -1}
    v11 := []int{127, (-(p2)), -99}
    switch "\n\t" {
    case "foo bar":
        if true {
            var v12 int16 = 32767
            for fi13 := 0; fi13 < 12; fi13++ {
                var v14 string = p1
                _ = fi13
                _ = v14
            }
            if false {
                p1 = p1
                fmt.Println(p1)
                v15 := GenS1{F0: 1, F1: (v12 + (-(v12)))}
                _ = v15
            } else {
                _ = len(v11)
            }
            _ = v12
        } else {
            p0 -= (p0 ^ p0)
            v16 := GenS2{F0: ((!true) || ((42 == 32767) && (-1 >= 29)))}
            if false {
                var v17 float32 = 1.5
                var v18 bool = (true || false)
                if len(v11) > 0 {
                    v11[0] = p2
                }
                _ = v17
                _ = v18
            }
            _ = v16
        }
    default:
        var v19 uint16 = 1
        _ = v19
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return 127
}
func genFn1(p0 uint64, p1 uint64) uint {
    for fi20 := 0; fi20 < 12; fi20++ {
        if (genFn0(971, "\n\t", fi20) < 1) {
            p0 = 4294967295
        }
        _ = fi20
    }
    fmt.Println(p0, p1)
    var v21 bool = true
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v21
    return 2
}

func main() {
    var v1 uint32 = 2147483648
    var v2 int32 = 0
    v3 := []int{2147483647, 815, -1, -1}
    var v4 string = (("\n\t" + "\n\t") + ("hello" + "hello"))
    var v5 string = v4
    v6 := make(map[int]uint32)
    if (14 == 0) {
        var v22 string = ("\n\t" + v4)
        v1 = (^(v1))
        fmt.Println(v22, v1, v4)
        _ = v22
    } else {
        v4 = "\n\t"
        v5 = v4
    }
    if (false || false) {
        switch ("\n\t" + v4) {
        case "hello":
            v23 := GenS2{F0: ((18446744073709551615 < 0) || (v5 == v4))}
            _ = v23
        case "a":
            for fi24 := 0; fi24 < 1; fi24++ {
                v2 += 943
                var v25 bool = (-128 < 1)
                var v26 float64 = -2.25
                _ = fi24
                _ = v25
                _ = v26
            }
        }
        v27 := GenS3{F0: false, F1: 0}
        v28 := GenS4{F0: 32767, F1: uint8(v1)}
        _ = v27
        _ = v28
    }
    v6[-87] = (v1 + 734)
    switch v5 {
    case "\n\t":
        var v29 uint32 = v1
        _ = v29
    case "foo bar":
        if (255 < 128) {
            if (0 < 0) {
                var v30 float64 = 1e10
                if (0 > v2) {
                    v31 := GenS5{F0: v1, F1: 18446744073709551615, F2: -751}
                    _ = len(v6)
                    _ = v31
                } else {
                    v1 += v1
                    var v32 float32 = (float32(v30) - 0.0)
                    _ = v32
                }
                v33 := GenS3{F0: (!(v1 != 1)), F1: -128}
                _ = v30
                _ = v33
            }
        }
    case "hello":
        if (0 >= 178) {
            for fi34 := 0; fi34 < 19; fi34++ {
                for fi35 := 0; fi35 < 15; fi35++ {
                    v36 := GenS1{F0: 65535, F1: 42}
                    _ = fi35
                    _ = v36
                }
                v1 += 966
                if false {
                    var v37 int = fi34
                    var v38 uint8 = 2
                    v39 := GenS1{F0: 358, F1: -170}
                    _ = v37
                    _ = v38
                    _ = v39
                }
                _ = fi34
            }
            var v40 uint16 = 179
            v1 -= v1
            _ = v40
        } else {
            if (false || (-2.25 >= 0.0)) {
                v4 = "\n\t"
            } else {
                if (695 != 18446744073709551615) {
                    if false {
                        switch "\n\t" {
                        case "a":
                            v5 = v4
                        case "\n\t":
                            var v41 int = 681
                            _ = v41
                        case "hello":
                            for fi42 := 0; fi42 < 18; fi42++ {
                                v5 = "hello"
                                _ = fi42
                            }
                        }
                    } else {
                        v4 = v5
                        v5 = ("\n\t" + "hello")
                    }
                } else {
                    if (!(-1 <= 1)) {
                        var v43 int64 = 43
                        if len(v3) > 0 {
                            v3[0] = 127
                        }
                        _ = v43
                    } else {
                        var v44 string = (v4 + "hello")
                        var v45 float32 = float32(0.0)
                        for fi46 := 0; fi46 < 8; fi46++ {
                            switch int(v2) {
                            case 32767:
                                v47 := GenS2{F0: true}
                                _ = v47
                            case 127:
                                v5 = ""
                            default:
                                var v48 uint8 = 1
                                _ = v48
                            }
                            var v49 uint32 = v1
                            _ = fi46
                            _ = v49
                        }
                        _ = v44
                        _ = v45
                    }
                }
                if true {
                    var v50 int16 = 597
                    var v51 string = ("hello" + v4)
                    var v52 uint = genFn1(42, 614)
                    _ = v50
                    _ = v51
                    _ = v52
                } else {
                    for fi53 := 0; fi53 < 1; fi53++ {
                        v54 := GenS1{F0: 306, F1: -32768}
                        _ = fi53
                        _ = v54
                    }
                    v55 := GenS6{F0: 42, F1: 70}
                    _ = v55
                }
                v2 = v2
            }
            if len(v3) > 0 {
                v3[0] = 32768
            }
            var v56 int = -9223372036854775808
            _ = v56
        }
    }
    var v57 int = -759
    if false {
        fmt.Println(v2, v2, v1)
        if false {
            for fi58 := 0; fi58 < 18; fi58++ {
                var v59 string = "\n\t"
                v59 = v5
                _ = fi58
                _ = v59
            }
            var v60 bool = true
            v61 := GenS7{F0: -1}
            _ = v60
            _ = v61
        } else {
            var v62 uint16 = 0
            v2 += (-(v2))
            _ = v62
        }
    } else {
        for fi63 := 0; fi63 < 5; fi63++ {
            if false {
                v4 = v5
                var v64 uint8 = 128
                var v65 uint8 = (^(v64))
                _ = v64
                _ = v65
            } else {
                var v66 uint = 59
                _ = len(v3)
                var v67 int64 = -9223372036854775808
                _ = v66
                _ = v67
            }
            v4 = "hello"
            var v68 uint8 = 127
            _ = fi63
            _ = v68
        }
    }
    v57 += (v57 ^ (v57 / ((128) | 1)))
    fmt.Println(v57, v1)
    var v69 int16 = 1000
    var v70 uint = 4294967295
    v71 := GenS2{F0: ((^(v1)) <= (v1 << ((2) & 63)))}
    v5 = (v5 + v5)
    switch (-(v57)) {
    case 42:
        for fi72 := 0; fi72 < 11; fi72++ {
            if (((v4 + (v4 + "a")) + (v5 + ((v4 + v4) + "\n\t"))) > v4) {
                v70 = (v70 | v70)
                v73 := GenS1{F0: 715, F1: (v69 + (v69 / ((v69) | 1)))}
                var v74 float32 = -2.25
                _ = v73
                _ = v74
            } else {
                fmt.Println(v4, v2)
                if (!true) {
                    var v75 bool = true
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v1 += 1
                    _ = v75
                }
                v76 := v6[42]
                _ = v76
            }
            v4 = v5
            _ = fi72
        }
    case -778:
        v6[(^(v57))] = uint32(v70)
    default:
        v77 := GenS5{F0: (v1 / ((v1) | 1)), F1: v70, F2: 42}
        _ = v77
    }
    for fi78 := 0; fi78 < 11; fi78++ {
        if (true || false) {
            switch "" {
            case "hello":
                v1 = 0
            case "foo bar":
                switch (v57 * (v57 + 273)) {
                case 0:
                    var v79 uint = 255
                    _ = v79
                case 846:
                    var v80 int64 = -1
                    _ = v80
                default:
                    fmt.Println(fi78)
                }
            }
        } else {
            v5 = "foo bar"
            v1 = v1
            if true {
                var v81 uint32 = v1
                _ = v81
            } else {
                v57 += 377
                v3 = append(v3, (fi78 ^ v57))
                v3 = append(v3, (v57 + (v57 >> ((255) & 63))))
            }
        }
        _ = fi78
    }
    fmt.Println(v69)
    v1 -= v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v6
    _ = v69
    _ = v70
    _ = v71
}
