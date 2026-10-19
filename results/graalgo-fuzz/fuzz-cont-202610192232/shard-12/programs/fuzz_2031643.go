package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string }

func main() {
    var v1 uint32 = 4294967295
    var v2 uint64 = 42
    var v3 int32 = 953
    if ("hello" == "a") {
        v2 -= (v2 * (v2 << ((171) & 63)))
        v4 := GenS1{F0: v3, F1: true}
        _ = v4
    }
    var v5 float32 = 3.14159
    if (false || (-128 >= 42)) {
        v6 := GenS1{F0: v3, F1: (!((!false) || true))}
        fmt.Println(v5)
        v7 := []int{-9223372036854775808, 0, -9223372036854775808, 4294967295}
        _ = v6
        _ = v7
    } else {
        var v8 int16 = 1000
        v9 := []int{-1, 32767}
        fmt.Println(v5, v1)
        _ = v8
        _ = v9
    }
    var v10 int32 = v3
    v11 := GenS1{F0: (v10 ^ int32(v2)), F1: (54 > 1)}
    switch -1 {
    case 32768:
        v12 := GenS1{F0: v10, F1: false}
        _ = v12
    case 32767:
        var v13 uint8 = 80
        _ = v13
    case 4294967295:
        var v14 uint16 = 1
        _ = v14
    default:
        v15 := []int{127}
        _ = v15
    }
    for fi16 := 0; fi16 < 5; fi16++ {
        v17 := GenS2{F0: v1}
        v3 = (^(v3))
        _ = fi16
        _ = v17
    }
    v18 := GenS3{F0: ("a" + ("hello" + "foo bar"))}
    var v19 bool = (v1 <= (^(v1)))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v5
}
