package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 bool; F1 int8; F2 int }
func genFn0(p0 uint32, p1 int8, p2 uint) uint16 {
    v7 := []int{}
    var v8 int32 = 309
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 2
}
func genFn1(p0 uint8, p1 int64, p2 int8) uint8 {
    for fi9 := 0; fi9 < 13; fi9++ {
        var v10 float32 = 1.5
        _ = fi9
        _ = v10
    }
    if ((!(false && true)) && (true || false)) {
        fmt.Println(p0, p0, p0)
        v11 := GenS1{F0: false}
        switch "" {
        case "foo bar":
            for fi12 := 0; fi12 < 2; fi12++ {
                p2 -= -128
                for fi13 := 0; fi13 < 15; fi13++ {
                    if false {
                        v14 := GenS2{F0: (p2 << ((4294967295) & 63))}
                        v15 := GenS3{F0: p0}
                        if true {
                            p0 -= p0
                            fmt.Println(fi12, p1)
                        } else {
                            p1 += int64(fi13)
                        }
                        _ = v14
                        _ = v15
                    }
                    p1 = 9223372036854775807
                    _ = fi13
                }
                _ = fi12
            }
        }
        _ = v11
    } else {
        var v16 uint32 = 2
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    return 128
}

func main() {
    var v1 uint16 = 42
    var v2 uint32 = 1
    v3 := make(map[uint32]bool)
    var v4 float32 = 0.1
    v5 := make(map[int]string)
    var v6 bool = true
    if (42 >= 32767) {
        if v6 {
            fmt.Println(v1, v4, v1)
            v17 := GenS4{F0: (2 <= 255), F1: 123, F2: -9223372036854775808}
            _ = v17
        }
    }
    if ((v1 > v1) || v6) {
        var v18 int8 = 42
        _ = v18
    } else {
        v1 += 255
        if (2 == 308) {
            var v19 int8 = 0
            var v20 float32 = (1.5 + v4)
            var v21 uint64 = 42
            _ = v19
            _ = v20
            _ = v21
        }
    }
    switch 1 {
    case 32768:
        var v22 int32 = 42
        _ = v22
    case -957:
        switch "hello" {
        case "\n\t":
            fmt.Println(v6)
        }
    }
    fmt.Println(v6, v1)
    v23 := []int{}
    var v24 int32 = -976
    switch ("a" + "\n\t") {
    case "hello":
        _ = len(v3)
    case "":
        switch "foo bar" {
        case "":
            fmt.Println(v24, v24, v1)
        default:
            var v25 int64 = 42
            _ = v25
        }
    case "a":
        if len(v23) > 0 {
            v23[0] = 32768
        }
    default:
        var v26 int8 = 100
        _ = v26
    }
    v23 = append(v23, 127)
    v27 := GenS1{F0: (!true)}
    v28 := GenS1{F0: v6}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
