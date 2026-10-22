package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 int }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 int8 }
func genFn0(p0 int64, p1 uint8, p2 uint16) uint32 {
    fmt.Println(p1)
    switch int(p1) {
    case 127:
        fmt.Println(p0)
    default:
        switch "foo bar" {
        case "foo bar":
            var v5 float32 = -2.25
            _ = v5
        case "\n\t":
            v6 := GenS1{F0: 2, F1: 2147483647, F2: 32767}
            _ = v6
        default:
            v7 := GenS2{F0: "\n\t", F1: 1}
            _ = v7
        }
    }
    p2 -= 1
    _ = p0
    _ = p1
    _ = p2
    return 42
}
func genFn1(p0 bool, p1 int64, p2 uint16) bool {
    var v8 string = ("" + ("" + "hello"))
    switch "\n\t" {
    case "":
        if p0 {
            p1 = (p1 | (^(p1)))
            var v9 uint64 = 746
            if p0 {
                var v10 uint = 255
                v11 := GenS2{F0: v8, F1: 0}
                var v12 string = (v8 + ("hello" + "foo bar"))
                _ = v10
                _ = v11
                _ = v12
            }
            _ = v9
        }
    case "a":
        for fi13 := 0; fi13 < 16; fi13++ {
            v14 := GenS2{F0: ((v8 + v8) + (v8 + v8)), F1: -1}
            var v15 int32 = -852
            v8 = "foo bar"
            _ = fi13
            _ = v14
            _ = v15
        }
    default:
        var v16 uint = 65535
        _ = v16
    }
    v17 := []int{-563, -361}
    p0 = ((3.14159 + (3.14159 - 0.1)) > (0.0 * 0.1))
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v8
    return (!false)
}

func main() {
    var v1 int8 = -33
    var v2 int8 = v1
    var v3 int = 351
    var v4 string = "hello"
    if false {
        switch 128 {
        case 128:
            v2 = (v2 % (((v2 * v2)) | 1))
        default:
            v18 := GenS3{F0: -1}
            _ = v18
        }
        fmt.Println(v4, v2)
    }
    for fi19 := 0; fi19 < 14; fi19++ {
        v20 := []int{(fi19 | fi19)}
        _ = fi19
        _ = v20
    }
    var v21 int8 = v2
    var v22 int32 = -958
    for fi23 := 0; fi23 < 7; fi23++ {
        var v24 uint32 = 4294967295
        _ = fi23
        _ = v24
    }
    v25 := []int{}
    var v26 uint = 0
    if len(v25) > 0 {
        v25[0] = v3
    }
    var v27 uint64 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
