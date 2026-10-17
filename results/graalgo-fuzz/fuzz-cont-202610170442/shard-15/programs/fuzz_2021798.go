package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int; F1 string }
type GenS4 struct { F0 uint8; F1 uint }
type GenS5 struct { F0 uint }
func genFn0(p0 uint8) bool {
    fmt.Println(p0)
    v5 := GenS1{F0: 0, F1: p0}
    if false {
        v6 := GenS2{F0: 255}
        v7 := []int{-1, 0, 4294967295, 289}
        _ = v6
        _ = v7
    }
    _ = p0
    _ = v5
    return false
}

func main() {
    var v1 uint = 42
    v2 := make(map[int]string)
    var v3 float32 = 0.0
    v4 := make(map[int]int)
    var v8 uint64 = 18446744073709551615
    for fi9 := 0; fi9 < 7; fi9++ {
        var v10 int32 = 1
        _ = fi9
        _ = v10
    }
    v3 = v3
    v8 -= (^(v8))
    switch 1 {
    case -326:
        fmt.Println(v8)
    case 1:
        v11 := []int{644, -168, 32767}
        _ = v11
    default:
        var v12 int64 = 42
        _ = v12
    }
    v13 := GenS2{F0: uint16(v8)}
    for fi14 := 0; fi14 < 11; fi14++ {
        v15 := []int{fi14, (fi14 ^ -9223372036854775808)}
        v8 = (v8 % (((^(v8))) | 1))
        if (!(uint32(v1) < 0)) {
            v3 = 0.0
            _ = len(v15)
            switch 1 {
            case 128:
                if (729 >= (v8 - v8)) {
                    v1 = v1
                    for fi16 := 0; fi16 < 8; fi16++ {
                        if false {
                            v8 = (v8 & (^(v8)))
                        }
                        v1 = v1
                        fmt.Println(v3, v3)
                        _ = fi16
                    }
                    for fi17 := 0; fi17 < 16; fi17++ {
                        fmt.Println(v8, v3, v1)
                        _ = fi17
                    }
                }
            }
        } else {
            v18 := GenS2{F0: 32767}
            _ = len(v15)
            if (-1 < 123) {
                fmt.Println(v3, v3, fi14)
                if (42 >= 1000) {
                    var v19 uint = (^(v1))
                    for fi20 := 0; fi20 < 4; fi20++ {
                        var v21 int32 = 1
                        _ = fi20
                        _ = v21
                    }
                    v22 := GenS3{F0: fi14, F1: "a"}
                    _ = v19
                    _ = v22
                }
                v1 -= (v1 + (v1 & v1))
            }
            _ = v18
        }
        _ = fi14
        _ = v15
    }
    if true {
        v23 := GenS4{F0: 255, F1: (v1 + (v1 % ((v1) | 1)))}
        _ = v23
    }
    v24 := []int{2147483647, 9223372036854775807, 1}
    v8 = (v8 - v8)
    if (false && true) {
        var v25 string = ""
        v3 -= 0.0
        _ = v25
    } else {
        var v26 uint64 = v8
        for fi27 := 0; fi27 < 8; fi27++ {
            if len(v24) > 0 {
                v24[0] = fi27
            }
            fmt.Println(v26, v8, v3)
            _ = fi27
        }
        _ = v26
    }
    if (0.0 <= 0.1) {
        if false {
            switch ("" + (("a" + "\n\t") + "hello")) {
            case "foo bar":
                var v28 uint8 = 255
                _ = v28
            case "hello":
                fmt.Println(v1, v8)
            case "a":
                if false {
                    fmt.Println(v3, v1)
                    if false {
                        v3 = v3
                        v29 := GenS5{F0: v1}
                        _ = v29
                    }
                    if len(v24) >= 2 {
                        v24 = v24[1:len(v24)]
                    }
                } else {
                    if true {
                        var v30 uint16 = 1
                        if len(v24) >= 2 {
                            v24 = v24[1:len(v24)]
                        }
                        if len(v24) > 0 {
                            v24[0] = 537
                        }
                        _ = v30
                    } else {
                        var v31 float64 = 3.14159
                        _ = v31
                    }
                }
            }
        }
        var v32 int32 = -2147483648
        _ = v32
    }
    var v33 uint64 = v8
    for fi34 := 0; fi34 < 7; fi34++ {
        v24 = append(v24, (^(fi34)))
        for fi35 := 0; fi35 < 14; fi35++ {
            v36 := v2[(-(fi35))]
            if len(v24) >= 2 {
                v24 = v24[1:len(v24)]
            }
            _ = fi35
            _ = v36
        }
        _ = fi34
    }
    var v37 int = -9223372036854775808
    for fi38 := 0; fi38 < 3; fi38++ {
        var v39 int64 = 647
        _ = fi38
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v24
    _ = v3
    _ = v33
    _ = v37
    _ = v4
    _ = v8
}
