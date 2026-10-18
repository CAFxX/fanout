package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
func genFn0() uint16 {
    var v5 float64 = ((1e10 * (0.0 - 1e10)) * float64(1e10))
    switch 4294967295 {
    case 32768:
        var v6 bool = false
        _ = v6
    case 135:
        v5 -= (v5 * (v5 * v5))
    case 511:
        fmt.Println(v5)
    default:
        if true {
            v5 = v5
        }
    }
    v5 -= v5
    for fi7 := 0; fi7 < 12; fi7++ {
        v5 -= (v5 * 3.14159)
        fmt.Println(fi7, v5)
        v5 -= 0.0
        _ = fi7
    }
    _ = v5
    return 1
}

func main() {
    v1 := []string{}
    var v2 int16 = 32767
    v3 := make(map[uint32]string)
    var v4 int32 = -1
    delete(v3, 2)
    for fi8 := 0; fi8 < 18; fi8++ {
        var v9 int = fi8
        switch fi8 {
        case 127:
            v1 = append(v1, ("hello" + "foo bar"))
        default:
            var v10 int32 = v4
            _ = v10
        }
        var v11 float64 = float64((0.1 * -2.25))
        _ = fi8
        _ = v11
        _ = v9
    }
    v12 := GenS1{F0: 65535, F1: -1}
    var v13 float64 = 3.14159
    v4 -= v4
    v4 -= (v4 ^ (^(v4)))
    var v14 string = "foo bar"
    v15 := v3[4294967295]
    var v16 int8 = 1
    var v17 float32 = 1e10
    for fi18 := 0; fi18 < 11; fi18++ {
        v13 -= (float64(0.1) * v13)
        fmt.Println(v15, v15)
        for fi19 := 0; fi19 < 20; fi19++ {
            fmt.Println(fi18, v13)
            if ((v16 >> ((2) & 63)) == (v16 - 42)) {
                fmt.Println(v14)
            }
            v20 := GenS1{F0: 16, F1: v4}
            _ = fi19
            _ = v20
        }
        _ = fi18
    }
    if (1 == 127) {
        if len(v1) > 0 {
            v1[0] = v14
        }
    } else {
        if len(v1) > 0 {
            v1[0] = v14
        }
        v15 = ("\n\t" + v14)
        delete(v3, 594)
    }
    v4 = v4
    fmt.Println(v13, v15, v2)
    fmt.Println(v15, v13)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
