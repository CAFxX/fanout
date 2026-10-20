package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 uint16 }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int; F1 uint64 }

func main() {
    var v1 uint8 = 139
    var v2 float64 = (1e10 * -2.25)
    v3 := make(map[string]int)
    v4 := []int{-1, -169, 603}
    fmt.Println(v1)
    var v5 float32 = 3.14159
    v2 += v2
    for fi6 := 0; fi6 < 4; fi6++ {
        var v7 float64 = float64(-2.25)
        for fi8 := 0; fi8 < 20; fi8++ {
            var v9 float64 = v2
            v10 := GenS1{F0: 32767, F1: 528, F2: 0}
            _ = fi8
            _ = v10
            _ = v9
        }
        v5 = v5
        _ = fi6
        _ = v7
    }
    if false {
        if false {
            for fi11 := 0; fi11 < 7; fi11++ {
                v4 = append(v4, (fi11 / (((fi11 / ((1) | 1))) | 1)))
                _ = fi11
            }
            var v12 uint8 = v1
            if (678 <= 827) {
                v1 = (v1 / (((v12 & v1)) | 1))
                fmt.Println(v1, v5, v12)
            } else {
                for fi13 := 0; fi13 < 19; fi13++ {
                    v14 := GenS1{F0: 2, F1: 0, F2: 42}
                    var v15 uint8 = 226
                    v16 := GenS1{F0: 32767, F1: 4294967295, F2: uint16(v15)}
                    _ = fi13
                    _ = v14
                    _ = v15
                    _ = v16
                }
            }
            _ = v12
        }
        var v17 int32 = -1
        _ = v17
    } else {
        v3[""] = 9223372036854775807
        switch ("a" + ("a" + ("hello" + ("\n\t" + "foo bar")))) {
        case "":
            v18 := GenS1{F0: 255, F1: 1, F2: 443}
            _ = v18
        }
        for fi19 := 0; fi19 < 3; fi19++ {
            v1 = (v1 | v1)
            v20 := GenS1{F0: 1, F1: 65535, F2: 0}
            fmt.Println(v5)
            _ = fi19
            _ = v20
        }
    }
    for fi21 := 0; fi21 < 15; fi21++ {
        v5 -= ((0.0 + (1e10 * v5)) * ((v5 * v5) - (v5 / v5)))
        _ = fi21
    }
    switch "" {
    case "foo bar":
        var v22 uint = 42
        _ = v22
    case "\n\t":
        v23 := GenS2{F0: -32768, F1: v1}
        _ = v23
    }
    _ = len(v4)
    delete(v3, "foo bar")
    var v24 int8 = -28
    switch "hello" {
    case "a":
        v1 -= 27
    default:
        if ((0 > uint32(v1)) && (false || false)) {
            v1 = (v1 - (v1 & 137))
        } else {
            var v25 int8 = v24
            v5 += v5
            switch -552 {
            case 127:
                for fi26 := 0; fi26 < 17; fi26++ {
                    for fi27 := 0; fi27 < 15; fi27++ {
                        var v28 int8 = 114
                        _ = fi27
                        _ = v28
                    }
                    _ = fi26
                }
            case 128:
                delete(v3, ("hello" + ""))
            }
            _ = v25
        }
    }
    switch ("" + (("\n\t" + ("" + "foo bar")) + "hello")) {
    case "\n\t":
        var v29 uint64 = 1
        _ = v29
    case "foo bar":
        v1 = 2
    default:
        if true {
            v30 := GenS1{F0: 2, F1: 65535, F2: 65535}
            var v31 uint64 = 2
            for fi32 := 0; fi32 < 1; fi32++ {
                v24 = -1
                v33 := GenS3{F0: -2147483648}
                v31 = v31
                _ = fi32
                _ = v33
            }
            _ = v30
            _ = v31
        }
    }
    _ = len(v3)
    v34 := GenS4{F0: 128, F1: 2}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}
