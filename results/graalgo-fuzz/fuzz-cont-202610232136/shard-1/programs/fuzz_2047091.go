package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 bool }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 int16; F1 int64; F2 string }
func genFn0(p0 uint) uint8 {
    v7 := GenS1{F0: -9223372036854775808, F1: -186, F2: (1 < 42)}
    for fi8 := 0; fi8 < 3; fi8++ {
        p0 += (^(p0))
        v9 := []int{fi8, (fi8 * 32767)}
        p0 = (p0 + (p0 + (^(p0))))
        _ = fi8
        _ = v9
    }
    p0 += p0
    if (-543 >= 378) {
        for fi10 := 0; fi10 < 20; fi10++ {
            v11 := GenS2{F0: 647, F1: 215}
            if ((fi10 >> ((660) & 63)) <= fi10) {
                var v12 uint32 = 65535
                if true {
                    v13 := []int{}
                    var v14 uint64 = 42
                    if len(v13) > 0 {
                        v13[0] = 32768
                    }
                    _ = v13
                    _ = v14
                } else {
                    if false {
                        fmt.Println(v12, v12, fi10)
                    } else {
                        v15 := []int{(fi10 / ((32767) | 1))}
                        var v16 float32 = (-2.25 + 3.14159)
                        _ = v15
                        _ = v16
                    }
                }
                v17 := GenS3{F0: 1, F1: -1, F2: "\n\t"}
                _ = v12
                _ = v17
            } else {
                for fi18 := 0; fi18 < 1; fi18++ {
                    p0 = (p0 >> (((p0 * p0)) & 63))
                    p0 = (p0 >> (((p0 | (p0 & 255))) & 63))
                    switch "foo bar" {
                    case "\n\t":
                        fmt.Println(fi18)
                    }
                    _ = fi18
                }
            }
            p0 = p0
            _ = fi10
            _ = v11
        }
    }
    _ = p0
    _ = v7
    return 2
}

func main() {
    v1 := make(map[int]int)
    var v2 uint32 = 65535
    var v3 uint64 = 1
    var v4 uint8 = 1
    v5 := make(map[uint32]bool)
    v6 := make(map[int]string)
    if false {
        var v19 int32 = 0
        _ = v19
    } else {
        var v20 uint = 42
        for fi21 := 0; fi21 < 3; fi21++ {
            switch "hello" {
            case "":
                var v22 uint64 = (v3 + v3)
                _ = v22
            default:
                var v23 uint64 = (v3 ^ v3)
                _ = v23
            }
            v3 = (v3 + v3)
            _ = fi21
        }
        _ = v20
    }
    if (0.0 >= 0.0) {
        v4 -= (v4 * v4)
        for fi24 := 0; fi24 < 11; fi24++ {
            v4 += v4
            _ = fi24
        }
        v25 := v6[0]
        _ = v25
    }
    for fi26 := 0; fi26 < 7; fi26++ {
        var v27 uint = 4294967295
        if (false && false) {
            var v28 int16 = 1
            _ = v28
        }
        _ = fi26
        _ = v27
    }
    for fi29 := 0; fi29 < 16; fi29++ {
        v4 = (v4 % (((v4 * (^(v4)))) | 1))
        _ = fi29
    }
    fmt.Println(v2, v3, v2)
    var v30 string = ""
    var v31 int32 = int32(v2)
    v32 := v5[0]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
