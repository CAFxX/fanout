package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool; F1 int64 }
type GenS4 struct { F0 uint64 }
func genFn0(p0 uint8, p1 int32) int16 {
    var v6 uint16 = 255
    v7 := GenS1{F0: -1, F1: (p0 & p0)}
    v8 := GenS2{F0: "foo bar"}
    for fi9 := 0; fi9 < 12; fi9++ {
        v10 := GenS2{F0: "foo bar"}
        for fi11 := 0; fi11 < 2; fi11++ {
            p1 -= p1
            v6 = 27
            _ = fi11
        }
        var v12 float64 = ((0.1 / 1.5) - ((-2.25 * -2.25) + -2.25))
        _ = fi9
        _ = v10
        _ = v12
    }
    for fi13 := 0; fi13 < 8; fi13++ {
        v14 := GenS2{F0: ""}
        _ = fi13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return -10
}

func main() {
    var v1 uint8 = 128
    var v2 bool = (127 != 2147483647)
    var v3 uint = 2
    var v4 bool = true
    var v5 float32 = (0.1 + 3.14159)
    var v15 float64 = ((0.1 * (1.5 + 1e10)) * 0.0)
    if v4 {
        if ((v3 >> (((^(v3))) & 63)) != (v3 & (^(v3)))) {
            switch "foo bar" {
            case "a":
                if (!v4) {
                    var v16 uint64 = 446
                    v1 = (v1 & v1)
                    _ = v16
                } else {
                    v17 := []int{44}
                    _ = v17
                }
            case "\n\t":
                v5 = v5
            default:
                v2 = v2
            }
            v1 += v1
        }
        fmt.Println(v2)
    }
    v1 += (v1 >> ((2) & 63))
    v15 = 1e10
    v15 = 3.14159
    if (v5 != v5) {
        fmt.Println(v15, v2, v1)
        for fi18 := 0; fi18 < 17; fi18++ {
            var v19 uint32 = 1
            _ = fi18
            _ = v19
        }
    }
    for fi20 := 0; fi20 < 18; fi20++ {
        var v21 int8 = -128
        if (65535 <= 32767) {
            v1 += v1
            var v22 int = fi20
            _ = v22
        } else {
            var v23 int32 = 0
            v24 := GenS3{F0: (!(42 > 2147483647)), F1: -935}
            _ = v23
            _ = v24
        }
        if (false && false) {
            switch fi20 {
            case -782:
                for fi25 := 0; fi25 < 7; fi25++ {
                    fmt.Println(v2)
                    v2 = ((0 < -754) && v4)
                    var v26 uint = (v3 >> ((v3) & 63))
                    _ = fi25
                    _ = v26
                }
            case 127:
                fmt.Println(fi20, v1)
            case -9223372036854775808:
                if v2 {
                    switch "a" {
                    case "foo bar":
                        v27 := GenS4{F0: 18446744073709551615}
                        _ = v27
                    case "":
                        if true {
                            var v28 int16 = 1000
                            v29 := []int{-9223372036854775808, (fi20 & fi20)}
                            _ = v28
                            _ = v29
                        }
                    default:
                        v30 := []int{fi20, 42, fi20}
                        _ = v30
                    }
                    v4 = v4
                    switch fi20 {
                    case 9223372036854775807:
                        v4 = (!true)
                    case 32768:
                        fmt.Println(v5)
                    }
                }
            default:
                if (v4 || ((!true) && v2)) {
                    switch fi20 {
                    case -9223372036854775808:
                        v21 -= (-(v21))
                    case -907:
                        if (v1 < v1) {
                            v2 = true
                        }
                    }
                } else {
                    if v4 {
                        v31 := GenS1{F0: -1, F1: v1}
                        _ = v31
                    }
                    fmt.Println(v4, v2, v2)
                }
            }
        } else {
            if false {
                v32 := []int{(-(fi20))}
                fmt.Println(v1)
                _ = v32
            } else {
                if (!(v5 != 1e10)) {
                    fmt.Println(v2)
                    v3 -= (v3 + (^(v3)))
                }
                v33 := []int{32768, (fi20 * 9223372036854775807)}
                var v34 uint64 = 4294967295
                _ = v33
                _ = v34
            }
        }
        _ = fi20
        _ = v21
    }
    var v35 uint16 = 1
    switch 128 {
    case 562:
        var v36 int16 = -32768
        _ = v36
    case 692:
        v37 := GenS3{F0: v4, F1: 0}
        _ = v37
    case 32768:
        v38 := GenS2{F0: "hello"}
        _ = v38
    }
    var v39 int32 = 2147483647
    fmt.Println(v15, v39, v39)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v35
    _ = v39
    _ = v4
    _ = v5
}
