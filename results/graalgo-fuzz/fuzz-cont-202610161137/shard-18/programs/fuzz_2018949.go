package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8; F1 int32 }
func genFn0(p0 int) int32 {
    v5 := []int{792, (p0 % ((4294967295) | 1))}
    v6 := GenS1{F0: 4294967295}
    p0 = p0
    _ = p0
    _ = v5
    _ = v6
    return 834
}
func genFn1(p0 int64, p1 int, p2 int32) uint32 {
    var v7 int64 = p0
    var v8 int16 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 697
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 bool = false
    v3 := make(map[int]int)
    var v4 int32 = 0
    v9 := GenS1{F0: (v1 ^ v1)}
    v10 := GenS1{F0: (v1 / ((v1) | 1))}
    var v11 uint64 = v1
    var v12 float32 = ((-2.25 / 3.14159) + 0.1)
    v13 := GenS2{F0: -446}
    switch 742 {
    case -2147483648:
        v11 = 42
    default:
        if false {
            for fi14 := 0; fi14 < 17; fi14++ {
                var v15 int8 = 127
                v16 := GenS1{F0: (v1 + 42)}
                _ = fi14
                _ = v15
                _ = v16
            }
        } else {
            v4 = 328
            var v17 float32 = v12
            var v18 float32 = (float32(-2.25) - (float32((float64((0.0 - (1.5 / (v12 - 1.5)))) / 1e10)) * ((0.0 - -2.25) * float32(0.0))))
            _ = v17
            _ = v18
        }
    }
    v19 := v3[146]
    if (v2 && ((float32((3.14159 * (0.0 * float64((((v12 * v12) * ((v12 + v12) - v12)) / v12))))) * ((float32(3.14159) + 0.1) + v12)) <= (v12 * float32(-2.25)))) {
        for fi20 := 0; fi20 < 19; fi20++ {
            delete(v3, 127)
            v3[-1] = 1
            fmt.Println(v1, v12)
            _ = fi20
        }
        fmt.Println(v4, v1, v19)
    } else {
        v12 -= 0.0
        for fi21 := 0; fi21 < 11; fi21++ {
            switch 32768 {
            case 32767:
                for fi22 := 0; fi22 < 6; fi22++ {
                    var v23 bool = (2 == genFn1(9223372036854775807, int(v1), v4))
                    _ = fi22
                    _ = v23
                }
            case 128:
                var v24 string = ("hello" + "foo bar")
                _ = v24
            case -892:
                v25 := []int{(^(v19)), v19, fi21}
                _ = v25
            }
            switch fi21 {
            case -1:
                v26 := GenS3{F0: 91, F1: genFn0((-(fi21)))}
                _ = v26
            case -2147483648:
                v27 := []int{127, v19}
                _ = v27
            default:
                var v28 uint32 = 0
                _ = v28
            }
            _ = fi21
        }
        v29 := []int{v19, (v19 + (v19 / (((v19 / (((v19 << ((622) & 63))) | 1))) | 1))), (v19 * (v19 << ((4294967295) & 63)))}
        _ = v29
    }
    var v30 int16 = 1
    v31 := GenS2{F0: (v30 + (v30 & v30))}
    if true {
        v19 -= (v19 << ((18446744073709551615) & 63))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v9
}
