package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8 }
func genFn0(p0 uint) uint8 {
    v4 := []int{923, 0, 4294967295}
    v5 := GenS1{F0: -32768, F1: (^(p0))}
    if false {
        p0 -= (p0 - 545)
    } else {
        var v6 uint16 = 65535
        p0 += (p0 + (p0 | (p0 >> ((p0) & 63))))
        p0 = (^(p0))
        _ = v6
    }
    var v7 int64 = 1
    var v8 float32 = 0.1
    _ = p0
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    return 79
}

func main() {
    v1 := make(map[string]int)
    var v2 uint = 65535
    var v3 float32 = 0.0
    v9 := []int{127, 9223372036854775807}
    for fi10 := 0; fi10 < 14; fi10++ {
        fmt.Println(v2, fi10)
        switch (fi10 << ((v2) & 63)) {
        case 127:
            var v11 bool = (!true)
            _ = v11
        }
        _ = fi10
    }
    var v12 int = int(v2)
    if false {
        var v13 int8 = -95
        _ = v13
    }
    for fi14 := 0; fi14 < 9; fi14++ {
        fmt.Println(fi14, v2)
        var v15 int32 = 42
        v16 := GenS1{F0: 1, F1: v2}
        _ = fi14
        _ = v15
        _ = v16
    }
    var v17 uint32 = 2
    v18 := GenS1{F0: 0, F1: v2}
    var v19 string = ("" + ("a" + ""))
    switch ("a" + "foo bar") {
    case "\n\t":
        fmt.Println(v17)
    case "":
        var v20 uint16 = 362
        _ = v20
    case "hello":
        v2 = (^(v2))
    }
    v21 := GenS2{F0: -2147483648}
    var v22 int32 = 1
    switch v19 {
    case "":
        fmt.Println(v2, v3)
    case "hello":
        switch (v12 - v12) {
        case 42:
            var v23 float64 = (3.14159 * ((float64(((v3 * 1e10) * ((v3 / 1.0) / v3))) + (((-2.25 * -2.25) * 0.1) - ((((1.5 / 1.0) * ((1e10 - 0.0) - (0.0 * float64(v3)))) / (0.0 + -2.25)) - -2.25))) / 1.5))
            _ = v23
        case 32768:
            fmt.Println(v12, v19)
        default:
            v24 := GenS1{F0: -1, F1: uint(v17)}
            _ = v24
        }
    default:
        v25 := GenS2{F0: int32(v12)}
        _ = v25
    }
    v2 = v2
    var v26 string = ("" + ("\n\t" + "foo bar"))
    for fi27 := 0; fi27 < 20; fi27++ {
        for fi28 := 0; fi28 < 15; fi28++ {
            v29 := GenS2{F0: (^(v22))}
            switch (fi28 + (fi28 << ((v2) & 63))) {
            case 9223372036854775807:
                var v30 float32 = v3
                _ = v30
            case 32768:
                var v31 uint8 = genFn0(v2)
                _ = v31
            default:
                var v32 string = v19
                _ = v32
            }
            _ = fi28
            _ = v29
        }
        _ = len(v9)
        if true {
            if (!true) {
                var v33 string = v19
                _ = v33
            }
            var v34 int = 128
            v34 = v12
            _ = v34
        } else {
            fmt.Println(fi27)
            var v35 int32 = 42
            _ = v35
        }
        _ = fi27
    }
    v3 = ((1e10 * 0.1) - 0.1)
    if true {
        var v36 uint8 = 221
        v17 += v17
        _ = v36
    } else {
        v37 := GenS3{F0: -92}
        for fi38 := 0; fi38 < 14; fi38++ {
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            if ((v17 * (v17 & (v17 / ((2147483648) | 1)))) < (v17 | (v17 * v17))) {
                switch v19 {
                case "":
                    v22 += -1
                }
            }
            _ = fi38
        }
        _ = v37
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v12
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v26
    _ = v3
    _ = v9
}
