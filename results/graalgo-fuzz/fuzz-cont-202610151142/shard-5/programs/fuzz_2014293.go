package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int64; F1 uint32; F2 int }
type GenS5 struct { F0 uint8; F1 uint8 }
type GenS6 struct { F0 uint16 }

func main() {
    v1 := make(map[int]string)
    v2 := []uint32{}
    v3 := make(map[int]uint32)
    v4 := make(map[uint32]int)
    if true {
        var v5 uint8 = 143
        if true {
            v6 := GenS1{F0: 0}
            _ = v6
        }
        _ = v5
    } else {
        if (42 >= 1) {
            var v7 uint16 = 32767
            var v8 uint16 = 65535
            _ = v7
            _ = v8
        } else {
            var v9 uint = 0
            _ = v9
        }
        var v10 int = -820
        v10 += v10
        _ = v10
    }
    if (!true) {
        var v11 string = ("a" + "a")
        if true {
            switch (((v11 + (("hello" + v11) + v11)) + v11) + "hello") {
            case "":
                if (!(1 < 1)) {
                    for fi12 := 0; fi12 < 20; fi12++ {
                        v11 = ("hello" + (v11 + (v11 + v11)))
                        for fi13 := 0; fi13 < 8; fi13++ {
                            switch (v11 + v11) {
                            case "":
                                var v14 int16 = 42
                                _ = v14
                            case "foo bar":
                                switch ("" + v11) {
                                case "":
                                    fmt.Println(v11)
                                case "\n\t":
                                    var v15 float64 = 0.0
                                    _ = v15
                                case "hello":
                                    var v16 int = 0
                                    _ = v16
                                }
                            case "\n\t":
                                fmt.Println(v11, v11, fi13)
                            }
                            _ = fi13
                        }
                        _ = fi12
                    }
                    delete(v1, 1)
                } else {
                    switch "a" {
                    case "":
                        if len(v2) > 0 {
                            v2[0] = 2147483648
                        }
                    case "foo bar":
                        v11 = "a"
                    }
                    v11 = ("hello" + ((v11 + "foo bar") + v11))
                    switch v11 {
                    case "a":
                        if false {
                            var v17 int64 = 100
                            _ = v17
                        } else {
                            var v18 uint32 = 501
                            _ = v18
                        }
                    }
                }
            case "\n\t":
                v19 := v3[-957]
                _ = v19
            }
        }
        for fi20 := 0; fi20 < 5; fi20++ {
            if (42 > 911) {
                var v21 int = fi20
                v11 = "a"
                for fi22 := 0; fi22 < 5; fi22++ {
                    var v23 float32 = (((0.0 / 1.0) + 1e10) / 1.0)
                    v11 = "\n\t"
                    _ = fi22
                    _ = v23
                }
                _ = v21
            } else {
                v11 = v11
                var v24 string = ((v11 + (v11 + v11)) + "hello")
                if (false || (333 != 0)) {
                    delete(v1, -631)
                }
                _ = v24
            }
            v11 = "a"
            for fi25 := 0; fi25 < 15; fi25++ {
                switch 245 {
                case 0:
                    v26 := GenS1{F0: 1}
                    _ = v26
                case 556:
                    v11 = ("foo bar" + ((v11 + "foo bar") + v11))
                    if ((false && true) || (-2.25 > ((0.0 / 0.1) + -2.25))) {
                        var v27 uint64 = 1
                        var v28 bool = (!false)
                        _ = v27
                        _ = v28
                    } else {
                        fmt.Println(fi25, fi25)
                    }
                    v11 = ("\n\t" + ("a" + v11))
                case -230:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                default:
                    v29 := GenS2{F0: 100, F1: "\n\t"}
                    _ = v29
                }
                var v30 float64 = 0.0
                _ = fi25
                _ = v30
            }
            _ = fi20
        }
        _ = v11
    }
    for fi31 := 0; fi31 < 20; fi31++ {
        v32 := GenS1{F0: 9223372036854775807}
        _ = fi31
        _ = v32
    }
    var v33 uint8 = 207
    var v34 int = -9223372036854775808
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi35 := 0; fi35 < 5; fi35++ {
        v34 = -248
        var v36 int32 = -632
        switch v34 {
        case 32767:
            var v37 string = "a"
            _ = v37
        case 9223372036854775807:
            var v38 uint16 = 32767
            _ = v38
        default:
            if true {
                fmt.Println(v33)
            } else {
                v36 -= v36
            }
        }
        _ = fi35
        _ = v36
    }
    if ((v34 % ((789) | 1)) != v34) {
        v39 := GenS3{F0: 32767}
        _ = v39
    } else {
        fmt.Println(v33)
        if ((v33 >> ((892) & 63)) >= 255) {
            v33 -= (v33 + (v33 - v33))
            if false {
                v33 = v33
                _ = len(v3)
            }
            var v40 int = v34
            _ = v40
        }
    }
    switch (-(v34)) {
    case 32767:
        var v41 float64 = 3.14159
        _ = v41
    case -248:
        _ = len(v2)
    case -386:
        fmt.Println(v33)
    }
    for fi42 := 0; fi42 < 18; fi42++ {
        if (0 == uint32(v33)) {
            for fi43 := 0; fi43 < 14; fi43++ {
                v34 = (v34 << ((1) & 63))
                v34 = (v34 ^ (fi43 * (fi43 / ((2147483647) | 1))))
                v34 = (fi43 / (((fi42 << ((415) & 63))) | 1))
                _ = fi43
            }
            for fi44 := 0; fi44 < 7; fi44++ {
                var v45 int64 = 365
                fmt.Println(fi42, v34)
                var v46 uint = 983
                _ = fi44
                _ = v45
                _ = v46
            }
            if false {
                v47 := GenS4{F0: 0, F1: 42, F2: fi42}
                v33 -= v33
                _ = v47
            } else {
                v48 := GenS5{F0: v33, F1: 189}
                var v49 uint64 = 4294967295
                _ = v48
                _ = v49
            }
        } else {
            if len(v2) > 0 {
                v2[0] = 42
            }
            v50 := GenS3{F0: 65535}
            switch (fi42 + 747) {
            case 2147483647:
                if (-1 < -159) {
                    if len(v2) > 0 {
                        v2[0] = 4294967295
                    }
                    fmt.Println(v33, v33, fi42)
                } else {
                    switch (v34 - fi42) {
                    case -965:
                        v34 += 635
                    case 42:
                        _ = len(v2)
                    }
                    for fi51 := 0; fi51 < 1; fi51++ {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        if true {
                            v52 := GenS3{F0: 255}
                            v34 = 42
                            var v53 uint16 = 255
                            _ = v52
                            _ = v53
                        }
                        if true {
                            var v54 int = fi51
                            _ = v54
                        } else {
                            fmt.Println(v33, fi51)
                        }
                        _ = fi51
                    }
                    var v55 float32 = 3.14159
                    _ = v55
                }
            case 32767:
                for fi56 := 0; fi56 < 3; fi56++ {
                    var v57 uint = 255
                    var v58 int = 9223372036854775807
                    _ = fi56
                    _ = v57
                    _ = v58
                }
            }
            _ = v50
        }
        v59 := GenS6{F0: 2}
        _ = fi42
        _ = v59
    }
    for fi60 := 0; fi60 < 13; fi60++ {
        delete(v1, (v34 + (v34 | (fi60 << ((180) & 63)))))
        _ = fi60
    }
    for fi61 := 0; fi61 < 9; fi61++ {
        fmt.Println(fi61)
        _ = fi61
    }
    fmt.Println(v34)
    var v62 float32 = 1.5
    switch 32768 {
    case -1:
        var v63 uint16 = 2
        _ = v63
    case 128:
        var v64 int8 = 105
        _ = v64
    case 4294967295:
        v34 -= v34
    default:
        _ = len(v2)
    }
    v65 := GenS5{F0: v33, F1: (v33 % (((v33 - (v33 + v33))) | 1))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v62
    _ = v65
}
