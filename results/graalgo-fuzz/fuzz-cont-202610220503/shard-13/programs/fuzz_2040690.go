package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 uint8 }
type GenS2 struct { F0 int8; F1 int16; F2 int }
type GenS3 struct { F0 int64; F1 int64 }
type GenS4 struct { F0 int; F1 int32; F2 int32 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint16; F1 uint }
type GenS7 struct { F0 int64; F1 uint16 }
type GenS8 struct { F0 uint8; F1 int32 }
func genFn0() bool {
    var v7 float32 = -2.25
    v8 := GenS1{F0: 179, F1: 95, F2: 2}
    var v9 bool = false
    _ = v7
    _ = v8
    _ = v9
    return (2147483647 <= 0)
}
func genFn1(p0 uint32) uint32 {
    if (-485 >= 590) {
        v10 := []int{}
        switch "foo bar" {
        case "hello":
            var v11 int8 = 105
            _ = v11
        case "\n\t":
            p0 -= 670
        case "foo bar":
            var v12 float64 = (0.1 + 3.14159)
            _ = v12
        }
        v13 := GenS2{F0: 1, F1: -1, F2: -539}
        _ = v10
        _ = v13
    }
    v14 := GenS3{F0: int64(p0), F1: 1}
    _ = p0
    _ = v14
    return p0
}

func main() {
    var v1 float32 = (0.1 + 1.5)
    var v2 int8 = -128
    var v3 int8 = 0
    var v4 float64 = (0.1 - ((-2.25 + ((1.5 + 1.5) + 1e10)) / (((float64(((-2.25 - ((1e10 + -2.25) - 0.1)) * v1)) - (float64(0.0) * 1.5)) - (1e10 - 0.1)) - (1.5 / 0.1))))
    v5 := []int{817, 723, 60, -110}
    var v6 int32 = -2147483648
    v3 = (-(v2))
    for fi15 := 0; fi15 < 7; fi15++ {
        if (0 > 1) {
            var v16 float64 = ((float64(((0.0 + (v1 - (v1 + (v1 + 0.1)))) - 1.5)) / ((v4 / v4) * 1.5)) + (v4 - v4))
            var v17 int32 = (v6 + -49)
            v2 = 0
            _ = v16
            _ = v17
        }
        fmt.Println(v3, v4, fi15)
        if true {
            fmt.Println(v2)
            for fi18 := 0; fi18 < 9; fi18++ {
                if ("a" < "foo bar") {
                    var v19 uint8 = 2
                    v2 += (v2 >> ((42) & 63))
                    var v20 float64 = 1e10
                    _ = v19
                    _ = v20
                }
                if len(v5) > 0 {
                    v5[0] = fi18
                }
                v3 = (v3 / ((v3) | 1))
                _ = fi18
            }
        }
        _ = fi15
    }
    v4 -= ((v4 * (float64(0.1) / (float64(float32(((float64(float32((v4 / float64(3.14159)))) - v4) * ((float64(v1) / float64(v1)) + (3.14159 / -2.25))))) / 0.1))) + 3.14159)
    var v21 string = "foo bar"
    v1 = (1.5 - (v1 + 1.5))
    var v22 int = 32768
    for fi23 := 0; fi23 < 12; fi23++ {
        v6 = v6
        for fi24 := 0; fi24 < 4; fi24++ {
            switch (v21 + v21) {
            case "":
                var v25 int32 = (v6 >> ((1) & 63))
                var v26 string = v21
                v27 := GenS1{F0: 42, F1: 2, F2: 38}
                _ = v25
                _ = v26
                _ = v27
            default:
                v22 = (-(v22))
            }
            var v28 int = (fi24 & (fi24 % ((fi24) | 1)))
            _ = fi24
            _ = v28
        }
        _ = fi23
    }
    for fi29 := 0; fi29 < 8; fi29++ {
        switch v21 {
        case "a":
            fmt.Println(v4, v4, v22)
        case "hello":
            for fi30 := 0; fi30 < 17; fi30++ {
                switch v21 {
                case "a":
                    v4 = v4
                case "\n\t":
                    v4 = v4
                default:
                    switch 2147483647 {
                    case 4294967295:
                        v31 := GenS4{F0: (v22 & v22), F1: v6, F2: v6}
                        _ = v31
                    }
                    var v32 float32 = 0.1
                    if len(v5) > 0 {
                        v5[0] = (v22 / ((516) | 1))
                    }
                    _ = v32
                }
                if false {
                    if false {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        v33 := GenS5{F0: (^(v3))}
                        if true {
                            if len(v5) > 0 {
                                v5[0] = 9223372036854775807
                            }
                        } else {
                            var v34 uint16 = 2
                            _ = v34
                        }
                        _ = v33
                    } else {
                        var v35 bool = (v1 < (-2.25 - v1))
                        v36 := GenS6{F0: 32767, F1: 2}
                        _ = v35
                        _ = v36
                    }
                    v37 := GenS7{F0: 42, F1: 32767}
                    _ = v37
                } else {
                    var v38 int = v22
                    _ = v38
                }
                if (((v1 * -2.25) >= float32(v4)) || false) {
                    v39 := GenS1{F0: 1, F1: 4294967295, F2: 128}
                    var v40 uint = 42
                    _ = len(v5)
                    _ = v39
                    _ = v40
                }
                _ = fi30
            }
        case "\n\t":
            var v41 uint16 = 255
            _ = v41
        default:
            switch (fi29 + (v22 - 356)) {
            case 9223372036854775807:
                var v42 uint16 = 725
                _ = v42
            case 1:
                for fi43 := 0; fi43 < 20; fi43++ {
                    var v44 uint16 = 255
                    fmt.Println(fi29, v6)
                    _ = fi43
                    _ = v44
                }
            case 523:
                v45 := GenS8{F0: 42, F1: (v6 & (v6 % ((v6) | 1)))}
                _ = v45
            }
        }
        var v46 uint64 = 488
        if len(v5) > 0 {
            v5[0] = (v22 * -9223372036854775808)
        }
        _ = fi29
        _ = v46
    }
    v22 = v22
    var v47 int32 = int32(v3)
    v6 = v6
    v48 := GenS2{F0: v2, F1: 32767, F2: (v22 - v22)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v5
    _ = v6
}
