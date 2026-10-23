package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 int; F1 uint64; F2 uint16 }
type GenS3 struct { F0 uint64; F1 bool }
func genFn0(p0 int, p1 bool, p2 bool) bool {
    p1 = (0 <= 516)
    if ((p0 + -2147483648) != (p0 % (((p0 / ((-437) | 1))) | 1))) {
        var v6 bool = (!p2)
        _ = v6
    } else {
        var v7 string = "foo bar"
        p0 -= (p0 >> ((255) & 63))
        fmt.Println(v7, p0, p2)
        _ = v7
    }
    v8 := GenS1{F0: 581, F1: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return (2 == 1)
}
func genFn1(p0 bool) int {
    switch ("hello" + ((("\n\t" + "hello") + (("hello" + "foo bar") + "")) + "foo bar")) {
    case "\n\t":
        var v9 int64 = -9223372036854775808
        _ = v9
    case "hello":
        v10 := []int{}
        _ = v10
    default:
        fmt.Println(p0)
    }
    p0 = true
    if true {
        var v11 uint8 = 0
        p0 = p0
        _ = v11
    } else {
        v12 := []int{}
        if (42 <= -113) {
            p0 = p0
        }
        _ = v12
    }
    for fi13 := 0; fi13 < 8; fi13++ {
        v14 := GenS2{F0: (fi13 - (fi13 | (fi13 % (((-(fi13))) | 1)))), F1: 42, F2: 0}
        _ = fi13
        _ = v14
    }
    _ = p0
    return -306
}

func main() {
    v1 := []int{127, 127, 128}
    v2 := make(map[uint32]int)
    v3 := make(map[int]bool)
    var v4 float32 = 0.0
    v5 := []uint32{65535}
    fmt.Println(v4)
    if (2 != 42) {
        v4 -= 0.1
    }
    delete(v3, 32767)
    v15 := GenS3{F0: 934, F1: (-1 > 42)}
    if false {
        v4 = 1e10
    }
    v4 += v4
    var v16 string = (("a" + ("a" + "hello")) + (("hello" + ("\n\t" + "\n\t")) + "foo bar"))
    v16 = v16
    fmt.Println(v4, v16)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v17 uint64 = 0
    if true {
        fmt.Println(v16)
        switch genFn1(false) {
        case 752:
            var v18 int16 = -857
            _ = v18
        case 32767:
            v4 = (v4 * 1e10)
        }
        v17 += v17
    } else {
        var v19 uint = 4294967295
        switch -515 {
        case 4294967295:
            for fi20 := 0; fi20 < 20; fi20++ {
                var v21 uint8 = 255
                v22 := GenS2{F0: fi20, F1: v17, F2: 962}
                _ = len(v5)
                _ = fi20
                _ = v21
                _ = v22
            }
        case 0:
            v16 = v16
        }
        var v23 uint = v19
        _ = v19
        _ = v23
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
