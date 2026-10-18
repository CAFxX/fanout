package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16 }
func genFn0() bool {
    v5 := []int{0, 558, 899}
    switch "a" {
    case "a":
        v6 := GenS1{F0: 0}
        _ = v6
    case "hello":
        var v7 int64 = 0
        _ = v7
    case "\n\t":
        var v8 float32 = 0.0
        _ = v8
    default:
        for fi9 := 0; fi9 < 12; fi9++ {
            var v10 float32 = -2.25
            v11 := GenS1{F0: 128}
            v12 := GenS1{F0: 255}
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
        }
    }
    _ = v5
    return false
}
func genFn1(p0 uint64) int16 {
    v13 := []int{42, 461, 4294967295}
    v14 := GenS1{F0: 29}
    switch "\n\t" {
    case "\n\t":
        fmt.Println(p0)
    case "a":
        p0 = (p0 | (^(p0)))
    default:
        var v15 int64 = -9223372036854775808
        _ = v15
    }
    _ = p0
    _ = v13
    _ = v14
    return -36
}

func main() {
    v1 := []string{"hello", ""}
    v2 := make(map[string]int)
    var v3 int64 = 0
    var v4 int64 = (v3 & v3)
    fmt.Println(v4)
    fmt.Println(v4, v3)
    switch "" {
    case "hello":
        v16 := GenS1{F0: 181}
        _ = v16
    case "a":
        v3 = (v4 * v3)
    case "":
        v17 := v2["\n\t"]
        _ = v17
    }
    v4 = (v4 >> ((182) & 63))
    var v18 int16 = 318
    v3 = (v4 << ((4294967295) & 63))
    if (1e10 > 0.0) {
        for fi19 := 0; fi19 < 17; fi19++ {
            var v20 float32 = 1e10
            _ = fi19
            _ = v20
        }
        var v21 uint64 = 4294967295
        v4 = v3
        _ = v21
    } else {
        v4 = (v3 << ((65535) & 63))
        v18 = v18
    }
    for fi22 := 0; fi22 < 14; fi22++ {
        if true {
            var v23 string = "foo bar"
            _ = v23
        } else {
            v4 += v4
        }
        var v24 int8 = -128
        _ = fi22
        _ = v24
    }
    v25 := GenS1{F0: 128}
    v18 += -1
    if true {
        v26 := GenS2{F0: 842}
        _ = v26
    }
    v18 = v18
    switch "hello" {
    case "foo bar":
        fmt.Println(v3, v18)
    }
    fmt.Println(v18, v4, v3)
    fmt.Println(v3, v3, v3)
    switch 1 {
    case 9223372036854775807:
        var v27 string = ("" + (("hello" + "") + ""))
        _ = v27
    case -210:
        var v28 int16 = 1
        _ = v28
    }
    fmt.Println(v4, v3)
    if len(v1) > 0 {
        v1[0] = ("" + "\n\t")
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v25
    _ = v3
    _ = v4
}
