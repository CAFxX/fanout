package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 int16 }
type GenS2 struct { F0 uint16; F1 uint16; F2 string }
func genFn0(p0 int16, p1 uint32) int32 {
    for fi7 := 0; fi7 < 1; fi7++ {
        v8 := GenS1{F0: 846, F1: 1, F2: p0}
        var v9 float64 = 1.5
        for fi10 := 0; fi10 < 10; fi10++ {
            var v11 uint = 42
            _ = fi10
            _ = v11
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    switch "foo bar" {
    case "hello":
        switch -202 {
        case 32767:
            v12 := GenS1{F0: 2147483647, F1: 127, F2: int16(p1)}
            _ = v12
        default:
            if false {
                for fi13 := 0; fi13 < 16; fi13++ {
                    var v14 bool = (false || (145 < 9223372036854775807))
                    var v15 uint32 = 2
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                v16 := GenS1{F0: p1, F1: 127, F2: (p0 + p0)}
                fmt.Println(p0)
                _ = v16
            }
        }
    case "foo bar":
        v17 := []int{903, 372}
        _ = v17
    }
    v18 := []int{-9223372036854775808, 0, 32767, 187}
    if (!(true || (29 <= 1))) {
        _ = len(v18)
        if true {
            var v19 float32 = 0.0
            if (1 < -113) {
                var v20 float32 = v19
                var v21 int32 = -419
                _ = v20
                _ = v21
            } else {
                v22 := GenS1{F0: p1, F1: 100, F2: p0}
                _ = v22
            }
            v23 := GenS1{F0: 42, F1: 1, F2: p0}
            _ = v19
            _ = v23
        }
        var v24 float32 = 0.1
        _ = v24
    } else {
        var v25 uint16 = 1
        for fi26 := 0; fi26 < 13; fi26++ {
            if (!true) {
                if false {
                    p1 += 2147483647
                } else {
                    var v27 float32 = 0.1
                    for fi28 := 0; fi28 < 3; fi28++ {
                        fmt.Println(v27)
                        if len(v18) > 0 {
                            v18[0] = (fi28 * 416)
                        }
                        _ = fi28
                    }
                    _ = v27
                }
                var v29 float64 = (-2.25 / -2.25)
                _ = v29
            }
            _ = len(v18)
            _ = fi26
        }
        v25 -= 65535
        _ = v25
    }
    _ = p0
    _ = p1
    _ = v18
    return 1
}
func genFn1(p0 uint32) string {
    v30 := []int{-920, -90}
    var v31 uint64 = 627
    if len(v30) > 0 {
        v30[0] = -1
    }
    var v32 uint16 = 1
    _ = p0
    _ = v30
    _ = v31
    _ = v32
    return "a"
}

func main() {
    v1 := []uint32{21, 875, 2147483648, 378}
    v2 := []uint32{}
    var v3 int8 = -122
    v4 := []int{2147483647, 2147483647}
    var v5 int16 = 0
    var v6 int32 = -1
    fmt.Println(v3)
    _ = len(v2)
    switch 32767 {
    case 32767:
        switch "foo bar" {
        case "hello":
            if true {
                for fi33 := 0; fi33 < 6; fi33++ {
                    v4 = append(v4, fi33)
                    var v34 int = 0
                    _ = fi33
                    _ = v34
                }
                v5 -= v5
            } else {
                if false {
                    var v35 uint8 = 1
                    _ = v35
                } else {
                    switch "" {
                    case "a":
                        v36 := GenS1{F0: 604, F1: (v3 + -5), F2: v5}
                        _ = v36
                    case "\n\t":
                        var v37 uint8 = 128
                        _ = v37
                    case "foo bar":
                        var v38 float64 = (0.0 - ((0.0 / (0.0 + 3.14159)) + (float64(0.1) / 1e10)))
                        _ = v38
                    default:
                        v39 := GenS1{F0: 42, F1: 89, F2: v5}
                        _ = v39
                    }
                }
                v6 = v6
                v40 := GenS1{F0: 2, F1: -1, F2: v5}
                _ = v40
            }
        case "\n\t":
            var v41 string = "\n\t"
            _ = v41
        }
    case 127:
        v6 += (v6 | 2147483647)
    }
    var v42 int16 = (v5 & (v5 | (v5 & 0)))
    if false {
        switch int(v3) {
        case 4294967295:
            if (9223372036854775807 != 32767) {
                switch ("\n\t" + "a") {
                case "\n\t":
                    var v43 float32 = 1.5
                    _ = v43
                case "":
                    fmt.Println(v6)
                case "hello":
                    v6 -= v6
                }
                var v44 int8 = int8(v6)
                _ = v44
            } else {
                v1 = append(v1, 65535)
                if true {
                    var v45 uint = 363
                    _ = v45
                }
            }
        default:
            v2 = append(v2, 874)
        }
        v2 = append(v2, 1)
    }
    for fi46 := 0; fi46 < 19; fi46++ {
        v42 -= v42
        switch ("foo bar" + "") {
        case "\n\t":
            v3 -= v3
        case "hello":
            switch "" {
            case "foo bar":
                v42 = -1
            default:
                fmt.Println(v5, v5)
            }
        default:
            var v47 string = "\n\t"
            _ = v47
        }
        _ = fi46
    }
    v3 = v3
    if ("hello" <= "a") {
        if (-120 < v3) {
            var v48 uint64 = 290
            v49 := GenS1{F0: 256, F1: -128, F2: -32768}
            _ = v48
            _ = v49
        }
    } else {
        for fi50 := 0; fi50 < 20; fi50++ {
            v6 += v6
            for fi51 := 0; fi51 < 16; fi51++ {
                if false {
                    v5 = v5
                }
                _ = fi51
            }
            _ = fi50
        }
        v6 -= (v6 + (v6 & genFn0((v42 >> ((0) & 63)), 2147483648)))
    }
    fmt.Println(v5, v42)
    fmt.Println(v6, v5, v6)
    v3 = v3
    if false {
        var v52 int8 = (v3 / ((v3) | 1))
        for fi53 := 0; fi53 < 8; fi53++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v54 uint8 = 255
            _ = fi53
            _ = v54
        }
        var v55 int16 = -108
        _ = v52
        _ = v55
    }
    if true {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v56 bool = (v42 > v5)
        var v57 int = -9223372036854775808
        _ = v56
        _ = v57
    } else {
        var v58 float32 = ((float32(0.1) / 0.1) - 3.14159)
        v59 := GenS2{F0: 255, F1: 255, F2: "foo bar"}
        _ = v58
        _ = v59
    }
    fmt.Println(v42, v42)
    switch 15 {
    case 9223372036854775807:
        v6 = v6
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
