package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 int }
type GenS2 struct { F0 uint32; F1 uint16; F2 int64 }
type GenS3 struct { F0 int8; F1 int16; F2 bool }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int; F1 int }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 string; F1 int16; F2 uint16 }

func main() {
    v1 := []uint32{2147483647, 42, 931, 42}
    var v2 int8 = 127
    var v3 float32 = (0.0 - (3.14159 * 1e10))
    var v4 uint8 = 128
    fmt.Println(v2, v2, v4)
    v5 := GenS1{F0: "\n\t", F1: 32768, F2: 32768}
    v6 := GenS2{F0: 777, F1: 1, F2: -9223372036854775808}
    switch 844 {
    case 32767:
        if true {
            if len(v1) > 0 {
                v1[0] = 4294967295
            }
        }
    case -246:
        v7 := GenS1{F0: "a", F1: 127, F2: 4294967295}
        _ = v7
    case 290:
        for fi8 := 0; fi8 < 2; fi8++ {
            var v9 uint = 2
            for fi10 := 0; fi10 < 10; fi10++ {
                if (float64(((v3 + ((0.0 + v3) / float32(((3.14159 * 3.14159) + (0.1 - 0.1))))) / v3)) <= 0.1) {
                    _ = len(v1)
                }
                _ = fi10
            }
            _ = fi8
            _ = v9
        }
    }
    var v11 string = "a"
    for fi12 := 0; fi12 < 15; fi12++ {
        v13 := GenS3{F0: 127, F1: -1, F2: (2 <= 18446744073709551615)}
        if true {
            switch fi12 {
            case 42:
                _ = len(v1)
            case 32767:
                v14 := GenS2{F0: 713, F1: 65535, F2: -1}
                _ = v14
            case 128:
                v11 = v11
            default:
                for fi15 := 0; fi15 < 14; fi15++ {
                    if len(v1) > 0 {
                        v1[0] = 2
                    }
                    _ = fi15
                }
            }
            switch ("foo bar" + (v11 + v11)) {
            case "\n\t":
                var v16 uint64 = 534
                _ = v16
            }
        } else {
            if ((42 != 1) && (-128 < 127)) {
                fmt.Println(v4, v2, v11)
                if (-833 == 42) {
                    var v17 uint16 = 255
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v18 := GenS3{F0: 100, F1: 32767, F2: true}
                    _ = v17
                    _ = v18
                } else {
                    for fi19 := 0; fi19 < 2; fi19++ {
                        if (v4 >= 255) {
                            fmt.Println(v2, v2, v2)
                            if (!(!(v4 == v4))) {
                                var v20 int64 = 9223372036854775807
                                _ = v20
                            } else {
                                v3 += (v3 * 3.14159)
                            }
                        }
                        if false {
                            v2 = (v2 - v2)
                        }
                        var v21 float32 = v3
                        _ = fi19
                        _ = v21
                    }
                }
                v22 := GenS4{F0: 32767}
                _ = v22
            }
        }
        v11 = (v11 + "foo bar")
        _ = fi12
        _ = v13
    }
    for fi23 := 0; fi23 < 15; fi23++ {
        for fi24 := 0; fi24 < 12; fi24++ {
            switch (-(fi23)) {
            case 128:
                if ((float32(((1e10 + ((float64(v3) / 1.5) + -2.25)) * -2.25)) * ((float32(float64(v3)) - (v3 + float32((1e10 - ((((0.1 * 0.1) + (((1e10 + ((((1.5 + 3.14159) * (1e10 / (1.5 - ((1.5 * float64(1.5)) / 1.0)))) + ((0.1 + 3.14159) + float64(-2.25))) - float64(-2.25))) / -2.25) / -2.25)) * 0.1) + (((((0.0 / (3.14159 + float64(v3))) / (0.1 / ((0.1 - 1.5) / 1.0))) / 1.0) * 1.5) + ((float64(float32(0.0)) / ((1.5 / 0.1) - (float64(0.1) - -2.25))) + 1.5))))))) + v3)) < v3) {
                    if (2147483647 > -2147483648) {
                        fmt.Println(v11, v3)
                        v25 := GenS2{F0: 535, F1: 255, F2: 1}
                        var v26 float64 = 3.14159
                        _ = v25
                        _ = v26
                    } else {
                        v27 := GenS5{F0: (^(fi23)), F1: fi23}
                        v4 = 128
                        _ = v27
                    }
                    var v28 uint8 = 80
                    _ = v28
                } else {
                    var v29 uint64 = 889
                    _ = v29
                }
                v30 := GenS6{F0: 397}
                if (false && false) {
                    var v31 int16 = 1
                    v11 = ("foo bar" + v11)
                    _ = v31
                }
                _ = v30
            }
            var v32 uint16 = 1
            _ = fi24
            _ = v32
        }
        _ = fi23
    }
    v33 := GenS7{F0: "foo bar", F1: 134, F2: 0}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
