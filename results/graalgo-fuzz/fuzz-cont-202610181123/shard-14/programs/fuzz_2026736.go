package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 bool }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 int32; F1 int32 }
type GenS4 struct { F0 int; F1 string; F2 uint16 }
func genFn0(p0 uint16, p1 int16) string {
    var v6 int32 = -2147483648
    fmt.Println(v6)
    var v7 uint16 = (p0 * p0)
    var v8 uint8 = 128
    v9 := []int{42, 299}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return "hello"
}
func genFn1(p0 int32) int8 {
    for fi10 := 0; fi10 < 16; fi10++ {
        v11 := GenS1{F0: true, F1: -84, F2: (2147483648 <= 42)}
        var v12 int64 = 0
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 uint64 = 42
    _ = p0
    _ = v13
    return -1
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]string)
    var v3 uint32 = 42
    v4 := make(map[uint32]uint32)
    v5 := []int{}
    v14 := GenS1{F0: (42 == 2), F1: 1, F2: true}
    var v15 int = 1
    var v16 uint = 42
    if (true && (!((true || (1000 != 1000)) || false))) {
        v17 := GenS2{F0: 2, F1: -1}
        v16 -= 22
        v16 -= v16
        _ = v17
    }
    if false {
        var v18 int8 = -128
        _ = v18
    }
    _ = len(v2)
    if (100 <= 41) {
        for fi19 := 0; fi19 < 19; fi19++ {
            if false {
                v15 -= (v15 * (v15 | v15))
                if len(v5) > 0 {
                    v5[0] = (fi19 << (((v16 & (v16 * v16))) & 63))
                }
                for fi20 := 0; fi20 < 19; fi20++ {
                    var v21 int16 = 1
                    v22 := GenS3{F0: -2147483648, F1: 1}
                    _ = fi20
                    _ = v21
                    _ = v22
                }
            } else {
                fmt.Println(v16)
            }
            v3 += 1
            _ = len(v5)
            _ = fi19
        }
        if ("" == "hello") {
            var v23 int8 = -110
            _ = v23
        }
        for fi24 := 0; fi24 < 6; fi24++ {
            if len(v5) > 0 {
                v5[0] = fi24
            }
            switch "a" {
            case "":
                if true {
                    if (-82 > 127) {
                        var v25 uint32 = 118
                        if false {
                            for fi26 := 0; fi26 < 8; fi26++ {
                                var v27 float64 = 3.14159
                                v15 = -2147483648
                                _ = fi26
                                _ = v27
                            }
                            for fi28 := 0; fi28 < 11; fi28++ {
                                v25 -= (^(v25))
                                if false {
                                    var v29 int16 = 42
                                    if (1 <= genFn1(-822)) {
                                        if (!(!false)) {
                                            v3 -= v25
                                            var v30 float64 = -2.25
                                            v31 := GenS4{F0: (-(v15)), F1: "hello", F2: 42}
                                            _ = v30
                                            _ = v31
                                        }
                                    } else {
                                        var v32 uint = (v16 ^ v16)
                                        _ = v32
                                    }
                                    _ = v29
                                } else {
                                    var v33 uint32 = 1
                                    _ = v33
                                }
                                _ = fi28
                            }
                            v34 := GenS2{F0: v16, F1: 1}
                            _ = v34
                        } else {
                            var v35 int16 = -1
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            _ = v35
                        }
                        v36 := GenS3{F0: -1, F1: 145}
                        _ = v25
                        _ = v36
                    }
                    switch (("a" + "hello") + ("" + "")) {
                    case "hello":
                        var v37 int = (v15 * fi24)
                        _ = v37
                    case "a":
                        for fi38 := 0; fi38 < 8; fi38++ {
                            if ((true && true) && true) {
                                v3 += (v3 << (((v16 & (v16 | 0))) & 63))
                            }
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            _ = fi38
                        }
                    case "":
                        var v39 int16 = -1
                        _ = v39
                    }
                    v16 = (^(v16))
                }
            case "foo bar":
                if len(v5) > 0 {
                    v5[0] = (-(v15))
                }
            default:
                if false {
                    if (false || (4294967295 >= 4294967295)) {
                        switch "" {
                        case "\n\t":
                            var v40 uint8 = 245
                            _ = v40
                        case "hello":
                            for fi41 := 0; fi41 < 17; fi41++ {
                                if len(v5) > 0 {
                                    v5[0] = (fi24 << (((v16 / (((v16 ^ v16)) | 1))) & 63))
                                }
                                v15 -= fi24
                                _ = fi41
                            }
                        case "":
                            var v42 int16 = 42
                            _ = v42
                        }
                        fmt.Println(v15, v16)
                    }
                    var v43 float32 = (float32(1.5) * (float32(0.1) - 0.1))
                    _ = len(v5)
                    _ = v43
                }
            }
            _ = fi24
        }
    }
    fmt.Println(v3, v3, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
