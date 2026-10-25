package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 uint32; F1 int16; F2 uint16 }

func main() {
    var v1 string = "\n\t"
    var v2 uint32 = 4294967295
    var v3 string = (("\n\t" + (("foo bar" + ((((((v1 + "") + "hello") + v1) + v1) + v1) + "")) + v1)) + "hello")
    v4 := []uint32{v2, (v2 & v2)}
    var v5 uint64 = 77
    fmt.Println(v1, v2)
    v2 = (v2 - (^(v2)))
    if ((v5 / ((v5) | 1)) == v5) {
        var v6 uint32 = 65535
        _ = v6
    }
    _ = len(v4)
    v5 -= v5
    for fi7 := 0; fi7 < 5; fi7++ {
        for fi8 := 0; fi8 < 11; fi8++ {
            switch 0 {
            case 9223372036854775807:
                v4 = append(v4, 835)
            case -731:
                switch fi8 {
                case 0:
                    var v9 uint32 = (^(v2))
                    _ = v9
                case 128:
                    v3 = v1
                case 42:
                    var v10 bool = (397 <= 9223372036854775807)
                    _ = v10
                default:
                    fmt.Println(fi7)
                    switch 969 {
                    case -9223372036854775808:
                        v4 = append(v4, 1)
                    case 148:
                        var v11 uint16 = 2
                        _ = v11
                    }
                    v12 := GenS1{F0: 804, F1: v2}
                    _ = v12
                }
            case 127:
                fmt.Println(fi7, fi7)
            default:
                if len(v4) > 0 {
                    v4[0] = v2
                }
            }
            _ = fi8
        }
        for fi13 := 0; fi13 < 20; fi13++ {
            var v14 int8 = 100
            var v15 string = "a"
            var v16 int32 = 2147483647
            _ = fi13
            _ = v14
            _ = v15
            _ = v16
        }
        _ = fi7
    }
    v2 = v2
    var v17 int = 9223372036854775807
    for fi18 := 0; fi18 < 11; fi18++ {
        v17 -= (fi18 & -1)
        if len(v4) > 0 {
            v4[0] = (v2 & 995)
        }
        _ = len(v4)
        _ = fi18
    }
    v19 := GenS1{F0: -9223372036854775808, F1: 1}
    switch v1 {
    case "a":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case "\n\t":
        v5 += (^(v5))
    }
    v20 := GenS2{F0: (v2 & v2), F1: -1, F2: 2}
    if len(v4) > 0 {
        v4[0] = (v2 % ((v2) | 1))
    }
    var v21 string = v1
    if true {
        v22 := GenS1{F0: 9223372036854775807, F1: (v2 >> ((4294967295) & 63))}
        if (!true) {
            var v23 uint8 = 159
            _ = v23
        }
        _ = v22
    }
    switch v17 {
    case 725:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case 128:
        v24 := GenS2{F0: v2, F1: 1000, F2: 730}
        _ = v24
    default:
        var v25 int = (v17 / (((v17 ^ v17)) | 1))
        _ = v25
    }
    var v26 uint = 42
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
