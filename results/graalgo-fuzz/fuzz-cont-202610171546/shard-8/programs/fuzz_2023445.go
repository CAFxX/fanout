package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 int32; F1 uint }
func genFn0(p0 string) int8 {
    switch "" {
    case "a":
        for fi5 := 0; fi5 < 12; fi5++ {
            var v6 int = -1
            if (0 <= 1) {
                var v7 string = p0
                var v8 uint16 = 0
                var v9 uint8 = 2
                _ = v7
                _ = v8
                _ = v9
            } else {
                v10 := []int{fi5, -2147483648}
                if true {
                    v11 := GenS1{F0: 127, F1: 32767}
                    _ = v11
                }
                _ = v10
            }
            _ = fi5
            _ = v6
        }
    case "foo bar":
        var v12 int16 = 0
        _ = v12
    case "":
        fmt.Println(p0)
    default:
        v13 := GenS1{F0: 0, F1: 95}
        _ = v13
    }
    for fi14 := 0; fi14 < 8; fi14++ {
        switch 127 {
        case 1:
            var v15 string = ("" + p0)
            _ = v15
        }
        p0 = "foo bar"
        _ = fi14
    }
    switch 2147483647 {
    case 9223372036854775807:
        fmt.Println(p0)
    case 579:
        var v16 uint8 = 0
        _ = v16
    case 438:
        var v17 uint8 = 1
        _ = v17
    default:
        v18 := GenS2{F0: 1, F1: 385}
        _ = v18
    }
    _ = p0
    return 42
}
func genFn1(p0 int64, p1 uint32, p2 uint64) uint16 {
    v19 := GenS2{F0: 42, F1: 2}
    v20 := []int{-750, 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    _ = v20
    return 1
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int8 = -1
    v3 := make(map[uint32]string)
    var v4 int16 = 1000
    v2 -= (v2 / ((v2) | 1))
    var v21 int8 = -128
    fmt.Println(v21, v2, v4)
    var v22 uint16 = 0
    var v23 uint64 = 1
    v1[("hello" + ("foo bar" + ""))] = 882
    v22 += v22
    delete(v3, 824)
    var v24 int64 = -87
    switch 128 {
    case 42:
        for fi25 := 0; fi25 < 4; fi25++ {
            v24 -= v24
            v26 := []int{fi25, 32767, 128}
            fmt.Println(v22)
            _ = fi25
            _ = v26
        }
    case 1:
        var v27 float32 = float32(-2.25)
        _ = v27
    default:
        fmt.Println(v2, v22, v21)
    }
    v28 := []int{0, -383, 43, -9223372036854775808}
    switch -1 {
    case -2147483648:
        switch -479 {
        case 15:
            var v29 uint64 = 0
            _ = v29
        case 32767:
            v30 := GenS2{F0: 42, F1: uint(v22)}
            _ = v30
        default:
            if (2 < 1) {
                switch -9223372036854775808 {
                case 32767:
                    v2 = 127
                default:
                    if true {
                        if false {
                            for fi31 := 0; fi31 < 13; fi31++ {
                                for fi32 := 0; fi32 < 9; fi32++ {
                                    v21 = genFn0((("foo bar" + "hello") + "hello"))
                                    if false {
                                        fmt.Println(v2)
                                    } else {
                                        v22 += v22
                                        if len(v28) >= 2 {
                                            v28 = v28[1:len(v28)]
                                        }
                                    }
                                    v23 += v23
                                    _ = fi32
                                }
                                v33 := GenS1{F0: 2, F1: genFn1(v24, 42, 42)}
                                _ = fi31
                                _ = v33
                            }
                            fmt.Println(v2, v24, v2)
                        }
                        for fi34 := 0; fi34 < 6; fi34++ {
                            v24 -= 856
                            var v35 int16 = (v4 & (v4 << ((0) & 63)))
                            var v36 float32 = 3.14159
                            _ = fi34
                            _ = v35
                            _ = v36
                        }
                    }
                }
            } else {
                delete(v3, 210)
                if (18446744073709551615 < 65535) {
                    v37 := v1["hello"]
                    var v38 uint = 255
                    _ = v37
                    _ = v38
                }
            }
        }
    }
    switch "foo bar" {
    case "a":
        v1[(("hello" + "foo bar") + "foo bar")] = 42
    }
    var v39 int32 = 386
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v28
    _ = v3
    _ = v39
    _ = v4
}
