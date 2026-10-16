package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 uint16 }
type GenS2 struct { F0 uint8 }
func genFn0() uint {
    v4 := GenS1{F0: 42, F1: -9223372036854775808, F2: 255}
    if true {
        v5 := []int{32768, 32767, -155, -2147483648}
        _ = v5
    }
    var v6 uint64 = 1
    var v7 float32 = 0.0
    _ = v4
    _ = v6
    _ = v7
    return 1
}
func genFn1() string {
    if (false && true) {
        if (341 < 42) {
            var v8 float32 = 0.0
            v8 += (v8 - v8)
            _ = v8
        }
        for fi9 := 0; fi9 < 11; fi9++ {
            v10 := []int{}
            _ = fi9
            _ = v10
        }
        v11 := []int{}
        _ = v11
    } else {
        switch 32768 {
        case 0:
            var v12 int8 = -128
            _ = v12
        case -448:
            switch 0 {
            case 1:
                v13 := []int{314, 4294967295, -9223372036854775808}
                _ = v13
            case 73:
                for fi14 := 0; fi14 < 4; fi14++ {
                    fmt.Println(fi14)
                    if (32768 <= (fi14 << ((488) & 63))) {
                        var v15 int32 = 2147483647
                        var v16 bool = false
                        _ = v15
                        _ = v16
                    } else {
                        switch fi14 {
                        case 42:
                            var v17 uint32 = 117
                            _ = v17
                        }
                        fmt.Println(fi14)
                        if (-26 <= -192) {
                            var v18 int = 36
                            if (v18 <= fi14) {
                                v19 := GenS1{F0: 1000, F1: 9223372036854775807, F2: 65535}
                                for fi20 := 0; fi20 < 9; fi20++ {
                                    var v21 string = (("a" + "a") + "foo bar")
                                    switch "hello" {
                                    case "hello":
                                        switch "a" {
                                        case "":
                                            if true {
                                                v22 := []int{}
                                                v21 = "\n\t"
                                                var v23 uint16 = 917
                                                _ = v22
                                                _ = v23
                                            }
                                        case "a":
                                            fmt.Println(v18)
                                        case "\n\t":
                                            var v24 uint8 = 0
                                            _ = v24
                                        }
                                    }
                                    _ = fi20
                                    _ = v21
                                }
                                _ = v19
                            } else {
                                switch (fi14 & 4294967295) {
                                case 32768:
                                    var v25 int = (fi14 / ((fi14) | 1))
                                    _ = v25
                                case 32767:
                                    v18 = v18
                                }
                            }
                            for fi26 := 0; fi26 < 16; fi26++ {
                                var v27 uint8 = 1
                                v28 := []int{2147483647, (^(fi14)), (^(v18)), (^(fi26))}
                                var v29 uint32 = 42
                                fmt.Println(v18)
                                fmt.Println(fi26, v29)
                                _ = fi26
                                _ = v27
                                _ = v28
                                _ = v29
                            }
                            _ = v18
                        }
                    }
                    _ = fi14
                }
            }
        default:
            v30 := []int{}
            _ = v30
        }
        var v31 int16 = 1
        _ = v31
    }
    var v32 uint32 = 0
    var v33 int16 = 42
    v34 := []int{811, 32767, 0}
    fmt.Println(v33, v32)
    _ = v32
    _ = v33
    _ = v34
    return ("" + "foo bar")
}

func main() {
    v1 := []uint32{42}
    var v2 uint64 = 0
    var v3 float64 = (1e10 - (1e10 / 1.0))
    var v35 uint16 = 281
    var v36 uint = 18446744073709551615
    v36 -= 65535
    fmt.Println(v3, v3, v35)
    fmt.Println(v35, v2, v35)
    var v37 int = 960
    v38 := GenS1{F0: -742, F1: -714, F2: uint16(v36)}
    for fi39 := 0; fi39 < 7; fi39++ {
        fmt.Println(v3, v37)
        _ = fi39
    }
    _ = len(v1)
    v2 -= v2
    var v40 bool = true
    v37 += 0
    switch v37 {
    case 42:
        var v41 uint64 = (v2 % ((v2) | 1))
        _ = v41
    default:
        if (255 == 2) {
            if ((v3 > (v3 / v3)) && (v2 != v2)) {
                v37 = v37
                v42 := GenS1{F0: -990, F1: 42, F2: v35}
                v2 = v2
                _ = v42
            } else {
                for fi43 := 0; fi43 < 16; fi43++ {
                    v37 -= (fi43 >> ((v36) & 63))
                    _ = fi43
                }
                fmt.Println(v40, v35)
            }
        } else {
            v40 = (!((10 == 255) || false))
            v2 = 0
            if (v40 && true) {
                var v44 uint64 = (v2 - (^(v2)))
                v1 = append(v1, 554)
                _ = v44
            } else {
                v36 = 2
            }
        }
    }
    if true {
        v45 := GenS2{F0: 42}
        _ = v45
    }
    var v46 int64 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v40:", v40)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v40
    _ = v46
}
