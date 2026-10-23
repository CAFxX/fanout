package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 int }
func genFn0(p0 bool, p1 uint8) uint16 {
    var v4 float64 = (1e10 + 1.5)
    for fi5 := 0; fi5 < 20; fi5++ {
        v6 := []int{(fi5 << ((816) & 63))}
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = p1
    _ = v4
    return 170
}

func main() {
    v1 := []string{"foo bar", ("\n\t" + (("" + ("" + "a")) + "\n\t")), "a"}
    var v2 float64 = (0.1 - 1.5)
    var v3 float64 = v2
    var v7 int64 = 42
    v1 = append(v1, "\n\t")
    v2 = 1.5
    var v8 int64 = v7
    var v9 uint32 = 0
    var v10 uint = 4294967295
    _ = len(v1)
    fmt.Println(v3)
    for fi11 := 0; fi11 < 8; fi11++ {
        v12 := GenS1{F0: v7}
        fmt.Println(v8, fi11, fi11)
        for fi13 := 0; fi13 < 6; fi13++ {
            var v14 int64 = v8
            var v15 int = fi11
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = fi11
        _ = v12
    }
    v10 -= 0
    fmt.Println(v9)
    if true {
        v10 += 2
        for fi16 := 0; fi16 < 18; fi16++ {
            var v17 uint16 = 672
            switch 127 {
            case 698:
                for fi18 := 0; fi18 < 3; fi18++ {
                    if len(v1) > 0 {
                        v1[0] = "foo bar"
                    }
                    _ = fi18
                }
            case -1:
                _ = len(v1)
            case 42:
                var v19 uint64 = 296
                _ = v19
            default:
                v8 += (^(v8))
            }
            _ = fi16
            _ = v17
        }
    } else {
        var v20 float32 = 0.1
        if (-631 < -1) {
            v21 := GenS2{F0: (v9 | v9), F1: 2147483647}
            _ = v21
        }
        var v22 uint64 = 0
        _ = v20
        _ = v22
    }
    _ = len(v1)
    switch "a" {
    case "a":
        var v23 int = -1
        _ = v23
    case "foo bar":
        fmt.Println(v10, v10, v8)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
