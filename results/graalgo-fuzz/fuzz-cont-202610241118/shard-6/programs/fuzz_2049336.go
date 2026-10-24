package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 int32 }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]uint32)
    var v3 int64 = 9223372036854775807
    v4 := make(map[uint32]string)
    v5 := make(map[uint32]uint32)
    var v6 uint64 = 892
    for fi7 := 0; fi7 < 15; fi7++ {
        for fi8 := 0; fi8 < 4; fi8++ {
            if false {
                switch "" {
                case "a":
                    v3 = (v3 % ((9223372036854775807) | 1))
                }
            } else {
                if (true || (!false)) {
                    v9 := GenS1{F0: 127, F1: -882, F2: 42}
                    _ = v9
                }
                v10 := []int{(fi7 ^ 268), (fi8 >> ((1) & 63)), fi7}
                _ = v10
            }
            _ = fi8
        }
        v11 := []int{(fi7 << ((67) & 63)), 0, (fi7 << ((19) & 63))}
        _ = fi7
        _ = v11
    }
    v12 := []int{9223372036854775807, -744, 32768}
    v6 -= 389
    v4[974] = ("hello" + "foo bar")
    for fi13 := 0; fi13 < 19; fi13++ {
        for fi14 := 0; fi14 < 4; fi14++ {
            var v15 int8 = -54
            _ = fi14
            _ = v15
        }
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
        _ = fi13
    }
    var v16 int8 = 47
    if (-2147483648 < 0) {
        v3 = (v3 / ((v3) | 1))
    }
    var v17 string = ""
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
