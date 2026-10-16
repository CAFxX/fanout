package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 int32; F1 uint }
type GenS4 struct { F0 bool; F1 uint64; F2 int64 }
func genFn0(p0 uint8, p1 int, p2 int64) bool {
    if false {
        var v7 uint64 = 19
        p1 = (p1 / (((p1 | (p1 / ((p1) | 1)))) | 1))
        p0 -= p0
        _ = v7
    }
    var v8 string = "foo bar"
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return false
}

func main() {
    var v1 string = "a"
    v2 := make(map[string]string)
    var v3 uint = 4294967295
    var v4 bool = (false || (-92 >= -1))
    var v5 uint32 = 786
    v6 := make(map[string]string)
    if (172 < 127) {
        var v9 float32 = -2.25
        v10 := GenS1{F0: 32767}
        _ = v10
        _ = v9
    } else {
        v11 := []int{42, 128, 32768}
        _ = v11
    }
    var v12 int64 = 0
    var v13 float32 = -2.25
    fmt.Println(v4)
    for fi14 := 0; fi14 < 9; fi14++ {
        v15 := GenS1{F0: 32767}
        if true {
            for fi16 := 0; fi16 < 19; fi16++ {
                switch (fi14 % ((32768) | 1)) {
                case -562:
                    v3 = (^(v3))
                }
                var v17 uint8 = 42
                _ = fi16
                _ = v17
            }
        } else {
            v5 -= uint32(v3)
        }
        v4 = v4
        _ = fi14
        _ = v15
    }
    v18 := GenS2{F0: -32768, F1: v4}
    v2["foo bar"] = v1
    _ = len(v6)
    var v19 int8 = 100
    switch (v1 + "foo bar") {
    case "a":
        var v20 uint32 = 42
        _ = v20
    default:
        if v4 {
            switch v1 {
            case "":
                var v21 int = 4294967295
                _ = v21
            case "foo bar":
                for fi22 := 0; fi22 < 3; fi22++ {
                    fmt.Println(fi22)
                    _ = fi22
                }
            case "a":
                var v23 uint = 98
                _ = v23
            default:
                v24 := GenS3{F0: 42, F1: 1}
                _ = v24
            }
            v3 = v3
        }
    }
    var v25 int32 = 1
    v26 := GenS4{F0: (0 >= (v5 % ((v5) | 1))), F1: 0, F2: (v12 | 463)}
    _ = len(v2)
    fmt.Println(v4)
    var v27 uint64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
