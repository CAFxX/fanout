package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 int }
type GenS3 struct { F0 string; F1 int8 }
type GenS4 struct { F0 uint16; F1 int16 }

func main() {
    var v1 uint = 4294967295
    v2 := make(map[string]int)
    var v3 string = ("hello" + ("foo bar" + ("a" + "\n\t")))
    v4 := make(map[int]uint32)
    v5 := make(map[int]string)
    v6 := make(map[uint32]string)
    v7 := GenS1{F0: 0}
    switch -2147483648 {
    case 2147483647:
        var v8 int64 = 9223372036854775807
        _ = v8
    default:
        v3 = (("foo bar" + ("foo bar" + (v3 + "\n\t"))) + v3)
    }
    if ((!true) && true) {
        var v9 uint64 = 1
        for fi10 := 0; fi10 < 7; fi10++ {
            if true {
                v11 := GenS1{F0: 128}
                var v12 float32 = (1.5 / (1e10 - 3.14159))
                _ = v11
                _ = v12
            }
            if (!false) {
                var v13 uint32 = 672
                for fi14 := 0; fi14 < 15; fi14++ {
                    v15 := GenS2{F0: true, F1: fi14}
                    v3 = v3
                    _ = fi14
                    _ = v15
                }
                var v16 int32 = 783
                _ = v13
                _ = v16
            } else {
                fmt.Println(fi10)
            }
            v1 += (v1 | v1)
            _ = fi10
        }
        fmt.Println(v1)
        _ = v9
    }
    for fi17 := 0; fi17 < 12; fi17++ {
        switch v3 {
        case "":
            var v18 uint8 = 2
            _ = v18
        case "a":
            if true {
                if true {
                    if false {
                        v1 = v1
                        v1 = 487
                    }
                } else {
                    var v19 int64 = 447
                    for fi20 := 0; fi20 < 1; fi20++ {
                        v21 := GenS3{F0: "hello", F1: -73}
                        _ = fi20
                        _ = v21
                    }
                    _ = v19
                }
                for fi22 := 0; fi22 < 14; fi22++ {
                    v23 := v5[fi22]
                    var v24 uint32 = 65535
                    _ = fi22
                    _ = v23
                    _ = v24
                }
                v1 -= v1
            }
        default:
            v1 += 1
        }
        v3 = v3
        var v25 int = fi17
        _ = fi17
        _ = v25
    }
    v1 = (v1 >> ((v1) & 63))
    var v26 int16 = 1
    v2[v3] = 32768
    v27 := GenS2{F0: (32767 != 1), F1: 4294967295}
    var v28 float64 = 1e10
    v29 := GenS4{F0: 0, F1: (v26 / ((32767) | 1))}
    switch v3 {
    case "a":
        switch 4294967295 {
        case -9223372036854775808:
            var v30 uint32 = 4294967295
            _ = v30
        }
    default:
        fmt.Println(v28, v1, v3)
    }
    v31 := []int{128, 0, -659, -2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
