package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 int32; F2 uint32 }
type GenS3 struct { F0 int8; F1 uint }

func main() {
    v1 := make(map[int]int)
    var v2 bool = (0 <= -1)
    var v3 uint16 = 494
    var v4 int64 = 9223372036854775807
    fmt.Println(v4)
    var v5 string = "a"
    v6 := GenS1{F0: "\n\t"}
    v7 := GenS1{F0: (v5 + ("\n\t" + v5))}
    v8 := GenS2{F0: 43, F1: -2147483648, F2: 608}
    v9 := []int{}
    v2 = ((v4 - (v4 & v4)) >= (v4 & v4))
    v2 = ((v4 >> ((42) & 63)) >= v4)
    for fi10 := 0; fi10 < 8; fi10++ {
        v2 = v2
        switch -471 {
        case -39:
            for fi11 := 0; fi11 < 14; fi11++ {
                v3 = v3
                _ = fi11
            }
        }
        _ = fi10
    }
    for fi12 := 0; fi12 < 16; fi12++ {
        v13 := GenS1{F0: v5}
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 18; fi14++ {
        fmt.Println(v4, fi14, fi14)
        _ = fi14
    }
    v2 = (2 != 128)
    if (true && false) {
        switch 1 {
        case -1:
            for fi15 := 0; fi15 < 2; fi15++ {
                v5 = (((("a" + v5) + ("foo bar" + (("\n\t" + "\n\t") + v5))) + "a") + ("foo bar" + "\n\t"))
                v16 := GenS2{F0: 32767, F1: 0, F2: 967}
                _ = fi15
                _ = v16
            }
        default:
            v3 -= 1
        }
        v1[-9223372036854775808] = 1
        var v17 uint = 65535
        _ = v17
    }
    v5 = v5
    v4 = (v4 & -3)
    v1[-2147483648] = 0
    v18 := GenS2{F0: -32768, F1: -1, F2: 137}
    for fi19 := 0; fi19 < 18; fi19++ {
        var v20 int8 = 42
        for fi21 := 0; fi21 < 3; fi21++ {
            v20 = v20
            var v22 int64 = v4
            v3 += 65535
            fmt.Println(v5, v2)
            _ = fi21
            _ = v22
        }
        if v2 {
            if ((true && ((uint32(v3) <= 2) && true)) && (float64(0.1) != 0.1)) {
                for fi23 := 0; fi23 < 16; fi23++ {
                    fmt.Println(fi23)
                    _ = fi23
                }
                var v24 string = ""
                var v25 uint8 = 2
                _ = v24
                _ = v25
            } else {
                var v26 int32 = 2147483647
                v27 := GenS3{F0: 127, F1: 745}
                var v28 uint8 = 0
                _ = v26
                _ = v27
                _ = v28
            }
            v4 = v4
        } else {
            v29 := GenS2{F0: 102, F1: 1, F2: uint32(v3)}
            v5 = v5
            _ = v29
        }
        _ = fi19
        _ = v20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
