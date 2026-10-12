package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16 }
func genFn0(p0 bool) int {
    var v4 uint8 = 13
    for fi5 := 0; fi5 < 6; fi5++ {
        var v6 string = "\n\t"
        switch v6 {
        case "\n\t":
            v7 := GenS1{F0: 1}
            _ = v7
        default:
            for fi8 := 0; fi8 < 10; fi8++ {
                switch (-(fi8)) {
                case -2147483648:
                    fmt.Println(p0, fi8)
                case 0:
                    v9 := GenS1{F0: 2}
                    _ = v9
                case 32768:
                    v6 = v6
                default:
                    var v10 uint32 = 2147483647
                    _ = v10
                }
                if ((("\n\t" + v6) + "\n\t") <= v6) {
                    var v11 uint = 0
                    _ = v11
                }
                if true {
                    var v12 float32 = 1.5
                    if (float32((1.5 + 0.1)) < ((3.14159 / v12) - float32((-2.25 * 1.5)))) {
                        var v13 int64 = 0
                        v13 = (v13 / (((-(v13))) | 1))
                        var v14 uint8 = 42
                        _ = v13
                        _ = v14
                    } else {
                        if p0 {
                            p0 = p0
                        }
                    }
                    _ = v12
                }
                _ = fi8
            }
        }
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = v4
    return 128
}

func main() {
    var v1 uint32 = 202
    v2 := make(map[int]int)
    var v3 int64 = 9223372036854775807
    for fi15 := 0; fi15 < 4; fi15++ {
        v3 = (^(v3))
        _ = fi15
    }
    v16 := []int{-1, 960, genFn0((true && true)), -96}
    v3 = 42
    v17 := GenS2{F0: 0}
    fmt.Println(v1, v1)
    var v18 float64 = ((float64(1e10) + 0.1) - 0.0)
    if false {
        var v19 uint = 564
        switch 128 {
        case 128:
            if true {
                v20 := GenS2{F0: 42}
                fmt.Println(v19, v19, v18)
                var v21 int8 = -58
                _ = v20
                _ = v21
            } else {
                var v22 int64 = -1
                _ = v22
            }
        case 1:
            var v23 uint = (v19 << (((^(v19))) & 63))
            _ = v23
        default:
            if true {
                _ = len(v2)
                switch 4294967295 {
                case 128:
                    var v24 uint32 = v1
                    _ = v24
                default:
                    fmt.Println(v3)
                }
            }
        }
        _ = len(v16)
        _ = v19
    } else {
        fmt.Println(v1, v3)
    }
    var v25 uint8 = 0
    var v26 int32 = -1
    for fi27 := 0; fi27 < 4; fi27++ {
        v16 = append(v16, 1)
        fmt.Println(v1, v3, v1)
        _ = fi27
    }
    v26 = (v26 ^ (v26 % (((v26 << ((979) & 63))) | 1)))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v25
    _ = v26
    _ = v3
}
