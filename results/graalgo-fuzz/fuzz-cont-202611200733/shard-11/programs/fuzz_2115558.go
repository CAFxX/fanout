package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 uint8; F1 uint32; F2 uint32 }
type GenS3 struct { F0 int64 }

func main() {
    var v1 int16 = 42
    v2 := make(map[uint32]bool)
    var v3 int = -2147483648
    v4 := GenS1{F0: 607, F1: 42}
    if false {
        v3 = v3
        fmt.Println(v1)
    }
    v3 = v3
    v3 += v3
    v1 -= (v1 & 1)
    v5 := GenS2{F0: 81, F1: 42, F2: 1}
    if true {
        v3 -= v3
        v1 -= v1
        v1 = 142
    }
    if (((-2.25 - (0.0 * float64(1e10))) * 0.0) == -2.25) {
        switch 1 {
        case 9223372036854775807:
            if true {
                var v6 uint32 = 834
                v7 := GenS1{F0: 570, F1: 100}
                for fi8 := 0; fi8 < 11; fi8++ {
                    v3 = (fi8 & v3)
                    var v9 uint = 65535
                    _ = fi8
                    _ = v9
                }
                _ = v6
                _ = v7
            } else {
                var v10 string = "hello"
                for fi11 := 0; fi11 < 12; fi11++ {
                    v12 := GenS1{F0: 607, F1: 100}
                    v13 := GenS2{F0: 0, F1: 4294967295, F2: 846}
                    var v14 int8 = -1
                    _ = fi11
                    _ = v12
                    _ = v13
                    _ = v14
                }
                v1 = (-(v1))
                _ = v10
            }
        default:
            var v15 float32 = (0.1 * float32(0.0))
            _ = v15
        }
        v3 = v3
        switch 127 {
        case 1:
            fmt.Println(v3, v1)
        case 128:
            var v16 uint64 = 0
            _ = v16
        }
    } else {
        v1 -= v1
        for fi17 := 0; fi17 < 17; fi17++ {
            v18 := []int{(fi17 & 1)}
            _ = fi17
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 10; fi19++ {
        fmt.Println(fi19, fi19)
        _ = fi19
    }
    v3 = (-(v3))
    fmt.Println(v3, v3)
    var v20 int16 = v1
    v20 = 804
    v1 += 42
    if (1 <= 1) {
        v21 := GenS3{F0: -1}
        _ = v21
    } else {
        v20 = v1
    }
    v2[704] = true
    var v22 uint16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
