package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16 }

func main() {
    v1 := []string{"hello"}
    v2 := []uint32{2147483648, 2}
    v3 := []int{9223372036854775807, 128, -2147483648, -2147483648}
    v4 := make(map[int]string)
    var v5 bool = true
    if v5 {
        v5 = false
        v4[-9223372036854775808] = ("foo bar" + "foo bar")
    }
    var v6 float64 = -2.25
    if (false && (4294967295 == 654)) {
        var v7 string = ""
        switch -807 {
        case 32767:
            v8 := GenS1{F0: 2147483647, F1: (((v7 + (v7 + v7)) + v7) > v7)}
            _ = v8
        default:
            var v9 float64 = 1.5
            _ = v9
        }
        var v10 bool = (605 != 1)
        _ = v10
        _ = v7
    } else {
        for fi11 := 0; fi11 < 1; fi11++ {
            fmt.Println(v5, v6, v5)
            _ = fi11
        }
    }
    var v12 int16 = 0
    if false {
        if (-1 <= 1) {
            v13 := GenS2{F0: 42}
            v14 := GenS1{F0: 4294967295, F1: v5}
            _ = v13
            _ = v14
        } else {
            var v15 uint64 = 2
            v12 -= (v12 << ((989) & 63))
            v16 := GenS1{F0: 246, F1: (!v5)}
            _ = v15
            _ = v16
        }
    } else {
        var v17 uint32 = 286
        if (0 < 0) {
            v12 = 32767
        } else {
            v18 := GenS2{F0: 42}
            fmt.Println(v6, v6, v5)
            _ = v18
        }
        _ = v17
    }
    for fi19 := 0; fi19 < 3; fi19++ {
        var v20 uint16 = 42
        if v5 {
            v4[(^(fi19))] = "foo bar"
            fmt.Println(v6, fi19, v6)
            v21 := GenS3{F0: 0}
            _ = v21
        } else {
            v22 := GenS3{F0: v12}
            _ = len(v2)
            fmt.Println(v12)
            _ = v22
        }
        fmt.Println(v6, v5)
        _ = fi19
        _ = v20
    }
    v5 = (0 < 243)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
