package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint8; F1 string }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int16; F1 uint }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]string)
    var v3 int64 = -9223372036854775808
    v4 := make(map[string]uint32)
    var v5 int8 = -128
    v6 := GenS1{F0: 1, F1: (v5 & 0)}
    if true {
        for fi7 := 0; fi7 < 17; fi7++ {
            var v8 int32 = 42
            if true {
                for fi9 := 0; fi9 < 17; fi9++ {
                    v10 := GenS2{F0: (v8 >> ((255) & 63))}
                    var v11 float64 = (3.14159 - (3.14159 * -2.25))
                    v12 := v1[(fi7 >> ((42) & 63))]
                    var v13 bool = (3.14159 >= 1e10)
                    _ = fi9
                    _ = v10
                    _ = v11
                    _ = v12
                    _ = v13
                }
                if false {
                    for fi14 := 0; fi14 < 15; fi14++ {
                        var v15 uint32 = 507
                        _ = fi14
                        _ = v15
                    }
                } else {
                    var v16 string = ("hello" + (("a" + ("" + ("" + "a"))) + ""))
                    var v17 uint64 = 18446744073709551615
                    var v18 uint64 = v17
                    _ = v16
                    _ = v17
                    _ = v18
                }
                var v19 int = int(v8)
                _ = v19
            }
            v20 := GenS2{F0: (v8 - v8)}
            _ = fi7
            _ = v20
            _ = v8
        }
    }
    v21 := GenS3{F0: 70, F1: "hello"}
    var v22 float32 = (0.1 / 3.14159)
    if true {
        for fi23 := 0; fi23 < 19; fi23++ {
            var v24 uint = 65535
            if true {
                var v25 bool = true
                _ = v25
            } else {
                var v26 uint64 = 2
                var v27 uint = 876
                _ = v26
                _ = v27
            }
            var v28 int64 = 1
            _ = fi23
            _ = v24
            _ = v28
        }
    }
    fmt.Println(v22)
    v22 += v22
    switch 9223372036854775807 {
    case -2147483648:
        v29 := GenS4{F0: 2}
        _ = v29
    case 2147483647:
        v30 := GenS5{F0: 1000, F1: 645}
        _ = v30
    }
    if true {
        v3 -= (v3 % (((v3 / ((0) | 1))) | 1))
    } else {
        v22 -= (-2.25 / 1.0)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
