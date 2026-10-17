package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool; F1 int32; F2 uint16 }
type GenS4 struct { F0 uint32; F1 uint32 }

func main() {
    var v1 uint8 = 2
    v2 := []string{}
    v3 := make(map[uint32]int)
    if (377 <= 2147483648) {
        for fi4 := 0; fi4 < 15; fi4++ {
            switch "" {
            case "a":
                v1 = 42
            case "":
                v5 := GenS1{F0: 1, F1: fi4}
                _ = v5
            }
            _ = fi4
        }
        var v6 uint8 = 127
        _ = v6
    }
    var v7 int16 = 582
    for fi8 := 0; fi8 < 19; fi8++ {
        v9 := GenS2{F0: true}
        _ = fi8
        _ = v9
    }
    var v10 uint = 0
    for fi11 := 0; fi11 < 6; fi11++ {
        if true {
            v7 = -1
            v12 := GenS3{F0: ("hello" == "foo bar"), F1: -118, F2: 255}
            v7 -= (-(v7))
            _ = v12
        }
        _ = fi11
    }
    v13 := GenS2{F0: (42 <= 582)}
    v14 := GenS4{F0: 667, F1: 2}
    delete(v3, 1)
    v1 += (v1 << ((v10) & 63))
    if true {
        v15 := GenS3{F0: true, F1: -2147483648, F2: 2}
        _ = v15
    }
    for fi16 := 0; fi16 < 19; fi16++ {
        switch "hello" {
        case "\n\t":
            for fi17 := 0; fi17 < 19; fi17++ {
                fmt.Println(fi17, v1)
                _ = fi17
            }
        default:
            fmt.Println(fi16, v10)
        }
        v10 = 65535
        _ = fi16
    }
    var v18 uint8 = v1
    v7 += v7
    for fi19 := 0; fi19 < 4; fi19++ {
        v1 = v18
        fmt.Println(v7)
        _ = fi19
    }
    v10 -= 42
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v7
}
