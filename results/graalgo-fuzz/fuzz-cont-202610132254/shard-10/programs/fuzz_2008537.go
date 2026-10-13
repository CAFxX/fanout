package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 string; F1 int16; F2 uint32 }
func genFn0() uint8 {
    for fi6 := 0; fi6 < 9; fi6++ {
        for fi7 := 0; fi7 < 19; fi7++ {
            var v8 float64 = 0.0
            v8 += v8
            fmt.Println(fi6, fi7, fi6)
            _ = fi7
            _ = v8
        }
        var v9 int32 = 0
        _ = fi6
        _ = v9
    }
    var v10 int64 = 938
    v11 := []int{1, -9223372036854775808, 4294967295}
    _ = v10
    _ = v11
    return 127
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 int = 4294967295
    var v3 int64 = (v1 | (v1 | (v1 + v1)))
    var v4 uint16 = 1
    var v5 uint32 = 727
    fmt.Println(v2, v5, v3)
    v12 := GenS1{F0: 2147483647, F1: (false || false)}
    var v13 int8 = 94
    v1 += -890
    v4 = (^(v4))
    v1 -= v1
    v2 -= v2
    if ((832 < 0) || (0 <= -1)) {
        fmt.Println(v5)
        var v14 uint64 = 4294967295
        _ = v14
    }
    if true {
        switch (("\n\t" + "a") + "foo bar") {
        case "foo bar":
            var v15 string = "\n\t"
            _ = v15
        }
        for fi16 := 0; fi16 < 1; fi16++ {
            v13 -= (^(v13))
            v17 := GenS2{F0: "\n\t", F1: 32767, F2: (v5 - (^(v5)))}
            v18 := []int{fi16}
            _ = fi16
            _ = v17
            _ = v18
        }
    }
    if true {
        var v19 uint8 = genFn0()
        _ = v19
    }
    var v20 bool = ((v4 != v4) || (uint(v5) > 2))
    for fi21 := 0; fi21 < 11; fi21++ {
        fmt.Println(v3, v3)
        v22 := GenS1{F0: 1, F1: true}
        _ = fi21
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
