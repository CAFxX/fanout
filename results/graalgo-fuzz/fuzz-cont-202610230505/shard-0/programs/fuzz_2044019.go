package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int64; F1 int }
type GenS4 struct { F0 bool; F1 uint8 }
type GenS5 struct { F0 uint16 }
func genFn0(p0 uint, p1 int16) string {
    p1 = -118
    var v6 float32 = 1.5
    v7 := GenS1{F0: -9223372036854775808, F1: 32767, F2: -2147483648}
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return ("\n\t" + "a")
}

func main() {
    v1 := []int{-423, 42, 9223372036854775807}
    v2 := make(map[uint32]string)
    v3 := []uint32{65535}
    v4 := make(map[uint32]int)
    v5 := make(map[uint32]int)
    v8 := v2[2]
    var v9 uint32 = 0
    fmt.Println(v9, v8)
    if (!true) {
        fmt.Println(v9, v9)
        var v10 float64 = 0.0
        _ = v10
    } else {
        var v11 uint32 = (v9 / ((2147483648) | 1))
        for fi12 := 0; fi12 < 15; fi12++ {
            v8 = v8
            var v13 int32 = -590
            _ = fi12
            _ = v13
        }
        _ = v11
    }
    switch -631 {
    case 4294967295:
        var v14 uint = 78
        _ = v14
    case 2147483647:
        fmt.Println(v8, v8)
    }
    for fi15 := 0; fi15 < 12; fi15++ {
        if (!true) {
            fmt.Println(v9)
        }
        _ = fi15
    }
    if false {
        v16 := GenS1{F0: 9223372036854775807, F1: 32767, F2: 285}
        for fi17 := 0; fi17 < 8; fi17++ {
            var v18 uint8 = 1
            _ = fi17
            _ = v18
        }
        _ = v16
    }
    v19 := v5[(^(v9))]
    for fi20 := 0; fi20 < 9; fi20++ {
        var v21 float32 = -2.25
        _ = fi20
        _ = v21
    }
    if false {
        v3 = append(v3, 1)
    }
    for fi22 := 0; fi22 < 3; fi22++ {
        if (((!(1 > (fi22 + fi22))) || (1 == 1)) && (4294967295 == 42)) {
            var v23 int8 = 125
            _ = v23
        } else {
            fmt.Println(v19, v9)
            if true {
                v8 = "a"
            }
            if true {
                var v24 uint = 1
                v19 = (fi22 * (^(v19)))
                _ = v24
            } else {
                for fi25 := 0; fi25 < 13; fi25++ {
                    var v26 int32 = -1
                    v27 := GenS2{F0: 4294967295}
                    if len(v3) > 0 {
                        v3[0] = (v9 << ((255) & 63))
                    }
                    _ = fi25
                    _ = v26
                    _ = v27
                }
            }
        }
        _ = fi22
    }
    v19 = -445
    v19 += 1
    var v28 uint8 = 2
    _ = len(v2)
    switch v19 {
    case 9223372036854775807:
        if true {
            v29 := GenS1{F0: -9223372036854775808, F1: 32767, F2: 0}
            switch "a" {
            case "a":
                v30 := GenS3{F0: 0, F1: (-(v19))}
                _ = v30
            default:
                v3 = append(v3, (v9 & (^(v9))))
            }
            switch (-(v19)) {
            case -68:
                switch v8 {
                case "a":
                    var v31 uint = 18446744073709551615
                    _ = v31
                case "\n\t":
                    if (true || false) {
                        var v32 float32 = 0.0
                        for fi33 := 0; fi33 < 13; fi33++ {
                            v34 := GenS4{F0: false, F1: (^(v28))}
                            for fi35 := 0; fi35 < 3; fi35++ {
                                if ((v8 + v8) > (("" + v8) + v8)) {
                                    switch fi33 {
                                    case -1:
                                        switch fi35 {
                                        case -384:
                                            v8 = v8
                                        default:
                                            v36 := GenS3{F0: 561, F1: (v19 + (-(fi33)))}
                                            _ = v36
                                        }
                                    case 128:
                                        v9 += 0
                                    }
                                    var v37 int8 = 42
                                    _ = v37
                                }
                                _ = fi35
                            }
                            _ = fi33
                            _ = v34
                        }
                        _ = v32
                    } else {
                        var v38 int64 = -1
                        v28 -= v28
                        v9 = 2147483648
                        _ = v38
                    }
                default:
                    v28 -= (v28 - 1)
                }
            case -362:
                v28 += (v28 + (v28 / (((v28 ^ (^(v28)))) | 1)))
            case -9223372036854775808:
                for fi39 := 0; fi39 < 9; fi39++ {
                    v9 = (v9 ^ 42)
                    switch v8 {
                    case "":
                        v40 := GenS2{F0: (^(fi39))}
                        _ = v40
                    }
                    v41 := GenS3{F0: -1, F1: (fi39 % ((4294967295) | 1))}
                    _ = fi39
                    _ = v41
                }
            default:
                var v42 int16 = 1000
                _ = v42
            }
            _ = v29
        } else {
            v43 := GenS5{F0: 713}
            var v44 float64 = 1e10
            _ = v43
            _ = v44
        }
    }
    v9 -= (v9 % ((v9) | 1))
    v9 = v9
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
