package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint }

func main() {
    var v1 int32 = 2147483647
    var v2 int16 = -338
    var v3 float32 = 3.14159
    var v4 float32 = v3
    v5 := GenS1{F0: 0}
    v6 := []int{}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    switch "\n\t" {
    case "":
        v7 := GenS1{F0: 2}
        _ = v7
    case "hello":
        for fi8 := 0; fi8 < 14; fi8++ {
            _ = len(v6)
            for fi9 := 0; fi9 < 1; fi9++ {
                var v10 int64 = 0
                _ = fi9
                _ = v10
            }
            for fi11 := 0; fi11 < 14; fi11++ {
                if false {
                    var v12 float32 = v3
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    _ = v12
                } else {
                    if (1 > 2) {
                        var v13 uint32 = 1
                        _ = v13
                    } else {
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    }
                    var v14 uint32 = 0
                    v3 -= (v4 + ((0.1 / (v4 / (0.0 / ((3.14159 - (v3 - 0.1)) / (float32(float64(v3)) / 1e10))))) * (v3 / -2.25)))
                    _ = v14
                }
                _ = fi11
            }
            _ = fi8
        }
    case "\n\t":
        for fi15 := 0; fi15 < 3; fi15++ {
            if false {
                for fi16 := 0; fi16 < 1; fi16++ {
                    if (v3 < ((v4 + ((float32(1e10) - v4) - ((v4 * float32(1e10)) + ((1e10 - (v3 / (v3 / 3.14159))) * float32(1e10))))) * v3)) {
                        v17 := GenS1{F0: 196}
                        _ = v17
                    } else {
                        fmt.Println(v3)
                    }
                    _ = fi16
                }
                v3 += float32(-2.25)
                if len(v6) > 0 {
                    v6[0] = 727
                }
            }
            _ = fi15
        }
    }
    fmt.Println(v4, v4, v2)
    for fi18 := 0; fi18 < 14; fi18++ {
        v19 := GenS1{F0: 255}
        if (v2 > v2) {
            if (690 < 2147483648) {
                var v20 float32 = 1e10
                v20 = 1e10
                v3 = 1.5
                _ = v20
            }
        }
        if true {
            v21 := GenS2{F0: 2147483648}
            var v22 int64 = 1
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = v21
            _ = v22
        }
        _ = fi18
        _ = v19
    }
    v1 = v1
    if (32767 != 65535) {
        for fi23 := 0; fi23 < 8; fi23++ {
            var v24 int = (fi23 << ((2) & 63))
            _ = fi23
            _ = v24
        }
        switch "" {
        case "hello":
            if false {
                v4 = v3
                var v25 uint = 65535
                _ = v25
            } else {
                v26 := GenS2{F0: 379}
                for fi27 := 0; fi27 < 5; fi27++ {
                    _ = len(v6)
                    switch (fi27 >> ((18446744073709551615) & 63)) {
                    case 0:
                        v4 += ((v4 / v4) + (float32(1.5) / v4))
                    case 32767:
                        v28 := GenS1{F0: 18446744073709551615}
                        _ = v28
                    case -860:
                        v1 -= (v1 % ((int32(fi27)) | 1))
                    }
                    _ = fi27
                }
                _ = v26
            }
        case "a":
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        case "foo bar":
            v3 -= (-2.25 + 0.1)
        default:
            v1 = (v1 + (v1 - v1))
        }
        if false {
            if true {
                v29 := GenS2{F0: 2147483648}
                var v30 int8 = 127
                _ = v29
                _ = v30
            } else {
                var v31 int64 = 0
                fmt.Println(v2, v2, v2)
                _ = v31
            }
        } else {
            var v32 float32 = 1.5
            var v33 uint32 = 247
            _ = v32
            _ = v33
        }
    }
    switch 127 {
    case 0:
        v34 := GenS2{F0: 798}
        _ = v34
    case 53:
        v3 = (0.0 - -2.25)
    default:
        var v35 float32 = (v4 / ((v3 + 0.1) - ((((0.1 + (float32((0.1 - 0.0)) - float32((1.5 + (((1e10 * float64(0.1)) * (((-2.25 / float64(v3)) * 0.0) / -2.25)) * 3.14159))))) * ((((v4 - (1.5 * v3)) - 1.5) + ((v3 - 3.14159) * (v3 + (-2.25 * 1e10)))) + ((1.5 - v3) * 0.1))) / -2.25) - v3)))
        _ = v35
    }
    fmt.Println(v1)
    v36 := GenS3{F0: 396}
    var v37 int = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
