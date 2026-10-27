package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 int; F2 string }
type GenS3 struct { F0 int64; F1 int }
type GenS4 struct { F0 string; F1 uint; F2 uint64 }

func main() {
    var v1 uint = 421
    var v2 uint = 0
    v3 := []int{}
    v4 := GenS1{F0: 583}
    fmt.Println(v2, v2)
    v5 := GenS2{F0: ("foo bar" + ""), F1: 4294967295, F2: "foo bar"}
    var v6 uint16 = 915
    v7 := GenS2{F0: "\n\t", F1: 1, F2: ""}
    for fi8 := 0; fi8 < 10; fi8++ {
        v1 += uint(v6)
        switch "" {
        case "\n\t":
            switch fi8 {
            case -712:
                v2 = v1
            default:
                for fi9 := 0; fi9 < 11; fi9++ {
                    v6 += v6
                    var v10 uint8 = 127
                    v11 := GenS3{F0: -1, F1: int(v1)}
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            }
        case "a":
            v6 += (v6 >> ((65535) & 63))
        }
        v1 = 4294967295
        _ = fi8
    }
    v12 := GenS4{F0: "hello", F1: v2, F2: 18446744073709551615}
    v2 = (v1 | v2)
    if false {
        v6 = 255
        fmt.Println(v6)
    }
    var v13 uint64 = 1
    for fi14 := 0; fi14 < 20; fi14++ {
        v3 = append(v3, (fi14 | (fi14 % ((0) | 1))))
        _ = fi14
    }
    v1 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
