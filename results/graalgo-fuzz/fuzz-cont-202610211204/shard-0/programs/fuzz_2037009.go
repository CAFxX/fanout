package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 uint64 }
func genFn0() bool {
    var v5 int32 = -2147483648
    fmt.Println(v5)
    if (!true) {
        v5 -= -1
        for fi6 := 0; fi6 < 9; fi6++ {
            var v7 int32 = v5
            v8 := GenS1{F0: fi6, F1: 42, F2: 18446744073709551615}
            _ = fi6
            _ = v7
            _ = v8
        }
        var v9 bool = true
        _ = v9
    }
    v5 -= (v5 + 1)
    _ = v5
    return (42 > 42)
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]string)
    var v3 float32 = 0.0
    v4 := []string{("a" + "foo bar"), ""}
    var v10 int64 = 0
    var v11 int8 = 100
    var v12 int32 = 42
    v10 = v10
    v13 := v1[1]
    if false {
        if len(v4) > 0 {
            v4[0] = ("\n\t" + "\n\t")
        }
        var v14 int16 = 291
        v15 := GenS1{F0: -833, F1: 995, F2: 18446744073709551615}
        _ = v14
        _ = v15
    }
    var v16 uint64 = 508
    switch 9223372036854775807 {
    case 4294967295:
        var v17 int64 = 236
        _ = v17
    case -2147483648:
        v16 = v16
    case 1:
        v12 += (v12 % (((v12 % ((-2147483648) | 1))) | 1))
    default:
        v10 -= v10
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
