package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
func genFn0(p0 string, p1 bool, p2 int16) uint16 {
    p2 += p2
    if p1 {
        for fi6 := 0; fi6 < 6; fi6++ {
            var v7 float64 = ((0.1 - 0.0) + (3.14159 + -2.25))
            if (32767 > 614) {
                v8 := []int{fi6, (fi6 % (((fi6 & (fi6 % ((-166) | 1)))) | 1))}
                _ = v8
            } else {
                fmt.Println(v7, fi6, p1)
                var v9 int8 = -1
                _ = v9
            }
            _ = fi6
            _ = v7
        }
    } else {
        v10 := []int{4294967295, 9223372036854775807}
        _ = v10
    }
    if p1 {
        if (p0 == p0) {
            p0 = p0
        } else {
            var v11 int = -88
            var v12 bool = (1.5 <= (-2.25 / (0.1 * (1e10 - 3.14159))))
            p2 = p2
            _ = v11
            _ = v12
        }
        var v13 int = -1
        _ = v13
    } else {
        for fi14 := 0; fi14 < 14; fi14++ {
            p1 = (!(0 < 128))
            _ = fi14
        }
        v15 := []int{32768, 4294967295}
        fmt.Println(p0, p0, p2)
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    return 65535
}

func main() {
    var v1 int32 = 266
    var v2 int8 = -1
    v3 := make(map[int]string)
    var v4 int64 = -9223372036854775808
    var v5 uint = 255
    var v16 uint8 = 127
    v17 := []int{295}
    var v18 int64 = (v4 % (((v4 & v4)) | 1))
    v4 = (v4 / (((v4 | (v4 + 0))) | 1))
    delete(v3, 1)
    var v19 int = 42
    var v20 uint16 = genFn0(("hello" + "hello"), false, 0)
    v21 := GenS1{F0: 0, F1: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
