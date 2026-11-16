package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 int }
type GenS2 struct { F0 int64; F1 uint64; F2 int16 }
type GenS3 struct { F0 uint; F1 uint64 }
type GenS4 struct { F0 int; F1 int16 }
type GenS5 struct { F0 int8 }
func genFn0(p0 uint8) uint8 {
    var v7 uint16 = 255
    v8 := GenS1{F0: 2147483647, F1: 255, F2: -664}
    var v9 int = 9223372036854775807
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return p0
}

func main() {
    var v1 int32 = 42
    v2 := make(map[string]string)
    var v3 uint = 18446744073709551615
    var v4 int32 = (v1 >> (((v3 % (((v3 + 0)) | 1))) & 63))
    var v5 uint16 = 967
    var v6 int16 = 0
    for fi10 := 0; fi10 < 10; fi10++ {
        v11 := GenS1{F0: (fi10 + (fi10 * fi10)), F1: (^(v3)), F2: fi10}
        _ = fi10
        _ = v11
    }
    var v12 int = 42
    if false {
        v13 := GenS2{F0: 9223372036854775807, F1: 2, F2: (v6 + (v6 & 32767))}
        _ = v13
    } else {
        var v14 float32 = 0.0
        _ = v14
    }
    v15 := v2["\n\t"]
    for fi16 := 0; fi16 < 2; fi16++ {
        v4 = 2147483647
        var v17 int64 = 9223372036854775807
        var v18 int = 127
        _ = fi16
        _ = v17
        _ = v18
    }
    v15 = v15
    var v19 int8 = -53
    v3 += v3
    var v20 int32 = (v4 / (((v1 >> ((v3) & 63))) | 1))
    var v21 float64 = 3.14159
    switch v12 {
    case 1:
        v15 = "a"
    case -9223372036854775808:
        var v22 int = int(v6)
        _ = v22
    }
    for fi23 := 0; fi23 < 2; fi23++ {
        fmt.Println(fi23, v4, fi23)
        _ = fi23
    }
    fmt.Println(v4, v4, v20)
    switch int(v6) {
    case -369:
        if ((!false) && false) {
            if (false || (v19 > (v19 >> ((1) & 63)))) {
                var v24 int = 276
                _ = v24
            }
            fmt.Println(v20, v4)
            v4 -= (-(v1))
        } else {
            var v25 int8 = v19
            _ = v25
        }
    case 982:
        var v26 int64 = -9223372036854775808
        _ = v26
    case 4294967295:
        if true {
            v27 := []int{-237, v12, -367}
            _ = v27
        } else {
            v5 -= v5
            v28 := GenS3{F0: v3, F1: 23}
            if (false || (false && ((1 <= 1) && (uint16(v3) != (v5 / (((v5 ^ 2)) | 1)))))) {
                v29 := GenS2{F0: -140, F1: 2, F2: v6}
                _ = v29
            } else {
                var v30 float32 = 3.14159
                _ = v30
            }
            _ = v28
        }
    default:
        if false {
            v2[(((((v15 + "a") + "\n\t") + "foo bar") + v15) + "hello")] = v15
            switch v15 {
            case "":
                for fi31 := 0; fi31 < 12; fi31++ {
                    v6 = -921
                    _ = fi31
                }
            case "\n\t":
                fmt.Println(v20)
            case "foo bar":
                if false {
                    v5 += uint16(v3)
                    v32 := []int{v12}
                    _ = v32
                } else {
                    if (v12 == 962) {
                        if true {
                            v33 := GenS4{F0: v12, F1: (-(v6))}
                            _ = len(v2)
                            var v34 int16 = -1
                            _ = v33
                            _ = v34
                        }
                        v35 := GenS4{F0: 9223372036854775807, F1: v6}
                        _ = v35
                    }
                    fmt.Println(v12, v5)
                    if true {
                        var v36 uint8 = genFn0(genFn0(0))
                        v37 := GenS1{F0: (v12 - (v12 % ((231) | 1))), F1: v3, F2: -390}
                        _ = v36
                        _ = v37
                    }
                }
            default:
                for fi38 := 0; fi38 < 5; fi38++ {
                    var v39 int16 = (v6 / ((v6) | 1))
                    if true {
                        v40 := GenS1{F0: fi38, F1: v3, F2: (v12 | (v12 << (((v3 * v3)) & 63)))}
                        var v41 uint64 = 18446744073709551615
                        _ = v40
                        _ = v41
                    }
                    if true {
                        _ = len(v2)
                        var v42 float32 = float32(v21)
                        _ = v42
                    }
                    _ = fi38
                    _ = v39
                }
            }
            v21 -= (((v21 / 1e10) - v21) + 0.1)
        } else {
            v43 := []int{v12}
            v44 := GenS2{F0: 485, F1: 42, F2: (v6 | 32767)}
            _ = v43
            _ = v44
        }
    }
    delete(v2, ((((v15 + v15) + "") + "") + v15))
    if ((1.5 - (3.14159 - float32(float64((float32(v21) + 0.0))))) != 0.0) {
        v45 := GenS5{F0: 42}
        switch 1 {
        case 32768:
            v2[v15] = ((v15 + "foo bar") + v15)
        case -9223372036854775808:
            switch v15 {
            case "foo bar":
                var v46 uint8 = 128
                _ = v46
            }
        case 32767:
            var v47 int8 = v19
            _ = v47
        }
        _ = v45
    } else {
        var v48 float64 = 3.14159
        _ = len(v2)
        _ = v48
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
