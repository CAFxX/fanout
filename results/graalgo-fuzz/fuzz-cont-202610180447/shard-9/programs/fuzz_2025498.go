package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint; F1 int64; F2 uint16 }
type GenS3 struct { F0 int8; F1 int8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 bool; F1 int }
func genFn0(p0 int64, p1 int, p2 bool) uint {
    fmt.Println(p0)
    var v4 uint = 50
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return (v4 - 65535)
}

func main() {
    v1 := make(map[int]bool)
    var v2 int = -9223372036854775808
    var v3 int32 = 0
    if (v3 < v3) {
        v5 := GenS1{F0: (v2 * v2)}
        v6 := GenS2{F0: 479, F1: 1, F2: 325}
        _ = v5
        _ = v6
    } else {
        for fi7 := 0; fi7 < 20; fi7++ {
            v8 := GenS3{F0: 127, F1: -64}
            _ = fi7
            _ = v8
        }
    }
    for fi9 := 0; fi9 < 7; fi9++ {
        fmt.Println(v3, v2, fi9)
        v1[v2] = (-1 > 127)
        _ = fi9
    }
    if (!true) {
        for fi10 := 0; fi10 < 3; fi10++ {
            switch v2 {
            case 9223372036854775807:
                var v11 uint16 = 982
                _ = v11
            case 4294967295:
                var v12 int = (v2 + fi10)
                _ = v12
            case -2147483648:
                var v13 bool = (2147483648 != 0)
                _ = v13
            }
            _ = fi10
        }
    }
    v3 = (v3 - (-(v3)))
    if (0 >= 845) {
        fmt.Println(v3, v3)
    } else {
        var v14 uint16 = 2
        _ = v14
    }
    v15 := GenS4{F0: 2}
    v2 = v2
    for fi16 := 0; fi16 < 13; fi16++ {
        v3 = v3
        switch v2 {
        case 395:
            v3 -= v3
        }
        _ = fi16
    }
    for fi17 := 0; fi17 < 11; fi17++ {
        if true {
            v18 := GenS3{F0: -1, F1: 73}
            _ = v18
        } else {
            for fi19 := 0; fi19 < 1; fi19++ {
                var v20 int = 4294967295
                _ = fi19
                _ = v20
            }
            v3 -= -75
            v21 := GenS3{F0: 1, F1: 0}
            _ = v21
        }
        _ = fi17
    }
    v2 = v2
    v22 := GenS5{F0: (593 < v2), F1: (v2 & (v2 / ((v2) | 1)))}
    switch "a" {
    case "hello":
        v23 := v1[v2]
        _ = v23
    default:
        v2 = (v2 / ((v2) | 1))
    }
    v24 := []int{1, 32, 32768, (v2 * (v2 % ((-9223372036854775808) | 1)))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v2
    _ = v22
    _ = v24
    _ = v3
}
