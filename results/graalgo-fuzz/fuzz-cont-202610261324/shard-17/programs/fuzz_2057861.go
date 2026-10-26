package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool; F1 int64; F2 int32 }
type GenS3 struct { F0 int64 }

func main() {
    var v1 uint16 = 0
    var v2 int8 = -1
    var v3 int64 = -9223372036854775808
    var v4 int32 = 605
    v5 := make(map[string]int)
    switch -2147483648 {
    case 0:
        v1 -= v1
    case -1:
        v6 := GenS1{F0: true}
        _ = v6
    }
    fmt.Println(v3, v4, v3)
    if (v2 > (v2 * v2)) {
        if (818 > 18446744073709551615) {
            v7 := GenS1{F0: true}
            fmt.Println(v4, v3)
            var v8 int32 = v4
            _ = v7
            _ = v8
        } else {
            v9 := GenS2{F0: true, F1: (v3 % ((v3) | 1)), F2: v4}
            _ = v9
        }
        var v10 int64 = 9223372036854775807
        v11 := []int{32767, 42, -1, 249}
        _ = v10
        _ = v11
    } else {
        v2 = v2
        v4 += 42
        if true {
            v4 += (-(v4))
        } else {
            v3 = -1
            v12 := GenS2{F0: (v4 > 42), F1: 42, F2: -1}
            var v13 float64 = float64((-2.25 / -2.25))
            _ = v12
            _ = v13
        }
    }
    v14 := GenS1{F0: (!(!false))}
    var v15 uint16 = v1
    var v16 uint64 = 2
    switch ("foo bar" + ("foo bar" + "")) {
    case "":
        v17 := GenS3{F0: v3}
        _ = v17
    }
    var v18 uint32 = 581
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
