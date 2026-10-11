package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 int16 }
type GenS2 struct { F0 int8; F1 int16; F2 int16 }
func genFn0(p0 int8) bool {
    p0 = (p0 - 57)
    v7 := []int{}
    _ = p0
    _ = v7
    return true
}

func main() {
    var v1 float32 = ((1e10 / 1.0) * -2.25)
    var v2 int16 = -423
    var v3 int8 = 1
    var v4 uint64 = 671
    v5 := make(map[string]string)
    var v6 uint = 215
    for fi8 := 0; fi8 < 18; fi8++ {
        v9 := []int{fi8, (-(fi8)), fi8}
        v1 -= float32((0.0 * (float64(v1) / float64((((float32((1e10 / 1.5)) + 1e10) * ((v1 / (3.14159 * 3.14159)) + v1)) + v1)))))
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 16; fi10++ {
        v11 := GenS1{F0: (v2 | v2), F1: v4, F2: v2}
        v12 := GenS1{F0: 42, F1: (^(v4)), F2: v2}
        _ = fi10
        _ = v11
        _ = v12
    }
    v13 := GenS2{F0: 119, F1: (v2 & (v2 << ((v6) & 63))), F2: v2}
    if ((v6 ^ 4294967295) < v6) {
        if (true || (1 <= 0)) {
            fmt.Println(v2)
        } else {
            v14 := GenS1{F0: (v2 | v2), F1: 980, F2: v2}
            _ = v14
        }
        if false {
            fmt.Println(v6, v3, v3)
            var v15 uint8 = 38
            _ = v15
        }
        switch -2147483648 {
        case 9223372036854775807:
            v16 := GenS2{F0: v3, F1: v2, F2: (v2 * v2)}
            _ = v16
        case 128:
            v17 := GenS1{F0: v2, F1: 42, F2: 32767}
            _ = v17
        }
    }
    v4 = (v4 | (v4 + v4))
    v18 := GenS2{F0: (v3 ^ v3), F1: (v2 ^ 1), F2: int16(v6)}
    var v19 int = 2147483647
    delete(v5, "foo bar")
    v20 := GenS1{F0: v2, F1: 42, F2: v2}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
