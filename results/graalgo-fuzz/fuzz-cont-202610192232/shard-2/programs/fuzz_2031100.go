package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 uint16; F2 bool }
type GenS3 struct { F0 int64; F1 int16 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int64; F1 uint16; F2 string }
type GenS7 struct { F0 int64; F1 uint32; F2 int64 }

func main() {
    var v1 uint16 = 2
    var v2 float32 = 0.1
    var v3 int64 = 1
    var v4 uint32 = 42
    v5 := GenS1{F0: 1}
    for fi6 := 0; fi6 < 15; fi6++ {
        v7 := GenS2{F0: 1000, F1: v1, F2: true}
        if ("a" > ("\n\t" + ("a" + "\n\t"))) {
            var v8 bool = (true && ((0 != 499) || (0 >= -1)))
            if v8 {
                v8 = false
                v2 = ((((v2 + float32(float64((0.1 + (float32(float64(v2)) / v2))))) / 1.0) * (v2 + 0.1)) + 1e10)
                var v9 int16 = 42
                _ = v9
            } else {
                for fi10 := 0; fi10 < 2; fi10++ {
                    v4 = (v4 % (((v4 << ((2) & 63))) | 1))
                    _ = fi10
                }
                v11 := GenS3{F0: (^(v3)), F1: 1000}
                _ = v11
            }
            _ = v8
        } else {
            v12 := GenS4{F0: v4}
            _ = v12
        }
        _ = fi6
        _ = v7
    }
    var v13 uint8 = 2
    var v14 uint = 18446744073709551615
    v15 := []int{941}
    var v16 int64 = v3
    for fi17 := 0; fi17 < 4; fi17++ {
        v18 := GenS3{F0: v16, F1: 205}
        for fi19 := 0; fi19 < 3; fi19++ {
            v20 := GenS1{F0: (v14 / ((v14) | 1))}
            if (-1 == 291) {
                var v21 bool = false
                if v21 {
                    if (1000 == -758) {
                        fmt.Println(v13, v4)
                        if (815 <= v1) {
                            var v22 int16 = 0
                            v22 = -32768
                            var v23 uint32 = (v4 - 42)
                            _ = v22
                            _ = v23
                        } else {
                            v24 := GenS5{F0: (v4 / ((v4) | 1))}
                            v4 = 2147483648
                            _ = v24
                        }
                    } else {
                        if len(v15) > 0 {
                            v15[0] = (fi17 + fi19)
                        }
                    }
                    if (42 <= 0) {
                        var v25 uint64 = 73
                        v1 = (v1 / ((v1) | 1))
                        _ = v25
                    }
                    switch ("\n\t" + "hello") {
                    case "a":
                        v16 = (v16 - (v16 | 0))
                    case "foo bar":
                        v26 := GenS3{F0: (-(v16)), F1: 1000}
                        _ = v26
                    case "":
                        v27 := GenS6{F0: v3, F1: 37, F2: "\n\t"}
                        _ = v27
                    }
                } else {
                    v3 = 9223372036854775807
                    if len(v15) >= 2 {
                        v15 = v15[1:len(v15)]
                    }
                }
                _ = v21
            } else {
                fmt.Println(fi17)
                v4 = v4
            }
            fmt.Println(v2)
            _ = fi19
            _ = v20
        }
        v28 := GenS5{F0: v4}
        _ = fi17
        _ = v18
        _ = v28
    }
    if len(v15) > 0 {
        v15[0] = -1
    }
    var v29 int = 179
    for fi30 := 0; fi30 < 10; fi30++ {
        v16 += int64(v14)
        v31 := GenS7{F0: 0, F1: 2147483648, F2: 185}
        v14 = 4294967295
        _ = fi30
        _ = v31
    }
    fmt.Println(v29, v16, v13)
    var v32 int16 = 32767
    v33 := GenS2{F0: v32, F1: (v1 / (((^(v1))) | 1)), F2: (!true)}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
