package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 int32 }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 uint8; F1 uint16 }
type GenS4 struct { F0 int8 }
func genFn0(p0 int32, p1 int64) int16 {
    if false {
        switch 32768 {
        case -1:
            fmt.Println(p0)
        case 9223372036854775807:
            v5 := []int{128, 0, 0, 127}
            _ = v5
        case 824:
            fmt.Println(p0)
        }
    }
    fmt.Println(p0, p0)
    fmt.Println(p1)
    if true {
        fmt.Println(p1, p1)
    }
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    return -32768
}

func main() {
    var v1 uint32 = 1
    var v2 float32 = 0.0
    var v3 float32 = v2
    v4 := make(map[int]int)
    if true {
        _ = len(v4)
        if false {
            var v6 uint64 = 211
            var v7 uint8 = 164
            if (!(1.5 <= (float32((0.1 - (-2.25 * 1e10))) - v2))) {
                var v8 uint8 = 255
                _ = v8
            } else {
                v9 := []int{-2147483648, 415}
                v10 := v4[-9223372036854775808]
                _ = v10
                _ = v9
            }
            _ = v6
            _ = v7
        } else {
            v11 := v4[32767]
            _ = v11
        }
    }
    fmt.Println(v1, v2, v3)
    var v12 int64 = 42
    switch "foo bar" {
    case "hello":
        for fi13 := 0; fi13 < 3; fi13++ {
            _ = len(v4)
            _ = fi13
        }
    case "foo bar":
        v2 = -2.25
    case "a":
        for fi14 := 0; fi14 < 3; fi14++ {
            v15 := v4[fi14]
            if true {
                var v16 int16 = 0
                _ = v16
            } else {
                for fi17 := 0; fi17 < 6; fi17++ {
                    v3 = (v2 + v2)
                    var v18 int = (-(fi17))
                    _ = fi17
                    _ = v18
                }
                switch fi14 {
                case -9223372036854775808:
                    v2 = (v3 - v3)
                case 128:
                    if (v12 <= v12) {
                        var v19 int32 = 2147483647
                        var v20 bool = (-9223372036854775808 > v12)
                        var v21 int16 = 0
                        _ = v19
                        _ = v20
                        _ = v21
                    } else {
                        delete(v4, (fi14 * (fi14 >> ((4294967295) & 63))))
                        var v22 int64 = (v12 ^ v12)
                        _ = v22
                    }
                default:
                    var v23 int64 = v12
                    _ = v23
                }
                v24 := GenS1{F0: 761, F1: uint8(v1), F2: 2147483647}
                _ = v24
            }
            _ = fi14
            _ = v15
        }
    default:
        v25 := []int{}
        _ = v25
    }
    v1 += v1
    switch "hello" {
    case "foo bar":
        fmt.Println(v12, v2, v2)
    case "\n\t":
        var v26 uint64 = 0
        _ = v26
    case "":
        for fi27 := 0; fi27 < 12; fi27++ {
            if false {
                var v28 uint32 = 1
                var v29 uint64 = 2
                v30 := GenS2{F0: 32767, F1: (("a" + (("" + "foo bar") + "hello")) + "\n\t")}
                _ = v28
                _ = v29
                _ = v30
            }
            _ = fi27
        }
    }
    v31 := GenS2{F0: 18, F1: ("foo bar" + "foo bar")}
    for fi32 := 0; fi32 < 8; fi32++ {
        switch "" {
        case "\n\t":
            fmt.Println(v1, v1)
        case "":
            if false {
                switch fi32 {
                case 4294967295:
                    switch fi32 {
                    case 2147483647:
                        var v33 uint64 = 42
                        _ = v33
                    case -586:
                        var v34 bool = false
                        _ = v34
                    }
                case 1:
                    v1 -= 65535
                }
            } else {
                fmt.Println(v12, v1)
                switch "hello" {
                case "\n\t":
                    v35 := GenS2{F0: 400, F1: ("\n\t" + ("" + "a"))}
                    _ = v35
                }
            }
        case "a":
            var v36 int32 = 2147483647
            _ = v36
        }
        if ((1 == 127) && true) {
            v37 := []int{9223372036854775807, 559, -1, -37}
            v38 := GenS3{F0: 1, F1: 32767}
            _ = v37
            _ = v38
        } else {
            switch ("\n\t" + "foo bar") {
            case "\n\t":
                var v39 int8 = 127
                _ = v39
            }
            v1 = v1
        }
        v40 := GenS4{F0: -128}
        _ = fi32
        _ = v40
    }
    v41 := []int{32768, -73, -849}
    var v42 int64 = (^(v12))
    var v43 uint64 = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v41
    _ = v42
    _ = v43
}
