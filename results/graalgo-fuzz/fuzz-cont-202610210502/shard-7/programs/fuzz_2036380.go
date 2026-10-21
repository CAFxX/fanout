package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int64; F1 int16 }

func main() {
    v1 := []int{-245, -575, 1, 542}
    v2 := make(map[string]string)
    var v3 bool = false
    var v4 float64 = float64(-2.25)
    var v5 string = "a"
    var v6 int64 = 729
    if (!v3) {
        var v7 int = -1
        v8 := v2["\n\t"]
        if (!false) {
            switch (v7 * v7) {
            case 816:
                fmt.Println(v8, v7, v8)
            case 839:
                if v3 {
                    if false {
                        v9 := GenS1{F0: 255, F1: (!(v5 < "a")), F2: 85}
                        _ = v9
                    }
                    v8 = "foo bar"
                } else {
                    if len(v1) > 0 {
                        v1[0] = (v7 / ((-9223372036854775808) | 1))
                    }
                }
            }
            var v10 int64 = 1
            var v11 string = (v8 + "hello")
            _ = v10
            _ = v11
        } else {
            switch (-(v7)) {
            case -9223372036854775808:
                var v12 float32 = 0.1
                _ = v12
            default:
                v13 := GenS2{F0: "\n\t"}
                _ = v13
            }
            var v14 uint = 255
            v15 := GenS3{F0: 4294967295}
            _ = v14
            _ = v15
        }
        _ = v7
        _ = v8
    }
    if v3 {
        v4 += v4
    }
    var v16 int16 = 0
    v16 = 1
    var v17 int16 = (v16 + v16)
    v18 := GenS4{F0: 4294967295}
    v4 = -2.25
    v19 := GenS3{F0: 18446744073709551615}
    var v20 bool = v3
    v21 := GenS5{F0: 0, F1: v17}
    fmt.Println(v3, v6, v20)
    v6 += (v6 >> ((65535) & 63))
    v1 = append(v1, 0)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
