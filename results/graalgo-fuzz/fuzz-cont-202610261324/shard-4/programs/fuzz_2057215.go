package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 uint8 }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 string; F1 uint32 }
func genFn0() uint8 {
    v5 := []int{-9223372036854775808, -9223372036854775808, -436, -538}
    for fi6 := 0; fi6 < 10; fi6++ {
        var v7 int = (fi6 >> ((201) & 63))
        var v8 int32 = -1
        switch "a" {
        case "a":
            var v9 uint8 = 39
            _ = v9
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    var v10 uint = 0
    var v11 uint16 = 0
    v12 := GenS1{F0: -412, F1: 18446744073709551615, F2: 255}
    _ = v10
    _ = v11
    _ = v12
    _ = v5
    return 116
}

func main() {
    var v1 int8 = -1
    v2 := make(map[string]int)
    var v3 uint8 = 127
    v4 := make(map[int]uint32)
    if true {
        if true {
            var v13 uint64 = 2
            v13 = (v13 >> ((42) & 63))
            _ = v13
        }
        fmt.Println(v3, v3)
    } else {
        switch "a" {
        case "a":
            if (1 >= (v1 / (((v1 - v1)) | 1))) {
                v1 -= (v1 + v1)
            }
        case "foo bar":
            v1 += 0
        }
        v1 += v1
    }
    for fi14 := 0; fi14 < 17; fi14++ {
        v4[(fi14 << ((1) & 63))] = 1
        _ = fi14
    }
    switch "hello" {
    case "a":
        var v15 int = 2147483647
        _ = v15
    case "hello":
        v16 := GenS1{F0: 42, F1: 1, F2: (v3 << ((65535) & 63))}
        _ = v16
    default:
        fmt.Println(v3)
    }
    if true {
        v4[-2147483648] = 2147483648
        v3 += v3
        var v17 int8 = v1
        _ = v17
    }
    v1 += 2
    for fi18 := 0; fi18 < 14; fi18++ {
        switch (("a" + "foo bar") + (("hello" + ("a" + ("" + ""))) + (("a" + "foo bar") + ("a" + "hello")))) {
        case "hello":
            var v19 int64 = -9223372036854775808
            _ = v19
        }
        switch (fi18 / (((-(fi18))) | 1)) {
        case 32767:
            v20 := []int{(fi18 + (fi18 & 32767)), (^(fi18)), 32768, fi18}
            _ = v20
        case -280:
            var v21 uint8 = v3
            _ = v21
        }
        v1 = v1
        _ = fi18
    }
    v22 := GenS2{F0: 0, F1: ""}
    fmt.Println(v3, v3)
    v1 -= (v1 ^ (v1 << ((255) & 63)))
    if (!true) {
        v23 := []int{}
        _ = v23
    }
    if (!true) {
        v1 -= v1
    } else {
        if (729 <= 692) {
            v3 = (^(v3))
        } else {
            v24 := GenS1{F0: -831, F1: 879, F2: (v3 * v3)}
            _ = v24
        }
    }
    v25 := []int{130, -469, 215}
    v26 := GenS3{F0: "", F1: 0}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v22
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
