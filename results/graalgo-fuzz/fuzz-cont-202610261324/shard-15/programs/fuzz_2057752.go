package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 int16 }
func genFn0(p0 int32) uint32 {
    if (42 > 631) {
        p0 = p0
    }
    p0 = (p0 % ((0) | 1))
    v7 := []int{128}
    if false {
        for fi8 := 0; fi8 < 7; fi8++ {
            p0 = (p0 / (((^(p0))) | 1))
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            _ = fi8
        }
        for fi9 := 0; fi9 < 9; fi9++ {
            switch (^(fi9)) {
            case 4294967295:
                for fi10 := 0; fi10 < 7; fi10++ {
                    var v11 int8 = -128
                    if len(v7) > 0 {
                        v7[0] = 32767
                    }
                    _ = fi10
                    _ = v11
                }
            }
            _ = fi9
        }
    }
    _ = p0
    _ = v7
    return 0
}

func main() {
    v1 := []uint32{65535}
    var v2 int16 = 1000
    v3 := []int{32767, 127, -455, 2147483647}
    var v4 uint16 = 255
    var v5 int = 128
    v6 := []int{v5}
    for fi12 := 0; fi12 < 1; fi12++ {
        v13 := GenS1{F0: -1}
        _ = fi12
        _ = v13
    }
    fmt.Println(v4, v2, v5)
    var v14 string = "a"
    fmt.Println(v5, v2)
    var v15 float32 = float32(0.1)
    v3 = append(v3, -1)
    v4 = v4
    var v16 int32 = -664
    v2 += v2
    v17 := GenS2{F0: 255, F1: (v2 & (-(v2)))}
    var v18 float32 = ((-2.25 * ((v15 - v15) - 0.0)) - v15)
    for fi19 := 0; fi19 < 15; fi19++ {
        for fi20 := 0; fi20 < 1; fi20++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v16 -= (v16 % ((v16) | 1))
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = fi20
        }
        v21 := GenS1{F0: (-(v16))}
        var v22 uint = 0
        _ = fi19
        _ = v21
        _ = v22
    }
    v5 -= 32767
    v2 += (v2 - v2)
    var v23 float64 = 0.1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
