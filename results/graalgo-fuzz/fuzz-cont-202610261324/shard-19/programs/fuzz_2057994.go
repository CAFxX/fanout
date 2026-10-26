package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int; F1 string; F2 bool }
func genFn0(p0 uint, p1 uint8) string {
    v4 := []int{256, 32767}
    fmt.Println(p0, p1)
    fmt.Println(p1)
    p0 = p0
    switch "foo bar" {
    case "a":
        fmt.Println(p0, p0)
    }
    _ = p0
    _ = p1
    _ = v4
    return (("\n\t" + ("\n\t" + "\n\t")) + "")
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]string)
    v3 := make(map[string]uint32)
    v5 := GenS1{F0: -374, F1: "\n\t"}
    if (1 >= 4294967295) {
        var v6 int32 = 2147483647
        var v7 string = "a"
        _ = v6
        _ = v7
    } else {
        if false {
            var v8 uint8 = 30
            v9 := []int{}
            _ = v8
            _ = v9
        } else {
            v10 := GenS2{F0: 4294967295}
            _ = v10
        }
        var v11 uint = 40
        _ = v11
    }
    v12 := GenS2{F0: 42}
    switch "hello" {
    case "\n\t":
        var v13 uint32 = 382
        _ = v13
    }
    var v14 bool = false
    switch "hello" {
    case "\n\t":
        v14 = (9223372036854775807 > -9223372036854775808)
    default:
        var v15 int32 = 42
        _ = v15
    }
    for fi16 := 0; fi16 < 1; fi16++ {
        switch "" {
        case "hello":
            for fi17 := 0; fi17 < 14; fi17++ {
                fmt.Println(v14, v14, fi16)
                v18 := GenS3{F0: 252}
                v14 = true
                _ = fi17
                _ = v18
            }
        }
        v19 := GenS4{F0: 1}
        _ = fi16
        _ = v19
    }
    delete(v1, 128)
    v20 := GenS3{F0: 255}
    v14 = false
    switch 128 {
    case -9223372036854775808:
        if (v14 && false) {
            fmt.Println(v14)
            var v21 int32 = -701
            var v22 int32 = -611
            _ = v21
            _ = v22
        }
    case 146:
        var v23 float32 = 1e10
        _ = v23
    case -1:
        v24 := GenS4{F0: 0}
        _ = v24
    }
    v3["hello"] = 65535
    switch 940 {
    case -9223372036854775808:
        _ = len(v3)
    case 867:
        var v25 bool = (!v14)
        _ = v25
    default:
        v26 := GenS5{F0: 1, F1: "foo bar", F2: (!v14)}
        _ = v26
    }
    v14 = false
    v14 = ((-2.25 * 0.0) != (0.1 / float64(1.5)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v14
    _ = v2
    _ = v20
    _ = v3
    _ = v5
}
