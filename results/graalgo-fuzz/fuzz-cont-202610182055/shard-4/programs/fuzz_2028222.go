package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 uint; F1 string; F2 uint32 }
func genFn0(p0 int32) int16 {
    if (128 < 196) {
        fmt.Println(p0)
        if false {
            if (true && true) {
                var v6 int64 = 677
                var v7 float64 = -2.25
                _ = v6
                _ = v7
            } else {
                p0 = 42
                v8 := GenS1{F0: false, F1: 128}
                var v9 uint64 = 1
                _ = v8
                _ = v9
            }
            fmt.Println(p0)
            var v10 int32 = (p0 ^ 2147483647)
            _ = v10
        }
    } else {
        switch 42 {
        case 9223372036854775807:
            var v11 uint64 = 2
            _ = v11
        }
        p0 += (^(p0))
        v12 := GenS1{F0: true, F1: 2}
        _ = v12
    }
    v13 := []int{-1, -672}
    var v14 uint = 124
    var v15 float32 = 3.14159
    _ = p0
    _ = v13
    _ = v14
    _ = v15
    return 1000
}
func genFn1() int8 {
    if (!true) {
        var v16 int = 32768
        _ = v16
    } else {
        var v17 uint16 = 0
        v18 := GenS2{F0: 2, F1: 463}
        _ = v17
        _ = v18
    }
    v19 := []int{}
    var v20 int16 = 0
    var v21 int8 = 1
    var v22 uint8 = 2
    _ = v19
    _ = v20
    _ = v21
    _ = v22
    return v21
}

func main() {
    var v1 uint16 = 815
    var v2 float64 = (0.0 / 3.14159)
    var v3 int64 = -951
    var v4 float32 = (1e10 / ((1.5 + 1.5) / -2.25))
    v5 := make(map[string]int)
    for fi23 := 0; fi23 < 16; fi23++ {
        var v24 uint8 = 42
        var v25 uint16 = v1
        v26 := GenS1{F0: (-1 != -2147483648), F1: 128}
        _ = fi23
        _ = v24
        _ = v25
        _ = v26
    }
    for fi27 := 0; fi27 < 19; fi27++ {
        var v28 int16 = 290
        _ = fi27
        _ = v28
    }
    fmt.Println(v3)
    v2 = 3.14159
    var v29 uint32 = 65535
    var v30 float64 = 1.5
    for fi31 := 0; fi31 < 19; fi31++ {
        v2 += v30
        fmt.Println(v2, v1, v2)
        _ = fi31
    }
    if false {
        var v32 int64 = v3
        if true {
            for fi33 := 0; fi33 < 8; fi33++ {
                for fi34 := 0; fi34 < 7; fi34++ {
                    fmt.Println(v29, fi33, v4)
                    _ = fi34
                }
                v4 = 1e10
                v35 := GenS3{F0: 65535, F1: ("" + ""), F2: 0}
                _ = fi33
                _ = v35
            }
            v32 -= (-(v32))
            switch "a" {
            case "hello":
                v2 = 3.14159
            default:
                v5["a"] = 1
            }
        }
        v36 := []int{-9223372036854775808, -89}
        _ = v32
        _ = v36
    } else {
        fmt.Println(v2, v30, v30)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
