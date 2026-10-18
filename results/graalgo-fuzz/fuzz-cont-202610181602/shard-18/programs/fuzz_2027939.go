package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
func genFn0(p0 string, p1 int8, p2 bool) int16 {
    var v4 float64 = -2.25
    if (!false) {
        for fi5 := 0; fi5 < 7; fi5++ {
            var v6 int64 = 127
            switch 994 {
            case 2147483647:
                for fi7 := 0; fi7 < 20; fi7++ {
                    fmt.Println(p0, v6)
                    p1 = (^(p1))
                    v6 = v6
                    if (p0 <= "\n\t") {
                        v8 := []int{fi7, fi7, (fi5 + (^(fi7))), (fi7 / ((fi7) | 1))}
                        var v9 float32 = (1e10 / (0.0 - -2.25))
                        _ = len(v8)
                        _ = v8
                        _ = v9
                    }
                    if (32 >= 42) {
                        var v10 uint = 67
                        if p2 {
                            v11 := []int{(^(fi7)), fi5, (^(fi5)), fi5}
                            v12 := GenS1{F0: 0, F1: p2}
                            var v13 int = 0
                            _ = v11
                            _ = v12
                            _ = v13
                        }
                        var v14 string = ((p0 + ((p0 + "") + p0)) + ((("" + "") + (p0 + ((p0 + p0) + "hello"))) + p0))
                        _ = v10
                        _ = v14
                    }
                    _ = fi7
                }
            default:
                p1 -= -128
            }
            _ = fi5
            _ = v6
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 0
}

func main() {
    var v1 int = -2147483648
    var v2 uint64 = 1
    v3 := make(map[string]int)
    switch "" {
    case "\n\t":
        var v15 uint64 = 42
        _ = v15
    case "":
        for fi16 := 0; fi16 < 19; fi16++ {
            v17 := GenS1{F0: -789, F1: false}
            switch (v1 * (v1 << ((4294967295) & 63))) {
            case 32767:
                if true {
                    fmt.Println(fi16, v1)
                }
            }
            _ = fi16
            _ = v17
        }
    case "foo bar":
        v18 := []int{(v1 + (^(v1))), 127, -2147483648, (^(v1))}
        _ = v18
    }
    var v19 uint8 = 1
    v1 -= (v1 >> ((407) & 63))
    if true {
        v19 = v19
    }
    v20 := []int{(v1 % (((v1 | -9223372036854775808)) | 1)), v1, (v1 ^ 9223372036854775807), -2147483648}
    switch "foo bar" {
    case "a":
        var v21 float64 = 0.1
        _ = v21
    case "hello":
        _ = len(v20)
    case "\n\t":
        var v22 float64 = 0.0
        _ = v22
    default:
        for fi23 := 0; fi23 < 1; fi23++ {
            for fi24 := 0; fi24 < 2; fi24++ {
                v2 += 4294967295
                var v25 uint8 = 2
                _ = fi24
                _ = v25
            }
            v20 = append(v20, v1)
            v1 -= (fi23 / ((v1) | 1))
            _ = fi23
        }
    }
    var v26 int = (v1 << ((600) & 63))
    v1 += v26
    var v27 uint = 36
    _ = len(v20)
    v28 := GenS1{F0: 833, F1: true}
    for fi29 := 0; fi29 < 20; fi29++ {
        v30 := GenS1{F0: -1, F1: (42 >= 1)}
        _ = fi29
        _ = v30
    }
    v20 = append(v20, (v1 + v1))
    var v31 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v31
}
