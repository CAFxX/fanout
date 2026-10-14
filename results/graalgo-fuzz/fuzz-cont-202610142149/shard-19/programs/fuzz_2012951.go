package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 string }
type GenS2 struct { F0 uint32; F1 uint32 }
func genFn0(p0 bool, p1 uint16) uint16 {
    for fi4 := 0; fi4 < 15; fi4++ {
        if p0 {
            p1 -= p1
            p1 = (p1 / (((p1 >> ((691) & 63))) | 1))
        } else {
            for fi5 := 0; fi5 < 5; fi5++ {
                p1 += (^(p1))
                fmt.Println(p1)
                p1 -= p1
                _ = fi5
            }
            if false {
                if false {
                    var v6 int16 = 32767
                    v7 := GenS1{F0: ((-35 >= 102) && (127 == 127)), F1: 42, F2: "hello"}
                    for fi8 := 0; fi8 < 3; fi8++ {
                        v9 := GenS1{F0: false, F1: 42, F2: "hello"}
                        switch "foo bar" {
                        case "hello":
                            fmt.Println(p1, v6)
                        case "\n\t":
                            p0 = p0
                        case "foo bar":
                            if (1 >= 488) {
                                v10 := []int{}
                                _ = v10
                            }
                        default:
                            var v11 uint64 = uint64(p1)
                            _ = v11
                        }
                        var v12 float32 = (3.14159 + -2.25)
                        _ = fi8
                        _ = v12
                        _ = v9
                    }
                    _ = v6
                    _ = v7
                } else {
                    p0 = p0
                }
                for fi13 := 0; fi13 < 6; fi13++ {
                    p0 = (true && p0)
                    var v14 uint16 = 1
                    var v15 float32 = ((3.14159 * (float32(-2.25) / 1.0)) * 3.14159)
                    _ = fi13
                    _ = v14
                    _ = v15
                }
            } else {
                v16 := []int{(fi4 * (fi4 - fi4))}
                p0 = (42 > 42)
                _ = v16
            }
        }
        var v17 int8 = -85
        _ = fi4
        _ = v17
    }
    v18 := []int{-414, -37, 459}
    var v19 uint = 20
    fmt.Println(v19, v19, p1)
    _ = p0
    _ = p1
    _ = v18
    _ = v19
    return 0
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int = 1
    var v3 uint64 = 4294967295
    v20 := GenS1{F0: ((true || false) && true), F1: 2147483647, F2: "a"}
    var v21 uint32 = 42
    v22 := GenS2{F0: v21, F1: 4294967295}
    v23 := GenS2{F0: v21, F1: (v21 | 42)}
    fmt.Println(v3)
    var v24 int8 = -1
    if (false || ((255 != 981) && true)) {
        var v25 bool = false
        switch "" {
        case "foo bar":
            var v26 int64 = 98
            _ = v26
        }
        if (!(681 > 18446744073709551615)) {
            delete(v1, 65535)
            v27 := GenS1{F0: false, F1: v21, F2: "a"}
            if v25 {
                for fi28 := 0; fi28 < 3; fi28++ {
                    for fi29 := 0; fi29 < 18; fi29++ {
                        v30 := GenS2{F0: 2, F1: 2}
                        var v31 int = fi28
                        fmt.Println(v24)
                        _ = fi29
                        _ = v30
                        _ = v31
                    }
                    v2 -= v2
                    _ = fi28
                }
            }
            _ = v27
        }
        _ = v25
    } else {
        var v32 uint32 = 4294967295
        for fi33 := 0; fi33 < 1; fi33++ {
            var v34 int64 = 0
            for fi35 := 0; fi35 < 18; fi35++ {
                var v36 float64 = 3.14159
                _ = fi35
                _ = v36
            }
            _ = fi33
            _ = v34
        }
        _ = v32
    }
    v37 := []int{32768, v2, v2, (v2 ^ 621)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v37:", len(v37), v37)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v37
}
