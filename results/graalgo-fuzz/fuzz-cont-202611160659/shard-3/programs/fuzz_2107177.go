package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 int32 = -2147483648
    var v2 int = 4294967295
    var v3 float64 = (float64(-2.25) * 0.1)
    var v4 int8 = 42
    var v5 float32 = 3.14159
    if true {
        v6 := []int{(v2 & v2), v2, (-(v2)), v2}
        v7 := GenS1{F0: 0, F1: (v4 - v4)}
        _ = v6
        _ = v7
    } else {
        v8 := []int{v2}
        var v9 string = ""
        v3 = 1e10
        _ = v8
        _ = v9
    }
    if false {
        for fi10 := 0; fi10 < 6; fi10++ {
            var v11 uint32 = 1
            _ = fi10
            _ = v11
        }
        var v12 int = v2
        if (true && ((255 > 4294967295) || true)) {
            if (4294967295 >= 65535) {
                v5 = float32(((1e10 / v3) * ((float64(v5) / v3) * (1.5 * float64(float32(float64(v5)))))))
                v13 := GenS1{F0: 1, F1: 1}
                _ = v13
            }
            v12 = (v12 * (v2 + v2))
            var v14 uint8 = 166
            _ = v14
        }
        _ = v12
    } else {
        switch "hello" {
        case "a":
            v15 := GenS1{F0: 19, F1: (v4 & v4)}
            _ = v15
        case "foo bar":
            fmt.Println(v4, v3, v4)
        case "":
            for fi16 := 0; fi16 < 15; fi16++ {
                if (9223372036854775807 == 966) {
                    v1 = v1
                    var v17 float32 = 0.0
                    v5 += v5
                    _ = v17
                } else {
                    v18 := GenS2{F0: 1}
                    _ = v18
                }
                v5 = v5
                for fi19 := 0; fi19 < 20; fi19++ {
                    v2 = v2
                    v20 := []int{(fi16 - -65), (fi16 / ((487) | 1)), 1, 0}
                    _ = fi19
                    _ = v20
                }
                _ = fi16
            }
        default:
            var v21 uint64 = 0
            _ = v21
        }
        v22 := []int{v2, -689}
        _ = v22
    }
    v23 := GenS1{F0: 127, F1: (^(v4))}
    v24 := []int{(^(v2)), (^(v2)), (^(v2))}
    if len(v24) > 0 {
        v24[0] = (v2 ^ int(v1))
    }
    var v25 int64 = -1
    fmt.Println(v4)
    v2 = -2147483648
    var v26 uint = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
