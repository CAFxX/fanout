package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 uint64 }

func main() {
    var v1 int64 = 42
    v2 := make(map[string]int)
    v3 := make(map[int]uint32)
    v1 = v1
    v4 := []int{32767, 128}
    switch -424 {
    case -968:
        for fi5 := 0; fi5 < 14; fi5++ {
            fmt.Println(v1, v1)
            fmt.Println(v1, fi5)
            _ = fi5
        }
    case 649:
        var v6 int64 = (v1 * v1)
        _ = v6
    case 42:
        delete(v3, -544)
    }
    v7 := GenS1{F0: 42, F1: -305, F2: 42}
    v8 := GenS1{F0: 32767, F1: 1, F2: 1}
    v9 := GenS1{F0: 0, F1: v1, F2: 1}
    if false {
        for fi10 := 0; fi10 < 6; fi10++ {
            switch ("foo bar" + "") {
            case "a":
                switch -935 {
                case -123:
                    var v11 int16 = 533
                    _ = v11
                default:
                    var v12 string = "a"
                    _ = v12
                }
            case "\n\t":
                fmt.Println(v1, fi10)
            default:
                var v13 int16 = 1000
                _ = v13
            }
            v1 -= (v1 ^ int64(fi10))
            var v14 uint8 = 2
            _ = fi10
            _ = v14
        }
    }
    if true {
        if true {
            for fi15 := 0; fi15 < 18; fi15++ {
                var v16 uint16 = 255
                _ = fi15
                _ = v16
            }
            v1 -= (-(v1))
        }
        v1 -= (-(v1))
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
