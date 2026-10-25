package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 uint; F2 uint }
type GenS3 struct { F0 uint16; F1 uint8 }
func genFn0() uint {
    v4 := GenS1{F0: 16}
    v5 := GenS1{F0: 0}
    var v6 int32 = -673
    v7 := GenS2{F0: 18446744073709551615, F1: 379, F2: 1}
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 0
}

func main() {
    var v1 uint64 = 517
    var v2 uint32 = 4294967295
    v3 := make(map[string]uint32)
    if ((0 != 0) || false) {
        for fi8 := 0; fi8 < 16; fi8++ {
            v1 += 980
            _ = fi8
        }
    }
    var v9 uint8 = 2
    var v10 int64 = 0
    switch "hello" {
    case "hello":
        for fi11 := 0; fi11 < 9; fi11++ {
            var v12 string = ("hello" + (("" + ("a" + "hello")) + ""))
            _ = fi11
            _ = v12
        }
    case "a":
        if false {
            if true {
                var v13 uint32 = 2
                var v14 uint16 = uint16(v9)
                _ = len(v3)
                _ = v13
                _ = v14
            }
        } else {
            var v15 float32 = 1e10
            for fi16 := 0; fi16 < 4; fi16++ {
                var v17 float32 = v15
                var v18 uint32 = 42
                if ((float64(((1.5 / v17) - (0.0 + v15))) != 1e10) && (1 <= -128)) {
                    v19 := GenS3{F0: 496, F1: (^(v9))}
                    var v20 uint = 436
                    _ = v19
                    _ = v20
                } else {
                    var v21 int16 = 0
                    for fi22 := 0; fi22 < 18; fi22++ {
                        v3[("\n\t" + "\n\t")] = v2
                        v23 := []int{}
                        fmt.Println(v1, v9, fi16)
                        _ = fi22
                        _ = v23
                    }
                    _ = v21
                }
                _ = fi16
                _ = v17
                _ = v18
            }
            var v24 int32 = -2147483648
            _ = v15
            _ = v24
        }
    default:
        if (!(-910 > -32768)) {
            _ = len(v3)
        } else {
            v25 := []int{-1, -286}
            var v26 uint = 65535
            v27 := GenS2{F0: (v1 << (((v26 % ((v26) | 1))) & 63)), F1: v26, F2: 2}
            _ = v25
            _ = v26
            _ = v27
        }
    }
    fmt.Println(v1, v2)
    switch 1 {
    case 32768:
        switch ((("hello" + "foo bar") + ("a" + "\n\t")) + ("\n\t" + (("foo bar" + (("\n\t" + "") + "\n\t")) + ("a" + "hello")))) {
        case "a":
            v28 := v3[""]
            _ = v28
        case "hello":
            var v29 uint8 = v9
            _ = v29
        }
    default:
        switch "a" {
        case "\n\t":
            v10 += (^(v10))
        case "a":
            var v30 float64 = 1e10
            _ = v30
        default:
            v31 := []int{-1, int(v9), -9223372036854775808, int(v10)}
            _ = v31
        }
    }
    v32 := GenS3{F0: 50, F1: 125}
    v9 = (v9 % ((2) | 1))
    var v33 uint32 = (v2 - v2)
    switch "" {
    case "\n\t":
        var v34 uint16 = 65535
        _ = v34
    case "hello":
        var v35 int = 0
        _ = v35
    case "a":
        switch "a" {
        case "\n\t":
            v1 = v1
        case "a":
            v10 -= v10
        case "":
            v2 += 65535
        }
    }
    fmt.Println(v1, v33)
    _ = len(v3)
    v10 += v10
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v9
}
