package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
func genFn0() int8 {
    var v4 float32 = 0.1
    fmt.Println(v4)
    _ = v4
    return 31
}

func main() {
    v1 := []string{"foo bar", "", ("hello" + "foo bar")}
    v2 := make(map[string]bool)
    var v3 int = -9223372036854775808
    switch "" {
    case "\n\t":
        var v5 int16 = 974
        _ = v5
    }
    v3 += v3
    if false {
        for fi6 := 0; fi6 < 18; fi6++ {
            if ((1000 < -44) && false) {
                delete(v2, "a")
                v3 = -715
                for fi7 := 0; fi7 < 17; fi7++ {
                    v8 := v2[("\n\t" + ("hello" + ""))]
                    _ = fi7
                    _ = v8
                }
            }
            v9 := GenS1{F0: 935, F1: 2}
            _ = fi6
            _ = v9
        }
        var v10 float32 = (((-2.25 + (((0.0 - 0.0) * -2.25) - float32(1e10))) + (3.14159 / 1.5)) * -2.25)
        var v11 uint8 = 1
        _ = v10
        _ = v11
    } else {
        var v12 float32 = 1e10
        for fi13 := 0; fi13 < 7; fi13++ {
            v3 = -985
            for fi14 := 0; fi14 < 9; fi14++ {
                v3 += v3
                _ = fi14
            }
            _ = fi13
        }
        _ = v12
    }
    v15 := GenS1{F0: -97, F1: 128}
    var v16 bool = (32767 <= 42)
    v17 := GenS1{F0: 32767, F1: 1}
    v18 := v2[("\n\t" + "foo bar")]
    v3 = (v3 << ((2) & 63))
    v3 -= v3
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    _ = len(v2)
    v16 = v18
    v19 := GenS1{F0: -1, F1: 0}
    for fi20 := 0; fi20 < 7; fi20++ {
        var v21 float64 = 1.5
        v3 = fi20
        _ = fi20
        _ = v21
    }
    var v22 uint32 = 2
    v23 := GenS1{F0: -32768, F1: 127}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v3
}
