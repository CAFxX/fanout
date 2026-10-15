package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 bool }

func main() {
    var v1 int32 = -594
    var v2 string = ("a" + ("foo bar" + ("\n\t" + "hello")))
    var v3 uint = 42
    var v4 int8 = -1
    var v5 int = int(v4)
    for fi6 := 0; fi6 < 17; fi6++ {
        var v7 bool = true
        _ = fi6
        _ = v7
    }
    fmt.Println(v2, v5, v3)
    switch v2 {
    case "hello":
        v8 := []int{v5, (^(v5)), (v5 ^ (v5 / ((42) | 1))), v5}
        _ = v8
    default:
        var v9 float32 = 0.1
        _ = v9
    }
    if false {
        v10 := GenS1{F0: 42, F1: true, F2: (v1 <= 2147483647)}
        _ = v10
    } else {
        var v11 int16 = 1000
        var v12 int32 = (v1 >> (((v3 % ((758) | 1))) & 63))
        v4 = (v4 ^ (v4 << ((42) & 63)))
        _ = v11
        _ = v12
    }
    fmt.Println(v5)
    v13 := []int{(v5 & 9223372036854775807), -462}
    v14 := GenS1{F0: 18446744073709551615, F1: true, F2: ((v3 | (^(v3))) > 4294967295)}
    v15 := GenS1{F0: 2, F1: (!(413 >= 2147483648)), F2: (-9223372036854775808 >= -1)}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
