package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 string; F1 bool; F2 uint16 }
type GenS3 struct { F0 int16; F1 int8; F2 int64 }
type GenS4 struct { F0 string; F1 int64 }
func genFn0() uint32 {
    var v6 uint32 = 598
    var v7 int32 = -459
    _ = v6
    _ = v7
    return v6
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int = -104
    v3 := make(map[uint32]int)
    var v4 uint32 = 0
    var v5 string = "foo bar"
    for fi8 := 0; fi8 < 7; fi8++ {
        v4 -= 653
        v5 = v5
        v9 := GenS1{F0: -1, F1: fi8}
        _ = fi8
        _ = v9
    }
    var v10 uint8 = 198
    var v11 uint8 = (v10 << ((4294967295) & 63))
    fmt.Println(v5, v11)
    for fi12 := 0; fi12 < 19; fi12++ {
        switch "hello" {
        case "":
            v13 := []int{}
            _ = v13
        default:
            v14 := []int{v2, v2, fi12, (-(fi12))}
            _ = v14
        }
        v5 = v5
        _ = fi12
    }
    v15 := GenS2{F0: (v5 + v5), F1: false, F2: 42}
    switch ("foo bar" + "foo bar") {
    case "\n\t":
        fmt.Println(v4, v4, v4)
    case "hello":
        switch (-(v2)) {
        case 32768:
            switch v2 {
            case -51:
                delete(v1, 773)
            case -270:
                v1[855] = v5
            default:
                if false {
                    fmt.Println(v11)
                } else {
                    var v16 uint16 = 1
                    v16 += (v16 - v16)
                    fmt.Println(v11, v10, v10)
                    _ = v16
                }
            }
        }
    case "a":
        fmt.Println(v2, v4, v4)
    default:
        var v17 uint16 = 32767
        _ = v17
    }
    switch v5 {
    case "hello":
        var v18 uint8 = v10
        _ = v18
    default:
        v19 := GenS3{F0: 14, F1: -1, F2: 1}
        _ = v19
    }
    var v20 string = v5
    if false {
        if false {
            var v21 uint16 = 32767
            _ = v21
        }
    } else {
        if false {
            v22 := GenS4{F0: ("" + v5), F1: -395}
            _ = v22
        } else {
            fmt.Println(v2)
        }
        switch (v5 + (v20 + v5)) {
        case "foo bar":
            v5 = v20
        case "\n\t":
            for fi23 := 0; fi23 < 9; fi23++ {
                _ = len(v3)
                _ = fi23
            }
        case "hello":
            v4 = 42
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
