package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 uint; F2 int32 }
func genFn0(p0 uint64, p1 bool) uint16 {
    switch -104 {
    case 128:
        v7 := GenS1{F0: 0}
        _ = v7
    }
    p1 = false
    p1 = p1
    for fi8 := 0; fi8 < 8; fi8++ {
        switch (("hello" + "\n\t") + "a") {
        case "foo bar":
            v9 := []int{fi8, (^(fi8)), fi8}
            _ = v9
        }
        switch (fi8 | fi8) {
        case 32768:
            fmt.Println(p1, p1)
        default:
            v10 := []int{-1}
            _ = v10
        }
        fmt.Println(p0, p0)
        _ = fi8
    }
    _ = p0
    _ = p1
    return 207
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]uint32)
    v3 := []int{315, 333}
    v4 := []int{128, 9223372036854775807, 9223372036854775807}
    var v5 uint8 = 122
    v6 := make(map[int]uint32)
    v11 := GenS1{F0: 306}
    var v12 string = "foo bar"
    if true {
        v12 = (v12 + v12)
        v12 = "hello"
        v13 := GenS1{F0: -467}
        _ = v13
    } else {
        var v14 uint64 = 992
        _ = v14
    }
    for fi15 := 0; fi15 < 14; fi15++ {
        if true {
            fmt.Println(v12, fi15)
        } else {
            if (!true) {
                if (1.5 >= 3.14159) {
                    v16 := GenS1{F0: -32768}
                    v5 = 0
                    _ = v16
                } else {
                    var v17 string = "a"
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    for fi18 := 0; fi18 < 2; fi18++ {
                        var v19 int8 = 85
                        _ = fi18
                        _ = v19
                    }
                    _ = v17
                }
                fmt.Println(v12, v5, fi15)
            } else {
                var v20 uint16 = 994
                _ = v20
            }
            if true {
                for fi21 := 0; fi21 < 2; fi21++ {
                    var v22 int32 = 175
                    _ = fi21
                    _ = v22
                }
                switch fi15 {
                case 300:
                    v2[fi15] = 24
                case 127:
                    v23 := GenS1{F0: -629}
                    _ = v23
                default:
                    v24 := GenS1{F0: -32768}
                    _ = v24
                }
            }
            var v25 uint16 = 899
            _ = v25
        }
        for fi26 := 0; fi26 < 7; fi26++ {
            v5 += (v5 << ((1) & 63))
            _ = fi26
        }
        _ = fi15
    }
    var v27 uint16 = 42
    var v28 int16 = 1000
    if (float64((-2.25 + -2.25)) <= (-2.25 / ((0.0 - float64((0.1 - 0.1))) - ((((float64((float32(1e10) * 1.5)) + 1.5) * 0.1) / 1e10) * 3.14159)))) {
        if true {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            var v29 uint = 0
            _ = v29
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v28 += v28
    }
    delete(v1, 42)
    v30 := GenS2{F0: v5, F1: 65535, F2: 42}
    var v31 uint8 = (v5 / ((v5) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
