package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 bool }
type GenS2 struct { F0 uint64 }
func genFn0(p0 uint64) uint16 {
    for fi6 := 0; fi6 < 15; fi6++ {
        for fi7 := 0; fi7 < 19; fi7++ {
            p0 = (^(p0))
            _ = fi7
        }
        p0 = (p0 & p0)
        for fi8 := 0; fi8 < 20; fi8++ {
            v9 := []int{2147483647, 32767}
            switch "foo bar" {
            case "\n\t":
                if (((3.14159 + 3.14159) != 1.5) && (0 >= 925)) {
                    v10 := GenS1{F0: -1, F1: 1000, F2: (-525 < -1)}
                    if (42 <= 42) {
                        p0 = p0
                        v9 = append(v9, (^(fi8)))
                        if true {
                            if (595 > 0) {
                                p0 = p0
                            } else {
                                if false {
                                    if true {
                                        v11 := GenS2{F0: p0}
                                        v12 := GenS1{F0: 0, F1: 1, F2: false}
                                        _ = v11
                                        _ = v12
                                    } else {
                                        p0 = (p0 - (p0 << ((250) & 63)))
                                        var v13 int32 = -1
                                        _ = v13
                                    }
                                }
                                var v14 float64 = ((0.1 + 0.0) - 1.5)
                                _ = v14
                            }
                        } else {
                            var v15 float32 = 1.5
                            _ = v15
                        }
                    } else {
                        var v16 uint = 65535
                        _ = v16
                    }
                    _ = v10
                }
            default:
                if len(v9) > 0 {
                    v9[0] = 32768
                }
            }
            switch "foo bar" {
            case "hello":
                p0 -= p0
            case "\n\t":
                var v17 float64 = 1e10
                _ = v17
            }
            _ = fi8
            _ = v9
        }
        _ = fi6
    }
    var v18 int64 = 42
    v19 := GenS2{F0: 1}
    for fi20 := 0; fi20 < 1; fi20++ {
        var v21 uint64 = 1
        v18 += 42
        v22 := []int{}
        _ = fi20
        _ = v21
        _ = v22
    }
    _ = p0
    _ = v18
    _ = v19
    return 255
}

func main() {
    var v1 uint = 533
    v2 := make(map[uint32]int)
    var v3 uint = 0
    var v4 uint = v3
    v5 := []string{(("foo bar" + ("a" + "foo bar")) + "a"), "foo bar", "\n\t"}
    var v23 float64 = (0.1 / (((3.14159 + 3.14159) - float64(0.0)) - 1.5))
    var v24 uint32 = 2147483647
    switch 315 {
    case -2147483648:
        var v25 float64 = ((float64(0.1) * (((((v23 / v23) / ((((float64(1.5) + (v23 * 1e10)) + 3.14159) + v23) / 1.0)) / 1.0) - v23) / v23)) * v23)
        _ = v25
    default:
        switch -2147483648 {
        case 832:
            var v26 int64 = 1
            _ = v26
        }
    }
    fmt.Println(v1, v23, v4)
    for fi27 := 0; fi27 < 20; fi27++ {
        if ((v1 >> ((v3) & 63)) != (v3 << ((2) & 63))) {
            v4 -= 863
            var v28 uint32 = v24
            delete(v2, v28)
            _ = v28
        }
        fmt.Println(v3, v3)
        _ = fi27
    }
    v24 += v24
    delete(v2, (v24 - v24))
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v1 = v4
    switch "hello" {
    case "hello":
        if true {
            var v29 uint32 = v24
            var v30 bool = (893 == 1)
            _ = v29
            _ = v30
        }
    }
    var v31 string = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
