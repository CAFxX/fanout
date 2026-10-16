package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32; F1 uint64 }

func main() {
    v1 := make(map[uint32]string)
    var v2 bool = false
    v3 := make(map[uint32]int)
    var v4 float64 = (float64(3.14159) + (3.14159 - ((((0.0 / float64(float32((3.14159 + 0.0)))) * -2.25) - 1e10) * (float64(0.1) - 0.0))))
    var v5 int = 9223372036854775807
    v6 := []int{}
    var v7 int32 = -812
    if v2 {
        delete(v1, 2)
        for fi8 := 0; fi8 < 8; fi8++ {
            var v9 float64 = ((v4 - 3.14159) / 1.5)
            var v10 string = "\n\t"
            _ = fi8
            _ = v10
            _ = v9
        }
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v2 = (false && true)
    if false {
        v4 -= (0.1 - 1e10)
    } else {
        for fi11 := 0; fi11 < 7; fi11++ {
            v12 := GenS1{F0: (v7 - -2147483648)}
            var v13 uint16 = 2
            _ = fi11
            _ = v12
            _ = v13
        }
        for fi14 := 0; fi14 < 3; fi14++ {
            for fi15 := 0; fi15 < 14; fi15++ {
                delete(v3, 679)
                var v16 uint32 = 0
                if (45 <= (v16 & 311)) {
                    v17 := GenS1{F0: v7}
                    v18 := GenS1{F0: v7}
                    var v19 uint8 = 255
                    _ = v17
                    _ = v18
                    _ = v19
                }
                _ = fi15
                _ = v16
            }
            _ = fi14
        }
    }
    v4 = (((0.0 - 3.14159) * v4) - 0.0)
    if (v2 && v2) {
        var v20 uint32 = 2147483647
        _ = v20
    }
    v5 -= v5
    if (!(v2 || false)) {
        v2 = v2
        delete(v3, 4294967295)
    } else {
        if len(v6) > 0 {
            v6[0] = -2147483648
        }
        v21 := GenS2{F0: true}
        _ = v21
    }
    for fi22 := 0; fi22 < 15; fi22++ {
        v23 := GenS1{F0: (v7 % (((v7 % ((int32(v5)) | 1))) | 1))}
        _ = fi22
        _ = v23
    }
    v3[0] = (v5 / ((v5) | 1))
    var v24 int32 = (-(v7))
    v24 -= (v24 % ((v7) | 1))
    fmt.Println(v4, v24, v2)
    v7 += v7
    for fi25 := 0; fi25 < 2; fi25++ {
        var v26 uint32 = 4294967295
        _ = len(v1)
        v27 := GenS3{F0: v26, F1: 232}
        _ = fi25
        _ = v26
        _ = v27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
