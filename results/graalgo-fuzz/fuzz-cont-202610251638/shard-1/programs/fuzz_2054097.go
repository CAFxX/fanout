package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 uint64 }
type GenS2 struct { F0 uint; F1 int8; F2 int64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint16 }
func genFn0(p0 bool) int64 {
    v5 := GenS1{F0: 1, F1: 477, F2: 42}
    for fi6 := 0; fi6 < 18; fi6++ {
        switch "foo bar" {
        case "foo bar":
            var v7 float32 = 3.14159
            _ = v7
        }
        p0 = (p0 || true)
        _ = fi6
    }
    _ = p0
    _ = v5
    return -1
}
func genFn1() int16 {
    if true {
        var v8 string = "a"
        var v9 int16 = 1
        v10 := GenS2{F0: 0, F1: -1, F2: -9223372036854775808}
        _ = v10
        _ = v8
        _ = v9
    } else {
        v11 := GenS2{F0: 147, F1: 127, F2: -9223372036854775808}
        v12 := []int{752}
        var v13 bool = false
        _ = v11
        _ = v12
        _ = v13
    }
    var v14 int16 = 42
    v15 := GenS2{F0: 2, F1: 100, F2: 1}
    var v16 int32 = -2147483648
    _ = v14
    _ = v15
    _ = v16
    return v14
}

func main() {
    var v1 uint32 = 2
    var v2 float32 = (3.14159 + 0.0)
    v3 := make(map[string]bool)
    v4 := []uint32{(v1 * (v1 / ((v1) | 1))), v1, 42}
    var v17 uint8 = 127
    var v18 uint64 = 1
    v19 := GenS2{F0: 2, F1: 127, F2: 0}
    v20 := GenS3{F0: 65535}
    if (!false) {
        fmt.Println(v1, v18, v17)
        fmt.Println(v2, v18, v18)
        v21 := GenS4{F0: 42}
        _ = v21
    } else {
        v18 = (v18 << ((945) & 63))
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    v22 := GenS5{F0: 0}
    switch "\n\t" {
    case "\n\t":
        var v23 float32 = v2
        _ = v23
    case "foo bar":
        v2 = 3.14159
    }
    v24 := GenS5{F0: 0}
    var v25 float64 = 0.1
    v26 := GenS3{F0: 0}
    var v27 float64 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
