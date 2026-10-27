package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int8; F1 int64; F2 int8 }
type GenS5 struct { F0 bool; F1 uint16; F2 uint16 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint64 }

func main() {
    var v1 int = 32767
    var v2 int = (v1 >> ((47) & 63))
    v3 := make(map[uint32]int)
    v4 := make(map[int]int)
    if (42 <= 65535) {
        v5 := GenS1{F0: -47, F1: 4294967295}
        v2 -= 2147483647
        var v6 int64 = 292
        _ = v5
        _ = v6
    } else {
        if true {
            var v7 bool = (false || (4294967295 <= 255))
            _ = v7
        }
    }
    for fi8 := 0; fi8 < 20; fi8++ {
        for fi9 := 0; fi9 < 7; fi9++ {
            v10 := GenS2{F0: 65535}
            v2 += (v1 * fi8)
            _ = fi9
            _ = v10
        }
        switch (fi8 | -2147483648) {
        case -9223372036854775808:
            if false {
                var v11 float64 = float64(((0.0 * 0.1) - 1e10))
                if true {
                    v12 := GenS2{F0: 4294967295}
                    for fi13 := 0; fi13 < 6; fi13++ {
                        v14 := GenS3{F0: 128}
                        _ = fi13
                        _ = v14
                    }
                    _ = v12
                } else {
                    v4[4294967295] = v2
                    if true {
                        v2 -= 37
                        if true {
                            for fi15 := 0; fi15 < 6; fi15++ {
                                var v16 int64 = 9223372036854775807
                                var v17 uint64 = 2
                                var v18 int = 128
                                v3[2147483648] = v18
                                _ = fi15
                                _ = v16
                                _ = v17
                                _ = v18
                            }
                            v19 := GenS4{F0: 100, F1: -976, F2: 0}
                            _ = v19
                        }
                    }
                }
                if (true || (1 <= 32767)) {
                    v20 := GenS3{F0: 224}
                    _ = v20
                } else {
                    if false {
                        var v21 uint32 = 0
                        _ = v21
                    }
                    for fi22 := 0; fi22 < 15; fi22++ {
                        var v23 uint8 = 41
                        var v24 int16 = 1000
                        v2 = -1
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                }
                _ = v11
            }
        case 32767:
            v25 := []int{fi8, v1}
            _ = v25
        case 42:
            fmt.Println(fi8)
        default:
            v2 = (fi8 | -9223372036854775808)
        }
        _ = fi8
    }
    var v26 uint = 249
    if (true || (!(65535 > 2147483647))) {
        if false {
            var v27 float64 = -2.25
            _ = v27
        }
        switch v1 {
        case 491:
            v28 := GenS5{F0: (1.5 >= 1e10), F1: 42, F2: 1}
            _ = v28
        case -848:
            switch "a" {
            case "foo bar":
                v26 -= 65535
            case "a":
                switch v2 {
                case 4294967295:
                    switch (v1 | 1) {
                    case -1:
                        if true {
                            if false {
                                _ = len(v4)
                                v26 = (v26 / (((v26 + (v26 & v26))) | 1))
                            }
                        }
                    default:
                        if false {
                            v26 = 255
                            v2 += (v2 << (((v26 | v26)) & 63))
                        } else {
                            if true {
                                if true {
                                    v26 = v26
                                }
                                var v29 uint32 = 2
                                _ = v29
                            } else {
                                var v30 uint16 = 42
                                _ = v30
                            }
                            var v31 uint = v26
                            v32 := []int{32767, (v2 - v2)}
                            _ = v31
                            _ = v32
                        }
                    }
                case 2147483647:
                    v33 := GenS5{F0: true, F1: 216, F2: 32767}
                    _ = v33
                default:
                    v26 = (^(v26))
                }
            default:
                if true {
                    for fi34 := 0; fi34 < 19; fi34++ {
                        var v35 uint64 = 4294967295
                        fmt.Println(fi34, v2)
                        _ = fi34
                        _ = v35
                    }
                    fmt.Println(v1, v1, v1)
                    v36 := GenS5{F0: ((false || ((113 < 1) && ((!((true || false) && false)) || (!((false && (!(-464 != 0))) || (1 >= 2147483648)))))) || ((79 >= 693) || true)), F1: 255, F2: 1}
                    _ = v36
                } else {
                    if true {
                        v37 := GenS5{F0: true, F1: 2, F2: 833}
                        if false {
                            for fi38 := 0; fi38 < 16; fi38++ {
                                for fi39 := 0; fi39 < 1; fi39++ {
                                    fmt.Println(v1, fi39, v26)
                                    v40 := GenS5{F0: (true || false), F1: 1, F2: 42}
                                    v41 := GenS6{F0: -2147483648}
                                    _ = fi39
                                    _ = v40
                                    _ = v41
                                }
                                for fi42 := 0; fi42 < 3; fi42++ {
                                    if (42 != 0) {
                                        fmt.Println(fi42, fi42)
                                        var v43 float32 = 0.1
                                        _ = v43
                                    }
                                    v44 := []int{-9223372036854775808}
                                    if len(v44) > 0 {
                                        v44[0] = v1
                                    }
                                    _ = fi42
                                    _ = v44
                                }
                                _ = fi38
                            }
                            if false {
                                var v45 float64 = float64((3.14159 / 1.5))
                                var v46 float64 = v45
                                fmt.Println(v1)
                                _ = v45
                                _ = v46
                            } else {
                                var v47 int8 = 0
                                v1 -= v2
                                var v48 uint64 = 4294967295
                                _ = v47
                                _ = v48
                            }
                            for fi49 := 0; fi49 < 19; fi49++ {
                                v3[124] = fi49
                                _ = fi49
                            }
                        }
                        v50 := GenS7{F0: 65535}
                        _ = v37
                        _ = v50
                    }
                }
            }
        default:
            v2 = 9223372036854775807
        }
    } else {
        v2 = v2
    }
    for fi51 := 0; fi51 < 10; fi51++ {
        v52 := GenS8{F0: 1}
        var v53 int8 = 32
        _ = fi51
        _ = v52
        _ = v53
    }
    v54 := []int{(v1 & -8), v1, v1, (-(v1))}
    v3[2] = -2147483648
    if true {
        if false {
            fmt.Println(v2, v26)
            var v55 float64 = ((float64((float32(-2.25) + float32(1.5))) * 1.5) - 1e10)
            _ = v55
        }
    }
    _ = len(v54)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v54:", len(v54), v54)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v54
}
