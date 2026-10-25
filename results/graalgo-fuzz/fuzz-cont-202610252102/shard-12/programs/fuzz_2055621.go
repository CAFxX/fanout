package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64; F1 uint8 }
type GenS4 struct { F0 uint16; F1 bool }
type GenS5 struct { F0 int16; F1 string }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int32 }
func genFn0(p0 int, p1 uint16, p2 int) uint8 {
    v6 := []int{}
    switch (p2 >> ((449) & 63)) {
    case 862:
        v7 := GenS1{F0: 1, F1: 1}
        _ = v7
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 255
}

func main() {
    var v1 int32 = -35
    v2 := []int{-539, 42}
    var v3 uint = 42
    var v4 int64 = 0
    var v5 float32 = 3.14159
    var v8 int8 = 42
    if true {
        v2 = append(v2, 741)
        var v9 uint64 = 42
        _ = v9
    }
    for fi10 := 0; fi10 < 2; fi10++ {
        fmt.Println(v3, v3)
        _ = fi10
    }
    if false {
        if (0 <= 0) {
            if true {
                v11 := GenS2{F0: 32767}
                for fi12 := 0; fi12 < 12; fi12++ {
                    for fi13 := 0; fi13 < 17; fi13++ {
                        var v14 float32 = float32(1e10)
                        _ = fi13
                        _ = v14
                    }
                    if len(v2) > 0 {
                        v2[0] = fi12
                    }
                    _ = fi12
                }
                _ = v11
            }
            for fi15 := 0; fi15 < 12; fi15++ {
                v16 := GenS3{F0: -1, F1: 0}
                switch ("hello" + ("foo bar" + "\n\t")) {
                case "foo bar":
                    fmt.Println(v3)
                case "a":
                    fmt.Println(v3)
                default:
                    v1 -= 42
                }
                v2 = append(v2, 128)
                _ = fi15
                _ = v16
            }
        }
        var v17 int16 = 1000
        var v18 int = 1
        _ = v17
        _ = v18
    }
    if true {
        v8 = v8
        if ((v1 != (v1 | v1)) && false) {
            v3 = 2
            for fi19 := 0; fi19 < 6; fi19++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                var v20 int = fi19
                _ = fi19
                _ = v20
            }
        }
    } else {
        var v21 uint32 = 2147483647
        for fi22 := 0; fi22 < 12; fi22++ {
            for fi23 := 0; fi23 < 8; fi23++ {
                v4 = v4
                var v24 uint = (v3 >> (((v3 ^ v3)) & 63))
                var v25 int16 = 0
                _ = fi23
                _ = v24
                _ = v25
            }
            _ = fi22
        }
        switch "foo bar" {
        case "foo bar":
            switch "" {
            case "\n\t":
                for fi26 := 0; fi26 < 6; fi26++ {
                    for fi27 := 0; fi27 < 17; fi27++ {
                        v5 = (v5 - (((0.0 / 1.5) / v5) * v5))
                        if true {
                            switch (fi27 & (-(fi27))) {
                            case -9223372036854775808:
                                fmt.Println(v21, v8)
                            case 32768:
                                fmt.Println(v8)
                            }
                            v5 = v5
                        } else {
                            if true {
                                v28 := GenS1{F0: (v3 & (^(v3))), F1: (v8 % ((v8) | 1))}
                                if true {
                                    v29 := GenS3{F0: (v4 * (v4 / ((9223372036854775807) | 1))), F1: 212}
                                    var v30 uint16 = 80
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    _ = v29
                                    _ = v30
                                }
                                _ = v28
                            } else {
                                fmt.Println(fi26, v8)
                            }
                            var v31 string = ""
                            fmt.Println(fi26)
                            _ = v31
                        }
                        _ = fi27
                    }
                    _ = fi26
                }
            }
        case "":
            var v32 float64 = (3.14159 * (-2.25 - (0.0 / 1.5)))
            _ = v32
        case "hello":
            for fi33 := 0; fi33 < 11; fi33++ {
                v3 = 65535
                if false {
                    if len(v2) > 0 {
                        v2[0] = (fi33 >> ((2) & 63))
                    }
                    fmt.Println(v21, fi33)
                    var v34 float64 = 0.0
                    _ = v34
                } else {
                    fmt.Println(fi33)
                    fmt.Println(v4)
                    for fi35 := 0; fi35 < 13; fi35++ {
                        var v36 int16 = -32768
                        v37 := GenS1{F0: v3, F1: 100}
                        _ = fi35
                        _ = v36
                        _ = v37
                    }
                }
                var v38 int16 = 0
                _ = fi33
                _ = v38
            }
        default:
            v2 = append(v2, 42)
        }
        _ = v21
    }
    if (!(false || (273 == int16(v3)))) {
        v39 := GenS4{F0: 32767, F1: true}
        v1 += (v1 & -132)
        _ = v39
    } else {
        switch "foo bar" {
        case "\n\t":
            if (1 >= 886) {
                switch 372 {
                case -2147483648:
                    v2 = append(v2, -999)
                }
                if (!true) {
                    v40 := GenS1{F0: v3, F1: (^(v8))}
                    var v41 float32 = v5
                    _ = v40
                    _ = v41
                } else {
                    for fi42 := 0; fi42 < 13; fi42++ {
                        v2 = append(v2, (^(fi42)))
                        switch fi42 {
                        case 280:
                            if false {
                                v1 = 2147483647
                                switch "" {
                                case "a":
                                    var v43 uint8 = 2
                                    _ = v43
                                case "":
                                    for fi44 := 0; fi44 < 14; fi44++ {
                                        switch (fi42 % (((fi42 * (fi44 + 4294967295))) | 1)) {
                                        case -814:
                                            var v45 int64 = (v4 >> ((v3) & 63))
                                            _ = v45
                                        case -395:
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        }
                                        fmt.Println(v5, fi42, v3)
                                        fmt.Println(fi42, v8, v1)
                                        var v46 int = 9223372036854775807
                                        var v47 uint64 = 2
                                        _ = fi44
                                        _ = v46
                                        _ = v47
                                    }
                                default:
                                    v8 = 1
                                }
                            } else {
                                var v48 float64 = -2.25
                                if len(v2) > 0 {
                                    v2[0] = -1
                                }
                                if true {
                                    for fi49 := 0; fi49 < 13; fi49++ {
                                        var v50 uint16 = 0
                                        v1 -= (^(v1))
                                        v3 = 601
                                        _ = fi49
                                        _ = v50
                                    }
                                    for fi51 := 0; fi51 < 5; fi51++ {
                                        v52 := GenS1{F0: v3, F1: (v8 | (v8 / ((-1) | 1)))}
                                        v53 := GenS2{F0: 0}
                                        if false {
                                            fmt.Println(v4)
                                        }
                                        _ = fi51
                                        _ = v52
                                        _ = v53
                                    }
                                } else {
                                    if len(v2) > 0 {
                                        v2[0] = fi42
                                    }
                                    for fi54 := 0; fi54 < 16; fi54++ {
                                        v55 := GenS3{F0: v4, F1: 1}
                                        var v56 bool = true
                                        if (v56 || (v48 >= ((0.1 / v48) + v48))) {
                                            var v57 bool = true
                                            var v58 int32 = v1
                                            if v57 {
                                                fmt.Println(v48, v1, v1)
                                                var v59 uint8 = 255
                                                _ = v59
                                            }
                                            _ = v57
                                            _ = v58
                                        }
                                        _ = fi54
                                        _ = v55
                                        _ = v56
                                    }
                                }
                                _ = v48
                            }
                        case -453:
                            v8 += v8
                        case -220:
                            switch fi42 {
                            case 128:
                                for fi60 := 0; fi60 < 15; fi60++ {
                                    fmt.Println(fi42)
                                    v61 := GenS5{F0: 42, F1: "hello"}
                                    if (v1 < (v1 & (v1 & (^(v1))))) {
                                        v62 := GenS6{F0: 42}
                                        v4 = v4
                                        _ = v62
                                    }
                                    _ = fi60
                                    _ = v61
                                }
                            case -353:
                                var v63 int16 = int16(v4)
                                _ = v63
                            default:
                                if len(v2) > 0 {
                                    v2[0] = 4294967295
                                }
                            }
                        default:
                            var v64 int64 = v4
                            _ = v64
                        }
                        if false {
                            fmt.Println(v8)
                            for fi65 := 0; fi65 < 19; fi65++ {
                                _ = len(v2)
                                _ = fi65
                            }
                            for fi66 := 0; fi66 < 18; fi66++ {
                                var v67 uint64 = 232
                                v68 := GenS4{F0: 741, F1: true}
                                v1 = v1
                                _ = fi66
                                _ = v67
                                _ = v68
                            }
                        } else {
                            v1 = v1
                        }
                        _ = fi42
                    }
                    var v69 int8 = (v8 ^ (v8 & v8))
                    _ = v69
                }
            }
        }
    }
    v70 := GenS7{F0: 42}
    v8 = v8
    v71 := GenS3{F0: 933, F1: 161}
    v3 = (v3 << (((v3 ^ 255)) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v70
    _ = v71
    _ = v8
}
