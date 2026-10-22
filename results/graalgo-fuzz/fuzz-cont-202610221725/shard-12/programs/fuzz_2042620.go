package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 int16; F1 uint64; F2 bool }
type GenS3 struct { F0 uint64; F1 uint32; F2 int8 }
func genFn0(p0 int16) int8 {
    var v5 uint64 = 4294967295
    var v6 int32 = 1
    _ = p0
    _ = v5
    _ = v6
    return -128
}

func main() {
    var v1 bool = true
    var v2 uint = 153
    v3 := []string{"\n\t", "foo bar"}
    var v4 uint8 = 42
    switch 32767 {
    case 0:
        fmt.Println(v2)
    case 9223372036854775807:
        var v7 bool = (157 != 4294967295)
        _ = v7
    default:
        fmt.Println(v2)
    }
    v8 := GenS1{F0: 2147483648, F1: 1}
    var v9 uint8 = 42
    v10 := GenS1{F0: 2147483648, F1: -1}
    var v11 float32 = -2.25
    v9 = (v4 ^ (v9 + (v4 ^ v9)))
    v12 := GenS2{F0: 42, F1: 4294967295, F2: (2 <= 2)}
    for fi13 := 0; fi13 < 3; fi13++ {
        fmt.Println(fi13, v1, v4)
        var v14 uint16 = 664
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 3; fi15++ {
        v11 = float32(0.0)
        for fi16 := 0; fi16 < 4; fi16++ {
            v3 = append(v3, "")
            if (18446744073709551615 >= v2) {
                v17 := GenS2{F0: -1, F1: 1, F2: true}
                _ = v17
            } else {
                if v1 {
                    switch 2147483647 {
                    case 1:
                        _ = len(v3)
                    case -23:
                        if (0.1 <= (3.14159 / (((((3.14159 + 3.14159) - -2.25) + 1.5) * float64(v11)) + 1.5))) {
                            v1 = v1
                            if (float64((((((v11 + 1.5) - v11) / (v11 + v11)) - 1.5) / v11)) > ((1e10 * (float64(((v11 / (v11 - v11)) * -2.25)) - (0.1 / 1.0))) - -2.25)) {
                                var v18 int64 = -521
                                fmt.Println(v1, v1)
                                _ = v18
                            } else {
                                v4 += (v4 + v9)
                                fmt.Println(fi15)
                            }
                        }
                    case -1:
                        switch fi16 {
                        case 424:
                            if len(v3) > 0 {
                                v3[0] = "a"
                            }
                        case 403:
                            var v19 float32 = v11
                            _ = v19
                        case 350:
                            v20 := GenS2{F0: 596, F1: 18446744073709551615, F2: false}
                            _ = v20
                        }
                    }
                    var v21 int64 = -9223372036854775808
                    v22 := GenS1{F0: 42, F1: v21}
                    _ = v21
                    _ = v22
                } else {
                    v23 := GenS3{F0: 471, F1: 895, F2: 127}
                    fmt.Println(v4)
                    _ = v23
                }
            }
            _ = fi16
        }
        switch (fi15 & (-(fi15))) {
        case -1:
            if (("" >= (("hello" + "\n\t") + "a")) || (!(v1 && (!false)))) {
                v24 := GenS3{F0: 0, F1: 0, F2: 1}
                var v25 string = ("foo bar" + (("" + "a") + ("hello" + ("hello" + ("hello" + "hello")))))
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = v24
                _ = v25
            } else {
                if (0.0 <= (((1.5 + float64(float32(0.1))) - (((float64(v11) + 3.14159) - (float64(v11) / (float64(-2.25) * float64(3.14159)))) / 1e10)) * -2.25)) {
                    _ = len(v3)
                    v9 = v9
                } else {
                    fmt.Println(v2)
                    var v26 float64 = float64(v11)
                    v9 -= v4
                    _ = v26
                }
            }
        default:
            v2 = v2
        }
        _ = fi15
    }
    for fi27 := 0; fi27 < 15; fi27++ {
        fmt.Println(v1, v9)
        v28 := GenS1{F0: 4294967295, F1: 42}
        _ = fi27
        _ = v28
    }
    var v29 int = -16
    v11 = (v11 + (((-2.25 * -2.25) / (v11 * v11)) + ((v11 * 0.0) / v11)))
    for fi30 := 0; fi30 < 15; fi30++ {
        v1 = ((-219 >= -32768) && v1)
        v29 = v29
        _ = fi30
    }
    v1 = v1
    v9 = (v9 & 53)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
