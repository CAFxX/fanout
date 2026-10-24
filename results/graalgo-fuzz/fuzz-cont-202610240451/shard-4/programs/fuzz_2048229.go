package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint16; F1 uint32 }
func genFn0(p0 string) string {
    p0 = ((p0 + p0) + ("hello" + "hello"))
    switch 4294967295 {
    case 32768:
        var v4 int = -651
        _ = v4
    }
    _ = p0
    return "foo bar"
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int8 = 0
    v3 := make(map[string]string)
    for fi5 := 0; fi5 < 18; fi5++ {
        v2 = 127
        var v6 int = fi5
        switch "\n\t" {
        case "":
            var v7 uint8 = 33
            _ = v7
        case "hello":
            v8 := GenS1{F0: 2147483647, F1: true, F2: v6}
            _ = v8
        case "a":
            v6 += (fi5 * (fi5 ^ fi5))
        default:
            v9 := []int{(^(fi5)), fi5}
            _ = v9
        }
        _ = fi5
        _ = v6
    }
    v10 := []int{2, 128, 127}
    v2 += (^(v2))
    fmt.Println(v2)
    if true {
        var v11 uint = 1
        var v12 uint64 = 629
        fmt.Println(v11)
        _ = v11
        _ = v12
    } else {
        v13 := GenS2{F0: "\n\t"}
        for fi14 := 0; fi14 < 19; fi14++ {
            v10 = append(v10, (^(fi14)))
            v15 := GenS1{F0: 42, F1: true, F2: fi14}
            var v16 uint = 889
            _ = fi14
            _ = v15
            _ = v16
        }
        if ((v2 * v2) < v2) {
            v17 := v1[619]
            _ = v17
        } else {
            if (628 == 1) {
                fmt.Println(v2)
            } else {
                for fi18 := 0; fi18 < 16; fi18++ {
                    v2 -= v2
                    v2 = -124
                    var v19 uint8 = 128
                    _ = fi18
                    _ = v19
                }
                v2 = v2
            }
        }
        _ = v13
    }
    for fi20 := 0; fi20 < 8; fi20++ {
        v2 += (v2 << ((372) & 63))
        _ = fi20
    }
    v2 = v2
    v2 += (v2 - v2)
    var v21 int16 = -28
    v21 -= v21
    if (219 != 0) {
        for fi22 := 0; fi22 < 16; fi22++ {
            v23 := GenS3{F0: 2}
            _ = fi22
            _ = v23
        }
        v2 = v2
    }
    fmt.Println(v21, v21)
    v2 += 127
    v24 := GenS4{F0: 465, F1: 4294967295}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v24
    _ = v3
}
