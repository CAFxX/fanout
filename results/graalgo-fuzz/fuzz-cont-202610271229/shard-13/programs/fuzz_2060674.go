package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int8; F1 bool; F2 uint32 }
type GenS5 struct { F0 uint64 }
func genFn0(p0 int16, p1 uint32) uint8 {
    if (-1 >= 42) {
        p1 = p1
        p1 -= (p1 - (p1 / ((1) | 1)))
        switch 32767 {
        case -1:
            for fi6 := 0; fi6 < 1; fi6++ {
                p1 = p1
                _ = fi6
            }
        }
    }
    if true {
        var v7 int8 = 100
        var v8 uint32 = 273
        _ = v7
        _ = v8
    }
    p1 = p1
    if (!true) {
        p0 = p0
    } else {
        if (!true) {
            v9 := GenS1{F0: false, F1: 42}
            v10 := []int{-9223372036854775808, -2147483648, -2147483648, 0}
            for fi11 := 0; fi11 < 15; fi11++ {
                var v12 bool = false
                v13 := GenS2{F0: (p0 / ((p0) | 1))}
                p0 = p0
                _ = fi11
                _ = v12
                _ = v13
            }
            _ = v10
            _ = v9
        }
        var v14 int = -9223372036854775808
        _ = v14
    }
    switch 127 {
    case -9223372036854775808:
        v15 := []int{-2147483648, 1}
        _ = v15
    case -1:
        p0 = p0
    case 127:
        switch ("foo bar" + "hello") {
        case "\n\t":
            v16 := GenS3{F0: 556}
            _ = v16
        case "hello":
            v17 := []int{2147483647}
            _ = v17
        case "":
            v18 := []int{-1, 9223372036854775807, int(p1), -2147483648}
            _ = v18
        default:
            for fi19 := 0; fi19 < 17; fi19++ {
                var v20 uint = 2
                var v21 uint = (v20 | (v20 << ((4294967295) & 63)))
                fmt.Println(v21, p1, v20)
                _ = fi19
                _ = v20
                _ = v21
            }
        }
    default:
        v22 := GenS4{F0: -1, F1: false, F2: (^(p1))}
        _ = v22
    }
    _ = p0
    _ = p1
    return uint8(p1)
}

func main() {
    var v1 int16 = 1
    var v2 int = 128
    var v3 int = 127
    var v4 bool = (("" > "\n\t") || false)
    v5 := make(map[string]uint32)
    if (37 < 2) {
        if true {
            var v23 string = "a"
            _ = v23
        }
        if (1 < 1) {
            switch 42 {
            case 4294967295:
                var v24 int32 = 418
                _ = v24
            }
            var v25 int8 = 50
            _ = v25
        }
        v26 := []int{v3, 42, v3}
        _ = v26
    }
    for fi27 := 0; fi27 < 6; fi27++ {
        fmt.Println(v1)
        var v28 int8 = 100
        _ = fi27
        _ = v28
    }
    v29 := GenS5{F0: 2}
    var v30 uint = 4294967295
    v31 := v5["a"]
    v32 := []int{}
    if ((v31 - v31) < (v31 >> ((v30) & 63))) {
        fmt.Println(v3)
        if len(v32) > 0 {
            v32[0] = v2
        }
    } else {
        var v33 uint = v30
        _ = v33
    }
    var v34 float32 = 3.14159
    if v4 {
        v35 := GenS2{F0: v1}
        var v36 int64 = 9223372036854775807
        fmt.Println(v4, v34)
        _ = v35
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v34
    _ = v4
    _ = v5
}
