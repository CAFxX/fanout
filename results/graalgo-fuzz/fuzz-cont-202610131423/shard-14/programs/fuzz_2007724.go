package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 bool; F1 bool; F2 uint8 }

func main() {
    var v1 uint = 4294967295
    v2 := make(map[uint32]int)
    v3 := make(map[int]bool)
    fmt.Println(v1)
    switch 32767 {
    case 32768:
        switch -593 {
        case 4294967295:
            fmt.Println(v1)
        case 127:
            v4 := GenS1{F0: 32768, F1: 42}
            _ = v4
        case 32767:
            delete(v3, 32767)
        default:
            v1 = 4294967295
        }
    default:
        v1 = v1
    }
    var v5 int = -896
    var v6 uint = v1
    switch "a" {
    case "foo bar":
        delete(v2, 42)
    default:
        if (!(-1 < -1)) {
            v7 := GenS1{F0: (^(v5)), F1: -1}
            _ = v7
        }
    }
    if false {
        v6 = (v6 | v1)
        if ((!(0 < 617)) && true) {
            if true {
                v8 := GenS1{F0: v5, F1: 0}
                if ((v5 % ((v5) | 1)) >= (^(v5))) {
                    var v9 float64 = 0.1
                    _ = v9
                } else {
                    v6 = v6
                }
                v10 := GenS2{F0: (4294967295 < 1), F1: false, F2: 1}
                _ = v10
                _ = v8
            } else {
                if false {
                    fmt.Println(v6)
                    v6 = v6
                } else {
                    if (!(-2.25 < 3.14159)) {
                        v5 -= v5
                        var v11 uint64 = 792
                        v12 := []int{(v5 * 2147483647), v5, v5, (-(v5))}
                        _ = v11
                        _ = v12
                    }
                    v1 += v6
                }
            }
        } else {
            fmt.Println(v5, v1)
        }
    }
    fmt.Println(v1, v5, v1)
    var v13 uint8 = 249
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v5
    _ = v6
}
