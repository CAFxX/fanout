package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint16; F1 string }
func genFn0(p0 int) int16 {
    p0 += 0
    if false {
        var v4 float64 = ((0.0 - -2.25) - 1.5)
        _ = v4
    } else {
        var v5 int32 = 903
        _ = v5
    }
    switch ("hello" + ("foo bar" + "a")) {
    case "foo bar":
        var v6 float64 = 1.5
        _ = v6
    case "\n\t":
        for fi7 := 0; fi7 < 15; fi7++ {
            p0 += p0
            if false {
                for fi8 := 0; fi8 < 11; fi8++ {
                    v9 := []int{fi7, (p0 - p0)}
                    _ = fi8
                    _ = v9
                }
            } else {
                var v10 uint16 = 2
                for fi11 := 0; fi11 < 19; fi11++ {
                    v12 := GenS1{F0: 10, F1: (v10 + (^(v10)))}
                    v13 := GenS2{F0: 781, F1: "hello"}
                    v14 := GenS3{F0: (v10 ^ 65535)}
                    _ = fi11
                    _ = v12
                    _ = v13
                    _ = v14
                }
                v10 = v10
                _ = v10
            }
            _ = fi7
        }
    }
    switch (^(p0)) {
    case 9223372036854775807:
        p0 = -802
    case 32768:
        v15 := []int{4294967295, (-(p0)), (p0 + p0), p0}
        _ = v15
    case 127:
        fmt.Println(p0)
    default:
        switch p0 {
        case 128:
            p0 += p0
        case 9223372036854775807:
            var v16 uint32 = 2147483647
            _ = v16
        case 32767:
            switch (("\n\t" + "a") + "") {
            case "a":
                p0 = 127
            case "hello":
                p0 += p0
            default:
                p0 -= (p0 << ((347) & 63))
            }
        default:
            v17 := GenS1{F0: 1, F1: 1}
            _ = v17
        }
    }
    var v18 string = "\n\t"
    _ = p0
    _ = v18
    return -345
}

func main() {
    var v1 int8 = 42
    v2 := []uint32{344}
    var v3 float64 = 3.14159
    for fi19 := 0; fi19 < 18; fi19++ {
        if (false || (9223372036854775807 == 42)) {
            v2 = append(v2, 574)
        } else {
            v1 = (v1 | v1)
            for fi20 := 0; fi20 < 8; fi20++ {
                fmt.Println(fi19)
                v21 := GenS3{F0: 42}
                fmt.Println(v1, fi20)
                _ = fi20
                _ = v21
            }
            var v22 int8 = (v1 << ((2) & 63))
            _ = v22
        }
        _ = fi19
    }
    for fi23 := 0; fi23 < 8; fi23++ {
        var v24 int16 = -32768
        v1 = (^(v1))
        _ = fi23
        _ = v24
    }
    v2 = append(v2, 4294967295)
    fmt.Println(v3, v3)
    v3 -= 0.0
    v25 := GenS1{F0: 255, F1: 32767}
    var v26 uint8 = 26
    v27 := GenS4{F0: 0}
    var v28 uint64 = 607
    v3 = -2.25
    var v29 int64 = -1
    switch "foo bar" {
    case "foo bar":
        v28 = v28
    case "a":
        if true {
            v2 = append(v2, 4294967295)
            v30 := GenS5{F0: v1}
            _ = v30
        } else {
            if len(v2) > 0 {
                v2[0] = 632
            }
            v31 := GenS4{F0: 1000}
            _ = v31
        }
    default:
        var v32 int = -9223372036854775808
        _ = v32
    }
    v2 = append(v2, 533)
    v33 := GenS5{F0: (v1 << ((2) & 63))}
    v26 -= (v26 / ((v26) | 1))
    v34 := GenS5{F0: v1}
    v35 := GenS6{F0: 484, F1: "a"}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v33
    _ = v34
    _ = v35
}
