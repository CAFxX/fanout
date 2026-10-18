package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 int16 }
type GenS2 struct { F0 int8; F1 int32 }

func main() {
    var v1 int32 = -706
    v2 := []uint32{636, 2, 1, 887}
    v3 := []string{("\n\t" + "\n\t"), "\n\t"}
    var v4 uint = 18446744073709551615
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v5 bool = false
    if v5 {
        v4 = 2
        if true {
            var v6 uint64 = 1
            _ = v6
        } else {
            switch int(v1) {
            case 273:
                switch 2147483647 {
                case 720:
                    v1 = 194
                case 2147483647:
                    var v7 uint32 = 2147483647
                    _ = v7
                case 4294967295:
                    if (v5 && (true && false)) {
                        var v8 bool = (true && v5)
                        _ = v8
                    }
                }
            case 127:
                for fi9 := 0; fi9 < 14; fi9++ {
                    for fi10 := 0; fi10 < 20; fi10++ {
                        fmt.Println(v4)
                        v5 = v5
                        v11 := GenS1{F0: -9223372036854775808, F1: -533, F2: -1}
                        _ = fi10
                        _ = v11
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi9
                }
            case 9223372036854775807:
                v12 := GenS2{F0: 42, F1: v1}
                _ = v12
            }
            for fi13 := 0; fi13 < 3; fi13++ {
                _ = len(v2)
                switch "hello" {
                case "":
                    v2 = append(v2, 4294967295)
                case "a":
                    var v14 int8 = -1
                    _ = v14
                }
                for fi15 := 0; fi15 < 20; fi15++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi15
                }
                _ = fi13
            }
            fmt.Println(v5, v5)
        }
    } else {
        var v16 int64 = 0
        fmt.Println(v1, v4, v1)
        var v17 uint8 = 255
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 14; fi18++ {
        var v19 bool = (((false || v5) || (-1 > -128)) && (65535 <= (v4 ^ 4294967295)))
        if (false || (-1 == -242)) {
            fmt.Println(fi18)
            fmt.Println(v5, fi18, v5)
            fmt.Println(v19, v5)
        }
        _ = fi18
        _ = v19
    }
    switch "a" {
    case "foo bar":
        v4 = v4
    case "a":
        v4 = (v4 & (^(v4)))
    case "\n\t":
        var v20 int32 = (^(v1))
        _ = v20
    }
    v4 = v4
    v1 = -1
    v1 -= -2147483648
    var v21 int16 = 42
    fmt.Println(v5, v5, v5)
    var v22 float32 = (3.14159 * 3.14159)
    v1 += (v1 - 1)
    switch 840 {
    case 4294967295:
        for fi23 := 0; fi23 < 18; fi23++ {
            v21 = -525
            _ = fi23
        }
    default:
        if ((1.5 - -2.25) >= v22) {
            for fi24 := 0; fi24 < 1; fi24++ {
                v1 = (v1 - (v1 & 1))
                _ = fi24
            }
        }
    }
    if true {
        v21 = v21
    }
    v1 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
