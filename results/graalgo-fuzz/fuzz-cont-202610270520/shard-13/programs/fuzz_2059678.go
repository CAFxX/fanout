package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 bool }
func genFn0(p0 uint64, p1 uint64, p2 uint) int {
    v5 := []int{783}
    p2 -= p2
    p1 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 0
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 int32 = 802
    v3 := make(map[int]uint32)
    var v4 float32 = (3.14159 / 3.14159)
    v4 = 0.1
    v6 := GenS1{F0: "\n\t", F1: 0, F2: 42}
    v7 := GenS2{F0: 32768}
    if (4294967295 >= 1) {
        for fi8 := 0; fi8 < 2; fi8++ {
            v1 = (^(v1))
            _ = fi8
        }
        var v9 float32 = 0.0
        _ = v9
    }
    v10 := []int{}
    var v11 bool = true
    for fi12 := 0; fi12 < 1; fi12++ {
        _ = len(v10)
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
        _ = fi12
    }
    fmt.Println(v4)
    var v13 bool = (427 >= 351)
    v14 := GenS3{F0: 127}
    v1 -= 1
    switch 127 {
    case -2147483648:
        var v15 int16 = 0
        _ = v15
    case -966:
        if (3.14159 >= 0.0) {
            var v16 int32 = v2
            for fi17 := 0; fi17 < 10; fi17++ {
                for fi18 := 0; fi18 < 18; fi18++ {
                    switch "" {
                    case "\n\t":
                        v13 = true
                    case "foo bar":
                        if v13 {
                            var v19 uint16 = 0
                            _ = v19
                        }
                    }
                    v20 := GenS4{F0: (v13 || v13)}
                    var v21 int16 = -32768
                    _ = fi18
                    _ = v20
                    _ = v21
                }
                for fi22 := 0; fi22 < 15; fi22++ {
                    var v23 int32 = (v2 ^ (-(v2)))
                    v16 += 2147483647
                    var v24 bool = (1.5 != float32(1e10))
                    _ = fi22
                    _ = v23
                    _ = v24
                }
                _ = fi17
            }
            _ = v16
        } else {
            var v25 int = genFn0(59, 2, 255)
            for fi26 := 0; fi26 < 13; fi26++ {
                for fi27 := 0; fi27 < 3; fi27++ {
                    v13 = v13
                    _ = fi27
                }
                v2 = v2
                for fi28 := 0; fi28 < 12; fi28++ {
                    v1 = 219
                    switch (v25 % ((14) | 1)) {
                    case 9223372036854775807:
                        v25 -= 881
                        v3[v25] = 1
                        var v29 float64 = 0.1
                        _ = v29
                    case 1:
                        if len(v10) >= 2 {
                            v10 = v10[1:len(v10)]
                        }
                    default:
                        switch fi28 {
                        case 0:
                            if len(v10) > 0 {
                                v10[0] = (fi26 & fi26)
                            }
                        }
                    }
                    var v30 int32 = -1
                    _ = fi28
                    _ = v30
                }
                _ = fi26
            }
            _ = v25
        }
    }
    v13 = false
    for fi31 := 0; fi31 < 14; fi31++ {
        v4 -= 1.5
        _ = fi31
    }
    v1 += 0
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
