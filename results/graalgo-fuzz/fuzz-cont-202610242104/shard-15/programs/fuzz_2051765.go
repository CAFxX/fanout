package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 uint64 }
func genFn0(p0 bool, p1 bool) int {
    for fi5 := 0; fi5 < 16; fi5++ {
        if false {
            var v6 int64 = 0
            switch (fi5 ^ (fi5 + 0)) {
            case 728:
                for fi7 := 0; fi7 < 14; fi7++ {
                    switch (fi5 * 425) {
                    case 127:
                        var v8 int32 = 42
                        _ = v8
                    case 32768:
                        if (p0 || p0) {
                            var v9 uint64 = 1
                            v6 = v6
                            _ = v9
                        } else {
                            fmt.Println(fi7)
                            var v10 string = "a"
                            v11 := []int{32768, fi5, fi7, 0}
                            _ = v10
                            _ = v11
                        }
                    }
                    var v12 uint8 = 1
                    _ = fi7
                    _ = v12
                }
            }
            _ = v6
        }
        p0 = (42 != 100)
        if (92 == 1) {
            v13 := []int{(fi5 | 4294967295)}
            v14 := GenS1{F0: (fi5 - (fi5 << ((1) & 63))), F1: 42, F2: 5}
            _ = v13
            _ = v14
        } else {
            var v15 float32 = (0.0 * -2.25)
            fmt.Println(p1, p0)
            for fi16 := 0; fi16 < 20; fi16++ {
                v15 += v15
                _ = fi16
            }
            _ = v15
        }
        _ = fi5
    }
    var v17 uint16 = 65535
    p1 = true
    for fi18 := 0; fi18 < 18; fi18++ {
        v19 := []int{}
        _ = fi18
        _ = v19
    }
    v17 = 325
    _ = p0
    _ = p1
    _ = v17
    return 538
}

func main() {
    v1 := []uint32{4294967295, 794, 687}
    var v2 int = 32767
    v3 := make(map[uint32]string)
    v4 := []int{v2, v2, v2}
    if false {
        switch "\n\t" {
        case "a":
            if (v2 > (v2 * (v2 - 1))) {
                var v20 int = 122
                _ = v20
            } else {
                v2 = (v2 & genFn0(true, true))
                var v21 int8 = 100
                _ = v21
            }
        case "foo bar":
            for fi22 := 0; fi22 < 4; fi22++ {
                v2 += (^(fi22))
                _ = fi22
            }
        case "hello":
            var v23 uint32 = 1
            _ = v23
        default:
            for fi24 := 0; fi24 < 2; fi24++ {
                var v25 int16 = 1
                fmt.Println(v25, v2)
                var v26 int16 = (-(v25))
                _ = fi24
                _ = v25
                _ = v26
            }
        }
    } else {
        var v27 bool = (-19 != 42)
        _ = v27
    }
    var v28 uint32 = 42
    if ((60 <= 1) || (18446744073709551615 <= 959)) {
        v2 = (v2 & (v2 >> ((972) & 63)))
        fmt.Println(v28, v28)
    }
    fmt.Println(v28, v2)
    var v29 float64 = float64(((0.0 + 3.14159) + ((0.1 / 0.1) * ((float32(-2.25) * 1.5) + 0.1))))
    switch "foo bar" {
    case "\n\t":
        for fi30 := 0; fi30 < 18; fi30++ {
            var v31 uint8 = 220
            _ = fi30
            _ = v31
        }
    default:
        for fi32 := 0; fi32 < 13; fi32++ {
            switch (v2 << ((65535) & 63)) {
            case -2147483648:
                fmt.Println(v29)
            case 32767:
                v28 -= (v28 / ((v28) | 1))
            default:
                fmt.Println(v28, v2)
            }
            _ = fi32
        }
    }
    if (((v29 / (v29 / (-2.25 / v29))) - v29) == 0.1) {
        fmt.Println(v28, v28, v28)
        var v33 uint32 = (v28 - (v28 | 0))
        _ = v33
    }
    switch "\n\t" {
    case "foo bar":
        fmt.Println(v28, v28, v29)
    default:
        var v34 uint = 18446744073709551615
        _ = v34
    }
    var v35 int16 = 32767
    v36 := GenS1{F0: (v2 >> ((313) & 63)), F1: v35, F2: 23}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v35
    _ = v36
    _ = v4
}
