package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 int32 }

func main() {
    var v1 uint = 473
    var v2 int8 = -128
    var v3 float64 = -2.25
    var v4 int16 = -434
    v5 := make(map[int]bool)
    var v6 bool = false
    fmt.Println(v2)
    for fi7 := 0; fi7 < 17; fi7++ {
        _ = len(v5)
        switch (fi7 << ((42) & 63)) {
        case 4294967295:
            var v8 int16 = 1
            _ = v8
        case 32768:
            fmt.Println(v2, fi7)
        default:
            for fi9 := 0; fi9 < 20; fi9++ {
                if true {
                    fmt.Println(v2, v2, fi7)
                    if ((-9223372036854775808 != 0) || v6) {
                        v4 = -217
                        v10 := GenS1{F0: 0, F1: (v4 * (-(v4)))}
                        var v11 float64 = v3
                        var v12 int32 = 2147483647
                        if (true || ((v11 == v3) || v6)) {
                            fmt.Println(fi7)
                            v4 = v4
                        }
                        var v13 int16 = 1
                        var v14 uint8 = 171
                        _ = v10
                        _ = v11
                        _ = v12
                        _ = v13
                        _ = v14
                    } else {
                        v4 = v4
                        var v15 uint64 = 4294967295
                        _ = v15
                    }
                    v16 := GenS2{F0: 0, F1: 2}
                    _ = v16
                } else {
                    v6 = (42 != 2147483647)
                    v17 := GenS3{F0: -813}
                    _ = v17
                }
                v2 += v2
                var v18 int = fi9
                _ = fi9
                _ = v18
            }
        }
        var v19 string = ("a" + "\n\t")
        _ = fi7
        _ = v19
    }
    var v20 uint16 = 2
    var v21 float64 = v3
    v2 = (v2 % ((v2) | 1))
    var v22 string = ""
    delete(v5, -9)
    var v23 float64 = (v3 + (((v21 + (v21 + v21)) / v3) - (((float64((3.14159 - float32(((v21 / (1.5 / v21)) + 3.14159)))) / 1.0) * (v3 + (float64((((-2.25 * -2.25) - 1e10) - (3.14159 / float32(v3)))) + (0.0 / (1.5 + 0.1))))) * -2.25)))
    if v6 {
        fmt.Println(v20)
    }
    v20 += 26
    v21 += (1.5 + v23)
    var v24 float64 = float64(0.1)
    var v25 uint = (v1 ^ 218)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
