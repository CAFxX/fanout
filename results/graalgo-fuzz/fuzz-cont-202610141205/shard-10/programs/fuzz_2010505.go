package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint; F1 int32 }

func main() {
    v1 := make(map[int]bool)
    var v2 uint8 = 255
    v3 := make(map[int]int)
    v4 := make(map[int]int)
    v5 := []string{("" + "foo bar"), ""}
    var v6 int16 = -1
    v6 = 0
    if (2147483647 < 0) {
        v7 := GenS1{F0: "foo bar"}
        var v8 uint = 146
        _ = v7
        _ = v8
    }
    v9 := GenS2{F0: 127, F1: 723}
    switch -2147483648 {
    case 128:
        v10 := GenS3{F0: 42}
        _ = v10
    case 32767:
        if false {
            if ((false || (!(true && true))) && true) {
                if (!true) {
                    var v11 int16 = -1
                    var v12 int = 127
                    v13 := GenS3{F0: 4294967295}
                    _ = v11
                    _ = v12
                    _ = v13
                }
            }
            var v14 int8 = 0
            _ = v14
        } else {
            v5 = append(v5, "hello")
            v6 += v6
        }
    default:
        var v15 bool = true
        _ = v15
    }
    v6 = v6
    delete(v3, int(v6))
    for fi16 := 0; fi16 < 10; fi16++ {
        var v17 bool = (255 >= 255)
        for fi18 := 0; fi18 < 9; fi18++ {
            _ = len(v5)
            _ = fi18
        }
        switch fi16 {
        case -407:
            switch "a" {
            case "hello":
                v3[fi16] = fi16
            }
        default:
            v17 = v17
        }
        _ = fi16
        _ = v17
    }
    var v19 uint = 42
    var v20 int16 = v6
    if len(v5) > 0 {
        v5[0] = "a"
    }
    var v21 float32 = 0.1
    v2 = 42
    var v22 bool = true
    if false {
        var v23 uint16 = 909
        var v24 int32 = -927
        v25 := GenS4{F0: (^(v19)), F1: v24}
        _ = v23
        _ = v24
        _ = v25
    } else {
        v22 = (-128 >= 1)
    }
    v4[-305] = 2147483647
    switch "\n\t" {
    case "\n\t":
        v6 += (v20 / ((0) | 1))
    default:
        var v26 int16 = 508
        _ = v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
