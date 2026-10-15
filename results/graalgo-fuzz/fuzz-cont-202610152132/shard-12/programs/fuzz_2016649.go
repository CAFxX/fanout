package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 bool; F1 uint8 }
type GenS4 struct { F0 int64; F1 int32 }
func genFn0() uint32 {
    v7 := GenS1{F0: -9223372036854775808, F1: 1}
    var v8 int16 = 42
    for fi9 := 0; fi9 < 7; fi9++ {
        v8 = (v8 / (((v8 / ((106) | 1))) | 1))
        _ = fi9
    }
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 float64 = (-2.25 - -2.25)
    v2 := []uint32{}
    var v3 int = 895
    var v4 int32 = 343
    v5 := make(map[string]string)
    var v6 int32 = v4
    v2 = append(v2, 1)
    v10 := v5[("hello" + ("foo bar" + "\n\t"))]
    fmt.Println(v10, v6, v6)
    if (false && ("" > v10)) {
        var v11 int16 = -52
        v12 := GenS1{F0: (v3 ^ 4294967295), F1: 1000}
        var v13 int32 = (v4 + 1)
        _ = v11
        _ = v12
        _ = v13
    }
    var v14 float32 = -2.25
    v14 = v14
    var v15 int64 = 0
    switch -2147483648 {
    case 32768:
        for fi16 := 0; fi16 < 8; fi16++ {
            if false {
                if len(v2) > 0 {
                    v2[0] = 65535
                }
                var v17 uint16 = 776
                _ = v17
            } else {
                v18 := GenS1{F0: 127, F1: -19}
                _ = v18
            }
            _ = fi16
        }
    default:
        v19 := GenS2{F0: 1}
        _ = v19
    }
    _ = len(v5)
    if false {
        if false {
            for fi20 := 0; fi20 < 12; fi20++ {
                if true {
                    var v21 uint64 = 657
                    fmt.Println(fi20)
                    fmt.Println(v4, v21)
                    _ = v21
                } else {
                    for fi22 := 0; fi22 < 2; fi22++ {
                        fmt.Println(fi20)
                        _ = fi22
                    }
                    v5["a"] = ("" + (v10 + v10))
                    var v23 float64 = v1
                    _ = v23
                }
                _ = fi20
            }
            var v24 uint = 4294967295
            _ = v24
        } else {
            v25 := GenS3{F0: (v10 < v10), F1: 128}
            _ = v25
        }
        switch (v10 + v10) {
        case "foo bar":
            _ = len(v5)
        default:
            for fi26 := 0; fi26 < 17; fi26++ {
                switch v3 {
                case 127:
                    fmt.Println(v15, fi26)
                }
                var v27 int32 = (-(v6))
                v28 := GenS3{F0: (!((((-(v3)) != fi26) && (v10 != v10)) && (!false))), F1: 128}
                _ = fi26
                _ = v27
                _ = v28
            }
        }
    }
    if (!(((128 < 255) && ((((float32(float64(v14)) + float32((v1 + v1))) + v14) / (-2.25 / (0.1 + v14))) > (v14 / v14))) && false)) {
        v2 = append(v2, genFn0())
        v6 -= 2147483647
    }
    v1 -= (3.14159 - v1)
    v4 = v6
    if (true || (229 < 127)) {
        for fi29 := 0; fi29 < 13; fi29++ {
            v30 := GenS2{F0: 32767}
            for fi31 := 0; fi31 < 10; fi31++ {
                v5["hello"] = (((v10 + v10) + v10) + ((v10 + v10) + ""))
                _ = fi31
            }
            v3 = v3
            _ = fi29
            _ = v30
        }
        var v32 int64 = (-(v15))
        _ = v32
    } else {
        var v33 int16 = 1000
        _ = v33
    }
    v34 := GenS4{F0: (v15 ^ v15), F1: -486}
    switch v10 {
    case "\n\t":
        v4 = v4
    case "foo bar":
        var v35 float32 = v14
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
