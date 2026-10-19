package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }

func main() {
    var v1 bool = false
    v2 := make(map[uint32]bool)
    v3 := []uint32{65535}
    var v4 int8 = 0
    var v5 bool = false
    var v6 int64 = -769
    switch 32768 {
    case 32768:
        var v7 uint16 = 2
        _ = v7
    case -9223372036854775808:
        var v8 uint16 = 547
        _ = v8
    default:
        if false {
            switch "" {
            case "\n\t":
                v9 := GenS1{F0: 35, F1: false}
                _ = v9
            default:
                switch "\n\t" {
                case "foo bar":
                    var v10 uint8 = 255
                    _ = v10
                default:
                    delete(v2, 2147483648)
                }
            }
            v11 := GenS1{F0: 65535, F1: false}
            _ = v11
        }
    }
    v4 -= (v4 >> ((721) & 63))
    switch "a" {
    case "hello":
        v4 = v4
    case "a":
        fmt.Println(v6)
    default:
        fmt.Println(v1)
    }
    if (-931 > v6) {
        v6 += -1
    } else {
        v12 := GenS1{F0: 562, F1: v5}
        _ = v12
    }
    v1 = v1
    for fi13 := 0; fi13 < 18; fi13++ {
        var v14 string = "a"
        v15 := GenS1{F0: 0, F1: true}
        fmt.Println(v6)
        _ = fi13
        _ = v14
        _ = v15
    }
    var v16 int16 = -32768
    var v17 uint32 = 951
    switch 128 {
    case 0:
        var v18 bool = (!true)
        _ = v18
    }
    for fi19 := 0; fi19 < 20; fi19++ {
        v4 = v4
        _ = fi19
    }
    for fi20 := 0; fi20 < 4; fi20++ {
        v1 = false
        var v21 float32 = (0.0 - -2.25)
        _ = fi20
        _ = v21
    }
    var v22 uint = 1
    v16 = -32768
    v16 = v16
    fmt.Println(v6, v22, v17)
    v4 = v4
    var v23 uint32 = (v17 * v17)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
