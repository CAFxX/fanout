package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 uint; F1 int8 }
func genFn0(p0 bool, p1 int64) uint8 {
    v4 := []int{}
    v5 := GenS1{F0: 32767, F1: 42}
    p1 = (p1 >> ((65535) & 63))
    p1 -= p1
    var v6 uint = 2
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 66
}

func main() {
    var v1 int16 = 1000
    var v2 uint = 2
    var v3 float64 = 0.1
    var v7 int16 = (-(v1))
    for fi8 := 0; fi8 < 13; fi8++ {
        fmt.Println(fi8, v2)
        _ = fi8
    }
    v9 := []int{-2147483648, -403}
    if ((v7 ^ (v1 + (v1 ^ -32768))) != v1) {
        if len(v9) > 0 {
            v9[0] = int(v2)
        }
    } else {
        if len(v9) > 0 {
            v9[0] = 1
        }
        v9 = append(v9, 383)
    }
    var v10 bool = true
    for fi11 := 0; fi11 < 15; fi11++ {
        v12 := GenS1{F0: 1000, F1: (v7 << ((v2) & 63))}
        var v13 uint8 = 0
        v2 += v2
        _ = fi11
        _ = v12
        _ = v13
    }
    fmt.Println(v7, v7)
    switch 9223372036854775807 {
    case 2147483647:
        switch ("" + "a") {
        case "foo bar":
            fmt.Println(v1, v2, v10)
        case "":
            fmt.Println(v7)
        case "a":
            switch -33 {
            case -9223372036854775808:
                if true {
                    if v10 {
                        var v14 float32 = 0.1
                        if (!true) {
                            v15 := GenS1{F0: (v7 - 42), F1: (v1 / ((v7) | 1))}
                            var v16 uint = v2
                            _ = v15
                            _ = v16
                        } else {
                            for fi17 := 0; fi17 < 16; fi17++ {
                                if (true && v10) {
                                    fmt.Println(fi17, v10)
                                }
                                _ = fi17
                            }
                            if (2 > 296) {
                                v9 = append(v9, -9223372036854775808)
                            } else {
                                v1 = 32767
                            }
                        }
                        for fi18 := 0; fi18 < 19; fi18++ {
                            v19 := GenS1{F0: 941, F1: -1}
                            fmt.Println(v10, v1)
                            for fi20 := 0; fi20 < 5; fi20++ {
                                fmt.Println(v1, fi18, v14)
                                if len(v9) >= 2 {
                                    v9 = v9[1:len(v9)]
                                }
                                _ = fi20
                            }
                            _ = fi18
                            _ = v19
                        }
                        _ = v14
                    }
                    fmt.Println(v3, v2)
                    for fi21 := 0; fi21 < 7; fi21++ {
                        if v10 {
                            for fi22 := 0; fi22 < 5; fi22++ {
                                v23 := GenS2{F0: v2, F1: 100}
                                _ = fi22
                                _ = v23
                            }
                            switch "a" {
                            case "a":
                                v2 = v2
                            case "foo bar":
                                fmt.Println(v3)
                            case "hello":
                                if false {
                                    var v24 float32 = 0.1
                                    _ = v24
                                }
                            }
                            if (float64(1e10) >= (1.5 / v3)) {
                                v9 = append(v9, 32768)
                            }
                        } else {
                            v25 := GenS1{F0: (v7 << (((v2 << ((v2) & 63))) & 63)), F1: (v7 / ((v1) | 1))}
                            _ = v25
                        }
                        v10 = (v10 || (v7 <= (v7 >> (((^(v2))) & 63))))
                        for fi26 := 0; fi26 < 3; fi26++ {
                            v10 = v10
                            _ = fi26
                        }
                        _ = fi21
                    }
                }
            case -1:
                for fi27 := 0; fi27 < 7; fi27++ {
                    if len(v9) > 0 {
                        v9[0] = fi27
                    }
                    _ = fi27
                }
            }
        }
    case -1:
        var v28 int = 2147483647
        _ = v28
    case 32768:
        if (((!v10) && v10) || v10) {
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            fmt.Println(v1, v7, v1)
        } else {
            var v29 uint32 = 555
            v29 -= (v29 << (((v2 << ((v2) & 63))) & 63))
            v30 := GenS1{F0: 1, F1: 42}
            _ = v29
            _ = v30
        }
    default:
        if true {
            v31 := GenS2{F0: v2, F1: 42}
            _ = v31
        }
    }
    if (!v10) {
        fmt.Println(v3, v3, v1)
    }
    v10 = (!((1 <= v2) && false))
    var v32 int64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v32
    _ = v7
    _ = v9
}
