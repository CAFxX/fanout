package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 uint8 }
type GenS2 struct { F0 uint16; F1 uint }
func genFn0(p0 int16) uint {
    v6 := GenS1{F0: 273, F1: ("hello" + ("\n\t" + "hello")), F2: 255}
    var v7 uint64 = 2
    v8 := []int{-169}
    switch 1 {
    case 128:
        v9 := GenS1{F0: 1, F1: ("foo bar" + "a"), F2: 42}
        _ = v9
    }
    var v10 bool = false
    _ = p0
    _ = v10
    _ = v6
    _ = v7
    _ = v8
    return 380
}
func genFn1() int {
    v11 := []int{2147483647, -717, 4294967295, 0}
    var v12 int64 = -9223372036854775808
    if (-270 < 127) {
        v13 := GenS2{F0: 2, F1: 42}
        fmt.Println(v12)
        v12 += v12
        _ = v13
    } else {
        for fi14 := 0; fi14 < 17; fi14++ {
            if len(v11) >= 2 {
                v11 = v11[1:len(v11)]
            }
            _ = fi14
        }
        var v15 uint8 = 81
        fmt.Println(v12, v15)
        _ = v15
    }
    switch "\n\t" {
    case "\n\t":
        v12 = 1
    case "":
        v11 = append(v11, 127)
    }
    if len(v11) > 0 {
        v11[0] = 1
    }
    _ = v11
    _ = v12
    return 0
}

func main() {
    var v1 float64 = 1.5
    var v2 uint16 = 65535
    v3 := make(map[uint32]int)
    var v4 string = ""
    v5 := make(map[uint32]uint32)
    switch "foo bar" {
    case "hello":
        v2 -= v2
    }
    if (!(v4 <= v4)) {
        v16 := GenS2{F0: v2, F1: 42}
        _ = v16
    }
    fmt.Println(v1)
    fmt.Println(v4, v1)
    v2 += v2
    v4 = "\n\t"
    switch 9223372036854775807 {
    case 42:
        fmt.Println(v1, v1, v4)
    case 2147483647:
        var v17 string = "hello"
        _ = v17
    case 127:
        v1 = v1
    }
    v18 := GenS2{F0: (v2 & 65535), F1: 2}
    var v19 uint = 2
    for fi20 := 0; fi20 < 5; fi20++ {
        fmt.Println(v4, v2)
        v1 = (1.5 - 0.0)
        _ = fi20
    }
    v21 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
