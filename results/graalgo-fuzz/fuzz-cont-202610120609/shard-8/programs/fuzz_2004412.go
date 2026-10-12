package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 int64; F1 string }
func genFn0(p0 int8, p1 int8, p2 string) uint32 {
    if true {
        v6 := []int{32767}
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        p1 = (p0 & (p0 - (p1 * (p1 + 31))))
        _ = v6
    }
    v7 := []int{32767, 0}
    p0 -= p0
    p1 = (^(p1))
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 2
}
func genFn1(p0 bool, p1 int8) uint {
    v8 := GenS1{F0: 2}
    for fi9 := 0; fi9 < 19; fi9++ {
        var v10 uint8 = 124
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 2; fi11++ {
        var v12 int8 = 0
        _ = fi11
        _ = v12
    }
    fmt.Println(p1, p0)
    v13 := GenS1{F0: 874}
    _ = p0
    _ = p1
    _ = v13
    _ = v8
    return 300
}

func main() {
    v1 := make(map[string]bool)
    var v2 int64 = -391
    var v3 string = "a"
    v4 := make(map[uint32]string)
    var v5 uint8 = 255
    for fi14 := 0; fi14 < 15; fi14++ {
        switch "hello" {
        case "\n\t":
            var v15 float32 = -2.25
            _ = v15
        case "foo bar":
            fmt.Println(v3, fi14, v3)
        default:
            var v16 uint64 = 1
            _ = v16
        }
        _ = fi14
    }
    for fi17 := 0; fi17 < 1; fi17++ {
        var v18 int = (fi17 >> ((458) & 63))
        _ = fi17
        _ = v18
    }
    v3 = ((v3 + "") + v3)
    var v19 uint8 = v5
    v2 = (v2 & v2)
    var v20 int8 = -7
    var v21 float64 = ((0.0 - 3.14159) * (1.5 / float64(1.5)))
    var v22 int16 = 289
    v19 -= 0
    v3 = "hello"
    var v23 float64 = (v21 + v21)
    switch "a" {
    case "foo bar":
        var v24 bool = (!(18446744073709551615 >= 527))
        _ = v24
    case "a":
        v25 := GenS1{F0: 18446744073709551615}
        _ = v25
    case "\n\t":
        v26 := GenS2{F0: 0, F1: v3}
        _ = v26
    }
    var v27 uint8 = (v5 - v19)
    v28 := GenS3{F0: v2, F1: ("" + "hello")}
    fmt.Println(v21, v22, v27)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
