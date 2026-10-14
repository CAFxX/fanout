package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 uint64 }
type GenS2 struct { F0 uint; F1 int16; F2 uint32 }

func main() {
    var v1 uint16 = 1
    var v2 uint8 = 127
    var v3 int8 = 42
    var v4 bool = (uint64(v1) >= 2)
    v5 := make(map[string]string)
    for fi6 := 0; fi6 < 15; fi6++ {
        var v7 int16 = 0
        var v8 string = "a"
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 float32 = (1.5 / 3.14159)
    v10 := []int{1, -614, 626, 32767}
    switch int(v2) {
    case 127:
        var v11 int32 = -412
        _ = v11
    }
    if ((v3 / (((v3 >> ((2) & 63))) | 1)) >= (v3 + (v3 & 42))) {
        v12 := GenS1{F0: 633, F1: 1, F2: 42}
        if v4 {
            var v13 int64 = 712
            _ = len(v10)
            for fi14 := 0; fi14 < 19; fi14++ {
                fmt.Println(v9, v4)
                _ = fi14
            }
            _ = v13
        }
        _ = v12
    } else {
        v2 -= 255
        for fi15 := 0; fi15 < 5; fi15++ {
            for fi16 := 0; fi16 < 5; fi16++ {
                fmt.Println(fi16)
                var v17 bool = (v3 == v3)
                if (v9 < 1e10) {
                    v18 := v5["a"]
                    _ = len(v10)
                    var v19 float32 = v9
                    _ = v18
                    _ = v19
                } else {
                    var v20 bool = v4
                    _ = v20
                }
                _ = fi16
                _ = v17
            }
            var v21 string = "foo bar"
            _ = fi15
            _ = v21
        }
        var v22 uint64 = 0
        _ = v22
    }
    for fi23 := 0; fi23 < 3; fi23++ {
        var v24 uint8 = v2
        if (v4 || (!((65535 < 373) || (-2.25 <= 0.0)))) {
            var v25 uint = 18446744073709551615
            fmt.Println(v25)
            _ = v25
        } else {
            fmt.Println(v2)
            fmt.Println(v9)
            v26 := GenS2{F0: 2, F1: 231, F2: 65535}
            _ = v26
        }
        _ = fi23
        _ = v24
    }
    var v27 uint16 = v1
    fmt.Println(v2)
    for fi28 := 0; fi28 < 11; fi28++ {
        for fi29 := 0; fi29 < 11; fi29++ {
            v30 := GenS2{F0: uint(v2), F1: 1, F2: 1}
            _ = fi29
            _ = v30
        }
        _ = fi28
    }
    var v31 int16 = 0
    v2 += v2
    v31 -= v31
    if (v9 > v9) {
        if v4 {
            fmt.Println(v31, v3)
            v32 := GenS1{F0: 220, F1: 2, F2: 0}
            _ = v32
        }
        if len(v10) > 0 {
            v10[0] = 1
        }
    } else {
        fmt.Println(v2)
        _ = len(v5)
    }
    switch 42 {
    case 975:
        v1 = 708
    case 877:
        v33 := GenS2{F0: 137, F1: (v31 - (v31 | (v31 / ((-1) | 1)))), F2: 691}
        _ = v33
    case -1:
        fmt.Println(v31, v4)
    default:
        _ = len(v10)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v27
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v9
}
