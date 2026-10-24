package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 uint16 }
type GenS2 struct { F0 string }
func genFn0(p0 uint16, p1 int32) uint8 {
    v4 := []int{584, -9223372036854775808, int(p1), 352}
    v5 := GenS1{F0: -1, F1: false, F2: 534}
    for fi6 := 0; fi6 < 11; fi6++ {
        if (false && false) {
            var v7 int8 = -128
            var v8 uint = 357
            _ = v7
            _ = v8
        }
        _ = fi6
    }
    var v9 int64 = 526
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v9
    return 201
}
func genFn1(p0 uint32, p1 int, p2 uint) uint32 {
    fmt.Println(p1)
    for fi10 := 0; fi10 < 14; fi10++ {
        v11 := GenS1{F0: -520, F1: (!true), F2: 0}
        p2 += (p2 * p2)
        _ = fi10
        _ = v11
    }
    v12 := []int{4294967295, 2147483647, (p1 | (p1 * (p1 ^ (^(p1))))), p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return uint32(p2)
}

func main() {
    v1 := []string{"foo bar", "hello"}
    v2 := []int{}
    v3 := make(map[uint32]int)
    for fi13 := 0; fi13 < 20; fi13++ {
        var v14 uint8 = 255
        fmt.Println(fi13, v14)
        v14 += genFn0(252, -454)
        _ = fi13
        _ = v14
    }
    delete(v3, 415)
    if (470 >= 286) {
        var v15 uint64 = 42
        _ = v15
    }
    var v16 int = 520
    fmt.Println(v16)
    if true {
        v17 := GenS1{F0: 0, F1: (!(42 <= -1)), F2: 2}
        var v18 string = "hello"
        var v19 uint = 2
        _ = v17
        _ = v18
        _ = v19
    } else {
        if false {
            if true {
                var v20 uint64 = 2
                if (0.0 <= 1.5) {
                    fmt.Println(v16, v20)
                    if len(v1) > 0 {
                        v1[0] = "foo bar"
                    }
                    var v21 int16 = 32767
                    _ = v21
                }
                _ = v20
            } else {
                fmt.Println(v16)
                v16 = v16
            }
            for fi22 := 0; fi22 < 4; fi22++ {
                fmt.Println(v16, fi22)
                _ = fi22
            }
        }
        v16 = v16
        v16 = (v16 + v16)
    }
    for fi23 := 0; fi23 < 6; fi23++ {
        var v24 float32 = (0.1 - 0.0)
        _ = fi23
        _ = v24
    }
    var v25 uint64 = 1
    fmt.Println(v25)
    if false {
        v26 := GenS2{F0: ("foo bar" + "\n\t")}
        fmt.Println(v16)
        _ = len(v1)
        _ = v26
    }
    var v27 float32 = 0.1
    var v28 uint32 = 535
    for fi29 := 0; fi29 < 20; fi29++ {
        if len(v1) > 0 {
            v1[0] = ("" + "foo bar")
        }
        switch ("hello" + "") {
        case "":
            for fi30 := 0; fi30 < 2; fi30++ {
                var v31 float32 = v27
                v25 -= (v25 | (^(v25)))
                _ = fi30
                _ = v31
            }
        case "\n\t":
            v16 = fi29
        case "a":
            fmt.Println(v25)
        }
        _ = fi29
    }
    delete(v3, (v28 ^ genFn1(v28, -2147483648, 696)))
    var v32 int32 = 2147483647
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    _ = v1
    _ = v16
    _ = v2
    _ = v25
    _ = v27
    _ = v28
    _ = v3
    _ = v32
}
