package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 int8; F1 uint; F2 uint32 }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint = 216
    v3 := make(map[int]string)
    for fi4 := 0; fi4 < 7; fi4++ {
        var v5 int32 = 49
        _ = fi4
        _ = v5
    }
    switch (("" + (("\n\t" + "a") + "\n\t")) + "a") {
    case "a":
        v2 += v2
    case "hello":
        if false {
            switch 407 {
            case 128:
                fmt.Println(v2)
            case 0:
                switch "\n\t" {
                case "\n\t":
                    var v6 int16 = 1000
                    _ = v6
                case "hello":
                    v7 := []int{-194, 42, 42, 2147483647}
                    _ = v7
                case "":
                    fmt.Println(v2)
                }
            }
        } else {
            for fi8 := 0; fi8 < 17; fi8++ {
                switch ("\n\t" + "\n\t") {
                case "foo bar":
                    fmt.Println(fi8, fi8)
                default:
                    v1[0] = 2
                }
                v2 = 4294967295
                _ = fi8
            }
        }
    default:
        v2 = 768
    }
    v2 -= (v2 >> ((v2) & 63))
    v2 = 2
    v9 := GenS1{F0: 1, F1: 483}
    v10 := GenS2{F0: 115, F1: v2, F2: 571}
    for fi11 := 0; fi11 < 11; fi11++ {
        switch "" {
        case "a":
            v2 += 0
        case "hello":
            v2 = 4294967295
        default:
            v2 += v2
        }
        fmt.Println(fi11, fi11)
        _ = fi11
    }
    var v12 int8 = 1
    var v13 uint8 = 172
    v14 := []int{-1, 42, -997, 32767}
    v14 = append(v14, 32767)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v9
}
