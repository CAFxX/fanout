package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 bool }
func genFn0(p0 uint16) uint {
    if true {
        v6 := []int{867, 1, -885}
        _ = v6
    }
    for fi7 := 0; fi7 < 19; fi7++ {
        switch "hello" {
        case "":
            for fi8 := 0; fi8 < 18; fi8++ {
                p0 = p0
                var v9 float64 = 1.5
                _ = fi8
                _ = v9
            }
        case "a":
            switch ("" + "hello") {
            case "\n\t":
                fmt.Println(fi7, fi7)
            default:
                v10 := []int{}
                _ = v10
            }
        }
        _ = fi7
    }
    v11 := GenS1{F0: -2147483648, F1: 65535}
    fmt.Println(p0)
    _ = p0
    _ = v11
    return 4294967295
}

func main() {
    var v1 uint32 = 0
    var v2 bool = true
    var v3 int32 = 0
    var v4 int64 = 9223372036854775807
    var v5 uint32 = 65535
    v2 = v2
    fmt.Println(v2)
    switch "a" {
    case "foo bar":
        var v12 int = 32767
        _ = v12
    }
    v13 := []int{9223372036854775807}
    var v14 int16 = 1
    switch "a" {
    case "a":
        var v15 bool = (true || (-2147483648 == 128))
        _ = v15
    default:
        var v16 float64 = 0.0
        _ = v16
    }
    var v17 int = -190
    for fi18 := 0; fi18 < 14; fi18++ {
        for fi19 := 0; fi19 < 2; fi19++ {
            if (464 < 0) {
                if len(v13) > 0 {
                    v13[0] = int(v1)
                }
                if (-9 > 1) {
                    if v2 {
                        fmt.Println(v4, v4)
                        fmt.Println(v2, v5, fi19)
                        v20 := GenS1{F0: fi19, F1: 76}
                        v21 := GenS2{F0: v2}
                        v3 = (v3 & (v3 & v3))
                        _ = v20
                        _ = v21
                    } else {
                        var v22 uint = 42
                        var v23 float64 = (-2.25 / 1.5)
                        fmt.Println(v3, fi18)
                        _ = v22
                        _ = v23
                    }
                    v13 = append(v13, fi19)
                } else {
                    var v24 float32 = 0.1
                    v13 = append(v13, v17)
                    var v25 uint64 = 660
                    _ = v24
                    _ = v25
                }
                v1 -= 2
            }
            var v26 float64 = float64(0.1)
            switch fi19 {
            case -862:
                var v27 int64 = (v4 * v4)
                _ = v27
            default:
                if (false || v2) {
                    switch v17 {
                    case 1:
                        v4 = -949
                    case -1:
                        var v28 int = fi18
                        _ = v28
                    }
                    if ((v14 <= v14) || v2) {
                        fmt.Println(v17)
                        v29 := GenS2{F0: (323 == 108)}
                        _ = v29
                    } else {
                        v4 = 42
                        fmt.Println(v2, v17, v2)
                        if ((v3 / ((v3) | 1)) >= (v3 * v3)) {
                            v3 -= v3
                            fmt.Println(v5, v26)
                        }
                    }
                    var v30 int = (v17 ^ 166)
                    _ = v30
                } else {
                    v2 = ((v2 || (919 == 1)) && ((-(v4)) < (^(v4))))
                }
            }
            v2 = (65535 != uint(v1))
            _ = fi19
            _ = v26
        }
        _ = fi18
    }
    var v31 int32 = (v3 ^ v3)
    v32 := GenS1{F0: (-(v17)), F1: 255}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
