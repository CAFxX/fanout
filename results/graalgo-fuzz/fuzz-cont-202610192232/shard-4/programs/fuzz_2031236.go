package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 uint; F1 uint64 }
func genFn0(p0 int32) int32 {
    switch 128 {
    case 128:
        var v4 int64 = -1
        _ = v4
    default:
        fmt.Println(p0)
    }
    v5 := []int{32768}
    v6 := GenS1{F0: 32768, F1: 0}
    fmt.Println(p0)
    _ = p0
    _ = v5
    _ = v6
    return (p0 & p0)
}
func genFn1() int {
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 int8 = 1
        v8 += (-(v8))
        var v9 uint16 = 32767
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 uint = 153
    _ = v10
    return 4294967295
}

func main() {
    var v1 string = "\n\t"
    v2 := make(map[string]bool)
    var v3 uint16 = 1
    if true {
        switch v1 {
        case "\n\t":
            v1 = v1
        case "a":
            v11 := []int{genFn1(), 127, genFn1(), 0}
            _ = v11
        }
    }
    switch "foo bar" {
    case "hello":
        switch v1 {
        case "hello":
            var v12 float32 = (1e10 - float32(3.14159))
            _ = v12
        }
    default:
        fmt.Println(v1, v3)
    }
    if (!false) {
        v13 := GenS1{F0: 32767, F1: 4294967295}
        _ = v13
    }
    v14 := []int{127, 32767, 127}
    for fi15 := 0; fi15 < 9; fi15++ {
        for fi16 := 0; fi16 < 5; fi16++ {
            if len(v14) > 0 {
                v14[0] = fi16
            }
            v3 += (v3 << ((65535) & 63))
            v17 := GenS2{F0: 462, F1: 18446744073709551615}
            _ = len(v2)
            if true {
                v3 = (v3 & (^(v3)))
            }
            _ = fi16
            _ = v17
        }
        _ = fi15
    }
    v18 := GenS1{F0: 926, F1: 839}
    var v19 bool = false
    v3 = v3
    v19 = false
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
