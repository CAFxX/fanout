package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 uint16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool; F1 int16; F2 string }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int; F1 int64 }

func main() {
    var v1 bool = (-1 <= -600)
    var v2 uint64 = 4294967295
    v3 := make(map[int]uint32)
    var v4 float64 = ((1.5 - (-2.25 - float64(((3.14159 - -2.25) / 3.14159)))) * (0.0 * (1e10 / (1e10 + -2.25))))
    v5 := make(map[uint32]int)
    for fi6 := 0; fi6 < 9; fi6++ {
        var v7 uint64 = 769
        _ = fi6
        _ = v7
    }
    if v1 {
        var v8 uint = 255
        v2 = (v2 & v2)
        _ = v8
    }
    v9 := GenS1{F0: uint16(v2), F1: 100, F2: 186}
    for fi10 := 0; fi10 < 17; fi10++ {
        v2 -= (v2 + 18446744073709551615)
        _ = fi10
    }
    if v1 {
        v11 := GenS1{F0: 2, F1: 10, F2: 65535}
        var v12 int64 = -1
        v13 := []int{-9223372036854775808, 386, int(v12)}
        _ = v11
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 8; fi14++ {
        if v1 {
            fmt.Println(v2, fi14)
            var v15 uint = 296
            _ = v15
        }
        v1 = (!v1)
        v16 := GenS2{F0: v2}
        _ = fi14
        _ = v16
    }
    var v17 string = "\n\t"
    switch ("hello" + v17) {
    case "foo bar":
        if (true && (((v1 || false) && (!v1)) || (-9223372036854775808 != -9223372036854775808))) {
            for fi18 := 0; fi18 < 8; fi18++ {
                if ((v2 << ((215) & 63)) != (v2 ^ (v2 << ((uint(v2)) & 63)))) {
                    if v1 {
                        fmt.Println(v4)
                        v4 = (v4 - (3.14159 * v4))
                        var v19 int = 128
                        _ = v19
                    }
                }
                if v1 {
                    v17 = v17
                    v4 += (((v4 / v4) / (v4 * v4)) - v4)
                } else {
                    v20 := GenS3{F0: (4294967295 < uint(v2)), F1: 32767, F2: v17}
                    for fi21 := 0; fi21 < 12; fi21++ {
                        var v22 float32 = (0.0 - -2.25)
                        v23 := GenS4{F0: 0}
                        _ = fi21
                        _ = v22
                        _ = v23
                    }
                    v24 := GenS5{F0: (fi18 + -1), F1: -901}
                    _ = v20
                    _ = v24
                }
                _ = fi18
            }
            v25 := []int{4294967295, -9223372036854775808, 665}
            if ((!false) && false) {
                v4 += -2.25
            }
            _ = v25
        }
    case "a":
        var v26 int16 = int16(v2)
        _ = v26
    case "":
        fmt.Println(v2)
    default:
        var v27 float32 = 0.0
        _ = v27
    }
    v28 := []int{32767, 32767}
    var v29 string = v17
    var v30 int = -2147483648
    if (65535 <= 2) {
        var v31 int64 = -1
        var v32 uint16 = 2
        _ = v31
        _ = v32
    } else {
        var v33 uint8 = 126
        fmt.Println(v2)
        if len(v28) >= 2 {
            v28 = v28[1:len(v28)]
        }
        _ = v33
    }
    if len(v28) > 0 {
        v28[0] = v30
    }
    var v34 uint32 = 2147483648
    for fi35 := 0; fi35 < 10; fi35++ {
        fmt.Println(v30, v29, v34)
        _ = fi35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v34
    _ = v4
    _ = v5
    _ = v9
}
