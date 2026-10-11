package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint; F1 bool; F2 int16 }
type GenS3 struct { F0 int32; F1 int8 }
func genFn0(p0 uint) uint {
    v7 := GenS1{F0: -9223372036854775808}
    var v8 int = -713
    v8 += -2147483648
    var v9 uint8 = 1
    switch "\n\t" {
    case "\n\t":
        v8 += -92
    }
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return (p0 ^ (^(p0)))
}

func main() {
    v1 := make(map[int]bool)
    var v2 int = -1
    var v3 uint8 = 1
    var v4 float32 = 1.5
    var v5 int32 = -2147483648
    var v6 int = (v2 - (v2 / (((v2 * (v2 * (v2 & 492)))) | 1)))
    for fi10 := 0; fi10 < 8; fi10++ {
        var v11 float64 = float64((0.1 / v4))
        v12 := GenS2{F0: 397, F1: true, F2: -1}
        fmt.Println(v6)
        _ = fi10
        _ = v11
        _ = v12
    }
    v3 -= v3
    var v13 int8 = 1
    switch "\n\t" {
    case "hello":
        v4 = v4
    default:
        switch ("hello" + "") {
        case "":
            var v14 uint8 = v3
            _ = v14
        }
    }
    switch v6 {
    case 881:
        fmt.Println(v6, v6)
    }
    v5 += (v5 ^ v5)
    var v15 bool = true
    v16 := []int{(v2 + (v6 | 702)), 2147483647, (v2 | v2), 149}
    if v15 {
        if (!(v5 <= (v5 >> ((18446744073709551615) & 63)))) {
            switch (v6 >> ((65535) & 63)) {
            case 128:
                var v17 float64 = 3.14159
                _ = v17
            case 4294967295:
                v3 = (v3 % ((v3) | 1))
            case 9223372036854775807:
                v6 = v2
            default:
                var v18 uint = 65535
                _ = v18
            }
        }
        if v15 {
            fmt.Println(v15, v5, v2)
            fmt.Println(v5)
            var v19 uint32 = 855
            _ = v19
        } else {
            if len(v16) > 0 {
                v16[0] = v6
            }
        }
    } else {
        for fi20 := 0; fi20 < 9; fi20++ {
            v21 := GenS2{F0: 564, F1: false, F2: -32768}
            _ = fi20
            _ = v21
        }
    }
    if (true && false) {
        for fi22 := 0; fi22 < 18; fi22++ {
            switch "a" {
            case "\n\t":
                var v23 float64 = -2.25
                _ = v23
            }
            var v24 int8 = (v13 + v13)
            v3 += (v3 + 127)
            _ = fi22
            _ = v24
        }
        var v25 float64 = (3.14159 / 1.0)
        for fi26 := 0; fi26 < 3; fi26++ {
            v4 = float32(0.1)
            v4 -= (v4 + 0.1)
            var v27 string = ""
            _ = fi26
            _ = v27
        }
        _ = v25
    }
    for fi28 := 0; fi28 < 16; fi28++ {
        _ = len(v1)
        if (!true) {
            fmt.Println(v6, v4, v4)
            fmt.Println(fi28)
            for fi29 := 0; fi29 < 7; fi29++ {
                fmt.Println(fi28, v3)
                v6 = v6
                _ = fi29
            }
        }
        v16 = append(v16, (fi28 + 127))
        _ = fi28
    }
    var v30 uint16 = 118
    if len(v16) > 0 {
        v16[0] = 582
    }
    if v15 {
        fmt.Println(v13, v30)
        switch 730 {
        case 42:
            if (v3 < v3) {
                if v15 {
                    var v31 int64 = 9223372036854775807
                    _ = v31
                } else {
                    fmt.Println(v3, v13)
                }
                fmt.Println(v15, v4)
            }
        default:
            var v32 int64 = -1
            _ = v32
        }
        var v33 float64 = 1e10
        _ = v33
    }
    v13 = v13
    var v34 uint16 = (v30 | (^(v30)))
    for fi35 := 0; fi35 < 6; fi35++ {
        var v36 uint8 = v3
        v37 := GenS3{F0: -667, F1: v13}
        if len(v16) > 0 {
            v16[0] = (-(fi35))
        }
        _ = fi35
        _ = v36
        _ = v37
    }
    v6 -= v6
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v30
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
