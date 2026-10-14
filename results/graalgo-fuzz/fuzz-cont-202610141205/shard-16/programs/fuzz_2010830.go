package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint; F1 uint16; F2 int32 }
type GenS3 struct { F0 int8; F1 uint64; F2 string }
type GenS4 struct { F0 uint64 }
func genFn0(p0 int) int {
    v7 := []int{}
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    switch "foo bar" {
    case "hello":
        var v8 uint16 = 267
        _ = v8
    default:
        switch p0 {
        case -834:
            for fi9 := 0; fi9 < 19; fi9++ {
                for fi10 := 0; fi10 < 20; fi10++ {
                    p0 = 128
                    _ = fi10
                }
                v11 := GenS1{F0: 42}
                var v12 int32 = 42
                _ = fi9
                _ = v11
                _ = v12
            }
        case 32767:
            if len(v7) > 0 {
                v7[0] = (p0 % ((876) | 1))
            }
        case 128:
            fmt.Println(p0)
        default:
            var v13 int64 = -9223372036854775808
            _ = v13
        }
    }
    p0 -= p0
    switch p0 {
    case 42:
        switch p0 {
        case -252:
            v14 := GenS2{F0: 562, F1: 42, F2: -998}
            _ = v14
        case 0:
            v15 := GenS2{F0: 65535, F1: 65535, F2: 1}
            _ = v15
        }
    default:
        var v16 float32 = (float32(0.0) + (1e10 + 3.14159))
        _ = v16
    }
    _ = p0
    _ = v7
    return (p0 << ((942) & 63))
}
func genFn1(p0 uint64, p1 int32, p2 bool) int64 {
    v17 := []int{32768, -907}
    var v18 int = 9223372036854775807
    var v19 int16 = 1
    for fi20 := 0; fi20 < 10; fi20++ {
        v21 := GenS3{F0: 100, F1: 2, F2: "\n\t"}
        _ = fi20
        _ = v21
    }
    var v22 uint64 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v18
    _ = v19
    _ = v22
    return -1
}

func main() {
    v1 := []int{4294967295, 1}
    var v2 bool = false
    var v3 uint16 = 42
    v4 := []uint32{4294967295, 167, 2147483647, 2147483647}
    v5 := make(map[string]string)
    v6 := make(map[int]string)
    var v23 int16 = -402
    var v24 bool = true
    var v25 int = 2147483647
    _ = len(v4)
    delete(v6, v25)
    delete(v5, ("" + ("foo bar" + (("\n\t" + "a") + ("" + ("foo bar" + ("a" + "foo bar")))))))
    v3 -= v3
    switch v25 {
    case 821:
        fmt.Println(v2, v2, v3)
    default:
        v2 = ((v2 || (false || v2)) || v24)
    }
    v26 := GenS4{F0: 1}
    var v27 string = ("a" + (("foo bar" + "hello") + "hello"))
    for fi28 := 0; fi28 < 10; fi28++ {
        var v29 uint16 = 258
        _ = fi28
        _ = v29
    }
    v25 = v25
    v6[(v25 << ((288) & 63))] = v27
    for fi30 := 0; fi30 < 17; fi30++ {
        fmt.Println(v2, fi30, v24)
        v24 = false
        for fi31 := 0; fi31 < 2; fi31++ {
            fmt.Println(v3, fi31, fi30)
            _ = fi31
        }
        _ = fi30
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
