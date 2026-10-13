package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 int64 }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 uint64; F1 int32; F2 int16 }

func main() {
    var v1 uint = 593
    var v2 uint64 = 216
    v3 := make(map[string]uint32)
    if (-1 == -9223372036854775808) {
        v3[("" + "a")] = 2147483648
    } else {
        var v4 uint8 = 2
        _ = v4
    }
    var v5 int64 = 9223372036854775807
    v6 := GenS1{F0: (v1 ^ (v1 + (v1 * v1))), F1: (v1 & 774), F2: (v5 % ((v5) | 1))}
    for fi7 := 0; fi7 < 7; fi7++ {
        v8 := []int{}
        v1 = (v1 ^ (v1 - 0))
        v9 := GenS2{F0: "foo bar", F1: 32767}
        _ = fi7
        _ = v8
        _ = v9
    }
    v2 += (v2 << ((v1) & 63))
    if true {
        if true {
            fmt.Println(v1, v2)
            var v10 uint32 = 607
            _ = v10
        } else {
            _ = len(v3)
            v11 := GenS3{F0: (v2 >> ((1) & 63)), F1: -2147483648, F2: 42}
            v2 -= v2
            _ = v11
        }
        v1 = (v1 + v1)
        v12 := GenS1{F0: uint(v2), F1: (v1 ^ (v1 / (((v1 / (((v1 - v1)) | 1))) | 1))), F2: v5}
        _ = v12
    } else {
        for fi13 := 0; fi13 < 2; fi13++ {
            if (!(float64(((((1e10 + -2.25) / -2.25) / (float32(((0.0 * 0.0) - -2.25)) - -2.25)) - 0.0)) != (float64(float32(1.5)) - 0.0))) {
                fmt.Println(v1, fi13)
            }
            _ = fi13
        }
    }
    v14 := []int{-976, -9223372036854775808, 944, 1}
    var v15 bool = ((false && (!false)) || (!(2 != 137)))
    var v16 int32 = 246
    v1 -= v1
    v15 = (v15 || (v15 && (!(!(!(0.1 >= 0.0))))))
    var v17 bool = (!("foo bar" >= ("hello" + "hello")))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v5
    _ = v6
}
