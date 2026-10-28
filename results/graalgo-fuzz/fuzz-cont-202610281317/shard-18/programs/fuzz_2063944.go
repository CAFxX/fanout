package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 uint8 }
type GenS2 struct { F0 bool }
func genFn0(p0 int32, p1 bool, p2 int8) int {
    for fi4 := 0; fi4 < 6; fi4++ {
        switch 2147483647 {
        case 899:
            p1 = (42 > 186)
        default:
            fmt.Println(p0, fi4, p1)
        }
        _ = fi4
    }
    v5 := GenS1{F0: p1, F1: (627 != 430), F2: 1}
    v6 := GenS1{F0: true, F1: p1, F2: 246}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return -9223372036854775808
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]string)
    var v3 uint = 357
    var v7 float64 = 0.1
    fmt.Println(v7, v7)
    var v8 bool = (!false)
    if false {
        v3 -= 410
        switch 32767 {
        case 9223372036854775807:
            v9 := GenS1{F0: (-32768 <= 602), F1: (!true), F2: 0}
            _ = v9
        }
        if (9223372036854775807 <= 9223372036854775807) {
            if (!(-1 == 2147483647)) {
                fmt.Println(v7)
                fmt.Println(v8, v3)
            }
        } else {
            v7 = (v7 - v7)
            var v10 float32 = -2.25
            for fi11 := 0; fi11 < 4; fi11++ {
                fmt.Println(v10, v7)
                _ = fi11
            }
            _ = v10
        }
    } else {
        var v12 string = ("foo bar" + "\n\t")
        _ = v12
    }
    v13 := GenS1{F0: (v8 && (!(88 <= 127))), F1: v8, F2: 128}
    v14 := GenS2{F0: (!(v8 || v8))}
    v15 := v1[471]
    v3 = (v3 + 0)
    switch -104 {
    case 32767:
        v15 = v15
    }
    fmt.Println(v3, v7)
    _ = len(v2)
    v3 = (v3 / (((v3 + 1)) | 1))
    if (!true) {
        var v16 int = 718
        v15 = v15
        _ = v16
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
