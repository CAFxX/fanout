package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32; F1 bool }
func genFn0(p0 uint16, p1 int16, p2 int16) int8 {
    var v4 int8 = 122
    p2 = p1
    for fi5 := 0; fi5 < 3; fi5++ {
        var v6 uint8 = 128
        _ = fi5
        _ = v6
    }
    v7 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v7
    return (v4 ^ -128)
}
func genFn1(p0 uint32) uint64 {
    v8 := GenS1{F0: (p0 + 53), F1: 32767, F2: 9223372036854775807}
    var v9 int32 = -1
    if true {
        for fi10 := 0; fi10 < 8; fi10++ {
            v11 := []int{(fi10 % ((fi10) | 1))}
            var v12 int16 = 42
            _ = len(v11)
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    fmt.Println(v9, v9)
    _ = p0
    _ = v8
    _ = v9
    return 868
}

func main() {
    var v1 int = 2147483647
    var v2 uint32 = 0
    v3 := make(map[string]int)
    v1 = (v1 / (((v1 + (v1 & (v1 >> ((465) & 63))))) | 1))
    var v13 float64 = (-2.25 + (3.14159 * (float64((((((3.14159 - float32(-2.25)) + 3.14159) + 0.1) * float32(0.0)) / 1.0)) + 0.1)))
    v14 := GenS1{F0: v2, F1: 1, F2: 376}
    v2 = (v2 & 2147483648)
    fmt.Println(v13)
    if true {
        var v15 uint16 = 0
        _ = v15
    } else {
        for fi16 := 0; fi16 < 8; fi16++ {
            var v17 float32 = (3.14159 - 0.0)
            _ = fi16
            _ = v17
        }
        var v18 int = v1
        _ = v18
    }
    var v19 string = ""
    for fi20 := 0; fi20 < 8; fi20++ {
        v13 = (v13 + ((float64(float32(v13)) + -2.25) - v13))
        v21 := GenS2{F0: true}
        v1 -= 4294967295
        _ = fi20
        _ = v21
    }
    if ((-1 > -9223372036854775808) && (((false && (v1 != v1)) && true) && true)) {
        switch v1 {
        case 141:
            for fi22 := 0; fi22 < 14; fi22++ {
                for fi23 := 0; fi23 < 11; fi23++ {
                    switch (-(fi22)) {
                    case -1:
                        v13 -= ((v13 - (float64((0.0 + (0.0 - 0.1))) * v13)) * float64(1e10))
                    case 4294967295:
                        v19 = v19
                        var v24 int32 = -2147483648
                        _ = v24
                    }
                    _ = fi23
                }
                _ = fi22
            }
        default:
            v25 := GenS3{F0: -1, F1: (false || (false && (127 <= -79)))}
            _ = v25
        }
        v13 -= ((v13 / ((float64(float32(v13)) + 0.1) + ((v13 + v13) + v13))) / 1.5)
        v26 := GenS1{F0: v2, F1: 255, F2: 41}
        _ = v26
    } else {
        switch v1 {
        case 907:
            v2 = 913
        case -931:
            for fi27 := 0; fi27 < 3; fi27++ {
                v28 := []int{127, v1}
                var v29 uint16 = 1
                _ = fi27
                _ = v28
                _ = v29
            }
        default:
            switch v1 {
            case 32767:
                delete(v3, v19)
            }
        }
        if (!false) {
            var v30 int = 4294967295
            _ = v30
        } else {
            var v31 uint = 42
            if false {
                var v32 uint16 = 1
                v13 -= 3.14159
                _ = v32
            } else {
                v33 := GenS1{F0: v2, F1: 42, F2: 756}
                _ = v33
            }
            _ = v31
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v3
}
