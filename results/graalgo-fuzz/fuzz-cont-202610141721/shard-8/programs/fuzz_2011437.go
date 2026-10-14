package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 bool }
type GenS2 struct { F0 int32 }
func genFn0(p0 int16, p1 uint32) bool {
    p1 = p1
    p0 = p0
    var v4 float32 = 1.5
    v4 -= v4
    _ = p0
    _ = p1
    _ = v4
    return (!false)
}

func main() {
    var v1 int8 = 127
    var v2 int8 = v1
    var v3 float64 = 1.5
    v5 := []int{877, 9223372036854775807, 915}
    switch 892 {
    case 127:
        v6 := GenS1{F0: 960, F1: 1, F2: false}
        _ = v6
    }
    var v7 string = "a"
    switch 42 {
    case 2147483647:
        if (9223372036854775807 > 42) {
            var v8 uint32 = 0
            if (139 == 18446744073709551615) {
                for fi9 := 0; fi9 < 16; fi9++ {
                    var v10 float64 = float64((0.1 - (1e10 * (3.14159 - 0.0))))
                    fmt.Println(v1)
                    for fi11 := 0; fi11 < 16; fi11++ {
                        if (!true) {
                            var v12 uint64 = 1
                            v7 = v7
                            var v13 int64 = -100
                            v14 := GenS1{F0: v13, F1: v8, F2: true}
                            _ = v12
                            _ = v13
                            _ = v14
                        }
                        _ = fi11
                    }
                    _ = fi9
                    _ = v10
                }
            }
            for fi15 := 0; fi15 < 14; fi15++ {
                for fi16 := 0; fi16 < 20; fi16++ {
                    v5 = append(v5, 676)
                    _ = fi16
                }
                switch (v7 + "\n\t") {
                case "":
                    switch fi15 {
                    case -753:
                        switch ("hello" + v7) {
                        case "\n\t":
                            fmt.Println(v1, v3, v7)
                        case "":
                            var v17 int32 = -2147483648
                            _ = v17
                        case "hello":
                            var v18 uint64 = 0
                            _ = v18
                        }
                    }
                default:
                    v19 := GenS2{F0: -1}
                    _ = v19
                }
                _ = fi15
            }
            _ = v8
        }
    case 1:
        var v20 uint8 = 42
        _ = v20
    case 32768:
        var v21 float64 = v3
        _ = v21
    default:
        var v22 uint8 = 1
        _ = v22
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v23 uint32 = 304
    v24 := GenS1{F0: 9223372036854775807, F1: 111, F2: false}
    var v25 uint32 = (v23 & (^(v23)))
    for fi26 := 0; fi26 < 9; fi26++ {
        var v27 uint64 = 18446744073709551615
        var v28 int8 = (v1 % ((int8(fi26)) | 1))
        _ = fi26
        _ = v27
        _ = v28
    }
    if len(v5) > 0 {
        v5[0] = -804
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v5
    _ = v7
}
