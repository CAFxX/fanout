package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint; F1 bool }
type GenS4 struct { F0 uint; F1 uint64 }
type GenS5 struct { F0 bool; F1 int32 }

func main() {
    var v1 uint32 = 507
    var v2 uint64 = 4294967295
    v3 := make(map[int]int)
    v4 := []uint32{(v1 ^ v1), v1, (^(v1))}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v2, v1)
    if (!(((-2.25 * float64(0.0)) / ((float64(1e10) * 1e10) + (float64(((0.0 / -2.25) * (((0.0 / 3.14159) * 1e10) / 3.14159))) / 3.14159))) >= 1e10)) {
        v1 = uint32(v2)
    } else {
        v5 := v3[4294967295]
        v2 = v2
        _ = v5
    }
    _ = len(v4)
    for fi6 := 0; fi6 < 9; fi6++ {
        var v7 uint32 = v1
        fmt.Println(v7, fi6)
        v3[1] = fi6
        _ = fi6
        _ = v7
    }
    fmt.Println(v2, v1)
    v8 := GenS1{F0: "", F1: "\n\t"}
    var v9 float64 = 0.1
    if false {
        for fi10 := 0; fi10 < 18; fi10++ {
            for fi11 := 0; fi11 < 17; fi11++ {
                var v12 int = fi11
                v9 += ((1e10 - 0.1) + 3.14159)
                _ = fi11
                _ = v12
            }
            switch "\n\t" {
            case "a":
                var v13 int8 = 104
                _ = v13
            }
            _ = fi10
        }
        for fi14 := 0; fi14 < 12; fi14++ {
            var v15 int = -2147483648
            if true {
                switch v15 {
                case -962:
                    switch (v15 ^ 4294967295) {
                    case 32768:
                        for fi16 := 0; fi16 < 16; fi16++ {
                            if (!true) {
                                if true {
                                    v4 = append(v4, (v1 & 837))
                                    switch (fi14 & v15) {
                                    case 832:
                                        v17 := GenS1{F0: "\n\t", F1: ""}
                                        _ = v17
                                    case 4294967295:
                                        var v18 uint32 = 2147483647
                                        _ = v18
                                    case 32768:
                                        v9 = (0.1 - v9)
                                    default:
                                        var v19 float64 = (3.14159 / (0.0 - (v9 - -2.25)))
                                        _ = v19
                                    }
                                } else {
                                    v2 = 4294967295
                                    v1 -= (v1 | (v1 / ((v1) | 1)))
                                }
                            } else {
                                if true {
                                    if (-2147483648 <= 353) {
                                        var v20 bool = (-653 >= -743)
                                        if v20 {
                                            fmt.Println(v1)
                                        }
                                        _ = v20
                                    }
                                    var v21 uint16 = 150
                                    v15 -= fi16
                                    _ = v21
                                } else {
                                    v22 := GenS2{F0: ("" + "")}
                                    v23 := GenS1{F0: "\n\t", F1: ""}
                                    _ = v22
                                    _ = v23
                                }
                                switch -9223372036854775808 {
                                case 32768:
                                    var v24 uint16 = 1
                                    _ = v24
                                default:
                                    v4 = append(v4, 4294967295)
                                }
                            }
                            if true {
                                if (!false) {
                                    if (255 == 1) {
                                        fmt.Println(v2, v9, fi14)
                                        delete(v3, fi16)
                                    }
                                    var v25 int32 = 782
                                    fmt.Println(v9)
                                    _ = v25
                                } else {
                                    if ((false && true) && (v2 > (v2 * v2))) {
                                        fmt.Println(v1, v15, v15)
                                        v26 := GenS3{F0: 733, F1: (false || false)}
                                        _ = v26
                                    }
                                    var v27 uint8 = uint8(v2)
                                    _ = v27
                                }
                                v1 = 1
                                v28 := GenS4{F0: 4294967295, F1: 2}
                                _ = v28
                            }
                            fmt.Println(fi16, v2)
                            if false {
                                var v29 uint64 = v2
                                v29 = 4294967295
                                v30 := GenS2{F0: ""}
                                _ = v29
                                _ = v30
                            } else {
                                delete(v3, v15)
                                fmt.Println(v9, v2, v15)
                            }
                            switch ("a" + "foo bar") {
                            case "hello":
                                var v31 int32 = 0
                                _ = v31
                            case "foo bar":
                                fmt.Println(fi16)
                            }
                            _ = fi16
                        }
                    default:
                        v32 := v3[v15]
                        _ = v32
                    }
                case 375:
                    var v33 int16 = 0
                    _ = v33
                case 32767:
                    var v34 uint32 = (v1 >> ((1) & 63))
                    _ = v34
                default:
                    v35 := GenS2{F0: "a"}
                    _ = v35
                }
                var v36 float64 = (float64(3.14159) - v9)
                _ = v36
            }
            if (-272 == 42) {
                v9 += v9
            }
            _ = fi14
            _ = v15
        }
    } else {
        v2 = (v2 >> ((uint(v2)) & 63))
    }
    v9 += v9
    v2 = v2
    fmt.Println(v2, v9, v1)
    v1 = (^(v1))
    if len(v4) > 0 {
        v4[0] = 1
    }
    v1 -= v1
    v9 = v9
    for fi37 := 0; fi37 < 2; fi37++ {
        v38 := GenS5{F0: false, F1: 2147483647}
        _ = fi37
        _ = v38
    }
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
