package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 bool }
type GenS2 struct { F0 int16 }
func genFn0(p0 bool) int16 {
    v7 := []int{658, 2147483647, 1, 32768}
    fmt.Println(p0)
    switch 32767 {
    case 9223372036854775807:
        switch -67 {
        case 127:
            fmt.Println(p0)
        }
    default:
        fmt.Println(p0)
    }
    _ = p0
    _ = v7
    return 490
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]bool)
    v3 := make(map[string]int)
    var v4 float64 = float64((0.0 + 1.5))
    v5 := make(map[int]string)
    v6 := make(map[string]uint32)
    var v8 bool = (610 < 0)
    for fi9 := 0; fi9 < 12; fi9++ {
        v10 := v2[""]
        var v11 uint64 = 1
        if ((0 > 569) || (110 < uint8(v11))) {
            v11 = 492
            switch ("hello" + "") {
            case "\n\t":
                for fi12 := 0; fi12 < 3; fi12++ {
                    var v13 uint = 18446744073709551615
                    _ = fi12
                    _ = v13
                }
            case "a":
                var v14 uint8 = 89
                _ = v14
            }
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    for fi15 := 0; fi15 < 2; fi15++ {
        for fi16 := 0; fi16 < 15; fi16++ {
            var v17 uint32 = 2147483647
            _ = fi16
            _ = v17
        }
        _ = fi15
    }
    v18 := []int{827, -173}
    var v19 string = ("foo bar" + "\n\t")
    if true {
        v8 = ((v4 <= (v4 * ((v4 + v4) / v4))) && ((true && v8) && (127 < 255)))
        for fi20 := 0; fi20 < 12; fi20++ {
            if true {
                v8 = v8
                if (!((1 != 2) && true)) {
                    _ = len(v18)
                    var v21 bool = v8
                    v4 -= -2.25
                    _ = v21
                }
                v22 := GenS1{F0: (v19 + "hello"), F1: 127, F2: v8}
                _ = v22
            }
            _ = fi20
        }
        v23 := GenS2{F0: 319}
        _ = v23
    }
    v6["\n\t"] = 2147483648
    var v24 bool = (0 >= 2147483648)
    v25 := v5[9223372036854775807]
    for fi26 := 0; fi26 < 16; fi26++ {
        v8 = v8
        v19 = "a"
        _ = fi26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
