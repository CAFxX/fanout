package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int64; F1 uint; F2 int16 }
type GenS5 struct { F0 int8; F1 int; F2 bool }
func genFn0(p0 uint16, p1 int32) int64 {
    var v4 int8 = 1
    var v5 string = "a"
    for fi6 := 0; fi6 < 19; fi6++ {
        var v7 string = v5
        p0 = (p0 * (p0 & p0))
        for fi8 := 0; fi8 < 17; fi8++ {
            v9 := []int{-9223372036854775808, fi6}
            _ = fi8
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    for fi10 := 0; fi10 < 12; fi10++ {
        v5 = "foo bar"
        v11 := []int{fi10, (fi10 + 18), 128}
        if len(v11) > 0 {
            v11[0] = (-(fi10))
        }
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 42
}
func genFn1(p0 bool, p1 int32) bool {
    v12 := GenS1{F0: 0, F1: genFn0(42, -457)}
    p0 = true
    p0 = (-1 < 0)
    fmt.Println(p1, p1)
    v13 := GenS1{F0: 9223372036854775807, F1: genFn0(65535, 42)}
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return (false && p0)
}

func main() {
    v1 := []string{"", "\n\t", "hello"}
    v2 := make(map[string]string)
    var v3 uint16 = 42
    for fi14 := 0; fi14 < 1; fi14++ {
        if false {
            var v15 string = "hello"
            _ = v15
        } else {
            var v16 uint32 = 1
            _ = v16
        }
        _ = fi14
    }
    v17 := GenS2{F0: false}
    v18 := GenS3{F0: 0}
    delete(v2, "foo bar")
    for fi19 := 0; fi19 < 11; fi19++ {
        v20 := GenS4{F0: 9223372036854775807, F1: 992, F2: 341}
        fmt.Println(fi19, fi19)
        var v21 uint8 = 0
        _ = fi19
        _ = v20
        _ = v21
    }
    var v22 float32 = float32(0.0)
    var v23 float64 = 1e10
    fmt.Println(v3, v23)
    if true {
        var v24 int8 = 0
        var v25 float64 = 1e10
        _ = v24
        _ = v25
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v26 := GenS1{F0: 456, F1: 607}
    switch ("foo bar" + ("hello" + "hello")) {
    case "hello":
        v27 := GenS5{F0: 42, F1: 9223372036854775807, F2: (!true)}
        _ = v27
    case "foo bar":
        v23 += 1e10
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v26
    _ = v3
}
