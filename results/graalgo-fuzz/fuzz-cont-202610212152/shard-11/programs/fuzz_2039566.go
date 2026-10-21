package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16 }
func genFn0() int32 {
    v5 := []int{0, -2147483648, 127, -1}
    var v6 float64 = ((0.0 * (3.14159 - 0.1)) * 3.14159)
    _ = v5
    _ = v6
    return -1
}

func main() {
    var v1 uint8 = 255
    var v2 float64 = -2.25
    var v3 float32 = 1e10
    var v4 uint16 = 32767
    for fi7 := 0; fi7 < 15; fi7++ {
        v8 := GenS1{F0: 18446744073709551615, F1: (fi7 - -987)}
        var v9 int8 = -75
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 uint32 = 0
    v11 := GenS2{F0: -9223372036854775808}
    v3 -= v3
    v1 = 127
    v12 := GenS1{F0: 246, F1: -9223372036854775808}
    v3 -= ((v3 + v3) - (v3 + v3))
    v13 := GenS1{F0: 42, F1: 2147483647}
    v14 := GenS2{F0: 9223372036854775807}
    v15 := GenS1{F0: 574, F1: 32768}
    if false {
        fmt.Println(v4)
    }
    v2 += v2
    var v16 uint = 4294967295
    for fi17 := 0; fi17 < 16; fi17++ {
        fmt.Println(v4, v4)
        _ = fi17
    }
    for fi18 := 0; fi18 < 4; fi18++ {
        var v19 bool = (-2147483648 == 0)
        v20 := GenS3{F0: -874}
        _ = fi18
        _ = v19
        _ = v20
    }
    v21 := GenS3{F0: -32768}
    for fi22 := 0; fi22 < 17; fi22++ {
        v2 -= 0.0
        v23 := GenS1{F0: 2, F1: (fi22 * fi22)}
        var v24 int8 = 42
        _ = fi22
        _ = v23
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
