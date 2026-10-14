package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 int; F2 bool }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int64 }

func main() {
    var v1 float64 = (((-2.25 * 0.1) + 0.1) - 1e10)
    var v2 float64 = ((v1 - 0.0) + (((float64(3.14159) / float64(3.14159)) - 0.1) + 0.1))
    v3 := make(map[int]int)
    v4 := make(map[int]bool)
    var v5 uint8 = 42
    var v6 bool = true
    v1 += (v1 / v1)
    for fi7 := 0; fi7 < 12; fi7++ {
        for fi8 := 0; fi8 < 1; fi8++ {
            v5 = v5
            var v9 bool = (!(!false))
            v10 := GenS1{F0: 171}
            _ = fi8
            _ = v10
            _ = v9
        }
        v5 = v5
        for fi11 := 0; fi11 < 17; fi11++ {
            var v12 bool = ((!v6) && (-941 == -1))
            _ = fi11
            _ = v12
        }
        _ = fi7
    }
    if true {
        var v13 uint8 = (v5 - v5)
        _ = v13
    } else {
        v5 = 2
        var v14 uint32 = 65535
        switch "a" {
        case "\n\t":
            var v15 float64 = (1e10 / float64(((float32((0.1 - -2.25)) / 0.1) + (-2.25 - -2.25))))
            _ = v15
        case "":
            v16 := GenS1{F0: uint(v14)}
            _ = v16
        case "hello":
            for fi17 := 0; fi17 < 1; fi17++ {
                if v6 {
                    if v6 {
                        switch fi17 {
                        case 899:
                            var v18 uint16 = 42
                            _ = v18
                        case 4294967295:
                            for fi19 := 0; fi19 < 9; fi19++ {
                                v14 -= v14
                                v20 := GenS2{F0: 1, F1: (fi17 >> ((0) & 63)), F2: (9223372036854775807 < 611)}
                                var v21 float64 = float64(1e10)
                                _ = fi19
                                _ = v20
                                _ = v21
                            }
                        default:
                            v6 = v6
                        }
                    } else {
                        fmt.Println(v2, v2)
                        v22 := []int{(fi17 ^ fi17)}
                        _ = v22
                    }
                    if (18446744073709551615 < 42) {
                        v23 := []int{(fi17 ^ 383), fi17}
                        var v24 uint16 = 1
                        if v6 {
                            v5 = v5
                            var v25 int32 = -1
                            v2 -= 0.0
                            _ = v25
                        } else {
                            switch ("\n\t" + "a") {
                            case "a":
                                delete(v4, (fi17 & (fi17 % ((-9223372036854775808) | 1))))
                            case "":
                                var v26 int8 = -128
                                _ = v26
                            default:
                                v27 := GenS1{F0: 4294967295}
                                _ = v27
                            }
                            var v28 int = (fi17 - (fi17 + fi17))
                            _ = v28
                        }
                        _ = v23
                        _ = v24
                    }
                }
                v29 := []int{fi17}
                for fi30 := 0; fi30 < 7; fi30++ {
                    switch "hello" {
                    case "\n\t":
                        var v31 uint16 = 630
                        _ = v31
                    case "foo bar":
                        var v32 uint8 = v5
                        _ = v32
                    case "a":
                        var v33 string = "hello"
                        _ = v33
                    default:
                        if v6 {
                            v14 = (v14 / ((v14) | 1))
                            var v34 uint64 = 18446744073709551615
                            v35 := GenS1{F0: 4294967295}
                            _ = v34
                            _ = v35
                        }
                    }
                    _ = fi30
                }
                _ = fi17
                _ = v29
            }
        default:
            for fi36 := 0; fi36 < 13; fi36++ {
                for fi37 := 0; fi37 < 9; fi37++ {
                    var v38 int32 = 175
                    v39 := GenS3{F0: 9223372036854775807}
                    _ = fi37
                    _ = v38
                    _ = v39
                }
                _ = fi36
            }
        }
        _ = v14
    }
    v2 += (1.5 + v1)
    var v40 int64 = 1
    v41 := GenS4{F0: v40}
    var v42 uint64 = 2
    _ = len(v4)
    v2 -= 1e10
    if true {
        v43 := GenS1{F0: 933}
        for fi44 := 0; fi44 < 6; fi44++ {
            v45 := v4[fi44]
            _ = fi44
            _ = v45
        }
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
