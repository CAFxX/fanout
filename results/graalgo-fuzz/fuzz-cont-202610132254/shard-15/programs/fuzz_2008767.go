package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 uint16 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 uint16) uint {
    fmt.Println(p0)
    fmt.Println(p0)
    _ = p0
    return 2
}
func genFn1(p0 int64) uint {
    v5 := GenS1{F0: 65535, F1: 2, F2: 32767}
    if false {
        var v6 bool = (128 == 127)
        _ = v6
    }
    v7 := GenS1{F0: 65535, F1: 568, F2: 1}
    _ = p0
    _ = v5
    _ = v7
    return 255
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]string)
    var v3 float32 = (0.1 + 1.5)
    var v4 int64 = 9223372036854775807
    switch "hello" {
    case "foo bar":
        v8 := []int{-268, 709, 4294967295}
        _ = v8
    default:
        v9 := []int{712, 127}
        _ = v9
    }
    v10 := []int{2147483647, -923, -144}
    v4 -= -9223372036854775808
    v11 := GenS1{F0: 2, F1: 952, F2: 32767}
    if false {
        v1[4294967295] = false
        var v12 bool = (245 == 42)
        _ = v12
    } else {
        for fi13 := 0; fi13 < 19; fi13++ {
            fmt.Println(v3)
            _ = fi13
        }
    }
    var v14 float32 = ((((v3 / (float32(-2.25) * 1.5)) * (v3 - 3.14159)) / (-2.25 + (v3 * float32(float64(v3))))) / v3)
    if false {
        fmt.Println(v3, v3, v4)
        for fi15 := 0; fi15 < 9; fi15++ {
            delete(v2, 2147483647)
            var v16 string = (("\n\t" + "\n\t") + "a")
            if (!true) {
                v17 := GenS2{F0: 128}
                _ = v17
            } else {
                v3 += v14
                for fi18 := 0; fi18 < 15; fi18++ {
                    v3 = 0.0
                    var v19 uint32 = 4294967295
                    v10 = append(v10, (-(fi15)))
                    _ = fi18
                    _ = v19
                }
                var v20 uint16 = 32767
                _ = v20
            }
            _ = fi15
            _ = v16
        }
    } else {
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
        v14 = ((-2.25 / v3) + 1e10)
        _ = len(v2)
    }
    var v21 int64 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
