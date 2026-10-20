package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 uint32 }
type GenS2 struct { F0 uint64; F1 int64 }
func genFn0(p0 int16) string {
    var v7 int32 = 189
    v8 := []int{942, 314, 127, 128}
    _ = p0
    _ = v7
    _ = v8
    return ""
}

func main() {
    var v1 uint = 255
    v2 := []string{"hello"}
    var v3 int32 = -2147483648
    var v4 uint32 = 2147483648
    v5 := []string{"a", "a", "foo bar"}
    v6 := make(map[string]bool)
    var v9 int8 = 127
    v6[""] = ((!false) && true)
    fmt.Println(v3, v4, v9)
    _ = len(v5)
    var v10 int = 128
    _ = len(v5)
    var v11 uint = v1
    fmt.Println(v4)
    v12 := GenS1{F0: 716, F1: 982, F2: (v4 | (v4 - (^(v4))))}
    for fi13 := 0; fi13 < 16; fi13++ {
        if true {
            var v14 int = 0
            var v15 int16 = 1000
            var v16 bool = false
            _ = v14
            _ = v15
            _ = v16
        }
        _ = fi13
    }
    var v17 int64 = 757
    var v18 int16 = -640
    for fi19 := 0; fi19 < 7; fi19++ {
        var v20 uint32 = v4
        if (v11 != (v1 / (((^(v1))) | 1))) {
            v1 = (v1 * v11)
        } else {
            var v21 uint = v11
            _ = v21
        }
        _ = fi19
        _ = v20
    }
    if ((v3 >> (((v11 / ((v11) | 1))) & 63)) == 42) {
        v4 = v4
        switch 1 {
        case -613:
            v3 = (v3 << (((v1 - v11)) & 63))
        case 1:
            v18 += v18
        case 789:
            v4 -= 2147483648
        }
    }
    for fi22 := 0; fi22 < 7; fi22++ {
        for fi23 := 0; fi23 < 2; fi23++ {
            v24 := GenS1{F0: (fi23 / ((127) | 1)), F1: 0, F2: (^(v4))}
            _ = fi23
            _ = v24
        }
        if true {
            var v25 int16 = int16(v1)
            fmt.Println(v18)
            _ = v25
        } else {
            if ((v9 + (v9 >> ((v1) & 63))) < (v9 ^ v9)) {
                v4 = 65535
                v10 -= (^(fi22))
            }
            v26 := GenS2{F0: 42, F1: v17}
            switch "hello" {
            case "foo bar":
                v27 := GenS1{F0: v10, F1: 255, F2: 2}
                _ = v27
            }
            _ = v26
        }
        _ = fi22
    }
    v28 := GenS2{F0: 0, F1: (^(v17))}
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v17
    _ = v18
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
