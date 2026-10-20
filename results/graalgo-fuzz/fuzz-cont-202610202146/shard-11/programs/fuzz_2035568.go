package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 uint16; F1 uint; F2 uint16 }
func genFn0(p0 int32, p1 uint32) bool {
    if false {
        fmt.Println(p0, p0)
        v6 := []int{}
        for fi7 := 0; fi7 < 18; fi7++ {
            switch "foo bar" {
            case "foo bar":
                p1 -= p1
            default:
                var v8 float64 = 1e10
                _ = v8
            }
            v9 := GenS1{F0: 42, F1: false}
            _ = fi7
            _ = v9
        }
        _ = v6
    }
    var v10 string = "\n\t"
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    _ = v10
    return true
}

func main() {
    var v1 uint = 18446744073709551615
    v2 := make(map[int]string)
    v3 := make(map[int]int)
    var v4 float32 = 0.1
    v5 := make(map[int]uint32)
    if false {
        v1 += v1
        v1 = (v1 + (v1 & v1))
    }
    v11 := GenS1{F0: -9223372036854775808, F1: true}
    v4 = 0.0
    for fi12 := 0; fi12 < 15; fi12++ {
        v4 = (v4 * 1.5)
        delete(v5, fi12)
        for fi13 := 0; fi13 < 13; fi13++ {
            var v14 uint32 = 65535
            _ = fi13
            _ = v14
        }
        _ = fi12
    }
    var v15 int8 = -1
    var v16 uint32 = 2147483647
    v5[9223372036854775807] = v16
    var v17 float32 = (v4 + (v4 + (v4 + v4)))
    v18 := []int{4294967295, -704, 0, -248}
    var v19 int16 = 1
    var v20 uint16 = 0
    delete(v5, -908)
    var v21 uint = v1
    v22 := GenS1{F0: -9223372036854775808, F1: false}
    for fi23 := 0; fi23 < 20; fi23++ {
        if true {
            v16 = v16
        }
        v17 += float32((0.1 - 3.14159))
        _ = fi23
    }
    if false {
        for fi24 := 0; fi24 < 19; fi24++ {
            v19 += (-(v19))
            for fi25 := 0; fi25 < 2; fi25++ {
                v26 := GenS2{F0: v20, F1: (v21 >> (((v1 & 2)) & 63)), F2: uint16(v1)}
                v27 := GenS1{F0: 0, F1: (1 > 0)}
                _ = fi25
                _ = v26
                _ = v27
            }
            var v28 uint = (v1 ^ 1)
            _ = fi24
            _ = v28
        }
        if true {
            var v29 int8 = (v15 * 27)
            v30 := GenS1{F0: 9223372036854775807, F1: (false && false)}
            v19 = -32768
            _ = v29
            _ = v30
        } else {
            v1 += v1
            if (42 != -618) {
                v31 := GenS1{F0: 661, F1: false}
                v20 = (v20 * v20)
                _ = v31
            }
        }
    }
    fmt.Println(v16, v17, v19)
    v21 = (v21 & 4294967295)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
