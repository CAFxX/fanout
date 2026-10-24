package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int; F1 string }

func main() {
    v1 := make(map[uint32]int)
    var v2 float32 = 1.5
    v3 := make(map[int]string)
    for fi4 := 0; fi4 < 15; fi4++ {
        v5 := GenS1{F0: false, F1: 32767}
        fmt.Println(fi4, fi4)
        _ = fi4
        _ = v5
    }
    if true {
        v2 -= 1e10
        var v6 int8 = 42
        _ = v6
    } else {
        switch 0 {
        case 9223372036854775807:
            v2 += v2
        case 42:
            _ = len(v3)
        }
        v2 -= ((float32(float64((3.14159 / v2))) / v2) - (v2 * 1.5))
        switch "\n\t" {
        case "":
            for fi7 := 0; fi7 < 7; fi7++ {
                switch fi7 {
                case 4294967295:
                    var v8 uint32 = 4294967295
                    _ = v8
                case -27:
                    v2 -= v2
                case -9223372036854775808:
                    for fi9 := 0; fi9 < 4; fi9++ {
                        v2 = v2
                        delete(v3, fi7)
                        _ = fi9
                    }
                default:
                    delete(v3, fi7)
                }
                v10 := []int{127}
                var v11 bool = (!(1 > 42))
                _ = fi7
                _ = v10
                _ = v11
            }
        case "hello":
            v2 += float32((-2.25 * 0.0))
        case "foo bar":
            for fi12 := 0; fi12 < 14; fi12++ {
                var v13 uint16 = 1
                _ = fi12
                _ = v13
            }
        }
    }
    if (127 == 128) {
        v3[-503] = "a"
        var v14 uint = 0
        var v15 uint32 = 4294967295
        _ = v14
        _ = v15
    } else {
        if false {
            fmt.Println(v2)
        } else {
            if (65535 <= 255) {
                fmt.Println(v2)
                v16 := GenS1{F0: true, F1: 42}
                _ = v16
            }
        }
    }
    switch -2147483648 {
    case 2147483647:
        var v17 int16 = 862
        _ = v17
    case 1:
        v18 := GenS2{F0: 499}
        _ = v18
    default:
        var v19 float32 = v2
        _ = v19
    }
    v20 := GenS2{F0: 2147483648}
    if (167 >= 1) {
        var v21 int64 = -1
        var v22 uint32 = 684
        v21 = 1
        _ = v21
        _ = v22
    }
    delete(v3, 2147483647)
    v23 := GenS3{F0: 128, F1: "\n\t"}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v3
}
