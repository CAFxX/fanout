package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 int; F2 int8 }
type GenS3 struct { F0 uint8; F1 int16 }
func genFn0(p0 uint32, p1 int32, p2 string) uint32 {
    v7 := GenS1{F0: 157}
    switch 317 {
    case -554:
        fmt.Println(p0)
    case 128:
        for fi8 := 0; fi8 < 13; fi8++ {
            v9 := []int{fi8}
            p0 = p0
            var v10 int = 4294967295
            _ = fi8
            _ = v10
            _ = v9
        }
    default:
        fmt.Println(p0, p1, p2)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return p0
}

func main() {
    v1 := make(map[string]string)
    var v2 uint32 = 981
    var v3 int8 = 34
    var v4 string = "foo bar"
    v5 := []uint32{v2, 65535, 42}
    var v6 float32 = -2.25
    fmt.Println(v2, v2)
    v11 := GenS1{F0: -1}
    fmt.Println(v6, v3, v4)
    var v12 uint16 = 0
    fmt.Println(v12, v6)
    if (!true) {
        v12 -= (v12 % (((^(v12))) | 1))
        v13 := GenS2{F0: (v12 ^ v12), F1: 128, F2: v3}
        v14 := GenS3{F0: 255, F1: 658}
        _ = v13
        _ = v14
    }
    var v15 string = ((((v4 + ("\n\t" + v4)) + v4) + "\n\t") + (((v4 + v4) + "foo bar") + v4))
    if len(v5) > 0 {
        v5[0] = v2
    }
    for fi16 := 0; fi16 < 4; fi16++ {
        v5 = append(v5, (v2 % ((68) | 1)))
        for fi17 := 0; fi17 < 19; fi17++ {
            fmt.Println(v3)
            var v18 uint32 = (v2 * 65535)
            v5 = append(v5, v2)
            _ = fi17
            _ = v18
        }
        _ = fi16
    }
    var v19 float64 = (((-2.25 / (0.1 - 3.14159)) - 0.0) - (((((1.5 * ((1.5 * -2.25) + float64(v6))) + (0.1 - ((3.14159 * 1.5) + 0.0))) * -2.25) / float64((v6 - v6))) * -2.25))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
