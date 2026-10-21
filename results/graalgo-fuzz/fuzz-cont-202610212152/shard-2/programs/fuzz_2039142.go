package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 int; F1 uint32; F2 uint64 }
func genFn0(p0 uint16, p1 string) uint8 {
    p0 += (^(p0))
    v6 := GenS1{F0: 497}
    var v7 uint = 0
    v8 := GenS1{F0: 0}
    v9 := []int{-472, 2147483647, 4294967295, 4294967295}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 0
}
func genFn1() uint32 {
    switch 42 {
    case 1:
        v10 := []int{-9223372036854775808, -1, 4294967295}
        _ = v10
    default:
        for fi11 := 0; fi11 < 12; fi11++ {
            var v12 int32 = 992
            v13 := GenS1{F0: 933}
            _ = fi11
            _ = v12
            _ = v13
        }
    }
    var v14 int8 = 100
    _ = v14
    return 4294967295
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 bool = (719 != 2147483647)
    v3 := make(map[string]int)
    var v4 uint = 885
    var v5 string = (("a" + ("\n\t" + "hello")) + "hello")
    var v15 uint8 = 216
    fmt.Println(v5)
    fmt.Println(v5)
    v4 = (^(v4))
    var v16 int32 = -1
    v17 := GenS2{F0: v5, F1: (65535 >= 2)}
    for fi18 := 0; fi18 < 4; fi18++ {
        fmt.Println(v5, v2, v5)
        v2 = (!true)
        _ = fi18
    }
    _ = len(v3)
    switch "hello" {
    case "foo bar":
        v19 := GenS3{F0: 2147483647, F1: 4294967295, F2: 18446744073709551615}
        _ = v19
    case "a":
        var v20 uint16 = 131
        _ = v20
    default:
        var v21 uint64 = 478
        _ = v21
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
