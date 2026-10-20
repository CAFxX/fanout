package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32; F1 uint64 }

func main() {
    var v1 uint32 = 2
    v2 := make(map[int]string)
    v3 := make(map[uint32]int)
    var v4 uint32 = v1
    var v5 uint32 = v4
    v6 := GenS1{F0: 42, F1: 127, F2: 0}
    v7 := GenS2{F0: ("a" + "\n\t")}
    if (0.0 == 1e10) {
        v8 := []int{-1, 400}
        _ = v8
    }
    if (false && (-1 >= 42)) {
        var v9 int16 = -291
        v10 := GenS1{F0: -1, F1: 56, F2: 127}
        _ = v10
        _ = v9
    } else {
        delete(v3, v4)
        var v11 int = 94
        var v12 uint16 = 32767
        _ = v11
        _ = v12
    }
    fmt.Println(v4, v1, v4)
    v13 := []int{}
    v1 -= v1
    v14 := GenS3{F0: -654, F1: 533}
    var v15 int32 = 2147483647
    if (0 == 255) {
        for fi16 := 0; fi16 < 17; fi16++ {
            fmt.Println(v5)
            _ = fi16
        }
    } else {
        v17 := GenS1{F0: -128, F1: 42, F2: 100}
        _ = v17
    }
    switch "hello" {
    case "hello":
        v18 := GenS1{F0: -128, F1: 210, F2: 35}
        _ = v18
    default:
        v5 -= (v5 / ((2) | 1))
    }
    v19 := GenS3{F0: (v15 & 42), F1: 18446744073709551615}
    if false {
        var v20 float32 = (-2.25 - 0.1)
        _ = v20
    } else {
        var v21 int = 42
        _ = v21
    }
    switch -579 {
    case 42:
        switch -1 {
        case 32767:
            for fi22 := 0; fi22 < 10; fi22++ {
                var v23 uint64 = 127
                v15 += v15
                v1 -= v1
                _ = fi22
                _ = v23
            }
        default:
            switch ("a" + "a") {
            case "a":
                var v24 int64 = 758
                _ = v24
            case "foo bar":
                for fi25 := 0; fi25 < 18; fi25++ {
                    v15 = 1
                    var v26 float32 = 0.0
                    v27 := GenS1{F0: 127, F1: 127, F2: 1}
                    _ = fi25
                    _ = v26
                    _ = v27
                }
            default:
                var v28 uint16 = 2
                _ = v28
            }
        }
    case 1:
        var v29 string = ""
        _ = v29
    }
    if false {
        var v30 int8 = 8
        v5 = v1
        for fi31 := 0; fi31 < 11; fi31++ {
            _ = len(v2)
            v15 = (-(v15))
            _ = fi31
        }
        _ = v30
    }
    for fi32 := 0; fi32 < 18; fi32++ {
        v13 = append(v13, 1)
        var v33 uint16 = 234
        _ = fi32
        _ = v33
    }
    v34 := GenS2{F0: "a"}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
