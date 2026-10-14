package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 uint64; F1 uint; F2 int }
func genFn0(p0 uint, p1 uint8) int8 {
    if (((3.14159 - -2.25) * 3.14159) <= (1e10 + 1.5)) {
        switch "" {
        case "hello":
            switch "a" {
            case "foo bar":
                var v7 uint = p0
                _ = v7
            case "hello":
                var v8 int8 = -72
                _ = v8
            }
        case "a":
            v9 := []int{42, -664}
            _ = v9
        case "":
            v10 := []int{0, 32767, 2147483647, 32768}
            _ = v10
        default:
            v11 := []int{-9223372036854775808, 32767, 2147483647, 0}
            _ = v11
        }
        for fi12 := 0; fi12 < 17; fi12++ {
            fmt.Println(fi12, p0)
            _ = fi12
        }
    } else {
        v13 := []int{0, 32768}
        p1 = (^(p1))
        switch 128 {
        case 32768:
            var v14 int16 = 32767
            _ = v14
        case -113:
            var v15 uint8 = 0
            _ = v15
        default:
            for fi16 := 0; fi16 < 2; fi16++ {
                v17 := GenS1{F0: 19, F1: (^(p1))}
                for fi18 := 0; fi18 < 11; fi18++ {
                    var v19 bool = false
                    p1 -= p1
                    var v20 float32 = (0.0 - (0.0 * 1.5))
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                _ = fi16
                _ = v17
            }
        }
        _ = v13
    }
    v21 := []int{137, 42}
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = v21
    return -67
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int = 9223372036854775807
    var v3 bool = (((0.0 * 0.1) / ((1.5 * 1.5) / (0.0 + 1.5))) >= (((0.1 - 0.1) + 3.14159) * (float64(-2.25) / 3.14159)))
    v4 := []int{-9223372036854775808, v2, (v2 | 128)}
    var v5 float32 = 0.1
    var v6 uint16 = 0
    v6 = v6
    for fi22 := 0; fi22 < 12; fi22++ {
        v6 = (v6 / ((32767) | 1))
        var v23 int64 = 205
        v23 += 1
        _ = fi22
        _ = v23
    }
    if ("" >= "hello") {
        v24 := GenS1{F0: (v2 + (v2 % ((9223372036854775807) | 1))), F1: 128}
        if v3 {
            var v25 int8 = 42
            fmt.Println(v2, v5, v25)
            _ = v25
        } else {
            v26 := GenS1{F0: (v2 - -673), F1: 2}
            var v27 int = 32767
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = v26
            _ = v27
        }
        fmt.Println(v6, v3, v3)
        _ = v24
    }
    v28 := GenS1{F0: v2, F1: uint8(v6)}
    var v29 bool = v3
    switch "hello" {
    case "\n\t":
        var v30 uint32 = 449
        _ = v30
    default:
        v5 += (v5 / v5)
    }
    fmt.Println(v2, v5)
    var v31 float64 = 1.5
    for fi32 := 0; fi32 < 15; fi32++ {
        switch v2 {
        case -12:
            var v33 uint64 = 0
            _ = v33
        default:
            fmt.Println(v3, v5, v2)
        }
        _ = fi32
    }
    fmt.Println(v5, v3, v6)
    switch (v2 & (v2 << ((1) & 63))) {
    case -141:
        var v34 string = ""
        _ = v34
    case 32767:
        switch (-(v2)) {
        case -615:
            v29 = v3
        case 613:
            delete(v1, v2)
        }
    default:
        v1[v2] = 773
    }
    var v35 uint16 = (v6 * 0)
    fmt.Println(v29)
    if (true && true) {
        v31 -= v31
        var v36 int32 = 784
        v29 = v29
        _ = v36
    } else {
        var v37 uint32 = 4294967295
        _ = v37
    }
    v38 := GenS1{F0: (^(v2)), F1: 200}
    v3 = false
    v39 := GenS2{F0: 1, F1: 486, F2: v2}
    if len(v4) > 0 {
        v4[0] = v2
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v35
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
