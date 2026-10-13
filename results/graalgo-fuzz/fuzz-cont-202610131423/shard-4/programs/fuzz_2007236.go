package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
func genFn0() uint8 {
    var v5 float32 = float32(1.5)
    v6 := GenS1{F0: 0, F1: 264}
    v5 = float32((((-2.25 * -2.25) * 0.1) / 3.14159))
    var v7 int8 = 42
    switch -658 {
    case 384:
        if (!true) {
            var v8 float64 = float64(v5)
            _ = v8
        }
    case 4294967295:
        v7 = v7
    case 127:
        fmt.Println(v7, v7)
    default:
        for fi9 := 0; fi9 < 14; fi9++ {
            var v10 int32 = 0
            for fi11 := 0; fi11 < 4; fi11++ {
                v12 := GenS1{F0: 42, F1: 32767}
                _ = fi11
                _ = v12
            }
            _ = fi9
            _ = v10
        }
    }
    _ = v5
    _ = v6
    _ = v7
    return 10
}
func genFn1(p0 int32) int64 {
    p0 = (p0 << ((2) & 63))
    p0 += p0
    _ = p0
    return 1
}

func main() {
    var v1 int64 = -482
    var v2 int32 = -2147483648
    v3 := make(map[string]string)
    var v4 bool = (370 <= 18446744073709551615)
    var v13 int8 = 42
    v4 = v4
    for fi14 := 0; fi14 < 13; fi14++ {
        for fi15 := 0; fi15 < 15; fi15++ {
            var v16 uint16 = 831
            _ = fi15
            _ = v16
        }
        _ = fi14
    }
    var v17 string = "\n\t"
    var v18 int64 = -1
    for fi19 := 0; fi19 < 13; fi19++ {
        switch v17 {
        case "hello":
            switch (^(fi19)) {
            case 32768:
                var v20 uint16 = 65535
                _ = v20
            case 1:
                var v21 uint = 259
                _ = v21
            }
        case "\n\t":
            var v22 bool = (v17 < (("foo bar" + v17) + "\n\t"))
            _ = v22
        default:
            var v23 int32 = v2
            _ = v23
        }
        var v24 uint32 = 4294967295
        _ = fi19
        _ = v24
    }
    var v25 bool = (v4 || ((((3.14159 + ((3.14159 / float64(0.1)) / (((float64(((-2.25 * -2.25) - 0.0)) - float64((float32(1e10) / 1e10))) + (float64(0.1) + 3.14159)) + (1.5 - float64(float32((((3.14159 - (3.14159 * (0.0 + (float64(float32(0.0)) * float64(-2.25))))) * 1e10) + (3.14159 * (-2.25 - (3.14159 + (-2.25 / 1.0))))))))))) - 0.0) * float64((float32(0.0) / -2.25))) != 1e10))
    for fi26 := 0; fi26 < 8; fi26++ {
        var v27 int16 = 0
        fmt.Println(v2, v17, v2)
        var v28 float32 = 3.14159
        _ = fi26
        _ = v27
        _ = v28
    }
    var v29 int16 = -1
    v30 := []int{-2147483648}
    var v31 int8 = (v13 << ((0) & 63))
    for fi32 := 0; fi32 < 6; fi32++ {
        v18 = v18
        v2 -= v2
        switch int(v29) {
        case -9223372036854775808:
            v33 := GenS1{F0: 42, F1: 50}
            _ = v33
        case 128:
            var v34 int32 = (v2 % ((-1) | 1))
            _ = v34
        default:
            var v35 int32 = -1
            _ = v35
        }
        _ = fi32
    }
    if (v25 || (!v25)) {
        fmt.Println(v1, v1)
        v29 = 1000
    }
    if ((1e10 >= 1.5) || (-128 != (v13 + v31))) {
        v36 := GenS1{F0: 1, F1: 1}
        _ = v36
    }
    if v4 {
        _ = len(v30)
        v1 -= v1
        if (!v25) {
            v2 -= v2
        }
    }
    var v37 int = 0
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v31:", v31)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v37
    _ = v4
}
