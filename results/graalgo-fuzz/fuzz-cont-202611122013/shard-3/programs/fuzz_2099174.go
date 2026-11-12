package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 int32 }
type GenS2 struct { F0 uint8; F1 int16; F2 int16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int16; F1 uint32 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int32; F1 int32 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 uint32; F1 uint32 }
type GenS11 struct { F0 uint8; F1 int16; F2 int }

func main() {
    var v1 uint64 = 2
    v2 := []int{2147483647}
    var v3 int64 = 767
    v4 := make(map[uint32]bool)
    v5 := GenS1{F0: 127, F1: 38, F2: 110}
    if false {
        v6 := GenS2{F0: 185, F1: 0, F2: -1}
        _ = v6
    } else {
        for fi7 := 0; fi7 < 6; fi7++ {
            var v8 float32 = ((0.1 / 1.0) * 0.1)
            v1 -= 18446744073709551615
            _ = fi7
            _ = v8
        }
        v9 := GenS3{F0: -813}
        v3 = (v3 & v3)
        _ = v9
    }
    switch 2147483647 {
    case 42:
        var v10 string = "\n\t"
        _ = v10
    case -2147483648:
        v1 = (^(v1))
    case -799:
        v11 := GenS4{F0: 4294967295}
        _ = v11
    }
    var v12 uint64 = (v1 - v1)
    if (4294967295 <= -858) {
        for fi13 := 0; fi13 < 4; fi13++ {
            v4[2] = false
            if (2147483648 > 65535) {
                var v14 float64 = 1e10
                _ = v14
            }
            v15 := GenS5{F0: -32768, F1: 42}
            _ = fi13
            _ = v15
        }
    }
    v4[2147483648] = (125 > -128)
    fmt.Println(v12)
    switch 128 {
    case 0:
        var v16 uint16 = 32767
        _ = v16
    default:
        v17 := GenS6{F0: 0}
        _ = v17
    }
    if (0 < 0) {
        switch -436 {
        case -496:
            v18 := GenS6{F0: 255}
            _ = v18
        case 9223372036854775807:
            v3 = (v3 - (v3 - 0))
        case 1:
            var v19 uint32 = 2
            _ = v19
        default:
            var v20 float64 = 0.0
            _ = v20
        }
        v1 = 280
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    } else {
        v12 = 18446744073709551615
        v21 := GenS7{F0: 42, F1: 1}
        _ = v21
    }
    v22 := GenS1{F0: 215, F1: 0, F2: 2147483647}
    switch ("\n\t" + (("hello" + ((("\n\t" + "") + "\n\t") + ("hello" + "foo bar"))) + "hello")) {
    case "foo bar":
        if true {
            fmt.Println(v12, v1)
            for fi23 := 0; fi23 < 4; fi23++ {
                v12 += (v1 | v12)
                switch "\n\t" {
                case "\n\t":
                    v24 := GenS4{F0: 549}
                    _ = v24
                case "":
                    v25 := GenS8{F0: -32768}
                    _ = v25
                case "a":
                    for fi26 := 0; fi26 < 19; fi26++ {
                        var v27 uint = 255
                        _ = fi26
                        _ = v27
                    }
                }
                v28 := GenS9{F0: (("hello" + ("foo bar" + "a")) <= ("foo bar" + "a"))}
                _ = fi23
                _ = v28
            }
            v1 = v12
        } else {
            if (uint32(v12) >= 4294967295) {
                for fi29 := 0; fi29 < 13; fi29++ {
                    _ = len(v2)
                    var v30 uint16 = 765
                    v3 = 108
                    _ = fi29
                    _ = v30
                }
            } else {
                switch -9223372036854775808 {
                case -1:
                    for fi31 := 0; fi31 < 13; fi31++ {
                        for fi32 := 0; fi32 < 18; fi32++ {
                            switch "\n\t" {
                            case "foo bar":
                                switch "hello" {
                                case "foo bar":
                                    if false {
                                        var v33 int64 = (-(v3))
                                        v1 = (v1 & v1)
                                        _ = v33
                                    } else {
                                        v34 := GenS10{F0: 42, F1: 1}
                                        _ = len(v4)
                                        _ = v34
                                    }
                                case "hello":
                                    var v35 int32 = 236
                                    _ = v35
                                case "\n\t":
                                    delete(v4, 4294967295)
                                }
                            case "hello":
                                v12 -= 1
                            }
                            if true {
                                if ("a" <= ("\n\t" + "foo bar")) {
                                    fmt.Println(v3, fi31, fi32)
                                    v36 := v4[2]
                                    if false {
                                        var v37 float64 = 0.1
                                        _ = v37
                                    } else {
                                        fmt.Println(v36)
                                        v38 := GenS11{F0: 2, F1: 1, F2: fi31}
                                        fmt.Println(fi32, v36)
                                        v1 = 686
                                        var v39 float64 = 1.5
                                        _ = v38
                                        _ = v39
                                    }
                                    _ = v36
                                }
                                if (-128 == 42) {
                                    var v40 float64 = 1.5
                                    v4[916] = (!(!(false || (((128 >= (fi31 | (fi31 | (^(fi31))))) || false) && (42 != uint32(v12))))))
                                    if len(v2) > 0 {
                                        v2[0] = 0
                                    }
                                    _ = v40
                                }
                                fmt.Println(fi31, fi31)
                            }
                            _ = fi32
                        }
                        _ = fi31
                    }
                case 228:
                    var v41 int8 = -3
                    _ = v41
                case 401:
                    if (!(255 > 180)) {
                        v12 = (v12 / ((4294967295) | 1))
                    }
                default:
                    if true {
                        v1 += 898
                        for fi42 := 0; fi42 < 3; fi42++ {
                            if (1 == 42) {
                                var v43 int64 = (-(v3))
                                if true {
                                    v1 += (^(v1))
                                } else {
                                    v3 -= (v43 * (v3 % ((v3) | 1)))
                                }
                                _ = v43
                            } else {
                                var v44 uint32 = 4294967295
                                fmt.Println(v1)
                                fmt.Println(v1)
                                _ = v44
                            }
                            _ = fi42
                        }
                    }
                }
            }
            for fi45 := 0; fi45 < 19; fi45++ {
                fmt.Println(v12)
                _ = fi45
            }
        }
    }
    v46 := GenS2{F0: 128, F1: -1, F2: -1}
    var v47 string = "\n\t"
    _ = len(v4)
    for fi48 := 0; fi48 < 16; fi48++ {
        v49 := v4[600]
        switch "hello" {
        case "":
            fmt.Println(v1, v49)
        case "foo bar":
            v1 = (v1 & v1)
        case "\n\t":
            v12 -= (v12 >> ((2) & 63))
        }
        _ = fi48
        _ = v49
    }
    v12 -= v12
    fmt.Println(v47, v1, v12)
    var v50 int = 557
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v12
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v5
    _ = v50
}
