package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 uint32; F1 uint16 }
type GenS4 struct { F0 int64; F1 int; F2 int8 }
type GenS5 struct { F0 uint64; F1 uint16 }
func genFn0(p0 int32, p1 string) int16 {
    for fi6 := 0; fi6 < 3; fi6++ {
        var v7 uint = 18446744073709551615
        v7 = (v7 << (((v7 + v7)) & 63))
        p0 = (^(p0))
        _ = fi6
        _ = v7
    }
    fmt.Println(p1)
    var v8 uint32 = 488
    _ = p0
    _ = p1
    _ = v8
    return int16(v8)
}
func genFn1(p0 uint) string {
    p0 += p0
    for fi9 := 0; fi9 < 2; fi9++ {
        var v10 uint32 = 2147483647
        fmt.Println(fi9)
        v11 := GenS1{F0: -2147483648}
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = p0
    return ""
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]uint32)
    var v3 int64 = 174
    var v4 float32 = (-2.25 - 3.14159)
    var v5 float32 = v4
    var v12 int64 = (-(v3))
    v13 := GenS1{F0: 42}
    v12 = -9223372036854775808
    if true {
        v3 = (v3 & v12)
        fmt.Println(v4, v3)
        fmt.Println(v5, v4)
    } else {
        v14 := GenS1{F0: 1}
        _ = v14
    }
    fmt.Println(v3)
    v15 := GenS1{F0: 281}
    v4 = 0.1
    for fi16 := 0; fi16 < 8; fi16++ {
        if (!(false && ((!false) || false))) {
            v17 := GenS2{F0: 2, F1: 0}
            _ = v17
        } else {
            fmt.Println(fi16)
        }
        v18 := GenS2{F0: 1, F1: 42}
        _ = fi16
        _ = v18
    }
    if (false && false) {
        if false {
            switch "foo bar" {
            case "\n\t":
                v19 := GenS2{F0: 2, F1: 0}
                _ = v19
            default:
                var v20 bool = false
                _ = v20
            }
        } else {
            switch 0 {
            case -1:
                v21 := v1[4294967295]
                _ = v21
            default:
                if true {
                    v22 := GenS1{F0: -358}
                    _ = v22
                }
            }
            fmt.Println(v5, v5)
            var v23 int = -505
            _ = v23
        }
        if true {
            if false {
                v3 -= (-(v12))
                switch "foo bar" {
                case "a":
                    var v24 int8 = -1
                    _ = v24
                case "hello":
                    v25 := []int{-2147483648, 32767}
                    _ = v25
                default:
                    switch 9223372036854775807 {
                    case 32767:
                        v26 := GenS3{F0: 65535, F1: 267}
                        _ = v26
                    case 4294967295:
                        var v27 uint16 = 255
                        _ = v27
                    case 2147483647:
                        v5 = (v5 * v5)
                    default:
                        for fi28 := 0; fi28 < 10; fi28++ {
                            v29 := GenS4{F0: v3, F1: fi28, F2: 100}
                            _ = fi28
                            _ = v29
                        }
                    }
                }
            } else {
                v4 = (float32(float64(0.0)) * (v4 / (((1.5 + v4) + v4) + v5)))
                var v30 float64 = 0.1
                var v31 int16 = genFn0(349, ("a" + "\n\t"))
                _ = v30
                _ = v31
            }
        }
    } else {
        switch ("" + ("a" + "foo bar")) {
        case "hello":
            fmt.Println(v4)
        default:
            v5 = -2.25
        }
        if false {
            var v32 uint32 = 1
            if (true || (!false)) {
                v32 = v32
                var v33 int8 = 73
                _ = v33
            } else {
                v34 := []int{1}
                _ = v34
            }
            _ = v32
        } else {
            v35 := GenS3{F0: 1, F1: 0}
            _ = v35
        }
        v5 -= ((v5 * (0.0 + 1.5)) + -2.25)
    }
    switch 1 {
    case 98:
        v36 := GenS4{F0: (v12 >> ((617) & 63)), F1: 0, F2: -81}
        _ = v36
    case 4294967295:
        v2[0] = 755
    default:
        v37 := GenS3{F0: 0, F1: 32767}
        _ = v37
    }
    if true {
        v12 = (v3 & v12)
        v3 = (v3 + 42)
        var v38 int64 = v3
        _ = v38
    }
    var v39 uint8 = 91
    v39 = v39
    if (!(692 <= 0)) {
        v40 := GenS5{F0: 0, F1: 255}
        v5 = (float32(1e10) / (((1.5 - float32(-2.25)) * float32(1e10)) - (v4 - float32(0.1))))
        _ = v40
    }
    var v41 int8 = 42
    if ((799 <= 695) && (true || false)) {
        if (106 == 0) {
            v42 := GenS4{F0: (^(v3)), F1: 547, F2: 127}
            _ = v42
        } else {
            v43 := GenS2{F0: 18446744073709551615, F1: 4294967295}
            var v44 float64 = (1.5 / 0.1)
            _ = v43
            _ = v44
        }
        if (-27 == 380) {
            v3 += v12
        }
    }
    v45 := v2[973]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v41
    _ = v45
    _ = v5
}
