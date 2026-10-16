package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 int64; F1 uint16; F2 string }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 bool }
func genFn0(p0 uint8, p1 int16) int32 {
    var v4 uint64 = 8
    var v5 string = "\n\t"
    switch 2147483647 {
    case 32767:
        switch 284 {
        case -9223372036854775808:
            for fi6 := 0; fi6 < 17; fi6++ {
                for fi7 := 0; fi7 < 15; fi7++ {
                    var v8 uint64 = v4
                    _ = fi7
                    _ = v8
                }
                switch -168 {
                case -1:
                    p1 = (-(p1))
                case 32768:
                    p1 += int16(fi6)
                case 579:
                    v9 := GenS1{F0: true, F1: 506}
                    _ = v9
                default:
                    v10 := []int{(fi6 & 32768)}
                    _ = v10
                }
                _ = fi6
            }
        case -2147483648:
            if true {
                v4 = v4
                var v11 int = 128
                _ = v11
            }
        }
    case -15:
        var v12 int32 = 42
        _ = v12
    case 0:
        v13 := []int{4294967295}
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 107
}
func genFn1(p0 uint64, p1 uint) int32 {
    fmt.Println(p0)
    fmt.Println(p1)
    fmt.Println(p1, p0)
    p1 = p1
    _ = p0
    _ = p1
    return int32(p0)
}

func main() {
    var v1 uint32 = 903
    var v2 int = 32768
    var v3 int8 = 42
    for fi14 := 0; fi14 < 12; fi14++ {
        v15 := []int{32767}
        fmt.Println(v2, v2)
        _ = len(v15)
        _ = fi14
        _ = v15
    }
    switch ("a" + "hello") {
    case "":
        v16 := []int{630, 0, (v2 % ((v2) | 1))}
        _ = v16
    default:
        fmt.Println(v2, v2)
    }
    switch "" {
    case "":
        v1 = v1
    default:
        var v17 uint32 = 868
        _ = v17
    }
    for fi18 := 0; fi18 < 12; fi18++ {
        fmt.Println(fi18, v1, v2)
        var v19 uint16 = 2
        _ = fi18
        _ = v19
    }
    v20 := GenS2{F0: 42, F1: 581, F2: "hello"}
    v1 = (v1 + 2147483647)
    v21 := []int{}
    fmt.Println(v2)
    if (42 == 908) {
        if len(v21) > 0 {
            v21[0] = v2
        }
        switch v2 {
        case -1:
            fmt.Println(v1, v2)
        case -458:
            var v22 float32 = float32(0.1)
            _ = v22
        case -9223372036854775808:
            v1 = (v1 & (^(v1)))
        }
        fmt.Println(v3, v1, v2)
    }
    fmt.Println(v1)
    v23 := GenS3{F0: 91}
    v24 := GenS4{F0: (!(!(true && (!false))))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v24
    _ = v3
}
