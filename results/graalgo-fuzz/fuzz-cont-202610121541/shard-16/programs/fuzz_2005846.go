package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 int32; F1 uint32; F2 uint16 }

func main() {
    v1 := []int{}
    var v2 uint32 = 214
    var v3 uint32 = 2147483647
    v4 := GenS1{F0: -1, F1: 4294967295}
    switch -2147483648 {
    case 0:
        var v5 string = "foo bar"
        _ = v5
    case 32768:
        var v6 int32 = 765
        _ = v6
    case 32767:
        v3 -= v3
    }
    if (((-964 != 9223372036854775807) && true) || (2 >= 66)) {
        var v7 uint16 = 175
        v2 = (v3 | (v2 >> ((2) & 63)))
        _ = v7
    } else {
        var v8 int8 = 83
        var v9 uint8 = uint8(v3)
        _ = v8
        _ = v9
    }
    var v10 int = 32767
    switch "hello" {
    case "":
        var v11 float32 = 0.0
        _ = v11
    default:
        if (!(false || true)) {
            v10 = v10
        }
    }
    v10 -= v10
    for fi12 := 0; fi12 < 7; fi12++ {
        v1 = append(v1, fi12)
        for fi13 := 0; fi13 < 7; fi13++ {
            v2 = 2147483648
            v10 += 1
            _ = fi13
        }
        switch fi12 {
        case 0:
            if true {
                for fi14 := 0; fi14 < 3; fi14++ {
                    var v15 uint8 = 1
                    v10 -= (v10 | 112)
                    if len(v1) > 0 {
                        v1[0] = (fi14 | v10)
                    }
                    _ = fi14
                    _ = v15
                }
            }
        default:
            switch (v10 >> ((1) & 63)) {
            case -5:
                v3 -= 381
            case 2147483647:
                for fi16 := 0; fi16 < 4; fi16++ {
                    _ = len(v1)
                    v17 := GenS2{F0: -2147483648, F1: 0, F2: 843}
                    fmt.Println(v3)
                    _ = fi16
                    _ = v17
                }
            case -171:
                var v18 uint16 = 60
                _ = v18
            default:
                var v19 string = "hello"
                _ = v19
            }
        }
        _ = fi12
    }
    var v20 uint8 = 2
    switch (("" + "") + "foo bar") {
    case "\n\t":
        var v21 uint16 = 32767
        _ = v21
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
