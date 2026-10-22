package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 int64; F1 uint64; F2 int }
func genFn0(p0 int8) string {
    if (18446744073709551615 >= 811) {
        fmt.Println(p0)
        v6 := []int{42}
        _ = v6
    }
    var v7 float32 = 1.5
    var v8 string = "foo bar"
    switch 4294967295 {
    case 0:
        fmt.Println(v7, v8)
    case 32768:
        p0 -= 42
    case 737:
        var v9 bool = false
        _ = v9
    default:
        var v10 int64 = -9223372036854775808
        _ = v10
    }
    _ = p0
    _ = v7
    _ = v8
    return v8
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]int)
    var v3 int16 = -32768
    v4 := make(map[string]int)
    var v5 int8 = 100
    fmt.Println(v3)
    if (float64(0.1) < (1.5 * 1.5)) {
        var v11 int = 128
        _ = len(v2)
        switch -678 {
        case 4294967295:
            v3 = v3
        }
        _ = v11
    } else {
        v12 := v1[65535]
        v13 := []int{9223372036854775807, 998, 42, 1}
        switch 127 {
        case 4294967295:
            if v12 {
                _ = len(v13)
            } else {
                var v14 float64 = 0.0
                _ = v14
            }
        default:
            v15 := GenS1{F0: 42, F1: 127}
            _ = v15
        }
        _ = v12
        _ = v13
    }
    v16 := []int{128, 756, 127, 32767}
    v17 := GenS1{F0: -9223372036854775808, F1: (-(v5))}
    v18 := GenS2{F0: 947, F1: 932, F2: 617}
    var v19 uint64 = 2
    var v20 int64 = -71
    v19 = v19
    if true {
        var v21 uint8 = 255
        if (("a" + "foo bar") != "\n\t") {
            _ = len(v16)
            v22 := v1[0]
            v3 = (v3 & (v3 << ((4294967295) & 63)))
            _ = v22
        } else {
            for fi23 := 0; fi23 < 11; fi23++ {
                v20 = v20
                fmt.Println(v5, v3)
                v21 = (^(v21))
                _ = fi23
            }
            if false {
                switch -1 {
                case -623:
                    delete(v4, "hello")
                }
            }
        }
        var v24 bool = (42 < 2)
        _ = v21
        _ = v24
    } else {
        v3 -= v3
        for fi25 := 0; fi25 < 8; fi25++ {
            if (!true) {
                var v26 float64 = -2.25
                if true {
                    v27 := GenS1{F0: 9223372036854775807, F1: v5}
                    v28 := GenS1{F0: (v20 / ((v20) | 1)), F1: v5}
                    for fi29 := 0; fi29 < 15; fi29++ {
                        v30 := v4[""]
                        v16 = append(v16, 9223372036854775807)
                        v30 = 622
                        v5 += v5
                        v30 = 42
                        _ = fi29
                        _ = v30
                    }
                    _ = v27
                    _ = v28
                }
                switch fi25 {
                case 42:
                    var v31 uint8 = 128
                    _ = v31
                case 9223372036854775807:
                    v26 -= v26
                case 1:
                    v20 -= -368
                default:
                    if true {
                        for fi32 := 0; fi32 < 18; fi32++ {
                            var v33 uint16 = 42
                            v33 = v33
                            v26 = 0.1
                            _ = len(v1)
                            _ = fi32
                            _ = v33
                        }
                    } else {
                        for fi34 := 0; fi34 < 14; fi34++ {
                            v26 -= (((float64(-2.25) * v26) / v26) - v26)
                            fmt.Println(v5)
                            _ = fi34
                        }
                    }
                }
                _ = v26
            }
            _ = fi25
        }
    }
    v19 = (v19 / ((v19) | 1))
    v19 -= (v19 - (v19 * (^(v19))))
    _ = len(v1)
    var v35 string = ""
    var v36 uint8 = 42
    v3 = 1000
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
