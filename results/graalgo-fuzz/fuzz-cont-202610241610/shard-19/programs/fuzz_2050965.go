package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 uint16; F1 int16 }
type GenS3 struct { F0 bool; F1 int; F2 uint64 }
func genFn0(p0 uint32, p1 int, p2 uint) uint64 {
    var v7 uint = (p2 + 1)
    fmt.Println(v7, v7, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 2
}
func genFn1(p0 uint16, p1 int32, p2 int64) int64 {
    var v8 float64 = 3.14159
    v9 := GenS1{F0: p2, F1: 128}
    if (!(777 == 905)) {
        fmt.Println(p2, v8)
        var v10 int8 = -128
        _ = v10
    } else {
        if true {
            fmt.Println(p0, v8)
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return (p2 - p2)
}

func main() {
    v1 := []int{-9223372036854775808, 32767, 2}
    v2 := []int{-1, 713, 679, 9223372036854775807}
    v3 := make(map[string]uint32)
    var v4 float64 = 1e10
    v5 := []int{128, 4294967295, -9223372036854775808}
    var v6 int16 = -1
    for fi11 := 0; fi11 < 1; fi11++ {
        var v12 int = (fi11 * 4294967295)
        _ = fi11
        _ = v12
    }
    v13 := GenS2{F0: 42, F1: v6}
    v14 := GenS2{F0: 42, F1: (v6 * (v6 % (((v6 & v6)) | 1)))}
    if false {
        if (1 >= 127) {
            var v15 int = -1
            if false {
                for fi16 := 0; fi16 < 16; fi16++ {
                    v4 += v4
                    var v17 int8 = 100
                    _ = fi16
                    _ = v17
                }
                v18 := GenS3{F0: false, F1: v15, F2: 925}
                for fi19 := 0; fi19 < 4; fi19++ {
                    v4 = v4
                    _ = fi19
                }
                _ = v18
            } else {
                var v20 int = v15
                _ = v20
            }
            fmt.Println(v6)
            _ = v15
        }
        for fi21 := 0; fi21 < 3; fi21++ {
            for fi22 := 0; fi22 < 7; fi22++ {
                v23 := v3[""]
                v3[""] = v23
                fmt.Println(fi22, v6, v6)
                _ = fi22
                _ = v23
            }
            _ = fi21
        }
    } else {
        var v24 uint64 = 198
        _ = v24
    }
    v6 -= -32768
    var v25 int16 = 1
    v25 = (v6 % ((v25) | 1))
    v26 := GenS2{F0: 1, F1: (^(v25))}
    v25 = v6
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
