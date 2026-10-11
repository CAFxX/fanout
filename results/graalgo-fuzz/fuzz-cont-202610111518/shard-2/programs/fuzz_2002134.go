package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 int64; F2 string }

func main() {
    var v1 uint32 = 2147483648
    v2 := make(map[int]string)
    v3 := make(map[int]uint32)
    var v4 uint = 1
    v5 := make(map[string]int)
    for fi6 := 0; fi6 < 8; fi6++ {
        for fi7 := 0; fi7 < 18; fi7++ {
            switch (fi7 ^ fi7) {
            case 9223372036854775807:
                v8 := GenS1{F0: (v4 << ((v4) & 63))}
                _ = v8
            case -213:
                v9 := []int{(fi7 >> ((v4) & 63)), fi7}
                _ = v9
            case 0:
                v4 = v4
            }
            _ = fi7
        }
        _ = fi6
    }
    fmt.Println(v1, v4)
    v10 := GenS1{F0: v4}
    v11 := v2[1]
    for fi12 := 0; fi12 < 14; fi12++ {
        var v13 bool = (fi12 != fi12)
        if (0 != -2147483648) {
            v11 = ""
            if v13 {
                var v14 int8 = 100
                v11 = v11
                for fi15 := 0; fi15 < 20; fi15++ {
                    var v16 float64 = (-2.25 + ((((3.14159 / 3.14159) + 3.14159) + 1.5) + (((3.14159 - 1e10) / 0.1) - float64(float32(float64(0.1))))))
                    if v13 {
                        var v17 bool = v13
                        var v18 uint32 = (v1 * v1)
                        _ = v17
                        _ = v18
                    } else {
                        v16 = ((v16 / (0.1 * v16)) * v16)
                    }
                    var v19 int32 = -1
                    _ = fi15
                    _ = v16
                    _ = v19
                }
                _ = v14
            }
        } else {
            switch "a" {
            case "\n\t":
                if (((v1 >= uint32(v4)) || (9223372036854775807 == 231)) || ((v1 >> ((v4) & 63)) == (^(v1)))) {
                    v20 := GenS2{F0: 9223372036854775807, F1: -876, F2: ("a" + v11)}
                    v21 := GenS2{F0: -9223372036854775808, F1: 892, F2: "hello"}
                    v22 := v3[-531]
                    _ = v20
                    _ = v21
                    _ = v22
                }
            case "foo bar":
                for fi23 := 0; fi23 < 17; fi23++ {
                    var v24 uint64 = 2
                    v1 = (v1 + 0)
                    v1 -= v1
                    _ = fi23
                    _ = v24
                }
            default:
                v25 := []int{fi12, (fi12 % ((127) | 1)), fi12, fi12}
                _ = v25
            }
            v26 := GenS1{F0: v4}
            switch v11 {
            case "\n\t":
                var v27 int = int(v4)
                _ = v27
            }
            _ = v26
        }
        _ = len(v5)
        _ = fi12
        _ = v13
    }
    var v28 int16 = -1
    var v29 uint = (v4 & uint(v1))
    var v30 float32 = 0.0
    var v31 uint16 = 42
    var v32 uint8 = 1
    var v33 int = 1
    switch (v33 - v33) {
    case -334:
        v2[v33] = "\n\t"
    default:
        var v34 float64 = 0.0
        _ = v34
    }
    v32 = 42
    v35 := []int{v33, 549, (v33 % (((v33 | v33)) | 1)), -2147483648}
    if false {
        fmt.Println(v33)
        fmt.Println(v1)
        fmt.Println(v28, v30)
    } else {
        var v36 uint32 = (^(v1))
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v35
    _ = v4
    _ = v5
}
