package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 uint }
func genFn0(p0 string, p1 uint64, p2 int16) int16 {
    if true {
        var v5 string = ""
        var v6 bool = (true || (!((!true) && (1 < 4294967295))))
        switch 127 {
        case 0:
            var v7 uint16 = 0
            _ = v7
        }
        _ = v5
        _ = v6
    } else {
        switch -9223372036854775808 {
        case 9223372036854775807:
            p0 = ((p0 + p0) + ("" + p0))
        case 127:
            for fi8 := 0; fi8 < 20; fi8++ {
                if (0.1 >= float32((-2.25 - float64(-2.25)))) {
                    var v9 int16 = (p2 << ((1) & 63))
                    v10 := []int{(fi8 % ((fi8) | 1)), (fi8 + 42), fi8}
                    _ = v10
                    _ = v9
                }
                for fi11 := 0; fi11 < 14; fi11++ {
                    var v12 uint8 = 126
                    var v13 int16 = (^(p2))
                    p1 = 841
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                _ = fi8
            }
        case 2147483647:
            v14 := GenS1{F0: ""}
            _ = v14
        }
        var v15 uint16 = 532
        p1 -= (p1 % ((42) | 1))
        _ = v15
    }
    p0 = ("a" + p0)
    v16 := []int{128, 9223372036854775807}
    var v17 int64 = 9223372036854775807
    v18 := GenS1{F0: (((p0 + ("a" + "")) + "") + (("\n\t" + p0) + ("" + ((p0 + ((p0 + p0) + "")) + "hello"))))}
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v17
    _ = v18
    return (^(p2))
}

func main() {
    var v1 uint32 = 160
    var v2 uint8 = 127
    v3 := make(map[uint32]uint32)
    v4 := make(map[int]int)
    v2 = 83
    var v19 int64 = 9223372036854775807
    for fi20 := 0; fi20 < 17; fi20++ {
        if false {
            v19 -= -1
        }
        _ = fi20
    }
    v21 := []int{556, 952, -2147483648, -1}
    v19 = (^(v19))
    delete(v3, (^(v1)))
    if len(v21) > 0 {
        v21[0] = 0
    }
    v22 := GenS2{F0: (155 == 0), F1: 45}
    v23 := GenS2{F0: false, F1: 797}
    if (false && (66 != 6)) {
        if len(v21) > 0 {
            v21[0] = 305
        }
    }
    var v24 uint16 = 411
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
