package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 int64 }
type GenS2 struct { F0 int16; F1 uint64 }
func genFn0(p0 int, p1 int64, p2 uint16) uint32 {
    fmt.Println(p2)
    var v4 float64 = float64(((0.1 / -2.25) * -2.25))
    for fi5 := 0; fi5 < 7; fi5++ {
        for fi6 := 0; fi6 < 15; fi6++ {
            v7 := GenS1{F0: 829, F1: uint(p2), F2: 0}
            if true {
                v8 := []int{p0}
                v9 := GenS1{F0: 2147483647, F1: 0, F2: p1}
                _ = v8
                _ = v9
            }
            switch 127 {
            case 9223372036854775807:
                p0 += p0
            }
            _ = fi6
            _ = v7
        }
        var v10 string = "hello"
        v11 := []int{p0}
        _ = fi5
        _ = v10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 778
}
func genFn1() uint8 {
    v12 := []int{-73}
    var v13 int16 = 32767
    _ = v12
    _ = v13
    return 255
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint16 = 2
    var v3 int64 = 763
    switch ("" + "a") {
    case "hello":
        var v14 float32 = 3.14159
        _ = v14
    case "foo bar":
        v15 := []int{127, 174}
        _ = v15
    case "\n\t":
        switch -1 {
        case -459:
            var v16 uint = 255
            _ = v16
        case -247:
            v3 = -1
        case 32768:
            var v17 uint8 = 124
            _ = v17
        }
    }
    _ = len(v1)
    fmt.Println(v3, v2)
    var v18 int64 = v3
    if (42 >= 0) {
        if ((-128 <= 42) && true) {
            fmt.Println(v2, v2, v2)
        } else {
            fmt.Println(v18)
        }
    }
    for fi19 := 0; fi19 < 8; fi19++ {
        v18 += v18
        delete(v1, 2)
        if (2 == v2) {
            fmt.Println(v18)
            v20 := GenS2{F0: 172, F1: 4294967295}
            switch "hello" {
            case "":
                v21 := v1[genFn0(0, (v3 * v3), 176)]
                _ = v21
            default:
                v22 := GenS2{F0: 32767, F1: 2}
                _ = v22
            }
            _ = v20
        } else {
            fmt.Println(v3, v18, fi19)
        }
        _ = fi19
    }
    fmt.Println(v3)
    v23 := GenS2{F0: -756, F1: 4294967295}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v2
    _ = v23
    _ = v3
}
