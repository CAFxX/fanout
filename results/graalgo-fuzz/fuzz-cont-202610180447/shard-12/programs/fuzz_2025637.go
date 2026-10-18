package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32; F1 uint64; F2 uint64 }
type GenS3 struct { F0 uint64; F1 uint8 }
type GenS4 struct { F0 uint16; F1 uint8 }
func genFn0(p0 int16, p1 uint) uint16 {
    p0 = p0
    for fi4 := 0; fi4 < 8; fi4++ {
        for fi5 := 0; fi5 < 14; fi5++ {
            switch (fi4 - -586) {
            case 32768:
                var v6 uint = 255
                _ = v6
            default:
                p0 += -466
            }
            v7 := []int{fi4}
            _ = fi5
            _ = v7
        }
        _ = fi4
    }
    v8 := []int{-950, -1, 9223372036854775807}
    _ = p0
    _ = p1
    _ = v8
    return 42
}
func genFn1(p0 int8) int64 {
    if false {
        p0 = -1
        var v9 uint8 = 128
        if (0 == 0) {
            switch 42 {
            case 648:
                p0 -= p0
            case 127:
                v10 := []int{}
                _ = v10
            }
            var v11 bool = true
            p0 = p0
            _ = v11
        }
        _ = v9
    } else {
        p0 += p0
    }
    p0 -= p0
    var v12 uint64 = 4294967295
    v13 := GenS1{F0: 2147483647}
    p0 += (p0 & p0)
    _ = p0
    _ = v12
    _ = v13
    return int64(v12)
}

func main() {
    v1 := []int{2147483647, 4294967295, 9223372036854775807, -9223372036854775808}
    var v2 int = 222
    var v3 uint16 = 1
    for fi14 := 0; fi14 < 9; fi14++ {
        if (((((float64(0.0) * (1e10 / -2.25)) * 0.1) / ((float64(float32(1e10)) + 1e10) / float64((-2.25 + float32((((-2.25 + (3.14159 / 1.0)) - 1.5) * (1e10 / (1e10 / 1e10)))))))) * (1e10 / 0.1)) >= (1e10 - float64((0.1 - 1e10)))) {
            v1 = append(v1, -8)
            v15 := GenS1{F0: -2147483648}
            _ = v15
        }
        if (2147483647 <= 886) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v16 int32 = 2147483647
            _ = v16
        }
        if true {
            _ = len(v1)
        }
        _ = fi14
    }
    var v17 uint8 = 0
    v3 = (v3 * v3)
    fmt.Println(v17, v17)
    switch "a" {
    case "\n\t":
        switch (v2 ^ (v2 + (v2 / ((v2) | 1)))) {
        case 0:
            var v18 int16 = 1
            _ = v18
        case -222:
            v19 := GenS2{F0: 0, F1: 811, F2: 877}
            _ = v19
        }
    case "foo bar":
        var v20 uint16 = (v3 * 65535)
        _ = v20
    case "a":
        var v21 int32 = 643
        _ = v21
    default:
        if true {
            v22 := GenS2{F0: -2147483648, F1: 137, F2: 2}
            _ = v22
        }
    }
    var v23 bool = (1000 >= -1)
    if v23 {
        fmt.Println(v23)
        switch ("\n\t" + "a") {
        case "a":
            var v24 uint32 = 4294967295
            _ = v24
        case "\n\t":
            var v25 int16 = int16(v17)
            _ = v25
        case "foo bar":
            for fi26 := 0; fi26 < 13; fi26++ {
                var v27 uint8 = v17
                _ = fi26
                _ = v27
            }
        }
    }
    v23 = (!v23)
    v28 := GenS3{F0: 18446744073709551615, F1: v17}
    switch -731 {
    case -63:
        switch (v2 / (((v2 << ((18446744073709551615) & 63))) | 1)) {
        case -2147483648:
            var v29 bool = v23
            _ = v29
        }
    }
    var v30 float64 = 3.14159
    if false {
        fmt.Println(v23, v23, v2)
    } else {
        var v31 int16 = 0
        var v32 uint8 = 128
        _ = v31
        _ = v32
    }
    _ = len(v1)
    v33 := GenS4{F0: (v3 + (v3 + (v3 ^ 65535))), F1: v17}
    _ = len(v1)
    var v34 string = (("\n\t" + ("a" + "hello")) + "foo bar")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v17
    _ = v2
    _ = v23
    _ = v28
    _ = v3
    _ = v30
    _ = v33
    _ = v34
}
