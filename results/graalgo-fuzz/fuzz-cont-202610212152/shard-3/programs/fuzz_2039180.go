package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0() int {
    v4 := []int{2147483647, 127, 0}
    v4 = append(v4, 4294967295)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if false {
        var v5 int = 9223372036854775807
        _ = v5
    } else {
        switch 4294967295 {
        case -1:
            var v6 int8 = -128
            _ = v6
        case -808:
            _ = len(v4)
        default:
            var v7 string = "hello"
            _ = v7
        }
        var v8 float32 = (1e10 / float32(1.5))
        _ = v8
    }
    if false {
        var v9 int16 = -32768
        for fi10 := 0; fi10 < 3; fi10++ {
            v9 = -32768
            v9 = (v9 - (-(v9)))
            _ = fi10
        }
        v9 = (-(v9))
        _ = v9
    }
    _ = v4
    return 2147483647
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []int{-468, 9223372036854775807, -49}
    var v3 bool = (((-9223372036854775808 == 127) || (100 < -128)) || ((-2147483648 <= 817) || true))
    v11 := GenS1{F0: 1}
    var v12 int32 = -797
    v13 := GenS1{F0: 128}
    fmt.Println(v3)
    v3 = (423 > 0)
    v12 -= (v12 >> ((42) & 63))
    v14 := GenS1{F0: 217}
    var v15 int32 = 656
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
}
