package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool; F1 uint16; F2 uint32 }
func genFn0(p0 uint16, p1 string, p2 uint) int16 {
    v5 := []int{32768, int(p2), -1}
    for fi6 := 0; fi6 < 4; fi6++ {
        _ = len(v5)
        for fi7 := 0; fi7 < 13; fi7++ {
            v8 := GenS1{F0: -1}
            _ = fi7
            _ = v8
        }
        var v9 int16 = -758
        _ = fi6
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 263
}

func main() {
    var v1 int64 = -9223372036854775808
    v2 := []string{}
    v3 := make(map[int]string)
    var v4 uint8 = 0
    if (!false) {
        fmt.Println(v4, v4)
        v1 = (v1 | (v1 << ((42) & 63)))
    }
    switch "hello" {
    case "foo bar":
        fmt.Println(v4)
    case "a":
        if false {
            v1 = 9223372036854775807
            v10 := GenS1{F0: v1}
            _ = v10
        } else {
            v1 = (^(v1))
            if (-589 < -1) {
                v11 := GenS2{F0: true, F1: 2, F2: 1}
                _ = v11
            }
        }
    case "\n\t":
        var v12 bool = false
        _ = v12
    default:
        v4 = v4
    }
    for fi13 := 0; fi13 < 15; fi13++ {
        v1 += (v1 * (-(v1)))
        for fi14 := 0; fi14 < 8; fi14++ {
            switch ("hello" + ("\n\t" + "\n\t")) {
            case "\n\t":
                v4 = (v4 - v4)
            case "a":
                if false {
                    v15 := GenS2{F0: false, F1: 2, F2: 42}
                    v4 = 132
                    _ = v15
                }
            }
            _ = fi14
        }
        _ = fi13
    }
    for fi16 := 0; fi16 < 7; fi16++ {
        if (42 > 2) {
            var v17 int8 = 127
            fmt.Println(fi16, v4)
            delete(v3, (fi16 << ((18446744073709551615) & 63)))
            _ = v17
        }
        _ = fi16
    }
    for fi18 := 0; fi18 < 3; fi18++ {
        if false {
            v19 := v3[fi18]
            _ = v19
        }
        _ = fi18
    }
    var v20 int8 = -1
    fmt.Println(v1)
    for fi21 := 0; fi21 < 5; fi21++ {
        v4 += 127
        fmt.Println(v20, fi21, fi21)
        v22 := GenS1{F0: (v1 | 154)}
        _ = fi21
        _ = v22
    }
    v3[128] = "a"
    var v23 int16 = genFn0(551, "foo bar", 1)
    fmt.Println(v4, v20, v23)
    for fi24 := 0; fi24 < 7; fi24++ {
        v25 := GenS1{F0: int64(v4)}
        var v26 string = "\n\t"
        v27 := v3[127]
        _ = fi24
        _ = v25
        _ = v26
        _ = v27
    }
    v23 = 1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
}
