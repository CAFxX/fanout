package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 bool }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint16 }
func genFn0(p0 string, p1 int) int8 {
    switch 127 {
    case 2147483647:
        fmt.Println(p0, p0)
    case -133:
        fmt.Println(p0)
    case -190:
        v4 := GenS1{F0: 255, F1: -596, F2: (p1 == p1)}
        _ = v4
    }
    v5 := []int{p1, p1, 4294967295}
    switch p1 {
    case -1:
        v6 := GenS2{F0: 352}
        _ = v6
    case -9223372036854775808:
        _ = len(v5)
    default:
        var v7 uint16 = 36
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v5
    return 1
}
func genFn1(p0 int32, p1 int, p2 uint32) uint8 {
    var v8 bool = false
    for fi9 := 0; fi9 < 6; fi9++ {
        if (!v8) {
            p1 += p1
        } else {
            for fi10 := 0; fi10 < 7; fi10++ {
                var v11 bool = true
                if false {
                    switch (fi10 % (((fi10 / ((int(p0)) | 1))) | 1)) {
                    case 9223372036854775807:
                        var v12 bool = (v8 || ((fi9 | 726) >= (-(fi10))))
                        _ = v12
                    case 358:
                        var v13 int8 = 100
                        _ = v13
                    case -2147483648:
                        v14 := GenS3{F0: 49}
                        _ = v14
                    }
                    if v8 {
                        if (v11 || (-9223372036854775808 == (p1 << ((0) & 63)))) {
                            var v15 uint8 = 255
                            v16 := GenS2{F0: -2147483648}
                            v17 := []int{fi9, (-(fi9)), fi9}
                            _ = v15
                            _ = v16
                            _ = v17
                        }
                    }
                } else {
                    v18 := GenS1{F0: 81, F1: int32(p1), F2: ((-2.25 - 1e10) > (float64(1e10) - 1.5))}
                    _ = v18
                }
                _ = fi10
                _ = v11
            }
        }
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 95
}

func main() {
    var v1 uint64 = 865
    v2 := make(map[uint32]bool)
    var v3 int8 = 101
    fmt.Println(v1, v1)
    v1 -= v1
    var v19 int8 = genFn0(("foo bar" + "hello"), 127)
    if false {
        fmt.Println(v19, v19, v19)
    } else {
        v3 += (v3 | (v19 | v19))
    }
    var v20 uint64 = (v1 - (^(v1)))
    fmt.Println(v20, v19, v1)
    for fi21 := 0; fi21 < 5; fi21++ {
        var v22 int32 = 42
        v23 := GenS4{F0: 65535}
        _ = fi21
        _ = v22
        _ = v23
    }
    v2[671] = (42 <= 2147483647)
    v3 = (v3 & v3)
    fmt.Println(v1, v19)
    var v24 int32 = 2147483647
    if false {
        delete(v2, 899)
        v20 -= v1
    }
    var v25 int = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v3
}
