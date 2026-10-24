package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 bool; F1 int16; F2 uint32 }
func genFn0() uint64 {
    v4 := []int{2147483647, 127}
    v5 := GenS1{F0: 4294967295}
    var v6 int32 = 2147483647
    for fi7 := 0; fi7 < 1; fi7++ {
        v8 := GenS2{F0: 42, F1: 388}
        fmt.Println(fi7)
        _ = fi7
        _ = v8
    }
    _ = v4
    _ = v5
    _ = v6
    return 0
}

func main() {
    var v1 int8 = -99
    var v2 int32 = 397
    var v3 bool = ((true || true) || true)
    v9 := GenS3{F0: v3, F1: 1000, F2: 4294967295}
    if v3 {
        var v10 int64 = -1
        for fi11 := 0; fi11 < 15; fi11++ {
            var v12 int64 = v10
            _ = fi11
            _ = v12
        }
        _ = v10
    } else {
        var v13 float64 = 3.14159
        var v14 float32 = 1e10
        _ = v13
        _ = v14
    }
    var v15 int64 = 42
    v3 = v3
    var v16 int = 4294967295
    var v17 uint16 = 518
    if (true || (42 != 0)) {
        v1 += -128
        switch (^(v16)) {
        case 9223372036854775807:
            v2 = v2
        case -2147483648:
            v3 = true
        case 32768:
            var v18 uint16 = 42
            _ = v18
        default:
            if true {
                v19 := []int{v16}
                v1 -= v1
                _ = v19
            }
        }
        v3 = true
    }
    var v20 uint32 = 657
    switch v16 {
    case 0:
        v21 := []int{(v16 ^ v16)}
        _ = v21
    default:
        if ((1e10 * 1e10) >= 1.5) {
            v20 += (v20 % ((v20) | 1))
            var v22 string = ("hello" + ("" + (("a" + "foo bar") + (("hello" + ("\n\t" + "foo bar")) + "foo bar"))))
            fmt.Println(v17, v17)
            _ = v22
        } else {
            v23 := GenS1{F0: v16}
            var v24 string = ""
            v16 -= (v16 ^ (-(v16)))
            _ = v23
            _ = v24
        }
    }
    var v25 uint64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v25
    _ = v3
    _ = v9
}
