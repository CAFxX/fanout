package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint; F1 string }
func genFn0(p0 int8, p1 bool, p2 uint32) uint8 {
    p0 = -128
    v5 := []int{}
    var v6 uint32 = (^(p2))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 128
}
func genFn1() uint8 {
    v7 := []int{127, -9223372036854775808, 127}
    var v8 float64 = (-2.25 * (3.14159 + 0.1))
    fmt.Println(v8)
    fmt.Println(v8)
    for fi9 := 0; fi9 < 1; fi9++ {
        switch "\n\t" {
        case "\n\t":
            v8 -= (float64(3.14159) * (-2.25 + float64(1e10)))
        case "foo bar":
            for fi10 := 0; fi10 < 1; fi10++ {
                v8 -= (v8 / (((v8 + v8) / (1e10 / ((v8 / (0.1 / (v8 + 1e10))) * (((v8 - (((v8 * float64(-2.25)) + -2.25) * float64(0.1))) - (v8 * v8)) + (v8 - (v8 / -2.25)))))) - (v8 - (v8 - (1e10 + (0.0 * (((1e10 + 1.5) + v8) * (v8 + 1.5))))))))
                v8 = v8
                v8 -= v8
                v8 += (v8 / float64(3.14159))
                _ = fi10
            }
        }
        _ = fi9
    }
    _ = v7
    _ = v8
    return 255
}

func main() {
    var v1 float32 = 0.0
    var v2 uint = 255
    var v3 int32 = 2147483647
    var v4 int32 = v3
    v11 := GenS1{F0: v2, F1: (v3 % (((v3 - -125)) | 1))}
    fmt.Println(v2)
    v1 += 1e10
    v2 = (^(v2))
    var v12 int8 = 0
    var v13 uint16 = 711
    v14 := []int{-433, 32768}
    v15 := GenS1{F0: 255, F1: v4}
    if (2147483647 <= v4) {
        fmt.Println(v13, v2, v3)
        if true {
            var v16 uint8 = 127
            var v17 float32 = 0.1
            _ = v16
            _ = v17
        } else {
            var v18 float32 = -2.25
            _ = v18
        }
    }
    v19 := GenS1{F0: v2, F1: (v3 >> (((v2 << ((v2) & 63))) & 63))}
    if ((219 == -1) || false) {
        v20 := GenS1{F0: (v2 / ((v2) | 1)), F1: v3}
        _ = v20
    } else {
        var v21 uint64 = 42
        _ = v21
    }
    v14 = append(v14, 128)
    v22 := GenS2{F0: 904}
    switch -2147483648 {
    case 32767:
        var v23 uint16 = v13
        _ = v23
    default:
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
    }
    if true {
        v3 -= v3
        if (v13 == (v13 | (v13 + (^(v13))))) {
            if false {
                for fi24 := 0; fi24 < 19; fi24++ {
                    var v25 int8 = v12
                    switch "hello" {
                    case "":
                        if len(v14) >= 2 {
                            v14 = v14[1:len(v14)]
                        }
                    case "\n\t":
                        v4 -= v4
                    default:
                        fmt.Println(v4, v12)
                    }
                    _ = fi24
                    _ = v25
                }
            } else {
                var v26 int16 = -223
                var v27 int = 1
                for fi28 := 0; fi28 < 17; fi28++ {
                    var v29 int = v27
                    switch v27 {
                    case 128:
                        v30 := GenS2{F0: 42}
                        _ = v30
                    case 4294967295:
                        v14 = append(v14, fi28)
                    default:
                        if (((("" + "\n\t") + "\n\t") + "\n\t") > "\n\t") {
                            if false {
                                fmt.Println(v27, v27, v13)
                                v29 = int(v4)
                                var v31 int32 = v4
                                _ = v31
                            } else {
                                var v32 int32 = (v4 >> (((v2 + (v2 << (((v2 | v2)) & 63)))) & 63))
                                _ = v32
                            }
                        } else {
                            var v33 uint16 = (v13 / (((v13 / ((v13) | 1))) | 1))
                            v26 = v26
                            v2 = (v2 ^ uint(v33))
                            _ = v33
                        }
                    }
                    _ = fi28
                    _ = v29
                }
                _ = v26
                _ = v27
            }
        } else {
            v14 = append(v14, 2147483647)
            v13 -= 2
            var v34 float32 = 3.14159
            _ = v34
        }
        var v35 int8 = (^(v12))
        _ = v35
    }
    v36 := GenS3{F0: 18446744073709551615, F1: "hello"}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v36
    _ = v4
}
