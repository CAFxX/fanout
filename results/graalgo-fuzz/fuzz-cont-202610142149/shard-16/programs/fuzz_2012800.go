package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 string; F1 uint64; F2 int8 }
func genFn0(p0 int8, p1 uint32, p2 uint32) int {
    fmt.Println(p2, p0)
    v4 := []int{}
    if (!false) {
        v5 := GenS1{F0: 127, F1: 9223372036854775807}
        _ = v5
    } else {
        _ = len(v4)
        p2 = p1
        v6 := GenS1{F0: 127, F1: 42}
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 2147483647
}

func main() {
    var v1 float64 = (0.1 / 1.5)
    v2 := []uint32{2147483648, 624}
    var v3 int16 = 708
    v1 = v1
    switch "foo bar" {
    case "\n\t":
        if len(v2) > 0 {
            v2[0] = 42
        }
    case "foo bar":
        v7 := GenS1{F0: 100, F1: 1}
        _ = v7
    }
    switch "" {
    case "":
        var v8 bool = (false || false)
        _ = v8
    case "foo bar":
        if false {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v1 = float64((1e10 - 0.1))
        } else {
            for fi9 := 0; fi9 < 4; fi9++ {
                for fi10 := 0; fi10 < 17; fi10++ {
                    if len(v2) > 0 {
                        v2[0] = 893
                    }
                    if len(v2) > 0 {
                        v2[0] = 4294967295
                    }
                    _ = fi10
                }
                _ = fi9
            }
            switch 128 {
            case 9223372036854775807:
                v11 := GenS1{F0: 127, F1: 9223372036854775807}
                _ = v11
            case 4294967295:
                var v12 int16 = v3
                _ = v12
            case 32768:
                switch "a" {
                case "":
                    v1 = v1
                case "foo bar":
                    fmt.Println(v3, v1)
                default:
                    v3 -= (-(v3))
                }
            default:
                var v13 uint16 = 593
                _ = v13
            }
        }
    case "\n\t":
        v14 := GenS2{F0: "\n\t", F1: 0, F2: 127}
        _ = v14
    }
    if false {
        v15 := GenS1{F0: 83, F1: -1}
        _ = v15
    }
    v3 = v3
    switch ("\n\t" + ("" + (("" + "\n\t") + "foo bar"))) {
    case "hello":
        v1 += (-2.25 - v1)
    case "":
        v1 = v1
    case "\n\t":
        var v16 float64 = v1
        _ = v16
    default:
        for fi17 := 0; fi17 < 14; fi17++ {
            if (!false) {
                for fi18 := 0; fi18 < 7; fi18++ {
                    _ = len(v2)
                    v3 = 42
                    _ = fi18
                }
                var v19 int8 = 126
                var v20 int16 = (v3 >> ((228) & 63))
                _ = v19
                _ = v20
            }
            _ = fi17
        }
    }
    v1 -= ((((3.14159 / v1) * -2.25) / ((-2.25 * 3.14159) * v1)) / (v1 - 0.0))
    v21 := GenS2{F0: "a", F1: 2, F2: 0}
    v1 = v1
    for fi22 := 0; fi22 < 11; fi22++ {
        v23 := GenS2{F0: "", F1: 0, F2: -128}
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi22
        _ = v23
    }
    var v24 int64 = -966
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if true {
        var v25 uint8 = 252
        switch -1 {
        case 506:
            v25 = (v25 | (^(v25)))
        case 1:
            v26 := GenS1{F0: -74, F1: (v24 % (((v24 % (((^(v24))) | 1))) | 1))}
            _ = v26
        case 2147483647:
            var v27 int = -579
            _ = v27
        }
        _ = v25
    }
    v3 -= (-(v3))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v24
    _ = v3
}
