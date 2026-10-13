package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 uint16 }
type GenS2 struct { F0 bool; F1 string; F2 int64 }
type GenS3 struct { F0 int; F1 int8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int16; F1 uint; F2 uint16 }

func main() {
    var v1 float64 = (0.1 * (float64(float32(float64((((-2.25 / ((((0.0 * float32((3.14159 - 1.5))) + 1e10) / 1.5) / 1.0)) * (-2.25 - 3.14159)) * 0.1)))) + 1.5))
    var v2 bool = true
    v3 := []int{}
    var v4 uint64 = 42
    fmt.Println(v2, v4)
    if (v2 && true) {
        v5 := GenS1{F0: -292, F1: 0, F2: 42}
        for fi6 := 0; fi6 < 17; fi6++ {
            if (255 >= 65535) {
                switch "\n\t" {
                case "a":
                    for fi7 := 0; fi7 < 15; fi7++ {
                        var v8 int32 = 582
                        _ = fi7
                        _ = v8
                    }
                case "foo bar":
                    switch fi6 {
                    case -877:
                        v4 += (v4 * 2)
                    case 4294967295:
                        fmt.Println(fi6)
                    }
                }
                for fi9 := 0; fi9 < 2; fi9++ {
                    v10 := GenS1{F0: -177, F1: 0, F2: 65535}
                    _ = fi9
                    _ = v10
                }
            } else {
                fmt.Println(v1, v4, fi6)
                for fi11 := 0; fi11 < 9; fi11++ {
                    v12 := GenS2{F0: (169 == 1), F1: "", F2: -145}
                    v1 = 1.5
                    _ = fi11
                    _ = v12
                }
            }
            v1 += ((v1 - (0.0 / ((0.1 + 0.1) / 0.1))) / v1)
            _ = fi6
        }
        _ = v5
    } else {
        for fi13 := 0; fi13 < 10; fi13++ {
            _ = len(v3)
            if (-32768 > 0) {
                var v14 float32 = 0.1
                _ = v14
            } else {
                v15 := GenS2{F0: false, F1: "hello", F2: 0}
                _ = v15
            }
            _ = fi13
        }
        fmt.Println(v4, v2)
    }
    fmt.Println(v2, v1, v4)
    v16 := GenS3{F0: -82, F1: 0}
    v17 := GenS3{F0: 577, F1: -61}
    if (2147483647 == 42) {
        var v18 int8 = 127
        v19 := GenS4{F0: (v4 ^ 0)}
        if v2 {
            var v20 float32 = 1e10
            var v21 int = 119
            _ = v20
            _ = v21
        } else {
            var v22 uint = 1
            if len(v3) > 0 {
                v3[0] = 2147483647
            }
            var v23 int64 = 9223372036854775807
            _ = v22
            _ = v23
        }
        _ = v18
        _ = v19
    } else {
        for fi24 := 0; fi24 < 10; fi24++ {
            if true {
                v25 := GenS1{F0: 32767, F1: 42, F2: 409}
                fmt.Println(v2)
                _ = v25
            }
            for fi26 := 0; fi26 < 20; fi26++ {
                var v27 uint16 = 2
                if (9223372036854775807 < -786) {
                    v28 := GenS5{F0: uint(v4)}
                    _ = v28
                }
                v3 = append(v3, (fi26 ^ (fi26 | fi26)))
                v1 += 0.1
                _ = fi26
                _ = v27
            }
            switch (fi24 * fi24) {
            case 32768:
                var v29 uint = 65535
                _ = v29
            default:
                for fi30 := 0; fi30 < 18; fi30++ {
                    v31 := GenS2{F0: false, F1: "foo bar", F2: -914}
                    switch -512 {
                    case -2147483648:
                        v32 := GenS6{F0: -500, F1: 42, F2: 820}
                        _ = v32
                    }
                    _ = fi30
                    _ = v31
                }
            }
            _ = fi24
        }
    }
    fmt.Println(v4)
    var v33 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v33
    _ = v4
}
