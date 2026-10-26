package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 int32 }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[string]int)
    var v3 int = 39
    v4 := GenS1{F0: (v3 * (v3 << ((286) & 63))), F1: 127, F2: 2147483647}
    for fi5 := 0; fi5 < 11; fi5++ {
        switch fi5 {
        case -2147483648:
            v3 -= (v3 ^ (v3 + (fi5 - -288)))
        case 756:
            fmt.Println(fi5, fi5)
        case 42:
            var v6 bool = (((0.0 * ((1e10 * 1e10) - float64(-2.25))) * 1e10) < (float64((float32((3.14159 * 0.1)) + float32(-2.25))) * (1.5 * -2.25)))
            _ = v6
        }
        for fi7 := 0; fi7 < 19; fi7++ {
            var v8 uint64 = 18446744073709551615
            _ = fi7
            _ = v8
        }
        if (36 > 128) {
            if (true || true) {
                var v9 bool = (929 <= -985)
                for fi10 := 0; fi10 < 11; fi10++ {
                    v9 = (-2.25 != float32(3.14159))
                    var v11 int8 = 42
                    _ = fi10
                    _ = v11
                }
                _ = v9
            } else {
                v3 = (v3 % (((fi5 >> ((2) & 63))) | 1))
            }
        } else {
            if (false || (1e10 != 1e10)) {
                v3 = 0
                v12 := GenS1{F0: v3, F1: -38, F2: 884}
                for fi13 := 0; fi13 < 7; fi13++ {
                    switch "a" {
                    case "":
                        switch "foo bar" {
                        case "\n\t":
                            if (42 != 16) {
                                v14 := []int{(v3 - (fi5 ^ fi13)), (-(v3)), v3, -1}
                                if true {
                                    fmt.Println(fi13, fi13)
                                }
                                _ = v14
                            }
                        case "a":
                            if (1 == 100) {
                                v1[541] = 496
                                delete(v1, 1)
                            }
                        case "hello":
                            var v15 int16 = 1
                            _ = v15
                        }
                    case "a":
                        var v16 int64 = -700
                        _ = v16
                    case "hello":
                        var v17 uint16 = 32767
                        _ = v17
                    }
                    if (true || (4294967295 == 86)) {
                        var v18 uint = 255
                        _ = v18
                    }
                    _ = fi13
                }
                _ = v12
            } else {
                switch v3 {
                case -2147483648:
                    if (21 <= 255) {
                        v19 := GenS1{F0: v3, F1: 1, F2: 42}
                        v20 := GenS1{F0: 128, F1: -16, F2: 0}
                        var v21 int8 = 42
                        _ = v19
                        _ = v20
                        _ = v21
                    }
                }
                v3 += v3
            }
        }
        _ = fi5
    }
    if false {
        if (true && (((1.5 - (0.0 * (3.14159 - 1e10))) / (1e10 + 0.1)) == -2.25)) {
            var v22 uint = 18446744073709551615
            _ = v22
        } else {
            fmt.Println(v3)
        }
        v3 -= 36
    } else {
        v3 = v3
        for fi23 := 0; fi23 < 7; fi23++ {
            v3 = (fi23 & fi23)
            var v24 int32 = -879
            _ = fi23
            _ = v24
        }
    }
    var v25 int = (v3 + v3)
    var v26 uint32 = 2147483648
    var v27 int8 = 1
    for fi28 := 0; fi28 < 6; fi28++ {
        v29 := GenS1{F0: (v25 & (v3 ^ fi28)), F1: v27, F2: 2147483647}
        fmt.Println(v27, v27)
        var v30 string = "\n\t"
        _ = fi28
        _ = v29
        _ = v30
    }
    switch 128 {
    case 4294967295:
        fmt.Println(v27, v3)
    case -289:
        var v31 uint = 255
        _ = v31
    default:
        fmt.Println(v27, v25)
    }
    if (32767 >= 804) {
        v27 = (-(v27))
    } else {
        var v32 int64 = 399
        if (1 == 42) {
            v33 := GenS1{F0: int(v27), F1: v27, F2: 42}
            fmt.Println(v32, v26, v25)
            fmt.Println(v32)
            _ = v33
        } else {
            var v34 int16 = 32767
            _ = v34
        }
        _ = v32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
