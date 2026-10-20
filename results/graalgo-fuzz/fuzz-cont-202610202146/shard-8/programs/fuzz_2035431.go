package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 uint32; F2 uint8 }
type GenS3 struct { F0 uint8; F1 int16; F2 uint }
type GenS4 struct { F0 uint64; F1 int64; F2 int32 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int32 }
func genFn0(p0 uint16) uint16 {
    p0 = (p0 * p0)
    var v4 string = "a"
    if (false || (false && (0.1 > 1.5))) {
        var v5 float64 = (-2.25 - (float64(1e10) / (((3.14159 - 3.14159) + 1.5) * (((((3.14159 + 3.14159) + 1e10) + (0.0 / 1.5)) + 3.14159) + (0.0 * 0.1)))))
        var v6 uint = 1
        v7 := []int{-9223372036854775808, 2147483647}
        _ = v5
        _ = v6
        _ = v7
    } else {
        var v8 int64 = 42
        v8 -= 9223372036854775807
        v8 -= v8
        _ = v8
    }
    _ = p0
    _ = v4
    return 0
}

func main() {
    var v1 float64 = float64(-2.25)
    v2 := make(map[string]string)
    v3 := make(map[int]bool)
    var v9 float32 = 3.14159
    if true {
        delete(v3, 64)
        var v10 uint64 = 4294967295
        fmt.Println(v9, v9, v1)
        _ = v10
    } else {
        v11 := GenS1{F0: 11}
        var v12 int = 32768
        _ = v11
        _ = v12
    }
    switch -9223372036854775808 {
    case 32767:
        switch ("foo bar" + ("foo bar" + "foo bar")) {
        case "":
            for fi13 := 0; fi13 < 15; fi13++ {
                v14 := GenS2{F0: 42, F1: 65535, F2: 255}
                v1 = (v1 - v1)
                var v15 int = (-(fi13))
                _ = fi13
                _ = v14
                _ = v15
            }
        case "foo bar":
            var v16 float32 = 1.5
            _ = v16
        default:
            var v17 uint = 42
            _ = v17
        }
    }
    var v18 uint8 = 128
    v1 -= float64(v9)
    v19 := GenS3{F0: 128, F1: 1000, F2: 728}
    if (uint32(v18) > 2147483647) {
        v18 = 1
    } else {
        var v20 uint8 = 42
        fmt.Println(v20)
        _ = v20
    }
    var v21 int = 0
    v22 := GenS2{F0: 2, F1: 2147483647, F2: 128}
    v23 := GenS1{F0: 4294967295}
    var v24 uint32 = 65535
    fmt.Println(v24)
    fmt.Println(v24, v9, v1)
    switch "" {
    case "\n\t":
        _ = len(v3)
    case "a":
        var v25 int8 = -128
        _ = v25
    case "foo bar":
        v26 := GenS4{F0: 2, F1: 1, F2: -1}
        _ = v26
    }
    if true {
        v27 := []int{v21}
        switch (v21 ^ -594) {
        case -2147483648:
            switch (v21 + v21) {
            case 4294967295:
                v27 = append(v27, v21)
            }
        case -886:
            for fi28 := 0; fi28 < 6; fi28++ {
                var v29 string = "\n\t"
                v30 := GenS5{F0: 18446744073709551615}
                _ = fi28
                _ = v29
                _ = v30
            }
        }
        _ = v27
    } else {
        for fi31 := 0; fi31 < 5; fi31++ {
            v32 := GenS6{F0: 1}
            fmt.Println(fi31, v18)
            for fi33 := 0; fi33 < 1; fi33++ {
                v24 = (v24 * 2147483648)
                v24 = v24
                v21 += v21
                _ = fi33
            }
            _ = fi31
            _ = v32
        }
        var v34 uint = 930
        _ = v34
    }
    fmt.Println(v9, v21, v18)
    for fi35 := 0; fi35 < 9; fi35++ {
        switch -275 {
        case 1:
            var v36 uint64 = 0
            _ = v36
        }
        v24 -= 104
        _ = fi35
    }
    v1 -= v1
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v9
}
