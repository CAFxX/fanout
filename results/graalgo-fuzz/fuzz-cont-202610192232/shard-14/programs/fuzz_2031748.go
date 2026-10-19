package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 int }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 string; F1 int64 }
func genFn0(p0 string, p1 int64, p2 uint8) uint {
    v5 := GenS1{F0: "\n\t", F1: 1, F2: -9223372036854775808}
    for fi6 := 0; fi6 < 11; fi6++ {
        v7 := GenS1{F0: p0, F1: 887, F2: -406}
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: (((p0 + p0) + ("hello" + ("\n\t" + "\n\t"))) + p0), F1: 42, F2: -2147483648}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v8
    return 1
}

func main() {
    var v1 uint64 = 4294967295
    var v2 bool = true
    var v3 float32 = 1.5
    var v4 string = "\n\t"
    for fi9 := 0; fi9 < 7; fi9++ {
        v10 := GenS1{F0: (("" + "hello") + ("hello" + v4)), F1: 0, F2: fi9}
        v11 := []int{}
        fmt.Println(fi9, v2)
        _ = fi9
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 12; fi12++ {
        v3 += 0.1
        v13 := GenS1{F0: "hello", F1: 82, F2: 32768}
        _ = fi12
        _ = v13
    }
    var v14 int = 127
    switch ("hello" + v4) {
    case "foo bar":
        v15 := []int{2147483647, (v14 & v14), (^(v14))}
        _ = v15
    case "":
        switch -899 {
        case 128:
            v14 = 32767
        default:
            v16 := []int{v14}
            _ = v16
        }
    case "hello":
        fmt.Println(v3, v1, v3)
    }
    v1 += v1
    if v2 {
        switch v14 {
        case -771:
            v17 := GenS2{F0: 255, F1: (uint8(v1) == 1)}
            _ = v17
        }
        if (v14 < (-(v14))) {
            var v18 uint64 = (v1 - (v1 | 168))
            fmt.Println(v2, v3)
            _ = v18
        } else {
            var v19 uint8 = 74
            if true {
                fmt.Println(v2)
                switch v14 {
                case 42:
                    var v20 uint16 = 353
                    _ = v20
                case 4294967295:
                    if (4294967295 != 20) {
                        v3 -= v3
                        v14 = (v14 + v14)
                        fmt.Println(v3)
                    } else {
                        v21 := []int{-2147483648, 128, (v14 ^ (v14 ^ (v14 ^ 4294967295))), v14}
                        _ = v21
                    }
                }
                v22 := []int{(v14 | 1), (v14 / ((851) | 1)), int(v19)}
                _ = v22
            }
            if true {
                switch v4 {
                case "":
                    if (-574 > 0) {
                        fmt.Println(v19, v19)
                        v23 := GenS1{F0: ((("" + v4) + v4) + v4), F1: 1, F2: (v14 | (v14 & (v14 ^ v14)))}
                        _ = v23
                    }
                case "\n\t":
                    var v24 int8 = -35
                    _ = v24
                default:
                    v3 += v3
                }
                v25 := []int{v14, (v14 % (((v14 ^ -1)) | 1))}
                for fi26 := 0; fi26 < 6; fi26++ {
                    for fi27 := 0; fi27 < 1; fi27++ {
                        var v28 int16 = -826
                        var v29 int32 = 1
                        v14 = 32767
                        v28 += v28
                        v2 = (v2 && false)
                        _ = fi27
                        _ = v28
                        _ = v29
                    }
                    for fi30 := 0; fi30 < 2; fi30++ {
                        if v2 {
                            v31 := GenS2{F0: 42, F1: ((-1 >= -32768) && (!v2))}
                            fmt.Println(v1, fi26, v1)
                            var v32 uint64 = (^(v1))
                            _ = v31
                            _ = v32
                        } else {
                            var v33 string = v4
                            v2 = false
                            _ = v33
                        }
                        var v34 int = (fi30 % ((-1) | 1))
                        v34 -= fi30
                        var v35 float32 = (v3 + 1.5)
                        _ = fi30
                        _ = v34
                        _ = v35
                    }
                    _ = fi26
                }
                _ = v25
            } else {
                v36 := GenS2{F0: 1, F1: true}
                _ = v36
            }
            _ = v19
        }
    }
    switch v4 {
    case "hello":
        for fi37 := 0; fi37 < 3; fi37++ {
            v3 = 3.14159
            _ = fi37
        }
    case "\n\t":
        for fi38 := 0; fi38 < 1; fi38++ {
            v39 := []int{611, fi38, (v14 ^ v14), v14}
            _ = fi38
            _ = v39
        }
    default:
        v40 := []int{}
        _ = v40
    }
    v4 = v4
    v41 := []int{v14, (v14 + -1)}
    v42 := GenS3{F0: (v4 + v4), F1: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v41:", len(v41), v41)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
}
