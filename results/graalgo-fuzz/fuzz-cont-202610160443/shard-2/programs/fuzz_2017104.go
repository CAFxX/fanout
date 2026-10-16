package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32; F1 uint8; F2 uint16 }
type GenS3 struct { F0 int64; F1 int64; F2 string }
func genFn0(p0 uint32, p1 int16) int {
    for fi7 := 0; fi7 < 19; fi7++ {
        p0 += 712
        _ = fi7
    }
    fmt.Println(p1)
    var v8 float64 = 3.14159
    _ = p0
    _ = p1
    _ = v8
    return -2147483648
}

func main() {
    var v1 uint8 = 42
    v2 := []uint32{4294967295, 2147483648, uint32(v1), 26}
    var v3 float64 = (1.5 * 1e10)
    var v4 int16 = 1
    var v5 float32 = 0.1
    v6 := []string{(("hello" + "a") + "\n\t")}
    if false {
        switch "\n\t" {
        case "":
            var v9 uint16 = uint16(v1)
            _ = v9
        case "foo bar":
            if false {
                fmt.Println(v4, v3, v3)
            } else {
                v5 += v5
                v10 := GenS1{F0: ("\n\t" + "")}
                _ = v10
            }
        default:
            v11 := GenS2{F0: 2147483647, F1: (v1 * (v1 - v1)), F2: 496}
            _ = v11
        }
    } else {
        var v12 int = -1
        for fi13 := 0; fi13 < 1; fi13++ {
            fmt.Println(v12, v5, v1)
            _ = fi13
        }
        for fi14 := 0; fi14 < 15; fi14++ {
            for fi15 := 0; fi15 < 19; fi15++ {
                if true {
                    v4 = 0
                    v12 = (fi15 / ((2147483647) | 1))
                    v3 = float64(((v5 / 1.0) * (v5 - v5)))
                }
                fmt.Println(v1, fi14)
                if false {
                    var v16 uint8 = v1
                    _ = v16
                } else {
                    var v17 uint = 0
                    _ = v17
                }
                _ = fi15
            }
            _ = fi14
        }
        _ = v12
    }
    fmt.Println(v3)
    _ = len(v2)
    if ((false || (!false)) || false) {
        if true {
            var v18 string = ((("a" + ("" + ((("\n\t" + ("hello" + "hello")) + "hello") + "\n\t"))) + "hello") + "hello")
            fmt.Println(v1, v5)
            v18 = ""
            _ = v18
        }
        var v19 float64 = ((v3 / 0.1) / float64(v5))
        _ = v19
    } else {
        v5 -= 1e10
        v20 := GenS2{F0: 65535, F1: 42, F2: 690}
        _ = v20
    }
    v4 = (-(v4))
    switch "a" {
    case "foo bar":
        var v21 int8 = 127
        _ = v21
    case "a":
        v5 = v5
    default:
        for fi22 := 0; fi22 < 9; fi22++ {
            var v23 bool = ((fi22 % ((genFn0(2147483648, 1000)) | 1)) <= (fi22 << ((18446744073709551615) & 63)))
            v4 = (^(v4))
            _ = fi22
            _ = v23
        }
    }
    fmt.Println(v1, v3)
    v24 := GenS3{F0: 0, F1: 974, F2: "\n\t"}
    v6 = append(v6, "\n\t")
    var v25 uint32 = 548
    switch ("a" + "foo bar") {
    case "hello":
        _ = len(v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
