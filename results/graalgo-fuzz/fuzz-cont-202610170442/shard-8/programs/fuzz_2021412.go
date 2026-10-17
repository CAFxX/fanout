package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32; F1 uint8 }
func genFn0(p0 int16, p1 int8, p2 int16) int8 {
    for fi6 := 0; fi6 < 16; fi6++ {
        p1 = (p1 >> ((1) & 63))
        var v7 float32 = -2.25
        var v8 int8 = 127
        _ = fi6
        _ = v7
        _ = v8
    }
    if (!false) {
        p1 += p1
        p2 = (-(p0))
        var v9 int64 = 9223372036854775807
        _ = v9
    }
    v10 := []int{9223372036854775807, -676}
    if len(v10) > 0 {
        v10[0] = 127
    }
    for fi11 := 0; fi11 < 10; fi11++ {
        if false {
            var v12 int64 = -9223372036854775808
            _ = v12
        }
        p1 = (p1 << ((273) & 63))
        _ = fi11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return p1
}

func main() {
    v1 := []int{2147483647, 32767, -9223372036854775808}
    var v2 uint16 = 1
    var v3 string = "foo bar"
    v4 := make(map[int]uint32)
    v5 := []string{}
    fmt.Println(v2)
    var v13 float32 = float32((1e10 + (3.14159 + -2.25)))
    v13 = v13
    var v14 uint8 = 156
    if true {
        v13 += float32(1e10)
    }
    var v15 int = 103
    v15 = (v15 - int(v14))
    var v16 int16 = 0
    v4[(v15 + v15)] = 42
    fmt.Println(v16)
    v15 = (v15 >> ((42) & 63))
    v15 -= 2147483647
    var v17 uint = 18446744073709551615
    v18 := GenS1{F0: 42}
    if (!(true || ((128 >= v15) || false))) {
        v17 += 42
    } else {
        for fi19 := 0; fi19 < 12; fi19++ {
            for fi20 := 0; fi20 < 14; fi20++ {
                v16 -= v16
                v5 = append(v5, v3)
                switch "hello" {
                case "\n\t":
                    if (v17 == v17) {
                        v21 := GenS2{F0: 2147483647}
                        _ = v21
                    } else {
                        v1 = append(v1, -9223372036854775808)
                    }
                default:
                    fmt.Println(v3)
                }
                v16 = v16
                _ = len(v4)
                _ = fi20
            }
            switch fi19 {
            case 9223372036854775807:
                if (((427 == 280) || true) || (int(v14) > (v15 % ((fi19) | 1)))) {
                    fmt.Println(v3)
                    v15 = -1
                } else {
                    v22 := GenS1{F0: v16}
                    var v23 int8 = 87
                    _ = v22
                    _ = v23
                }
            case 32767:
                for fi24 := 0; fi24 < 17; fi24++ {
                    v2 -= 2
                    var v25 uint32 = 2147483647
                    var v26 float32 = v13
                    v26 = v13
                    v27 := GenS3{F0: -448, F1: 127}
                    v13 += 0.0
                    _ = fi24
                    _ = v25
                    _ = v26
                    _ = v27
                }
            }
            var v28 uint = (^(v17))
            _ = fi19
            _ = v28
        }
    }
    v29 := GenS3{F0: -708, F1: (v14 % (((v14 * (v14 & v14))) | 1))}
    var v30 uint8 = (v14 | (^(v14)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
