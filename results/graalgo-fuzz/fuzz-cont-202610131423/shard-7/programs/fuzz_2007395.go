package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
func genFn0(p0 int32, p1 int16) uint64 {
    v5 := []int{127}
    p1 += (p1 >> ((818) & 63))
    _ = p0
    _ = p1
    _ = v5
    return 0
}
func genFn1(p0 uint32, p1 int64) int16 {
    var v6 int64 = 42
    v6 = p1
    _ = p0
    _ = p1
    _ = v6
    return -551
}

func main() {
    v1 := []string{"", "", "foo bar"}
    var v2 float32 = 0.1
    v3 := make(map[string]bool)
    v4 := make(map[string]int)
    if false {
        var v7 int32 = 414
        var v8 int8 = 127
        _ = v7
        _ = v8
    }
    var v9 int8 = 3
    for fi10 := 0; fi10 < 14; fi10++ {
        fmt.Println(v2)
        switch fi10 {
        case 127:
            var v11 uint8 = 241
            _ = v11
        case -9223372036854775808:
            v2 += v2
        default:
            v12 := GenS1{F0: (!(!(2147483647 != 95))), F1: 1}
            _ = v12
        }
        switch "hello" {
        case "":
            var v13 string = "foo bar"
            _ = v13
        case "foo bar":
            _ = len(v1)
        case "a":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        _ = fi10
    }
    var v14 float64 = -2.25
    var v15 uint64 = genFn0(1, 808)
    var v16 int = 4294967295
    var v17 int32 = 42
    for fi18 := 0; fi18 < 1; fi18++ {
        var v19 uint64 = 534
        _ = fi18
        _ = v19
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
