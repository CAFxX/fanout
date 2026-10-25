package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }

func main() {
    v1 := []string{"hello", ("hello" + ("\n\t" + ("" + ("hello" + "\n\t")))), (("foo bar" + "foo bar") + "a")}
    var v2 bool = false
    v3 := make(map[int]bool)
    var v4 uint8 = 102
    switch 128 {
    case 2147483647:
        var v5 float64 = float64(1.5)
        _ = v5
    }
    switch "a" {
    case "foo bar":
        v6 := GenS1{F0: 65535, F1: int16(v4)}
        _ = v6
    }
    switch 878 {
    case -2147483648:
        if (0.1 != 1.5) {
            var v7 int16 = -1
            if v2 {
                fmt.Println(v2, v4, v4)
                v7 = 32767
            }
            _ = v7
        }
    case -125:
        var v8 uint16 = 69
        _ = v8
    default:
        v9 := GenS1{F0: 42, F1: 0}
        _ = v9
    }
    for fi10 := 0; fi10 < 2; fi10++ {
        v3[(^(fi10))] = ((v4 - (v4 % (((v4 << ((961) & 63))) | 1))) < (v4 ^ v4))
        _ = fi10
    }
    if v2 {
        for fi11 := 0; fi11 < 13; fi11++ {
            _ = len(v1)
            _ = fi11
        }
        var v12 float64 = (1e10 / float64((1.5 * -2.25)))
        _ = v12
    }
    var v13 int32 = -615
    if v2 {
        v14 := GenS1{F0: 32767, F1: -1}
        var v15 int16 = 1
        _ = v14
        _ = v15
    } else {
        v3[-888] = (float32(1.5) < 0.0)
        v16 := GenS1{F0: 65535, F1: 631}
        _ = v16
    }
    v2 = (688 < 42)
    var v17 int = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
