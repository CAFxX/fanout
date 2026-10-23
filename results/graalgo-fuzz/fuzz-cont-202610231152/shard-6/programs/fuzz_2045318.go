package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 bool }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint32; F1 uint8 }
func genFn0(p0 uint64, p1 int16) uint8 {
    v4 := GenS1{F0: 42, F1: 255, F2: true}
    if false {
        p1 -= p1
        p1 = (p1 << ((66) & 63))
    }
    _ = p0
    _ = p1
    _ = v4
    return 127
}
func genFn1(p0 int64) uint16 {
    var v5 float64 = (0.1 * -2.25)
    fmt.Println(p0, p0)
    _ = p0
    _ = v5
    return 695
}

func main() {
    var v1 float32 = -2.25
    var v2 uint16 = 32767
    v3 := make(map[uint32]bool)
    v1 = ((float32(float64(float32(1.5))) + v1) - v1)
    v3[0] = (float64((v1 + v1)) == 1e10)
    if false {
        v6 := []int{2147483647, 128}
        _ = v6
    }
    v2 += v2
    var v7 int = 32768
    for fi8 := 0; fi8 < 14; fi8++ {
        v9 := GenS2{F0: 0, F1: (!(!true))}
        for fi10 := 0; fi10 < 2; fi10++ {
            v1 = (v1 / (v1 / (((v1 / v1) + v1) - 1.5)))
            v11 := v3[2]
            v12 := GenS3{F0: 255}
            _ = fi10
            _ = v11
            _ = v12
        }
        var v13 uint8 = 0
        _ = fi8
        _ = v13
        _ = v9
    }
    v1 = v1
    for fi14 := 0; fi14 < 15; fi14++ {
        v15 := GenS4{F0: 4294967295, F1: 1}
        var v16 uint16 = v2
        _ = fi14
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 3; fi17++ {
        v18 := v3[609]
        v2 = (v2 * v2)
        var v19 uint8 = 111
        _ = fi17
        _ = v18
        _ = v19
    }
    var v20 float32 = ((1.5 * v1) + (v1 - (v1 * v1)))
    v21 := GenS4{F0: 2, F1: 36}
    v22 := GenS4{F0: 4294967295, F1: 11}
    var v23 int = v7
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v7
}
