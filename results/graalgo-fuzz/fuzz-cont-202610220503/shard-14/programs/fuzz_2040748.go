package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int64; F1 int64 }
func genFn0(p0 int, p1 int, p2 int64) uint8 {
    fmt.Println(p2, p0, p0)
    var v5 string = ((((("foo bar" + "\n\t") + "foo bar") + "a") + ("\n\t" + (("" + ("" + "foo bar")) + ""))) + "")
    var v6 int16 = 1000
    switch (p1 % ((p0) | 1)) {
    case 1:
        fmt.Println(p1, v6, v5)
    case 0:
        var v7 bool = true
        _ = v7
    default:
        for fi8 := 0; fi8 < 18; fi8++ {
            for fi9 := 0; fi9 < 13; fi9++ {
                fmt.Println(v6)
                _ = fi9
            }
            for fi10 := 0; fi10 < 16; fi10++ {
                v11 := GenS1{F0: 901, F1: -1, F2: 4294967295}
                var v12 int = (p0 + (p1 | p0))
                _ = fi10
                _ = v11
                _ = v12
            }
            var v13 bool = true
            _ = fi8
            _ = v13
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 1
}

func main() {
    var v1 int = 705
    var v2 float64 = 1e10
    var v3 int32 = -2147483648
    var v4 int8 = 100
    v2 += float64(float32(float64(1e10)))
    v14 := GenS2{F0: 42}
    v15 := GenS2{F0: v3}
    v3 += v3
    var v16 int8 = (-(v4))
    v17 := GenS1{F0: 206, F1: 0, F2: 250}
    for fi18 := 0; fi18 < 10; fi18++ {
        v19 := GenS3{F0: 0}
        switch "a" {
        case "":
            v20 := []int{(fi18 - v1)}
            _ = v20
        case "a":
            var v21 uint8 = 128
            _ = v21
        case "foo bar":
            if true {
                v22 := []int{}
                v22 = append(v22, fi18)
                v3 = 205
                _ = v22
            } else {
                switch v1 {
                case 127:
                    fmt.Println(v16, v1)
                case -754:
                    v23 := GenS2{F0: (-(v3))}
                    _ = v23
                }
                var v24 int32 = v3
                var v25 int = 4294967295
                _ = v24
                _ = v25
            }
        default:
            if ((true && (113 > 42)) && (-2.25 > (0.1 * 1e10))) {
                switch v1 {
                case -9223372036854775808:
                    for fi26 := 0; fi26 < 13; fi26++ {
                        fmt.Println(v16, v2)
                        v27 := GenS2{F0: (-(v3))}
                        v1 += (fi26 ^ (fi26 >> ((63) & 63)))
                        v28 := []int{(fi26 + v1), -555, fi26, 91}
                        var v29 uint8 = 193
                        _ = fi26
                        _ = v27
                        _ = v28
                        _ = v29
                    }
                case 9223372036854775807:
                    for fi30 := 0; fi30 < 11; fi30++ {
                        v3 = (^(v3))
                        v31 := GenS2{F0: v3}
                        _ = fi30
                        _ = v31
                    }
                case 4294967295:
                    var v32 uint16 = 1
                    _ = v32
                default:
                    v33 := GenS4{F0: -1, F1: 9223372036854775807}
                    _ = v33
                }
                var v34 uint32 = 0
                if false {
                    if false {
                        v35 := []int{v1}
                        v34 = (v34 + v34)
                        _ = v35
                    }
                    if true {
                        var v36 uint8 = 42
                        var v37 int8 = v16
                        v37 += (-(v37))
                        _ = v36
                        _ = v37
                    }
                    v38 := GenS1{F0: 717, F1: 100, F2: (v34 & v34)}
                    _ = v38
                }
                _ = v34
            } else {
                var v39 uint8 = 208
                _ = v39
            }
        }
        var v40 float32 = (1.5 + (0.1 - 1.5))
        _ = fi18
        _ = v19
        _ = v40
    }
    var v41 int8 = 0
    for fi42 := 0; fi42 < 18; fi42++ {
        fmt.Println(v1)
        var v43 uint8 = 127
        var v44 int64 = 9223372036854775807
        _ = fi42
        _ = v43
        _ = v44
    }
    for fi45 := 0; fi45 < 14; fi45++ {
        v3 = v3
        v46 := []int{4294967295, (v1 - 128)}
        _ = fi45
        _ = v46
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v41
}
