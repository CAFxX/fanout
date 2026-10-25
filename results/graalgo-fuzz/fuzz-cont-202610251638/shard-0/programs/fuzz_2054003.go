package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint; F1 string; F2 int }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]string)
    v3 := make(map[int]bool)
    v4 := make(map[int]uint32)
    v5 := make(map[int]string)
    v6 := GenS1{F0: 1}
    v7 := v4[-9223372036854775808]
    switch "hello" {
    case "\n\t":
        for fi8 := 0; fi8 < 16; fi8++ {
            if (v7 > v7) {
                var v9 float32 = (0.1 / 1.5)
                _ = v9
            } else {
                switch (fi8 + fi8) {
                case 473:
                    v10 := []int{(-(fi8)), 42, fi8, 42}
                    _ = v10
                case -9223372036854775808:
                    fmt.Println(fi8, fi8)
                default:
                    if true {
                        var v11 int8 = -1
                        _ = v11
                    }
                }
            }
            _ = fi8
        }
    case "":
        var v12 float32 = ((1e10 * 1.5) / 1.0)
        _ = v12
    case "a":
        fmt.Println(v7)
    default:
        v7 += 82
    }
    var v13 int64 = -9223372036854775808
    if false {
        if false {
            for fi14 := 0; fi14 < 2; fi14++ {
                v15 := v5[(fi14 >> ((1) & 63))]
                _ = fi14
                _ = v15
            }
            if (!(!(2147483647 > (v7 >> ((2) & 63))))) {
                var v16 uint64 = 1
                var v17 int64 = v13
                v18 := []int{}
                _ = v16
                _ = v17
                _ = v18
            } else {
                var v19 int16 = 0
                v19 = v19
                if false {
                    v20 := GenS1{F0: 536}
                    v21 := []int{4294967295, 127, 127, -969}
                    _ = v20
                    _ = v21
                } else {
                    var v22 float64 = 1e10
                    v23 := v4[-9223372036854775808]
                    switch -1 {
                    case 1:
                        var v24 int32 = 1
                        _ = v24
                    case 42:
                        fmt.Println(v13)
                    case 128:
                        v25 := []int{0}
                        _ = v25
                    }
                    _ = v22
                    _ = v23
                }
                _ = v19
            }
            delete(v2, 127)
        }
        v26 := GenS1{F0: 404}
        var v27 bool = ((-1 <= -1) && (false && false))
        _ = v26
        _ = v27
    }
    v28 := []int{202}
    if ((v13 >> ((230) & 63)) == 0) {
        var v29 uint32 = (v7 % ((v7) | 1))
        v13 = -1
        _ = v29
    } else {
        fmt.Println(v13)
        var v30 bool = true
        _ = v30
    }
    if (376 > 201) {
        _ = len(v2)
        v13 -= (v13 ^ (v13 + (v13 >> ((42) & 63))))
    }
    var v31 uint8 = 128
    switch ("hello" + "foo bar") {
    case "a":
        if true {
            v13 = (v13 ^ 670)
        }
    case "":
        v32 := GenS1{F0: 18446744073709551615}
        _ = v32
    }
    v33 := GenS2{F0: 65535, F1: ("" + ""), F2: -740}
    v13 -= (v13 << ((255) & 63))
    var v34 uint64 = 927
    var v35 int16 = 32767
    v36 := v5[127]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v2
    _ = v28
    _ = v3
    _ = v31
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
