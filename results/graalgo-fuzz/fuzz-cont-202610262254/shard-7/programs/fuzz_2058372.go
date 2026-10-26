package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 uint32 }
func genFn0(p0 int16) int {
    fmt.Println(p0)
    fmt.Println(p0)
    var v4 bool = (false || ((2 == 65535) && false))
    _ = p0
    _ = v4
    return 32768
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []string{"\n\t", "\n\t"}
    v3 := make(map[int]uint32)
    switch -1 {
    case 42:
        for fi5 := 0; fi5 < 8; fi5++ {
            fmt.Println(fi5)
            for fi6 := 0; fi6 < 2; fi6++ {
                var v7 float64 = float64(-2.25)
                _ = fi6
                _ = v7
            }
            _ = fi5
        }
    case -818:
        var v8 int64 = 42
        _ = v8
    case 32767:
        var v9 float64 = (3.14159 / 1e10)
        _ = v9
    default:
        v10 := v3[genFn0(-32768)]
        _ = v10
    }
    var v11 uint32 = 4294967295
    if (!(-128 <= -128)) {
        v11 += 1
    }
    v11 += (v11 << ((1) & 63))
    var v12 bool = (!(1 <= 2))
    if (1 >= 1) {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    } else {
        switch "foo bar" {
        case "foo bar":
            for fi13 := 0; fi13 < 17; fi13++ {
                for fi14 := 0; fi14 < 10; fi14++ {
                    _ = len(v1)
                    if v12 {
                        v12 = (true || (v12 && true))
                        switch "foo bar" {
                        case "":
                            v11 = v11
                        case "\n\t":
                            var v15 int = fi13
                            _ = v15
                        default:
                            v11 = 0
                        }
                    }
                    v16 := GenS1{F0: 42, F1: -55, F2: v11}
                    var v17 float64 = 0.0
                    _ = fi14
                    _ = v16
                    _ = v17
                }
                delete(v1, "\n\t")
                _ = fi13
            }
        case "a":
            fmt.Println(v11, v12)
        case "":
            v11 -= v11
        }
        if ((((v11 & 38) < (v11 / ((v11) | 1))) && false) || v12) {
            switch 1 {
            case 9223372036854775807:
                fmt.Println(v12, v11)
            default:
                v11 = 4294967295
            }
        }
    }
    fmt.Println(v11, v12)
    var v18 uint32 = v11
    v18 = (v18 | v11)
    v12 = v12
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v12
    _ = v18
    _ = v2
    _ = v3
}
