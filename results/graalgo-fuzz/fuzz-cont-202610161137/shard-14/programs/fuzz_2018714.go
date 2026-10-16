package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 int }

func main() {
    v1 := []uint32{0, 4294967295}
    v2 := []uint32{1, 792, 4294967295}
    v3 := []int{-866, 4294967295}
    v4 := make(map[int]string)
    var v5 int32 = -1
    for fi6 := 0; fi6 < 20; fi6++ {
        for fi7 := 0; fi7 < 3; fi7++ {
            v5 = (v5 + -413)
            _ = len(v4)
            v5 += (-(v5))
            _ = fi7
        }
        switch (("a" + "") + "a") {
        case "hello":
            v8 := GenS1{F0: 2, F1: 217, F2: fi6}
            _ = v8
        default:
            for fi9 := 0; fi9 < 20; fi9++ {
                _ = len(v4)
                _ = fi9
            }
        }
        _ = fi6
    }
    var v10 uint32 = 82
    for fi11 := 0; fi11 < 13; fi11++ {
        var v12 int8 = 42
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 18; fi13++ {
        v10 = 99
        _ = fi13
    }
    switch -1 {
    case 795:
        for fi14 := 0; fi14 < 20; fi14++ {
            v10 = v10
            switch fi14 {
            case 0:
                fmt.Println(v5)
            case 9223372036854775807:
                v5 = (v5 - (v5 >> ((2) & 63)))
            case -284:
                if len(v2) > 0 {
                    v2[0] = v10
                }
            default:
                if (false || (!true)) {
                    v5 = -2147483648
                }
            }
            for fi15 := 0; fi15 < 1; fi15++ {
                var v16 string = "foo bar"
                if len(v1) > 0 {
                    v1[0] = v10
                }
                _ = fi15
                _ = v16
            }
            _ = fi14
        }
    default:
        v10 = (v10 / (((^(v10))) | 1))
    }
    delete(v4, 9223372036854775807)
    var v17 int8 = 1
    if (-1 <= -32768) {
        var v18 int32 = v5
        _ = v18
    } else {
        var v19 uint8 = 42
        for fi20 := 0; fi20 < 11; fi20++ {
            switch "a" {
            case "":
                var v21 uint32 = 719
                _ = v21
            case "hello":
                if true {
                    v17 -= v17
                }
            case "a":
                var v22 uint64 = 18446744073709551615
                _ = v22
            default:
                var v23 uint8 = v19
                _ = v23
            }
            delete(v4, (fi20 + fi20))
            v19 -= (v19 / ((v19) | 1))
            _ = fi20
        }
        var v24 uint16 = 42
        _ = v19
        _ = v24
    }
    fmt.Println(v5, v5, v17)
    v17 = v17
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
