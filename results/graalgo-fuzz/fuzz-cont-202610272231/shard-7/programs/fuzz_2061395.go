package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 string; F1 int8 }

func main() {
    var v1 uint16 = 642
    var v2 bool = true
    var v3 uint8 = 117
    v4 := []int{610}
    var v5 int = -2147483648
    v3 += 2
    for fi6 := 0; fi6 < 10; fi6++ {
        v7 := GenS1{F0: (v1 >> ((241) & 63))}
        v2 = v2
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 4; fi8++ {
        fmt.Println(v1, v2, v3)
        if v2 {
            v9 := GenS1{F0: v1}
            v2 = v2
            switch 0 {
            case 4294967295:
                v10 := GenS1{F0: (v1 << ((18446744073709551615) & 63))}
                _ = v10
            case 207:
                v5 = int(v1)
            case 32767:
                v3 -= (^(v3))
            default:
                for fi11 := 0; fi11 < 8; fi11++ {
                    var v12 bool = true
                    v13 := GenS1{F0: v1}
                    _ = fi11
                    _ = v12
                    _ = v13
                }
            }
            _ = v9
        }
        _ = fi8
    }
    var v14 int64 = -945
    var v15 int = (v5 - (v5 | v5))
    var v16 int8 = 42
    switch "" {
    case "\n\t":
        if v2 {
            var v17 int64 = v14
            _ = v17
        } else {
            v18 := GenS1{F0: v1}
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 20; fi19++ {
        v14 -= (^(v14))
        for fi20 := 0; fi20 < 10; fi20++ {
            fmt.Println(v2, v2)
            _ = fi20
        }
        var v21 uint8 = (^(v3))
        _ = fi19
        _ = v21
    }
    if false {
        v5 -= (v5 % ((32767) | 1))
        if (!(true && v2)) {
            if (false && (((v16 / (((v16 ^ (-(v16)))) | 1)) > int8(v14)) || (!false))) {
                if len(v4) > 0 {
                    v4[0] = 2147483647
                }
            } else {
                for fi22 := 0; fi22 < 11; fi22++ {
                    var v23 uint64 = 953
                    _ = fi22
                    _ = v23
                }
                var v24 int8 = v16
                v25 := GenS1{F0: v1}
                _ = v24
                _ = v25
            }
        } else {
            if (float64(((3.14159 + 1.5) - (((((-2.25 * 1.5) / 3.14159) / 1e10) / 1.0) + ((-2.25 - (1.5 - -2.25)) / -2.25)))) == (1.5 + float64(0.1))) {
                for fi26 := 0; fi26 < 8; fi26++ {
                    v1 = 42
                    v27 := GenS2{F0: 0, F1: v16}
                    _ = fi26
                    _ = v27
                }
                fmt.Println(v5, v3, v14)
            } else {
                var v28 bool = v2
                _ = v28
            }
            var v29 string = "a"
            var v30 uint8 = v3
            _ = v29
            _ = v30
        }
        var v31 int8 = (v16 * (v16 - (v16 ^ (v16 - v16))))
        _ = v31
    }
    v32 := GenS3{F0: "\n\t", F1: v16}
    if v2 {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v2 = v2
    } else {
        v33 := GenS1{F0: v1}
        if (739 >= 42) {
            v1 = v1
            var v34 int8 = v16
            _ = v34
        } else {
            if len(v4) > 0 {
                v4[0] = (v5 >> ((0) & 63))
            }
            if len(v4) > 0 {
                v4[0] = (v15 + (-(v5)))
            }
        }
        _ = v33
    }
    var v35 uint8 = v3
    var v36 uint16 = 255
    _ = len(v4)
    v5 -= v15
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v32
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
