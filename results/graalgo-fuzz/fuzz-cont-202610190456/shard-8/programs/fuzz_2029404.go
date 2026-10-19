package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 int8 }
func genFn0() uint64 {
    for fi4 := 0; fi4 < 12; fi4++ {
        v5 := []int{-9223372036854775808, fi4}
        fmt.Println(fi4)
        _ = fi4
        _ = v5
    }
    var v6 uint32 = 2147483647
    for fi7 := 0; fi7 < 12; fi7++ {
        v6 += (v6 ^ v6)
        var v8 int32 = 1
        _ = fi7
        _ = v8
    }
    v9 := []int{32768}
    _ = v6
    _ = v9
    return 637
}

func main() {
    var v1 uint16 = 42
    var v2 bool = true
    v3 := make(map[string]bool)
    switch 375 {
    case 5:
        v2 = ((v2 && (42 != 1000)) && v2)
    case 230:
        switch "a" {
        case "\n\t":
            v1 = (v1 * (v1 % (((v1 * v1)) | 1)))
        case "hello":
            for fi10 := 0; fi10 < 20; fi10++ {
                fmt.Println(v1)
                var v11 int32 = 675
                _ = fi10
                _ = v11
            }
        }
    case -9223372036854775808:
        if (v2 || v2) {
            var v12 int = 0
            if (0 < 1) {
                fmt.Println(v1)
                if v2 {
                    v13 := []int{}
                    _ = v13
                }
            }
            _ = v12
        }
    default:
        if (872 > -725) {
            var v14 uint16 = (v1 / ((v1) | 1))
            var v15 uint64 = 1
            _ = v14
            _ = v15
        }
    }
    if ((float32(0.1) > -2.25) || v2) {
        v16 := GenS1{F0: -32768, F1: 9223372036854775807, F2: -128}
        v1 += 32767
        if v2 {
            v1 -= 255
            var v17 float64 = 0.1
            _ = v17
        } else {
            v18 := GenS1{F0: 792, F1: 4294967295, F2: -1}
            _ = v18
        }
        _ = v16
    }
    if (v2 || (4294967295 < 4294967295)) {
        v19 := GenS1{F0: 1, F1: 128, F2: 42}
        var v20 float32 = (3.14159 / (1.5 - (3.14159 * 1e10)))
        _ = v19
        _ = v20
    }
    var v21 uint = 0
    for fi22 := 0; fi22 < 16; fi22++ {
        if false {
            v23 := GenS1{F0: 0, F1: (fi22 ^ fi22), F2: -128}
            _ = v23
        }
        var v24 int32 = -1
        v24 += 977
        _ = fi22
        _ = v24
    }
    v25 := GenS1{F0: -755, F1: 385, F2: 0}
    switch 1 {
    case -456:
        var v26 string = (("" + "hello") + "\n\t")
        _ = v26
    }
    var v27 int16 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v21
    _ = v25
    _ = v27
    _ = v3
}
