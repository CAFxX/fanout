package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32; F1 bool; F2 string }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int8; F1 bool }
type GenS5 struct { F0 int }
func genFn0(p0 int8) uint16 {
    var v4 int32 = 672
    switch "foo bar" {
    case "foo bar":
        p0 -= 94
    case "\n\t":
        for fi5 := 0; fi5 < 5; fi5++ {
            switch fi5 {
            case 968:
                fmt.Println(v4, v4, v4)
            case -2147483648:
                fmt.Println(fi5)
            default:
                v6 := []int{32767}
                _ = v6
            }
            for fi7 := 0; fi7 < 11; fi7++ {
                p0 += 1
                if false {
                    var v8 int16 = 42
                    switch "a" {
                    case "":
                        v8 -= v8
                    case "\n\t":
                        if false {
                            var v9 int8 = p0
                            v10 := []int{}
                            _ = v10
                            _ = v9
                        }
                    default:
                        if false {
                            p0 = (-(p0))
                        }
                    }
                    _ = v8
                }
                v11 := GenS1{F0: "foo bar"}
                _ = fi7
                _ = v11
            }
            v12 := GenS1{F0: "\n\t"}
            _ = fi5
            _ = v12
        }
    case "a":
        v4 = (-(v4))
    default:
        v4 = v4
    }
    switch 128 {
    case 221:
        fmt.Println(v4)
    case 1:
        for fi13 := 0; fi13 < 2; fi13++ {
            v14 := GenS1{F0: "a"}
            if (3.14159 > (0.1 + 0.0)) {
                v15 := []int{(-(fi13)), (fi13 & fi13), fi13}
                _ = v15
            }
            _ = fi13
            _ = v14
        }
    case 65:
        v16 := []int{}
        _ = v16
    default:
        v4 += 42
    }
    for fi17 := 0; fi17 < 17; fi17++ {
        var v18 int64 = -1
        _ = fi17
        _ = v18
    }
    _ = p0
    _ = v4
    return 32767
}

func main() {
    var v1 uint8 = 123
    v2 := make(map[uint32]string)
    var v3 int32 = -1
    fmt.Println(v1, v1)
    var v19 int64 = -9223372036854775808
    for fi20 := 0; fi20 < 6; fi20++ {
        v19 += v19
        _ = fi20
    }
    fmt.Println(v19)
    if false {
        if false {
            v21 := GenS1{F0: ("" + "foo bar")}
            _ = v21
        } else {
            v22 := GenS2{F0: 71, F1: (-2.25 == -2.25), F2: ("" + ("foo bar" + "\n\t"))}
            var v23 uint64 = 2
            var v24 uint32 = 4294967295
            _ = v22
            _ = v23
            _ = v24
        }
        if (1 < -1) {
            v25 := GenS3{F0: "hello"}
            _ = v25
        } else {
            for fi26 := 0; fi26 < 13; fi26++ {
                v27 := GenS4{F0: 100, F1: (false || false)}
                _ = fi26
                _ = v27
            }
            v28 := GenS5{F0: -1}
            _ = v28
        }
    }
    var v29 uint8 = v1
    var v30 uint16 = genFn0(-125)
    var v31 int8 = -1
    switch -2147483648 {
    case 42:
        var v32 float64 = 0.0
        _ = v32
    }
    _ = len(v2)
    for fi33 := 0; fi33 < 8; fi33++ {
        v29 -= 1
        switch "a" {
        case "hello":
            fmt.Println(fi33, v30, v1)
        case "\n\t":
            v34 := GenS5{F0: (fi33 | (fi33 & (fi33 - int(v19))))}
            _ = v34
        case "":
            v1 = (v1 + v29)
        default:
            v29 = v1
        }
        _ = fi33
    }
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v19
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
}
