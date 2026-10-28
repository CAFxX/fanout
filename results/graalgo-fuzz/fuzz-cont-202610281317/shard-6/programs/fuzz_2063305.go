package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8; F1 uint8 }
type GenS4 struct { F0 int; F1 int32 }
type GenS5 struct { F0 int; F1 uint32 }
type GenS6 struct { F0 int16; F1 string; F2 uint16 }
type GenS7 struct { F0 uint64; F1 string; F2 bool }

func main() {
    v1 := make(map[string]string)
    var v2 float64 = float64(((float32(((1.5 - 1.5) * float64(1.5))) * 3.14159) - 0.1))
    v3 := make(map[string]string)
    v4 := make(map[string]string)
    if (1 > 788) {
        v2 = 1e10
    }
    v5 := []int{32767}
    if ((v2 / v2) != float64(0.0)) {
        if true {
            fmt.Println(v2)
            var v6 int8 = 42
            _ = v6
        }
        v7 := GenS1{F0: 42}
        _ = v7
    } else {
        v2 += 0.0
        if ((-347 != 805) || true) {
            if len(v5) > 0 {
                v5[0] = 596
            }
        } else {
            var v8 uint = 255
            switch 42 {
            case -1:
                for fi9 := 0; fi9 < 11; fi9++ {
                    v8 = 2
                    _ = fi9
                }
            default:
                switch 9223372036854775807 {
                case 4294967295:
                    var v10 int64 = -170
                    _ = v10
                }
            }
            _ = v8
        }
        if (false && (42 > 4294967295)) {
            for fi11 := 0; fi11 < 17; fi11++ {
                v12 := GenS2{F0: 42}
                _ = fi11
                _ = v12
            }
        } else {
            if (!false) {
                if (!(32767 <= -32768)) {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    for fi13 := 0; fi13 < 10; fi13++ {
                        if true {
                            if ((v2 / 1.0) > (v2 + (v2 * v2))) {
                                v2 = (float64(float32((3.14159 / v2))) - v2)
                                if true {
                                    var v14 float32 = 3.14159
                                    v14 = (3.14159 - v14)
                                    _ = v14
                                } else {
                                    delete(v1, "\n\t")
                                }
                                for fi15 := 0; fi15 < 8; fi15++ {
                                    var v16 bool = true
                                    v17 := v4[("" + ("hello" + "a"))]
                                    var v18 int64 = -885
                                    _ = fi15
                                    _ = v16
                                    _ = v17
                                    _ = v18
                                }
                            } else {
                                var v19 uint16 = 1
                                _ = v19
                            }
                            var v20 bool = (!((("hello" + (("foo bar" + "") + ("foo bar" + ("hello" + ("foo bar" + ("a" + "\n\t")))))) > ("" + "a")) || true))
                            _ = v20
                        } else {
                            v1["a"] = ("hello" + "a")
                            v2 += 3.14159
                        }
                        var v21 int8 = 42
                        for fi22 := 0; fi22 < 7; fi22++ {
                            fmt.Println(v2, fi22, fi22)
                            fmt.Println(v2, v2)
                            if (2147483648 > 2147483648) {
                                switch (fi22 / ((1) | 1)) {
                                case -2147483648:
                                    v23 := GenS2{F0: -984}
                                    _ = v23
                                case -207:
                                    var v24 string = "foo bar"
                                    _ = v24
                                default:
                                    _ = len(v3)
                                }
                                var v25 int8 = (v21 / (((v21 >> ((131) & 63))) | 1))
                                switch (fi13 + (fi22 ^ (fi13 & (fi22 << ((532) & 63))))) {
                                case -2147483648:
                                    var v26 uint64 = 0
                                    _ = v26
                                }
                                _ = v25
                            } else {
                                if (65535 >= 199) {
                                    fmt.Println(v21, v2, v2)
                                    v5 = append(v5, fi13)
                                    v2 -= 0.1
                                } else {
                                    if (0 <= 18446744073709551615) {
                                        v27 := v1["foo bar"]
                                        var v28 float32 = -2.25
                                        fmt.Println(fi22, fi13, v28)
                                        _ = v27
                                        _ = v28
                                    } else {
                                        v29 := GenS2{F0: -578}
                                        v30 := v3["a"]
                                        _ = v29
                                        _ = v30
                                    }
                                    var v31 string = "a"
                                    var v32 uint16 = 2
                                    var v33 uint64 = 765
                                    if (592 < (fi22 - -1)) {
                                        switch -777 {
                                        case 2147483647:
                                            v34 := GenS2{F0: int32(v33)}
                                            _ = v34
                                        }
                                        var v35 int8 = int8(fi13)
                                        _ = v35
                                    } else {
                                        if true {
                                            switch "hello" {
                                            case "hello":
                                                v21 -= int8(v33)
                                            case "":
                                                fmt.Println(fi22, v33, fi22)
                                            default:
                                                v21 = v21
                                            }
                                            v36 := GenS3{F0: v21, F1: 128}
                                            if (true && true) {
                                                fmt.Println(v2)
                                                var v37 uint16 = (v32 - v32)
                                                v37 += v37
                                                _ = v37
                                            }
                                            _ = v36
                                        }
                                        v2 += v2
                                    }
                                    var v38 int32 = -1
                                    _ = v31
                                    _ = v32
                                    _ = v33
                                    _ = v38
                                }
                                v39 := GenS4{F0: fi13, F1: -1}
                                _ = v39
                            }
                            v21 += (v21 >> ((0) & 63))
                            _ = len(v5)
                            _ = fi22
                        }
                        _ = fi13
                        _ = v21
                    }
                    switch 755 {
                    case -2147483648:
                        v4["a"] = ("\n\t" + "hello")
                    default:
                        var v40 int8 = -1
                        _ = v40
                    }
                } else {
                    if (18446744073709551615 <= 2) {
                        var v41 int8 = -88
                        _ = v41
                    } else {
                        for fi42 := 0; fi42 < 10; fi42++ {
                            switch -46 {
                            case 60:
                                for fi43 := 0; fi43 < 14; fi43++ {
                                    var v44 int8 = -4
                                    _ = fi43
                                    _ = v44
                                }
                            case 4294967295:
                                v45 := GenS1{F0: 2}
                                _ = v45
                            }
                            _ = fi42
                        }
                        var v46 int = 128
                        v5 = append(v5, (v46 | 2147483647))
                        _ = v46
                    }
                    switch "hello" {
                    case "foo bar":
                        var v47 uint32 = 1
                        _ = v47
                    case "":
                        v48 := GenS5{F0: 1, F1: 4294967295}
                        _ = v48
                    default:
                        v49 := GenS6{F0: 1000, F1: ("foo bar" + "a"), F2: 255}
                        _ = v49
                    }
                    for fi50 := 0; fi50 < 16; fi50++ {
                        for fi51 := 0; fi51 < 5; fi51++ {
                            v52 := GenS4{F0: fi51, F1: -2147483648}
                            v2 = (-2.25 / (v2 / ((v2 / v2) + (1e10 * v2))))
                            _ = fi51
                            _ = v52
                        }
                        fmt.Println(fi50)
                        v2 += 3.14159
                        _ = fi50
                    }
                }
                var v53 int8 = -66
                _ = v53
            }
            switch "a" {
            case "foo bar":
                v54 := GenS3{F0: 100, F1: 1}
                _ = v54
            case "a":
                v55 := GenS3{F0: -30, F1: 2}
                _ = v55
            case "hello":
                if false {
                    fmt.Println(v2)
                }
            }
        }
    }
    v2 = v2
    switch ("foo bar" + "foo bar") {
    case "\n\t":
        v2 = (v2 / v2)
    case "foo bar":
        v2 -= v2
    default:
        v2 -= (((v2 + v2) + float64(-2.25)) * (float64(3.14159) + v2))
    }
    for fi56 := 0; fi56 < 11; fi56++ {
        var v57 bool = true
        v2 -= float64(3.14159)
        _ = fi56
        _ = v57
    }
    v58 := GenS7{F0: 2, F1: "", F2: false}
    if true {
        fmt.Println(v2)
        for fi59 := 0; fi59 < 16; fi59++ {
            fmt.Println(fi59)
            _ = fi59
        }
        v2 = (v2 / (1.5 + (0.0 / 1.0)))
    } else {
        v60 := v4[("hello" + ("" + ("a" + "a")))]
        fmt.Println(v2, v2)
        if (255 <= 248) {
            v2 = (0.1 * v2)
            v60 = "foo bar"
            if false {
                var v61 int16 = 1
                var v62 uint32 = 0
                _ = v61
                _ = v62
            } else {
                if (true && (120 < 2)) {
                    if (!false) {
                        for fi63 := 0; fi63 < 14; fi63++ {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            _ = fi63
                        }
                        var v64 int64 = 9223372036854775807
                        if true {
                            v60 = v60
                            var v65 uint32 = 408
                            var v66 uint8 = 17
                            _ = v65
                            _ = v66
                        }
                        _ = v64
                    }
                    v2 += v2
                    v60 = v60
                }
                v2 -= (((0.1 / (v2 / (0.0 - ((v2 + (((v2 / -2.25) - -2.25) / float64(1.5))) / (1e10 + v2))))) * 1e10) - (v2 * (v2 + float64(3.14159))))
                var v67 uint16 = 2
                _ = v67
            }
        } else {
            var v68 int16 = 1000
            v60 = ("a" + v60)
            _ = v68
        }
        _ = v60
    }
    v5 = append(v5, 749)
    v69 := GenS5{F0: 42, F1: 65535}
    v2 -= float64((1e10 / 0.1))
    var v70 float64 = v2
    for fi71 := 0; fi71 < 13; fi71++ {
        _ = len(v3)
        _ = fi71
    }
    if true {
        var v72 uint32 = 534
        _ = v72
    }
    v2 += 0.1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v69
    _ = v70
}
