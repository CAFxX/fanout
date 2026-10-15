package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 bool }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 int; F1 int }
func genFn0(p0 string, p1 uint8) uint64 {
    switch (p0 + ("a" + ((p0 + ((p0 + p0) + (p0 + p0))) + p0))) {
    case "":
        var v5 int8 = 1
        _ = v5
    case "foo bar":
        var v6 float32 = 0.0
        _ = v6
    case "hello":
        if (!false) {
            v7 := []int{1}
            switch 42 {
            case 2147483647:
                fmt.Println(p0, p0)
            case -9223372036854775808:
                if true {
                    v8 := GenS1{F0: -319, F1: 0, F2: true}
                    if false {
                        p0 = p0
                    }
                    _ = v8
                } else {
                    p1 = (p1 * p1)
                }
            }
            var v9 int8 = 100
            _ = v7
            _ = v9
        }
    default:
        v10 := []int{}
        _ = v10
    }
    p1 = (p1 / (((p1 * 1)) | 1))
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    return 4294967295
}
func genFn1(p0 int64, p1 int, p2 uint32) int64 {
    var v11 string = "hello"
    p2 += (p2 >> ((1) & 63))
    fmt.Println(p0, v11)
    if true {
        var v12 int = p1
        v13 := GenS1{F0: (p0 >> ((808) & 63)), F1: (p2 | p2), F2: false}
        fmt.Println(p1, v12)
        _ = v12
        _ = v13
    } else {
        v14 := GenS1{F0: p0, F1: p2, F2: (!(-1 > 0))}
        v15 := GenS2{F0: uint8(p2), F1: v11}
        _ = v14
        _ = v15
    }
    v11 = "hello"
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 0
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int = -2147483648
    var v3 uint32 = 130
    v4 := make(map[string]string)
    if (true || false) {
        if (42 <= 0) {
            var v16 int = 42
            _ = v16
        }
    }
    if true {
        fmt.Println(v3)
        if false {
            for fi17 := 0; fi17 < 18; fi17++ {
                fmt.Println(v3, v3)
                var v18 int32 = -933
                for fi19 := 0; fi19 < 16; fi19++ {
                    v18 = 2147483647
                    _ = fi19
                }
                _ = fi17
                _ = v18
            }
        } else {
            v20 := GenS1{F0: genFn1(-15, v2, 0), F1: (v3 - 784), F2: true}
            v21 := []int{-9223372036854775808, -2147483648, v2}
            if true {
                v22 := GenS2{F0: 181, F1: "hello"}
                _ = v22
            }
            _ = v20
            _ = v21
        }
        delete(v4, "a")
    } else {
        v2 = (v2 % ((217) | 1))
    }
    v2 += v2
    var v23 uint32 = (v3 / ((v3) | 1))
    var v24 float32 = 1.5
    var v25 float64 = (1e10 / 1.5)
    v26 := GenS3{F0: v2, F1: v2}
    var v27 uint = 384
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
