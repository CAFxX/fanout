package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 int64 }
type GenS2 struct { F0 uint64; F1 int8; F2 int8 }

func main() {
    v1 := make(map[int]uint32)
    v2 := []uint32{527, 55, 2147483647}
    var v3 float32 = float32(3.14159)
    var v4 string = (("a" + ("hello" + ("a" + ("foo bar" + "hello")))) + "a")
    if true {
        v5 := GenS1{F0: false, F1: 1, F2: 169}
        fmt.Println(v3, v4)
        v3 = (0.1 * (0.1 - v3))
        _ = v5
    }
    var v6 int8 = 1
    var v7 bool = true
    switch -2147483648 {
    case -1:
        v6 = v6
    case -352:
        if v7 {
            v6 += v6
        }
    case -2147483648:
        var v8 int16 = 1000
        _ = v8
    }
    var v9 uint64 = 42
    switch -9223372036854775808 {
    case 4294967295:
        var v10 int16 = -32768
        _ = v10
    case 0:
        var v11 uint16 = 2
        _ = v11
    case 127:
        switch "" {
        case "foo bar":
            v12 := GenS2{F0: v9, F1: v6, F2: (-(v6))}
            _ = v12
        case "":
            for fi13 := 0; fi13 < 3; fi13++ {
                var v14 bool = v7
                v14 = ((float32(3.14159) - (v3 - (v3 - v3))) >= v3)
                _ = fi13
                _ = v14
            }
        case "a":
            v6 -= -128
        default:
            switch v4 {
            case "a":
                switch v4 {
                case "foo bar":
                    for fi15 := 0; fi15 < 4; fi15++ {
                        for fi16 := 0; fi16 < 19; fi16++ {
                            v6 = (v6 - -1)
                            v7 = false
                            _ = fi16
                        }
                        _ = fi15
                    }
                case "\n\t":
                    v9 += 796
                case "a":
                    fmt.Println(v9)
                default:
                    var v17 uint = 4294967295
                    _ = v17
                }
            default:
                fmt.Println(v7, v9)
            }
        }
    default:
        v9 -= (v9 * (^(v9)))
    }
    if (((v4 + v4) + (v4 + v4)) <= v4) {
        var v18 string = "foo bar"
        if true {
            if v7 {
                var v19 uint8 = 109
                var v20 bool = v7
                _ = v19
                _ = v20
            }
            var v21 int64 = 0
            v22 := GenS2{F0: (v9 / (((^(v9))) | 1)), F1: (^(v6)), F2: v6}
            _ = v21
            _ = v22
        }
        _ = v18
    } else {
        for fi23 := 0; fi23 < 17; fi23++ {
            _ = len(v1)
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            for fi24 := 0; fi24 < 8; fi24++ {
                var v25 uint64 = (v9 & v9)
                fmt.Println(v9)
                _ = fi24
                _ = v25
            }
            _ = fi23
        }
    }
    if v7 {
        for fi26 := 0; fi26 < 4; fi26++ {
            v1[(-(fi26))] = 575
            var v27 bool = false
            fmt.Println(v9, v7, fi26)
            _ = fi26
            _ = v27
        }
    }
    v7 = ((!v7) && (v7 || true))
    v4 = v4
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v9
}
