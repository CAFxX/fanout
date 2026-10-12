package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 int16; F1 string; F2 string }
type GenS3 struct { F0 bool; F1 int }

func main() {
    var v1 int32 = 726
    var v2 int16 = -195
    var v3 int32 = (^(v1))
    v4 := GenS1{F0: false, F1: (v1 ^ (v1 * 484))}
    var v5 int16 = v2
    fmt.Println(v3, v2)
    var v6 float32 = 1.5
    if false {
        fmt.Println(v5, v6)
        v7 := GenS1{F0: false, F1: v1}
        v5 -= (v2 * (v2 >> ((452) & 63)))
        _ = v7
    } else {
        v8 := GenS1{F0: (!(42 < 4294967295)), F1: (v3 - 1)}
        v2 = 32767
        _ = v8
    }
    if true {
        v9 := GenS2{F0: -389, F1: ("" + "foo bar"), F2: "a"}
        _ = v9
    }
    v1 += (v3 ^ (-(v1)))
    v10 := GenS3{F0: false, F1: 1}
    fmt.Println(v3)
    var v11 bool = (127 != 2)
    var v12 uint8 = 0
    var v13 uint = 255
    v3 -= (v3 / (((v3 >> ((v13) & 63))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
