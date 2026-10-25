package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint; F1 uint32; F2 int8 }
func genFn0(p0 int16) uint8 {
    switch -9223372036854775808 {
    case 220:
        var v4 float32 = -2.25
        _ = v4
    case 0:
        var v5 uint16 = 1
        _ = v5
    default:
        fmt.Println(p0)
    }
    for fi6 := 0; fi6 < 7; fi6++ {
        var v7 uint = 255
        if ((fi6 << (((v7 + 2)) & 63)) <= (^(fi6))) {
            p0 = (-(p0))
        } else {
            fmt.Println(fi6)
            if (627 >= 0) {
                v8 := GenS1{F0: -786}
                _ = v8
            } else {
                var v9 float64 = 1.5
                _ = v9
            }
        }
        _ = fi6
        _ = v7
    }
    _ = p0
    return 2
}
func genFn1(p0 string, p1 int32) uint {
    p0 = (("foo bar" + p0) + "\n\t")
    for fi10 := 0; fi10 < 17; fi10++ {
        for fi11 := 0; fi11 < 7; fi11++ {
            v12 := []int{fi10}
            var v13 float32 = (1.5 * float32((-2.25 / 1.0)))
            v14 := GenS1{F0: p1}
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
        _ = fi10
    }
    for fi15 := 0; fi15 < 18; fi15++ {
        v16 := []int{(fi15 / ((fi15) | 1))}
        _ = fi15
        _ = v16
    }
    switch 1 {
    case -214:
        v17 := GenS1{F0: 1}
        _ = v17
    default:
        var v18 int16 = 1
        _ = v18
    }
    p1 = (-(p1))
    _ = p0
    _ = p1
    return 0
}

func main() {
    var v1 uint32 = 2
    var v2 uint32 = v1
    v3 := []uint32{v2, (v1 & v1), (v2 + (v2 / (((v1 | (v2 + 1))) | 1)))}
    v2 -= 255
    var v19 int8 = -1
    switch 9223372036854775807 {
    case 32767:
        v20 := GenS1{F0: -2147483648}
        _ = v20
    case 128:
        var v21 bool = (!true)
        _ = v21
    }
    v2 = 295
    var v22 uint16 = 85
    fmt.Println(v1, v22, v2)
    v2 -= (^(v1))
    var v23 uint32 = (v2 - (v2 * v2))
    var v24 float64 = float64(((0.1 / ((-2.25 / 3.14159) * float32(3.14159))) * float32((0.1 * -2.25))))
    if true {
        fmt.Println(v24)
        if false {
            v2 = (^(v2))
        } else {
            v25 := GenS2{F0: 2}
            _ = v25
        }
    }
    var v26 uint64 = 2
    v24 -= 0.1
    v27 := GenS3{F0: 0, F1: 646, F2: (v19 * 0)}
    v3 = append(v3, v1)
    var v28 uint8 = 2
    var v29 uint32 = v23
    if true {
        var v30 float32 = (1e10 / 1.0)
        for fi31 := 0; fi31 < 8; fi31++ {
            v19 -= int8(v28)
            _ = fi31
        }
        _ = v30
    }
    v32 := GenS2{F0: 230}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v32
}
