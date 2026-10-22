package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string; F1 int64; F2 uint }
type GenS3 struct { F0 bool }
func genFn0(p0 int, p1 int32, p2 uint8) int64 {
    var v4 float64 = 0.1
    v5 := []int{p0}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    switch "\n\t" {
    case "\n\t":
        var v6 int32 = -105
        _ = v6
    case "hello":
        for fi7 := 0; fi7 < 18; fi7++ {
            if true {
                var v8 uint = 938
                _ = v8
            }
            _ = fi7
        }
    case "a":
        switch (p0 << ((65535) & 63)) {
        case 2147483647:
            for fi9 := 0; fi9 < 5; fi9++ {
                var v10 uint64 = 819
                _ = fi9
                _ = v10
            }
        default:
            switch (p0 - (^(p0))) {
            case 32768:
                if (-551 < 32767) {
                    p2 = 1
                    var v11 bool = true
                    _ = v11
                }
            case 2147483647:
                for fi12 := 0; fi12 < 12; fi12++ {
                    var v13 float64 = -2.25
                    _ = fi12
                    _ = v13
                }
            default:
                p2 -= p2
            }
        }
    }
    var v14 string = "a"
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v4
    _ = v5
    return 0
}
func genFn1(p0 uint64, p1 uint32, p2 uint16) int {
    var v15 uint8 = 5
    fmt.Println(p0, v15)
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return 127
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []uint32{0, 2147483647, 900}
    var v3 float32 = float32((0.0 / 1e10))
    v3 = ((((1.5 - v3) * 3.14159) * v3) / v3)
    for fi16 := 0; fi16 < 17; fi16++ {
        var v17 int16 = -1
        _ = fi16
        _ = v17
    }
    v18 := GenS1{F0: 42}
    if false {
        v19 := GenS1{F0: -193}
        var v20 uint = 1
        switch 128 {
        case 668:
            var v21 string = ("foo bar" + "a")
            _ = v21
        case 775:
            v22 := v1[-779]
            _ = v22
        case 437:
            switch "a" {
            case "hello":
                v23 := GenS2{F0: ("a" + ("" + "hello")), F1: 42, F2: 2}
                _ = v23
            case "":
                for fi24 := 0; fi24 < 20; fi24++ {
                    for fi25 := 0; fi25 < 8; fi25++ {
                        v1[fi25] = 29
                        _ = fi25
                    }
                    switch (fi24 % (((fi24 - (fi24 / ((fi24) | 1)))) | 1)) {
                    case 9223372036854775807:
                        var v26 float32 = v3
                        _ = v26
                    }
                    _ = fi24
                }
            case "foo bar":
                var v27 float32 = v3
                _ = v27
            default:
                v1[209] = 559
            }
        }
        _ = v19
        _ = v20
    }
    fmt.Println(v3)
    v3 -= v3
    switch -1 {
    case -119:
        for fi28 := 0; fi28 < 1; fi28++ {
            if true {
                v3 += 0.0
                v3 -= (1.5 / 1e10)
                if len(v2) > 0 {
                    v2[0] = 2147483648
                }
            } else {
                fmt.Println(fi28, fi28)
                _ = len(v1)
                v3 = (v3 + (v3 * v3))
            }
            _ = fi28
        }
    case 32767:
        v29 := GenS2{F0: "foo bar", F1: 9223372036854775807, F2: 18446744073709551615}
        _ = v29
    }
    for fi30 := 0; fi30 < 3; fi30++ {
        v1[(fi30 | fi30)] = 65535
        if (2 >= 500) {
            v3 = ((1.5 / 3.14159) * ((v3 * v3) - ((v3 * 3.14159) - v3)))
        }
        _ = fi30
    }
    v31 := GenS3{F0: (9223372036854775807 <= 32768)}
    fmt.Println(v3)
    for fi32 := 0; fi32 < 20; fi32++ {
        var v33 uint16 = 908
        switch fi32 {
        case 32767:
            fmt.Println(v33, fi32, v33)
        default:
            fmt.Println(v3, v33)
        }
        var v34 uint = 65535
        _ = fi32
        _ = v33
        _ = v34
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v31
}
