package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 uint16 = 2
    var v2 int64 = 274
    v3 := make(map[int]uint32)
    v4 := make(map[int]string)
    if true {
        var v5 uint8 = uint8(v1)
        switch 1 {
        case -9223372036854775808:
            v5 = v5
        }
        _ = v5
    } else {
        for fi6 := 0; fi6 < 19; fi6++ {
            fmt.Println(fi6, v1)
            var v7 uint = 0
            _ = fi6
            _ = v7
        }
    }
    if false {
        var v8 int = -1
        _ = v8
    } else {
        var v9 int = 9223372036854775807
        v9 = (v9 << ((1) & 63))
        if (!(!(v1 == (^(v1))))) {
            for fi10 := 0; fi10 < 5; fi10++ {
                v11 := GenS1{F0: -53}
                _ = fi10
                _ = v11
            }
            var v12 float32 = (-2.25 * -2.25)
            _ = len(v3)
            _ = v12
        }
        _ = v9
    }
    v2 -= v2
    var v13 int64 = 0
    delete(v4, 0)
    for fi14 := 0; fi14 < 17; fi14++ {
        fmt.Println(fi14, v13)
        v15 := []int{fi14}
        v1 += v1
        _ = fi14
        _ = v15
    }
    v16 := v4[-588]
    delete(v4, 9223372036854775807)
    var v17 int64 = 876
    v18 := v3[32768]
    var v19 int32 = 1
    fmt.Println(v16)
    v20 := []int{-628}
    for fi21 := 0; fi21 < 4; fi21++ {
        v13 -= v17
        v2 -= (v17 / (((v13 + v2)) | 1))
        _ = fi21
    }
    v17 += (-(v17))
    switch -2147483648 {
    case -502:
        v17 += (v17 >> ((461) & 63))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
