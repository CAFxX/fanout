package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 int8 }

func main() {
    var v1 int16 = 42
    var v2 int64 = -1
    v3 := []int{-2147483648, -584, 1, -9223372036854775808}
    v4 := make(map[string]uint32)
    var v5 uint8 = 225
    var v6 uint = 28
    if len(v3) > 0 {
        v3[0] = -2147483648
    }
    switch "hello" {
    case "foo bar":
        var v7 string = "\n\t"
        _ = v7
    case "a":
        _ = len(v3)
    case "\n\t":
        var v8 int16 = (v1 - v1)
        _ = v8
    }
    v2 = (v2 ^ 1)
    _ = len(v4)
    fmt.Println(v2, v2)
    fmt.Println(v1)
    if len(v3) > 0 {
        v3[0] = -9223372036854775808
    }
    for fi9 := 0; fi9 < 19; fi9++ {
        v2 += v2
        var v10 float32 = 3.14159
        v11 := GenS1{F0: ("a" != "\n\t"), F1: (v2 >> ((v6) & 63)), F2: -1}
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 int16 = (v1 << ((627) & 63))
    fmt.Println(v2, v1, v6)
    v13 := GenS1{F0: ((^(v5)) <= v5), F1: (-(v2)), F2: 100}
    v14 := GenS1{F0: false, F1: -125, F2: 16}
    var v15 int16 = (v1 << ((v6) & 63))
    for fi16 := 0; fi16 < 10; fi16++ {
        for fi17 := 0; fi17 < 13; fi17++ {
            v18 := GenS1{F0: false, F1: v2, F2: int8(v1)}
            _ = fi17
            _ = v18
        }
        var v19 uint = (v6 | 479)
        if ((18446744073709551615 <= 42) && false) {
            if (v6 != v6) {
                var v20 int = 1
                v20 = (v20 ^ (v20 - fi16))
                _ = v20
            } else {
                var v21 int32 = 0
                _ = v21
            }
            v22 := GenS1{F0: (!((-(v2)) > v2)), F1: v2, F2: -78}
            var v23 uint64 = 4294967295
            _ = v22
            _ = v23
        }
        _ = fi16
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
