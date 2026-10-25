package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    var v1 uint16 = 32767
    var v2 string = "hello"
    v3 := make(map[uint32]int)
    v4 := make(map[int]bool)
    var v5 uint = 65535
    if false {
        for fi6 := 0; fi6 < 20; fi6++ {
            var v7 int32 = 42
            _ = fi6
            _ = v7
        }
        v8 := GenS1{F0: (v1 % ((v1) | 1))}
        _ = v8
    } else {
        v9 := GenS1{F0: v1}
        switch -1 {
        case 32768:
            var v10 float64 = (((1.5 * 1.5) / 1e10) + float64(3.14159))
            _ = v10
        case -9223372036854775808:
            v11 := GenS1{F0: v1}
            _ = v11
        case -2147483648:
            v12 := []int{127, -562}
            _ = v12
        default:
            switch 127 {
            case -915:
                v13 := GenS1{F0: (v1 << (((v5 / ((v5) | 1))) & 63))}
                _ = v13
            default:
                v14 := GenS1{F0: (v1 | 2)}
                _ = v14
            }
        }
        var v15 uint8 = 127
        _ = v15
        _ = v9
    }
    var v16 uint32 = 0
    var v17 int8 = -1
    var v18 uint32 = 111
    for fi19 := 0; fi19 < 18; fi19++ {
        var v20 uint32 = (v16 * uint32(v5))
        _ = fi19
        _ = v20
    }
    var v21 uint32 = (^(v18))
    v2 = "a"
    v22 := []int{9223372036854775807}
    fmt.Println(v17, v21)
    v5 -= v5
    if true {
        v5 -= 4294967295
    } else {
        v23 := GenS1{F0: 42}
        v24 := v4[887]
        for fi25 := 0; fi25 < 14; fi25++ {
            var v26 uint = 1
            v17 -= -60
            var v27 uint8 = 247
            _ = fi25
            _ = v26
            _ = v27
        }
        _ = v23
        _ = v24
    }
    fmt.Println(v17, v16, v16)
    var v28 int16 = -32768
    var v29 uint = v5
    for fi30 := 0; fi30 < 1; fi30++ {
        v31 := GenS1{F0: v1}
        _ = fi30
        _ = v31
    }
    var v32 uint = 4294967295
    v18 += (^(v16))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
