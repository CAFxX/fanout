package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 int32 = 1
    v2 := make(map[int]string)
    var v3 int16 = -682
    v4 := make(map[string]string)
    v5 := make(map[uint32]uint32)
    fmt.Println(v1, v1)
    for fi6 := 0; fi6 < 11; fi6++ {
        if true {
            var v7 uint32 = 42
            for fi8 := 0; fi8 < 14; fi8++ {
                var v9 uint8 = 201
                _ = fi8
                _ = v9
            }
            _ = v7
        }
        var v10 int = (fi6 % (((fi6 << ((353) & 63))) | 1))
        _ = fi6
        _ = v10
    }
    delete(v5, 499)
    var v11 uint64 = 4294967295
    for fi12 := 0; fi12 < 3; fi12++ {
        var v13 uint16 = 466
        v14 := GenS1{F0: 255}
        _ = fi12
        _ = v13
        _ = v14
    }
    if false {
        if true {
            fmt.Println(v1, v1, v11)
            var v15 int8 = 100
            for fi16 := 0; fi16 < 1; fi16++ {
                v1 = (v1 | v1)
                fmt.Println(v11, v15, fi16)
                _ = fi16
            }
            _ = v15
        } else {
            for fi17 := 0; fi17 < 2; fi17++ {
                var v18 float64 = (0.1 * 0.0)
                v11 = v11
                _ = fi17
                _ = v18
            }
            delete(v2, 0)
            var v19 int8 = 127
            _ = v19
        }
    }
    if false {
        var v20 uint16 = 70
        var v21 bool = (!(!(127 < -128)))
        v21 = false
        _ = v20
        _ = v21
    } else {
        fmt.Println(v3, v11, v1)
        v3 -= 42
        if (42 != 1) {
            if true {
                v22 := []int{0, 828, -2147483648}
                var v23 int32 = 0
                _ = v22
                _ = v23
            }
            var v24 float32 = ((((3.14159 * (0.1 / 1.5)) * -2.25) * 1e10) - 3.14159)
            _ = v24
        } else {
            fmt.Println(v11)
            fmt.Println(v3)
            v1 = -251
        }
    }
    var v25 uint8 = 42
    v26 := []int{32767}
    fmt.Println(v11)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
