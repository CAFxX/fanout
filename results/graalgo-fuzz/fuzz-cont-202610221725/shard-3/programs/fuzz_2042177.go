package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 int32; F2 uint8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int32; F1 int8; F2 int8 }

func main() {
    var v1 float32 = 3.14159
    var v2 uint8 = 2
    var v3 bool = (4294967295 >= 1)
    var v4 float32 = v1
    v5 := []int{127, 690, -798}
    v6 := make(map[string]string)
    for fi7 := 0; fi7 < 11; fi7++ {
        fmt.Println(v4, v4, fi7)
        var v8 int8 = -100
        if (-32768 > 32767) {
            v9 := GenS1{F0: 693}
            v8 -= (v8 & v8)
            for fi10 := 0; fi10 < 6; fi10++ {
                v4 = 3.14159
                var v11 int16 = -1
                if (v3 || v3) {
                    v5 = append(v5, (fi10 << ((4294967295) & 63)))
                }
                _ = fi10
                _ = v11
            }
            _ = v9
        }
        _ = fi7
        _ = v8
    }
    v12 := GenS1{F0: 0}
    fmt.Println(v1)
    v1 += (v4 * float32((0.1 + 1e10)))
    switch 32768 {
    case 2147483647:
        v13 := GenS1{F0: 2}
        _ = v13
    }
    switch "a" {
    case "a":
        v5 = append(v5, -776)
    case "\n\t":
        for fi14 := 0; fi14 < 15; fi14++ {
            for fi15 := 0; fi15 < 11; fi15++ {
                v4 -= (((v4 * ((v4 / -2.25) + v4)) * v4) - v1)
                if (0 > 356) {
                    v3 = ((^(v2)) < v2)
                    v4 -= 0.0
                    switch fi14 {
                    case -595:
                        v1 = (v4 - (-2.25 - v1))
                    case 593:
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    case 128:
                        v16 := GenS1{F0: 283}
                        _ = v16
                    }
                }
                if (0 == 85) {
                    switch (fi15 & 259) {
                    case 0:
                        var v17 string = "a"
                        _ = v17
                    default:
                        v3 = true
                    }
                } else {
                    v2 -= v2
                    var v18 int8 = -121
                    switch fi14 {
                    case 32768:
                        switch (^(fi15)) {
                        case 0:
                            if true {
                                var v19 uint16 = 1
                                var v20 int8 = -48
                                var v21 uint = 18446744073709551615
                                _ = v19
                                _ = v20
                                _ = v21
                            } else {
                                v22 := GenS1{F0: 1}
                                _ = v22
                            }
                        case 128:
                            v2 -= v2
                        case -9223372036854775808:
                            var v23 uint32 = 821
                            _ = v23
                        }
                    case 1:
                        var v24 int8 = 127
                        _ = v24
                    case 42:
                        if v3 {
                            if (false && v3) {
                                var v25 float64 = (((0.0 - (1.5 + 3.14159)) * (1e10 - ((-2.25 * -2.25) - (((3.14159 - (-2.25 - (float64(((v4 - (0.1 * v1)) / v1)) * (float64(v4) * 3.14159)))) / ((1.5 * ((0.0 * (1.5 - 3.14159)) / 1e10)) / ((1e10 / (0.0 + (float64(v4) * (1.5 / 0.1)))) / 0.1))) / 0.1)))) - (1e10 * float64((float32((0.0 - -2.25)) * float32(0.1)))))
                                _ = v25
                            } else {
                                var v26 int = (fi14 ^ fi14)
                                v1 += v4
                                if (((v18 + v18) < v18) || (-447 == -1)) {
                                    fmt.Println(v18, v3)
                                } else {
                                    switch (v26 * -956) {
                                    case -9223372036854775808:
                                        fmt.Println(v2, v1)
                                    case 2147483647:
                                        v27 := GenS2{F0: 18446744073709551615, F1: 42, F2: (^(v2))}
                                        _ = v27
                                    default:
                                        switch "a" {
                                        case "":
                                            var v28 int64 = 1
                                            _ = v28
                                        default:
                                            v2 += 17
                                        }
                                    }
                                }
                                var v29 uint16 = 1
                                _ = v26
                                _ = v29
                            }
                            var v30 int8 = v18
                            v31 := GenS3{F0: -496}
                            _ = v30
                            _ = v31
                        }
                    }
                    _ = v18
                }
                _ = fi15
            }
            delete(v6, "")
            _ = fi14
        }
    }
    v3 = (!true)
    v32 := GenS4{F0: 2147483647, F1: -78, F2: -68}
    var v33 int16 = -32768
    var v34 int16 = v33
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
