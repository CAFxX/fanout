package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
func genFn0() int {
    var v7 int8 = -1
    var v8 int16 = -355
    fmt.Println(v8, v7)
    _ = v7
    _ = v8
    return 522
}

func main() {
    var v1 int16 = 651
    v2 := []int{32767, 1, -1, -9223372036854775808}
    v3 := make(map[int]bool)
    var v4 int64 = 1
    var v5 uint64 = 2
    v6 := make(map[int]string)
    var v9 int64 = (v4 << ((881) & 63))
    _ = len(v2)
    var v10 float32 = (1e10 / 1.0)
    v11 := GenS1{F0: v4, F1: -84}
    var v12 uint32 = 860
    for fi13 := 0; fi13 < 16; fi13++ {
        v14 := GenS1{F0: v9, F1: 5}
        _ = fi13
        _ = v14
    }
    fmt.Println(v5, v9, v5)
    var v15 uint8 = 0
    for fi16 := 0; fi16 < 9; fi16++ {
        for fi17 := 0; fi17 < 13; fi17++ {
            v12 = v12
            _ = fi17
        }
        var v18 bool = true
        _ = fi16
        _ = v18
    }
    var v19 uint8 = v15
    var v20 float64 = ((0.1 / 3.14159) + ((-2.25 / (((0.0 + ((1e10 - float64((0.0 + v10))) - ((((0.0 - -2.25) - 3.14159) + 0.1) - -2.25))) * 0.0) / (0.1 / 3.14159))) / 0.1))
    v20 -= -2.25
    for fi21 := 0; fi21 < 10; fi21++ {
        var v22 int32 = 726
        switch "a" {
        case "foo bar":
            switch int(v1) {
            case 0:
                var v23 string = ((("" + ("\n\t" + "a")) + "hello") + "")
                _ = v23
            case 9223372036854775807:
                fmt.Println(v20)
            case 42:
                v1 = 1
            }
        case "hello":
            v24 := GenS1{F0: 9223372036854775807, F1: 89}
            _ = v24
        case "\n\t":
            var v25 float64 = v20
            _ = v25
        default:
            v20 += float64(1.5)
        }
        v26 := v3[fi21]
        _ = fi21
        _ = v22
        _ = v26
    }
    var v27 uint16 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
