package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int }
func genFn0(p0 uint32, p1 uint64, p2 uint32) uint64 {
    var v7 int64 = 9223372036854775807
    var v8 float64 = (0.0 + 0.1)
    p1 = (p1 / ((p1) | 1))
    v7 = (v7 + v7)
    var v9 uint64 = (p1 % (((^(p1))) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return (v9 % (((^(v9))) | 1))
}

func main() {
    var v1 float64 = float64(0.0)
    v2 := make(map[string]bool)
    var v3 int64 = 1
    var v4 int8 = -2
    var v5 int8 = v4
    v6 := make(map[int]string)
    v10 := []int{758, 2147483647, 1, 2147483647}
    v1 += v1
    for fi11 := 0; fi11 < 14; fi11++ {
        v12 := GenS1{F0: 176}
        fmt.Println(v4, fi11, v4)
        var v13 int = fi11
        _ = fi11
        _ = v12
        _ = v13
    }
    v1 -= v1
    switch "hello" {
    case "hello":
        var v14 float32 = (float32(0.0) + (1e10 + 0.1))
        _ = v14
    case "":
        if false {
            for fi15 := 0; fi15 < 12; fi15++ {
                for fi16 := 0; fi16 < 13; fi16++ {
                    v1 = ((v1 / 1.0) + (v1 - float64(0.1)))
                    var v17 uint64 = 2
                    fmt.Println(v3, fi15)
                    _ = fi16
                    _ = v17
                }
                for fi18 := 0; fi18 < 14; fi18++ {
                    var v19 bool = true
                    _ = len(v2)
                    var v20 float32 = 3.14159
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                v5 = v5
                _ = fi15
            }
            v4 = v5
            for fi21 := 0; fi21 < 13; fi21++ {
                if false {
                    if len(v10) > 0 {
                        v10[0] = (^(fi21))
                    }
                }
                v5 = v5
                _ = fi21
            }
        } else {
            v22 := GenS1{F0: 0}
            for fi23 := 0; fi23 < 6; fi23++ {
                var v24 uint8 = 2
                fmt.Println(v24, v4)
                var v25 uint64 = 232
                _ = fi23
                _ = v24
                _ = v25
            }
            if ((42 >= 95) && (true && (2147483648 > 1))) {
                var v26 uint16 = 484
                _ = v26
            } else {
                var v27 int8 = v5
                v28 := GenS1{F0: 127}
                v4 = 1
                _ = v27
                _ = v28
            }
            _ = v22
        }
    }
    v29 := GenS2{F0: 881}
    switch ("foo bar" + ("\n\t" + ("hello" + ""))) {
    case "a":
        var v30 int8 = (v5 ^ -1)
        _ = v30
    default:
        v5 += v4
    }
    if false {
        v3 += 88
    }
    v5 = 1
    v31 := GenS3{F0: 128}
    _ = len(v10)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v2
    _ = v29
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
