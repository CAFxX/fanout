package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int8; F1 int64; F2 uint8 }
type GenS5 struct { F0 uint16; F1 int; F2 uint32 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int; F1 int64; F2 uint8 }
func genFn0() int32 {
    var v5 int64 = 0
    for fi6 := 0; fi6 < 19; fi6++ {
        switch (fi6 << ((362) & 63)) {
        case 42:
            v7 := GenS1{F0: 128, F1: false}
            _ = v7
        case -1:
            v8 := GenS2{F0: 2, F1: 2}
            _ = v8
        case 32767:
            var v9 float64 = 1.5
            _ = v9
        default:
            fmt.Println(v5, fi6)
        }
        _ = fi6
    }
    if ((42 < -92) && false) {
        var v10 int8 = 1
        v11 := []int{42, -804}
        _ = v10
        _ = v11
    }
    _ = v5
    return 2147483647
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]uint32)
    var v3 int = -1
    var v4 int = v3
    var v12 uint32 = 2147483647
    if (!true) {
        for fi13 := 0; fi13 < 17; fi13++ {
            var v14 uint8 = 0
            for fi15 := 0; fi15 < 2; fi15++ {
                v16 := GenS2{F0: (^(v12)), F1: 65535}
                if (!false) {
                    v4 = fi13
                    var v17 uint16 = 65535
                    switch (v3 + (v4 / (((fi15 & fi15)) | 1))) {
                    case 0:
                        if true {
                            if false {
                                var v18 uint = 18446744073709551615
                                fmt.Println(v14)
                                _ = v18
                            }
                            v19 := []int{}
                            v20 := GenS3{F0: v3}
                            _ = v19
                            _ = v20
                        }
                    case 2147483647:
                        v17 += v17
                    case 32767:
                        v21 := GenS2{F0: 1, F1: v17}
                        _ = v21
                    default:
                        var v22 int64 = 42
                        _ = v22
                    }
                    v23 := GenS2{F0: 309, F1: v17}
                    v24 := []int{fi13, (-(fi13))}
                    switch v4 {
                    case 32768:
                        if true {
                            var v25 int32 = -955
                            v26 := GenS4{F0: 109, F1: 811, F2: (^(v14))}
                            _ = v25
                            _ = v26
                        } else {
                            v27 := GenS3{F0: int(v12)}
                            _ = v27
                        }
                    default:
                        v28 := GenS1{F0: v14, F1: false}
                        _ = v28
                    }
                    if ((int32(fi15) > 339) || true) {
                        var v29 float32 = 0.1
                        var v30 uint8 = (v14 * uint8(v12))
                        _ = len(v24)
                        _ = v29
                        _ = v30
                    }
                    _ = v17
                    _ = v23
                    _ = v24
                } else {
                    switch "hello" {
                    case "":
                        v14 = (v14 | 255)
                    }
                }
                var v31 float32 = 0.0
                v3 = (v3 - fi15)
                _ = fi15
                _ = v16
                _ = v31
            }
            _ = fi13
            _ = v14
        }
        if false {
            var v32 int64 = 938
            _ = v32
        }
        v3 -= (-(v3))
    }
    v12 -= v12
    var v33 uint32 = v12
    v4 = 76
    switch (v4 | (-(v4))) {
    case 128:
        if true {
            switch 1 {
            case 42:
                switch "foo bar" {
                case "a":
                    var v34 float64 = (0.1 - 3.14159)
                    _ = v34
                case "\n\t":
                    v35 := GenS5{F0: 32767, F1: v3, F2: v33}
                    _ = v35
                case "":
                    v36 := GenS5{F0: 194, F1: 128, F2: v12}
                    _ = v36
                default:
                    switch "foo bar" {
                    case "foo bar":
                        v37 := []int{(-(v3)), 4294967295, (v3 - (v4 ^ (v3 - v4)))}
                        _ = v37
                    default:
                        var v38 uint8 = 205
                        _ = v38
                    }
                }
            default:
                for fi39 := 0; fi39 < 12; fi39++ {
                    for fi40 := 0; fi40 < 11; fi40++ {
                        if false {
                            v41 := GenS5{F0: 65535, F1: fi39, F2: 0}
                            _ = v41
                        }
                        switch ("hello" + "foo bar") {
                        case "foo bar":
                            if true {
                                var v42 string = "hello"
                                var v43 bool = (!(32767 >= 93))
                                v3 += (^(v4))
                                _ = v42
                                _ = v43
                            }
                        }
                        _ = fi40
                    }
                    fmt.Println(v33, fi39)
                    if (true && (37 != 128)) {
                        for fi44 := 0; fi44 < 18; fi44++ {
                            v12 = (v33 & (v12 >> ((42) & 63)))
                            _ = fi44
                        }
                    } else {
                        if (42 < 9223372036854775807) {
                            var v45 int8 = -114
                            _ = v45
                        } else {
                            for fi46 := 0; fi46 < 4; fi46++ {
                                v47 := GenS6{F0: 31}
                                _ = fi46
                                _ = v47
                            }
                            for fi48 := 0; fi48 < 8; fi48++ {
                                v3 = (fi48 & (v4 / ((int(v12)) | 1)))
                                fmt.Println(v4, fi39, v3)
                                v3 -= v4
                                _ = fi48
                            }
                        }
                        var v49 int = v4
                        fmt.Println(v3, fi39)
                        _ = v49
                    }
                    _ = fi39
                }
            }
            var v50 float32 = (3.14159 * 3.14159)
            _ = v50
        }
    default:
        if (1 <= 42) {
            switch ("hello" + ("foo bar" + ("hello" + "hello"))) {
            case "\n\t":
                if false {
                    v4 = v4
                } else {
                    v3 = (^(v3))
                    if false {
                        for fi51 := 0; fi51 < 17; fi51++ {
                            if false {
                                if (v3 >= 42) {
                                    fmt.Println(fi51, v12, fi51)
                                    for fi52 := 0; fi52 < 13; fi52++ {
                                        if false {
                                            v53 := v2[42]
                                            _ = v53
                                        }
                                        var v54 uint8 = 2
                                        _ = fi52
                                        _ = v54
                                    }
                                } else {
                                    for fi55 := 0; fi55 < 18; fi55++ {
                                        v56 := GenS4{F0: -1, F1: 42, F2: 42}
                                        _ = fi55
                                        _ = v56
                                    }
                                    v4 -= 127
                                }
                                v12 -= 259
                            }
                            for fi57 := 0; fi57 < 8; fi57++ {
                                v4 -= (v3 >> ((1) & 63))
                                _ = fi57
                            }
                            for fi58 := 0; fi58 < 3; fi58++ {
                                if false {
                                    v3 -= (^(fi51))
                                    fmt.Println(v33, v33)
                                }
                                switch "a" {
                                case "\n\t":
                                    v12 -= (v12 + v12)
                                case "":
                                    var v59 string = "hello"
                                    _ = v59
                                default:
                                    if (fi58 < v3) {
                                        if (4294967295 >= 0) {
                                            var v60 uint = 0
                                            var v61 float32 = 0.0
                                            _ = v60
                                            _ = v61
                                        } else {
                                            v12 += (^(v33))
                                            var v62 int32 = -208
                                            _ = v62
                                        }
                                    } else {
                                        v63 := []int{(-(v4))}
                                        _ = v63
                                    }
                                }
                                v64 := v2[-548]
                                _ = fi58
                                _ = v64
                            }
                            _ = fi51
                        }
                        for fi65 := 0; fi65 < 20; fi65++ {
                            v4 -= (^(v4))
                            var v66 bool = (!false)
                            _ = fi65
                            _ = v66
                        }
                        fmt.Println(v33)
                    } else {
                        switch -936 {
                        case 883:
                            fmt.Println(v3, v3, v4)
                        }
                        switch ("\n\t" + ("a" + "a")) {
                        case "foo bar":
                            fmt.Println(v3, v12)
                        case "a":
                            v33 = (^(v33))
                        default:
                            var v67 uint = 18446744073709551615
                            _ = v67
                        }
                    }
                }
            default:
                v3 -= v4
            }
        }
    }
    v68 := GenS7{F0: (v4 & (v3 / (((v4 | -9223372036854775808)) | 1))), F1: 42, F2: 127}
    v12 -= (^(v33))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v68
}
