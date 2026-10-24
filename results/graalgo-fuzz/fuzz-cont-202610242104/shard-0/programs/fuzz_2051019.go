package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 string }
type GenS2 struct { F0 uint16; F1 int16 }
type GenS3 struct { F0 int16; F1 int8; F2 string }
type GenS4 struct { F0 string }
func genFn0(p0 uint8, p1 string) uint32 {
    if ("\n\t" > "foo bar") {
        var v5 uint16 = 32767
        var v6 int8 = 42
        _ = v5
        _ = v6
    } else {
        fmt.Println(p0, p1)
        var v7 float64 = 0.0
        _ = v7
    }
    fmt.Println(p0)
    for fi8 := 0; fi8 < 15; fi8++ {
        var v9 int32 = 1
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    return 619
}
func genFn1(p0 int64, p1 uint8) string {
    v10 := []int{270, 1}
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    _ = v10
    return ""
}

func main() {
    var v1 float32 = 0.1
    var v2 uint16 = 65535
    var v3 bool = (0.1 < (0.0 * 0.0))
    v4 := []uint32{629, 710, 1, 1}
    v11 := GenS1{F0: 0, F1: -2147483648, F2: "\n\t"}
    v12 := GenS1{F0: 2, F1: 42, F2: (("" + "foo bar") + "\n\t")}
    for fi13 := 0; fi13 < 2; fi13++ {
        var v14 bool = v3
        v2 -= v2
        _ = fi13
        _ = v14
    }
    var v15 float64 = (-2.25 / (1e10 / 3.14159))
    for fi16 := 0; fi16 < 12; fi16++ {
        v3 = false
        fmt.Println(v1, v3)
        _ = fi16
    }
    v4 = append(v4, uint32(v2))
    v17 := GenS1{F0: 65535, F1: 128, F2: ("\n\t" + "")}
    if v3 {
        v18 := GenS2{F0: v2, F1: 1000}
        v19 := GenS3{F0: -482, F1: 1, F2: "\n\t"}
        _ = v18
        _ = v19
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        fmt.Println(v15, v3, v15)
        var v20 int8 = 42
        _ = v20
    }
    var v21 int = 32768
    if (0 == 127) {
        v21 -= 9223372036854775807
        for fi22 := 0; fi22 < 8; fi22++ {
            var v23 uint = 962
            _ = fi22
            _ = v23
        }
        for fi24 := 0; fi24 < 19; fi24++ {
            v25 := GenS4{F0: ("\n\t" + ("a" + ""))}
            v15 = ((v15 * float64(v1)) / v15)
            fmt.Println(v15, fi24, v2)
            _ = fi24
            _ = v25
        }
    }
    var v26 uint64 = 42
    if v3 {
        fmt.Println(v26, v3, v21)
        v27 := GenS2{F0: (v2 % ((v2) | 1)), F1: 1000}
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v17
    _ = v2
    _ = v21
    _ = v26
    _ = v3
    _ = v4
}
