package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8; F1 uint64 }
func genFn0(p0 int16, p1 uint) int32 {
    v6 := GenS1{F0: p1}
    fmt.Println(p0, p1)
    switch "\n\t" {
    case "hello":
        var v7 float64 = (1e10 + (1.5 * 1.5))
        _ = v7
    }
    if false {
        p1 += p1
    } else {
        var v8 uint16 = 514
        var v9 int64 = 264
        for fi10 := 0; fi10 < 14; fi10++ {
            var v11 uint = (p1 * p1)
            _ = fi10
            _ = v11
        }
        _ = v8
        _ = v9
    }
    if (!true) {
        var v12 string = ""
        switch v12 {
        case "foo bar":
            p1 -= (p1 % (((p1 >> ((1) & 63))) | 1))
        case "\n\t":
            if (-84 > -1) {
                var v13 int8 = -1
                for fi14 := 0; fi14 < 3; fi14++ {
                    for fi15 := 0; fi15 < 16; fi15++ {
                        v16 := GenS2{F0: (p1 >= 65535)}
                        switch "\n\t" {
                        case "\n\t":
                            var v17 int32 = 559
                            _ = v17
                        case "":
                            v18 := GenS2{F0: false}
                            _ = v18
                        }
                        _ = fi15
                        _ = v16
                    }
                    var v19 uint = 18446744073709551615
                    _ = fi14
                    _ = v19
                }
                v20 := GenS2{F0: (42 <= 18446744073709551615)}
                _ = v13
                _ = v20
            } else {
                v21 := []int{1, 744, 4294967295}
                var v22 int = 128
                _ = v21
                _ = v22
            }
        case "":
            if false {
                fmt.Println(p1, p0, v12)
                p1 = p1
                p0 -= (p0 * p0)
            } else {
                v23 := []int{2147483647, -89}
                _ = v23
            }
        default:
            fmt.Println(p0, v12)
        }
        switch 445 {
        case -2147483648:
            switch 0 {
            case 42:
                var v24 int = -127
                _ = v24
            }
        case 0:
            switch 42 {
            case 32767:
                v12 = ((v12 + "\n\t") + v12)
            case -311:
                v25 := []int{128}
                _ = v25
            default:
                var v26 uint16 = 0
                _ = v26
            }
        case 2147483647:
            p1 -= (p1 << (((p1 * (p1 / ((538) | 1)))) & 63))
        default:
            if false {
                if (!false) {
                    p1 = (p1 * p1)
                    for fi27 := 0; fi27 < 1; fi27++ {
                        var v28 float64 = 1.5
                        _ = fi27
                        _ = v28
                    }
                } else {
                    v12 = v12
                }
            }
        }
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v6
    return -1
}
func genFn1() int {
    switch 907 {
    case 32767:
        for fi29 := 0; fi29 < 14; fi29++ {
            v30 := []int{-2147483648}
            _ = fi29
            _ = v30
        }
    case 127:
        v31 := []int{1}
        _ = v31
    }
    var v32 int32 = 609
    var v33 string = ""
    if (true && false) {
        var v34 uint64 = 42
        _ = v34
    } else {
        for fi35 := 0; fi35 < 12; fi35++ {
            v36 := GenS1{F0: 255}
            _ = fi35
            _ = v36
        }
        var v37 string = v33
        _ = v37
    }
    v33 = v33
    _ = v32
    _ = v33
    return 32768
}

func main() {
    v1 := []int{9223372036854775807, 4294967295, 2147483647, -2147483648}
    var v2 float32 = 3.14159
    var v3 uint = 2
    v4 := make(map[uint32]uint32)
    var v5 uint = 909
    switch -711 {
    case 217:
        v38 := v4[544]
        _ = v38
    case 4294967295:
        var v39 int32 = 0
        _ = v39
    }
    v2 = (0.1 * v2)
    var v40 uint = (^(v3))
    v41 := GenS3{F0: -128, F1: 882}
    if len(v1) > 0 {
        v1[0] = -9223372036854775808
    }
    v3 += 594
    var v42 uint16 = 225
    fmt.Println(v5, v40, v42)
    if true {
        fmt.Println(v5)
    }
    var v43 string = "a"
    v3 = v40
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v5
}
