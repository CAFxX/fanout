package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
func genFn0(p0 uint64, p1 uint, p2 int) int64 {
    p1 += p1
    if true {
        fmt.Println(p1, p2, p1)
    } else {
        if (false || (-67 > -621)) {
            for fi5 := 0; fi5 < 5; fi5++ {
                switch fi5 {
                case -431:
                    for fi6 := 0; fi6 < 13; fi6++ {
                        p0 -= p0
                        p0 += p0
                        _ = fi6
                    }
                }
                _ = fi5
            }
            fmt.Println(p1)
            v7 := GenS1{F0: p1, F1: 0}
            _ = v7
        }
        var v8 float32 = 1.5
        v9 := GenS1{F0: p1, F1: 1}
        _ = v8
        _ = v9
    }
    v10 := []int{p2, 37, p2, p2}
    var v11 uint32 = 2147483648
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return 0
}

func main() {
    v1 := []int{}
    var v2 string = ""
    v3 := make(map[uint32]uint32)
    var v4 uint8 = 1
    v1 = append(v1, 330)
    if (!(!true)) {
        var v12 int = 4294967295
        var v13 float32 = (float32((-2.25 * 1.5)) * 0.0)
        if len(v1) > 0 {
            v1[0] = (v12 + (v12 | (v12 - 134)))
        }
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 14; fi14++ {
        var v15 int32 = 44
        for fi16 := 0; fi16 < 12; fi16++ {
            if false {
                var v17 uint8 = v4
                _ = v17
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v15 -= (v15 << ((18446744073709551615) & 63))
            _ = fi16
        }
        _ = fi14
        _ = v15
    }
    v4 = (v4 | (^(v4)))
    v2 = v2
    for fi18 := 0; fi18 < 3; fi18++ {
        switch "a" {
        case "hello":
            v2 = "hello"
        case "a":
            if true {
                var v19 uint8 = v4
                if true {
                    var v20 int32 = 1
                    _ = v20
                }
                _ = v19
            }
        }
        for fi21 := 0; fi21 < 3; fi21++ {
            if ((fi18 ^ fi18) < 42) {
                v22 := GenS1{F0: 65535, F1: -1}
                _ = v22
            } else {
                v23 := GenS1{F0: 4294967295, F1: -1}
                _ = v23
            }
            fmt.Println(fi21, v4)
            _ = fi21
        }
        fmt.Println(v4, fi18)
        _ = fi18
    }
    fmt.Println(v2)
    v4 = v4
    v1 = append(v1, 0)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
