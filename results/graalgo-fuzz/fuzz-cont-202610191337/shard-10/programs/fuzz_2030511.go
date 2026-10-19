package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 uint64 }
type GenS2 struct { F0 uint; F1 bool; F2 int }
type GenS3 struct { F0 uint; F1 int8 }
type GenS4 struct { F0 int32 }
func genFn0(p0 uint, p1 uint32) uint32 {
    fmt.Println(p1, p1)
    v4 := GenS1{F0: 42, F1: 32768, F2: 0}
    p1 -= 4294967295
    v5 := []int{}
    v6 := GenS1{F0: 42, F1: 42, F2: 742}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return (p1 - (p1 - 2147483648))
}

func main() {
    var v1 int32 = 2147483647
    v2 := make(map[uint32]string)
    var v3 float64 = (-2.25 + 1.5)
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 string = ("foo bar" + ("a" + "\n\t"))
        _ = fi7
        _ = v8
    }
    fmt.Println(v1, v3)
    var v9 uint32 = 0
    for fi10 := 0; fi10 < 18; fi10++ {
        var v11 float64 = v3
        var v12 uint16 = 2
        var v13 bool = false
        _ = fi10
        _ = v11
        _ = v12
        _ = v13
    }
    if (((3.14159 / 3.14159) - (float32(float64(float32((((v3 / v3) + ((v3 * v3) - v3)) + ((v3 - v3) + 1e10))))) + (-2.25 - (((-2.25 * 1.5) * 1e10) + (1.5 / 1e10))))) == 0.1) {
        if ("foo bar" < "\n\t") {
            for fi14 := 0; fi14 < 5; fi14++ {
                v1 = (v1 % (((v1 ^ (v1 + (v1 << ((42) & 63))))) | 1))
                var v15 uint8 = 42
                _ = fi14
                _ = v15
            }
        } else {
            var v16 int32 = 42
            switch "\n\t" {
            case "":
                v16 = (v16 % (((v16 + (v16 % ((v1) | 1)))) | 1))
            case "a":
                if (true && (255 != 1)) {
                    fmt.Println(v16, v1)
                    v2[genFn0(4294967295, genFn0(0, (v9 - v9)))] = ((("foo bar" + ("a" + "hello")) + "a") + ("a" + "a"))
                    v17 := GenS2{F0: 2, F1: true, F2: 1}
                    _ = v17
                }
            case "foo bar":
                switch 243 {
                case 32768:
                    switch 98 {
                    case 128:
                        var v18 uint8 = 0
                        _ = v18
                    }
                default:
                    v19 := []int{}
                    _ = v19
                }
            default:
                var v20 uint64 = 42
                _ = v20
            }
            _ = v16
        }
        switch "foo bar" {
        case "a":
            delete(v2, v9)
        default:
            if ((565 > 18446744073709551615) && false) {
                var v21 uint8 = 136
                v22 := GenS2{F0: uint(v9), F1: (!true), F2: -163}
                fmt.Println(v21, v1)
                _ = v21
                _ = v22
            } else {
                v23 := v2[448]
                if false {
                    for fi24 := 0; fi24 < 3; fi24++ {
                        v25 := []int{}
                        switch -1 {
                        case 1:
                            for fi26 := 0; fi26 < 18; fi26++ {
                                var v27 int64 = 9223372036854775807
                                var v28 float32 = (1e10 * 0.0)
                                delete(v2, (^(v9)))
                                _ = fi26
                                _ = v27
                                _ = v28
                            }
                        default:
                            if len(v25) >= 2 {
                                v25 = v25[1:len(v25)]
                            }
                        }
                        _ = fi24
                        _ = v25
                    }
                    for fi29 := 0; fi29 < 11; fi29++ {
                        for fi30 := 0; fi30 < 14; fi30++ {
                            fmt.Println(v3)
                            v9 = v9
                            if ((v9 >> ((0) & 63)) != 42) {
                                var v31 bool = (675 < 0)
                                v31 = true
                                fmt.Println(fi30, v1, fi30)
                                _ = v31
                            } else {
                                v32 := GenS1{F0: 302, F1: 128, F2: 511}
                                _ = v32
                            }
                            _ = fi30
                        }
                        fmt.Println(v9, fi29)
                        _ = fi29
                    }
                } else {
                    v33 := GenS2{F0: 314, F1: (false && true), F2: 520}
                    if true {
                        var v34 float32 = (1.5 * (0.1 * 3.14159))
                        v3 = (v3 * (float64(0.0) * -2.25))
                        _ = v34
                    }
                    _ = v33
                }
                for fi35 := 0; fi35 < 11; fi35++ {
                    v23 = ("a" + v23)
                    fmt.Println(v3, v9)
                    v1 -= (v1 >> ((1) & 63))
                    _ = fi35
                }
                _ = v23
            }
        }
        for fi36 := 0; fi36 < 8; fi36++ {
            for fi37 := 0; fi37 < 18; fi37++ {
                switch "hello" {
                case "":
                    if (false || (true && false)) {
                        v38 := GenS1{F0: int64(v1), F1: (fi36 - (fi36 | fi37)), F2: 0}
                        var v39 int = fi37
                        var v40 uint16 = 490
                        _ = v38
                        _ = v39
                        _ = v40
                    }
                default:
                    v3 += v3
                }
                if false {
                    v41 := GenS3{F0: 2, F1: 42}
                    v42 := GenS2{F0: 143, F1: false, F2: 9223372036854775807}
                    var v43 uint16 = 42
                    _ = v41
                    _ = v42
                    _ = v43
                }
                _ = fi37
            }
            _ = fi36
        }
    } else {
        v44 := []int{315}
        _ = v44
    }
    v45 := GenS4{F0: (v1 ^ v1)}
    var v46 string = ((("" + ("a" + "\n\t")) + "hello") + "a")
    var v47 int16 = 1
    if (-1 > v1) {
        var v48 uint16 = 1
        _ = v48
    }
    v49 := []int{127, 127}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v49:", len(v49), v49)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v45
    _ = v46
    _ = v47
    _ = v49
    _ = v9
}
