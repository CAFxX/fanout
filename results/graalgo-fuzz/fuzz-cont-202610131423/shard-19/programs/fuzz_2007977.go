package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 uint64 }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 bool; F1 int64; F2 uint64 }
type GenS5 struct { F0 int32; F1 int64 }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]uint32)
    var v3 uint32 = 4294967295
    var v4 int8 = 100
    v5 := []int{192, -729}
    v3 = v3
    var v6 uint16 = 42
    for fi7 := 0; fi7 < 14; fi7++ {
        v8 := GenS1{F0: 42, F1: v4, F2: 372}
        _ = fi7
        _ = v8
    }
    var v9 uint64 = 2
    for fi10 := 0; fi10 < 7; fi10++ {
        fmt.Println(v3, v3)
        _ = fi10
    }
    fmt.Println(v9)
    var v11 int32 = -304
    var v12 float64 = 1e10
    switch 230 {
    case -9223372036854775808:
        v13 := GenS1{F0: 9223372036854775807, F1: (v4 | (v4 ^ (v4 - (-(v4))))), F2: v9}
        _ = v13
    }
    v14 := GenS1{F0: 0, F1: (v4 >> ((255) & 63)), F2: 0}
    switch "\n\t" {
    case "a":
        if true {
            v15 := GenS2{F0: 99, F1: v9}
            if false {
                if (!false) {
                    if true {
                        v4 += (^(v4))
                        v9 = v9
                    } else {
                        var v16 string = ("hello" + ("hello" + ("\n\t" + ("hello" + "\n\t"))))
                        if ((v3 / (((^(v3))) | 1)) != v3) {
                            fmt.Println(v6)
                        }
                        v17 := v2[v3]
                        _ = v16
                        _ = v17
                    }
                } else {
                    var v18 uint8 = 2
                    for fi19 := 0; fi19 < 5; fi19++ {
                        v12 = ((-2.25 / v12) - (-2.25 - v12))
                        fmt.Println(v4, v9, v4)
                        _ = fi19
                    }
                    v12 = 0.1
                    _ = v18
                }
                switch -162 {
                case 9223372036854775807:
                    for fi20 := 0; fi20 < 17; fi20++ {
                        var v21 uint16 = 0
                        v12 = 0.0
                        v6 = v21
                        _ = fi20
                        _ = v21
                    }
                case 127:
                    var v22 uint64 = (v9 << ((255) & 63))
                    _ = v22
                case 32767:
                    switch "foo bar" {
                    case "\n\t":
                        var v23 int8 = (v4 ^ (v4 * v4))
                        _ = v23
                    case "foo bar":
                        var v24 int8 = (v4 & v4)
                        _ = v24
                    default:
                        switch 32767 {
                        case 489:
                            var v25 string = "foo bar"
                            _ = v25
                        case 520:
                            var v26 float32 = 1.5
                            _ = v26
                        default:
                            var v27 uint = 18446744073709551615
                            _ = v27
                        }
                    }
                }
                for fi28 := 0; fi28 < 13; fi28++ {
                    if (-40 >= 275) {
                        if ((^(v4)) >= (v4 % ((v4) | 1))) {
                            if false {
                                for fi29 := 0; fi29 < 6; fi29++ {
                                    fmt.Println(v9)
                                    var v30 int64 = -301
                                    _ = len(v2)
                                    _ = fi29
                                    _ = v30
                                }
                                v2[v3] = (^(v3))
                            }
                            for fi31 := 0; fi31 < 19; fi31++ {
                                if ((v9 / ((v9) | 1)) <= uint64(v6)) {
                                    switch 2147483647 {
                                    case 2147483647:
                                        var v32 int16 = 1
                                        _ = v32
                                    case 127:
                                        v6 = (v6 >> ((18446744073709551615) & 63))
                                    case 0:
                                        var v33 int16 = 1
                                        _ = v33
                                    }
                                    fmt.Println(v12, v11, v9)
                                }
                                v9 = (v9 + (v9 * 718))
                                _ = fi31
                            }
                        } else {
                            delete(v2, (^(v3)))
                            v6 = v6
                            v34 := GenS3{F0: (-(v4))}
                            _ = v34
                        }
                        v35 := GenS1{F0: -577, F1: (v4 - v4), F2: (^(v9))}
                        for fi36 := 0; fi36 < 12; fi36++ {
                            var v37 uint32 = (v3 * v3)
                            v38 := GenS1{F0: -15, F1: (^(v4)), F2: 1}
                            _ = fi36
                            _ = v37
                            _ = v38
                        }
                        _ = v35
                    } else {
                        var v39 float64 = v12
                        _ = v39
                    }
                    _ = fi28
                }
            } else {
                fmt.Println(v11)
                if false {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                } else {
                    v40 := GenS4{F0: false, F1: 0, F2: 0}
                    if false {
                        for fi41 := 0; fi41 < 8; fi41++ {
                            if (int8(fi41) < (v4 + (v4 % ((int8(v3)) | 1)))) {
                                fmt.Println(v12, fi41)
                                var v42 float64 = (((v12 + (v12 / 1.5)) / 0.1) - 0.1)
                                _ = v42
                            } else {
                                if true {
                                    switch "hello" {
                                    case "a":
                                        v9 += v9
                                    }
                                    for fi43 := 0; fi43 < 5; fi43++ {
                                        if len(v5) > 0 {
                                            v5[0] = fi43
                                        }
                                        if true {
                                            _ = len(v5)
                                        }
                                        switch "\n\t" {
                                        case "":
                                            var v44 int32 = (v11 ^ v11)
                                            _ = v44
                                        case "a":
                                            v45 := GenS1{F0: 747, F1: (v4 * v4), F2: 274}
                                            _ = v45
                                        case "\n\t":
                                            v46 := GenS4{F0: false, F1: 1, F2: v9}
                                            _ = v46
                                        }
                                        _ = fi43
                                    }
                                } else {
                                    v9 = 1
                                    if (!false) {
                                        v47 := GenS5{F0: -1, F1: 42}
                                        _ = v47
                                    } else {
                                        v3 = (v3 * (v3 | (v3 - (v3 >> ((42) & 63)))))
                                        v48 := GenS3{F0: 0}
                                        var v49 float64 = v12
                                        _ = v48
                                        _ = v49
                                    }
                                    v9 = 283
                                }
                            }
                            v4 = v4
                            _ = fi41
                        }
                    } else {
                        v12 = (1e10 + (v12 / (3.14159 - v12)))
                        if (-571 >= 1) {
                            v6 = (v6 * 32767)
                            var v50 uint64 = 18446744073709551615
                            _ = v50
                        } else {
                            if true {
                                var v51 int32 = (v11 << ((2) & 63))
                                var v52 int = -154
                                _ = v51
                                _ = v52
                            }
                        }
                        var v53 int = -170
                        _ = v53
                    }
                    for fi54 := 0; fi54 < 1; fi54++ {
                        v55 := GenS3{F0: v4}
                        v3 -= (v3 / (((v3 % (((^(v3))) | 1))) | 1))
                        v11 += 2147483647
                        _ = fi54
                        _ = v55
                    }
                    _ = v40
                }
            }
            _ = v15
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
