package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 uint16 }

func main() {
    var v1 uint64 = 4294967295
    var v2 uint = 4294967295
    var v3 bool = true
    v4 := []uint32{}
    switch -2147483648 {
    case 1:
        switch "hello" {
        case "":
            switch -1 {
            case 32767:
                if (!true) {
                    _ = len(v4)
                } else {
                    if (v2 <= 42) {
                        var v5 bool = (!(-1 == -1))
                        var v6 int8 = -121
                        _ = v5
                        _ = v6
                    }
                    switch "hello" {
                    case "":
                        v2 -= v2
                    case "a":
                        v3 = v3
                    case "hello":
                        switch 127 {
                        case -9223372036854775808:
                            v3 = v3
                        }
                    default:
                        var v7 int8 = 103
                        _ = v7
                    }
                }
            }
        default:
            var v8 float64 = (0.1 * float64((float32(-2.25) / (3.14159 - 0.1))))
            _ = v8
        }
    case 32768:
        v9 := GenS1{F0: 2147483647, F1: 0, F2: 0}
        _ = v9
    }
    var v10 string = ("hello" + "hello")
    v11 := GenS1{F0: 984, F1: 65535, F2: 2}
    var v12 float64 = 1.5
    if (!false) {
        v13 := GenS1{F0: 42, F1: 42, F2: 477}
        v12 = (v12 * float64((float32((float64(0.1) / (((v12 * (((v12 + (-2.25 * v12)) / float64((3.14159 / 1.0))) * v12)) + v12) / 1.0))) - 0.1)))
        _ = v13
    }
    v3 = v3
    var v14 uint32 = uint32(v1)
    var v15 float32 = 0.1
    if (0 == 4294967295) {
        for fi16 := 0; fi16 < 13; fi16++ {
            var v17 int16 = int16(v2)
            v18 := GenS1{F0: (v14 ^ v14), F1: (v14 ^ 1), F2: 42}
            v2 = (^(v2))
            _ = fi16
            _ = v17
            _ = v18
        }
    }
    _ = len(v4)
    var v19 int32 = -1
    switch 32768 {
    case 2147483647:
        var v20 string = (v10 + v10)
        _ = v20
    case -862:
        v19 += (v19 * (^(v19)))
    case 902:
        for fi21 := 0; fi21 < 13; fi21++ {
            v22 := GenS1{F0: (v14 & (^(v14))), F1: (v14 << ((v2) & 63)), F2: 255}
            var v23 uint64 = (v1 ^ (v1 | (v1 * 4294967295)))
            var v24 int32 = (v19 << ((v2) & 63))
            _ = fi21
            _ = v22
            _ = v23
            _ = v24
        }
    default:
        v2 -= (v2 & (v2 & 2))
    }
    _ = len(v4)
    v12 = (v12 / (float64(v15) + float64((-2.25 - v15))))
    if (v19 > (-(v19))) {
        v15 += 0.0
        if v3 {
            v1 += (v1 & 4294967295)
            var v25 float64 = 1.5
            _ = v25
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
