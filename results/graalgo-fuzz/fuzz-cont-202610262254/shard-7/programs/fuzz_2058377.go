package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64; F1 string }
type GenS3 struct { F0 bool; F1 uint32; F2 string }
type GenS4 struct { F0 string; F1 int64; F2 uint16 }
func genFn0() uint64 {
    var v5 uint8 = 215
    v5 += 114
    var v6 float32 = 0.0
    _ = v5
    _ = v6
    return 0
}
func genFn1() string {
    if true {
        var v7 bool = false
        for fi8 := 0; fi8 < 10; fi8++ {
            if (42 != -2147483648) {
                v7 = v7
                fmt.Println(v7, v7)
                v7 = (("" + "\n\t") < "")
            }
            var v9 bool = v7
            _ = fi8
            _ = v9
        }
        if (!(true || v7)) {
            v7 = v7
            v7 = false
            v7 = (-171 <= 0)
        }
        _ = v7
    } else {
        switch "\n\t" {
        case "foo bar":
            var v10 int16 = 0
            _ = v10
        case "a":
            var v11 int8 = 127
            _ = v11
        }
        var v12 uint8 = 8
        _ = v12
    }
    v13 := []int{32767}
    _ = len(v13)
    var v14 bool = (!false)
    if v14 {
        fmt.Println(v14)
        v14 = v14
        switch 42 {
        case 4294967295:
            fmt.Println(v14)
        case -9223372036854775808:
            if v14 {
                v13 = append(v13, -396)
            } else {
                v14 = v14
            }
        case 32767:
            v14 = (!(0.0 < 1.5))
        default:
            switch "a" {
            case "\n\t":
                var v15 uint = 65535
                _ = v15
            case "foo bar":
                var v16 uint16 = 32767
                _ = v16
            }
        }
    }
    _ = v13
    _ = v14
    return "\n\t"
}

func main() {
    var v1 uint8 = 128
    var v2 int8 = -128
    v3 := make(map[uint32]int)
    v4 := make(map[string]uint32)
    if false {
        v2 = (v2 / (((v2 * v2)) | 1))
    } else {
        v17 := GenS1{F0: 89}
        _ = v17
    }
    v3[0] = 521
    var v18 uint8 = (^(v1))
    var v19 uint8 = (v1 - v1)
    for fi20 := 0; fi20 < 19; fi20++ {
        fmt.Println(fi20, fi20, v18)
        _ = len(v4)
        _ = fi20
    }
    v21 := GenS1{F0: (v18 | (v19 << ((121) & 63)))}
    if true {
        fmt.Println(v2)
    } else {
        v22 := GenS1{F0: (v19 & (^(v1)))}
        var v23 string = "a"
        _ = v22
        _ = v23
    }
    v24 := []int{2147483647, -2147483648, 127}
    if (42 != -274) {
        var v25 int16 = 0
        v19 += (v18 & v19)
        _ = v25
    } else {
        fmt.Println(v18, v18)
        v1 = 20
        fmt.Println(v2, v1)
    }
    for fi26 := 0; fi26 < 5; fi26++ {
        _ = len(v24)
        _ = fi26
    }
    v2 -= (v2 >> ((42) & 63))
    v18 = v1
    v27 := GenS1{F0: (^(v18))}
    if true {
        var v28 uint64 = 18446744073709551615
        _ = v28
    } else {
        if (2 <= 69) {
            switch "a" {
            case "\n\t":
                switch "hello" {
                case "":
                    fmt.Println(v18)
                default:
                    if true {
                        if false {
                            v24 = append(v24, 127)
                            v29 := GenS2{F0: 0, F1: "foo bar"}
                            _ = v29
                        } else {
                            var v30 uint32 = 1
                            for fi31 := 0; fi31 < 16; fi31++ {
                                fmt.Println(v19, fi31)
                                var v32 uint16 = 1
                                _ = fi31
                                _ = v32
                            }
                            _ = v30
                        }
                    } else {
                        if (true || (-1 > -2147483648)) {
                            var v33 bool = ((!(((3.14159 / 1.5) <= 1.5) && (!true))) && ((!false) || (!true)))
                            fmt.Println(v2)
                            _ = v33
                        }
                        _ = len(v24)
                    }
                }
            default:
                v18 = 255
            }
            for fi34 := 0; fi34 < 10; fi34++ {
                _ = len(v24)
                delete(v3, 2)
                for fi35 := 0; fi35 < 15; fi35++ {
                    v36 := GenS3{F0: (!(((v2 % ((v2) | 1)) < (-(v2))) || (true || (!(2147483647 != -1))))), F1: 2147483647, F2: ("a" + ("\n\t" + "hello"))}
                    _ = fi35
                    _ = v36
                }
                _ = fi34
            }
        } else {
            var v37 float32 = (3.14159 + (float32((0.1 * 0.1)) / 1e10))
            _ = v37
        }
    }
    switch -1 {
    case 128:
        switch "foo bar" {
        case "a":
            switch "" {
            case "":
                v38 := GenS4{F0: "\n\t", F1: -9223372036854775808, F2: 65535}
                _ = v38
            }
        case "":
            switch 9223372036854775807 {
            case 32767:
                fmt.Println(v2, v19, v19)
            }
        case "hello":
            delete(v3, 302)
        }
    case 357:
        v39 := GenS4{F0: "\n\t", F1: 42, F2: 42}
        _ = v39
    case 2147483647:
        fmt.Println(v18)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v24
    _ = v27
    _ = v3
    _ = v4
}
