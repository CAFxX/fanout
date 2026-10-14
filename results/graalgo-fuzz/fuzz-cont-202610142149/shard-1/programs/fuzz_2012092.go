package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 uint8; F1 uint16; F2 int32 }
func genFn0(p0 int16, p1 int8, p2 int32) bool {
    v4 := GenS1{F0: 233}
    if true {
        p1 += (p1 * (p1 << ((2) & 63)))
    } else {
        var v5 float32 = 0.1
        _ = v5
    }
    p2 += p2
    p0 -= -32768
    v6 := GenS1{F0: 408}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v6
    return false
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 bool = (false || ((-2.25 / 3.14159) >= 1e10))
    var v3 int = -566
    fmt.Println(v2)
    for fi7 := 0; fi7 < 5; fi7++ {
        delete(v1, 364)
        var v8 uint32 = 0
        _ = fi7
        _ = v8
    }
    var v9 uint32 = 4294967295
    var v10 float32 = (-2.25 - 0.1)
    for fi11 := 0; fi11 < 9; fi11++ {
        v12 := v1[(v9 / (((v9 | 2147483647)) | 1))]
        var v13 bool = (v2 || v12)
        _ = fi11
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        v3 -= (^(fi14))
        if (v2 && ((1 > 2147483647) && (false || (0 >= 275)))) {
            var v15 float32 = v10
            v16 := GenS2{F0: 42, F1: -1}
            _ = v15
            _ = v16
        } else {
            v17 := GenS3{F0: 191, F1: 65535, F2: -1}
            _ = v17
        }
        _ = fi14
    }
    v18 := GenS3{F0: 193, F1: 32767, F2: 2147483647}
    v2 = v2
    v19 := GenS1{F0: 32767}
    v20 := GenS1{F0: 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v9
}
