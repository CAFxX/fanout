package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 int8 }
type GenS2 struct { F0 int16; F1 int8 }
func genFn0() uint8 {
    var v5 uint16 = 118
    fmt.Println(v5)
    v5 += (v5 * v5)
    var v6 uint = 18446744073709551615
    _ = v5
    _ = v6
    return 1
}

func main() {
    var v1 uint8 = 50
    var v2 int64 = 9223372036854775807
    var v3 int = 2147483647
    var v4 string = "a"
    v4 = v4
    var v7 uint16 = 2
    var v8 float64 = float64((0.0 + 1.5))
    v3 = v3
    var v9 int32 = -2147483648
    fmt.Println(v4, v4)
    v1 -= (v1 << ((255) & 63))
    var v10 int8 = 1
    var v11 uint32 = 1
    switch v3 {
    case -46:
        v12 := []int{(v3 % (((v3 + int(v11))) | 1)), 127}
        _ = v12
    case 127:
        if false {
            if (!true) {
                var v13 uint64 = 0
                _ = v13
            }
            v14 := []int{(-(v3)), v3, v3}
            if len(v14) > 0 {
                v14[0] = v3
            }
            _ = v14
        } else {
            v15 := GenS1{F0: uint64(v1), F1: (v11 & (v11 % ((42) | 1))), F2: v10}
            if (!false) {
                for fi16 := 0; fi16 < 18; fi16++ {
                    v9 = int32(v1)
                    for fi17 := 0; fi17 < 3; fi17++ {
                        if (true && (0.1 >= float32(v8))) {
                            var v18 uint16 = v7
                            v10 = v10
                            _ = v18
                        }
                        var v19 string = "foo bar"
                        v11 += v11
                        _ = fi17
                        _ = v19
                    }
                    _ = fi16
                }
            } else {
                v3 = 4294967295
            }
            _ = v15
        }
    default:
        v8 += ((v8 * v8) - (v8 * float64(0.0)))
    }
    for fi20 := 0; fi20 < 19; fi20++ {
        if false {
            if (!(18446744073709551615 > 22)) {
                v11 -= v11
            }
            if (!true) {
                v3 = 4294967295
                switch v4 {
                case "foo bar":
                    v7 += 593
                case "\n\t":
                    for fi21 := 0; fi21 < 11; fi21++ {
                        var v22 int32 = (v9 >> ((2) & 63))
                        _ = fi21
                        _ = v22
                    }
                default:
                    v23 := []int{9223372036854775807}
                    _ = v23
                }
                var v24 int16 = 535
                _ = v24
            } else {
                switch fi20 {
                case 2147483647:
                    fmt.Println(v3, v2, fi20)
                }
                v25 := []int{(^(fi20))}
                for fi26 := 0; fi26 < 7; fi26++ {
                    v3 += (fi26 >> ((944) & 63))
                    _ = fi26
                }
                _ = v25
            }
            var v27 float64 = (float64(0.1) / (v8 - float64(3.14159)))
            _ = v27
        } else {
            var v28 uint16 = (v7 - (v7 | (v7 ^ v7)))
            var v29 uint32 = (v11 << ((0) & 63))
            for fi30 := 0; fi30 < 9; fi30++ {
                v31 := []int{fi30, int(v28)}
                v4 = "hello"
                var v32 uint64 = 305
                _ = fi30
                _ = v31
                _ = v32
            }
            _ = v28
            _ = v29
        }
        fmt.Println(fi20)
        _ = fi20
    }
    v33 := GenS2{F0: 1, F1: -128}
    v34 := []int{}
    var v35 float64 = v8
    var v36 float32 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
