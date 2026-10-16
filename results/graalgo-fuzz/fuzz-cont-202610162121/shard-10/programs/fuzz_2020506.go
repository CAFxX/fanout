package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 uint }
type GenS2 struct { F0 uint8; F1 int32; F2 int64 }
type GenS3 struct { F0 uint32 }
func genFn0() int8 {
    v4 := GenS1{F0: 42, F1: true, F2: 18446744073709551615}
    if (0 != 1) {
        var v5 uint16 = 2
        _ = v5
    }
    var v6 string = "a"
    var v7 int8 = -128
    _ = v4
    _ = v6
    _ = v7
    return v7
}

func main() {
    var v1 uint = 406
    v2 := make(map[string]int)
    var v3 int32 = 42
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 bool = true
        var v10 int32 = v3
        _ = fi8
        _ = v10
        _ = v9
    }
    v1 = (v1 * v1)
    v11 := GenS1{F0: 9223372036854775807, F1: false, F2: (v1 - (^(v1)))}
    v3 += (v3 % (((v3 | 2147483647)) | 1))
    fmt.Println(v3)
    fmt.Println(v3)
    var v12 float32 = (3.14159 / -2.25)
    v13 := GenS2{F0: 38, F1: (v3 % ((v3) | 1)), F2: -1}
    var v14 uint8 = 128
    switch "\n\t" {
    case "\n\t":
        for fi15 := 0; fi15 < 20; fi15++ {
            fmt.Println(v1, v3, fi15)
            v12 = 1e10
            switch "foo bar" {
            case "foo bar":
                if false {
                    var v16 int32 = 1
                    if (-128 != 1) {
                        var v17 uint = v1
                        v12 = ((v12 - 0.1) + (v12 - (v12 + (-2.25 + v12))))
                        fmt.Println(fi15, v17, v3)
                        _ = v17
                    } else {
                        v16 = -663
                        v12 -= (v12 / ((v12 - (1.5 + (-2.25 - ((v12 / (((3.14159 * v12) - (1e10 / (v12 / ((v12 * (((3.14159 + v12) * (0.1 + (v12 / (((v12 * v12) * 0.1) + v12)))) - (0.0 * 1e10))) + -2.25)))) / -2.25)) / ((v12 / v12) + v12))))) / v12))
                        var v18 string = "a"
                        _ = v18
                    }
                    _ = v16
                }
            case "":
                for fi19 := 0; fi19 < 18; fi19++ {
                    switch (("foo bar" + ("" + "\n\t")) + "foo bar") {
                    case "\n\t":
                        if (127 <= genFn0()) {
                            v14 = 127
                            v20 := []int{(fi19 >> (((v1 / (((^(v1))) | 1))) & 63))}
                            _ = v20
                        }
                    case "":
                        _ = len(v2)
                    default:
                        var v21 int16 = 593
                        _ = v21
                    }
                    v12 -= v12
                    _ = fi19
                }
            default:
                var v22 int32 = (v3 & (v3 ^ -391))
                _ = v22
            }
            _ = fi15
        }
    default:
        var v23 int = 128
        _ = v23
    }
    var v24 float64 = 0.1
    for fi25 := 0; fi25 < 17; fi25++ {
        switch "" {
        case "":
            _ = len(v2)
        case "hello":
            v1 = 18446744073709551615
        case "foo bar":
            fmt.Println(fi25, v24)
        }
        switch "a" {
        case "foo bar":
            var v26 int16 = 411
            _ = v26
        case "hello":
            if ("\n\t" <= "foo bar") {
                var v27 float64 = v24
                if false {
                    v2["\n\t"] = fi25
                    fmt.Println(v3, v12, v12)
                } else {
                    for fi28 := 0; fi28 < 9; fi28++ {
                        if false {
                            switch (fi25 % ((fi28) | 1)) {
                            case 9223372036854775807:
                                v29 := GenS2{F0: (v14 / ((147) | 1)), F1: v3, F2: 667}
                                _ = v29
                            case -2147483648:
                                switch 32767 {
                                case -362:
                                    var v30 int8 = 58
                                    _ = v30
                                case 2147483647:
                                    v1 = (v1 & v1)
                                case 127:
                                    v1 += (v1 << ((v1) & 63))
                                default:
                                    if (65535 == 32767) {
                                        var v31 bool = true
                                        v32 := GenS2{F0: 0, F1: v3, F2: -586}
                                        _ = v31
                                        _ = v32
                                    }
                                }
                            default:
                                v33 := GenS1{F0: 9223372036854775807, F1: false, F2: (v1 + v1)}
                                v14 = uint8(v1)
                                var v34 float64 = v27
                                _ = v33
                                _ = v34
                            }
                        } else {
                            v14 = (^(v14))
                            v3 = (v3 & v3)
                            v27 = (-2.25 - (v27 / v27))
                        }
                        v14 = v14
                        _ = fi28
                    }
                }
                v27 += (v27 - 0.1)
                _ = v27
            } else {
                v14 = v14
            }
        case "":
            v35 := GenS3{F0: 849}
            _ = v35
        default:
            v1 += v1
        }
        _ = fi25
    }
    if true {
        var v36 uint32 = 48
        v37 := v2["\n\t"]
        _ = v36
        _ = v37
    }
    var v38 string = "\n\t"
    var v39 bool = ((true || true) || false)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v24
    _ = v3
    _ = v38
    _ = v39
}
