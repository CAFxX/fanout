package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 int32; F1 int8; F2 uint8 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 int16, p1 int64) uint64 {
    switch -2147483648 {
    case 32768:
        v6 := GenS1{F0: 132, F1: -180}
        _ = v6
    default:
        if false {
            var v7 int16 = -676
            v8 := GenS2{F0: 552, F1: 127, F2: 1}
            _ = v7
            _ = v8
        } else {
            p1 = (p1 | p1)
            p0 += p0
            p1 += (p1 & p1)
        }
    }
    for fi9 := 0; fi9 < 5; fi9++ {
        for fi10 := 0; fi10 < 5; fi10++ {
            fmt.Println(p1, fi9)
            var v11 float32 = 0.0
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    _ = p0
    _ = p1
    return 42
}

func main() {
    var v1 uint32 = 414
    var v2 uint16 = 65535
    var v3 uint = 1
    v4 := make(map[uint32]string)
    v5 := make(map[uint32]int)
    v12 := []int{4294967295, -74, -867, 705}
    if ((!true) || (((((-2.25 - (-2.25 - (1.5 - 1.5))) + 1e10) - float32((((3.14159 * 0.1) + 1e10) * 1e10))) * (0.0 * 0.0)) > 0.0)) {
        fmt.Println(v3, v1)
        v2 = (v2 - (v2 >> ((65535) & 63)))
        v13 := GenS3{F0: 569}
        _ = v13
    } else {
        switch "foo bar" {
        case "":
            var v14 float32 = -2.25
            _ = v14
        case "foo bar":
            v12 = append(v12, 1)
        default:
            v15 := GenS1{F0: 255, F1: -9223372036854775808}
            _ = v15
        }
        fmt.Println(v1, v2, v1)
        var v16 bool = (-2147483648 != -640)
        _ = v16
    }
    v12 = append(v12, -9223372036854775808)
    var v17 uint16 = (v2 << ((955) & 63))
    v1 = uint32(v3)
    v17 = v2
    var v18 uint = (v3 / ((v3) | 1))
    if len(v12) > 0 {
        v12[0] = 470
    }
    var v19 int8 = -109
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
