package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int16; F1 int32; F2 int }
type GenS5 struct { F0 uint8; F1 uint; F2 uint64 }
type GenS6 struct { F0 int64; F1 uint64; F2 int }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]int)
    v3 := []int{9223372036854775807, -778, 255}
    var v4 int8 = -128
    var v5 int64 = int64(v4)
    v6 := GenS1{F0: v4, F1: 724, F2: 2}
    fmt.Println(v4)
    var v7 uint = 0
    for fi8 := 0; fi8 < 11; fi8++ {
        if true {
            if (false && true) {
                for fi9 := 0; fi9 < 4; fi9++ {
                    v5 = v5
                    _ = fi9
                }
                if false {
                    for fi10 := 0; fi10 < 1; fi10++ {
                        if false {
                            var v11 float32 = (-2.25 / -2.25)
                            _ = v11
                        } else {
                            if len(v3) > 0 {
                                v3[0] = (fi8 << (((v7 / ((v7) | 1))) & 63))
                            }
                        }
                        _ = fi10
                    }
                    v4 = (v4 / ((v4) | 1))
                    var v12 float64 = (0.1 + 0.0)
                    _ = v12
                }
                v13 := v2["\n\t"]
                _ = v13
            }
        }
        var v14 int16 = int16(v5)
        _ = fi8
        _ = v14
    }
    var v15 uint32 = 0
    if (-9223372036854775808 > 466) {
        v16 := GenS2{F0: 2147483647}
        _ = v16
    } else {
        for fi17 := 0; fi17 < 10; fi17++ {
            if true {
                v4 = v4
                fmt.Println(v7)
                v18 := GenS3{F0: 687}
                _ = v18
            }
            fmt.Println(fi17, fi17)
            _ = fi17
        }
        v15 -= (v15 >> (((v7 | (^(v7)))) & 63))
        for fi19 := 0; fi19 < 15; fi19++ {
            v15 -= (v15 * v15)
            _ = fi19
        }
    }
    var v20 string = "foo bar"
    for fi21 := 0; fi21 < 1; fi21++ {
        for fi22 := 0; fi22 < 6; fi22++ {
            v5 += (v5 >> (((v7 - (v7 >> ((v7) & 63)))) & 63))
            fmt.Println(v20, v20)
            _ = fi22
        }
        delete(v2, v20)
        var v23 uint8 = 42
        _ = fi21
        _ = v23
    }
    var v24 float64 = 0.0
    var v25 int64 = (v5 >> (((v7 - 0)) & 63))
    if true {
        v3 = append(v3, 910)
        var v26 int64 = v25
        if len(v3) > 0 {
            v3[0] = 1
        }
        _ = v26
    } else {
        switch v20 {
        case "foo bar":
            v4 += (v4 << ((18446744073709551615) & 63))
        case "":
            for fi27 := 0; fi27 < 10; fi27++ {
                v24 = v24
                if true {
                    v28 := GenS4{F0: -408, F1: -522, F2: fi27}
                    _ = v28
                } else {
                    for fi29 := 0; fi29 < 3; fi29++ {
                        switch (v20 + v20) {
                        case "a":
                            if false {
                                v15 += v15
                                var v30 bool = (v15 >= 853)
                                v31 := GenS5{F0: 132, F1: (^(v7)), F2: 2}
                                _ = v30
                                _ = v31
                            }
                        case "hello":
                            if (fi29 > fi29) {
                                v15 -= (v15 - (v15 % ((4294967295) | 1)))
                                var v32 uint32 = (^(v15))
                                _ = v32
                            }
                        case "foo bar":
                            if (9223372036854775807 > (v25 + -629)) {
                                v4 += v4
                                v1[(v15 & 2147483648)] = true
                                fmt.Println(fi29)
                            } else {
                                var v33 uint8 = 255
                                v15 += v15
                                if (1 != int16(v7)) {
                                    v3 = append(v3, fi27)
                                    var v34 uint8 = v33
                                    var v35 uint8 = (v33 + (v33 / ((v34) | 1)))
                                    _ = v34
                                    _ = v35
                                } else {
                                    switch (fi29 >> ((v7) & 63)) {
                                    case 9223372036854775807:
                                        if false {
                                            var v36 uint = uint(v33)
                                            var v37 int16 = 0
                                            _ = v36
                                            _ = v37
                                        }
                                    default:
                                        fmt.Println(v15)
                                    }
                                    var v38 int32 = 42
                                    v4 = v4
                                    v20 = ((((v20 + ((v20 + v20) + "foo bar")) + (v20 + v20)) + (v20 + v20)) + v20)
                                    var v39 float64 = ((float64((0.0 + 3.14159)) / v24) * (float64((0.1 / 1.5)) / (v24 + (v24 * (v24 - -2.25)))))
                                    if (42 != 0) {
                                        v15 += (v15 * v15)
                                    }
                                    v24 += (v39 / v24)
                                    _ = v38
                                    _ = v39
                                }
                                _ = v33
                            }
                        }
                        switch (("hello" + ("a" + ((v20 + v20) + v20))) + "foo bar") {
                        case "hello":
                            var v40 uint8 = 223
                            _ = v40
                        case "\n\t":
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        default:
                            v41 := GenS4{F0: 1000, F1: -582, F2: 2147483647}
                            _ = v41
                        }
                        v7 += (v7 % ((v7) | 1))
                        _ = fi29
                    }
                }
                var v42 float64 = (float64(float32((((v24 / (float64(1.5) + (v24 / 1.0))) + v24) + v24))) + 1e10)
                _ = fi27
                _ = v42
            }
        }
        var v43 string = ""
        _ = v43
    }
    if (-668 > 1000) {
        v44 := GenS3{F0: (v15 | (v15 + 65535))}
        _ = v44
    }
    v45 := GenS6{F0: (v25 << ((v7) & 63)), F1: 1, F2: 2147483647}
    v46 := GenS2{F0: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v5
    _ = v6
    _ = v7
}
