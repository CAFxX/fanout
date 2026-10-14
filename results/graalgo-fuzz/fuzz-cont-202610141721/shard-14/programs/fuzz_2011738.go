package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 uint16 }
type GenS2 struct { F0 int16; F1 int64; F2 uint8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint8 }
func genFn0(p0 uint, p1 uint32, p2 uint16) uint64 {
    for fi4 := 0; fi4 < 1; fi4++ {
        if false {
            if true {
                p1 += p1
                v5 := GenS1{F0: int8(fi4), F1: (p0 / ((p0) | 1)), F2: (p2 / ((65535) | 1))}
                var v6 int64 = -225
                _ = v5
                _ = v6
            }
            v7 := GenS1{F0: -1, F1: (p0 << (((p0 * 255)) & 63)), F2: (p2 + (p2 >> (((p0 ^ p0)) & 63)))}
            _ = v7
        }
        if (2147483647 > 1) {
            v8 := []int{}
            _ = v8
        } else {
            p0 -= (p0 % ((p0) | 1))
            var v9 int32 = 213
            for fi10 := 0; fi10 < 7; fi10++ {
                p0 += p0
                _ = fi10
            }
            _ = v9
        }
        _ = fi4
    }
    v11 := []int{-518, 42, 17}
    if (1 > 4294967295) {
        v11 = append(v11, -31)
        var v12 uint = p0
        _ = v12
    }
    p0 = 736
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 2
}

func main() {
    var v1 uint8 = 0
    v2 := make(map[uint32]uint32)
    var v3 uint64 = 2
    var v13 string = ("" + ("a" + "foo bar"))
    if ((true && (uint(v1) <= 0)) && (false || (!true))) {
        v14 := GenS1{F0: -128, F1: 0, F2: 723}
        v15 := GenS2{F0: 42, F1: -9223372036854775808, F2: v1}
        _ = v14
        _ = v15
    }
    fmt.Println(v1, v3, v13)
    v1 = v1
    switch "" {
    case "":
        if false {
            for fi16 := 0; fi16 < 3; fi16++ {
                for fi17 := 0; fi17 < 8; fi17++ {
                    if true {
                        v3 = 36
                        fmt.Println(v3, fi16, v1)
                        v1 -= 14
                    }
                    fmt.Println(v13, fi16)
                    var v18 uint64 = v3
                    _ = fi17
                    _ = v18
                }
                var v19 string = v13
                v3 += genFn0(65535, 1, 32767)
                _ = fi16
                _ = v19
            }
            var v20 uint32 = 2
            _ = v20
        } else {
            var v21 uint32 = 2
            _ = v21
        }
    }
    v22 := GenS3{F0: 100}
    v23 := []int{-372, 128, 127, -984}
    if false {
        v3 = (^(v3))
        var v24 string = v13
        v25 := GenS3{F0: 100}
        _ = v24
        _ = v25
    } else {
        v1 -= v1
        v3 -= genFn0(2, 2, 65535)
        var v26 uint = 65535
        _ = v26
    }
    switch 127 {
    case 59:
        for fi27 := 0; fi27 < 15; fi27++ {
            for fi28 := 0; fi28 < 20; fi28++ {
                switch v13 {
                case "hello":
                    var v29 float32 = 0.0
                    _ = v29
                case "foo bar":
                    v13 = (v13 + v13)
                case "a":
                    v23 = append(v23, fi27)
                default:
                    var v30 int = fi28
                    _ = v30
                }
                v31 := GenS4{F0: uint8(v3)}
                _ = fi28
                _ = v31
            }
            _ = fi27
        }
    case 847:
        switch "a" {
        case "":
            switch 404 {
            case 42:
                if true {
                    var v32 float64 = float64(1e10)
                    _ = v32
                }
            case 0:
                for fi33 := 0; fi33 < 17; fi33++ {
                    v1 = v1
                    _ = fi33
                }
            case -448:
                var v34 int8 = 42
                _ = v34
            default:
                var v35 uint8 = v1
                _ = v35
            }
        case "hello":
            fmt.Println(v13)
        }
    case 127:
        for fi36 := 0; fi36 < 17; fi36++ {
            var v37 int64 = -440
            for fi38 := 0; fi38 < 9; fi38++ {
                var v39 float32 = -2.25
                _ = fi38
                _ = v39
            }
            var v40 uint = 518
            _ = fi36
            _ = v37
            _ = v40
        }
    default:
        if false {
            switch -9223372036854775808 {
            case 127:
                v1 += 128
            case -368:
                v13 = "foo bar"
            }
            v13 = v13
        }
    }
    v41 := GenS2{F0: 458, F1: 0, F2: (v1 >> ((4294967295) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v41
}
