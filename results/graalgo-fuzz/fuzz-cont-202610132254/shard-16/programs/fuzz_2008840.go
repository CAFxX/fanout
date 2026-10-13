package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64; F1 uint64 }
type GenS3 struct { F0 int16; F1 bool; F2 int16 }
type GenS4 struct { F0 int8 }

func main() {
    var v1 int = 128
    v2 := make(map[string]uint32)
    var v3 bool = true
    var v4 uint16 = 0
    v5 := []int{}
    v1 = (-(v1))
    v5 = append(v5, (v1 + (v1 >> ((65535) & 63))))
    var v6 uint16 = 414
    v7 := GenS1{F0: "a"}
    v8 := GenS1{F0: ("foo bar" + "a")}
    if (616 == 715) {
        fmt.Println(v3, v1, v1)
        fmt.Println(v3)
        if (82 != -1) {
            v9 := GenS2{F0: 1, F1: 0}
            _ = v9
        } else {
            v10 := GenS2{F0: 42, F1: 42}
            v1 = -554
            for fi11 := 0; fi11 < 8; fi11++ {
                v6 = v6
                delete(v2, "")
                _ = fi11
            }
            _ = v10
        }
    } else {
        switch 0 {
        case -782:
            fmt.Println(v3, v3)
        case 128:
            v12 := GenS3{F0: 32767, F1: v3, F2: 489}
            _ = v12
        case 4294967295:
            v4 += v6
        }
        if v3 {
            v13 := GenS4{F0: -1}
            _ = v13
        } else {
            v3 = (42 <= 65535)
            fmt.Println(v4)
        }
    }
    var v14 int64 = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
