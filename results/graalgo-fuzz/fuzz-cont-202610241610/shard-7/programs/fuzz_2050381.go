package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64; F1 int; F2 int }
func genFn0(p0 uint, p1 uint32) int64 {
    var v6 uint16 = 1
    var v7 uint8 = 111
    for fi8 := 0; fi8 < 20; fi8++ {
        v9 := []int{(fi8 & fi8), fi8, fi8, int(v6)}
        v9 = append(v9, (fi8 % ((fi8) | 1)))
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return -9223372036854775808
}

func main() {
    var v1 bool = (0.1 < 0.1)
    v2 := make(map[int]bool)
    v3 := []int{}
    var v4 uint = 42
    v5 := make(map[string]int)
    if (0 != 42) {
        for fi10 := 0; fi10 < 20; fi10++ {
            var v11 bool = ((fi10 / (((fi10 << ((v4) & 63))) | 1)) < int(v4))
            _ = fi10
            _ = v11
        }
        if v1 {
            switch "" {
            case "a":
                var v12 uint16 = 42
                _ = v12
            case "\n\t":
                delete(v2, -317)
            case "hello":
                v4 = (v4 ^ v4)
            default:
                var v13 int64 = 42
                _ = v13
            }
            var v14 int8 = -101
            v15 := GenS1{F0: (v4 >> ((162) & 63))}
            _ = v14
            _ = v15
        } else {
            var v16 int8 = 100
            _ = v16
        }
        switch "\n\t" {
        case "hello":
            v1 = v1
        case "a":
            v1 = (0 >= 128)
        }
    }
    var v17 int16 = 0
    v18 := GenS1{F0: 65535}
    v19 := GenS2{F0: true}
    switch 32767 {
    case -30:
        if v1 {
            if len(v3) > 0 {
                v3[0] = -918
            }
        }
    case 0:
        v3 = append(v3, int(v17))
    default:
        switch ("\n\t" + ("\n\t" + "")) {
        case "":
            fmt.Println(v4, v17)
        default:
            v4 -= (v4 | (v4 / (((v4 & v4)) | 1)))
        }
    }
    v20 := v2[1]
    fmt.Println(v20, v4)
    for fi21 := 0; fi21 < 10; fi21++ {
        for fi22 := 0; fi22 < 11; fi22++ {
            switch 318 {
            case 584:
                fmt.Println(v20)
            }
            fmt.Println(v1, v4, v4)
            _ = fi22
        }
        _ = fi21
    }
    if v20 {
        v23 := GenS3{F0: genFn0(228, 134), F1: 4294967295, F2: 316}
        fmt.Println(v4, v4)
        switch 9223372036854775807 {
        case 42:
            var v24 bool = v20
            _ = v24
        }
        _ = v23
    } else {
        switch ("" + ("" + ("foo bar" + "a"))) {
        case "hello":
            var v25 uint = v4
            _ = v25
        default:
            v4 -= v4
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
