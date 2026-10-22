package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64; F1 int64; F2 uint16 }
func genFn0(p0 bool, p1 uint32) bool {
    var v5 uint = 255
    v6 := GenS1{F0: -9223372036854775808, F1: int32(p1)}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return p0
}
func genFn1(p0 uint, p1 int, p2 int) uint16 {
    if false {
        if (!false) {
            fmt.Println(p0, p0)
        } else {
            if true {
                p1 += 215
            } else {
                v7 := GenS2{F0: 127}
                if ((-2.25 - (((1.5 - float32(0.0)) + -2.25) / (float32(3.14159) - (-2.25 - 3.14159)))) == float32((float64(float32(0.1)) * 1e10))) {
                    v8 := GenS2{F0: 0}
                    fmt.Println(p2, p0, p1)
                    _ = v8
                } else {
                    var v9 int16 = -1
                    if (1 >= (p0 & (p0 * 128))) {
                        for fi10 := 0; fi10 < 3; fi10++ {
                            fmt.Println(fi10, p2)
                            _ = fi10
                        }
                        v11 := GenS1{F0: -23, F1: 615}
                        _ = v11
                    }
                    _ = v9
                }
                var v12 bool = (false || true)
                _ = v12
                _ = v7
            }
            for fi13 := 0; fi13 < 14; fi13++ {
                var v14 uint32 = 0
                if (!false) {
                    v15 := GenS3{F0: uint64(v14), F1: 0, F2: 0}
                    p2 = 128
                    _ = v15
                }
                p1 = (-(fi13))
                _ = fi13
                _ = v14
            }
            var v16 uint32 = 65535
            _ = v16
        }
        v17 := []int{}
        p1 = (p2 ^ (p1 | p2))
        _ = v17
    } else {
        var v18 float64 = (3.14159 * ((-2.25 - 1e10) * -2.25))
        var v19 string = ("hello" + "\n\t")
        _ = v18
        _ = v19
    }
    var v20 float32 = 3.14159
    v20 = ((v20 / -2.25) + -2.25)
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    return 32767
}

func main() {
    var v1 uint = 2
    v2 := make(map[int]uint32)
    var v3 float64 = 0.0
    v4 := []string{}
    v3 += 0.0
    var v21 uint8 = 42
    var v22 uint = (^(v1))
    v1 += 1
    if true {
        v23 := GenS3{F0: 4294967295, F1: 42, F2: 255}
        _ = v23
    }
    var v24 int64 = 70
    v1 = 0
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v24
    _ = v3
    _ = v4
}
