package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
func genFn0(p0 int64) uint64 {
    v6 := []int{-1, 0, 127}
    switch "a" {
    case "foo bar":
        var v7 uint32 = 2147483648
        _ = v7
    default:
        p0 += 296
    }
    _ = p0
    _ = v6
    return 42
}

func main() {
    var v1 int = 32768
    var v2 uint32 = 2
    v3 := make(map[uint32]int)
    var v4 int8 = 42
    var v5 int16 = 0
    var v8 int32 = int32(v5)
    v9 := []int{v1, v1, v1}
    _ = len(v9)
    var v10 float32 = 0.0
    for fi11 := 0; fi11 < 15; fi11++ {
        v10 += (float32((1.5 / ((1.5 + 3.14159) + (1.5 - 0.0)))) * v10)
        v1 += v1
        if true {
            fmt.Println(v2, v10, v1)
            v5 = v5
            if (true && true) {
                fmt.Println(v4, fi11, v8)
                switch -825 {
                case 32768:
                    for fi12 := 0; fi12 < 1; fi12++ {
                        v8 -= (^(v8))
                        switch "foo bar" {
                        case "hello":
                            if false {
                                _ = len(v9)
                            } else {
                                v4 = (-(v4))
                                var v13 int8 = v4
                                fmt.Println(v4)
                                _ = v13
                            }
                        case "\n\t":
                            var v14 uint32 = (v2 << ((18446744073709551615) & 63))
                            _ = v14
                        }
                        _ = fi12
                    }
                case -353:
                    v1 = -2147483648
                case -591:
                    fmt.Println(fi11, v10)
                default:
                    v4 = (v4 ^ v4)
                }
            }
        } else {
            for fi15 := 0; fi15 < 2; fi15++ {
                var v16 bool = false
                _ = fi15
                _ = v16
            }
        }
        _ = fi11
    }
    v10 -= 1e10
    v17 := GenS1{F0: 255, F1: 272}
    fmt.Println(v5)
    var v18 float64 = 1.5
    fmt.Println(v18, v2)
    if ((("\n\t" + "\n\t") + "foo bar") > "") {
        fmt.Println(v1, v1, v10)
    }
    var v19 float64 = (3.14159 + v18)
    var v20 string = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
