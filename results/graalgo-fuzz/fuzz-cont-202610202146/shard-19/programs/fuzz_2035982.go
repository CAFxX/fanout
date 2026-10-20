package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 int; F1 int }
func genFn0(p0 uint32) uint64 {
    v7 := GenS1{F0: 2, F1: 284}
    v8 := []int{127, 4294967295, 9223372036854775807, -2147483648}
    v8 = append(v8, -860)
    var v9 bool = (2 != 4294967295)
    v10 := GenS2{F0: -156, F1: 127}
    _ = p0
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 1
}

func main() {
    var v1 uint8 = 127
    var v2 uint32 = 2
    var v3 uint = 1
    var v4 uint = v3
    var v5 int16 = -1
    var v6 uint16 = 0
    v3 -= v3
    v11 := []int{-1}
    for fi12 := 0; fi12 < 2; fi12++ {
        var v13 string = "foo bar"
        fmt.Println(v3)
        if (4294967295 != 1) {
            var v14 int16 = (v5 - v5)
            _ = v14
        } else {
            var v15 int32 = -2147483648
            _ = v15
        }
        _ = fi12
        _ = v13
    }
    var v16 int8 = 47
    v17 := GenS2{F0: -266, F1: 2147483647}
    for fi18 := 0; fi18 < 2; fi18++ {
        v5 += (v5 << ((v3) & 63))
        var v19 bool = ((v1 >= (v1 + v1)) && ((3.14159 / 1e10) == 1.5))
        v1 -= (v1 << ((v3) & 63))
        _ = fi18
        _ = v19
    }
    if false {
        v5 = v5
        v20 := GenS1{F0: v1, F1: v5}
        var v21 string = (("\n\t" + "hello") + ("a" + "foo bar"))
        _ = v20
        _ = v21
    } else {
        v22 := GenS1{F0: v1, F1: 1}
        fmt.Println(v2, v2, v2)
        _ = v22
    }
    v3 = 4294967295
    fmt.Println(v2, v4, v1)
    fmt.Println(v4)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
