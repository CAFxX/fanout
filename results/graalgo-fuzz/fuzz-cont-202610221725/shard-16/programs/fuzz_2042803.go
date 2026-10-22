package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 string }
func genFn0(p0 uint, p1 int64, p2 uint) int32 {
    v7 := []int{}
    p0 = p2
    var v8 uint64 = 752
    switch "a" {
    case "":
        fmt.Println(p2, p2)
    case "a":
        v9 := GenS1{F0: v8}
        _ = v9
    }
    var v10 float64 = 1.5
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    _ = v8
    return -820
}

func main() {
    v1 := []uint32{42, 2147483647, 2147483647, 65535}
    v2 := make(map[string]uint32)
    var v3 int = 221
    v4 := make(map[int]uint32)
    var v5 int64 = -9223372036854775808
    var v6 uint8 = 42
    for fi11 := 0; fi11 < 1; fi11++ {
        var v12 uint64 = 937
        fmt.Println(v3)
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 19; fi13++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi13
    }
    for fi14 := 0; fi14 < 14; fi14++ {
        v3 = (v3 >> ((451) & 63))
        v5 -= (v5 % ((9223372036854775807) | 1))
        var v15 uint8 = (v6 & (v6 >> ((42) & 63)))
        _ = fi14
        _ = v15
    }
    v16 := GenS2{F0: v3, F1: "\n\t"}
    v17 := v2[("foo bar" + "\n\t")]
    switch ("hello" + ("hello" + "")) {
    case "":
        v18 := GenS1{F0: 4294967295}
        _ = v18
    default:
        var v19 uint16 = 72
        _ = v19
    }
    v20 := GenS1{F0: 189}
    switch (("\n\t" + ("\n\t" + "")) + ("" + "\n\t")) {
    case "\n\t":
        if (((((3.14159 / (((0.1 + 0.1) / float64((-2.25 * (-2.25 * 1.5)))) + -2.25)) - 0.1) - (1e10 + float64(0.0))) + (((0.0 * 1.5) - 3.14159) + (((0.0 / float64(((1e10 - (0.0 * (0.0 * float32(float64((0.0 + float32(((0.0 - 0.0) / ((0.0 + 1.5) / (float64(1e10) + 1e10)))))))))) - -2.25))) - ((3.14159 - 0.1) + (0.0 + 1.5))) - -2.25))) < 0.0) {
            switch "" {
            case "\n\t":
                var v21 uint16 = uint16(v6)
                _ = v21
            case "hello":
                v4[(v3 - (-(v3)))] = (v17 / ((65535) | 1))
            default:
                for fi22 := 0; fi22 < 11; fi22++ {
                    fmt.Println(v17, v6, v17)
                    var v23 int64 = 0
                    _ = fi22
                    _ = v23
                }
            }
            v17 = (v17 / ((973) | 1))
        } else {
            fmt.Println(v3, v5, v5)
        }
    case "hello":
        var v24 int = 127
        _ = v24
    }
    v25 := GenS1{F0: 18446744073709551615}
    v6 = (v6 / (((^(v6))) | 1))
    v26 := GenS2{F0: (v3 >> ((908) & 63)), F1: "\n\t"}
    if false {
        v27 := GenS1{F0: 2}
        v4[v3] = (v17 - (^(v17)))
        fmt.Println(v5)
        _ = v27
    } else {
        if ((3.14159 - 1.5) == 3.14159) {
            var v28 uint = 682
            _ = v28
        } else {
            if true {
                if ((0.0 / 1e10) < (0.0 - (0.1 + 1.5))) {
                    var v29 uint16 = 32767
                    _ = v29
                }
            } else {
                v30 := GenS2{F0: v3, F1: ("\n\t" + "\n\t")}
                _ = v30
            }
            for fi31 := 0; fi31 < 9; fi31++ {
                for fi32 := 0; fi32 < 8; fi32++ {
                    var v33 int64 = v5
                    _ = fi32
                    _ = v33
                }
                _ = fi31
            }
            fmt.Println(v5)
        }
    }
    var v34 int16 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v25
    _ = v26
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
