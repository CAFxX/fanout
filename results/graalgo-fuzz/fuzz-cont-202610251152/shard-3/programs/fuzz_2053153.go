package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 bool; F1 int16; F2 uint }
type GenS3 struct { F0 uint32; F1 int32 }
func genFn0(p0 uint8) int {
    p0 = p0
    if (1 >= 874) {
        if (p0 > (^(p0))) {
            if true {
                p0 = 2
            }
            v6 := GenS1{F0: p0, F1: (p0 / ((p0) | 1))}
            _ = v6
        } else {
            if (-1 > 0) {
                if ((42 > 883) || (false || true)) {
                    fmt.Println(p0)
                    if false {
                        var v7 uint16 = 65535
                        _ = v7
                    }
                } else {
                    fmt.Println(p0)
                }
                var v8 bool = (!true)
                _ = v8
            }
            var v9 int = 32767
            p0 += (p0 + 1)
            _ = v9
        }
        for fi10 := 0; fi10 < 12; fi10++ {
            v11 := []int{(fi10 * (fi10 + (fi10 | (fi10 & -535)))), (fi10 & fi10), (fi10 | -93)}
            _ = fi10
            _ = v11
        }
    }
    for fi12 := 0; fi12 < 3; fi12++ {
        p0 -= (^(p0))
        v13 := GenS2{F0: false, F1: 1, F2: 65535}
        _ = fi12
        _ = v13
    }
    _ = p0
    return -538
}
func genFn1(p0 int16, p1 uint8) uint8 {
    v14 := []int{}
    if len(v14) > 0 {
        v14[0] = 128
    }
    v14 = append(v14, 4294967295)
    v15 := GenS3{F0: 4294967295, F1: 2147483647}
    for fi16 := 0; fi16 < 3; fi16++ {
        for fi17 := 0; fi17 < 15; fi17++ {
            var v18 float32 = -2.25
            _ = fi17
            _ = v18
        }
        p0 = (p0 | p0)
        p0 = p0
        _ = fi16
    }
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    return (^(p1))
}

func main() {
    v1 := []uint32{891}
    var v2 uint8 = 127
    var v3 bool = (3.14159 == 3.14159)
    var v4 uint16 = 1
    v5 := []int{-614, 4294967295, 9223372036854775807}
    for fi19 := 0; fi19 < 18; fi19++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi19
    }
    if (127 == 127) {
        for fi20 := 0; fi20 < 14; fi20++ {
            var v21 int16 = 42
            _ = fi20
            _ = v21
        }
        var v22 float64 = (1e10 / ((3.14159 / (0.0 + (3.14159 - 1.5))) + -2.25))
        _ = v22
    }
    v23 := GenS2{F0: ((^(v4)) != (v4 & (v4 | v4))), F1: -1, F2: 42}
    for fi24 := 0; fi24 < 9; fi24++ {
        var v25 uint32 = 722
        _ = fi24
        _ = v25
    }
    fmt.Println(v2, v4)
    for fi26 := 0; fi26 < 8; fi26++ {
        v4 -= 0
        fmt.Println(v4, fi26, v2)
        _ = fi26
    }
    for fi27 := 0; fi27 < 8; fi27++ {
        switch "hello" {
        case "hello":
            fmt.Println(v3, v3)
        }
        switch fi27 {
        case 0:
            var v28 int = (fi27 * (^(fi27)))
            _ = v28
        }
        _ = fi27
    }
    v5 = append(v5, 453)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
