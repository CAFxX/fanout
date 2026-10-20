package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 int8; F1 uint32; F2 uint8 }
type GenS3 struct { F0 int64; F1 int32; F2 string }
func genFn0(p0 uint16) int32 {
    var v4 float32 = 1.5
    v5 := []int{605, 128, 758}
    _ = p0
    _ = v4
    _ = v5
    return 42
}
func genFn1(p0 int32) int32 {
    p0 = (-(p0))
    switch ("\n\t" + "") {
    case "\n\t":
        var v6 uint64 = 18446744073709551615
        _ = v6
    case "":
        p0 = (p0 % ((p0) | 1))
    case "foo bar":
        var v7 int16 = -1
        _ = v7
    }
    _ = p0
    return (p0 - p0)
}

func main() {
    v1 := []int{4294967295, -798, 128, -1}
    v2 := make(map[uint32]string)
    v3 := make(map[int]bool)
    var v8 float32 = 0.1
    v9 := GenS1{F0: 18446744073709551615, F1: "\n\t"}
    switch "foo bar" {
    case "foo bar":
        v8 = 0.0
    case "a":
        if (!(!false)) {
            var v10 uint32 = 499
            for fi11 := 0; fi11 < 15; fi11++ {
                fmt.Println(fi11)
                _ = fi11
            }
            _ = v10
        }
    }
    v8 -= v8
    switch 42 {
    case -2147483648:
        var v12 uint = 439
        _ = v12
    default:
        var v13 uint8 = 255
        _ = v13
    }
    if false {
        switch (("a" + (("a" + "hello") + "foo bar")) + "") {
        case "hello":
            v8 = 3.14159
        }
        var v14 int = 977
        fmt.Println(v14, v14)
        _ = v14
    } else {
        for fi15 := 0; fi15 < 6; fi15++ {
            v8 -= v8
            _ = fi15
        }
        for fi16 := 0; fi16 < 5; fi16++ {
            var v17 uint16 = 255
            for fi18 := 0; fi18 < 18; fi18++ {
                v8 = 3.14159
                _ = fi18
            }
            var v19 int = (fi16 << ((65535) & 63))
            _ = fi16
            _ = v17
            _ = v19
        }
    }
    for fi20 := 0; fi20 < 8; fi20++ {
        v21 := GenS2{F0: 1, F1: 42, F2: 255}
        _ = fi20
        _ = v21
    }
    fmt.Println(v8)
    if false {
        var v22 uint8 = 127
        v23 := GenS1{F0: 18446744073709551615, F1: "hello"}
        v24 := GenS1{F0: 242, F1: "foo bar"}
        _ = v22
        _ = v23
        _ = v24
    }
    if (!(1e10 == (float64((v8 + 0.0)) * -2.25))) {
        fmt.Println(v8)
    }
    if ((false || true) || true) {
        delete(v3, 1)
        v25 := GenS3{F0: 9223372036854775807, F1: 0, F2: ""}
        _ = v25
    } else {
        v8 -= float32(0.1)
    }
    v26 := GenS1{F0: 2, F1: ("a" + "")}
    var v27 int16 = 1000
    v27 = 42
    if true {
        var v28 uint16 = 1
        _ = v28
    }
    v29 := GenS3{F0: -1, F1: 1, F2: "a"}
    v3[-9223372036854775808] = true
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v29
    _ = v3
    _ = v8
    _ = v9
}
