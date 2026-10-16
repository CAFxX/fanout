package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8; F1 uint16; F2 int }
func genFn0(p0 uint, p1 int, p2 int32) int32 {
    var v5 uint16 = 2
    var v6 int64 = int64(v5)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return (p2 << ((p0) & 63))
}
func genFn1() int32 {
    var v7 int16 = 956
    v8 := []int{-40}
    for fi9 := 0; fi9 < 14; fi9++ {
        v10 := GenS1{F0: 255}
        _ = fi9
        _ = v10
    }
    v11 := GenS1{F0: 4294967295}
    _ = v11
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 int32 = 1
    v2 := make(map[int]int)
    v3 := make(map[int]uint32)
    var v4 uint32 = 269
    fmt.Println(v1, v4)
    if (false && true) {
        switch 42 {
        case -448:
            switch "hello" {
            case "a":
                v12 := []int{32767, 32768, 476, 487}
                _ = v12
            case "foo bar":
                fmt.Println(v4, v4)
            case "":
                if true {
                    v13 := GenS1{F0: 18446744073709551615}
                    var v14 int8 = 100
                    _ = v13
                    _ = v14
                }
            default:
                v4 = (v4 - v4)
            }
        case 967:
            v3[9223372036854775807] = 93
        case -702:
            fmt.Println(v1)
        }
        delete(v3, 128)
    } else {
        v15 := GenS2{F0: true}
        _ = v15
    }
    var v16 int8 = 100
    v4 -= (v4 ^ 42)
    if ("a" == "") {
        fmt.Println(v1)
        v17 := GenS3{F0: 1, F1: 255, F2: 1}
        _ = v17
    }
    fmt.Println(v1, v16, v4)
    for fi18 := 0; fi18 < 13; fi18++ {
        v19 := GenS3{F0: (v16 / ((v16) | 1)), F1: 2, F2: fi18}
        v20 := v2[fi18]
        fmt.Println(v4, v4, fi18)
        _ = fi18
        _ = v19
        _ = v20
    }
    var v21 int32 = (^(v1))
    var v22 string = "hello"
    if false {
        fmt.Println(v22, v16)
        v23 := []int{-2147483648, 449, 0}
        _ = v23
    }
    v22 = ""
    if false {
        var v24 uint16 = 697
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
