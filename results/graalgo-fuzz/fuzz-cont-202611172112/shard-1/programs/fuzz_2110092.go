package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 int16 }

func main() {
    v1 := make(map[int]int)
    var v2 uint16 = 177
    var v3 int32 = -1
    v4 := make(map[int]bool)
    var v5 int64 = -230
    v6 := make(map[int]string)
    fmt.Println(v2, v3, v2)
    if false {
        v7 := GenS1{F0: 600}
        for fi8 := 0; fi8 < 17; fi8++ {
            if false {
                for fi9 := 0; fi9 < 18; fi9++ {
                    v3 -= (^(v3))
                    _ = fi9
                }
                v5 += 42
            }
            _ = fi8
        }
        if true {
            for fi10 := 0; fi10 < 3; fi10++ {
                switch "a" {
                case "foo bar":
                    fmt.Println(v3, v2, v5)
                case "a":
                    var v11 uint8 = 218
                    _ = v11
                case "\n\t":
                    v12 := []int{}
                    _ = v12
                }
                _ = fi10
            }
        }
        _ = v7
    } else {
        var v13 float64 = (-2.25 + (1.5 - 0.1))
        _ = v13
    }
    var v14 int = 221
    v14 = 73
    switch "" {
    case "hello":
        if ((false || true) || true) {
            if (387 >= 42) {
                switch -1 {
                case -1:
                    var v15 int64 = v5
                    _ = v15
                case 912:
                    for fi16 := 0; fi16 < 5; fi16++ {
                        var v17 int32 = v3
                        _ = fi16
                        _ = v17
                    }
                case -690:
                    v18 := GenS1{F0: 875}
                    _ = v18
                }
                _ = len(v1)
                var v19 float64 = ((float64(1.5) - 1.5) / 0.1)
                _ = v19
            }
            var v20 uint32 = 925
            var v21 uint32 = 2
            _ = v20
            _ = v21
        } else {
            fmt.Println(v2)
            var v22 int16 = 0
            _ = v22
        }
    default:
        v23 := []int{v14, 127, -9223372036854775808, (v14 << ((2) & 63))}
        _ = v23
    }
    v14 = v14
    if ((3.14159 * 1.5) == (0.0 * (0.1 / (1e10 + 1e10)))) {
        v24 := GenS1{F0: 1}
        _ = v24
    } else {
        for fi25 := 0; fi25 < 4; fi25++ {
            if false {
                var v26 int32 = (v3 - (v3 * (v3 / (((v3 << ((65535) & 63))) | 1))))
                switch (v14 - (fi25 ^ (fi25 ^ (v14 & v14)))) {
                case -1:
                    fmt.Println(fi25, fi25)
                case 77:
                    switch v14 {
                    case 612:
                        v27 := []int{v14, -13, (^(v14)), (^(fi25))}
                        _ = v27
                    default:
                        v28 := []int{}
                        _ = v28
                    }
                }
                v29 := GenS2{F0: 1, F1: 568}
                _ = v26
                _ = v29
            } else {
                fmt.Println(v14)
            }
            fmt.Println(v5, v5)
            var v30 int = v14
            _ = fi25
            _ = v30
        }
        var v31 uint8 = 0
        var v32 int8 = -66
        _ = v31
        _ = v32
    }
    v5 += (v5 % ((v5) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
