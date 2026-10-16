package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 uint16; F1 int8 }
func genFn0() bool {
    v4 := []int{0, -9223372036854775808}
    var v5 int32 = -2147483648
    for fi6 := 0; fi6 < 13; fi6++ {
        var v7 uint16 = 987
        var v8 bool = true
        var v9 uint16 = (v7 & (v7 << ((0) & 63)))
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    _ = v4
    _ = v5
    return false
}

func main() {
    var v1 float64 = 0.1
    v2 := make(map[uint32]uint32)
    var v3 uint = 255
    if (false || (42 != 32767)) {
        v3 = 477
    } else {
        v1 = -2.25
        var v10 string = ("hello" + "a")
        _ = v10
    }
    v11 := GenS1{F0: 2, F1: 2147483648}
    fmt.Println(v3, v3)
    v3 = 0
    v1 = v1
    v12 := GenS2{F0: 0, F1: 100}
    switch -9223372036854775808 {
    case -872:
        v2[7] = 2147483647
    case 4294967295:
        fmt.Println(v3)
    default:
        v2[535] = 2147483648
    }
    if true {
        for fi13 := 0; fi13 < 14; fi13++ {
            v3 = v3
            v1 = v1
            v1 = v1
            _ = fi13
        }
        v3 = (v3 << (((v3 << ((v3) & 63))) & 63))
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
}
