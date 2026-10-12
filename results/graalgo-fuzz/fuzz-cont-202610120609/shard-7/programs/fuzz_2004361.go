package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
func genFn0(p0 int64, p1 bool, p2 uint32) string {
    var v6 bool = false
    v7 := GenS1{F0: 7, F1: 128}
    switch 4294967295 {
    case -782:
        if v6 {
            var v8 string = ("a" + "\n\t")
            fmt.Println(p2, v6, p1)
            for fi9 := 0; fi9 < 20; fi9++ {
                v6 = true
                _ = fi9
            }
            _ = v8
        }
    case 1:
        for fi10 := 0; fi10 < 16; fi10++ {
            fmt.Println(p1, p1, p0)
            _ = fi10
        }
    default:
        v11 := GenS1{F0: 0, F1: 32}
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return "foo bar"
}

func main() {
    var v1 float32 = -2.25
    var v2 int = -593
    var v3 int32 = 1
    v4 := make(map[int]bool)
    var v5 int16 = -32768
    v1 -= 3.14159
    v5 -= (v5 & 32767)
    if (!(true || false)) {
        v12 := GenS1{F0: 1, F1: 1}
        _ = v12
    } else {
        var v13 float32 = 1.5
        v5 -= (v5 | (^(v5)))
        var v14 int64 = 42
        _ = v13
        _ = v14
    }
    v15 := []int{}
    delete(v4, v2)
    var v16 int64 = 0
    v2 -= (v2 >> ((65535) & 63))
    var v17 uint64 = 18446744073709551615
    v18 := GenS1{F0: int8(v17), F1: 0}
    fmt.Println(v5, v3, v17)
    v5 = v5
    for fi19 := 0; fi19 < 5; fi19++ {
        for fi20 := 0; fi20 < 18; fi20++ {
            fmt.Println(fi20, fi19, v2)
            if false {
                switch "hello" {
                case "\n\t":
                    fmt.Println(fi20, v5)
                }
                fmt.Println(v2)
            }
            var v21 int64 = v16
            _ = fi20
            _ = v21
        }
        _ = fi19
    }
    for fi22 := 0; fi22 < 8; fi22++ {
        _ = len(v4)
        _ = fi22
    }
    fmt.Println(v2, v17, v2)
    if true {
        _ = len(v4)
    } else {
        v15 = append(v15, (v2 % (((v2 * (^(v2)))) | 1)))
        v3 += (v3 + v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
