package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 int; F1 int64 }
type GenS3 struct { F0 int8; F1 uint16 }
type GenS4 struct { F0 int16; F1 int }
func genFn0(p0 uint32, p1 uint64, p2 uint) int8 {
    for fi7 := 0; fi7 < 20; fi7++ {
        p1 = p1
        p1 -= p1
        _ = fi7
    }
    for fi8 := 0; fi8 < 11; fi8++ {
        fmt.Println(p0, fi8, p1)
        _ = fi8
    }
    v9 := GenS1{F0: 1, F1: 127}
    v10 := []int{32768}
    var v11 string = ""
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v9
    return 0
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []string{((("\n\t" + "\n\t") + (("a" + ("hello" + "foo bar")) + "hello")) + "")}
    v3 := make(map[string]bool)
    v4 := []uint32{2147483648, 0}
    v5 := []int{0, 32767, -1}
    var v6 int8 = 0
    fmt.Println(v6)
    switch "" {
    case "hello":
        var v12 uint32 = 940
        _ = v12
    case "":
        var v13 string = (("hello" + "hello") + "a")
        _ = v13
    case "a":
        v6 -= v6
    default:
        v6 = v6
    }
    var v14 int8 = v6
    v15 := GenS1{F0: -299, F1: 128}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v6 += (v6 | (v14 * (v6 ^ (v14 - (v14 * 18)))))
    v6 += (v6 - v6)
    v2 = append(v2, "")
    v16 := GenS2{F0: -107, F1: 42}
    for fi17 := 0; fi17 < 15; fi17++ {
        v18 := v3["foo bar"]
        var v19 float64 = (-2.25 + float64(3.14159))
        _ = fi17
        _ = v18
        _ = v19
    }
    switch "hello" {
    case "a":
        v20 := GenS3{F0: (v14 * -119), F1: 255}
        _ = v20
    case "":
        var v21 uint64 = 650
        _ = v21
    default:
        var v22 uint64 = 2
        _ = v22
    }
    v14 = v6
    delete(v1, 508)
    v6 -= -128
    fmt.Println(v6, v6)
    switch "\n\t" {
    case "hello":
        v23 := GenS4{F0: 729, F1: -2147483648}
        _ = v23
    }
    _ = len(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
