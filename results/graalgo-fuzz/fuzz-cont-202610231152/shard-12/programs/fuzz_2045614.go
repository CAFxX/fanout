package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 string }
func genFn0(p0 bool, p1 uint, p2 uint16) string {
    if (!false) {
        fmt.Println(p1, p1)
        v6 := []int{-2147483648, 2147483647}
        if p0 {
            p1 += 751
        }
        _ = v6
    } else {
        switch -263 {
        case -2147483648:
            for fi7 := 0; fi7 < 4; fi7++ {
                if p0 {
                    v8 := []int{1, fi7, fi7, 127}
                    var v9 uint8 = 127
                    _ = v8
                    _ = v9
                }
                _ = fi7
            }
        }
    }
    v10 := GenS1{F0: 0, F1: 762}
    v11 := []int{}
    p2 = 65535
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return "hello"
}

func main() {
    v1 := []uint32{4294967295, 2, 1}
    v2 := make(map[string]uint32)
    v3 := []string{}
    var v4 string = (("\n\t" + "a") + "\n\t")
    var v5 float64 = 0.0
    var v12 float64 = (-2.25 / ((v5 * v5) + -2.25))
    for fi13 := 0; fi13 < 1; fi13++ {
        v12 -= 1e10
        switch "" {
        case "\n\t":
            if (0.1 > ((1.5 + 0.0) * 0.1)) {
                if false {
                    v4 = (v4 + ((v4 + "") + "hello"))
                } else {
                    if true {
                        v5 = (1e10 / (v12 * 0.0))
                    }
                }
            } else {
                _ = len(v3)
                _ = len(v3)
            }
        case "a":
            fmt.Println(v5, v12)
        case "":
            v12 -= v12
        }
        switch (((v4 + (v4 + "hello")) + v4) + v4) {
        case "":
            if ((0 >= 149) || false) {
                var v14 float64 = 1.5
                var v15 int16 = -32768
                _ = v14
                _ = v15
            }
        case "foo bar":
            v5 = v5
        }
        _ = fi13
    }
    var v16 bool = true
    v4 = (v4 + (v4 + (v4 + "hello")))
    v17 := GenS2{F0: v4}
    for fi18 := 0; fi18 < 9; fi18++ {
        var v19 int = 32767
        v20 := GenS2{F0: (v4 + v4)}
        _ = fi18
        _ = v19
        _ = v20
    }
    fmt.Println(v16)
    var v21 uint8 = 128
    var v22 int32 = -32
    for fi23 := 0; fi23 < 12; fi23++ {
        for fi24 := 0; fi24 < 8; fi24++ {
            if (18446744073709551615 > 1) {
                v4 = v4
                v25 := v2[v4]
                _ = len(v3)
                _ = v25
            }
            _ = fi24
        }
        if true {
            v21 += (v21 << ((65535) & 63))
            switch 2147483647 {
            case -2147483648:
                switch fi23 {
                case 42:
                    if len(v1) > 0 {
                        v1[0] = 941
                    }
                case 865:
                    var v26 bool = (((!(v16 && true)) && ((-(v22)) != int32(v21))) && (!true))
                    _ = v26
                case 4294967295:
                    var v27 uint = 203
                    _ = v27
                }
            case -9223372036854775808:
                var v28 uint64 = 4294967295
                _ = v28
            default:
                v21 = 127
            }
        } else {
            v12 -= v12
        }
        _ = fi23
    }
    var v29 float64 = v5
    var v30 uint8 = v21
    v21 = v21
    if v16 {
        var v31 int = 2147483647
        _ = v31
    }
    fmt.Println(v22)
    v5 = v29
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
