package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 int16; F1 uint8 }

func main() {
    var v1 uint64 = 704
    var v2 uint64 = v1
    v3 := make(map[uint32]string)
    var v4 float32 = 1.5
    v5 := make(map[uint32]int)
    v6 := GenS1{F0: -9223372036854775808}
    if ((1.5 / 1.5) != ((float64(v4) - 0.1) + ((-2.25 * 3.14159) * 0.1))) {
        v7 := []int{-384, -9223372036854775808}
        _ = v7
    }
    if (!((255 != 537) || (!false))) {
        v8 := []int{}
        _ = v8
    } else {
        v9 := []int{}
        _ = v9
    }
    for fi10 := 0; fi10 < 15; fi10++ {
        if false {
            var v11 float64 = 0.1
            v2 -= v1
            _ = v11
        } else {
            v12 := []int{fi10, (-(fi10))}
            v13 := GenS1{F0: fi10}
            _ = v12
            _ = v13
        }
        v14 := []int{fi10, (-(fi10)), fi10, (fi10 + (fi10 % (((fi10 - fi10)) | 1)))}
        v4 = v4
        _ = fi10
        _ = v14
    }
    var v15 uint64 = 927
    var v16 uint64 = 0
    v17 := GenS1{F0: -2147483648}
    switch -2147483648 {
    case 4294967295:
        v16 += 0
    case 2147483647:
        if false {
            switch 427 {
            case 127:
                var v18 uint = 2
                _ = v18
            case 0:
                switch 42 {
                case -80:
                    v2 = v15
                default:
                    for fi19 := 0; fi19 < 9; fi19++ {
                        v20 := GenS2{F0: 2147483648, F1: 127}
                        _ = fi19
                        _ = v20
                    }
                }
            }
            v21 := GenS3{F0: 174, F1: 2}
            _ = v21
        }
    case 42:
        v22 := []int{463}
        _ = v22
    }
    var v23 float32 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
