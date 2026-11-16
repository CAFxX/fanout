package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16; F1 int64 }
func genFn0(p0 bool, p1 uint8, p2 uint32) uint8 {
    fmt.Println(p0, p0, p2)
    v5 := GenS1{F0: 43, F1: 1}
    p1 += 0
    v6 := GenS1{F0: 88, F1: 499}
    if (true && p0) {
        var v7 uint = 65535
        for fi8 := 0; fi8 < 19; fi8++ {
            p0 = p0
            _ = fi8
        }
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return p1
}
func genFn1(p0 uint) int {
    var v9 string = ("foo bar" + "\n\t")
    v10 := []int{2147483647, 4294967295, 1}
    v11 := GenS2{F0: "\n\t"}
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    _ = p0
    _ = v10
    _ = v11
    _ = v9
    return 1
}

func main() {
    var v1 uint8 = 1
    var v2 string = ("hello" + ("" + ""))
    v3 := make(map[int]bool)
    var v4 uint64 = 0
    var v12 int32 = 273
    v13 := []int{}
    v12 = (v12 + (v12 / ((42) | 1)))
    switch "" {
    case "hello":
        v14 := GenS3{F0: 1, F1: 0}
        _ = v14
    }
    var v15 int8 = 19
    v16 := GenS2{F0: v2}
    for fi17 := 0; fi17 < 5; fi17++ {
        v4 -= v4
        v3[fi17] = false
        _ = fi17
    }
    if false {
        _ = len(v13)
        v15 += (v15 & v15)
    } else {
        if (false || false) {
            v4 = v4
            var v18 bool = (v4 == 252)
            _ = v18
        }
        for fi19 := 0; fi19 < 12; fi19++ {
            v13 = append(v13, genFn1(65535))
            for fi20 := 0; fi20 < 9; fi20++ {
                fmt.Println(v15, v4, v4)
                var v21 uint = 594
                v21 = v21
                _ = fi20
                _ = v21
            }
            _ = fi19
        }
        if (true && (1 == 0)) {
            _ = len(v13)
            v12 -= (v12 - v12)
        }
    }
    v22 := GenS1{F0: 127, F1: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}
