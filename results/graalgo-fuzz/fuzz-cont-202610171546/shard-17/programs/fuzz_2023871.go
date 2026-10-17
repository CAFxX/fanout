package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 int32; F1 int8; F2 uint64 }
func genFn0(p0 int8, p1 uint8) string {
    p1 = (p1 ^ p1)
    p0 = 127
    for fi7 := 0; fi7 < 14; fi7++ {
        v8 := GenS1{F0: 32767, F1: ("\n\t" + "a")}
        fmt.Println(fi7)
        switch ("" + "foo bar") {
        case "":
            p1 -= (p1 * p1)
        default:
            v9 := GenS2{F0: 0, F1: p0, F2: 723}
            _ = v9
        }
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    return "foo bar"
}
func genFn1(p0 int, p1 uint64, p2 int8) uint8 {
    var v10 float32 = 0.0
    v11 := GenS1{F0: 32767, F1: ("\n\t" + ("foo bar" + ""))}
    p0 += (-(p0))
    var v12 float64 = -2.25
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v12
    return 123
}

func main() {
    var v1 uint16 = 870
    var v2 uint = 351
    var v3 uint32 = 655
    var v4 int8 = -85
    v5 := make(map[int]uint32)
    v6 := make(map[string]string)
    for fi13 := 0; fi13 < 6; fi13++ {
        v4 = v4
        switch fi13 {
        case -1:
            switch 702 {
            case 9223372036854775807:
                fmt.Println(v2, v3)
            default:
                var v14 uint16 = v1
                _ = v14
            }
        case 108:
            v4 += (v4 * v4)
        default:
            v15 := []int{}
            _ = v15
        }
        _ = fi13
    }
    v6["foo bar"] = "\n\t"
    if false {
        var v16 uint = (v2 ^ v2)
        v2 = v16
        _ = v16
    }
    v17 := GenS1{F0: -987, F1: ""}
    fmt.Println(v2, v4)
    for fi18 := 0; fi18 < 1; fi18++ {
        v3 += (v3 >> ((v2) & 63))
        _ = fi18
    }
    for fi19 := 0; fi19 < 3; fi19++ {
        v3 -= v3
        var v20 int32 = -974
        fmt.Println(v3, v4, v4)
        _ = fi19
        _ = v20
    }
    delete(v6, "\n\t")
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
