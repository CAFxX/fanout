package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 uint32 }
func genFn0(p0 bool) int32 {
    for fi6 := 0; fi6 < 11; fi6++ {
        p0 = (p0 || false)
        var v7 uint = 1
        for fi8 := 0; fi8 < 5; fi8++ {
            v9 := GenS1{F0: v7, F1: 0, F2: 0}
            var v10 bool = (128 >= 2)
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    switch "hello" {
    case "":
        fmt.Println(p0)
    case "hello":
        fmt.Println(p0)
    case "foo bar":
        p0 = true
    }
    _ = p0
    return -217
}
func genFn1(p0 bool) int8 {
    v11 := GenS1{F0: 4294967295, F1: 255, F2: 598}
    fmt.Println(p0)
    fmt.Println(p0)
    p0 = p0
    for fi12 := 0; fi12 < 19; fi12++ {
        var v13 string = "foo bar"
        switch (fi12 - 184) {
        case 0:
            var v14 uint = 2
            _ = v14
        case -2147483648:
            if ((2147483647 != 279) || (!false)) {
                v13 = v13
                switch "a" {
                case "foo bar":
                    if (!(65535 == 897)) {
                        v15 := GenS1{F0: 941, F1: 0, F2: 560}
                        switch fi12 {
                        case 32768:
                            var v16 uint8 = 240
                            _ = v16
                        case 2147483647:
                            if (false && (float64((((1e10 + ((-2.25 + (float32((0.1 + 0.1)) + -2.25)) * ((3.14159 + 0.1) / (1e10 * 0.1)))) + (3.14159 * 3.14159)) - -2.25)) >= 1e10)) {
                                var v17 bool = p0
                                var v18 uint64 = 18446744073709551615
                                _ = v17
                                _ = v18
                            }
                        }
                        _ = v15
                    }
                case "a":
                    fmt.Println(fi12)
                default:
                    v19 := GenS1{F0: 0, F1: 255, F2: 65535}
                    _ = v19
                }
            }
        case -1:
            var v20 uint = 0
            _ = v20
        default:
            v21 := []int{}
            _ = v21
        }
        var v22 float32 = 0.0
        _ = fi12
        _ = v13
        _ = v22
    }
    _ = p0
    _ = v11
    return 71
}

func main() {
    var v1 int = 673
    var v2 int8 = 100
    var v3 uint64 = 2
    v4 := make(map[string]int)
    var v5 float64 = (float64(0.1) - 0.1)
    switch 128 {
    case 127:
        delete(v4, ("\n\t" + "hello"))
    case 1:
        v5 -= (v5 + v5)
    default:
        var v23 string = "\n\t"
        _ = v23
    }
    switch 32768 {
    case 1:
        for fi24 := 0; fi24 < 15; fi24++ {
            fmt.Println(v1)
            var v25 uint16 = 65535
            var v26 uint16 = (^(v25))
            _ = fi24
            _ = v25
            _ = v26
        }
    default:
        v27 := GenS1{F0: 18446744073709551615, F1: 1, F2: 42}
        _ = v27
    }
    fmt.Println(v3, v3, v1)
    v3 = v3
    switch ("hello" + "\n\t") {
    case "hello":
        v3 = (v3 * (v3 ^ v3))
    case "a":
        var v28 float64 = (1.5 - float64(-2.25))
        _ = v28
    case "\n\t":
        var v29 uint16 = 118
        _ = v29
    default:
        var v30 float64 = v5
        _ = v30
    }
    v1 -= (v1 / ((-224) | 1))
    switch (-(v1)) {
    case -559:
        v31 := []int{(v1 >> ((532) & 63))}
        _ = v31
    case 64:
        fmt.Println(v3)
    }
    v32 := []int{}
    var v33 string = "\n\t"
    delete(v4, (v33 + ("" + v33)))
    v33 = v33
    if ((!true) || true) {
        v1 = -839
        if true {
            v34 := GenS1{F0: 2, F1: 1, F2: 2147483648}
            var v35 uint32 = 2147483647
            _ = v34
            _ = v35
        } else {
            for fi36 := 0; fi36 < 4; fi36++ {
                var v37 string = v33
                var v38 uint64 = v3
                _ = fi36
                _ = v37
                _ = v38
            }
            _ = len(v32)
        }
    } else {
        v1 = -9223372036854775808
    }
    v39 := GenS1{F0: 18446744073709551615, F1: 244, F2: 725}
    for fi40 := 0; fi40 < 6; fi40++ {
        v1 = -2147483648
        if len(v32) > 0 {
            v32[0] = (v1 / (((v1 | (v1 ^ (v1 * 42)))) | 1))
        }
        v41 := GenS1{F0: 0, F1: 127, F2: 387}
        _ = fi40
        _ = v41
    }
    switch v1 {
    case 32768:
        fmt.Println(v1, v1, v5)
    default:
        v42 := GenS1{F0: 4294967295, F1: 2, F2: 2}
        _ = v42
    }
    fmt.Println(v33, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v39
    _ = v4
    _ = v5
}
