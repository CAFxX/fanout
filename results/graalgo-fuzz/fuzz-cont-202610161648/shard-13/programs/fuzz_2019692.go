package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 int; F1 int32; F2 uint }
type GenS3 struct { F0 bool; F1 int32; F2 uint64 }

func main() {
    v1 := make(map[uint32]int)
    v2 := []int{32768, 4294967295}
    var v3 int = -158
    var v4 int8 = 100
    var v5 int32 = -2147483648
    v6 := GenS1{F0: 2147483647, F1: 128}
    v4 = 42
    switch 127 {
    case -982:
        if false {
            var v7 int32 = (v5 | v5)
            fmt.Println(v3, v5, v3)
            _ = len(v2)
            _ = v7
        }
    default:
        if ((v5 >> ((42) & 63)) > 1) {
            for fi8 := 0; fi8 < 11; fi8++ {
                v4 -= (v4 | v4)
                _ = fi8
            }
            v4 += v4
            v2 = append(v2, v3)
        } else {
            v4 += (v4 << ((967) & 63))
        }
    }
    switch v3 {
    case 0:
        var v9 int8 = -128
        _ = v9
    case 42:
        var v10 uint64 = 42
        _ = v10
    case 4294967295:
        if false {
            if (65535 <= 529) {
                for fi11 := 0; fi11 < 14; fi11++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    for fi12 := 0; fi12 < 1; fi12++ {
                        fmt.Println(v5, v5)
                        v5 = (v5 & 0)
                        v4 += v4
                        _ = fi12
                    }
                    _ = fi11
                }
                switch v3 {
                case 9223372036854775807:
                    switch "\n\t" {
                    case "":
                        v5 = v5
                    case "a":
                        var v13 int16 = 1
                        _ = v13
                    default:
                        v3 -= 128
                    }
                default:
                    var v14 string = "hello"
                    _ = v14
                }
            } else {
                _ = len(v2)
                v4 += v4
                for fi15 := 0; fi15 < 14; fi15++ {
                    var v16 uint8 = 234
                    v17 := GenS1{F0: 2147483647, F1: (^(v16))}
                    v4 = v4
                    _ = fi15
                    _ = v16
                    _ = v17
                }
            }
            for fi18 := 0; fi18 < 8; fi18++ {
                var v19 string = ("hello" + ("a" + "hello"))
                for fi20 := 0; fi20 < 8; fi20++ {
                    v21 := GenS2{F0: (fi18 | v3), F1: -2147483648, F2: 42}
                    var v22 float32 = 1.5
                    _ = fi20
                    _ = v21
                    _ = v22
                }
                var v23 bool = ((-(v3)) != (v3 * v3))
                _ = fi18
                _ = v19
                _ = v23
            }
        }
    default:
        v24 := GenS1{F0: 49, F1: 42}
        _ = v24
    }
    switch (v3 * 0) {
    case 32768:
        switch ("" + "foo bar") {
        case "foo bar":
            var v25 bool = ((false || false) && (!true))
            _ = v25
        case "hello":
            v3 = v3
        case "\n\t":
            v4 -= (v4 >> ((470) & 63))
        }
    case -347:
        var v26 float32 = 0.1
        _ = v26
    case 32767:
        var v27 int32 = v5
        _ = v27
    default:
        var v28 bool = (true && (true && ((v5 / (((-(v5))) | 1)) != v5)))
        _ = v28
    }
    var v29 float64 = (3.14159 / 1.0)
    switch v3 {
    case 4294967295:
        v3 -= (^(v3))
    case 128:
        delete(v1, 235)
    case 42:
        v30 := GenS1{F0: 2, F1: 128}
        _ = v30
    }
    for fi31 := 0; fi31 < 18; fi31++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v3 -= (fi31 << ((18446744073709551615) & 63))
        _ = fi31
    }
    var v32 int8 = v4
    var v33 float32 = (1e10 + 0.1)
    for fi34 := 0; fi34 < 9; fi34++ {
        if (v5 >= v5) {
            switch (("foo bar" + "hello") + ("" + "foo bar")) {
            case "":
                v35 := GenS3{F0: false, F1: 1, F2: 2}
                _ = v35
            default:
                v3 += fi34
            }
            v4 = (v32 / (((v4 | (v4 ^ v4))) | 1))
            v36 := v1[4294967295]
            _ = v36
        }
        _ = fi34
    }
    if (v3 > v3) {
        fmt.Println(v32)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
