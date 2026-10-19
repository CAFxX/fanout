package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 bool }
func genFn0() uint32 {
    var v6 uint32 = 4294967295
    var v7 int8 = 100
    var v8 bool = ((v6 < (v6 * (v6 + (v6 % ((885) | 1))))) && true)
    v9 := GenS1{F0: false, F1: 4294967295, F2: v8}
    v10 := []int{927, -9223372036854775808}
    _ = v10
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 722
}

func main() {
    var v1 int8 = -22
    var v2 float64 = float64(((0.1 - float32(-2.25)) * (float32(3.14159) - 1.5)))
    var v3 int = 20
    var v4 uint32 = 205
    v5 := make(map[string]string)
    if false {
        var v11 int = -512
        v12 := GenS1{F0: (!false), F1: 2147483647, F2: (4294967295 != 963)}
        _ = v11
        _ = v12
    }
    v1 = 23
    fmt.Println(v2)
    for fi13 := 0; fi13 < 13; fi13++ {
        var v14 bool = (false || true)
        _ = fi13
        _ = v14
    }
    var v15 uint64 = 4294967295
    for fi16 := 0; fi16 < 3; fi16++ {
        v3 = (v3 << ((uint(v15)) & 63))
        _ = fi16
    }
    v17 := GenS1{F0: ((!true) || true), F1: (v3 - 4294967295), F2: false}
    for fi18 := 0; fi18 < 20; fi18++ {
        var v19 float32 = (0.0 + 1.5)
        _ = fi18
        _ = v19
    }
    v20 := GenS1{F0: (("hello" != "a") && true), F1: v3, F2: true}
    v4 += (v4 + 4294967295)
    for fi21 := 0; fi21 < 19; fi21++ {
        var v22 bool = ((-9223372036854775808 <= 0) && false)
        v23 := GenS1{F0: true, F1: (fi21 & 4294967295), F2: ((1.5 != v2) || (true && v22))}
        v1 += 100
        _ = fi21
        _ = v22
        _ = v23
    }
    var v24 int = v3
    v25 := GenS1{F0: ((v2 - 0.1) >= v2), F1: v24, F2: (42 == 2147483647)}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v17
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
