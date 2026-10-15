package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool }
func genFn0() uint16 {
    v4 := []int{}
    if true {
        v5 := GenS1{F0: 18446744073709551615}
        _ = v5
    } else {
        var v6 int16 = -32768
        var v7 int64 = 0
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 8; fi8++ {
        var v9 uint16 = 229
        _ = fi8
        _ = v9
    }
    if true {
        var v10 bool = false
        var v11 uint8 = 234
        v12 := GenS1{F0: 334}
        _ = v10
        _ = v11
        _ = v12
    }
    v13 := GenS1{F0: 630}
    _ = v13
    _ = v4
    return 65535
}

func main() {
    v1 := []uint32{65535, 665, 42}
    v2 := []uint32{}
    v3 := []int{}
    var v14 float32 = -2.25
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v15 uint16 = 65535
    v14 += 3.14159
    if false {
        v16 := GenS2{F0: false}
        var v17 int32 = -2147483648
        fmt.Println(v15, v17)
        _ = v16
        _ = v17
    } else {
        switch 42 {
        case 0:
            v14 = v14
        case 730:
            v14 += v14
        default:
            v15 = (v15 << ((327) & 63))
        }
        switch 1 {
        case 499:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case -1:
            var v18 uint8 = uint8(v15)
            _ = v18
        }
        if (v14 != v14) {
            v19 := GenS2{F0: (((v15 ^ v15) <= genFn0()) && (127 < 2147483647))}
            _ = v19
        }
    }
    fmt.Println(v14, v15)
    var v20 uint16 = (v15 * v15)
    var v21 int8 = 109
    fmt.Println(v20, v20, v20)
    if len(v2) > 0 {
        v2[0] = 690
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
