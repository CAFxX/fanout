package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8; F1 int16; F2 int32 }
type GenS3 struct { F0 bool; F1 uint64; F2 bool }
type GenS4 struct { F0 int16; F1 int32; F2 uint64 }
func genFn0(p0 string) uint32 {
    v7 := GenS1{F0: 889}
    for fi8 := 0; fi8 < 17; fi8++ {
        if true {
            v9 := GenS1{F0: 2147483648}
            for fi10 := 0; fi10 < 13; fi10++ {
                p0 = p0
                fmt.Println(fi10, fi10, fi10)
                _ = fi10
            }
            switch "" {
            case "foo bar":
                p0 = p0
            case "a":
                var v11 float64 = 1e10
                _ = v11
            case "\n\t":
                v12 := []int{-935}
                _ = v12
            default:
                var v13 bool = false
                _ = v13
            }
            _ = v9
        } else {
            v14 := []int{-9223372036854775808, fi8, fi8, fi8}
            _ = v14
        }
        fmt.Println(p0)
        _ = fi8
    }
    fmt.Println(p0)
    p0 = (("a" + p0) + p0)
    _ = p0
    _ = v7
    return 2147483647
}

func main() {
    var v1 int16 = 0
    v2 := []int{2147483647, 32768, -440}
    var v3 string = ("" + "hello")
    var v4 uint = 18446744073709551615
    var v5 uint32 = 417
    var v6 int32 = -1
    v2 = append(v2, 4294967295)
    switch "foo bar" {
    case "\n\t":
        v3 = (((v3 + v3) + v3) + v3)
    case "a":
        var v15 int32 = v6
        _ = v15
    default:
        v16 := GenS1{F0: v5}
        _ = v16
    }
    var v17 bool = false
    for fi18 := 0; fi18 < 16; fi18++ {
        for fi19 := 0; fi19 < 11; fi19++ {
            v6 = v6
            switch fi19 {
            case 2147483647:
                var v20 uint8 = 128
                _ = v20
            default:
                v21 := GenS2{F0: 63, F1: (v1 | (v1 >> ((v4) & 63))), F2: (v6 - v6)}
                v22 := GenS2{F0: 1, F1: v1, F2: v6}
                var v23 float64 = (1e10 - (3.14159 + 3.14159))
                if len(v2) > 0 {
                    v2[0] = fi19
                }
                _ = v21
                _ = v22
                _ = v23
            }
            v3 = "foo bar"
            _ = fi19
        }
        if true {
            var v24 int16 = (v1 % ((v1) | 1))
            _ = v24
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            if v17 {
                v3 = (v3 + v3)
            } else {
                fmt.Println(v17)
            }
        }
        _ = fi18
    }
    v4 -= v4
    if v17 {
        v25 := GenS1{F0: v5}
        fmt.Println(v1, v4)
        _ = v25
    }
    var v26 bool = (!v17)
    fmt.Println(v3, v1)
    var v27 uint = v4
    switch -9223372036854775808 {
    case -2147483648:
        var v28 bool = true
        _ = v28
    case -1:
        v29 := GenS3{F0: (!true), F1: uint64(v5), F2: v26}
        _ = v29
    }
    v3 = (v3 + ("hello" + v3))
    v30 := GenS4{F0: 886, F1: -2147483648, F2: 0}
    for fi31 := 0; fi31 < 3; fi31++ {
        var v32 float64 = ((((0.0 + (float64(1.5) / 1.0)) / 1e10) * float64(0.0)) + -2.25)
        if false {
            if len(v2) > 0 {
                v2[0] = -635
            }
            v32 -= (v32 - (float64(0.0) - 0.0))
        } else {
            var v33 int16 = v1
            fmt.Println(v33, v33, v5)
            fmt.Println(v27, v27, v5)
            _ = v33
        }
        _ = fi31
        _ = v32
    }
    switch ("\n\t" + v3) {
    case "a":
        var v34 int16 = 1
        _ = v34
    default:
        var v35 uint16 = 213
        _ = v35
    }
    v36 := GenS2{F0: 108, F1: (v1 + 32767), F2: v6}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
