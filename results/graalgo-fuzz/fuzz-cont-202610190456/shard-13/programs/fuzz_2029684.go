package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string }

func main() {
    var v1 int16 = -32768
    v2 := []int{-2147483648, 32768, 591, 797}
    var v3 uint64 = 42
    var v4 int64 = 0
    var v5 uint64 = v3
    v6 := make(map[int]uint32)
    var v7 string = ("foo bar" + "hello")
    for fi8 := 0; fi8 < 17; fi8++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi8
    }
    _ = len(v6)
    var v9 int64 = (v4 | v4)
    if false {
        if false {
            v7 = v7
            v9 = (v4 ^ (v4 >> ((281) & 63)))
            var v10 uint16 = 65535
            _ = v10
        }
        if (-128 != 0) {
            v6[127] = 2147483647
        } else {
            var v11 uint = 0
            _ = v11
        }
        var v12 int8 = 1
        _ = v12
    } else {
        if true {
            if false {
                v13 := GenS1{F0: 0}
                var v14 bool = (int(v4) < 0)
                var v15 int32 = 0
                _ = v13
                _ = v14
                _ = v15
            }
        }
    }
    switch 1 {
    case -643:
        for fi16 := 0; fi16 < 10; fi16++ {
            v17 := GenS1{F0: 127}
            for fi18 := 0; fi18 < 5; fi18++ {
                if ((-1 > 127) || ((0.1 + ((3.14159 - float32(float64(-2.25))) + (1e10 * ((((0.1 * 3.14159) * 3.14159) + ((3.14159 * -2.25) / -2.25)) / 1.0)))) != (((((-2.25 / 0.1) * (1.5 / 1e10)) / (-2.25 + 0.1)) / -2.25) * ((1e10 + 1e10) * (0.0 - 1.5))))) {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                } else {
                    v9 += v9
                }
                v19 := GenS1{F0: 127}
                _ = fi18
                _ = v19
            }
            v4 -= v4
            _ = fi16
            _ = v17
        }
    }
    var v20 string = ""
    v21 := GenS2{F0: v7}
    var v22 int8 = 106
    v23 := GenS2{F0: ("a" + v20)}
    if true {
        fmt.Println(v20)
        for fi24 := 0; fi24 < 2; fi24++ {
            fmt.Println(v4, v5)
            v4 += v9
            fmt.Println(v22, v3)
            _ = fi24
        }
        fmt.Println(v4, v20)
    }
    _ = len(v2)
    for fi25 := 0; fi25 < 15; fi25++ {
        var v26 uint = 457
        _ = fi25
        _ = v26
    }
    var v27 int16 = (v1 ^ -982)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
