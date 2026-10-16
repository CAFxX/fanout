package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 uint32 }
func genFn0(p0 uint32, p1 int16, p2 int) uint64 {
    if ((^(p2)) > 127) {
        v6 := GenS1{F0: -707}
        _ = v6
    } else {
        p0 = p0
    }
    if false {
        for fi7 := 0; fi7 < 6; fi7++ {
            fmt.Println(p1)
            _ = fi7
        }
        for fi8 := 0; fi8 < 14; fi8++ {
            switch (p2 * p2) {
            case 115:
                if ("\n\t" >= "hello") {
                    switch (^(fi8)) {
                    case -253:
                        p1 += -752
                    }
                    fmt.Println(p0, p0)
                    p0 = p0
                } else {
                    for fi9 := 0; fi9 < 10; fi9++ {
                        v10 := GenS1{F0: 40}
                        p0 += 2147483648
                        v11 := GenS2{F0: 65535, F1: (p0 + p0)}
                        _ = fi9
                        _ = v10
                        _ = v11
                    }
                    fmt.Println(p2)
                }
            case 9223372036854775807:
                p1 = p1
            }
            switch "" {
            case "hello":
                var v12 int16 = -928
                _ = v12
            case "":
                if false {
                    switch (^(fi8)) {
                    case 361:
                        var v13 float32 = -2.25
                        _ = v13
                    case 1:
                        var v14 float64 = (3.14159 * float64(float32(0.0)))
                        _ = v14
                    }
                    switch fi8 {
                    case 116:
                        v15 := []int{(^(p2)), 117}
                        _ = v15
                    default:
                        var v16 float64 = 1e10
                        _ = v16
                    }
                    v17 := []int{}
                    _ = v17
                }
            case "a":
                v18 := GenS2{F0: 0, F1: p0}
                _ = v18
            }
            _ = fi8
        }
        fmt.Println(p2)
    } else {
        p1 = 42
        var v19 string = "hello"
        _ = v19
    }
    var v20 uint = uint(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    return 18446744073709551615
}

func main() {
    v1 := make(map[int]int)
    var v2 float32 = -2.25
    var v3 int64 = 9223372036854775807
    var v4 bool = (!(-145 <= 32767))
    var v5 int64 = v3
    var v21 float64 = 1e10
    var v22 uint64 = 805
    v5 -= v3
    v22 = 18446744073709551615
    fmt.Println(v22, v2)
    for fi23 := 0; fi23 < 10; fi23++ {
        switch ((("foo bar" + "foo bar") + "a") + "") {
        case "":
            fmt.Println(v2, fi23, v4)
        case "\n\t":
            fmt.Println(v3, v4, v4)
        case "hello":
            var v24 uint32 = 650
            _ = v24
        }
        switch fi23 {
        case 4294967295:
            v21 = v21
        case 63:
            if (65535 >= 2) {
                if true {
                    for fi25 := 0; fi25 < 18; fi25++ {
                        _ = len(v1)
                        _ = fi25
                    }
                    v26 := GenS2{F0: 42, F1: 29}
                    switch 4294967295 {
                    case 4294967295:
                        v21 -= v21
                    case 42:
                        v27 := []int{fi23, (fi23 % (((fi23 << ((4294967295) & 63))) | 1)), -480, (fi23 ^ (fi23 * fi23))}
                        _ = v27
                    case 2147483647:
                        if (v22 < v22) {
                            v28 := []int{128, fi23, fi23, 578}
                            v4 = true
                            v4 = (float64(v2) >= 0.1)
                            _ = v28
                        }
                    default:
                        if (!(!(!((v2 + 3.14159) != v2)))) {
                            var v29 float64 = v21
                            if (v4 && v4) {
                                _ = len(v1)
                                v29 = 0.0
                                fmt.Println(v2, v22, v21)
                            }
                            _ = v29
                        }
                    }
                    _ = v26
                } else {
                    var v30 uint64 = 18446744073709551615
                    _ = v30
                }
            }
        case 11:
            var v31 int64 = v5
            _ = v31
        }
        _ = fi23
    }
    fmt.Println(v5)
    v22 = (v22 & (v22 - 1))
    v32 := GenS2{F0: 65535, F1: 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
