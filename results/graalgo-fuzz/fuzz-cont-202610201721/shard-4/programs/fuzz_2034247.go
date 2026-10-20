package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 uint16 }
type GenS2 struct { F0 uint16; F1 int8; F2 uint32 }

func main() {
    var v1 uint64 = 814
    v2 := make(map[int]bool)
    var v3 bool = true
    var v4 uint32 = 2147483648
    v5 := make(map[uint32]bool)
    var v6 uint8 = 255
    if v3 {
        for fi7 := 0; fi7 < 19; fi7++ {
            v6 -= 128
            _ = fi7
        }
        if (-474 > -34) {
            var v8 uint64 = v1
            _ = v8
        }
    }
    if false {
        v9 := GenS1{F0: 695, F1: 42, F2: 716}
        _ = v9
    }
    v1 -= (v1 + (v1 + (v1 - (v1 / (((v1 * (v1 ^ (v1 - uint64(v6))))) | 1)))))
    v4 = v4
    var v10 string = "\n\t"
    var v11 int = 1
    var v12 int32 = -63
    var v13 int = 0
    fmt.Println(v13, v6)
    if ((true || (0.1 > 3.14159)) && ((!(42 <= 1)) && v3)) {
        fmt.Println(v13, v4)
        v14 := GenS1{F0: 530, F1: 42, F2: 32767}
        _ = v14
    } else {
        v15 := GenS2{F0: 0, F1: int8(v13), F2: (^(v4))}
        v3 = v3
        _ = v15
    }
    v16 := GenS2{F0: 255, F1: -128, F2: 354}
    _ = len(v2)
    for fi17 := 0; fi17 < 18; fi17++ {
        v12 = -692
        v13 = (^(v11))
        switch v13 {
        case 128:
            v6 -= v6
        case -999:
            v1 = (v1 / ((v1) | 1))
        case -466:
            for fi18 := 0; fi18 < 11; fi18++ {
                if ((v4 / (((v4 & (^(v4)))) | 1)) <= 145) {
                    v19 := v5[v4]
                    v1 = v1
                    var v20 float64 = 1e10
                    _ = v19
                    _ = v20
                }
                _ = fi18
            }
        }
        _ = fi17
    }
    v21 := []int{}
    var v22 uint16 = 0
    v3 = (1.5 <= ((float32((1e10 - float64((-2.25 * (((1.5 / -2.25) / 0.1) + -2.25))))) / float32(((3.14159 - float64((0.0 / ((1.5 + 3.14159) + ((3.14159 - -2.25) / 1.5))))) * 0.1))) * 3.14159))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
