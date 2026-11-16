package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8; F1 int64; F2 uint }
func genFn0() uint32 {
    var v7 uint64 = 42
    var v8 uint64 = 114
    switch -9223372036854775808 {
    case 9223372036854775807:
        var v9 int32 = 189
        _ = v9
    case 127:
        fmt.Println(v7, v8)
    }
    v10 := []int{32767}
    _ = v10
    _ = v7
    _ = v8
    return 42
}
func genFn1() int64 {
    var v11 string = ("\n\t" + "")
    var v12 int64 = 1
    _ = v11
    _ = v12
    return -9223372036854775808
}

func main() {
    var v1 int8 = -85
    var v2 string = "a"
    var v3 string = "hello"
    v4 := []uint32{466}
    v5 := []string{}
    v6 := make(map[int]uint32)
    if len(v4) > 0 {
        v4[0] = 2
    }
    v1 += 100
    fmt.Println(v2)
    switch 32767 {
    case 42:
        if true {
            v13 := GenS1{F0: 24, F1: 118}
            _ = v13
        }
    }
    delete(v6, 32767)
    v14 := GenS2{F0: false}
    if len(v5) > 0 {
        v5[0] = (v2 + (v2 + "\n\t"))
    }
    switch 2147483647 {
    case 834:
        fmt.Println(v3, v1)
    }
    v4 = append(v4, 47)
    v1 += (v1 % ((v1) | 1))
    switch v2 {
    case "\n\t":
        v15 := GenS1{F0: 936, F1: 216}
        _ = v15
    case "":
        fmt.Println(v1, v3, v1)
    case "a":
        fmt.Println(v2, v3)
    }
    var v16 int32 = -949
    v17 := GenS1{F0: 2, F1: 0}
    if false {
        if true {
            v6[541] = genFn0()
        }
    }
    if ((v16 ^ (v16 / ((v16) | 1))) < (v16 - v16)) {
        var v18 uint16 = 201
        _ = v18
    } else {
        v3 = v3
        var v19 float32 = -2.25
        _ = v19
    }
    v20 := GenS3{F0: (v1 % (((v1 ^ v1)) | 1)), F1: int64(v16), F2: 4294967295}
    v16 += (-(v16))
    v16 += (v16 * -372)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
