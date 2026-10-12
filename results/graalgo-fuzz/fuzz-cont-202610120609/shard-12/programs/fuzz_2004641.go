package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 int32 }
func genFn0(p0 int) int64 {
    var v6 int64 = -9223372036854775808
    if false {
        v6 += v6
    } else {
        var v7 uint64 = 1
        if (!(p0 <= (p0 << ((42) & 63)))) {
            if false {
                var v8 int32 = -1
                _ = v8
            } else {
                var v9 int = (^(p0))
                if true {
                    p0 -= (v9 - (^(v9)))
                    switch (p0 | (v9 << ((1) & 63))) {
                    case 0:
                        for fi10 := 0; fi10 < 14; fi10++ {
                            v11 := GenS1{F0: 65535, F1: (("\n\t" + "a") + "foo bar"), F2: 219}
                            _ = fi10
                            _ = v11
                        }
                    }
                    var v12 int8 = -1
                    _ = v12
                }
                _ = v9
            }
            if false {
                v13 := GenS1{F0: 1, F1: ("\n\t" + ("hello" + ("hello" + ""))), F2: -229}
                var v14 string = "a"
                _ = v13
                _ = v14
            } else {
                fmt.Println(p0, v7)
                var v15 int8 = 0
                _ = v15
            }
            fmt.Println(v7, v7, p0)
        }
        _ = v7
    }
    _ = p0
    _ = v6
    return v6
}

func main() {
    var v1 int16 = -1
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]bool)
    var v4 int = -9223372036854775808
    v5 := make(map[uint32]bool)
    delete(v5, 702)
    fmt.Println(v1)
    switch -1 {
    case -1:
        var v16 int16 = (v1 + 1)
        _ = v16
    default:
        v17 := []int{v4, v4}
        _ = v17
    }
    v1 += int16(v4)
    for fi18 := 0; fi18 < 16; fi18++ {
        v19 := GenS1{F0: 477, F1: ("a" + "foo bar"), F2: -692}
        v4 -= (-(v4))
        _ = fi18
        _ = v19
    }
    v1 = 437
    var v20 uint16 = 0
    v4 -= (v4 | v4)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
