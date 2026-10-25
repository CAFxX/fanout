package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 int8; F1 int32; F2 uint16 }
type GenS3 struct { F0 uint32; F1 bool; F2 bool }

func main() {
    var v1 string = "hello"
    var v2 uint64 = 1
    var v3 int32 = -2147483648
    v3 += v3
    for fi4 := 0; fi4 < 14; fi4++ {
        if false {
            for fi5 := 0; fi5 < 10; fi5++ {
                v2 += (v2 / ((113) | 1))
                v2 += (v2 ^ v2)
                var v6 bool = ((float32(-2.25) * (1e10 / (float32(float64(1.5)) + (float32((((0.1 / 0.1) - (((0.0 / 1.0) * 3.14159) * 3.14159)) * (3.14159 - (1.5 / 1.0)))) - (3.14159 - 1.5))))) > 0.1)
                _ = fi5
                _ = v6
            }
            if false {
                for fi7 := 0; fi7 < 18; fi7++ {
                    var v8 int8 = 116
                    _ = fi7
                    _ = v8
                }
                v9 := []int{-690, (fi4 << ((255) & 63)), fi4}
                _ = v9
            } else {
                fmt.Println(v1)
                v1 = v1
                v10 := GenS1{F0: int16(v3), F1: 0}
                _ = v10
            }
            v3 += (-(v3))
        }
        v11 := GenS2{F0: 66, F1: (^(v3)), F2: 255}
        _ = fi4
        _ = v11
    }
    v12 := GenS3{F0: 113, F1: false, F2: (-2.25 < ((-2.25 * -2.25) / 3.14159))}
    fmt.Println(v1)
    v2 -= (v2 % ((v2) | 1))
    v1 = "a"
    for fi13 := 0; fi13 < 4; fi13++ {
        v14 := []int{(fi13 - (fi13 * (fi13 - fi13)))}
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 8; fi15++ {
        if false {
            fmt.Println(fi15, v1, v3)
            var v16 string = ("a" + v1)
            var v17 int16 = 412
            _ = v16
            _ = v17
        } else {
            v18 := []int{fi15, fi15, fi15}
            var v19 int16 = 32767
            var v20 int64 = -1
            _ = v18
            _ = v19
            _ = v20
        }
        for fi21 := 0; fi21 < 19; fi21++ {
            v22 := []int{(fi21 | 32767), (fi15 - fi15)}
            v23 := GenS1{F0: 1, F1: 146}
            _ = fi21
            _ = v22
            _ = v23
        }
        fmt.Println(v3, fi15, v3)
        _ = fi15
    }
    v2 += 4294967295
    var v24 int8 = -67
    switch -1 {
    case 9223372036854775807:
        fmt.Println(v24)
    case -309:
        var v25 int8 = (v24 - (v24 | (v24 + 127)))
        _ = v25
    case 0:
        fmt.Println(v24)
    }
    if (true && false) {
        var v26 string = "\n\t"
        var v27 uint64 = (v2 * (v2 / ((1) | 1)))
        if ((v27 > 4294967295) && true) {
            v27 = (v27 & 301)
            v3 = v3
        } else {
            v3 -= 471
            if (true || false) {
                if false {
                    switch v26 {
                    case "a":
                        v26 = v26
                    case "\n\t":
                        switch 107 {
                        case 9223372036854775807:
                            v2 = v27
                        }
                    default:
                        v28 := GenS1{F0: 0, F1: 4294967295}
                        _ = v28
                    }
                    if (v3 >= (v3 << ((65535) & 63))) {
                        v29 := []int{}
                        switch "hello" {
                        case "foo bar":
                            for fi30 := 0; fi30 < 14; fi30++ {
                                var v31 uint64 = (^(v27))
                                if len(v29) >= 2 {
                                    v29 = v29[1:len(v29)]
                                }
                                _ = fi30
                                _ = v31
                            }
                        }
                        fmt.Println(v3, v3)
                        _ = v29
                    }
                    v32 := []int{}
                    _ = v32
                } else {
                    var v33 uint = 166
                    _ = v33
                }
            }
        }
        _ = v26
        _ = v27
    } else {
        v3 = 0
    }
    var v34 int32 = (-(v3))
    fmt.Println(v34, v1, v34)
    var v35 uint64 = 836
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v12
    _ = v2
    _ = v24
    _ = v3
    _ = v34
    _ = v35
}
