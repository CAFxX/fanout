package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32; F1 int64; F2 int16 }
func genFn0(p0 uint64, p1 int64) int {
    var v6 float32 = 0.1
    switch -1 {
    case 560:
        p0 = 54
    }
    var v7 float64 = (3.14159 + (3.14159 * 0.0))
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 9223372036854775807
}
func genFn1(p0 int32, p1 bool, p2 int32) uint32 {
    for fi8 := 0; fi8 < 19; fi8++ {
        p2 += (p0 + p0)
        _ = fi8
    }
    var v9 string = "foo bar"
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 1
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint16 = 422
    var v3 uint32 = 2147483648
    var v4 uint64 = 610
    var v5 int64 = 0
    var v10 uint64 = v4
    if false {
        v3 -= v3
        fmt.Println(v4)
    } else {
        var v11 uint8 = 128
        _ = v11
    }
    v1["\n\t"] = false
    v12 := v1["\n\t"]
    switch "foo bar" {
    case "\n\t":
        v1["foo bar"] = ((127 == -1) || true)
    case "hello":
        fmt.Println(v2)
    }
    var v13 float32 = -2.25
    var v14 uint16 = (v2 + 157)
    v12 = true
    if (v12 && (((v10 / ((1) | 1)) == (v4 | v4)) && (-116 > 119))) {
        var v15 uint8 = 127
        switch 0 {
        case -759:
            if (!v12) {
                for fi16 := 0; fi16 < 7; fi16++ {
                    v17 := GenS1{F0: (uint(v4) >= uint(v2))}
                    _ = fi16
                    _ = v17
                }
                var v18 float64 = 1.5
                v15 = (^(v15))
                _ = v18
            } else {
                if v12 {
                    v19 := []int{422, 4294967295, 32767}
                    if ((v14 + 42) <= (v14 + 32767)) {
                        for fi20 := 0; fi20 < 5; fi20++ {
                            var v21 float32 = v13
                            if len(v19) >= 2 {
                                v19 = v19[1:len(v19)]
                            }
                            _ = fi20
                            _ = v21
                        }
                    } else {
                        v22 := GenS2{F0: 0, F1: v5, F2: -1}
                        v13 += (1.5 - -2.25)
                        switch genFn0(uint64(v2), (v5 | v5)) {
                        case 69:
                            fmt.Println(v13)
                        }
                        _ = v22
                    }
                    _ = v19
                }
                if (v12 && v12) {
                    v23 := GenS1{F0: v12}
                    _ = v23
                } else {
                    v24 := v1["a"]
                    _ = v24
                }
            }
        case 9223372036854775807:
            switch ("" + "foo bar") {
            case "":
                v25 := GenS1{F0: (9223372036854775807 != v5)}
                _ = v25
            case "foo bar":
                v26 := GenS2{F0: 1, F1: v5, F2: 0}
                _ = v26
            case "\n\t":
                var v27 string = "foo bar"
                _ = v27
            }
        case -1:
            var v28 uint8 = 189
            _ = v28
        }
        _ = v15
    }
    if (v5 > -211) {
        v2 = v14
    }
    v29 := []int{}
    var v30 float32 = v13
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
