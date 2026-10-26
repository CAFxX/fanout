package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int64 }

func main() {
    var v1 float32 = (1.5 + 0.0)
    var v2 bool = true
    var v3 int64 = -1
    var v4 bool = (v2 || ((-(v3)) >= -76))
    var v5 bool = false
    for fi6 := 0; fi6 < 11; fi6++ {
        v7 := GenS1{F0: 266}
        v4 = (65535 >= 2)
        _ = fi6
        _ = v7
    }
    var v8 float64 = 3.14159
    v4 = v2
    if true {
        var v9 int64 = (v3 >> ((4294967295) & 63))
        _ = v9
    } else {
        v4 = (-74 <= 0)
        v10 := []int{-1, 32767}
        v11 := GenS2{F0: 1}
        _ = v10
        _ = v11
    }
    if (-2.25 > ((3.14159 / ((v8 + v8) / float64(1e10))) * v8)) {
        v12 := GenS1{F0: 6}
        _ = v12
    } else {
        v13 := GenS2{F0: 42}
        v14 := []int{-969}
        _ = len(v14)
        _ = v13
        _ = v14
    }
    v15 := []int{9223372036854775807, 32767, -9223372036854775808}
    switch 32767 {
    case 949:
        fmt.Println(v1, v5, v5)
    case 42:
        var v16 int8 = 82
        _ = v16
    }
    if (0 != 943) {
        switch -988 {
        case -1:
            fmt.Println(v4, v4, v5)
        case -2147483648:
            v17 := GenS3{F0: v3}
            _ = v17
        default:
            if len(v15) > 0 {
                v15[0] = 128
            }
        }
    } else {
        var v18 int16 = -838
        v19 := GenS1{F0: 65535}
        _ = v18
        _ = v19
    }
    v4 = v5
    if len(v15) >= 2 {
        v15 = v15[1:len(v15)]
    }
    for fi20 := 0; fi20 < 11; fi20++ {
        var v21 uint8 = 212
        v15 = append(v15, fi20)
        v21 += (v21 & (^(v21)))
        _ = fi20
        _ = v21
    }
    fmt.Println(v1)
    var v22 int64 = v3
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
