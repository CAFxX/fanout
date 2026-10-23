package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 uint8 }
type GenS2 struct { F0 uint16; F1 uint8; F2 string }

func main() {
    v1 := []string{"a", ("" + "\n\t"), "a"}
    var v2 uint = 2
    v3 := make(map[uint32]bool)
    v4 := []uint32{1, 301}
    var v5 uint16 = 65535
    if true {
        var v6 float32 = 1e10
        var v7 float32 = v6
        _ = v6
        _ = v7
    }
    if false {
        v2 = v2
    }
    switch "\n\t" {
    case "foo bar":
        if len(v1) > 0 {
            v1[0] = "\n\t"
        }
    default:
        if false {
            v8 := GenS1{F0: false, F1: 93, F2: 255}
            v2 = v2
            _ = v8
        } else {
            if (65535 <= 760) {
                v9 := GenS1{F0: (9223372036854775807 > 603), F1: 0, F2: 0}
                _ = v9
            } else {
                v5 = (v5 + (v5 % (((v5 ^ (v5 - (v5 * v5)))) | 1)))
            }
            var v10 uint8 = 42
            var v11 uint = v2
            _ = v10
            _ = v11
        }
    }
    v12 := GenS1{F0: (42 >= 1), F1: 75, F2: 0}
    if (1e10 >= float32(-2.25)) {
        v5 = 2
        if true {
            if (42 == 18446744073709551615) {
                var v13 uint8 = uint8(v2)
                _ = v13
            } else {
                v14 := GenS2{F0: v5, F1: 2, F2: "\n\t"}
                v5 += (v5 & v5)
                _ = v14
            }
            for fi15 := 0; fi15 < 4; fi15++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v16 := v3[1]
                _ = fi15
                _ = v16
            }
        } else {
            var v17 string = ("" + "hello")
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = v17
        }
    } else {
        _ = len(v3)
    }
    v3[207] = false
    for fi18 := 0; fi18 < 11; fi18++ {
        v2 += (v2 | (v2 / (((v2 & 376)) | 1)))
        _ = fi18
    }
    v3[1] = false
    fmt.Println(v5, v2)
    v19 := v3[0]
    switch "foo bar" {
    case "\n\t":
        if len(v4) > 0 {
            v4[0] = 174
        }
    case "foo bar":
        for fi20 := 0; fi20 < 6; fi20++ {
            for fi21 := 0; fi21 < 14; fi21++ {
                var v22 int = fi20
                v4 = append(v4, 412)
                var v23 int8 = -128
                _ = fi21
                _ = v22
                _ = v23
            }
            switch "a" {
            case "\n\t":
                var v24 uint16 = (v5 / (((v5 | (v5 + (v5 >> ((v2) & 63))))) | 1))
                _ = v24
            }
            _ = fi20
        }
    case "hello":
        switch "hello" {
        case "\n\t":
            var v25 int32 = 686
            _ = v25
        }
    default:
        for fi26 := 0; fi26 < 17; fi26++ {
            for fi27 := 0; fi27 < 19; fi27++ {
                switch "a" {
                case "":
                    switch "\n\t" {
                    case "foo bar":
                        fmt.Println(v2)
                    case "hello":
                        v2 = 318
                    case "\n\t":
                        v2 = 255
                    }
                case "hello":
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                default:
                    v19 = (!(int16(v5) != -828))
                }
                v28 := GenS1{F0: v19, F1: 97, F2: 127}
                _ = fi27
                _ = v28
            }
            fmt.Println(v5, v2, fi26)
            for fi29 := 0; fi29 < 17; fi29++ {
                v30 := GenS1{F0: (!true), F1: 79, F2: 155}
                _ = fi29
                _ = v30
            }
            _ = fi26
        }
    }
    var v31 uint16 = (^(v5))
    var v32 uint = (v2 << (((v2 | v2)) & 63))
    v33 := GenS1{F0: (v19 || v19), F1: -128, F2: 50}
    v31 -= v31
    fmt.Println(v32, v32, v31)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
