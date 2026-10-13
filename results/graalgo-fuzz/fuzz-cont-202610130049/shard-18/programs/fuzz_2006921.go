package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint }
func genFn0() int {
    var v5 uint16 = 65535
    v5 += v5
    for fi6 := 0; fi6 < 7; fi6++ {
        v7 := GenS1{F0: 2147483647}
        var v8 uint16 = (v5 / ((v5) | 1))
        v9 := GenS1{F0: 2147483648}
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    if false {
        v10 := GenS1{F0: 0}
        for fi11 := 0; fi11 < 10; fi11++ {
            v5 -= v5
            _ = fi11
        }
        _ = v10
    }
    v12 := []int{127, 128, -9223372036854775808}
    _ = v12
    _ = v5
    return 42
}
func genFn1(p0 uint) uint {
    v13 := []int{}
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    p0 -= 741
    _ = p0
    _ = v13
    return (p0 * (p0 << ((p0) & 63)))
}

func main() {
    var v1 uint8 = 2
    var v2 bool = true
    v3 := make(map[int]bool)
    v4 := make(map[int]bool)
    v14 := GenS2{F0: 4294967295}
    switch "foo bar" {
    case "":
        for fi15 := 0; fi15 < 3; fi15++ {
            fmt.Println(fi15, v1)
            _ = fi15
        }
    default:
        var v16 int64 = 42
        _ = v16
    }
    for fi17 := 0; fi17 < 14; fi17++ {
        var v18 int = 128
        var v19 uint16 = 2
        v2 = (2 >= v1)
        _ = fi17
        _ = v18
        _ = v19
    }
    v2 = v2
    var v20 int8 = 1
    v21 := GenS1{F0: 1}
    v22 := GenS2{F0: uint(v1)}
    v23 := GenS1{F0: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
