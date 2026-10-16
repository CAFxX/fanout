package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 uint64; F1 uint8 }
func genFn0(p0 uint8) bool {
    if ((0.1 - 1.5) == float32(0.0)) {
        var v5 int32 = -2147483648
        _ = v5
    } else {
        if true {
            p0 -= (p0 | (p0 + 79))
            if true {
                if (-728 != 534) {
                    var v6 int64 = 42
                    _ = v6
                }
            } else {
                p0 = (p0 * (p0 / ((1) | 1)))
                var v7 uint = 2
                _ = v7
            }
            v8 := []int{}
            _ = v8
        } else {
            fmt.Println(p0)
        }
        fmt.Println(p0)
    }
    switch 32768 {
    case 231:
        v9 := GenS1{F0: 8, F1: 14}
        _ = v9
    case -2147483648:
        for fi10 := 0; fi10 < 3; fi10++ {
            var v11 string = ""
            var v12 int8 = 127
            for fi13 := 0; fi13 < 17; fi13++ {
                if (!false) {
                    var v14 uint16 = 513
                    _ = v14
                }
                _ = fi13
            }
            _ = fi10
            _ = v11
            _ = v12
        }
    case 985:
        fmt.Println(p0)
    default:
        v15 := GenS1{F0: -736, F1: 2}
        _ = v15
    }
    fmt.Println(p0)
    v16 := GenS1{F0: -2147483648, F1: 255}
    _ = p0
    _ = v16
    return (128 >= 9223372036854775807)
}

func main() {
    var v1 int8 = 127
    var v2 uint32 = 642
    v3 := make(map[string]bool)
    v4 := make(map[string]uint32)
    if false {
        var v17 string = "hello"
        _ = v17
    }
    v18 := GenS1{F0: 32768, F1: 0}
    var v19 uint8 = 138
    v20 := GenS2{F0: 4294967295, F1: v19}
    v2 += v2
    for fi21 := 0; fi21 < 9; fi21++ {
        if true {
            v2 += v2
            var v22 int = (fi21 | (fi21 >> ((4294967295) & 63)))
            v1 = v1
            _ = v22
        }
        _ = fi21
    }
    if (2147483647 <= -629) {
        var v23 int32 = -436
        _ = v23
    }
    fmt.Println(v2)
    v19 += 127
    v24 := GenS1{F0: -2147483648, F1: 255}
    v25 := []int{4294967295, -20}
    switch "" {
    case "a":
        var v26 int32 = -2147483648
        _ = v26
    case "foo bar":
        if ("" == "hello") {
            v19 = (v19 + v19)
            for fi27 := 0; fi27 < 9; fi27++ {
                switch fi27 {
                case 127:
                    v19 -= (v19 - (v19 % (((^(v19))) | 1)))
                }
                var v28 float32 = (1e10 - (-2.25 / 1e10))
                _ = fi27
                _ = v28
            }
            for fi29 := 0; fi29 < 8; fi29++ {
                v19 += 30
                fmt.Println(v19)
                _ = fi29
            }
        } else {
            var v30 int32 = -2147483648
            v1 = v1
            var v31 int8 = v1
            _ = v30
            _ = v31
        }
    default:
        var v32 uint64 = 2
        _ = v32
    }
    v19 = 128
    var v33 string = "foo bar"
    v33 = "hello"
    for fi34 := 0; fi34 < 5; fi34++ {
        var v35 int16 = 979
        var v36 int32 = -2147483648
        var v37 uint = 255
        _ = fi34
        _ = v35
        _ = v36
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v3
    _ = v33
    _ = v4
}
