package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 int; F1 int16; F2 string }
type GenS3 struct { F0 int16 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint64 = 2
    var v3 uint32 = 42
    v4 := GenS1{F0: 832, F1: (89 >= 1)}
    if false {
        var v5 int32 = -1
        if false {
            fmt.Println(v2)
            v2 = v2
        } else {
            switch -2147483648 {
            case 32767:
                v5 = (v5 ^ 2147483647)
            case 9223372036854775807:
                fmt.Println(v5, v2, v5)
            default:
                fmt.Println(v2, v5)
            }
            var v6 string = "hello"
            if true {
                v7 := GenS1{F0: 1, F1: false}
                _ = v7
            }
            _ = v6
        }
        var v8 int64 = -270
        _ = v5
        _ = v8
    } else {
        for fi9 := 0; fi9 < 18; fi9++ {
            v3 = (^(v3))
            _ = fi9
        }
        var v10 bool = true
        v1[1] = v10
        _ = v10
    }
    if true {
        v2 = (v2 << ((255) & 63))
    } else {
        _ = len(v1)
        switch 340 {
        case -2147483648:
            v11 := GenS1{F0: 65535, F1: false}
            _ = v11
        case -942:
            for fi12 := 0; fi12 < 15; fi12++ {
                for fi13 := 0; fi13 < 9; fi13++ {
                    v14 := []int{fi12, fi12}
                    v14 = append(v14, fi12)
                    _ = fi13
                    _ = v14
                }
                var v15 uint64 = (v2 & v2)
                fmt.Println(v15, v2)
                _ = fi12
                _ = v15
            }
        }
    }
    v2 = (v2 ^ (v2 / ((0) | 1)))
    if (9223372036854775807 > -367) {
        switch -1 {
        case -116:
            v16 := GenS1{F0: 255, F1: true}
            _ = v16
        case 1:
            var v17 uint8 = 127
            _ = v17
        default:
            var v18 uint32 = (v3 << ((36) & 63))
            _ = v18
        }
        var v19 int32 = 2147483647
        _ = v19
    } else {
        var v20 int16 = 0
        _ = v20
    }
    for fi21 := 0; fi21 < 8; fi21++ {
        v3 -= (^(v3))
        var v22 bool = (-9223372036854775808 == -1)
        for fi23 := 0; fi23 < 16; fi23++ {
            switch "hello" {
            case "foo bar":
                var v24 uint64 = 1
                _ = v24
            case "a":
                v25 := []int{fi23}
                _ = v25
            case "hello":
                v26 := []int{(fi21 | (fi21 & (fi21 - -2147483648))), fi21, fi21, 815}
                _ = v26
            }
            _ = fi23
        }
        _ = fi21
        _ = v22
    }
    v27 := GenS2{F0: -903, F1: 267, F2: "a"}
    fmt.Println(v2, v2)
    v3 = 0
    switch -2147483648 {
    case 127:
        _ = len(v1)
    case 420:
        fmt.Println(v3, v3)
    default:
        if false {
            if (!(false && (3.14159 == -2.25))) {
                v3 = (v3 << ((269) & 63))
            }
            var v28 int16 = 0
            if (v28 <= v28) {
                _ = len(v1)
            }
            _ = v28
        } else {
            v2 -= 0
            v2 -= 18446744073709551615
        }
    }
    switch "hello" {
    case "foo bar":
        v2 = (v2 * v2)
    case "\n\t":
        for fi29 := 0; fi29 < 1; fi29++ {
            var v30 int8 = 100
            _ = fi29
            _ = v30
        }
    default:
        v31 := GenS3{F0: -1}
        _ = v31
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
}
