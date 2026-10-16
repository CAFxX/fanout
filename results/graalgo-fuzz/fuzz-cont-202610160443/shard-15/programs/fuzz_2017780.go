package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 uint16 }

func main() {
    v1 := []uint32{2147483648, 2147483647, 1, 2147483648}
    var v2 uint = 2
    var v3 uint8 = 42
    var v4 int32 = 1
    v5 := GenS1{F0: (v3 & (^(v3))), F1: (false && (4294967295 > 42)), F2: 2}
    switch -1 {
    case -1:
        v6 := GenS1{F0: (v3 & (v3 << ((v2) & 63))), F1: (true || true), F2: 255}
        _ = v6
    default:
        var v7 uint16 = 2
        _ = v7
    }
    for fi8 := 0; fi8 < 17; fi8++ {
        var v9 float32 = 1.5
        _ = fi8
        _ = v9
    }
    if false {
        switch -151 {
        case 42:
            for fi10 := 0; fi10 < 3; fi10++ {
                var v11 float64 = (3.14159 - (-2.25 - float64(1.5)))
                var v12 uint64 = 18446744073709551615
                _ = fi10
                _ = v11
                _ = v12
            }
        case 32768:
            v3 += v3
        default:
            switch "" {
            case "":
                for fi13 := 0; fi13 < 6; fi13++ {
                    v14 := GenS1{F0: 0, F1: (57 != -128), F2: 823}
                    _ = fi13
                    _ = v14
                }
            case "foo bar":
                v15 := GenS1{F0: v3, F1: true, F2: 42}
                _ = v15
            case "hello":
                v1 = append(v1, 4294967295)
            default:
                var v16 uint64 = 0
                _ = v16
            }
        }
        var v17 uint64 = 2
        _ = v17
    } else {
        for fi18 := 0; fi18 < 12; fi18++ {
            v3 += (^(v3))
            _ = fi18
        }
    }
    v2 = 258
    fmt.Println(v2, v2)
    var v19 string = (("a" + (("foo bar" + "a") + ("hello" + ""))) + "\n\t")
    for fi20 := 0; fi20 < 14; fi20++ {
        if false {
            var v21 float32 = (0.1 / 0.1)
            v19 = v19
            if false {
                fmt.Println(v19, fi20)
            }
            _ = v21
        } else {
            fmt.Println(v19)
            var v22 float64 = float64((3.14159 + 0.0))
            switch fi20 {
            case -372:
                v3 = 23
            case -9223372036854775808:
                v3 = uint8(v2)
            case 0:
                for fi23 := 0; fi23 < 11; fi23++ {
                    var v24 int8 = 127
                    if len(v1) > 0 {
                        v1[0] = 42
                    }
                    _ = fi23
                    _ = v24
                }
            }
            _ = v22
        }
        v19 = ((v19 + v19) + "")
        _ = fi20
    }
    var v25 int8 = 0
    v26 := GenS1{F0: v3, F1: true, F2: 255}
    v27 := GenS1{F0: v3, F1: (0.1 <= ((float32(((1e10 / 0.1) / -2.25)) * 1e10) - 1e10)), F2: 600}
    switch -2147483648 {
    case 9223372036854775807:
        var v28 int32 = 0
        _ = v28
    case -503:
        v3 = (v3 ^ v3)
    default:
        if ((^(v4)) != v4) {
            fmt.Println(v2, v25)
        } else {
            if true {
                switch ((v19 + "foo bar") + v19) {
                case "\n\t":
                    _ = len(v1)
                }
            }
            switch 128 {
            case -308:
                fmt.Println(v4, v4)
            case 33:
                v3 += 105
            case 128:
                v29 := GenS1{F0: (v3 & v3), F1: false, F2: 65535}
                _ = v29
            }
        }
    }
    v25 = (v25 ^ 1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
