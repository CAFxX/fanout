package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
func genFn0(p0 uint32, p1 uint16, p2 bool) uint8 {
    switch "\n\t" {
    case "hello":
        v6 := []int{9223372036854775807, -9223372036854775808, 2147483647}
        _ = v6
    case "\n\t":
        fmt.Println(p0)
    default:
        for fi7 := 0; fi7 < 19; fi7++ {
            v8 := []int{}
            switch "hello" {
            case "foo bar":
                fmt.Println(p1)
            default:
                p1 = 925
            }
            _ = fi7
            _ = v8
        }
    }
    for fi9 := 0; fi9 < 2; fi9++ {
        v10 := GenS1{F0: (("hello" + "\n\t") + ("a" + (("a" + (("\n\t" + "hello") + "")) + "a"))), F1: 42}
        for fi11 := 0; fi11 < 17; fi11++ {
            fmt.Println(p2, p0, fi11)
            var v12 int8 = int8(fi9)
            _ = fi11
            _ = v12
        }
        fmt.Println(p0)
        _ = fi9
        _ = v10
    }
    var v13 int32 = 679
    p0 = p0
    v14 := GenS1{F0: "\n\t", F1: uint64(p1)}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    return 127
}

func main() {
    v1 := []uint32{1, 65535}
    v2 := []string{"a", "\n\t", ("a" + "hello")}
    v3 := make(map[string]string)
    v4 := make(map[int]int)
    v5 := make(map[uint32]int)
    var v15 uint8 = 42
    var v16 bool = (v15 != genFn0(42, 32767, false))
    switch ("a" + "") {
    case "foo bar":
        fmt.Println(v15)
    case "":
        _ = len(v4)
    }
    v2 = append(v2, ("" + ""))
    v15 += v15
    fmt.Println(v16)
    var v17 float32 = 3.14159
    if (1000 != 42) {
        v17 = 0.1
    } else {
        var v18 string = "\n\t"
        fmt.Println(v17)
        v18 = v18
        _ = v18
    }
    var v19 float32 = v17
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
