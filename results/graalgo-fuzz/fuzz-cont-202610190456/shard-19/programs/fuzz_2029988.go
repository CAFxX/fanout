package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8; F1 int32 }
type GenS4 struct { F0 int32 }

func main() {
    var v1 bool = true
    var v2 int64 = -1
    var v3 int32 = -2147483648
    v4 := []int{0}
    var v5 float32 = 0.0
    switch "hello" {
    case "a":
        fmt.Println(v3, v1)
    case "\n\t":
        var v6 int32 = v3
        _ = v6
    default:
        if ((v5 / ((v5 + 1.5) / (v5 / (v5 - 3.14159)))) <= (((v5 + 3.14159) + v5) / v5)) {
            switch 42 {
            case 9223372036854775807:
                _ = len(v4)
            case 1:
                v4 = append(v4, 127)
            case 174:
                _ = len(v4)
            default:
                if true {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    fmt.Println(v2)
                    switch 9223372036854775807 {
                    case 2147483647:
                        var v7 float32 = ((1.5 / v5) - v5)
                        _ = v7
                    case 0:
                        var v8 bool = (455 < 2147483648)
                        _ = v8
                    default:
                        if len(v4) > 0 {
                            v4[0] = -2147483648
                        }
                    }
                } else {
                    for fi9 := 0; fi9 < 3; fi9++ {
                        var v10 float64 = 1e10
                        _ = fi9
                        _ = v10
                    }
                    if (v1 || (v3 <= (v3 ^ (v3 >> ((2) & 63))))) {
                        v11 := GenS1{F0: 84}
                        if v1 {
                            _ = len(v4)
                        } else {
                            for fi12 := 0; fi12 < 2; fi12++ {
                                for fi13 := 0; fi13 < 17; fi13++ {
                                    var v14 uint = 2
                                    var v15 uint = (v14 / ((v14) | 1))
                                    _ = fi13
                                    _ = v14
                                    _ = v15
                                }
                                _ = fi12
                            }
                            v16 := GenS2{F0: 0}
                            var v17 uint64 = 18446744073709551615
                            _ = v16
                            _ = v17
                        }
                        _ = v11
                    }
                    var v18 int64 = 0
                    _ = v18
                }
            }
        }
    }
    _ = len(v4)
    v19 := GenS3{F0: 127, F1: v3}
    if (102 >= -128) {
        fmt.Println(v3)
    }
    _ = len(v4)
    v2 = -569
    v20 := GenS4{F0: (^(v3))}
    fmt.Println(v2)
    var v21 int32 = (v3 / (((^(v3))) | 1))
    var v22 float64 = float64(v5)
    var v23 int16 = 42
    fmt.Println(v21, v23)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
