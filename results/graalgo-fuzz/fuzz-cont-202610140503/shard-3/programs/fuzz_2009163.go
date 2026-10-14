package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 bool }
func genFn0() int8 {
    var v5 int64 = -9223372036854775808
    var v6 float64 = float64(-2.25)
    for fi7 := 0; fi7 < 1; fi7++ {
        v5 = int64(fi7)
        var v8 bool = (true || false)
        _ = fi7
        _ = v8
    }
    v5 = v5
    if (42 == 42) {
        switch "a" {
        case "a":
            var v9 uint32 = 2147483648
            _ = v9
        case "":
            fmt.Println(v5, v5)
        case "hello":
            var v10 int8 = 54
            _ = v10
        }
        var v11 int64 = v5
        _ = v11
    }
    _ = v5
    _ = v6
    return 100
}
func genFn1(p0 uint32, p1 string) int16 {
    v12 := GenS1{F0: genFn0(), F1: 827}
    switch 32768 {
    case 42:
        v13 := GenS1{F0: 42, F1: 42}
        _ = v13
    case 4294967295:
        v14 := []int{4294967295, 4294967295}
        _ = v14
    case 757:
        p1 = p1
    }
    _ = p0
    _ = p1
    _ = v12
    return 0
}

func main() {
    var v1 int64 = -1
    var v2 int32 = -612
    var v3 uint64 = 42
    var v4 uint = 18446744073709551615
    v4 -= 42
    v2 += int32(v4)
    v1 += (v1 ^ (v1 | (-(v1))))
    var v15 uint64 = v3
    v16 := GenS1{F0: -86, F1: v3}
    v1 += (v1 << (((v4 ^ (^(v4)))) & 63))
    v1 += (-(v1))
    if false {
        v4 = (^(v4))
        var v17 uint16 = 1
        _ = v17
    } else {
        switch 9223372036854775807 {
        case -9223372036854775808:
            if true {
                v18 := []int{}
                _ = len(v18)
                _ = v18
            }
        case -1:
            var v19 int16 = 42
            _ = v19
        default:
            v20 := []int{4294967295, 378, -9223372036854775808, 42}
            _ = v20
        }
        fmt.Println(v2, v3)
    }
    switch ((("" + (("" + (("foo bar" + (("" + "") + "hello")) + "hello")) + "foo bar")) + ("hello" + "a")) + "a") {
    case "\n\t":
        v21 := []int{}
        _ = v21
    }
    v1 = v1
    for fi22 := 0; fi22 < 8; fi22++ {
        v1 += v1
        _ = fi22
    }
    var v23 uint = v4
    if true {
        v23 += 675
    }
    v15 += (v3 << ((v23) & 63))
    v3 = 0
    fmt.Println(v2, v23)
    v24 := []int{-303, 951}
    v25 := GenS2{F0: false}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
