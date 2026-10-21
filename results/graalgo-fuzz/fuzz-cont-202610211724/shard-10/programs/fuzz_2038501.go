package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 uint8 }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 uint8; F1 uint32 }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]uint32)
    var v3 string = "\n\t"
    v4 := make(map[string]string)
    if false {
        v3 = ""
        for fi5 := 0; fi5 < 17; fi5++ {
            var v6 float32 = float32(1.5)
            var v7 bool = (185 <= 1)
            _ = fi5
            _ = v6
            _ = v7
        }
        var v8 float64 = 0.0
        _ = v8
    } else {
        v9 := GenS1{F0: 9223372036854775807, F1: -1, F2: 236}
        v10 := []int{-47}
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 16; fi11++ {
        v3 = ((v3 + v3) + (v3 + "hello"))
        _ = fi11
    }
    v3 = "foo bar"
    if false {
        v12 := GenS1{F0: 128, F1: 42, F2: 42}
        switch 0 {
        case -637:
            v13 := GenS2{F0: 138, F1: -74}
            _ = v13
        default:
            v3 = ("a" + v3)
        }
        fmt.Println(v3)
        _ = v12
    }
    switch "a" {
    case "hello":
        var v14 bool = false
        _ = v14
    default:
        if (!(false && true)) {
            var v15 float32 = 1.5
            _ = v15
        }
    }
    v3 = (("" + (v3 + v3)) + "\n\t")
    if true {
        var v16 uint32 = 0
        _ = v16
    } else {
        v17 := GenS3{F0: 1, F1: 0}
        v3 = (("\n\t" + "") + v3)
        v3 = v3
        _ = v17
    }
    _ = len(v2)
    switch 32768 {
    case 4294967295:
        fmt.Println(v3)
    default:
        v3 = ((v3 + v3) + v3)
    }
    fmt.Println(v3)
    fmt.Println(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
