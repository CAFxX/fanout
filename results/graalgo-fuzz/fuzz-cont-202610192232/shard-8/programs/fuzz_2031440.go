package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 string; F1 uint32 }

func main() {
    var v1 float32 = float32((float64(3.14159) * 0.1))
    v2 := []int{268}
    v3 := make(map[int]bool)
    for fi4 := 0; fi4 < 8; fi4++ {
        var v5 string = "a"
        _ = fi4
        _ = v5
    }
    _ = len(v2)
    if false {
        v1 = (((float32((0.0 / (0.0 * float64(((v1 / v1) - (0.0 + 3.14159)))))) - 0.1) + (v1 - 3.14159)) + v1)
        for fi6 := 0; fi6 < 9; fi6++ {
            v1 = v1
            var v7 string = ""
            _ = fi6
            _ = v7
        }
    }
    var v8 int32 = 2147483647
    for fi9 := 0; fi9 < 13; fi9++ {
        v10 := GenS1{F0: "a", F1: -1}
        _ = fi9
        _ = v10
    }
    switch 32768 {
    case 4294967295:
        if false {
            for fi11 := 0; fi11 < 5; fi11++ {
                fmt.Println(v1, v8, v8)
                var v12 int16 = -556
                var v13 int8 = 42
                _ = fi11
                _ = v12
                _ = v13
            }
            v14 := v3[0]
            v1 -= v1
            _ = v14
        }
    case -232:
        var v15 uint8 = 2
        _ = v15
    }
    switch -170 {
    case -842:
        var v16 uint = 255
        _ = v16
    case 32767:
        fmt.Println(v1)
    default:
        var v17 uint16 = 32767
        _ = v17
    }
    fmt.Println(v8, v8)
    if true {
        v8 -= v8
        v1 += ((v1 + v1) + 1.5)
    } else {
        if true {
            if (4294967295 != 0) {
                for fi18 := 0; fi18 < 16; fi18++ {
                    switch "foo bar" {
                    case "":
                        v19 := GenS1{F0: "hello", F1: 127}
                        _ = v19
                    case "foo bar":
                        var v20 uint64 = 2
                        _ = v20
                    }
                    fmt.Println(v8, v8)
                    v21 := GenS1{F0: ("a" + ("\n\t" + ("\n\t" + "\n\t"))), F1: 1}
                    _ = fi18
                    _ = v21
                }
            }
            var v22 float32 = -2.25
            _ = v22
        } else {
            switch "\n\t" {
            case "hello":
                var v23 uint8 = 42
                _ = v23
            case "a":
                _ = len(v3)
            case "":
                fmt.Println(v8, v8)
            }
            var v24 int32 = (v8 & (^(v8)))
            _ = v24
        }
        var v25 float32 = (float32((3.14159 + 0.0)) + (float32(3.14159) + v1))
        var v26 uint = 4294967295
        _ = v25
        _ = v26
    }
    var v27 int32 = v8
    switch "\n\t" {
    case "":
        for fi28 := 0; fi28 < 2; fi28++ {
            fmt.Println(fi28, fi28, v8)
            v29 := GenS1{F0: "", F1: 100}
            _ = fi28
            _ = v29
        }
    case "a":
        switch 51 {
        case 32768:
            for fi30 := 0; fi30 < 5; fi30++ {
                var v31 float64 = ((-2.25 / (3.14159 * ((1e10 / float64(((v1 * (float32((1.5 + (0.0 - (float64(v1) / 1.0)))) * v1)) - -2.25))) + (1e10 / 3.14159)))) + -2.25)
                v1 = ((v1 * (-2.25 * (v1 * v1))) / v1)
                _ = fi30
                _ = v31
            }
        case -758:
            v2 = append(v2, -9223372036854775808)
        case 4294967295:
            _ = len(v3)
        default:
            fmt.Println(v8, v27, v1)
        }
    }
    fmt.Println(v1, v1, v8)
    var v32 int16 = -32768
    var v33 int = 4294967295
    var v34 int16 = v32
    if ((0.0 + float64(v1)) >= 1e10) {
        v32 = v34
        v35 := v3[(v33 | v33)]
        _ = v35
    } else {
        var v36 uint8 = 2
        v2 = append(v2, v33)
        for fi37 := 0; fi37 < 20; fi37++ {
            v3[(fi37 >> ((0) & 63))] = (v36 > v36)
            for fi38 := 0; fi38 < 2; fi38++ {
                fmt.Println(v1)
                _ = fi38
            }
            v39 := GenS2{F0: "hello", F1: 42}
            _ = fi37
            _ = v39
        }
        _ = v36
    }
    v3[9223372036854775807] = false
    var v40 string = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v40:", v40)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v40
    _ = v8
}
