package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string; F1 int32 }
type GenS3 struct { F0 uint16; F1 int64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint64; F1 uint16; F2 uint }
func genFn0() string {
    if false {
        if true {
            var v6 uint = 65535
            v6 = (^(v6))
            _ = v6
        } else {
            switch 298 {
            case 42:
                var v7 int64 = 42
                _ = v7
            case -9223372036854775808:
                var v8 uint = 0
                _ = v8
            case -536:
                var v9 uint64 = 2
                _ = v9
            }
            var v10 int = 9223372036854775807
            _ = v10
        }
        var v11 int8 = 0
        _ = v11
    } else {
        v12 := []int{-1}
        v13 := GenS1{F0: 42}
        _ = v12
        _ = v13
    }
    var v14 int = 32768
    if true {
        v15 := GenS2{F0: (("hello" + ("foo bar" + "\n\t")) + ("" + "\n\t")), F1: 0}
        fmt.Println(v14)
        _ = v15
    }
    var v16 string = "\n\t"
    v17 := GenS3{F0: 508, F1: 42}
    _ = v14
    _ = v16
    _ = v17
    return ("hello" + v16)
}
func genFn1(p0 bool, p1 int16) uint16 {
    fmt.Println(p0)
    fmt.Println(p1, p1)
    var v18 int8 = 127
    var v19 int32 = 642
    _ = p0
    _ = p1
    _ = v18
    _ = v19
    return 255
}

func main() {
    var v1 int8 = -12
    v2 := make(map[int]uint32)
    var v3 int8 = -128
    var v4 uint16 = 255
    var v5 int16 = 0
    v20 := GenS4{F0: "\n\t"}
    for fi21 := 0; fi21 < 13; fi21++ {
        for fi22 := 0; fi22 < 2; fi22++ {
            v23 := []int{fi21}
            var v24 uint64 = 18446744073709551615
            v2[fi21] = 42
            _ = fi22
            _ = v23
            _ = v24
        }
        var v25 int = fi21
        v4 -= v4
        _ = fi21
        _ = v25
    }
    fmt.Println(v1)
    switch 128 {
    case 466:
        fmt.Println(v1, v5)
    case -9223372036854775808:
        v26 := GenS5{F0: 18446744073709551615, F1: (v4 | v4), F2: 42}
        _ = v26
    case 32767:
        v4 = (v4 % (((v4 & v4)) | 1))
    default:
        v1 -= 1
    }
    v27 := []int{2147483647, 0, 32767, 501}
    v4 = v4
    v1 -= (v3 | v1)
    v28 := v2[32767]
    for fi29 := 0; fi29 < 10; fi29++ {
        var v30 uint = 670
        _ = fi29
        _ = v30
    }
    var v31 int64 = -615
    if len(v27) >= 2 {
        v27 = v27[1:len(v27)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
