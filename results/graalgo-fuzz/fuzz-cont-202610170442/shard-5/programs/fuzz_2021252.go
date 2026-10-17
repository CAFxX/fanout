package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 uint8 }
type GenS2 struct { F0 uint64; F1 string }
type GenS3 struct { F0 int32; F1 bool; F2 string }
func genFn0(p0 uint16, p1 int, p2 uint) uint16 {
    switch "a" {
    case "a":
        var v4 int8 = 1
        _ = v4
    default:
        var v5 uint32 = 2147483647
        _ = v5
    }
    p2 -= 1
    v6 := GenS1{F0: 100, F1: 42, F2: 216}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return (p0 & p0)
}
func genFn1(p0 int8, p1 int8, p2 uint16) int32 {
    if (((false || false) && true) || true) {
        if false {
            v7 := []int{2147483647, 0}
            _ = v7
        } else {
            p1 = p1
            v8 := GenS2{F0: 1, F1: ("hello" + "\n\t")}
            var v9 string = "foo bar"
            _ = v8
            _ = v9
        }
        if false {
            var v10 float32 = 1.5
            var v11 int64 = 1
            _ = v10
            _ = v11
        } else {
            switch -1 {
            case -1:
                var v12 int8 = (^(p0))
                _ = v12
            case 1:
                var v13 uint32 = 86
                _ = v13
            case 127:
                var v14 bool = false
                _ = v14
            default:
                var v15 int32 = 1
                _ = v15
            }
            p1 = 0
        }
        v16 := []int{-413, 9223372036854775807, 61, 32767}
        _ = v16
    }
    switch "foo bar" {
    case "hello":
        fmt.Println(p2, p1, p1)
    case "a":
        v17 := GenS3{F0: -653, F1: (!(1 > 370)), F2: ""}
        _ = v17
    case "\n\t":
        p2 = p2
    }
    var v18 float32 = (3.14159 + 0.1)
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    return -620
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int64 = -436
    v3 := make(map[string]string)
    var v19 float32 = -2.25
    var v20 int32 = 2147483647
    v21 := []int{-9223372036854775808, 32767}
    for fi22 := 0; fi22 < 13; fi22++ {
        var v23 int64 = v2
        _ = fi22
        _ = v23
    }
    v20 = genFn1(1, -71, 65535)
    v3["hello"] = "\n\t"
    fmt.Println(v20)
    v19 += float32(3.14159)
    switch 532 {
    case 42:
        var v24 uint = 65535
        _ = v24
    case 128:
        var v25 uint16 = genFn0(210, 603, 2)
        _ = v25
    }
    var v26 uint64 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v26
    _ = v3
}
