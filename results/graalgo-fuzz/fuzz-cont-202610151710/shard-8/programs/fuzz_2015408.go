package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int64; F1 uint32 }
type GenS5 struct { F0 uint8; F1 int }
type GenS6 struct { F0 string; F1 uint16 }
func genFn0() int8 {
    var v6 string = "a"
    if false {
        if (65535 <= 42) {
            v6 = (v6 + "")
            var v7 uint = 255
            _ = v7
        } else {
            v8 := []int{32767, -202, 128}
            _ = v8
        }
        var v9 uint64 = 1
        v10 := GenS1{F0: -1, F1: 223}
        _ = v10
        _ = v9
    }
    _ = v6
    return -10
}

func main() {
    var v1 uint64 = 495
    v2 := []int{-457, 128}
    v3 := []string{""}
    v4 := make(map[int]string)
    var v5 uint64 = (v1 + v1)
    v11 := GenS2{F0: 243, F1: v5}
    v12 := GenS3{F0: 311}
    if (0.0 <= (float32(((1.5 + ((-2.25 * 1e10) / ((-2.25 / (1.5 / 1.0)) + 3.14159))) - (float64(((3.14159 + (((1e10 * -2.25) * float32(float64(-2.25))) * 0.1)) / ((float32(0.1) - 1.5) / ((-2.25 - 1.5) - 1.5)))) - 1e10))) - 0.0)) {
        if false {
            v1 = (v5 ^ (^(v1)))
        } else {
            var v13 int16 = -1
            _ = v13
        }
        if ((3.14159 >= 1e10) && ("foo bar" < "a")) {
            v14 := GenS4{F0: 9223372036854775807, F1: 65535}
            _ = v14
        }
    }
    fmt.Println(v5, v5)
    var v15 int64 = -9223372036854775808
    for fi16 := 0; fi16 < 15; fi16++ {
        v17 := GenS1{F0: v15, F1: 42}
        var v18 int16 = 32767
        _ = fi16
        _ = v17
        _ = v18
    }
    for fi19 := 0; fi19 < 14; fi19++ {
        var v20 uint32 = 1
        fmt.Println(v1)
        var v21 int32 = -149
        _ = fi19
        _ = v20
        _ = v21
    }
    if (1000 < -32768) {
        v5 = v1
        if false {
            v15 = (v15 / ((v15) | 1))
            if (false || true) {
                fmt.Println(v1, v1, v5)
                v5 += (v5 + v1)
                v22 := GenS1{F0: 9223372036854775807, F1: 255}
                _ = v22
            } else {
                var v23 int = 127
                v15 += 915
                _ = v23
            }
            v24 := GenS5{F0: 42, F1: 2147483647}
            _ = v24
        }
        for fi25 := 0; fi25 < 15; fi25++ {
            v26 := GenS6{F0: "foo bar", F1: 283}
            var v27 int = (fi25 << ((0) & 63))
            _ = fi25
            _ = v26
            _ = v27
        }
    } else {
        for fi28 := 0; fi28 < 18; fi28++ {
            v29 := GenS1{F0: int64(fi28), F1: 795}
            var v30 int32 = 42
            var v31 int8 = int8(v1)
            _ = fi28
            _ = v29
            _ = v30
            _ = v31
        }
        if true {
            var v32 int8 = 1
            var v33 int16 = 547
            _ = v32
            _ = v33
        }
        v34 := GenS2{F0: 11, F1: v5}
        _ = v34
    }
    v15 += (v15 << ((0) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
