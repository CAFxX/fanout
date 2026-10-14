package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 uint64 }
type GenS2 struct { F0 string; F1 int8 }
func genFn0(p0 int16, p1 int8) int {
    var v5 int = 42
    for fi6 := 0; fi6 < 12; fi6++ {
        if (0 < 66) {
            if (255 == 449) {
                p1 = p1
            }
            if false {
                fmt.Println(p1, v5)
            }
        } else {
            var v7 float32 = float32((3.14159 + (float64(1.5) / 1.0)))
            p1 += (p1 / ((1) | 1))
            _ = v7
        }
        _ = fi6
    }
    var v8 string = "hello"
    for fi9 := 0; fi9 < 17; fi9++ {
        var v10 int16 = (^(p0))
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v8
    return (v5 >> ((661) & 63))
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint64 = 0
    var v3 float64 = ((0.0 / 1.0) * 1e10)
    v4 := make(map[int]bool)
    switch 32768 {
    case 128:
        v11 := []int{}
        _ = v11
    }
    if (v3 != ((3.14159 / 3.14159) * float64(1.5))) {
        for fi12 := 0; fi12 < 1; fi12++ {
            var v13 uint16 = 1
            v14 := GenS1{F0: 1, F1: "foo bar", F2: (v2 | v2)}
            v15 := GenS1{F0: 454, F1: ("a" + "hello"), F2: (v2 << ((1) & 63))}
            _ = fi12
            _ = v13
            _ = v14
            _ = v15
        }
        v2 = 150
    } else {
        switch genFn0(32767, 0) {
        case 9223372036854775807:
            fmt.Println(v2)
        }
        for fi16 := 0; fi16 < 16; fi16++ {
            v3 = 0.1
            _ = fi16
        }
        fmt.Println(v2, v3)
    }
    fmt.Println(v3)
    switch 32767 {
    case 0:
        v17 := []int{genFn0(-1, 127), 32768, -2147483648, -915}
        _ = v17
    }
    switch 9223372036854775807 {
    case 0:
        switch -1 {
        case 32768:
            if (224 < 4294967295) {
                if (910 != -645) {
                    fmt.Println(v2, v2)
                    var v18 int32 = 0
                    _ = v18
                } else {
                    switch -499 {
                    case 2147483647:
                        switch -759 {
                        case 32768:
                            var v19 int16 = 1
                            _ = v19
                        }
                    case 32768:
                        v3 += v3
                    case 4294967295:
                        if false {
                            v3 -= v3
                            fmt.Println(v2, v2)
                            switch 42 {
                            case 223:
                                v3 += (v3 - (((v3 / v3) - -2.25) + float64(1e10)))
                            case 32767:
                                delete(v4, 32767)
                            default:
                                if (!(557 > 974)) {
                                    var v20 int = 32767
                                    switch ("foo bar" + "") {
                                    case "\n\t":
                                        v21 := GenS2{F0: "a", F1: 0}
                                        _ = v21
                                    case "hello":
                                        if (0 >= (v2 << ((2) & 63))) {
                                            var v22 uint32 = 42
                                            _ = v22
                                        }
                                    case "":
                                        var v23 uint64 = 583
                                        _ = v23
                                    default:
                                        v2 = 499
                                    }
                                    if false {
                                        var v24 uint32 = 42
                                        fmt.Println(v20, v3)
                                        if true {
                                            v20 = v20
                                        }
                                        _ = v24
                                    }
                                    _ = v20
                                }
                            }
                        }
                    }
                }
            } else {
                for fi25 := 0; fi25 < 16; fi25++ {
                    for fi26 := 0; fi26 < 6; fi26++ {
                        v1[65535] = (false || true)
                        fmt.Println(fi25)
                        v27 := GenS1{F0: 0, F1: "a", F2: v2}
                        _ = fi26
                        _ = v27
                    }
                    _ = fi25
                }
            }
        case -506:
            if (!(20 != -9223372036854775808)) {
                if (false && false) {
                    v2 = 0
                    delete(v1, 2)
                }
            }
        case 4294967295:
            v28 := []int{-2147483648, 2147483647, -9223372036854775808, 2147483647}
            _ = v28
        default:
            v1[2] = false
        }
    }
    var v29 float32 = 0.1
    for fi30 := 0; fi30 < 6; fi30++ {
        v2 -= (^(v2))
        _ = fi30
    }
    var v31 int16 = 0
    v3 = v3
    for fi32 := 0; fi32 < 17; fi32++ {
        var v33 float64 = 1.5
        v33 = (v3 - ((v3 - v33) * v3))
        var v34 uint8 = 78
        _ = fi32
        _ = v33
        _ = v34
    }
    v35 := GenS2{F0: "a", F1: 100}
    var v36 uint16 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v31
    _ = v35
    _ = v36
    _ = v4
}
