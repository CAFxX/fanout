package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64; F1 int16 }

func main() {
    var v1 int16 = 0
    v2 := []uint32{42, 283, 2147483648}
    v3 := make(map[string]bool)
    var v4 int64 = 9223372036854775807
    var v5 int8 = -128
    if (false || true) {
        v3[""] = false
        if true {
            if ((!false) && ((v1 | (v1 | (v1 - 1))) >= (v1 % ((42) | 1)))) {
                var v6 uint = 610
                switch -585 {
                case -955:
                    v6 -= v6
                case 1:
                    if (0 == 328) {
                        v7 := GenS1{F0: v6}
                        v8 := GenS1{F0: v6}
                        fmt.Println(v5)
                        _ = v7
                        _ = v8
                    }
                case -165:
                    _ = len(v3)
                default:
                    if false {
                        var v9 int8 = 103
                        v5 = (v5 << (((v6 & 255)) & 63))
                        _ = v9
                    }
                }
                _ = v6
            } else {
                switch "a" {
                case "hello":
                    switch "a" {
                    case "hello":
                        var v10 int64 = v4
                        _ = v10
                    case "a":
                        v3[""] = false
                    case "foo bar":
                        v4 = -731
                    }
                case "foo bar":
                    _ = len(v2)
                default:
                    if len(v2) > 0 {
                        v2[0] = 2147483648
                    }
                }
                for fi11 := 0; fi11 < 12; fi11++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi11
                }
                fmt.Println(v1, v5, v1)
            }
        }
        delete(v3, "hello")
    }
    v2 = append(v2, 2)
    v5 += -70
    v12 := GenS2{F0: 2, F1: v1}
    var v13 int = 32767
    v1 = (^(v1))
    var v14 string = "hello"
    var v15 int32 = -610
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
