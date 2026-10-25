package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16; F1 int8; F2 int32 }
func genFn0() int8 {
    for fi5 := 0; fi5 < 13; fi5++ {
        v6 := GenS1{F0: 0, F1: -32768}
        var v7 uint8 = 127
        _ = fi5
        _ = v6
        _ = v7
    }
    v8 := GenS2{F0: 18446744073709551615}
    v9 := GenS3{F0: 427, F1: -79, F2: 2147483647}
    v10 := []int{32768, 127, -659}
    var v11 uint64 = 412
    _ = v10
    _ = v11
    _ = v8
    _ = v9
    return -128
}
func genFn1() bool {
    var v12 int16 = 0
    var v13 uint = 399
    _ = v12
    _ = v13
    return (709 == 219)
}

func main() {
    var v1 uint16 = 232
    v2 := make(map[uint32]bool)
    var v3 int8 = 97
    var v4 int32 = -921
    v14 := []int{-9223372036854775808, 0}
    v15 := GenS3{F0: 255, F1: (^(v3)), F2: (v4 + v4)}
    fmt.Println(v3, v4, v3)
    v16 := GenS2{F0: 2}
    var v17 uint64 = 141
    for fi18 := 0; fi18 < 7; fi18++ {
        for fi19 := 0; fi19 < 10; fi19++ {
            fmt.Println(fi19)
            v20 := GenS1{F0: 0, F1: 1}
            v21 := GenS3{F0: (^(v1)), F1: (v3 | (v3 << ((255) & 63))), F2: (^(v4))}
            _ = fi19
            _ = v20
            _ = v21
        }
        if false {
            for fi22 := 0; fi22 < 13; fi22++ {
                var v23 float32 = ((((float32(0.0) - 3.14159) * float32((3.14159 * (1e10 - 1e10)))) * 3.14159) + float32(((1e10 / 1.5) * 1e10)))
                if false {
                    v24 := v2[0]
                    _ = v24
                }
                _ = fi22
                _ = v23
            }
            switch 2147483647 {
            case 128:
                if len(v14) > 0 {
                    v14[0] = fi18
                }
            case 32768:
                v25 := v2[uint32(v17)]
                _ = v25
            default:
                var v26 int = (fi18 & (fi18 / ((206) | 1)))
                _ = v26
            }
        }
        _ = fi18
    }
    var v27 string = "\n\t"
    v28 := GenS1{F0: 0, F1: 1}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
