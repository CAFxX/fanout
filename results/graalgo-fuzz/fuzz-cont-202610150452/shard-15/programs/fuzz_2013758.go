package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 int64; F1 bool; F2 uint64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint32; F1 int; F2 bool }
func genFn0(p0 uint16, p1 bool, p2 int32) uint64 {
    p2 = p2
    var v6 string = "foo bar"
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 1
}
func genFn1(p0 bool, p1 uint32) int32 {
    switch 0 {
    case 0:
        v7 := []int{1, 4294967295}
        _ = v7
    }
    p1 = (p1 >> ((65535) & 63))
    v8 := []int{32768, -382, -329}
    var v9 string = "a"
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 1
}

func main() {
    var v1 string = "hello"
    var v2 uint8 = 2
    var v3 int32 = 1
    v4 := []int{4294967295, 4294967295, 4294967295}
    var v5 uint64 = 4294967295
    if (!((0.1 + (3.14159 + (0.0 - float64(0.1)))) <= -2.25)) {
        v2 = (v2 & (v2 - 128))
        var v10 int16 = -32768
        _ = v10
    }
    var v11 int64 = 746
    fmt.Println(v1, v11, v11)
    if true {
        v12 := GenS1{F0: (v2 & 180), F1: v3}
        _ = v12
    } else {
        switch 1 {
        case 9223372036854775807:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case 1:
            if false {
                var v13 int32 = (v3 - v3)
                fmt.Println(v13, v3)
                var v14 uint16 = 688
                _ = v13
                _ = v14
            }
        default:
            fmt.Println(v3)
        }
    }
    v15 := GenS2{F0: v11, F1: true, F2: 1}
    var v16 int64 = (-(v11))
    for fi17 := 0; fi17 < 8; fi17++ {
        switch -9223372036854775808 {
        case 42:
            if (!false) {
                var v18 int8 = -123
                var v19 int8 = 42
                _ = v18
                _ = v19
            }
        case 32767:
            for fi20 := 0; fi20 < 5; fi20++ {
                var v21 string = (v1 + "")
                var v22 uint64 = 4294967295
                _ = fi20
                _ = v21
                _ = v22
            }
        }
        fmt.Println(v1)
        v4 = append(v4, (^(fi17)))
        _ = fi17
    }
    var v23 uint16 = 65535
    if false {
        var v24 uint32 = 4294967295
        _ = v24
    } else {
        var v25 uint16 = 255
        _ = v25
    }
    if (!(!true)) {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    } else {
        fmt.Println(v5, v2)
        for fi26 := 0; fi26 < 7; fi26++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            for fi27 := 0; fi27 < 5; fi27++ {
                var v28 int = (fi27 << ((0) & 63))
                v29 := GenS3{F0: fi26}
                v3 -= int32(v16)
                switch v1 {
                case "":
                    if ((3.14159 < 0.0) || false) {
                        var v30 bool = (!(false && (2 >= 0)))
                        _ = v30
                    } else {
                        if false {
                            v11 = 0
                        } else {
                            if (v1 > v1) {
                                var v31 int = (fi27 - (-(fi26)))
                                _ = v31
                            } else {
                                v32 := GenS4{F0: 42, F1: -820, F2: (0.0 < 0.0)}
                                var v33 float64 = 1.5
                                _ = v32
                                _ = v33
                            }
                        }
                    }
                case "hello":
                    if (false || false) {
                        var v34 uint = 672
                        v3 += v3
                        _ = v34
                    }
                default:
                    var v35 int8 = 127
                    _ = v35
                }
                _ = fi27
                _ = v28
                _ = v29
            }
            _ = fi26
        }
    }
    v16 -= v16
    v36 := GenS4{F0: 323, F1: 128, F2: false}
    v16 += (v16 / ((42) | 1))
    var v37 uint = 67
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v23
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
