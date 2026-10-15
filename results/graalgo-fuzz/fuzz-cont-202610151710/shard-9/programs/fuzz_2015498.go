package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 string; F1 uint32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint8 }
func genFn0() int64 {
    var v5 bool = true
    var v6 uint32 = 488
    for fi7 := 0; fi7 < 9; fi7++ {
        if v5 {
            v6 = v6
        }
        fmt.Println(fi7)
        for fi8 := 0; fi8 < 5; fi8++ {
            var v9 string = ("" + "foo bar")
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    if v5 {
        for fi10 := 0; fi10 < 11; fi10++ {
            v5 = ((!true) && v5)
            v6 = v6
            _ = fi10
        }
        var v11 uint64 = 4294967295
        _ = v11
    } else {
        for fi12 := 0; fi12 < 1; fi12++ {
            v13 := GenS1{F0: 600}
            var v14 int64 = 1
            _ = fi12
            _ = v13
            _ = v14
        }
        fmt.Println(v5)
    }
    _ = v5
    _ = v6
    return 575
}

func main() {
    var v1 float64 = (((1.5 + 1.5) * ((-2.25 + 1.5) / 0.1)) * 0.0)
    var v2 int64 = 1
    v3 := make(map[string]bool)
    var v4 int8 = 78
    v15 := GenS1{F0: 4294967295}
    if true {
        fmt.Println(v4, v1, v2)
        var v16 int16 = 876
        var v17 bool = true
        _ = v16
        _ = v17
    } else {
        v4 -= (v4 / ((-128) | 1))
        for fi18 := 0; fi18 < 5; fi18++ {
            v19 := []int{(fi18 / ((fi18) | 1)), (fi18 / (((fi18 >> ((288) & 63))) | 1)), fi18}
            var v20 int8 = -1
            _ = fi18
            _ = v19
            _ = v20
        }
    }
    v21 := GenS2{F0: (1 < 65535), F1: 4294967295}
    v4 = v4
    if true {
        var v22 bool = (0.0 < 3.14159)
        v1 = v1
        v22 = ((!v22) && v22)
        _ = v22
    }
    for fi23 := 0; fi23 < 20; fi23++ {
        v2 += v2
        v2 -= (v2 | (v2 | (v2 & v2)))
        _ = fi23
    }
    v24 := GenS3{F0: "a", F1: 4294967295}
    for fi25 := 0; fi25 < 17; fi25++ {
        switch "a" {
        case "":
            v26 := GenS3{F0: "", F1: 504}
            _ = v26
        }
        var v27 int8 = 1
        _ = fi25
        _ = v27
    }
    v28 := []int{int(v4), -2147483648, 32767}
    v1 = 1.5
    v29 := GenS4{F0: "\n\t"}
    v30 := GenS4{F0: "foo bar"}
    switch -1 {
    case 1:
        var v31 float32 = 0.0
        _ = v31
    case 128:
        var v32 int8 = v4
        _ = v32
    }
    for fi33 := 0; fi33 < 6; fi33++ {
        v2 -= -400
        var v34 uint64 = 2
        _ = fi33
        _ = v34
    }
    v35 := GenS1{F0: 65535}
    if true {
        v36 := GenS1{F0: 4294967295}
        var v37 bool = false
        var v38 int16 = 1
        _ = v36
        _ = v37
        _ = v38
    } else {
        switch 127 {
        case 2147483647:
            if true {
                switch "" {
                case "foo bar":
                    switch -900 {
                    case 996:
                        v4 = (v4 % ((v4) | 1))
                    case -513:
                        switch 857 {
                        case 127:
                            v39 := GenS5{F0: 1}
                            _ = v39
                        case -478:
                            var v40 uint32 = 42
                            _ = v40
                        case -1:
                            fmt.Println(v1)
                        default:
                            v28 = append(v28, -269)
                        }
                    default:
                        v1 -= -2.25
                    }
                case "\n\t":
                    fmt.Println(v2, v2, v4)
                }
                fmt.Println(v2, v1)
                v2 -= (^(v2))
            }
        default:
            v1 += v1
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v24
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v35
    _ = v4
}
