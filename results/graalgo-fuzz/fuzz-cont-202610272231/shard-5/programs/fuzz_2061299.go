package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint; F1 bool }
type GenS4 struct { F0 int64; F1 uint32 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int8; F1 int8 }
func genFn0() int64 {
    switch 9223372036854775807 {
    case 241:
        v7 := []int{1}
        _ = v7
    case 4294967295:
        var v8 bool = true
        _ = v8
    case 42:
        for fi9 := 0; fi9 < 2; fi9++ {
            for fi10 := 0; fi10 < 19; fi10++ {
                var v11 int = fi9
                var v12 int8 = -1
                _ = fi10
                _ = v11
                _ = v12
            }
            _ = fi9
        }
    default:
        var v13 int8 = 95
        _ = v13
    }
    v14 := []int{-425, -9223372036854775808, -9223372036854775808}
    if false {
        switch 0 {
        case 646:
            var v15 float32 = (float32(1e10) / ((3.14159 + (0.1 + 1e10)) + 1e10))
            _ = v15
        }
        if true {
            var v16 uint = 255
            v16 = (v16 * (^(v16)))
            v16 = 103
            _ = v16
        } else {
            for fi17 := 0; fi17 < 8; fi17++ {
                if true {
                    v18 := GenS1{F0: (255 > 255)}
                    _ = v18
                }
                _ = fi17
            }
        }
    }
    var v19 uint64 = 408
    _ = v14
    _ = v19
    return -324
}
func genFn1(p0 uint32, p1 uint8, p2 int) int16 {
    var v20 int16 = 32767
    v21 := GenS2{F0: -1}
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    _ = v21
    return 917
}

func main() {
    var v1 uint32 = 0
    var v2 int8 = 42
    var v3 uint = 18446744073709551615
    v4 := make(map[int]int)
    v5 := make(map[string]uint32)
    v6 := []uint32{(v1 >> (((v3 & 4294967295)) & 63)), (v1 ^ (v1 * v1))}
    var v22 float32 = -2.25
    v23 := GenS3{F0: 317, F1: (42 != 1)}
    v1 = v1
    if false {
        v24 := GenS3{F0: v3, F1: false}
        var v25 uint16 = 32767
        for fi26 := 0; fi26 < 20; fi26++ {
            var v27 int8 = (v2 >> ((v3) & 63))
            if true {
                v2 = v2
            } else {
                var v28 uint8 = 255
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = v28
            }
            v3 -= v3
            _ = fi26
            _ = v27
        }
        _ = v24
        _ = v25
    } else {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        switch 69 {
        case -107:
            fmt.Println(v2, v3)
        case 49:
            v29 := GenS3{F0: v3, F1: (!true)}
            _ = v29
        default:
            if len(v6) > 0 {
                v6[0] = v1
            }
        }
        switch "hello" {
        case "\n\t":
            v1 -= v1
        case "":
            v3 -= v3
        }
    }
    if true {
        var v30 float64 = (3.14159 + 3.14159)
        _ = v30
    }
    v31 := GenS1{F0: ("foo bar" != ("hello" + "\n\t"))}
    fmt.Println(v3, v3)
    v32 := v5[(("a" + (("\n\t" + "") + ("\n\t" + ("foo bar" + ("foo bar" + (("\n\t" + "") + "foo bar")))))) + "")]
    delete(v4, 128)
    if false {
        if false {
            if (!false) {
                var v33 float64 = 0.0
                var v34 int8 = v2
                v35 := GenS4{F0: -442, F1: v1}
                _ = v33
                _ = v34
                _ = v35
            }
            switch "\n\t" {
            case "\n\t":
                switch -364 {
                case 127:
                    v36 := GenS4{F0: 9223372036854775807, F1: (^(v1))}
                    _ = v36
                case 1:
                    v5["foo bar"] = (^(v1))
                case 32768:
                    switch 631 {
                    case 42:
                        switch "" {
                        case "foo bar":
                            var v37 int64 = 0
                            _ = v37
                        case "":
                            _ = len(v6)
                        }
                    case 9223372036854775807:
                        for fi38 := 0; fi38 < 20; fi38++ {
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            v1 += v1
                            v39 := GenS5{F0: 0}
                            _ = fi38
                            _ = v39
                        }
                    case 615:
                        v3 = 679
                    }
                }
            case "":
                fmt.Println(v22, v2, v32)
            default:
                v2 -= (v2 * v2)
            }
            for fi40 := 0; fi40 < 10; fi40++ {
                v4[(fi40 - (^(fi40)))] = 0
                var v41 uint8 = 0
                for fi42 := 0; fi42 < 13; fi42++ {
                    v3 = v3
                    var v43 int = 512
                    if false {
                        fmt.Println(v2)
                        if true {
                            fmt.Println(fi42)
                        } else {
                            fmt.Println(v41, fi40, v32)
                        }
                        var v44 uint64 = 643
                        _ = v44
                    } else {
                        v32 += 0
                    }
                    v32 = 631
                    v32 = v32
                    _ = fi42
                    _ = v43
                }
                _ = fi40
                _ = v41
            }
        }
        for fi45 := 0; fi45 < 11; fi45++ {
            var v46 int64 = 9223372036854775807
            if len(v6) > 0 {
                v6[0] = (^(v32))
            }
            for fi47 := 0; fi47 < 13; fi47++ {
                if true {
                    if false {
                        fmt.Println(fi47)
                    }
                    v1 += (v32 / (((v1 % ((v1) | 1))) | 1))
                    var v48 int32 = -585
                    fmt.Println(fi47, fi45)
                    _ = v48
                }
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = fi47
            }
            _ = fi45
            _ = v46
        }
    }
    v49 := GenS6{F0: (v2 / (((v2 ^ (v2 + v2))) | 1)), F1: (v2 ^ v2)}
    var v50 int = 409
    v1 -= (v1 >> (((v3 | (v3 - 66))) & 63))
    if (332 > -1) {
        fmt.Println(v3, v2)
        if len(v6) > 0 {
            v6[0] = 65535
        }
    } else {
        v51 := GenS6{F0: -1, F1: 1}
        _ = v51
    }
    var v52 uint = 822
    for fi53 := 0; fi53 < 13; fi53++ {
        var v54 int64 = 9223372036854775807
        _ = fi53
        _ = v54
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("v52:", v52)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v52
    _ = v6
}
