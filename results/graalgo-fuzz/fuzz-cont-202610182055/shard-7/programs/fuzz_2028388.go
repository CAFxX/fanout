package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 uint; F1 bool; F2 string }
func genFn0() int64 {
    if (!true) {
        for fi6 := 0; fi6 < 19; fi6++ {
            var v7 uint8 = 2
            _ = fi6
            _ = v7
        }
        v8 := GenS1{F0: 32768, F1: 127}
        _ = v8
    }
    var v9 int16 = -1
    v10 := GenS2{F0: 337, F1: 18446744073709551615}
    v11 := []int{423, -9223372036854775808, -1, 4294967295}
    if false {
        var v12 int32 = 1
        _ = v12
    }
    _ = v10
    _ = v11
    _ = v9
    return -9223372036854775808
}

func main() {
    var v1 float64 = 1e10
    var v2 uint8 = 128
    var v3 int = 42
    var v4 string = "foo bar"
    var v5 uint8 = v2
    var v13 uint = 18446744073709551615
    fmt.Println(v1)
    fmt.Println(v3, v3, v2)
    v14 := GenS3{F0: (v13 / ((v13) | 1)), F1: (-9223372036854775808 != genFn0()), F2: v4}
    v15 := GenS1{F0: (v3 ^ v3), F1: 42}
    v4 = (v4 + ("\n\t" + v4))
    v13 -= 255
    v2 = (^(v5))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
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
