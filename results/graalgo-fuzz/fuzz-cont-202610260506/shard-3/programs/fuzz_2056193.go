package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint16 }

func main() {
    v1 := make(map[int]uint32)
    v2 := []uint32{1}
    v3 := make(map[int]int)
    switch 50 {
    case -198:
        v4 := GenS1{F0: 0, F1: 0}
        _ = v4
    }
    for fi5 := 0; fi5 < 19; fi5++ {
        var v6 string = "hello"
        _ = fi5
        _ = v6
    }
    switch "hello" {
    case "foo bar":
        v2 = append(v2, 2147483647)
    default:
        v3[546] = 1
    }
    if false {
        var v7 string = "foo bar"
        v8 := GenS2{F0: v7}
        var v9 uint64 = 603
        _ = v7
        _ = v8
        _ = v9
    } else {
        v10 := GenS1{F0: 127, F1: 1}
        delete(v1, 128)
        var v11 int16 = -879
        _ = v10
        _ = v11
    }
    if len(v2) > 0 {
        v2[0] = 94
    }
    var v12 uint32 = 1
    v12 = (^(v12))
    switch -401 {
    case 32767:
        if (false && false) {
            v13 := GenS3{F0: ""}
            switch -251 {
            case 2147483647:
                v12 = (v12 >> ((1) & 63))
            case 0:
                if false {
                    v14 := GenS4{F0: uint16(v12)}
                    _ = v14
                }
            case -2147483648:
                var v15 bool = true
                _ = v15
            }
            v12 = 149
            _ = v13
        } else {
            if false {
                var v16 int8 = -108
                if true {
                    for fi17 := 0; fi17 < 7; fi17++ {
                        v12 += v12
                        _ = fi17
                    }
                } else {
                    if (65535 > 44) {
                        v16 += (v16 * v16)
                        v16 -= (^(v16))
                        v16 = (v16 - v16)
                    } else {
                        v12 -= (v12 | (v12 % ((2) | 1)))
                    }
                    fmt.Println(v12)
                    var v18 uint32 = v12
                    _ = v18
                }
                v3[-302] = 9223372036854775807
                _ = v16
            } else {
                switch 608 {
                case -604:
                    var v19 int16 = -541
                    _ = v19
                case -9223372036854775808:
                    v12 += 1
                default:
                    var v20 float32 = 1.5
                    _ = v20
                }
            }
            v12 = (^(v12))
        }
    case 246:
        v21 := v1[906]
        _ = v21
    case 128:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    default:
        for fi22 := 0; fi22 < 4; fi22++ {
            if len(v2) > 0 {
                v2[0] = (v12 + v12)
            }
            v23 := GenS2{F0: "\n\t"}
            _ = fi22
            _ = v23
        }
    }
    v12 -= 556
    fmt.Println(v12)
    if false {
        for fi24 := 0; fi24 < 20; fi24++ {
            if (!true) {
                v12 = 587
            }
            var v25 uint32 = (v12 ^ v12)
            v25 = 691
            _ = fi24
            _ = v25
        }
        v12 = 942
        for fi26 := 0; fi26 < 16; fi26++ {
            fmt.Println(v12, fi26)
            for fi27 := 0; fi27 < 17; fi27++ {
                v28 := GenS2{F0: "hello"}
                var v29 string = "hello"
                _ = fi27
                _ = v28
                _ = v29
            }
            var v30 int32 = -750
            _ = fi26
            _ = v30
        }
    } else {
        switch 32768 {
        case -2147483648:
            v12 = v12
        case -543:
            v12 += v12
        case 941:
            for fi31 := 0; fi31 < 10; fi31++ {
                var v32 int64 = 42
                v32 += (v32 >> ((917) & 63))
                _ = fi31
                _ = v32
            }
        default:
            if (504 <= 155) {
                _ = len(v1)
            }
        }
    }
    var v33 string = "hello"
    v33 = (v33 + ("\n\t" + ("\n\t" + v33)))
    var v34 float32 = -2.25
    fmt.Println(v12)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v33
    _ = v34
}
