package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 int64, p1 uint16) uint8 {
    v6 := []int{2147483647}
    if false {
        p1 -= p1
    }
    _ = p0
    _ = p1
    _ = v6
    return 0
}
func genFn1(p0 int16, p1 uint64, p2 int8) uint {
    p2 += p2
    p0 = p0
    fmt.Println(p0)
    var v7 uint = 163
    v8 := GenS1{F0: -190}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return v7
}

func main() {
    var v1 uint8 = 255
    v2 := make(map[int]bool)
    var v3 uint = 2
    var v4 int64 = -495
    var v5 uint32 = 0
    var v9 uint64 = 1
    v9 = (v9 + (v9 & v9))
    if false {
        v10 := GenS1{F0: 221}
        fmt.Println(v5)
        switch 4294967295 {
        case 32768:
            var v11 float64 = 3.14159
            _ = v11
        default:
            v12 := []int{}
            _ = v12
        }
        _ = v10
    } else {
        delete(v2, 128)
        v1 = v1
    }
    v13 := []int{34, 128}
    fmt.Println(v4)
    var v14 int16 = 0
    var v15 uint8 = genFn0((v4 >> (((v3 >> (((v3 | uint(v5))) & 63))) & 63)), uint16(v3))
    if (true || true) {
        v16 := GenS1{F0: 42}
        _ = v16
    } else {
        switch 484 {
        case 127:
            for fi17 := 0; fi17 < 20; fi17++ {
                for fi18 := 0; fi18 < 17; fi18++ {
                    v15 = uint8(v3)
                    _ = fi18
                }
                _ = fi17
            }
        case 2147483647:
            v19 := GenS1{F0: 1}
            _ = v19
        }
    }
    v2[32768] = (!("foo bar" == ""))
    for fi20 := 0; fi20 < 18; fi20++ {
        if ((((1e10 / ((1.5 - ((-2.25 / 1.5) + (3.14159 / float64(0.1)))) * (float64(3.14159) + (float64(1e10) * 0.0)))) + (3.14159 + ((float64(float32(0.1)) / 1.5) - (1e10 + 1e10)))) >= -2.25) || ((v14 * v14) >= (v14 - (v14 - v14)))) {
            v15 -= genFn0((v4 - (v4 + (v4 - (v4 >> ((2) & 63))))), 65535)
        }
        v5 -= v5
        v21 := GenS1{F0: 42}
        _ = fi20
        _ = v21
    }
    if false {
        v9 = (v9 ^ (v9 << ((106) & 63)))
    }
    v13 = append(v13, 127)
    v1 += 2
    v5 -= (v5 * 2147483647)
    fmt.Println(v5, v4, v5)
    v22 := GenS1{F0: 2147483647}
    fmt.Println(v9)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
