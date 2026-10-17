package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
func genFn0(p0 uint64, p1 uint32, p2 uint32) uint16 {
    if false {
        switch "\n\t" {
        case "\n\t":
            var v4 string = "hello"
            _ = v4
        default:
            for fi5 := 0; fi5 < 5; fi5++ {
                var v6 int8 = 42
                fmt.Println(p1)
                _ = fi5
                _ = v6
            }
        }
    }
    if (true || true) {
        v7 := []int{0, 42, -308}
        var v8 string = ("a" + "foo bar")
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: 42, F1: uint64(p2)}
    p1 += p2
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 0
}
func genFn1() uint16 {
    v10 := []int{42, 4294967295, 4294967295, 128}
    var v11 uint64 = 42
    _ = v10
    _ = v11
    return 42
}

func main() {
    var v1 int = 4294967295
    var v2 int = 128
    var v3 int32 = int32(v2)
    var v12 int = (v2 - (v2 % (((v1 / ((v1) | 1))) | 1)))
    v3 = v3
    if (1 < 0) {
        v13 := GenS1{F0: (v3 << ((18446744073709551615) & 63)), F1: 18446744073709551615}
        v14 := []int{v1, v2, v1}
        v2 = v1
        _ = v13
        _ = v14
    }
    if true {
        fmt.Println(v2, v3)
    }
    for fi15 := 0; fi15 < 15; fi15++ {
        var v16 int8 = 42
        v3 -= (v3 << ((0) & 63))
        fmt.Println(v3, v3)
        _ = fi15
        _ = v16
    }
    var v17 bool = true
    v18 := []int{v2, (v2 - (^(v1))), v12, v1}
    var v19 int = -219
    var v20 uint = 65535
    var v21 int = v1
    var v22 uint32 = 2147483648
    _ = len(v18)
    v23 := GenS1{F0: v3, F1: 2}
    var v24 int8 = -128
    for fi25 := 0; fi25 < 14; fi25++ {
        for fi26 := 0; fi26 < 9; fi26++ {
            v3 -= v3
            _ = fi26
        }
        _ = fi25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
}
