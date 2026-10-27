package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 string = "foo bar"
    v2 := make(map[int]uint32)
    var v3 float32 = ((float32(float64(-2.25)) + -2.25) - 1e10)
    var v4 int16 = 328
    var v5 int32 = 1
    var v6 string = ((v1 + v1) + "")
    _ = len(v2)
    v3 = 0.1
    v7 := GenS1{F0: ((float64(0.1) / 1.0) < 3.14159), F1: 2}
    var v8 int16 = v4
    if true {
        for fi9 := 0; fi9 < 15; fi9++ {
            var v10 uint32 = 792
            if false {
                var v11 int8 = 1
                v12 := GenS1{F0: ((false && true) && (831 >= 2)), F1: (v10 & (v10 / (((^(v10))) | 1)))}
                _ = v11
                _ = v12
            }
            var v13 uint64 = 646
            _ = fi9
            _ = v10
            _ = v13
        }
    }
    if (255 <= 0) {
        delete(v2, int(v5))
    }
    if false {
        for fi14 := 0; fi14 < 9; fi14++ {
            fmt.Println(v5, v3, v6)
            if (!false) {
                var v15 int16 = v8
                v1 = (v1 + (v1 + v1))
                fmt.Println(fi14, fi14, fi14)
                _ = v15
            } else {
                if false {
                    v5 = v5
                    v16 := v2[127]
                    _ = v16
                }
            }
            var v17 uint16 = 42
            _ = fi14
            _ = v17
        }
        fmt.Println(v4, v1, v3)
        if true {
            v2[-434] = 142
        }
    }
    switch "a" {
    case "foo bar":
        switch -9223372036854775808 {
        case 32768:
            v18 := GenS2{F0: 2}
            _ = v18
        }
    case "hello":
        var v19 int16 = v4
        _ = v19
    default:
        fmt.Println(v8, v5, v8)
    }
    var v20 int8 = 56
    if false {
        var v21 int8 = -49
        if ("hello" >= v6) {
            if (1 < 65535) {
                for fi22 := 0; fi22 < 8; fi22++ {
                    fmt.Println(v3)
                    _ = fi22
                }
                v2[-9223372036854775808] = 0
                var v23 float64 = 0.1
                _ = v23
            } else {
                fmt.Println(v4, v6, v20)
                v6 = v6
                var v24 uint64 = 673
                _ = v24
            }
        } else {
            for fi25 := 0; fi25 < 20; fi25++ {
                v26 := v2[fi25]
                _ = fi25
                _ = v26
            }
        }
        _ = v21
    }
    v27 := []int{320}
    v1 = (v6 + v1)
    v20 -= v20
    switch -889 {
    case -9223372036854775808:
        v20 -= v20
    case 128:
        if true {
            fmt.Println(v4)
        } else {
            if len(v27) > 0 {
                v27[0] = 9223372036854775807
            }
            v28 := GenS1{F0: false, F1: 690}
            _ = v28
        }
    case 934:
        _ = len(v27)
    default:
        v29 := v2[-2147483648]
        _ = v29
    }
    var v30 int64 = -1
    for fi31 := 0; fi31 < 9; fi31++ {
        v1 = v1
        _ = fi31
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v27
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
