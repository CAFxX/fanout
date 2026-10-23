package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 int16 }
func genFn0(p0 int64, p1 uint) string {
    fmt.Println(p0, p0)
    var v5 float64 = 0.0
    _ = p0
    _ = p1
    _ = v5
    return "\n\t"
}

func main() {
    v1 := []string{"foo bar", ("hello" + "a"), ("\n\t" + "a")}
    var v2 uint64 = 0
    v3 := []string{"a", "foo bar", "\n\t", "foo bar"}
    v4 := make(map[string]string)
    for fi6 := 0; fi6 < 6; fi6++ {
        var v7 string = ""
        _ = fi6
        _ = v7
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if false {
        v2 += (v2 % ((v2) | 1))
        v3 = append(v3, "")
    }
    v2 = (v2 | (v2 % ((2) | 1)))
    for fi8 := 0; fi8 < 12; fi8++ {
        var v9 bool = ((^(fi8)) >= -58)
        for fi10 := 0; fi10 < 12; fi10++ {
            switch fi8 {
            case -199:
                var v11 uint = 0
                _ = v11
            case -1:
                _ = len(v1)
            }
            v2 += (v2 << ((65535) & 63))
            _ = fi10
        }
        fmt.Println(v9, v9, fi8)
        _ = fi8
        _ = v9
    }
    for fi12 := 0; fi12 < 7; fi12++ {
        fmt.Println(fi12)
        _ = fi12
    }
    for fi13 := 0; fi13 < 7; fi13++ {
        var v14 int8 = -100
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 9; fi15++ {
        if (116 != 188) {
            if (float32((0.1 + 0.1)) == 0.1) {
                v16 := GenS1{F0: (fi15 >> ((65535) & 63)), F1: (^(v2)), F2: 42}
                _ = v16
            }
        }
        _ = fi15
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
