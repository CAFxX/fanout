package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 uint16; F2 uint }

func main() {
    var v1 string = "a"
    var v2 int64 = -1
    var v3 int16 = 0
    v4 := make(map[string]string)
    if true {
        var v5 uint32 = 4294967295
        v6 := GenS1{F0: 42}
        _ = v5
        _ = v6
    }
    for fi7 := 0; fi7 < 3; fi7++ {
        if (!true) {
            var v8 bool = false
            _ = v8
        } else {
            switch 128 {
            case 127:
                for fi9 := 0; fi9 < 14; fi9++ {
                    var v10 uint8 = 128
                    _ = fi9
                    _ = v10
                }
            case 419:
                fmt.Println(fi7, v1, v3)
            default:
                if (!(-1 > 0)) {
                    var v11 float32 = 0.0
                    _ = v11
                } else {
                    v12 := v4[v1]
                    v13 := v4[""]
                    _ = v12
                    _ = v13
                }
            }
            if (-909 != 0) {
                var v14 bool = ((v1 + (v1 + "")) > (((("hello" + v1) + "") + v1) + v1))
                _ = v14
            } else {
                if (false || (((v3 > 649) && false) && false)) {
                    v2 += v2
                    for fi15 := 0; fi15 < 11; fi15++ {
                        var v16 float64 = ((0.0 * 0.0) * 1.5)
                        _ = fi15
                        _ = v16
                    }
                } else {
                    var v17 uint = 255
                    fmt.Println(v1)
                    _ = v17
                }
            }
            var v18 int = fi7
            _ = v18
        }
        _ = fi7
    }
    switch "foo bar" {
    case "":
        v3 = v3
    default:
        v3 = v3
    }
    if ((193 < 65535) && true) {
        if (!(279 != 1000)) {
            if (42 <= 0) {
                v3 += 1000
            } else {
                fmt.Println(v2, v3)
            }
        }
        v19 := []int{}
        var v20 int8 = 127
        _ = v19
        _ = v20
    } else {
        v21 := []int{32767, 2147483647, 4294967295, 4294967295}
        _ = v21
    }
    var v22 int8 = 100
    var v23 uint64 = 622
    if true {
        v24 := v4["a"]
        _ = v24
    } else {
        v4[v1] = ((v1 + (v1 + "")) + v1)
        if (int32(v2) == 42) {
            fmt.Println(v2)
            _ = len(v4)
        }
    }
    switch 32768 {
    case -913:
        fmt.Println(v22, v3, v22)
    case 32768:
        v25 := GenS2{F0: 0, F1: 42, F2: 1}
        _ = v25
    case 2147483647:
        for fi26 := 0; fi26 < 13; fi26++ {
            if (4294967295 < 2147483647) {
                fmt.Println(v2)
                var v27 int8 = (v22 ^ (v22 % ((100) | 1)))
                _ = v27
            } else {
                v28 := GenS1{F0: 2}
                switch v1 {
                case "":
                    var v29 int8 = int8(v23)
                    _ = v29
                }
                v2 -= (v2 | (v2 * (v2 | (v2 + -789))))
                _ = v28
            }
            _ = fi26
        }
    }
    v2 = 1
    v3 = -595
    switch v1 {
    case "":
        v1 = (v1 + "a")
    case "a":
        var v30 uint16 = 326
        _ = v30
    }
    var v31 int16 = (^(v3))
    v22 = v22
    switch v1 {
    case "foo bar":
        for fi32 := 0; fi32 < 15; fi32++ {
            var v33 uint8 = 209
            var v34 float32 = -2.25
            var v35 float32 = float32((((-2.25 + float64(v34)) - float64(v34)) - 1e10))
            _ = fi32
            _ = v33
            _ = v34
            _ = v35
        }
    default:
        v36 := GenS2{F0: 0, F1: 430, F2: 255}
        _ = v36
    }
    v37 := GenS1{F0: 44}
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v31
    _ = v37
    _ = v4
}
