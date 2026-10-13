package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64; F1 int32; F2 string }
type GenS4 struct { F0 uint64; F1 uint8 }
func genFn0(p0 uint16, p1 int32, p2 int16) int8 {
    v7 := []int{}
    switch 44 {
    case 32767:
        if false {
            for fi8 := 0; fi8 < 20; fi8++ {
                var v9 int32 = 42
                _ = fi8
                _ = v9
            }
            v10 := GenS1{F0: (p1 | 42)}
            v7 = append(v7, 2147483647)
            _ = v10
        } else {
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            v11 := GenS1{F0: -2147483648}
            _ = v11
        }
    case 4294967295:
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
    }
    var v12 int64 = 0
    var v13 uint64 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    _ = v7
    return 1
}

func main() {
    var v1 float64 = (-2.25 + (1e10 + 0.1))
    var v2 uint8 = 255
    var v3 uint64 = 2
    var v4 int64 = 9223372036854775807
    v5 := make(map[uint32]int)
    v6 := []int{-810, int(v3), 9223372036854775807, 243}
    fmt.Println(v3, v3, v3)
    var v14 bool = (42 != 0)
    v15 := GenS1{F0: 0}
    if true {
        for fi16 := 0; fi16 < 17; fi16++ {
            if v14 {
                if v14 {
                    var v17 int16 = -869
                    v18 := GenS1{F0: 0}
                    for fi19 := 0; fi19 < 14; fi19++ {
                        v4 += v4
                        if false {
                            switch fi16 {
                            case 42:
                                v20 := GenS1{F0: 2147483647}
                                _ = v20
                            case -197:
                                v21 := GenS2{F0: 969}
                                _ = v21
                            case 32767:
                                switch 136 {
                                case 2147483647:
                                    v5[65535] = (^(fi16))
                                }
                            }
                            if v14 {
                                fmt.Println(v1, v2)
                            } else {
                                v14 = (255 > 924)
                                _ = len(v5)
                            }
                        }
                        _ = fi19
                    }
                    _ = v17
                    _ = v18
                } else {
                    v3 += (v3 & (v3 % ((0) | 1)))
                }
                for fi22 := 0; fi22 < 7; fi22++ {
                    v23 := v5[0]
                    v1 -= 1e10
                    _ = fi22
                    _ = v23
                }
            }
            switch "a" {
            case "\n\t":
                v24 := GenS1{F0: 1}
                _ = v24
            case "hello":
                var v25 int32 = 468
                _ = v25
            case "":
                v2 = 0
            default:
                v26 := GenS3{F0: (v3 / ((276) | 1)), F1: 2147483647, F2: ("foo bar" + "hello")}
                _ = v26
            }
            if v14 {
                v27 := GenS4{F0: 18446744073709551615, F1: 255}
                _ = v27
            }
            _ = fi16
        }
        switch 128 {
        case -736:
            switch -9223372036854775808 {
            case -9223372036854775808:
                for fi28 := 0; fi28 < 7; fi28++ {
                    for fi29 := 0; fi29 < 6; fi29++ {
                        var v30 float64 = v1
                        v5[240] = (fi29 + 4294967295)
                        _ = fi29
                        _ = v30
                    }
                    _ = fi28
                }
            case 128:
                if false {
                    v3 = (v3 % (((v3 << ((uint(v3)) & 63))) | 1))
                } else {
                    if len(v6) > 0 {
                        v6[0] = -9223372036854775808
                    }
                    switch ("\n\t" + ("hello" + ("hello" + "foo bar"))) {
                    case "\n\t":
                        v5[65535] = -910
                    case "":
                        if (127 > genFn0(2, 42, 1000)) {
                            v31 := GenS3{F0: (^(v3)), F1: -1, F2: "foo bar"}
                            v4 += v4
                            _ = v31
                        }
                    default:
                        for fi32 := 0; fi32 < 6; fi32++ {
                            v5[2] = (fi32 / ((fi32) | 1))
                            _ = fi32
                        }
                    }
                    v33 := GenS1{F0: 42}
                    _ = v33
                }
            case 4294967295:
                for fi34 := 0; fi34 < 3; fi34++ {
                    v4 -= v4
                    if (false || false) {
                        var v35 int8 = -1
                        v35 -= (v35 << ((65535) & 63))
                        _ = v35
                    } else {
                        var v36 bool = true
                        _ = v36
                    }
                    _ = fi34
                }
            default:
                switch -1 {
                case 67:
                    for fi37 := 0; fi37 < 3; fi37++ {
                        delete(v5, 42)
                        v14 = (1 >= 238)
                        for fi38 := 0; fi38 < 17; fi38++ {
                            var v39 uint32 = 42
                            _ = fi38
                            _ = v39
                        }
                        _ = fi37
                    }
                case 706:
                    v4 -= -934
                case -692:
                    for fi40 := 0; fi40 < 12; fi40++ {
                        v2 = 0
                        _ = fi40
                    }
                default:
                    fmt.Println(v1, v3, v4)
                }
            }
        case 0:
            if v14 {
                var v41 uint = 65535
                v2 = (v2 % ((v2) | 1))
                switch 986 {
                case 2147483647:
                    v1 += v1
                }
                _ = v41
            }
        }
        v2 = (^(v2))
    }
    v14 = v14
    var v42 uint = 2
    var v43 float32 = (1e10 - 0.1)
    if ((^(v3)) >= (v3 ^ (v3 >> (((v42 & (v42 >> ((v42) & 63)))) & 63)))) {
        switch "foo bar" {
        case "\n\t":
            v44 := GenS2{F0: 785}
            _ = v44
        case "foo bar":
            for fi45 := 0; fi45 < 10; fi45++ {
                var v46 uint64 = 4294967295
                v2 -= (v2 % (((v2 ^ (v2 % ((v2) | 1)))) | 1))
                v4 = v4
                _ = fi45
                _ = v46
            }
        default:
            v47 := GenS4{F0: 18446744073709551615, F1: (v2 | v2)}
            _ = v47
        }
        v48 := GenS4{F0: (v3 | (v3 >> (((v42 ^ v42)) & 63))), F1: v2}
        _ = v48
    }
    var v49 int32 = 954
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v49
    _ = v5
    _ = v6
}
