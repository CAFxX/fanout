package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 uint8; F1 int8 }
type GenS4 struct { F0 uint32; F1 int8; F2 string }

func main() {
    var v1 uint64 = 2
    v2 := make(map[string]int)
    var v3 float32 = 0.1
    var v4 int8 = 100
    v5 := []uint32{}
    var v6 string = "\n\t"
    var v7 uint = 66
    _ = len(v5)
    var v8 int8 = (v4 * v4)
    if true {
        var v9 bool = (int64(v1) >= 621)
        v8 -= (v8 & int8(v7))
        v5 = append(v5, 147)
        _ = v9
    } else {
        v10 := GenS1{F0: 2147483647, F1: true}
        var v11 string = (v6 + "hello")
        _ = v10
        _ = v11
    }
    switch (v6 + "a") {
    case "\n\t":
        v5 = append(v5, 2147483648)
    case "":
        v4 += v8
    default:
        if true {
            v4 = v8
            for fi12 := 0; fi12 < 10; fi12++ {
                v4 = (-(v8))
                if (false && false) {
                    for fi13 := 0; fi13 < 11; fi13++ {
                        var v14 float32 = (((float32((((((1e10 + (((1.5 + (((float64(v3) - 0.0) - (1e10 - 1.5)) * ((-2.25 / (float64(1.5) * (float64(0.1) + 1.5))) - 0.0))) / -2.25) - 0.1)) - (1e10 - 1e10)) * (0.1 - (((0.1 + -2.25) * 1e10) + 0.0))) - 1e10) + ((1.5 - ((-2.25 * 1e10) + 1.5)) * 3.14159))) + 1e10) - 0.1) / v3)
                        var v15 string = ""
                        v16 := GenS1{F0: 42, F1: false}
                        var v17 uint32 = 42
                        _ = fi13
                        _ = v14
                        _ = v15
                        _ = v16
                        _ = v17
                    }
                    var v18 uint32 = 65535
                    _ = v18
                } else {
                    var v19 int16 = 32767
                    v20 := v2[""]
                    switch (v20 - v20) {
                    case 0:
                        var v21 int16 = v19
                        _ = v21
                    case -1:
                        for fi22 := 0; fi22 < 7; fi22++ {
                            if ((1 == -1) && ((!(!true)) && true)) {
                                v23 := GenS1{F0: 4294967295, F1: (42 >= 9223372036854775807)}
                                if false {
                                    delete(v2, "\n\t")
                                    fmt.Println(fi12, v3)
                                    var v24 float64 = 1.5
                                    _ = v24
                                }
                                _ = v23
                            }
                            switch ("hello" + v6) {
                            case "":
                                var v25 int8 = -128
                                _ = v25
                            case "hello":
                                v26 := GenS2{F0: -2147483648, F1: 2147483647}
                                _ = v26
                            default:
                                if ((!((v20 >> (((v7 ^ v7)) & 63)) != fi22)) || ((fi22 - fi12) >= (fi12 * (-(fi22))))) {
                                    switch (-(fi22)) {
                                    case 4294967295:
                                        if ((((float32(((0.0 + 3.14159) + (1.5 / 3.14159))) - float32(0.0)) / 3.14159) / 1.5) != v3) {
                                            var v27 int8 = -128
                                            _ = v27
                                        } else {
                                            v28 := GenS3{F0: 199, F1: -1}
                                            fmt.Println(v3, v20)
                                            _ = v28
                                        }
                                    case 2147483647:
                                        v5 = append(v5, 65535)
                                    default:
                                        var v29 float32 = v3
                                        _ = v29
                                    }
                                    if true {
                                        var v30 int16 = v19
                                        var v31 int16 = v19
                                        _ = v30
                                        _ = v31
                                    }
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                }
                            }
                            v32 := GenS1{F0: 65535, F1: true}
                            _ = fi22
                            _ = v32
                        }
                    case -505:
                        v33 := GenS4{F0: 600, F1: (v4 + v4), F2: "hello"}
                        _ = v33
                    }
                    _ = v19
                    _ = v20
                }
                for fi34 := 0; fi34 < 7; fi34++ {
                    switch v6 {
                    case "":
                        var v35 int32 = -286
                        _ = v35
                    case "a":
                        var v36 float64 = 0.0
                        _ = v36
                    }
                    var v37 int32 = 2147483647
                    v4 = (v4 & (v4 >> ((v7) & 63)))
                    var v38 int32 = 42
                    fmt.Println(v37, v4, v8)
                    _ = fi34
                    _ = v37
                    _ = v38
                }
                _ = fi12
            }
        }
    }
    var v39 bool = (!false)
    for fi40 := 0; fi40 < 17; fi40++ {
        v7 = v7
        _ = fi40
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
