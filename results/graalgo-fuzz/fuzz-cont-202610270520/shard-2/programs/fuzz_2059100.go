package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32 }
func genFn0(p0 uint) uint {
    for fi5 := 0; fi5 < 15; fi5++ {
        if true {
            var v6 uint = (p0 & p0)
            v7 := []int{fi5, fi5}
            fmt.Println(p0, p0)
            _ = v6
            _ = v7
        } else {
            for fi8 := 0; fi8 < 12; fi8++ {
                p0 -= p0
                _ = fi8
            }
        }
        var v9 int64 = -9223372036854775808
        _ = fi5
        _ = v9
    }
    for fi10 := 0; fi10 < 15; fi10++ {
        p0 -= 42
        var v11 uint = (^(p0))
        if ((1 <= -1) && (100 <= -128)) {
            p0 -= v11
            v11 = p0
        } else {
            switch fi10 {
            case 1:
                if (0 <= 42) {
                    v12 := []int{(fi10 >> (((p0 / ((v11) | 1))) & 63)), int(v11), fi10}
                    _ = v12
                } else {
                    switch fi10 {
                    case -9223372036854775808:
                        var v13 bool = true
                        _ = v13
                    case -813:
                        v11 = p0
                    }
                    var v14 uint8 = 127
                    _ = v14
                }
            case -2147483648:
                if (false || true) {
                    v15 := []int{(fi10 >> (((v11 - v11)) & 63)), fi10}
                    v11 = p0
                    p0 = p0
                    _ = v15
                } else {
                    fmt.Println(fi10)
                }
            default:
                v16 := []int{fi10, 9223372036854775807}
                _ = v16
            }
        }
        _ = fi10
        _ = v11
    }
    fmt.Println(p0)
    for fi17 := 0; fi17 < 3; fi17++ {
        var v18 int32 = 0
        p0 = (^(p0))
        var v19 int64 = int64(p0)
        _ = fi17
        _ = v18
        _ = v19
    }
    _ = p0
    return (p0 / ((p0) | 1))
}

func main() {
    v1 := []uint32{}
    v2 := make(map[uint32]uint32)
    var v3 int16 = 1000
    v4 := make(map[uint32]uint32)
    for fi20 := 0; fi20 < 2; fi20++ {
        if (fi20 == -1) {
            v21 := GenS1{F0: 32768}
            for fi22 := 0; fi22 < 3; fi22++ {
                v3 += v3
                v3 += -32768
                _ = fi22
            }
            v3 += v3
            _ = v21
        }
        v3 = (v3 + v3)
        v3 = v3
        _ = fi20
    }
    delete(v2, 65535)
    switch 1 {
    case 0:
        var v23 bool = false
        _ = v23
    case 9223372036854775807:
        v24 := GenS1{F0: -2147483648}
        _ = v24
    }
    var v25 float32 = ((float32(0.1) * float32(0.0)) / float32(float64(-2.25)))
    v3 -= (v3 * (v3 | 991))
    var v26 int16 = 704
    v27 := GenS1{F0: 32767}
    var v28 int = 0
    v29 := GenS1{F0: (v28 ^ v28)}
    var v30 uint16 = 2
    fmt.Println(v3)
    if ((0 == 255) || (!("a" <= ("a" + (("" + "") + "a"))))) {
        v4[42] = 65535
    }
    fmt.Println(v28)
    v31 := GenS2{F0: 710}
    delete(v2, 0)
    var v32 int16 = v26
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
}
