package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 int32 }
type GenS2 struct { F0 int16; F1 bool; F2 uint8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint; F1 int8; F2 uint64 }
type GenS5 struct { F0 int16; F1 int32 }
func genFn0(p0 uint32) uint16 {
    var v4 int16 = 463
    var v5 string = ""
    v6 := GenS1{F0: 255, F1: (!true), F2: 42}
    v7 := []int{1, 30, 127}
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 810
}
func genFn1(p0 int16) int32 {
    switch 171 {
    case 4294967295:
        v8 := GenS2{F0: -1, F1: false, F2: 127}
        _ = v8
    case -1:
        if false {
            fmt.Println(p0)
            v9 := GenS1{F0: 42, F1: (!(p0 > (p0 >> ((42) & 63)))), F2: 42}
            v10 := []int{975, -729, 0, 674}
            _ = v10
            _ = v9
        } else {
            v11 := []int{610}
            _ = v11
        }
    default:
        p0 = (p0 << ((42) & 63))
    }
    v12 := GenS3{F0: 42}
    if (((true && true) && false) && false) {
        fmt.Println(p0)
        switch "\n\t" {
        case "a":
            var v13 float32 = -2.25
            _ = v13
        case "":
            p0 = (p0 >> ((18446744073709551615) & 63))
        case "foo bar":
            var v14 string = "hello"
            _ = v14
        }
        var v15 uint64 = 42
        _ = v15
    }
    p0 += p0
    _ = p0
    _ = v12
    return 0
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint16 = 255
    v3 := make(map[string]bool)
    v16 := GenS4{F0: 0, F1: 42, F2: 2}
    switch 46 {
    case 9223372036854775807:
        v2 = 32767
    default:
        v17 := GenS2{F0: 42, F1: true, F2: 197}
        _ = v17
    }
    v1[9223372036854775807] = (9223372036854775807 < -1)
    for fi18 := 0; fi18 < 14; fi18++ {
        v2 -= genFn0(65535)
        _ = fi18
    }
    v2 -= genFn0(1)
    switch 9223372036854775807 {
    case -9223372036854775808:
        if (true && true) {
            var v19 int8 = 127
            fmt.Println(v2, v19)
            for fi20 := 0; fi20 < 18; fi20++ {
                if (2 >= 42) {
                    var v21 uint16 = v2
                    v2 += (v2 / ((569) | 1))
                    _ = v21
                } else {
                    if ("foo bar" == "foo bar") {
                        var v22 uint16 = v2
                        var v23 uint32 = 65535
                        _ = v22
                        _ = v23
                    } else {
                        switch "foo bar" {
                        case "hello":
                            switch fi20 {
                            case -1:
                                v2 = (v2 << ((141) & 63))
                            }
                        case "":
                            for fi24 := 0; fi24 < 16; fi24++ {
                                _ = len(v3)
                                v19 += v19
                                v19 -= 100
                                v19 -= -110
                                _ = fi24
                            }
                        }
                    }
                    v2 += (v2 + v2)
                }
                _ = fi20
            }
            _ = v19
        }
    case 466:
        if false {
            for fi25 := 0; fi25 < 1; fi25++ {
                fmt.Println(fi25, fi25)
                _ = fi25
            }
        }
    case -1:
        v2 += v2
    default:
        for fi26 := 0; fi26 < 17; fi26++ {
            var v27 bool = (-498 == -897)
            v28 := GenS5{F0: -32768, F1: int32(v2)}
            _ = fi26
            _ = v27
            _ = v28
        }
    }
    for fi29 := 0; fi29 < 3; fi29++ {
        var v30 uint = 1
        delete(v3, ("" + (("hello" + "\n\t") + "")))
        _ = fi29
        _ = v30
    }
    v31 := v1[1]
    v32 := []int{468}
    switch "foo bar" {
    case "":
        v33 := GenS5{F0: 1, F1: 42}
        _ = v33
    case "\n\t":
        v2 = (v2 * v2)
    default:
        switch -2147483648 {
        case -1:
            v2 -= (^(v2))
        case 416:
            var v34 string = ""
            _ = v34
        default:
            fmt.Println(v2, v2)
        }
    }
    if len(v32) > 0 {
        v32[0] = 0
    }
    var v35 bool = (!v31)
    v36 := GenS1{F0: 0, F1: (!(2 >= 42)), F2: 183}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v35
    _ = v36
}
