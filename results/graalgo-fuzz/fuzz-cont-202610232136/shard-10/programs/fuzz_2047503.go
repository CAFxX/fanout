package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 int32 }
type GenS2 struct { F0 int32; F1 int; F2 bool }

func main() {
    var v1 string = "a"
    var v2 int16 = -929
    var v3 int8 = -107
    var v4 uint64 = 4294967295
    var v5 uint = 1000
    var v6 uint = v5
    for fi7 := 0; fi7 < 17; fi7++ {
        v8 := []int{fi7, (-(fi7)), fi7, fi7}
        _ = fi7
        _ = v8
    }
    var v9 int8 = (v3 & -58)
    var v10 int8 = (-(v9))
    var v11 int32 = 337
    switch 9223372036854775807 {
    case 0:
        var v12 int16 = v2
        _ = v12
    case 42:
        for fi13 := 0; fi13 < 11; fi13++ {
            v14 := []int{}
            var v15 float64 = 0.1
            _ = fi13
            _ = v14
            _ = v15
        }
    case 9223372036854775807:
        if (!true) {
            fmt.Println(v11, v4, v5)
            v16 := GenS1{F0: -1, F1: 42, F2: v11}
            _ = v16
        }
    }
    v6 = (v6 * v6)
    if false {
        v17 := []int{863}
        _ = v17
    } else {
        switch 2147483647 {
        case -750:
            if false {
                fmt.Println(v11)
                v18 := GenS2{F0: (v11 + v11), F1: 943, F2: true}
                _ = v18
            }
        case 42:
            var v19 float32 = ((0.0 * -2.25) / (((((1.5 * 1.5) - 0.0) * 1e10) / (1e10 - 3.14159)) * 1.5))
            _ = v19
        case -2147483648:
            var v20 uint16 = 862
            _ = v20
        }
        var v21 uint8 = 255
        v22 := []int{-1}
        _ = v21
        _ = v22
    }
    var v23 int16 = v2
    fmt.Println(v11)
    v9 -= v9
    if (1 != 65535) {
        v6 = v6
    } else {
        fmt.Println(v6)
    }
    for fi24 := 0; fi24 < 6; fi24++ {
        fmt.Println(v3)
        for fi25 := 0; fi25 < 10; fi25++ {
            var v26 int = fi24
            var v27 uint16 = uint16(v5)
            fmt.Println(v2)
            var v28 bool = (v3 == 1)
            var v29 uint16 = 1
            var v30 int8 = (v9 % (((v3 >> ((v6) & 63))) | 1))
            _ = fi25
            _ = v26
            _ = v27
            _ = v28
            _ = v29
            _ = v30
        }
        var v31 uint8 = 127
        _ = fi24
        _ = v31
    }
    var v32 float64 = ((0.1 / 1.0) + (0.0 + ((0.1 * float64(0.0)) / 0.1)))
    v23 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v23
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
