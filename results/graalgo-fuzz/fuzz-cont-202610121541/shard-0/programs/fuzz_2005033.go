package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 int32 }
type GenS2 struct { F0 int64; F1 int }
func genFn0(p0 int) uint32 {
    p0 += (p0 / ((42) | 1))
    v4 := GenS1{F0: 127, F1: 127, F2: 856}
    fmt.Println(p0)
    fmt.Println(p0)
    _ = p0
    _ = v4
    return 65535
}

func main() {
    var v1 int32 = -1
    v2 := []int{32768}
    v3 := make(map[int]uint32)
    for fi5 := 0; fi5 < 14; fi5++ {
        fmt.Println(v1)
        switch ("\n\t" + "hello") {
        case "a":
            var v6 string = "a"
            _ = v6
        }
        var v7 float64 = -2.25
        _ = fi5
        _ = v7
    }
    _ = len(v2)
    for fi8 := 0; fi8 < 14; fi8++ {
        v1 = (-(v1))
        _ = fi8
    }
    v9 := v3[32768]
    v10 := v3[int(v1)]
    _ = len(v2)
    for fi11 := 0; fi11 < 1; fi11++ {
        if (2 >= 255) {
            if false {
                if len(v2) > 0 {
                    v2[0] = 117
                }
            } else {
                switch -818 {
                case 42:
                    var v12 float64 = (float64(0.1) * -2.25)
                    _ = v12
                case 32767:
                    for fi13 := 0; fi13 < 18; fi13++ {
                        if (127 >= 255) {
                            v1 -= 2147483647
                            var v14 int8 = 42
                            _ = v14
                        } else {
                            v15 := GenS1{F0: 42, F1: 42, F2: (v1 ^ v1)}
                            _ = v15
                        }
                        v1 = v1
                        _ = fi13
                    }
                }
                fmt.Println(v1, v10)
                for fi16 := 0; fi16 < 1; fi16++ {
                    if (!true) {
                        v17 := GenS1{F0: 42, F1: 27, F2: v1}
                        _ = v17
                    } else {
                        v18 := GenS2{F0: 42, F1: fi11}
                        fmt.Println(v10, fi11, fi16)
                        if true {
                            if (!(!true)) {
                                v9 = genFn0((fi11 - -634))
                                switch "" {
                                case "\n\t":
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                case "":
                                    var v19 float64 = (3.14159 + 1e10)
                                    _ = v19
                                case "a":
                                    v1 -= (^(v1))
                                }
                                var v20 int = fi11
                                _ = v20
                            } else {
                                var v21 string = ("" + "\n\t")
                                _ = v21
                            }
                            v10 -= v10
                            delete(v3, fi16)
                            v9 += v10
                            delete(v3, 0)
                            var v22 int16 = 42
                            v23 := GenS2{F0: -628, F1: fi11}
                            _ = len(v2)
                            var v24 int8 = 119
                            _ = v22
                            _ = v23
                            _ = v24
                        }
                        _ = v18
                    }
                    v1 = (v1 | v1)
                    _ = fi16
                }
            }
            if true {
                switch -1 {
                case 2147483647:
                    switch 127 {
                    case -1:
                        if len(v2) > 0 {
                            v2[0] = fi11
                        }
                    case -18:
                        var v25 int8 = 100
                        _ = v25
                    }
                case -949:
                    if false {
                        switch "" {
                        case "":
                            var v26 uint16 = 557
                            _ = v26
                        case "a":
                            v10 -= (v10 & (v9 / ((v10) | 1)))
                        }
                        var v27 int8 = 76
                        _ = v27
                    } else {
                        v2 = append(v2, (fi11 | fi11))
                        fmt.Println(fi11, v9, v9)
                    }
                default:
                    var v28 int = 127
                    _ = v28
                }
                var v29 int = fi11
                _ = v29
            } else {
                var v30 uint64 = 255
                var v31 uint16 = 457
                _ = v30
                _ = v31
            }
        }
        var v32 uint = 65535
        _ = fi11
        _ = v32
    }
    v33 := GenS2{F0: 1, F1: -2147483648}
    var v34 int16 = -32768
    for fi35 := 0; fi35 < 16; fi35++ {
        v10 -= v10
        _ = fi35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v9
}
