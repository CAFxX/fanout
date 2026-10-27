package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 uint8; F1 bool; F2 uint8 }
type GenS4 struct { F0 uint16; F1 uint; F2 int16 }
type GenS5 struct { F0 uint16; F1 int32; F2 uint16 }

func main() {
    var v1 int64 = 279
    var v2 int8 = 100
    var v3 uint = 255
    fmt.Println(v1, v1, v3)
    for fi4 := 0; fi4 < 11; fi4++ {
        v5 := []int{}
        var v6 float64 = (0.0 + 3.14159)
        _ = fi4
        _ = v5
        _ = v6
    }
    v1 = (^(v1))
    var v7 float64 = (0.0 - 1e10)
    v8 := []int{-130, -2147483648, 32768}
    for fi9 := 0; fi9 < 18; fi9++ {
        var v10 uint32 = 42
        if false {
            switch ("a" + "") {
            case "hello":
                if ((v2 + (v2 | (v2 & v2))) <= (v2 * (^(v2)))) {
                    v10 += (v10 * (^(v10)))
                }
            case "\n\t":
                v10 -= 2
            default:
                for fi11 := 0; fi11 < 5; fi11++ {
                    fmt.Println(v10, v3)
                    _ = fi11
                }
            }
            switch int(v3) {
            case 32768:
                var v12 uint64 = 18446744073709551615
                _ = v12
            case 0:
                for fi13 := 0; fi13 < 15; fi13++ {
                    if (1 == 856) {
                        v3 = v3
                        if false {
                            v14 := GenS1{F0: 0, F1: 255}
                            var v15 float64 = (1e10 * 0.0)
                            _ = v14
                            _ = v15
                        }
                    } else {
                        var v16 int16 = 1
                        _ = v16
                    }
                    v1 += (v1 ^ (v1 & -1))
                    v17 := GenS2{F0: 182, F1: (^(fi9))}
                    var v18 uint32 = v10
                    _ = fi13
                    _ = v17
                    _ = v18
                }
            }
        }
        v19 := GenS3{F0: 2, F1: true, F2: 0}
        _ = fi9
        _ = v10
        _ = v19
    }
    switch "" {
    case "\n\t":
        fmt.Println(v1)
    case "a":
        var v20 bool = true
        _ = v20
    }
    for fi21 := 0; fi21 < 18; fi21++ {
        v1 -= v1
        var v22 int = (-(fi21))
        _ = fi21
        _ = v22
    }
    var v23 int = -788
    v24 := GenS4{F0: 868, F1: 2, F2: -68}
    v25 := GenS5{F0: 2, F1: -391, F2: 255}
    var v26 int32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v7
    _ = v8
}
