package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 int16 }
type GenS2 struct { F0 uint16 }
func genFn0() uint8 {
    var v4 uint16 = 1
    var v5 int32 = 0
    v6 := GenS1{F0: 42, F1: 18446744073709551615, F2: 32767}
    var v7 string = "a"
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 int64, p1 bool) bool {
    v8 := GenS1{F0: 32767, F1: 823, F2: -754}
    v9 := []int{127}
    if len(v9) > 0 {
        v9[0] = 42
    }
    v10 := GenS2{F0: 32767}
    _ = p0
    _ = p1
    _ = v10
    _ = v8
    _ = v9
    return ((1 >= -7) && ((false || true) && p1))
}

func main() {
    var v1 int32 = -2147483648
    var v2 int64 = -1
    var v3 int8 = 27
    v11 := GenS2{F0: 950}
    fmt.Println(v3)
    v3 += 100
    v12 := GenS2{F0: 1}
    if false {
        var v13 uint = 255
        var v14 int64 = 1
        v3 = -98
        _ = v13
        _ = v14
    } else {
        switch 982 {
        case 42:
            var v15 int16 = -465
            _ = v15
        }
        fmt.Println(v1, v3)
    }
    for fi16 := 0; fi16 < 7; fi16++ {
        v1 = (v1 % (((v1 % ((2147483647) | 1))) | 1))
        _ = fi16
    }
    fmt.Println(v3)
    v17 := GenS1{F0: 1000, F1: 42, F2: 0}
    v18 := GenS1{F0: 0, F1: 374, F2: 32767}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
