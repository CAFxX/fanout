package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 uint8 }

func main() {
    var v1 int64 = 618
    v2 := make(map[string]string)
    v3 := make(map[string]bool)
    var v4 bool = (!false)
    v5 := GenS1{F0: 102, F1: 100}
    for fi6 := 0; fi6 < 20; fi6++ {
        if (0 >= 1) {
            v4 = (-575 < 0)
            v3["hello"] = (!(!false))
        }
        _ = fi6
    }
    var v7 uint8 = 127
    if true {
        for fi8 := 0; fi8 < 15; fi8++ {
            var v9 int64 = v1
            _ = fi8
            _ = v9
        }
    } else {
        var v10 int64 = (^(v1))
        _ = v10
    }
    if v4 {
        for fi11 := 0; fi11 < 6; fi11++ {
            v12 := GenS2{F0: 973, F1: (127 > 0)}
            _ = fi11
            _ = v12
        }
        v13 := GenS3{F0: (^(v7))}
        switch 32767 {
        case 32767:
            for fi14 := 0; fi14 < 11; fi14++ {
                v15 := GenS3{F0: v7}
                var v16 uint64 = 42
                _ = fi14
                _ = v15
                _ = v16
            }
        case -155:
            var v17 uint64 = 1
            _ = v17
        case 127:
            switch "foo bar" {
            case "hello":
                for fi18 := 0; fi18 < 11; fi18++ {
                    var v19 float32 = 0.0
                    v20 := GenS3{F0: (v7 + v7)}
                    v21 := []int{(fi18 % (((^(fi18))) | 1))}
                    _ = fi18
                    _ = v19
                    _ = v20
                    _ = v21
                }
            default:
                v1 = -1
            }
        default:
            v22 := GenS2{F0: 0, F1: false}
            _ = v22
        }
        _ = v13
    }
    v7 -= (v7 / ((v7) | 1))
    fmt.Println(v4, v4)
    switch "foo bar" {
    case "hello":
        var v23 int = 1
        _ = v23
    case "\n\t":
        v7 -= (v7 ^ v7)
    }
    fmt.Println(v1, v4, v7)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
