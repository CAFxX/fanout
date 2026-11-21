package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 uint8; F1 int64; F2 int }
type GenS4 struct { F0 uint8 }
func genFn0(p0 uint8, p1 uint16) uint8 {
    p0 += (p0 >> ((42) & 63))
    if false {
        switch 42 {
        case 128:
            if (355 != -50) {
                v5 := GenS1{F0: 0}
                _ = v5
            }
        case -1:
            v6 := []int{}
            _ = v6
        }
    } else {
        var v7 uint16 = (p1 | (^(p1)))
        v8 := GenS1{F0: 171}
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: 65535}
    _ = p0
    _ = p1
    _ = v9
    return (p0 % ((p0) | 1))
}
func genFn1(p0 uint16) int32 {
    v10 := GenS2{F0: 100, F1: 128}
    var v11 int = 685
    v12 := []int{(v11 ^ v11)}
    _ = p0
    _ = v10
    _ = v11
    _ = v12
    return -981
}

func main() {
    var v1 int8 = -6
    var v2 uint8 = 127
    var v3 uint = 2
    var v4 uint16 = 255
    fmt.Println(v1)
    var v13 int = int(v2)
    v3 = (v3 ^ uint(v2))
    v2 = (v2 / ((v2) | 1))
    var v14 int16 = 1
    for fi15 := 0; fi15 < 18; fi15++ {
        v3 = (v3 >> ((255) & 63))
        var v16 uint = (v3 ^ v3)
        _ = fi15
        _ = v16
    }
    var v17 float32 = -2.25
    var v18 bool = ((v1 - (v1 - (v1 >> ((v3) & 63)))) != v1)
    for fi19 := 0; fi19 < 20; fi19++ {
        v20 := GenS3{F0: (^(v2)), F1: -1, F2: (v13 >> (((^(v3))) & 63))}
        v21 := GenS4{F0: 255}
        v17 += 0.0
        _ = fi19
        _ = v20
        _ = v21
    }
    for fi22 := 0; fi22 < 16; fi22++ {
        var v23 float64 = ((((-2.25 / 1.5) + (1e10 * 0.1)) * 1.5) - 0.0)
        v2 += v2
        v24 := GenS2{F0: (v1 ^ v1), F1: v13}
        _ = fi22
        _ = v23
        _ = v24
    }
    switch v13 {
    case 74:
        for fi25 := 0; fi25 < 11; fi25++ {
            v1 += (v1 % (((v1 ^ (^(v1)))) | 1))
            for fi26 := 0; fi26 < 13; fi26++ {
                v2 -= genFn0(v2, v4)
                _ = fi26
            }
            var v27 uint8 = uint8(v4)
            _ = fi25
            _ = v27
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
