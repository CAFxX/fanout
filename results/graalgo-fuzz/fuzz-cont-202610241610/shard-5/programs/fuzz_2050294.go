package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }

func main() {
    v1 := make(map[uint32]int)
    var v2 float32 = 1e10
    var v3 float64 = float64((((3.14159 - v2) / 3.14159) * (v2 + v2)))
    v4 := make(map[string]uint32)
    if true {
        v5 := v4[""]
        for fi6 := 0; fi6 < 11; fi6++ {
            v5 = v5
            _ = fi6
        }
        _ = v5
    } else {
        v3 -= (v3 - 0.1)
        v2 = 1.5
    }
    fmt.Println(v2)
    switch 1 {
    case 0:
        for fi7 := 0; fi7 < 13; fi7++ {
            var v8 uint8 = 0
            for fi9 := 0; fi9 < 8; fi9++ {
                fmt.Println(fi7, v2)
                _ = len(v1)
                var v10 uint16 = 1
                _ = fi9
                _ = v10
            }
            _ = fi7
            _ = v8
        }
    case -71:
        v3 = 1.5
    case 4294967295:
        _ = len(v4)
    default:
        v11 := GenS1{F0: 42, F1: 0}
        _ = v11
    }
    var v12 uint16 = 219
    fmt.Println(v3, v12)
    for fi13 := 0; fi13 < 15; fi13++ {
        switch (^(fi13)) {
        case -919:
            if true {
                for fi14 := 0; fi14 < 3; fi14++ {
                    v3 -= v3
                    switch fi14 {
                    case -187:
                        v3 -= (v3 - (float64(((((1e10 + 1.5) - -2.25) * (((1.5 + 0.0) / (v2 / (v2 * v2))) * float32(v3))) / v2)) - (((float64((((v2 - 0.1) + v2) / (v2 * ((0.0 / v2) * v2)))) * v3) + ((v3 - v3) * float64(v2))) + v3)))
                    }
                    _ = fi14
                }
            }
        case 42:
            var v15 int32 = 0
            _ = v15
        }
        v16 := []int{}
        _ = len(v16)
        _ = fi13
        _ = v16
    }
    v17 := GenS1{F0: uint(v12), F1: -274}
    if (!((("\n\t" + ("\n\t" + "\n\t")) + "a") < "a")) {
        v3 = v3
        var v18 uint64 = 18446744073709551615
        _ = v18
    }
    v1[2147483648] = 42
    for fi19 := 0; fi19 < 4; fi19++ {
        v2 = float32((v3 / 1.5))
        for fi20 := 0; fi20 < 3; fi20++ {
            v12 = (^(v12))
            fmt.Println(v2, fi20)
            _ = fi20
        }
        _ = fi19
    }
    if (!true) {
        delete(v1, 65535)
    } else {
        var v21 int32 = 42
        v22 := GenS1{F0: 308, F1: 992}
        _ = v21
        _ = v22
    }
    var v23 float32 = 0.0
    v2 = v23
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
