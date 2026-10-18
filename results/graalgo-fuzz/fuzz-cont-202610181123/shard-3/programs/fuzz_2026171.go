package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 bool }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 uint16; F1 int }
func genFn0() string {
    v6 := GenS1{F0: 2, F1: -718, F2: (!true)}
    v7 := []int{83, -270, -9223372036854775808}
    var v8 float64 = 1e10
    if (!false) {
        if (4294967295 > 2) {
            fmt.Println(v8)
            for fi9 := 0; fi9 < 12; fi9++ {
                if true {
                    v10 := GenS2{F0: 65535, F1: 2}
                    if true {
                        v8 = v8
                        if false {
                            _ = len(v7)
                            fmt.Println(v8, v8)
                        } else {
                            for fi11 := 0; fi11 < 8; fi11++ {
                                v12 := GenS2{F0: 12, F1: 42}
                                var v13 float32 = float32(1e10)
                                _ = fi11
                                _ = v12
                                _ = v13
                            }
                            var v14 int8 = 5
                            v15 := GenS1{F0: 128, F1: -612, F2: true}
                            _ = v14
                            _ = v15
                        }
                    }
                    _ = v10
                } else {
                    if len(v7) > 0 {
                        v7[0] = 128
                    }
                    fmt.Println(v8, v8)
                    v8 += v8
                }
                _ = fi9
            }
            fmt.Println(v8)
        } else {
            if (0 < 2147483648) {
                v8 -= v8
                if len(v7) > 0 {
                    v7[0] = 32768
                }
                v16 := GenS3{F0: 382, F1: 0}
                _ = v16
            } else {
                var v17 uint = 193
                switch "" {
                case "":
                    fmt.Println(v17, v17)
                case "a":
                    v8 = (-2.25 - 0.0)
                default:
                    v17 -= 803
                }
                if true {
                    v8 -= (-2.25 + v8)
                    v18 := GenS2{F0: 4294967295, F1: 296}
                    switch "" {
                    case "\n\t":
                        v7 = append(v7, 1)
                    case "hello":
                        v19 := GenS2{F0: (v17 + 18446744073709551615), F1: 2}
                        _ = v19
                    }
                    _ = v18
                }
                _ = v17
            }
            var v20 float32 = (float32(float64(float32(((v8 * v8) - (v8 / (v8 - 1.5)))))) * 3.14159)
            fmt.Println(v20)
            _ = v20
        }
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
    }
    var v21 uint32 = 2147483647
    _ = v21
    _ = v6
    _ = v7
    _ = v8
    return "a"
}
func genFn1(p0 int32) uint32 {
    switch "" {
    case "\n\t":
        v22 := GenS1{F0: 255, F1: p0, F2: (!(!false))}
        _ = v22
    case "hello":
        var v23 uint16 = 65535
        _ = v23
    default:
        var v24 float64 = 0.0
        _ = v24
    }
    fmt.Println(p0)
    switch -2147483648 {
    case 1:
        switch "" {
        case "a":
            var v25 uint = 65535
            _ = v25
        default:
            p0 = p0
        }
    default:
        p0 = (^(p0))
    }
    fmt.Println(p0)
    v26 := []int{2147483647, 32767}
    _ = p0
    _ = v26
    return 355
}

func main() {
    var v1 int32 = 0
    v2 := make(map[string]int)
    v3 := []int{712}
    var v4 bool = (100 < 0)
    var v5 string = ("foo bar" + "foo bar")
    v5 = v5
    var v27 string = "hello"
    fmt.Println(v5, v27, v4)
    v3 = append(v3, 4294967295)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v4 = (192 <= 32768)
    switch 42 {
    case 0:
        var v28 int16 = 42
        _ = v28
    case 4294967295:
        fmt.Println(v27)
    }
    var v29 int = 942
    fmt.Println(v1)
    v30 := GenS1{F0: 246, F1: (v1 | (v1 ^ v1)), F2: (31 <= -1)}
    for fi31 := 0; fi31 < 2; fi31++ {
        var v32 float32 = (0.0 / 1.0)
        _ = fi31
        _ = v32
    }
    var v33 uint = 2
    switch -617 {
    case -1:
        delete(v2, ((("hello" + ("hello" + (v5 + v5))) + v5) + v27))
    case 128:
        fmt.Println(v5, v27, v5)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v33
    _ = v4
    _ = v5
}
