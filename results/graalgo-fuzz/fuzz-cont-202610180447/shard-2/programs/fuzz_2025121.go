package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 uint16 }
type GenS2 struct { F0 bool; F1 uint; F2 int32 }
type GenS3 struct { F0 int; F1 int32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint8; F1 uint8 }
type GenS6 struct { F0 bool; F1 int }
type GenS7 struct { F0 uint64; F1 uint }
type GenS8 struct { F0 uint; F1 int32; F2 int16 }

func main() {
    v1 := []uint32{2, 117, 38}
    v2 := make(map[uint32]uint32)
    v3 := []int{518, 207, 2147483647, 42}
    var v4 string = "foo bar"
    if false {
        delete(v2, 34)
    } else {
        var v5 float64 = 1e10
        var v6 uint = 753
        for fi7 := 0; fi7 < 19; fi7++ {
            v8 := GenS1{F0: 42, F1: ((v4 + v4) + v4), F2: 2}
            _ = fi7
            _ = v8
        }
        _ = v5
        _ = v6
    }
    switch 1 {
    case 42:
        var v9 bool = false
        _ = v9
    case -1:
        var v10 int8 = 0
        _ = v10
    case 1:
        for fi11 := 0; fi11 < 6; fi11++ {
            if true {
                v4 = (("" + ((v4 + v4) + v4)) + (v4 + v4))
                var v12 uint16 = 247
                _ = v12
            } else {
                for fi13 := 0; fi13 < 1; fi13++ {
                    if false {
                        var v14 int16 = 162
                        var v15 uint16 = 42
                        var v16 bool = true
                        _ = v14
                        _ = v15
                        _ = v16
                    } else {
                        var v17 float64 = float64(-2.25)
                        _ = v17
                    }
                    _ = fi13
                }
                v18 := v2[2147483648]
                if (1 != -1) {
                    var v19 float32 = (0.0 * 3.14159)
                    v20 := GenS1{F0: 156, F1: (v4 + ("hello" + "")), F2: 548}
                    fmt.Println(v4)
                    _ = v19
                    _ = v20
                } else {
                    var v21 int32 = 211
                    _ = v21
                }
                _ = v18
            }
            _ = fi11
        }
    }
    v1 = append(v1, 795)
    for fi22 := 0; fi22 < 7; fi22++ {
        v23 := GenS2{F0: (1 < 255), F1: 524, F2: -961}
        if true {
            var v24 int8 = 100
            for fi25 := 0; fi25 < 16; fi25++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                if true {
                    v26 := GenS3{F0: fi22, F1: 2147483647}
                    v27 := GenS4{F0: true}
                    var v28 uint = 4294967295
                    _ = v26
                    _ = v27
                    _ = v28
                }
                var v29 uint = 4294967295
                v4 = v4
                _ = fi25
                _ = v29
            }
            _ = v24
        } else {
            v30 := GenS5{F0: 1, F1: 255}
            v4 = v4
            _ = v30
        }
        for fi31 := 0; fi31 < 16; fi31++ {
            var v32 int8 = -128
            var v33 uint = 18446744073709551615
            _ = fi31
            _ = v32
            _ = v33
        }
        _ = fi22
        _ = v23
    }
    var v34 int8 = 78
    if false {
        v35 := GenS4{F0: (65535 != 0)}
        if (3.14159 > 0.0) {
            v4 = "foo bar"
            var v36 uint8 = 2
            for fi37 := 0; fi37 < 20; fi37++ {
                for fi38 := 0; fi38 < 1; fi38++ {
                    v36 += (v36 + v36)
                    v4 = v4
                    _ = fi38
                }
                switch (-(fi37)) {
                case -635:
                    var v39 uint = 4294967295
                    _ = v39
                }
                v4 = (v4 + (((v4 + v4) + ("a" + ("a" + (((v4 + v4) + v4) + "")))) + v4))
                _ = fi37
            }
            _ = v36
        }
        _ = v35
    } else {
        v4 = "foo bar"
        v40 := GenS5{F0: 240, F1: 205}
        _ = v40
    }
    for fi41 := 0; fi41 < 2; fi41++ {
        _ = len(v3)
        _ = fi41
    }
    for fi42 := 0; fi42 < 1; fi42++ {
        var v43 float32 = float32(1.5)
        _ = fi42
        _ = v43
    }
    fmt.Println(v34, v4)
    var v44 float64 = ((0.0 - 3.14159) * 0.0)
    if true {
        v44 = 0.0
        if false {
            for fi45 := 0; fi45 < 6; fi45++ {
                if (true || (fi45 >= fi45)) {
                    v44 -= v44
                    v46 := GenS1{F0: 18446744073709551615, F1: (v4 + v4), F2: 878}
                    _ = v46
                } else {
                    for fi47 := 0; fi47 < 20; fi47++ {
                        fmt.Println(v34, fi45, fi45)
                        v48 := GenS6{F0: (0 < 487), F1: 128}
                        var v49 float32 = (-2.25 + -2.25)
                        _ = fi47
                        _ = v48
                        _ = v49
                    }
                }
                _ = fi45
            }
            if false {
                v50 := GenS5{F0: 0, F1: 128}
                fmt.Println(v44, v34)
                v1 = append(v1, 65535)
                _ = v50
            } else {
                v51 := GenS7{F0: 18446744073709551615, F1: 255}
                if (!(false || true)) {
                    fmt.Println(v44, v4)
                }
                var v52 uint = 18446744073709551615
                _ = v51
                _ = v52
            }
        }
        _ = len(v2)
    } else {
        v44 += 0.1
    }
    if (!(2 == 74)) {
        v4 = "\n\t"
        if ((-122 < (v34 ^ v34)) || true) {
            if true {
                var v53 uint8 = 168
                var v54 float32 = float32(v44)
                _ = v53
                _ = v54
            }
            v44 = (v44 * 1.5)
        } else {
            for fi55 := 0; fi55 < 15; fi55++ {
                switch -511 {
                case 276:
                    switch fi55 {
                    case 127:
                        v56 := GenS4{F0: true}
                        _ = v56
                    case 9223372036854775807:
                        if true {
                            var v57 bool = false
                            _ = v57
                        }
                    default:
                        var v58 uint32 = 2147483648
                        _ = v58
                    }
                case 127:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                case 0:
                    _ = len(v2)
                }
                fmt.Println(v44, fi55)
                _ = fi55
            }
            v59 := GenS5{F0: 0, F1: 172}
            if len(v1) > 0 {
                v1[0] = 2147483648
            }
            _ = v59
        }
        for fi60 := 0; fi60 < 16; fi60++ {
            v34 = 0
            v34 -= 1
            fmt.Println(v34, v4)
            _ = fi60
        }
    } else {
        switch v4 {
        case "\n\t":
            if (65535 == 2147483648) {
                var v61 int64 = 192
                _ = v61
            } else {
                v4 = v4
                var v62 uint = 1
                fmt.Println(v34, v34)
                _ = v62
            }
        case "":
            for fi63 := 0; fi63 < 2; fi63++ {
                if false {
                    switch "" {
                    case "\n\t":
                        for fi64 := 0; fi64 < 4; fi64++ {
                            if false {
                                switch v4 {
                                case "a":
                                    switch (fi64 / ((fi64) | 1)) {
                                    case 128:
                                        var v65 float64 = v44
                                        v66 := GenS6{F0: (65535 != 2), F1: (fi63 | fi64)}
                                        _ = len(v2)
                                        _ = v65
                                        _ = v66
                                    default:
                                        var v67 uint16 = 2
                                        _ = v67
                                    }
                                default:
                                    v44 = v44
                                }
                                v4 = "a"
                            } else {
                                var v68 float64 = (1.5 * v44)
                                switch "a" {
                                case "foo bar":
                                    if false {
                                        var v69 uint = 13
                                        if false {
                                            v44 = (v68 - (v44 - v68))
                                            v68 = (((((v68 - (v68 * 1e10)) / float64(1e10)) + v68) / 1e10) * 1.5)
                                            var v70 float32 = 0.0
                                            _ = v70
                                        }
                                        v71 := GenS8{F0: v69, F1: 1, F2: 0}
                                        _ = v69
                                        _ = v71
                                    }
                                }
                                _ = v68
                            }
                            _ = fi64
                        }
                    case "a":
                        var v72 string = ""
                        _ = v72
                    case "hello":
                        fmt.Println(v44, v44, v34)
                    }
                } else {
                    v73 := GenS6{F0: false, F1: fi63}
                    _ = v73
                }
                for fi74 := 0; fi74 < 12; fi74++ {
                    v44 += (1e10 + (0.1 - v44))
                    _ = fi74
                }
                var v75 int16 = 550
                _ = fi63
                _ = v75
            }
        case "foo bar":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
    }
    var v76 uint = 18446744073709551615
    var v77 int16 = 317
    v78 := v2[4294967295]
    _ = len(v3)
    v79 := GenS7{F0: uint64(v78), F1: 2}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v76:", v76)
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v44
    _ = v76
    _ = v77
    _ = v78
    _ = v79
}
