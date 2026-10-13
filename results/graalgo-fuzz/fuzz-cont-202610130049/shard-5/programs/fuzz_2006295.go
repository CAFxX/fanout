package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint8; F1 bool; F2 uint32 }
type GenS4 struct { F0 uint; F1 uint64 }
type GenS5 struct { F0 bool }
func genFn0(p0 uint16, p1 string) int64 {
    var v7 string = ""
    v8 := GenS1{F0: 230, F1: 428}
    v9 := []int{-2147483648, 9223372036854775807}
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 42
}
func genFn1(p0 uint16, p1 int16, p2 string) string {
    v10 := GenS2{F0: -1}
    for fi11 := 0; fi11 < 13; fi11++ {
        v12 := []int{(fi11 << ((65535) & 63)), (fi11 >> ((65535) & 63))}
        p1 += p1
        _ = fi11
        _ = v12
    }
    v13 := GenS3{F0: 0, F1: false, F2: 2}
    for fi14 := 0; fi14 < 12; fi14++ {
        v15 := []int{(fi14 % ((fi14) | 1))}
        var v16 int32 = 1
        _ = fi14
        _ = v15
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v13
    return p2
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint16 = 65535
    var v3 uint = 2
    var v4 float64 = 0.0
    var v5 float64 = v4
    var v6 uint8 = 42
    for fi17 := 0; fi17 < 6; fi17++ {
        v2 = 255
        _ = len(v1)
        _ = fi17
    }
    v18 := GenS4{F0: (v3 | 18446744073709551615), F1: 4294967295}
    var v19 uint8 = v6
    v20 := []int{}
    var v21 uint = v3
    v20 = append(v20, -95)
    switch "\n\t" {
    case "\n\t":
        fmt.Println(v2, v6, v6)
    case "a":
        v22 := GenS5{F0: (!true)}
        _ = v22
    }
    var v23 string = "a"
    if true {
        var v24 bool = true
        _ = v24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
