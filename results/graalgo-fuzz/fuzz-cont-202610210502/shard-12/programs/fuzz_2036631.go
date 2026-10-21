package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 uint32 }

func main() {
    v1 := make(map[uint32]int)
    var v2 bool = (65535 != 42)
    var v3 uint32 = 109
    var v4 uint16 = 0
    var v5 uint64 = 4294967295
    v6 := make(map[int]uint32)
    _ = len(v1)
    v7 := []int{128, -794, -9223372036854775808}
    switch ("foo bar" + ("\n\t" + (("a" + "foo bar") + ""))) {
    case "foo bar":
        switch "" {
        case "":
            v8 := GenS1{F0: (!((127 == 89) && v2)), F1: "foo bar"}
            _ = v8
        default:
            v4 = 875
        }
    case "hello":
        for fi9 := 0; fi9 < 13; fi9++ {
            var v10 uint64 = uint64(v3)
            _ = fi9
            _ = v10
        }
    case "\n\t":
        var v11 int16 = -32768
        _ = v11
    default:
        if v2 {
            if v2 {
                var v12 int64 = -9223372036854775808
                if ((-1 >= 100) && (false && (v3 <= v3))) {
                    v3 = 1
                }
                v2 = v2
                _ = v12
            }
            v5 -= (v5 & (v5 * (^(v5))))
            v3 = 151
        } else {
            var v13 uint = 350
            _ = v13
        }
    }
    v3 -= v3
    v4 = 42
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    v14 := GenS2{F0: v3}
    var v15 string = "\n\t"
    v15 = (v15 + "")
    v7 = append(v7, 0)
    v15 = "hello"
    v5 -= 2
    var v16 float32 = 3.14159
    switch 1 {
    case -121:
        var v17 bool = true
        _ = v17
    default:
        var v18 int16 = -1
        _ = v18
    }
    fmt.Println(v3, v16, v15)
    for fi19 := 0; fi19 < 11; fi19++ {
        var v20 int32 = 0
        var v21 int16 = 42
        _ = fi19
        _ = v20
        _ = v21
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
