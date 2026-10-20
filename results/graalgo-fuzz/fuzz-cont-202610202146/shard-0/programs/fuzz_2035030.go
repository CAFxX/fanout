package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 int8 }
type GenS2 struct { F0 int8; F1 int16 }
type GenS3 struct { F0 int8; F1 bool }
type GenS4 struct { F0 uint32; F1 int16; F2 uint8 }
type GenS5 struct { F0 uint64; F1 uint32; F2 uint }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint16 }

func main() {
    v1 := []uint32{2147483648, 1, 2147483648, 2147483647}
    var v2 float64 = 3.14159
    var v3 int32 = 2147483647
    var v4 string = "hello"
    if true {
        var v5 string = v4
        v5 = "\n\t"
        for fi6 := 0; fi6 < 8; fi6++ {
            var v7 uint32 = 42
            _ = fi6
            _ = v7
        }
        _ = v5
    } else {
        if true {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        } else {
            for fi8 := 0; fi8 < 19; fi8++ {
                v9 := GenS1{F0: 201, F1: 0, F2: 100}
                _ = fi8
                _ = v9
            }
            v3 = v3
        }
    }
    if (v4 == "foo bar") {
        for fi10 := 0; fi10 < 16; fi10++ {
            _ = len(v1)
            for fi11 := 0; fi11 < 2; fi11++ {
                if (525 < 1) {
                    if (!true) {
                        var v12 int16 = 0
                        _ = v12
                    } else {
                        if true {
                            var v13 bool = true
                            var v14 float64 = (3.14159 + float64(0.1))
                            if (false || v13) {
                                fmt.Println(fi11, fi10)
                                if (-25 > -128) {
                                    var v15 uint64 = 42
                                    v15 = 42
                                    _ = v15
                                } else {
                                    fmt.Println(v2, fi11, fi10)
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    v16 := GenS1{F0: 42, F1: 42, F2: -18}
                                    v4 = "hello"
                                    _ = v16
                                }
                                v17 := GenS2{F0: -128, F1: 918}
                                _ = v17
                            }
                            _ = v13
                            _ = v14
                        } else {
                            var v18 int16 = 1000
                            _ = v18
                        }
                        v19 := GenS3{F0: 0, F1: (!(v4 >= (v4 + (((v4 + (v4 + "a")) + v4) + v4))))}
                        switch -2147483648 {
                        case -9223372036854775808:
                            v20 := GenS1{F0: 2, F1: 42, F2: 100}
                            _ = v20
                        case 0:
                            fmt.Println(fi10)
                        default:
                            var v21 int8 = 127
                            _ = v21
                        }
                        _ = v19
                    }
                    v22 := GenS4{F0: 2, F1: 1, F2: 255}
                    _ = v22
                } else {
                    _ = len(v1)
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi11
            }
            _ = fi10
        }
    } else {
        if ((0 != 71) && true) {
            v23 := GenS5{F0: 2, F1: 535, F2: 963}
            for fi24 := 0; fi24 < 13; fi24++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi24
            }
            _ = v23
        } else {
            if true {
                var v25 int8 = -1
                if (v4 != ((v4 + v4) + (v4 + (v4 + "")))) {
                    for fi26 := 0; fi26 < 20; fi26++ {
                        var v27 bool = false
                        var v28 float64 = float64(0.1)
                        _ = len(v1)
                        _ = fi26
                        _ = v27
                        _ = v28
                    }
                }
                _ = v25
            } else {
                if (!true) {
                    v2 -= (((((0.1 / v2) / float64(-2.25)) * 0.1) * v2) - ((((v2 - v2) + (float64(0.1) / -2.25)) - (float64((0.0 / 3.14159)) - float64(3.14159))) + (v2 - float64((3.14159 + (1e10 * 1e10))))))
                    var v29 float64 = 0.1
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = v29
                }
                for fi30 := 0; fi30 < 8; fi30++ {
                    for fi31 := 0; fi31 < 5; fi31++ {
                        switch (-(fi31)) {
                        case 9223372036854775807:
                            if (165 <= 1) {
                                var v32 string = "foo bar"
                                v33 := GenS4{F0: 0, F1: 42, F2: 1}
                                _ = v32
                                _ = v33
                            } else {
                                fmt.Println(fi31, fi30, v4)
                            }
                        case 128:
                            v34 := GenS6{F0: 1}
                            _ = v34
                        case 0:
                            fmt.Println(v2)
                        default:
                            var v35 int64 = 42
                            _ = v35
                        }
                        var v36 uint64 = 42
                        fmt.Println(v2, v36, fi30)
                        _ = fi31
                        _ = v36
                    }
                    _ = fi30
                }
                v37 := GenS7{F0: 1}
                _ = v37
            }
            v4 = v4
        }
        if (127 > 42) {
            for fi38 := 0; fi38 < 8; fi38++ {
                var v39 int16 = 0
                _ = fi38
                _ = v39
            }
            fmt.Println(v2, v4)
            v4 = ""
        } else {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v40 int32 = v3
            v40 += (v3 % ((317) | 1))
            _ = v40
        }
        for fi41 := 0; fi41 < 2; fi41++ {
            if len(v1) > 0 {
                v1[0] = 42
            }
            _ = fi41
        }
    }
    v42 := GenS8{F0: 1}
    v2 = 0.0
    var v43 uint8 = 255
    fmt.Println(v4)
    v44 := GenS3{F0: -128, F1: (1 < 42)}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
}
