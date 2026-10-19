package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 string; F1 int16; F2 int64 }

func main() {
    var v1 uint32 = 485
    v2 := []int{32767, 4294967295, -9223372036854775808, 4294967295}
    var v3 float64 = float64((3.14159 / (-2.25 - (0.0 / (-2.25 * 1e10)))))
    var v4 int64 = 552
    v5 := GenS1{F0: 18446744073709551615, F1: 65535}
    v3 -= v3
    v2 = append(v2, 0)
    v2 = append(v2, 0)
    if false {
        if true {
            var v6 int64 = (v4 ^ (v4 ^ (^(v4))))
            var v7 int32 = -376
            _ = v6
            _ = v7
        } else {
            var v8 int = int(v1)
            var v9 int32 = int32(v8)
            v2 = append(v2, -988)
            _ = v8
            _ = v9
        }
        _ = len(v2)
        fmt.Println(v1, v3, v1)
    }
    if true {
        v2 = append(v2, 4294967295)
        if ((!(42 >= 573)) && (9223372036854775807 != 274)) {
            v10 := GenS1{F0: 18446744073709551615, F1: 1}
            _ = v10
        } else {
            v11 := GenS2{F0: "hello", F1: -988, F2: 0}
            var v12 string = ""
            _ = v11
            _ = v12
        }
        var v13 string = "a"
        _ = v13
    }
    _ = len(v2)
    var v14 int64 = (v4 >> ((42) & 63))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
