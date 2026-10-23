package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 int; F1 bool }
func genFn0(p0 int8) int {
    switch "hello" {
    case "":
        p0 = 127
    }
    switch 1 {
    case 32767:
        if true {
            p0 -= (p0 & (p0 - -4))
            fmt.Println(p0)
        } else {
            if (1 != 32767) {
                p0 -= (^(p0))
                for fi4 := 0; fi4 < 8; fi4++ {
                    v5 := []int{fi4}
                    switch "foo bar" {
                    case "a":
                        fmt.Println(p0, p0)
                    case "foo bar":
                        v6 := GenS1{F0: 76, F1: 42}
                        _ = v6
                    case "\n\t":
                        fmt.Println(fi4)
                    default:
                        if (!(!true)) {
                            p0 -= (p0 | (p0 | (p0 + (p0 - p0))))
                            var v7 uint8 = 128
                            _ = v7
                        } else {
                            switch 32768 {
                            case 2147483647:
                                p0 += p0
                            default:
                                p0 += (p0 + 0)
                            }
                            v8 := GenS1{F0: 933, F1: 127}
                            for fi9 := 0; fi9 < 15; fi9++ {
                                p0 -= p0
                                _ = fi9
                            }
                            _ = v8
                        }
                    }
                    _ = fi4
                    _ = v5
                }
                v10 := []int{128}
                _ = v10
            }
            if true {
                var v11 uint16 = 565
                var v12 uint = 0
                for fi13 := 0; fi13 < 17; fi13++ {
                    for fi14 := 0; fi14 < 18; fi14++ {
                        v11 = (v11 & v11)
                        _ = fi14
                    }
                    _ = fi13
                }
                _ = v11
                _ = v12
            } else {
                v15 := []int{-764, 2147483647}
                var v16 uint = 18446744073709551615
                switch "" {
                case "":
                    _ = len(v15)
                }
                _ = v15
                _ = v16
            }
            var v17 uint64 = 1
            _ = v17
        }
    default:
        var v18 int = 746
        _ = v18
    }
    _ = p0
    return 128
}
func genFn1(p0 int32, p1 int16) uint16 {
    if true {
        switch int(p1) {
        case 2147483647:
            var v19 bool = false
            _ = v19
        }
        fmt.Println(p0)
    }
    for fi20 := 0; fi20 < 6; fi20++ {
        if (!(p1 <= p1)) {
            var v21 string = "\n\t"
            for fi22 := 0; fi22 < 6; fi22++ {
                v21 = "\n\t"
                _ = fi22
            }
            _ = v21
        }
        _ = fi20
    }
    fmt.Println(p0)
    for fi23 := 0; fi23 < 11; fi23++ {
        p0 += p0
        fmt.Println(p1, p0)
        p1 = (p1 + p1)
        _ = fi23
    }
    _ = p0
    _ = p1
    return 65535
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]uint32)
    var v3 uint8 = 128
    var v24 float64 = 0.1
    fmt.Println(v3, v3)
    fmt.Println(v24, v3)
    if false {
        _ = len(v1)
        v25 := GenS1{F0: 312, F1: 97}
        v26 := GenS2{F0: 641, F1: false}
        _ = v25
        _ = v26
    } else {
        for fi27 := 0; fi27 < 12; fi27++ {
            if true {
                v24 -= v24
            }
            v28 := GenS1{F0: 585, F1: (v3 * (v3 - v3))}
            v3 -= v3
            _ = fi27
            _ = v28
        }
        v3 = (v3 << ((42) & 63))
    }
    fmt.Println(v24)
    for fi29 := 0; fi29 < 17; fi29++ {
        var v30 int32 = 0
        v31 := GenS2{F0: fi29, F1: (65535 > 0)}
        _ = fi29
        _ = v30
        _ = v31
    }
    var v32 int8 = 42
    v32 -= (v32 / ((v32) | 1))
    v33 := []int{}
    v3 = v3
    switch ((((("" + ((("" + "hello") + "a") + "")) + "") + "foo bar") + ("foo bar" + "a")) + ("foo bar" + "")) {
    case "foo bar":
        v3 = v3
    case "\n\t":
        var v34 float32 = (-2.25 + 0.1)
        _ = v34
    }
    if true {
        var v35 int = -9223372036854775808
        var v36 int8 = 0
        var v37 uint32 = 4294967295
        _ = v35
        _ = v36
        _ = v37
    } else {
        fmt.Println(v32)
        if true {
            for fi38 := 0; fi38 < 4; fi38++ {
                if false {
                    if true {
                        var v39 bool = ((286 == 852) && (2 == 2))
                        _ = v39
                    }
                }
                var v40 string = "\n\t"
                _ = fi38
                _ = v40
            }
        } else {
            _ = len(v33)
            var v41 float32 = 0.0
            if false {
                v42 := GenS1{F0: 2147483648, F1: 1}
                v3 = v3
                fmt.Println(v24, v41, v24)
                _ = v42
            }
            _ = v41
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v33:", len(v33), v33)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v32
    _ = v33
}
