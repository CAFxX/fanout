package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 uint8; F2 int64 }
func genFn0() int32 {
    var v7 uint16 = 1
    v8 := []int{32768, -1, 2147483647, -909}
    fmt.Println(v7)
    fmt.Println(v7)
    v9 := GenS1{F0: -54}
    _ = v7
    _ = v8
    _ = v9
    return 415
}

func main() {
    var v1 float64 = 0.0
    var v2 uint32 = 0
    var v3 int8 = -24
    var v4 int = 4294967295
    var v5 bool = false
    var v6 float32 = (0.1 * 0.1)
    v10 := GenS2{F0: 65535, F1: 128, F2: 0}
    v2 = (v2 | v2)
    v11 := GenS2{F0: 65535, F1: 222, F2: int64(v3)}
    if false {
        if v5 {
            v5 = false
            v12 := GenS2{F0: 1, F1: 84, F2: 42}
            v13 := []int{}
            _ = v12
            _ = v13
        }
        v6 += (v6 + -2.25)
        v5 = v5
    } else {
        var v14 int64 = -9223372036854775808
        v5 = true
        _ = v14
    }
    v2 -= (v2 + 702)
    var v15 uint8 = 0
    var v16 bool = (0 >= 32767)
    for fi17 := 0; fi17 < 19; fi17++ {
        for fi18 := 0; fi18 < 8; fi18++ {
            v19 := GenS2{F0: 18446744073709551615, F1: 62, F2: -1}
            v20 := GenS2{F0: 18446744073709551615, F1: 128, F2: -1}
            v5 = (v16 || v16)
            _ = fi18
            _ = v19
            _ = v20
        }
        _ = fi17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
