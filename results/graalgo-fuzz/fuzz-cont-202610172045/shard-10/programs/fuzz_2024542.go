package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
func genFn0(p0 int8, p1 uint) string {
    var v4 string = ("hello" + "")
    p0 = (p0 + p0)
    v5 := []int{}
    var v6 uint = 1
    v6 = 1
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return v4
}

func main() {
    v1 := []string{("a" + "foo bar")}
    var v2 int8 = 100
    v3 := make(map[uint32]int)
    v2 -= 33
    for fi7 := 0; fi7 < 19; fi7++ {
        var v8 bool = true
        _ = fi7
        _ = v8
    }
    var v9 uint64 = 22
    for fi10 := 0; fi10 < 7; fi10++ {
        var v11 uint = 255
        v9 = (v9 << ((v11) & 63))
        _ = fi10
        _ = v11
    }
    var v12 uint32 = 2
    for fi13 := 0; fi13 < 14; fi13++ {
        switch "\n\t" {
        case "a":
            if (65535 < 2) {
                switch ("" + ("a" + "")) {
                case "hello":
                    var v14 int8 = (v2 / (((^(v2))) | 1))
                    _ = v14
                default:
                    var v15 int32 = -597
                    _ = v15
                }
            }
        case "":
            fmt.Println(v2)
        case "hello":
            for fi16 := 0; fi16 < 13; fi16++ {
                v2 = 42
                if (v2 <= v2) {
                    _ = len(v3)
                }
                _ = fi16
            }
        }
        _ = fi13
    }
    var v17 uint16 = 255
    delete(v3, v12)
    var v18 float64 = (1e10 - 1.5)
    for fi19 := 0; fi19 < 9; fi19++ {
        var v20 float64 = v18
        v21 := GenS1{F0: 0, F1: 532}
        _ = fi19
        _ = v20
        _ = v21
    }
    v12 = (^(v12))
    var v22 int = 32768
    fmt.Println(v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v9
}
