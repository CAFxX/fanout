package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 uint32; F2 uint }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 uint16 }

func main() {
    var v1 float32 = (3.14159 * (float32(float64(3.14159)) - 0.0))
    var v2 bool = true
    v3 := []int{9223372036854775807, 0, 4294967295, 4294967295}
    v4 := []int{32768}
    var v5 int = 1
    v2 = false
    switch (v5 | (v5 / ((713) | 1))) {
    case -2147483648:
        v5 -= 32767
    }
    switch (v5 << ((42) & 63)) {
    case 80:
        v3 = append(v3, 225)
    default:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    if (!((((((1e10 / (float64(v1) + -2.25)) / 1.0) + 1e10) / 1.0) / -2.25) >= (1.5 + (3.14159 - 1e10)))) {
        var v6 int64 = 1
        _ = v6
    } else {
        v5 = v5
        if (65535 == 2147483648) {
            if (v2 && v2) {
                v7 := GenS1{F0: 4294967295}
                v1 += v1
                for fi8 := 0; fi8 < 18; fi8++ {
                    for fi9 := 0; fi9 < 8; fi9++ {
                        v5 -= fi8
                        _ = fi9
                    }
                    v10 := GenS2{F0: 18446744073709551615, F1: 956, F2: 255}
                    var v11 int = 939
                    _ = fi8
                    _ = v10
                    _ = v11
                }
                _ = v7
            } else {
                for fi12 := 0; fi12 < 18; fi12++ {
                    var v13 uint64 = 990
                    if ((!v2) && true) {
                        fmt.Println(fi12, v2, fi12)
                        v14 := GenS3{F0: -1, F1: 65535}
                        _ = v14
                    }
                    v2 = ((-274 <= 42) || (true && (-629 == 904)))
                    _ = fi12
                    _ = v13
                }
                if v2 {
                    switch ("hello" + "hello") {
                    case "a":
                        for fi15 := 0; fi15 < 5; fi15++ {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = fi15
                        }
                    case "foo bar":
                        var v16 int = 42
                        _ = v16
                    }
                } else {
                    if (102 < 2) {
                        var v17 int16 = -32768
                        _ = v17
                    } else {
                        v1 = (v1 * 0.0)
                    }
                    v1 += (v1 - ((((v1 - (3.14159 / (v1 + v1))) * (v1 / ((-2.25 * v1) * (((3.14159 + v1) * 0.1) + v1)))) + float32((1.5 / 0.1))) - (1.5 + v1)))
                }
                if (!(((42 < 42) && ((-1 > 34) || (v2 && (-9223372036854775808 >= 455)))) && (v2 || (-528 != 9223372036854775807)))) {
                    for fi18 := 0; fi18 < 15; fi18++ {
                        var v19 string = ""
                        fmt.Println(v1, fi18, v2)
                        v1 += 0.0
                        _ = fi18
                        _ = v19
                    }
                    for fi20 := 0; fi20 < 9; fi20++ {
                        if v2 {
                            v3 = append(v3, 32767)
                            v5 += 127
                        } else {
                            v21 := GenS2{F0: 2, F1: 0, F2: 65535}
                            v22 := GenS2{F0: 255, F1: 0, F2: 2}
                            _ = v21
                            _ = v22
                        }
                        v1 += 1.5
                        _ = fi20
                    }
                }
            }
            v23 := GenS2{F0: 18446744073709551615, F1: 4294967295, F2: 65535}
            _ = v23
        } else {
            v24 := GenS4{F0: 2}
            if v2 {
                if ((-79 <= 1) && v2) {
                    if (76 != 1) {
                        for fi25 := 0; fi25 < 14; fi25++ {
                            v26 := GenS3{F0: 100, F1: 0}
                            fmt.Println(v1, v2, v1)
                            v27 := GenS4{F0: 2}
                            _ = fi25
                            _ = v26
                            _ = v27
                        }
                        if v2 {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            var v28 uint32 = 2147483648
                            _ = v28
                        }
                        var v29 float64 = 1.5
                        _ = v29
                    }
                    v5 -= v5
                    var v30 uint = 180
                    _ = v30
                }
                v1 = 0.0
                v1 -= (float32(-2.25) * 1.5)
            }
            _ = v24
        }
    }
    var v31 uint = 18446744073709551615
    var v32 uint8 = 1
    _ = len(v3)
    if len(v3) > 0 {
        v3[0] = v5
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
