package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 uint8 }
type GenS2 struct { F0 bool }

func main() {
    var v1 float32 = (float32(0.1) * (0.1 + 3.14159))
    var v2 float32 = ((v1 + float32(float64(0.0))) * (-2.25 - (3.14159 * ((v1 + v1) * v1))))
    var v3 uint64 = 0
    v4 := make(map[string]uint32)
    v5 := make(map[uint32]string)
    switch 32767 {
    case 334:
        for fi6 := 0; fi6 < 8; fi6++ {
            v7 := []int{}
            v2 = (0.0 - (v1 * 0.0))
            _ = fi6
            _ = v7
        }
    case 128:
        if ((1 > v3) || false) {
            var v8 float32 = (((((v2 - (v2 * v2)) / v1) / v2) / (v1 + float32((3.14159 - float64(v1))))) * v1)
            v2 = v8
            _ = v8
        } else {
            v1 += (v2 / v1)
            for fi9 := 0; fi9 < 13; fi9++ {
                v2 += 1.5
                v1 = (-2.25 / float32(0.1))
                if false {
                    switch "" {
                    case "foo bar":
                        if (127 != 255) {
                            v10 := GenS1{F0: -32768, F1: 180, F2: 9}
                            v3 = (v3 * v3)
                            _ = v10
                        } else {
                            switch fi9 {
                            case -2147483648:
                                v11 := GenS1{F0: 0, F1: 24, F2: 2}
                                _ = v11
                            }
                            if (fi9 == fi9) {
                                v12 := []int{-1, (-(fi9)), (-(fi9))}
                                var v13 uint32 = 624
                                _ = v12
                                _ = v13
                            } else {
                                for fi14 := 0; fi14 < 18; fi14++ {
                                    if ((18446744073709551615 > 364) || (false && false)) {
                                        v15 := []int{32768, -788, fi9}
                                        _ = v15
                                    }
                                    if (4294967295 == 202) {
                                        var v16 uint64 = 444
                                        _ = v16
                                    }
                                    _ = fi14
                                }
                            }
                        }
                    case "hello":
                        v17 := v5[1]
                        _ = v17
                    case "":
                        v18 := GenS1{F0: 1, F1: 160, F2: 0}
                        _ = v18
                    }
                    for fi19 := 0; fi19 < 10; fi19++ {
                        switch ("hello" + ("\n\t" + ("\n\t" + "\n\t"))) {
                        case "\n\t":
                            fmt.Println(v3, v2, fi19)
                        case "":
                            v3 += 1
                        case "foo bar":
                            if true {
                                if (!(((v3 ^ v3) != (v3 - (v3 - (^(v3))))) && false)) {
                                    if false {
                                        var v20 int8 = -128
                                        v21 := []int{1}
                                        _ = v20
                                        _ = v21
                                    }
                                }
                                delete(v5, 227)
                            }
                        }
                        v22 := []int{-1, (fi19 << ((862) & 63)), 4294967295}
                        fmt.Println(v2, fi19, v3)
                        v2 = (0.0 / v2)
                        v2 -= float32(1.5)
                        _ = fi19
                        _ = v22
                    }
                }
                _ = fi9
            }
            v3 -= v3
        }
    case 127:
        if false {
            switch "\n\t" {
            case "":
                switch "" {
                case "a":
                    var v23 uint8 = 2
                    _ = v23
                case "\n\t":
                    if false {
                        v24 := []int{42, -728, 32767}
                        var v25 int8 = 100
                        _ = v24
                        _ = v25
                    }
                default:
                    fmt.Println(v1, v2)
                }
            case "hello":
                v1 -= 1.5
            case "\n\t":
                var v26 bool = (42 == 2147483647)
                _ = v26
            }
            if (-1 >= 0) {
                for fi27 := 0; fi27 < 18; fi27++ {
                    for fi28 := 0; fi28 < 6; fi28++ {
                        var v29 float64 = (float64(v1) + (1.5 + (3.14159 + 1.5)))
                        _ = fi28
                        _ = v29
                    }
                    _ = fi27
                }
                switch 128 {
                case 854:
                    v30 := GenS2{F0: true}
                    _ = v30
                case 4294967295:
                    var v31 uint = 255
                    _ = v31
                default:
                    var v32 bool = true
                    _ = v32
                }
            }
        }
    default:
        v2 = 1.5
    }
    switch 4294967295 {
    case 128:
        v33 := v4["foo bar"]
        _ = v33
    case -1:
        var v34 uint64 = v3
        _ = v34
    case 42:
        var v35 uint32 = 1
        _ = v35
    default:
        v36 := GenS2{F0: false}
        _ = v36
    }
    v1 += v1
    var v37 int64 = 1
    v38 := GenS2{F0: true}
    var v39 float64 = 3.14159
    var v40 uint16 = 632
    var v41 string = "hello"
    var v42 bool = true
    v43 := GenS1{F0: 0, F1: 124, F2: 127}
    for fi44 := 0; fi44 < 11; fi44++ {
        fmt.Println(v39, v3, v40)
        var v45 float32 = v1
        _ = fi44
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v5
}
