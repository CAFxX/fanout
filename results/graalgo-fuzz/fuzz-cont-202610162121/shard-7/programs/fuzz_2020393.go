package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 uint16 }
type GenS4 struct { F0 int64; F1 int16; F2 uint16 }
type GenS5 struct { F0 int32; F1 int }
func genFn0(p0 uint8, p1 uint16) int {
    v7 := GenS1{F0: (^(p0)), F1: 1}
    var v8 uint8 = p0
    fmt.Println(p1, v8, p1)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 4294967295
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 uint64 = (^(v1))
    v3 := []uint32{780, 885}
    var v4 int64 = -9223372036854775808
    var v5 int16 = 1000
    v6 := []int{-2147483648, -487, -1}
    v4 -= int64(v2)
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v9 int64 = 9223372036854775807
    v10 := GenS1{F0: 1, F1: v5}
    var v11 float32 = -2.25
    var v12 float64 = -2.25
    fmt.Println(v12)
    v6 = append(v6, 714)
    var v13 uint = 18446744073709551615
    var v14 uint16 = 65535
    v2 = (v1 | (v1 + 60))
    switch 9223372036854775807 {
    case -9223372036854775808:
        var v15 uint8 = 2
        _ = v15
    case 1:
        var v16 bool = (!((^(v14)) != v14))
        _ = v16
    default:
        if len(v3) > 0 {
            v3[0] = uint32(v13)
        }
    }
    if true {
        if true {
            if true {
                var v17 int16 = (v5 + 1000)
                if len(v6) > 0 {
                    v6[0] = 795
                }
                switch 209 {
                case -129:
                    fmt.Println(v9)
                case 931:
                    for fi18 := 0; fi18 < 13; fi18++ {
                        v3 = append(v3, 2147483647)
                        switch 32767 {
                        case -2147483648:
                            v19 := GenS2{F0: 52}
                            _ = v19
                        default:
                            var v20 uint = 65535
                            _ = v20
                        }
                        _ = fi18
                    }
                default:
                    if len(v3) > 0 {
                        v3[0] = 2
                    }
                }
                _ = v17
            } else {
                v5 += v5
                for fi21 := 0; fi21 < 3; fi21++ {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    switch (^(fi21)) {
                    case -9223372036854775808:
                        if false {
                            v12 += (-2.25 - v12)
                            for fi22 := 0; fi22 < 18; fi22++ {
                                v23 := GenS3{F0: -1, F1: 478}
                                _ = fi22
                                _ = v23
                            }
                            v4 -= (v4 - v9)
                        } else {
                            var v24 uint = 0
                            fmt.Println(v14, v24, v14)
                            _ = v24
                        }
                    case -528:
                        switch (fi21 >> ((v13) & 63)) {
                        case 127:
                            if (-9223372036854775808 >= fi21) {
                                var v25 float64 = v12
                                _ = v25
                            } else {
                                v26 := GenS2{F0: 127}
                                _ = v26
                            }
                        case 32768:
                            switch ("foo bar" + "hello") {
                            case "foo bar":
                                v27 := GenS3{F0: -195, F1: v14}
                                _ = v27
                            case "hello":
                                switch "foo bar" {
                                case "hello":
                                    v6 = append(v6, fi21)
                                }
                            default:
                                for fi28 := 0; fi28 < 7; fi28++ {
                                    if (!true) {
                                        fmt.Println(v12, fi21)
                                    } else {
                                        if ((false && false) || true) {
                                            var v29 int = fi21
                                            switch int(v14) {
                                            case 42:
                                                var v30 float64 = float64((v11 + (3.14159 * v11)))
                                                _ = v30
                                            case 9223372036854775807:
                                                v3 = append(v3, 2147483647)
                                            case -1:
                                                var v31 uint64 = v2
                                                _ = v31
                                            }
                                            _ = v29
                                        } else {
                                            var v32 int64 = (v9 << (((^(v13))) & 63))
                                            _ = v32
                                        }
                                        fmt.Println(fi21)
                                        switch ("\n\t" + "a") {
                                        case "":
                                            v12 += 0.0
                                        case "a":
                                            v9 += (v4 * -748)
                                        case "foo bar":
                                            v4 = 42
                                        }
                                    }
                                    _ = fi28
                                }
                            }
                        default:
                            if true {
                                v2 = v2
                                if false {
                                    v33 := GenS4{F0: v4, F1: v5, F2: (v14 ^ (v14 ^ 255))}
                                    v34 := GenS2{F0: 117}
                                    var v35 bool = (42 != 1)
                                    _ = v33
                                    _ = v34
                                    _ = v35
                                }
                                switch 531 {
                                case 582:
                                    if false {
                                        v6 = append(v6, fi21)
                                        v2 += uint64(v14)
                                        v1 = (v2 >> (((v13 % ((v13) | 1))) & 63))
                                    }
                                case 326:
                                    if (2147483647 >= 800) {
                                        for fi36 := 0; fi36 < 17; fi36++ {
                                            fmt.Println(v12)
                                            _ = fi36
                                        }
                                    } else {
                                        var v37 int = fi21
                                        _ = v37
                                    }
                                default:
                                    v38 := GenS1{F0: 255, F1: -623}
                                    _ = v38
                                }
                            } else {
                                v13 = 479
                                fmt.Println(v2, v14, v13)
                            }
                        }
                    case 407:
                        v39 := GenS3{F0: -1, F1: 279}
                        _ = v39
                    default:
                        v40 := GenS4{F0: v4, F1: int16(v14), F2: (v14 / (((v14 ^ (v14 ^ (v14 | 2)))) | 1))}
                        _ = v40
                    }
                    var v41 uint8 = 255
                    _ = fi21
                    _ = v41
                }
            }
            v4 = 42
        }
    }
    for fi42 := 0; fi42 < 9; fi42++ {
        for fi43 := 0; fi43 < 7; fi43++ {
            fmt.Println(v4)
            v5 = (v5 << ((v13) & 63))
            if len(v3) > 0 {
                v3[0] = 31
            }
            _ = fi43
        }
        if (!false) {
            if (!true) {
                v44 := GenS3{F0: -1, F1: (v14 ^ (v14 - 1))}
                _ = v44
            }
        }
        _ = fi42
    }
    v45 := GenS5{F0: 1, F1: 479}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v6
    _ = v9
}
