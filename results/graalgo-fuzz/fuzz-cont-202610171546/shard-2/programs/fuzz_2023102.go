package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 int }
type GenS2 struct { F0 int32; F1 uint32 }

func main() {
    var v1 float32 = 3.14159
    var v2 uint64 = 42
    var v3 int32 = 1
    v4 := make(map[int]uint32)
    for fi5 := 0; fi5 < 2; fi5++ {
        v6 := GenS1{F0: 4294967295, F1: 32767, F2: (fi5 >> ((24) & 63))}
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 13; fi7++ {
        if (!("a" <= "foo bar")) {
            v2 += (^(v2))
            fmt.Println(fi7, fi7, v1)
            switch fi7 {
            case -170:
                fmt.Println(v1, v3)
            case 9223372036854775807:
                v8 := v4[fi7]
                _ = v8
            default:
                fmt.Println(fi7, v3)
            }
        } else {
            var v9 int16 = -543
            _ = v9
        }
        _ = fi7
    }
    for fi10 := 0; fi10 < 12; fi10++ {
        switch "a" {
        case "":
            var v11 uint = 255
            _ = v11
        case "a":
            var v12 float64 = (-2.25 + (0.1 * (((-2.25 + 3.14159) / 3.14159) - ((1e10 + 0.1) * 3.14159))))
            _ = v12
        default:
            v13 := GenS1{F0: 1, F1: 32767, F2: (fi10 & 4294967295)}
            _ = v13
        }
        var v14 int16 = 42
        _ = fi10
        _ = v14
    }
    v15 := v4[2147483647]
    v16 := []int{598, 771, 128, 2147483647}
    v15 = (^(v15))
    v16 = append(v16, 569)
    if ((!(!true)) || false) {
        if (507 < 18446744073709551615) {
            v17 := GenS1{F0: (v15 | 0), F1: 153, F2: 630}
            fmt.Println(v1, v2, v15)
            var v18 bool = (true && false)
            _ = v17
            _ = v18
        }
        switch 0 {
        case 42:
            for fi19 := 0; fi19 < 19; fi19++ {
                for fi20 := 0; fi20 < 16; fi20++ {
                    switch fi19 {
                    case 127:
                        v4[2147483647] = (v15 + v15)
                    }
                    _ = fi20
                }
                _ = fi19
            }
        case 32767:
            for fi21 := 0; fi21 < 10; fi21++ {
                v16 = append(v16, int(v2))
                v15 += (v15 ^ (v15 + (v15 >> ((879) & 63))))
                fmt.Println(v2, v2)
                _ = fi21
            }
        case -1:
            v22 := GenS1{F0: 195, F1: 305, F2: -2147483648}
            _ = v22
        default:
            var v23 int8 = 100
            _ = v23
        }
    }
    v16 = append(v16, 127)
    fmt.Println(v3)
    var v24 float64 = float64((((float32(((((1.5 / 1.0) - 1e10) + (0.1 / 1e10)) * 1e10)) / 0.1) - (v1 - 0.1)) / 1.5))
    v16 = append(v16, 32767)
    switch "\n\t" {
    case "a":
        v25 := GenS1{F0: 2, F1: -1, F2: 9223372036854775807}
        _ = v25
    case "hello":
        switch 9223372036854775807 {
        case 32767:
            var v26 int8 = 42
            _ = v26
        }
    }
    var v27 int = 32768
    if false {
        var v28 uint = 964
        var v29 int = (v27 / ((int(v15)) | 1))
        v30 := GenS1{F0: v15, F1: 936, F2: (v27 ^ (v27 >> (((v28 * v28)) & 63)))}
        _ = v28
        _ = v29
        _ = v30
    } else {
        v27 = v27
        fmt.Println(v2, v15, v1)
    }
    switch -1 {
    case 128:
        v31 := GenS2{F0: -1, F1: (^(v15))}
        _ = v31
    case -1:
        fmt.Println(v3, v2)
    case 1:
        var v32 int = (-(v27))
        _ = v32
    }
    v33 := GenS2{F0: 2147483647, F1: v15}
    if true {
        switch ("hello" + ("\n\t" + "a")) {
        case "\n\t":
            v34 := GenS2{F0: (-(v3)), F1: (v15 << ((255) & 63))}
            _ = v34
        case "":
            v27 = (v27 - (v27 | v27))
        }
        for fi35 := 0; fi35 < 10; fi35++ {
            fmt.Println(v27, v3, v15)
            v15 = v15
            _ = fi35
        }
        _ = len(v16)
    } else {
        for fi36 := 0; fi36 < 16; fi36++ {
            fmt.Println(v24)
            var v37 uint8 = 139
            fmt.Println(v1)
            _ = fi36
            _ = v37
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v24
    _ = v27
    _ = v3
    _ = v33
    _ = v4
}
