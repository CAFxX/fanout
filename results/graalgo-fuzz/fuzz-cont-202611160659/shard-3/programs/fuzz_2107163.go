package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 uint64 }
type GenS2 struct { F0 uint32; F1 int8; F2 int32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int8; F1 uint }
func genFn0(p0 string, p1 int32, p2 bool) uint {
    v4 := GenS1{F0: 61, F1: 754, F2: 4294967295}
    v5 := GenS2{F0: 703, F1: 127, F2: 997}
    var v6 int64 = -9223372036854775808
    switch ("\n\t" + "a") {
    case "hello":
        var v7 int8 = -128
        _ = v7
    case "a":
        switch p0 {
        case "\n\t":
            p0 = (p0 + ("a" + (p0 + "\n\t")))
        case "hello":
            if (((!false) && ((42 >= 832) && p2)) && ((p2 && p2) && false)) {
                v8 := GenS2{F0: 4294967295, F1: 100, F2: 0}
                v9 := GenS3{F0: 105}
                _ = v8
                _ = v9
            }
        }
    case "":
        var v10 uint64 = 1
        _ = v10
    default:
        if p2 {
            p2 = true
        } else {
            p0 = p0
            v11 := []int{32767, 9223372036854775807}
            _ = v11
        }
    }
    for fi12 := 0; fi12 < 19; fi12++ {
        var v13 string = (p0 + p0)
        var v14 int32 = -1
        if p2 {
            var v15 int64 = (v6 ^ v6)
            if (0.1 <= 3.14159) {
                if p2 {
                    var v16 string = "hello"
                    _ = v16
                } else {
                    p1 = -2147483648
                    p0 = v13
                    v6 -= (v6 % ((v6) | 1))
                }
                v17 := GenS1{F0: 42, F1: 192, F2: 450}
                for fi18 := 0; fi18 < 16; fi18++ {
                    v19 := GenS2{F0: 2, F1: 0, F2: 1}
                    var v20 uint8 = 42
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                _ = v17
            } else {
                var v21 bool = false
                switch fi12 {
                case -9223372036854775808:
                    var v22 int8 = 0
                    _ = v22
                default:
                    fmt.Println(fi12)
                }
                _ = v21
            }
            v23 := GenS4{F0: -24}
            _ = v15
            _ = v23
        }
        _ = fi12
        _ = v13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return 4294967295
}

func main() {
    var v1 uint64 = 0
    var v2 float32 = float32(float64(1e10))
    var v3 int = -357
    v24 := []int{v3, v3, v3}
    switch "\n\t" {
    case "hello":
        fmt.Println(v3, v2, v1)
    case "foo bar":
        var v25 uint8 = 81
        _ = v25
    default:
        var v26 int16 = 1
        _ = v26
    }
    fmt.Println(v1, v1)
    var v27 float64 = 1.5
    v2 -= v2
    var v28 float32 = v2
    if len(v24) > 0 {
        v24[0] = 127
    }
    switch (v3 % ((v3) | 1)) {
    case 4294967295:
        if false {
            v28 = float32(((float64(0.0) + v27) - v27))
            var v29 uint16 = 42
            _ = v29
        }
    case 82:
        for fi30 := 0; fi30 < 3; fi30++ {
            var v31 int64 = 1
            for fi32 := 0; fi32 < 20; fi32++ {
                _ = len(v24)
                _ = fi32
            }
            switch (v3 + 32768) {
            case -212:
                if true {
                    if (3.14159 < v27) {
                        var v33 uint8 = 128
                        _ = v33
                    }
                }
            }
            _ = fi30
            _ = v31
        }
    default:
        _ = len(v24)
    }
    v24 = append(v24, (v3 * (v3 % ((v3) | 1))))
    v2 = v28
    if len(v24) >= 2 {
        v24 = v24[1:len(v24)]
    }
    switch (v3 >> ((2) & 63)) {
    case 32767:
        v2 -= float32((float64(v28) - v27))
    case 4294967295:
        if len(v24) >= 2 {
            v24 = v24[1:len(v24)]
        }
    case 1:
        var v34 int = 0
        _ = v34
    }
    v28 += v2
    v35 := GenS1{F0: -32768, F1: 2, F2: (v1 / ((979) | 1))}
    switch 128 {
    case -221:
        for fi36 := 0; fi36 < 4; fi36++ {
            var v37 float64 = -2.25
            _ = fi36
            _ = v37
        }
    }
    fmt.Println(v28, v27, v28)
    switch (v3 & (v3 + 4294967295)) {
    case 32768:
        switch v3 {
        case 127:
            for fi38 := 0; fi38 < 7; fi38++ {
                v28 = -2.25
                switch "hello" {
                case "a":
                    var v39 uint8 = 128
                    _ = v39
                case "\n\t":
                    var v40 string = "\n\t"
                    _ = v40
                case "":
                    if len(v24) >= 2 {
                        v24 = v24[1:len(v24)]
                    }
                }
                v41 := GenS1{F0: 301, F1: 42, F2: v1}
                _ = fi38
                _ = v41
            }
        case -9223372036854775808:
            v42 := GenS5{F0: 0, F1: 18446744073709551615}
            _ = v42
        }
    }
    var v43 string = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v24
    _ = v27
    _ = v28
    _ = v3
    _ = v35
    _ = v43
}
