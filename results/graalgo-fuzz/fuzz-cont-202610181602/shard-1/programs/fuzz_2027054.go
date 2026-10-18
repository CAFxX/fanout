package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 uint8; F1 int16; F2 int32 }
type GenS4 struct { F0 int16 }
func genFn0(p0 int) bool {
    v6 := []int{}
    p0 = (-(p0))
    _ = p0
    _ = v6
    return true
}
func genFn1() uint16 {
    v7 := GenS1{F0: 127}
    if (65535 >= 2) {
        var v8 string = ""
        if false {
            var v9 int = 659
            var v10 int32 = 1
            _ = v10
            _ = v9
        }
        _ = v8
    }
    if (0 > 87) {
        if false {
            v11 := []int{127}
            var v12 uint8 = 42
            for fi13 := 0; fi13 < 18; fi13++ {
                for fi14 := 0; fi14 < 11; fi14++ {
                    var v15 float64 = ((((1e10 / (((-2.25 + ((0.1 / (-2.25 / (3.14159 - 1e10))) * float64((3.14159 / 1.5)))) - 1.5) / 1e10)) * -2.25) * 1.5) * -2.25)
                    fmt.Println(v15, v15, fi14)
                    if true {
                        var v16 uint64 = 1
                        v17 := GenS1{F0: v12}
                        if len(v11) > 0 {
                            v11[0] = (-(fi13))
                        }
                        _ = v16
                        _ = v17
                    }
                    _ = fi14
                    _ = v15
                }
                v12 -= v12
                _ = fi13
            }
            _ = v11
            _ = v12
        } else {
            v18 := GenS2{F0: 2147483648, F1: -1}
            v19 := GenS2{F0: 1, F1: 0}
            _ = v18
            _ = v19
        }
        var v20 uint = 1
        v20 += v20
        _ = v20
    }
    _ = v7
    return 822
}

func main() {
    var v1 float32 = ((0.1 / 0.1) + (0.1 / (1e10 - -2.25)))
    var v2 uint64 = 1
    var v3 uint = 18446744073709551615
    v4 := make(map[int]bool)
    var v5 uint16 = 65535
    switch "a" {
    case "":
        v21 := GenS3{F0: 1, F1: -919, F2: 982}
        _ = v21
    case "hello":
        v22 := GenS2{F0: 1, F1: -419}
        _ = v22
    }
    v23 := []int{-195}
    v24 := GenS4{F0: 325}
    v1 = -2.25
    var v25 uint = 2
    var v26 float32 = (0.1 + v1)
    var v27 int64 = 1
    delete(v4, 4294967295)
    fmt.Println(v26, v2)
    var v28 string = ""
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
