package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 int32; F1 bool }
func genFn0(p0 uint16) int {
    for fi7 := 0; fi7 < 16; fi7++ {
        v8 := []int{}
        v8 = append(v8, fi7)
        _ = fi7
        _ = v8
    }
    if (((false || false) || true) || true) {
        fmt.Println(p0)
        fmt.Println(p0)
    } else {
        var v9 int64 = 9223372036854775807
        _ = v9
    }
    for fi10 := 0; fi10 < 7; fi10++ {
        var v11 int8 = -84
        if true {
            v12 := []int{}
            _ = v12
        }
        _ = fi10
        _ = v11
    }
    v13 := []int{}
    for fi14 := 0; fi14 < 19; fi14++ {
        var v15 int8 = 42
        for fi16 := 0; fi16 < 2; fi16++ {
            var v17 int8 = v15
            fmt.Println(v15, v15)
            if (42 >= (p0 * p0)) {
                var v18 uint8 = 0
                p0 -= 0
                if (!(!true)) {
                    _ = len(v13)
                    var v19 uint = 712
                    var v20 float64 = 1e10
                    _ = v19
                    _ = v20
                }
                _ = v18
            }
            switch (^(fi14)) {
            case -9223372036854775808:
                switch fi14 {
                case 127:
                    v15 -= (v15 << ((42) & 63))
                    var v21 uint = 4294967295
                    _ = v21
                case 1:
                    var v22 float32 = 3.14159
                    _ = v22
                default:
                    if ((v15 | (^(v17))) <= -128) {
                        v15 = (v15 - (v17 >> ((2) & 63)))
                    }
                }
                var v23 int8 = (v15 & (v17 >> ((21) & 63)))
                _ = v23
            }
            _ = fi16
            _ = v17
        }
        v24 := GenS1{F0: 255, F1: 42}
        _ = fi14
        _ = v15
        _ = v24
    }
    _ = p0
    _ = v13
    return 0
}

func main() {
    v1 := make(map[int]string)
    var v2 float64 = 3.14159
    var v3 float64 = 1e10
    var v4 bool = false
    var v5 string = "\n\t"
    var v6 uint8 = 128
    if (!v4) {
        var v25 int = 32768
        var v26 int = (v25 % ((genFn0(2)) | 1))
        _ = v25
        _ = v26
    } else {
        var v27 float64 = v3
        delete(v1, 127)
        _ = v27
    }
    if (v3 > (1.5 + 0.0)) {
        if (!v4) {
            v28 := []int{genFn0(32767), 1, 4294967295}
            var v29 uint8 = (^(v6))
            v5 = v5
            _ = v28
            _ = v29
        }
        var v30 bool = ((419 >= 42) || (v4 || (-1 >= -1)))
        _ = v30
    }
    var v31 int8 = 89
    v6 = v6
    if (798 <= 18446744073709551615) {
        var v32 float64 = v3
        _ = v32
    }
    v33 := v1[823]
    var v34 int = 793
    fmt.Println(v6, v2, v2)
    var v35 uint16 = 632
    switch "a" {
    case "hello":
        switch v34 {
        case 637:
            var v36 uint16 = 255
            _ = v36
        case -591:
            v3 = v2
        default:
            var v37 uint8 = 255
            _ = v37
        }
    case "\n\t":
        var v38 uint = 1
        _ = v38
    case "":
        v39 := GenS2{F0: -205, F1: (-1 > -483)}
        _ = v39
    default:
        var v40 float32 = 1e10
        _ = v40
    }
    v35 = v35
    _ = len(v1)
    v31 += v31
    v41 := []int{genFn0(v35), genFn0(v35), 128}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
