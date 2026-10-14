package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 int }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]uint32)
    var v3 int8 = 42
    var v4 uint32 = 0
    v5 := []int{0, 27, 187}
    v6 := GenS1{F0: 255}
    _ = len(v2)
    fmt.Println(v3, v3)
    var v7 int = 4294967295
    var v8 int = (-(v7))
    var v9 uint16 = 596
    fmt.Println(v9, v8)
    switch 32768 {
    case 127:
        var v10 int = (v7 % (((v7 & v7)) | 1))
        _ = v10
    default:
        for fi11 := 0; fi11 < 18; fi11++ {
            if ((v3 % (((-(v3))) | 1)) != (v3 << ((42) & 63))) {
                if false {
                    v5 = append(v5, v8)
                    switch "a" {
                    case "a":
                        for fi12 := 0; fi12 < 19; fi12++ {
                            fmt.Println(v7, v7, v3)
                            _ = fi12
                        }
                    case "\n\t":
                        var v13 int64 = 1
                        _ = v13
                    default:
                        _ = len(v2)
                    }
                    switch "\n\t" {
                    case "a":
                        if false {
                            v1[("hello" + "hello")] = 1
                            v14 := GenS1{F0: 255}
                            _ = v14
                        }
                    case "foo bar":
                        v15 := GenS2{F0: ("foo bar" + "foo bar"), F1: -9223372036854775808}
                        _ = v15
                    }
                }
                v7 = v7
                fmt.Println(v7, v4, v9)
            }
            _ = fi11
        }
    }
    if (872 != 248) {
        v4 = 447
        _ = len(v5)
    }
    var v16 uint32 = v4
    v17 := GenS2{F0: ("hello" + ("" + "\n\t")), F1: v8}
    for fi18 := 0; fi18 < 5; fi18++ {
        v19 := GenS1{F0: 18446744073709551615}
        var v20 int64 = 1
        _ = fi18
        _ = v19
        _ = v20
    }
    var v21 int16 = 796
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
