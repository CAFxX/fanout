package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int32 }
func genFn0(p0 uint16) int8 {
    v5 := GenS1{F0: 56}
    if true {
        fmt.Println(p0)
    }
    p0 += p0
    var v6 float32 = -2.25
    v7 := GenS1{F0: -1}
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 bool) int64 {
    switch "" {
    case "a":
        p0 = false
    case "foo bar":
        var v8 float64 = -2.25
        _ = v8
    }
    v9 := []int{}
    _ = p0
    _ = v9
    return -880
}

func main() {
    var v1 float32 = (1.5 - 0.1)
    var v2 int16 = 1
    v3 := []string{}
    var v4 uint = 255
    var v10 int64 = 559
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v10 = v10
    var v11 uint32 = 0
    var v12 int8 = -74
    v10 = (-(v10))
    v13 := GenS1{F0: v2}
    v3 = append(v3, "hello")
    v2 -= v2
    v14 := GenS2{F0: -855, F1: (v10 % ((v10) | 1))}
    if ((v11 | 7) == (v11 - (v11 ^ (v11 | (v11 - v11))))) {
        for fi15 := 0; fi15 < 13; fi15++ {
            switch "a" {
            case "hello":
                if (v12 >= 1) {
                    v16 := GenS3{F0: "foo bar"}
                    v10 = v10
                    var v17 uint64 = 1
                    _ = v16
                    _ = v17
                }
            case "\n\t":
                v10 = (v10 / ((0) | 1))
            case "foo bar":
                v1 = (v1 - (((((1e10 + (v1 * v1)) * -2.25) * v1) + (v1 / 1.5)) + float32(3.14159)))
            }
            fmt.Println(fi15, fi15)
            _ = fi15
        }
        v12 += 81
    } else {
        _ = len(v3)
    }
    switch "a" {
    case "\n\t":
        for fi18 := 0; fi18 < 7; fi18++ {
            v12 += (-(v12))
            for fi19 := 0; fi19 < 16; fi19++ {
                v20 := GenS3{F0: "foo bar"}
                var v21 bool = (((359 >= (v10 % ((42) | 1))) && (!false)) || (float32(0.0) <= ((float32((((0.0 - (1e10 + 1e10)) * 0.1) * 1.5)) * v1) * ((1.5 * (float32(3.14159) + (((((float32(0.1) / float32(-2.25)) - ((float32((0.0 * ((1e10 + 1.5) * 3.14159))) * v1) + ((-2.25 * float32(0.1)) / v1))) + (0.1 / 1.0)) + (0.1 * (v1 - (v1 + -2.25)))) * 0.0))) * v1))))
                if (genFn1((!v21)) != (v10 ^ genFn1((v11 >= 2147483648)))) {
                    v22 := GenS1{F0: (v2 * -1)}
                    _ = v22
                } else {
                    var v23 float64 = 3.14159
                    if (!v21) {
                        if false {
                            switch "a" {
                            case "\n\t":
                                v10 = 1
                            case "hello":
                                var v24 uint8 = uint8(v11)
                                _ = v24
                            }
                            v21 = v21
                            if true {
                                v1 = (v1 - v1)
                                switch fi19 {
                                case 119:
                                    v25 := GenS1{F0: (v2 / ((v2) | 1))}
                                    if (4294967295 == 2) {
                                        fmt.Println(v1, v21, fi18)
                                    } else {
                                        v3 = append(v3, "")
                                        if false {
                                            v11 = (v11 >> ((v4) & 63))
                                            var v26 float64 = (((v23 + 3.14159) + float64(v1)) * 1.5)
                                            v27 := GenS1{F0: (^(v2))}
                                            _ = v26
                                            _ = v27
                                        } else {
                                            if ((v2 * -373) == v2) {
                                                v28 := GenS4{F0: -1}
                                                _ = v28
                                            }
                                        }
                                    }
                                    var v29 uint32 = (v11 / ((v11) | 1))
                                    var v30 uint64 = 4294967295
                                    fmt.Println(v23, v1, v30)
                                    _ = v25
                                    _ = v29
                                    _ = v30
                                case 32767:
                                    v31 := GenS2{F0: 248, F1: (v10 - int64(v11))}
                                    _ = v31
                                }
                            }
                        }
                    }
                    var v32 uint8 = uint8(v4)
                    _ = v23
                    _ = v32
                }
                fmt.Println(v4, v12)
                _ = fi19
                _ = v20
                _ = v21
            }
            _ = fi18
        }
    }
    var v33 int16 = 868
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v33
    _ = v4
}
