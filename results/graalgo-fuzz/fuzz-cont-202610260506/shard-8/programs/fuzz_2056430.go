package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64; F1 int32; F2 int32 }
type GenS3 struct { F0 int8; F1 int }
type GenS4 struct { F0 uint32; F1 int }
type GenS5 struct { F0 int16 }
func genFn0() uint32 {
    var v5 uint32 = 4294967295
    switch "\n\t" {
    case "foo bar":
        if (-397 > 678) {
            for fi6 := 0; fi6 < 11; fi6++ {
                v5 = (^(v5))
                _ = fi6
            }
            v7 := GenS1{F0: (!((true || ((v5 & v5) > 2)) || true))}
            v8 := []int{350, 544}
            _ = v7
            _ = v8
        }
    case "hello":
        var v9 uint8 = 1
        _ = v9
    default:
        if (1 < 1) {
            v5 = 769
        }
    }
    fmt.Println(v5)
    for fi10 := 0; fi10 < 1; fi10++ {
        v11 := []int{127, (fi10 & fi10), fi10, 128}
        var v12 int32 = 1
        if false {
            v12 += v12
        } else {
            v13 := GenS2{F0: 542, F1: 1, F2: v12}
            fmt.Println(v5, v5, v5)
            _ = v13
        }
        _ = fi10
        _ = v11
        _ = v12
    }
    v14 := []int{-9223372036854775808, 42, -294, 415}
    _ = v14
    _ = v5
    return v5
}

func main() {
    v1 := []int{32767, -669, -150, 682}
    var v2 int8 = 127
    v3 := []int{}
    v4 := make(map[string]uint32)
    var v15 float32 = 0.0
    v16 := GenS1{F0: (!(false || (true && true)))}
    v17 := GenS3{F0: v2, F1: -1}
    var v18 uint = 1
    if (!((!(2 > 404)) && true)) {
        fmt.Println(v18, v2, v15)
        fmt.Println(v15, v2, v15)
        switch (("a" + "foo bar") + "foo bar") {
        case "":
            v15 = 0.1
        }
    }
    for fi19 := 0; fi19 < 17; fi19++ {
        v20 := GenS4{F0: 4294967295, F1: 239}
        _ = fi19
        _ = v20
    }
    v2 = 1
    if ((v15 * v15) <= v15) {
        var v21 int8 = v2
        switch "foo bar" {
        case "":
            v22 := GenS4{F0: 42, F1: -1}
            _ = v22
        case "foo bar":
            var v23 int16 = -1
            _ = v23
        default:
            var v24 int16 = 0
            _ = v24
        }
        switch 4294967295 {
        case 42:
            switch "" {
            case "hello":
                if true {
                    v15 += float32(0.1)
                    v18 -= 2
                } else {
                    var v25 int8 = (v21 % (((v2 * v2)) | 1))
                    if true {
                        v21 -= v25
                        v15 += v15
                        var v26 uint32 = 2147483647
                        _ = v26
                    } else {
                        var v27 string = "hello"
                        var v28 uint32 = 2147483648
                        var v29 uint8 = 127
                        _ = v27
                        _ = v28
                        _ = v29
                    }
                    _ = v25
                }
            default:
                for fi30 := 0; fi30 < 9; fi30++ {
                    var v31 uint = (^(v18))
                    v31 = (v18 - v18)
                    _ = fi30
                    _ = v31
                }
            }
        case -331:
            v32 := GenS5{F0: -1}
            _ = v32
        case 1:
            var v33 uint8 = 0
            _ = v33
        default:
            var v34 int = -9223372036854775808
            _ = v34
        }
        _ = v21
    } else {
        switch 0 {
        case 127:
            var v35 int32 = -1
            _ = v35
        case 789:
            var v36 uint16 = 491
            _ = v36
        }
        if len(v3) > 0 {
            v3[0] = 42
        }
        var v37 int32 = 42
        _ = v37
    }
    if (!false) {
        var v38 int64 = -375
        v38 -= v38
        _ = v38
    } else {
        for fi39 := 0; fi39 < 15; fi39++ {
            _ = len(v1)
            _ = fi39
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
