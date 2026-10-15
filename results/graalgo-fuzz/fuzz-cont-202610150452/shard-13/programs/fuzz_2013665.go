package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 string }
type GenS2 struct { F0 string; F1 int; F2 uint16 }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]bool)
    var v3 uint8 = 42
    for fi4 := 0; fi4 < 13; fi4++ {
        var v5 int32 = -1
        _ = fi4
        _ = v5
    }
    for fi6 := 0; fi6 < 3; fi6++ {
        v3 = 207
        _ = fi6
    }
    var v7 uint = 1
    v3 = (v3 % (((v3 & v3)) | 1))
    v7 = v7
    for fi8 := 0; fi8 < 3; fi8++ {
        for fi9 := 0; fi9 < 3; fi9++ {
            var v10 uint8 = v3
            v7 = 2
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    v11 := []int{-1}
    switch 9223372036854775807 {
    case 32768:
        switch "hello" {
        case "hello":
            v7 = 0
        case "":
            for fi12 := 0; fi12 < 10; fi12++ {
                for fi13 := 0; fi13 < 4; fi13++ {
                    v14 := GenS1{F0: -290, F1: v7, F2: "\n\t"}
                    var v15 uint64 = 4294967295
                    fmt.Println(fi13, fi13)
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                _ = fi12
            }
        default:
            var v16 int64 = -9223372036854775808
            _ = v16
        }
    case 954:
        _ = len(v11)
    case 903:
        if true {
            v17 := v1[944]
            var v18 bool = ("a" >= ("hello" + ""))
            _ = v17
            _ = v18
        } else {
            v7 += v7
            v7 = v7
            v2[914] = true
        }
    default:
        var v19 int8 = -10
        _ = v19
    }
    fmt.Println(v7, v3)
    switch 32767 {
    case 127:
        v20 := GenS2{F0: "", F1: -1, F2: 255}
        _ = v20
    }
    if (9223372036854775807 >= 890) {
        delete(v1, 0)
        v3 = v3
    }
    v7 += (v7 % ((v7) | 1))
    for fi21 := 0; fi21 < 8; fi21++ {
        var v22 bool = false
        v23 := GenS1{F0: -319, F1: (v7 * v7), F2: "foo bar"}
        _ = fi21
        _ = v22
        _ = v23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v7
}
