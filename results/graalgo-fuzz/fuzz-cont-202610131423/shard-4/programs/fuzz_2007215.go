package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 int16 }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint64; F1 int64; F2 int8 }
type GenS5 struct { F0 uint8; F1 int16 }
type GenS6 struct { F0 int32; F1 int16 }
func genFn0(p0 uint8) bool {
    var v5 int64 = 9223372036854775807
    switch -904 {
    case 128:
        var v6 string = "\n\t"
        _ = v6
    case -2147483648:
        var v7 float32 = ((0.1 - 1e10) / 1.5)
        _ = v7
    case -9223372036854775808:
        v8 := GenS1{F0: 42, F1: 32767, F2: 42}
        _ = v8
    }
    var v9 int32 = -878
    var v10 uint16 = 65535
    v11 := GenS1{F0: uint32(p0), F1: (v10 << ((1) & 63)), F2: -132}
    _ = p0
    _ = v10
    _ = v11
    _ = v5
    _ = v9
    return true
}

func main() {
    var v1 float64 = ((0.0 + 3.14159) + 0.0)
    var v2 int32 = 519
    var v3 uint32 = 1
    v4 := []string{("" + ""), "\n\t"}
    for fi12 := 0; fi12 < 13; fi12++ {
        v13 := GenS2{F0: 2, F1: 2}
        v3 -= v3
        _ = fi12
        _ = v13
    }
    v2 = -339
    var v14 string = ((((((("hello" + "") + "\n\t") + "hello") + (("a" + ("" + "a")) + (("a" + "hello") + "foo bar"))) + "a") + "") + "")
    v2 -= v2
    switch 32767 {
    case 32768:
        v15 := GenS3{F0: 127}
        _ = v15
    case -2147483648:
        for fi16 := 0; fi16 < 17; fi16++ {
            v14 = v14
            _ = fi16
        }
    default:
        var v17 int = 32768
        _ = v17
    }
    for fi18 := 0; fi18 < 4; fi18++ {
        v19 := GenS4{F0: 42, F1: 42, F2: -1}
        _ = fi18
        _ = v19
    }
    var v20 string = v14
    var v21 int16 = 42
    for fi22 := 0; fi22 < 9; fi22++ {
        v14 = v14
        var v23 string = (v20 + ("a" + ((v20 + "foo bar") + v14)))
        _ = fi22
        _ = v23
    }
    v24 := GenS5{F0: 42, F1: (-(v21))}
    switch -920 {
    case -1:
        var v25 float64 = 0.1
        _ = v25
    case 2147483647:
        switch v14 {
        case "a":
            if len(v4) > 0 {
                v4[0] = v14
            }
        case "foo bar":
            var v26 int32 = (v2 & v2)
            _ = v26
        }
    }
    v27 := GenS5{F0: 127, F1: 156}
    switch "a" {
    case "hello":
        if true {
            var v28 uint16 = 1
            _ = v28
        } else {
            var v29 int16 = 0
            _ = v29
        }
    }
    switch int(v21) {
    case -203:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case -9223372036854775808:
        for fi30 := 0; fi30 < 3; fi30++ {
            _ = len(v4)
            _ = fi30
        }
    case -94:
        v14 = (v20 + "hello")
    default:
        var v31 int16 = 0
        _ = v31
    }
    v3 += v3
    switch v14 {
    case "a":
        v32 := GenS6{F0: v2, F1: v21}
        _ = v32
    case "hello":
        var v33 int16 = (v21 + (v21 | v21))
        _ = v33
    case "\n\t":
        fmt.Println(v3, v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v27
    _ = v3
    _ = v4
}
