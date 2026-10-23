package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 uint64 }
type GenS2 struct { F0 uint32; F1 int16; F2 int16 }
type GenS3 struct { F0 int16; F1 bool }

func main() {
    v1 := []uint32{65535, 2, 2}
    var v2 float64 = 1e10
    var v3 uint32 = 440
    v4 := []string{"foo bar"}
    v5 := GenS1{F0: -326, F1: (v3 | v3), F2: 2}
    v2 += v2
    if false {
        var v6 uint64 = 18446744073709551615
        if true {
            v2 += v2
            switch "\n\t" {
            case "foo bar":
                for fi7 := 0; fi7 < 7; fi7++ {
                    for fi8 := 0; fi8 < 4; fi8++ {
                        _ = len(v1)
                        _ = fi8
                    }
                    fmt.Println(v6)
                    _ = fi7
                }
            case "a":
                v6 = v6
            case "":
                _ = len(v1)
            default:
                v9 := GenS1{F0: -761, F1: (v3 + (v3 * (v3 % (((^(v3))) | 1)))), F2: (v6 + (v6 << ((2) & 63)))}
                _ = v9
            }
            v10 := GenS1{F0: 0, F1: 42, F2: 0}
            _ = v10
        }
        _ = v6
    }
    _ = len(v1)
    v3 -= v3
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    _ = len(v4)
    if (!true) {
        if false {
            var v11 int32 = 664
            _ = v11
        }
    }
    switch "a" {
    case "":
        var v12 int = 42
        _ = v12
    }
    var v13 int8 = 83
    if ((v3 >> ((42) & 63)) > (v3 >> ((0) & 63))) {
        if ((false && ("foo bar" == "foo bar")) || false) {
            _ = len(v4)
            var v14 uint8 = uint8(v3)
            _ = v14
        } else {
            fmt.Println(v2, v3)
        }
        v15 := GenS2{F0: (v3 ^ v3), F1: 869, F2: 32767}
        v13 -= -1
        _ = v15
    } else {
        v2 += float64(float32((v2 - v2)))
        if (v13 != -1) {
            switch 0 {
            case 127:
                for fi16 := 0; fi16 < 14; fi16++ {
                    switch fi16 {
                    case 128:
                        fmt.Println(v2, v3)
                    default:
                        v1 = append(v1, (v3 | (v3 * (v3 * (v3 ^ 4294967295)))))
                    }
                    v13 += 0
                    if false {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    } else {
                        var v17 bool = true
                        _ = v17
                    }
                    _ = fi16
                }
            case 209:
                v18 := GenS3{F0: -389, F1: (!(1e10 <= -2.25))}
                _ = v18
            case 152:
                var v19 uint16 = 578
                _ = v19
            }
            var v20 bool = false
            v13 -= v13
            _ = v20
        } else {
            var v21 uint32 = v3
            _ = v21
        }
        var v22 string = "foo bar"
        _ = v22
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
