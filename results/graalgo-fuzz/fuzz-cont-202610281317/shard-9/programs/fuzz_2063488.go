package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 int64 }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 uint64; F1 uint }
func genFn0() uint64 {
    var v5 uint32 = 4294967295
    v6 := GenS1{F0: -32768, F1: true, F2: 42}
    v7 := []int{42, 42, 9223372036854775807}
    var v8 int32 = 2147483647
    if false {
        var v9 int16 = 32767
        v7 = append(v7, 42)
        v7 = append(v7, int(v9))
        _ = v9
    } else {
        var v10 uint16 = 980
        var v11 float64 = (3.14159 / 1.0)
        _ = v10
        _ = v11
    }
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 467
}

func main() {
    var v1 string = "\n\t"
    var v2 uint16 = 42
    v3 := make(map[int]int)
    var v4 float64 = 0.0
    fmt.Println(v1, v4, v4)
    v1 = (((v1 + "\n\t") + (v1 + (v1 + v1))) + "hello")
    v1 = "foo bar"
    v4 = v4
    for fi12 := 0; fi12 < 17; fi12++ {
        if false {
            v13 := []int{}
            var v14 uint64 = 2
            _ = v13
            _ = v14
        }
        _ = fi12
    }
    v4 -= v4
    if (v1 < v1) {
        fmt.Println(v2, v1)
    } else {
        _ = len(v3)
    }
    if true {
        if false {
            var v15 int64 = 9223372036854775807
            v2 = (v2 | (v2 / ((v2) | 1)))
            v1 = v1
            _ = v15
        } else {
            var v16 float64 = 1e10
            _ = v16
        }
        var v17 uint64 = 42
        v17 = (v17 >> ((uint(v2)) & 63))
        _ = v17
    } else {
        fmt.Println(v2, v2)
        if (896 == 42) {
            v2 = (v2 & v2)
            fmt.Println(v4, v1)
        } else {
            _ = len(v3)
            fmt.Println(v4, v2, v2)
            v4 += v4
        }
    }
    v1 = ""
    var v18 int = 9223372036854775807
    var v19 int = 128
    if true {
        var v20 float64 = (float64((-2.25 + 3.14159)) / v4)
        var v21 int16 = 32767
        _ = v20
        _ = v21
    }
    v3[(^(v18))] = v19
    if false {
        v22 := GenS1{F0: -32768, F1: (v4 >= v4), F2: -1}
        var v23 uint64 = 4294967295
        _ = v22
        _ = v23
    }
    for fi24 := 0; fi24 < 20; fi24++ {
        v4 += v4
        v2 = v2
        v1 = ("\n\t" + v1)
        _ = fi24
    }
    for fi25 := 0; fi25 < 8; fi25++ {
        v4 = v4
        v26 := GenS2{F0: -89, F1: -9223372036854775808}
        v27 := GenS3{F0: 4294967295, F1: 974}
        _ = fi25
        _ = v26
        _ = v27
    }
    v2 -= 415
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
