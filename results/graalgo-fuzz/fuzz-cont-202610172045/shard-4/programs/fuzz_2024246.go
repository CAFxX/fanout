package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 int64; F1 int64; F2 int32 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int8 = 0
    v3 := []int{}
    var v4 string = ("foo bar" + "foo bar")
    var v5 int16 = 396
    var v6 int = 0
    var v7 uint64 = 18446744073709551615
    switch -336 {
    case 9223372036854775807:
        switch "foo bar" {
        case "\n\t":
            for fi8 := 0; fi8 < 4; fi8++ {
                v3 = append(v3, (v6 % ((v6) | 1)))
                var v9 int8 = (v2 + v2)
                _ = fi8
                _ = v9
            }
        }
    default:
        for fi10 := 0; fi10 < 18; fi10++ {
            _ = len(v1)
            var v11 int8 = (v2 | v2)
            v2 += v2
            _ = fi10
            _ = v11
        }
    }
    var v12 uint = 18446744073709551615
    v3 = append(v3, v6)
    for fi13 := 0; fi13 < 10; fi13++ {
        if (!true) {
            if true {
                var v14 uint64 = v7
                v15 := GenS1{F0: v6, F1: 851}
                if (1.5 == 1.5) {
                    v5 = 42
                }
                _ = v14
                _ = v15
            }
            v7 -= (^(v7))
        }
        v7 = v7
        var v16 uint64 = (v7 + (v7 * (^(v7))))
        _ = fi13
        _ = v16
    }
    for fi17 := 0; fi17 < 2; fi17++ {
        switch v4 {
        case "foo bar":
            _ = len(v1)
        default:
            var v18 int64 = 9223372036854775807
            _ = v18
        }
        if false {
            if true {
                v5 += (^(v5))
            }
            var v19 uint64 = 1
            if false {
                for fi20 := 0; fi20 < 4; fi20++ {
                    var v21 float32 = float32((-2.25 - float64(0.0)))
                    _ = fi20
                    _ = v21
                }
            } else {
                v5 -= (-(v5))
            }
            _ = v19
        } else {
            if len(v3) > 0 {
                v3[0] = fi17
            }
        }
        _ = fi17
    }
    for fi22 := 0; fi22 < 18; fi22++ {
        delete(v1, 198)
        v12 = v12
        _ = fi22
    }
    _ = len(v1)
    if true {
        if true {
            var v23 uint = v12
            for fi24 := 0; fi24 < 4; fi24++ {
                v12 -= 563
                _ = fi24
            }
            v25 := GenS2{F0: 1, F1: (^(v23))}
            _ = v23
            _ = v25
        } else {
            if ((v2 != v2) && true) {
                v12 = v12
            } else {
                v3 = append(v3, (v6 / ((42) | 1)))
            }
            if (4 != 255) {
                for fi26 := 0; fi26 < 17; fi26++ {
                    switch (v6 - (^(v6))) {
                    case 127:
                        var v27 uint = 42
                        _ = v27
                    default:
                        var v28 int8 = (v2 ^ v2)
                        _ = v28
                    }
                    _ = fi26
                }
            } else {
                for fi29 := 0; fi29 < 16; fi29++ {
                    var v30 string = "a"
                    _ = fi29
                    _ = v30
                }
            }
        }
        if true {
            if true {
                for fi31 := 0; fi31 < 8; fi31++ {
                    fmt.Println(v4, fi31)
                    fmt.Println(v7)
                    for fi32 := 0; fi32 < 16; fi32++ {
                        v33 := GenS1{F0: -885, F1: 1}
                        fmt.Println(v4, fi32, v12)
                        _ = fi32
                        _ = v33
                    }
                    _ = fi31
                }
                if len(v3) > 0 {
                    v3[0] = v6
                }
                switch 32768 {
                case -9223372036854775808:
                    var v34 bool = false
                    _ = v34
                case -214:
                    v12 = v12
                case 4294967295:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                default:
                    fmt.Println(v7, v6, v6)
                }
            } else {
                fmt.Println(v2, v2)
            }
            var v35 int = int(v7)
            v12 = v12
            _ = v35
        }
        if (500 >= (v7 / ((18446744073709551615) | 1))) {
            if true {
                if (((("hello" + v4) + v4) + v4) >= v4) {
                    _ = len(v1)
                    delete(v1, 1)
                }
            }
            switch 32767 {
            case 1:
                switch "a" {
                case "foo bar":
                    switch (v6 >> (((v12 ^ (^(v12)))) & 63)) {
                    case -9223372036854775808:
                        switch "" {
                        case "hello":
                            var v36 uint16 = 182
                            _ = v36
                        case "\n\t":
                            v6 += v6
                        default:
                            fmt.Println(v4, v5)
                        }
                    case -1:
                        v37 := GenS3{F0: -9223372036854775808, F1: -9223372036854775808, F2: 628}
                        _ = v37
                    }
                case "\n\t":
                    _ = len(v1)
                }
            case 4294967295:
                v7 -= v7
            case 9223372036854775807:
                if (int(v5) != (^(v6))) {
                    _ = len(v3)
                }
            default:
                _ = len(v3)
            }
        }
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    v38 := v1[385]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
