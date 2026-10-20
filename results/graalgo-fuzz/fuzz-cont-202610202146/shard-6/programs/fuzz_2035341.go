package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 uint64 }
func genFn0(p0 uint8, p1 uint64, p2 int) int {
    v4 := []int{}
    for fi5 := 0; fi5 < 12; fi5++ {
        switch 127 {
        case 32768:
            for fi6 := 0; fi6 < 13; fi6++ {
                switch "a" {
                case "\n\t":
                    var v7 int16 = 1
                    _ = v7
                case "hello":
                    _ = len(v4)
                case "a":
                    p1 = (p1 & p1)
                default:
                    var v8 float32 = 0.1
                    _ = v8
                }
                _ = fi6
            }
        case -9223372036854775808:
            var v9 uint16 = 0
            _ = v9
        }
        v10 := GenS1{F0: 0, F1: (fi5 << ((65535) & 63))}
        var v11 string = ("hello" + "a")
        _ = fi5
        _ = v10
        _ = v11
    }
    if true {
        switch p2 {
        case 32768:
            var v12 int16 = -776
            _ = v12
        case 42:
            p2 -= (p2 + 32768)
        default:
            for fi13 := 0; fi13 < 7; fi13++ {
                p1 = p1
                v14 := GenS2{F0: p1}
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi13
                _ = v14
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return (p2 << ((289) & 63))
}

func main() {
    var v1 uint = 2
    v2 := make(map[int]int)
    var v3 int32 = -1
    switch "foo bar" {
    case "hello":
        v3 = -2147483648
    case "foo bar":
        fmt.Println(v3, v3)
    }
    for fi15 := 0; fi15 < 14; fi15++ {
        v1 = v1
        _ = fi15
    }
    v1 += 4294967295
    fmt.Println(v3, v1)
    var v16 int64 = 9223372036854775807
    var v17 int16 = -1
    v18 := v2[42]
    v3 = v3
    var v19 float32 = 1e10
    v20 := v2[v18]
    if (v1 >= v1) {
        var v21 uint64 = 53
        _ = v21
    }
    v2[(v20 >> ((v1) & 63))] = (v20 << ((748) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
