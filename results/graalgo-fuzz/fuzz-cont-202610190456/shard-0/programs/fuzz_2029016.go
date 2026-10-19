package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }

func main() {
    var v1 float64 = 1.5
    v2 := make(map[uint32]bool)
    v3 := make(map[string]bool)
    v4 := []int{-9223372036854775808, 127, 4294967295, 32768}
    delete(v2, 591)
    var v5 int16 = 424
    var v6 int = -1
    if true {
        for fi7 := 0; fi7 < 20; fi7++ {
            for fi8 := 0; fi8 < 19; fi8++ {
                if len(v4) > 0 {
                    v4[0] = 1
                }
                v5 = 1
                _ = fi8
            }
            for fi9 := 0; fi9 < 13; fi9++ {
                var v10 int8 = 42
                _ = fi9
                _ = v10
            }
            _ = fi7
        }
        for fi11 := 0; fi11 < 8; fi11++ {
            var v12 int64 = -1
            if false {
                var v13 int32 = -148
                _ = v13
            } else {
                var v14 int64 = 9223372036854775807
                _ = v14
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi11
            _ = v12
        }
        var v15 uint8 = 1
        _ = v15
    } else {
        var v16 uint64 = 1
        var v17 uint32 = 65535
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 15; fi18++ {
        _ = len(v4)
        _ = fi18
    }
    if (-128 > 0) {
        v19 := GenS1{F0: (^(v6)), F1: 2147483647}
        _ = v19
    }
    var v20 uint32 = 2
    if false {
        v1 -= v1
        var v21 uint16 = 42
        _ = v21
    } else {
        v22 := GenS1{F0: (v6 + (v6 % ((v6) | 1))), F1: (v20 % (((^(v20))) | 1))}
        fmt.Println(v5, v6)
        _ = v22
    }
    v1 -= v1
    var v23 string = ("foo bar" + "\n\t")
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
