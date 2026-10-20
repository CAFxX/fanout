package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 uint64 }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 uint; F1 uint64; F2 uint32 }
func genFn0(p0 int32, p1 int) uint {
    v5 := GenS1{F0: (p1 % (((p1 + (p1 + p1))) | 1)), F1: 42, F2: 1}
    p1 = p1
    _ = p0
    _ = p1
    _ = v5
    return 2
}
func genFn1(p0 uint32, p1 uint16, p2 uint8) int {
    v6 := []int{268, -9223372036854775808, 1}
    v7 := GenS1{F0: 507, F1: 205, F2: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 320
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int16 = 319
    var v3 string = (((("\n\t" + "foo bar") + "hello") + ("\n\t" + ("foo bar" + ("\n\t" + "foo bar")))) + "\n\t")
    var v4 bool = true
    for fi8 := 0; fi8 < 16; fi8++ {
        for fi9 := 0; fi9 < 5; fi9++ {
            v10 := []int{(fi9 % (((fi9 & (^(fi9)))) | 1)), fi8, fi8}
            var v11 uint8 = 2
            _ = fi9
            _ = v10
            _ = v11
        }
        v1["\n\t"] = 65535
        var v12 uint8 = 2
        _ = fi8
        _ = v12
    }
    var v13 bool = true
    var v14 uint8 = 42
    v15 := GenS2{F0: v14, F1: (v14 / ((178) | 1))}
    if (!(!v13)) {
        for fi16 := 0; fi16 < 2; fi16++ {
            var v17 int8 = 0
            fmt.Println(fi16)
            v18 := GenS3{F0: 18446744073709551615, F1: 4294967295, F2: 4294967295}
            _ = fi16
            _ = v17
            _ = v18
        }
        v19 := []int{genFn1(2147483648, 197, 255), 9223372036854775807, 2147483647, -2147483648}
        switch v3 {
        case "\n\t":
            if (!(2 != 18446744073709551615)) {
                var v20 float32 = -2.25
                v21 := v1[v3]
                if len(v19) > 0 {
                    v19[0] = 1
                }
                _ = v20
                _ = v21
            }
        case "foo bar":
            fmt.Println(v13)
        case "hello":
            var v22 float32 = -2.25
            _ = v22
        default:
            v3 = v3
        }
        _ = v19
    }
    v23 := []int{32767, 4294967295, -2147483648}
    if (v4 || v13) {
        v13 = v4
        fmt.Println(v13)
        switch int(v2) {
        case 399:
            v13 = (-21 >= 32767)
        case 9223372036854775807:
            fmt.Println(v3, v3, v2)
        case 465:
            var v24 bool = (127 != 0)
            _ = v24
        default:
            var v25 int16 = v2
            _ = v25
        }
    }
    v3 = "hello"
    v4 = (1 != 378)
    for fi26 := 0; fi26 < 11; fi26++ {
        for fi27 := 0; fi27 < 17; fi27++ {
            if v13 {
                var v28 int64 = 9223372036854775807
                if false {
                    var v29 float64 = 1e10
                    _ = v29
                } else {
                    switch int(v2) {
                    case 0:
                        _ = len(v1)
                    }
                }
                var v30 uint8 = (v14 & 2)
                _ = v28
                _ = v30
            }
            if (!v13) {
                var v31 uint64 = 0
                _ = v31
            }
            v13 = (3.14159 != 3.14159)
            _ = fi27
        }
        _ = fi26
    }
    v32 := GenS1{F0: 42, F1: v14, F2: 4294967295}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v23
    _ = v3
    _ = v32
    _ = v4
}
