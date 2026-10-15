package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 int; F1 uint8; F2 int }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int64; F1 int32; F2 bool }
func genFn0(p0 uint32) bool {
    var v7 float32 = ((float32(0.0) * float32(-2.25)) + 0.0)
    var v8 string = "\n\t"
    v8 = "foo bar"
    for fi9 := 0; fi9 < 1; fi9++ {
        for fi10 := 0; fi10 < 9; fi10++ {
            p0 += p0
            var v11 int8 = -12
            var v12 float64 = 3.14159
            _ = fi10
            _ = v11
            _ = v12
        }
        fmt.Println(v8, p0, v8)
        fmt.Println(v7, p0, v8)
        _ = fi9
    }
    switch v8 {
    case "foo bar":
        if (-2.25 != 0.0) {
            var v13 uint64 = 42
            v14 := []int{-9223372036854775808, 2147483647}
            _ = v13
            _ = v14
        }
    case "hello":
        var v15 float32 = 1e10
        _ = v15
    case "a":
        fmt.Println(p0)
    default:
        for fi16 := 0; fi16 < 4; fi16++ {
            if true {
                v17 := GenS1{F0: 32767, F1: (v8 + v8)}
                fmt.Println(fi16, p0)
                v18 := GenS2{F0: fi16, F1: 1, F2: fi16}
                _ = v17
                _ = v18
            }
            _ = fi16
        }
    }
    _ = p0
    _ = v7
    _ = v8
    return (!(!false))
}

func main() {
    v1 := []uint32{356, 2147483647, 591, 917}
    v2 := []uint32{42, 2}
    var v3 uint64 = 42
    var v4 uint32 = 1
    v5 := []int{4294967295, 32768}
    var v6 uint64 = v3
    for fi19 := 0; fi19 < 6; fi19++ {
        var v20 string = "foo bar"
        v4 -= v4
        v20 = (v20 + v20)
        _ = fi19
        _ = v20
    }
    v6 -= (v6 + (v3 & (v6 - (v6 * 4294967295))))
    fmt.Println(v3)
    var v21 int = 21
    v22 := GenS2{F0: 42, F1: 186, F2: int(v4)}
    fmt.Println(v6, v4)
    if len(v1) > 0 {
        v1[0] = (^(v4))
    }
    v3 += v6
    switch v21 {
    case 938:
        v23 := GenS1{F0: 117, F1: ("a" + ("" + ""))}
        _ = v23
    case 712:
        var v24 uint = 402
        _ = v24
    case 42:
        v25 := GenS3{F0: 411}
        _ = v25
    default:
        if false {
            var v26 int64 = 42
            _ = v26
        }
    }
    if len(v5) > 0 {
        v5[0] = v21
    }
    if (v3 >= (v6 / ((731) | 1))) {
        v21 = v21
    } else {
        var v27 bool = true
        v6 += (v6 ^ 763)
        if (v3 < v6) {
            var v28 uint32 = (^(v4))
            _ = v28
        }
        _ = v27
    }
    if false {
        var v29 uint64 = 4294967295
        for fi30 := 0; fi30 < 16; fi30++ {
            var v31 float32 = 0.1
            for fi32 := 0; fi32 < 19; fi32++ {
                fmt.Println(fi32, fi30, v3)
                fmt.Println(v4, v4, v31)
                fmt.Println(v3, v29)
                _ = fi32
            }
            _ = fi30
            _ = v31
        }
        v33 := GenS4{F0: true}
        _ = v29
        _ = v33
    } else {
        for fi34 := 0; fi34 < 6; fi34++ {
            var v35 float64 = 3.14159
            for fi36 := 0; fi36 < 10; fi36++ {
                fmt.Println(v6)
                _ = fi36
            }
            _ = fi34
            _ = v35
        }
        v37 := GenS4{F0: true}
        v38 := GenS5{F0: 981, F1: 2147483647, F2: (false || (true && true))}
        _ = v37
        _ = v38
    }
    v4 = v4
    v4 = 42
    if true {
        v39 := GenS4{F0: false}
        _ = v39
    } else {
        _ = len(v2)
        switch "a" {
        case "":
            var v40 int8 = 110
            _ = v40
        }
        var v41 bool = false
        _ = v41
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
