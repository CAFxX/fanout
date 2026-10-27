package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint; F1 uint8; F2 int16 }
type GenS5 struct { F0 uint64; F1 int32; F2 int }
func genFn0(p0 uint64, p1 int8) uint8 {
    for fi4 := 0; fi4 < 16; fi4++ {
        for fi5 := 0; fi5 < 20; fi5++ {
            v6 := []int{}
            p1 += (-(p1))
            switch "hello" {
            case "foo bar":
                v7 := GenS1{F0: 0, F1: 0, F2: 792}
                _ = v7
            case "hello":
                v8 := GenS2{F0: 2}
                _ = v8
            default:
                fmt.Println(fi4)
            }
            _ = fi5
            _ = v6
        }
        _ = fi4
    }
    p1 = 100
    _ = p0
    _ = p1
    return 255
}
func genFn1() string {
    var v9 uint64 = 1
    v10 := []int{0, 127, 127, 127}
    switch "" {
    case "\n\t":
        for fi11 := 0; fi11 < 1; fi11++ {
            v9 = (v9 - (v9 | (^(v9))))
            v9 += (v9 / (((v9 & 1)) | 1))
            _ = fi11
        }
    case "":
        v9 = v9
    case "hello":
        var v12 int16 = -1
        _ = v12
    default:
        var v13 string = ("hello" + "")
        _ = v13
    }
    v9 -= 4294967295
    _ = v10
    _ = v9
    return "\n\t"
}

func main() {
    var v1 uint64 = 4294967295
    var v2 uint = 162
    v3 := make(map[uint32]int)
    v3[622] = -510
    v14 := GenS1{F0: 42, F1: 496, F2: 1}
    v1 = v1
    v15 := []int{128, -9223372036854775808}
    if true {
        if (127 != 32767) {
            if false {
                v16 := GenS2{F0: 4294967295}
                _ = v16
            } else {
                v3[206] = 1
            }
            v2 += v2
            switch 779 {
            case -1:
                var v17 int = -188
                _ = v17
            case -444:
                for fi18 := 0; fi18 < 18; fi18++ {
                    for fi19 := 0; fi19 < 17; fi19++ {
                        if false {
                            v15 = append(v15, fi19)
                        }
                        v20 := GenS2{F0: 2147483647}
                        _ = len(v3)
                        switch (("\n\t" + "foo bar") + ("a" + "hello")) {
                        case "hello":
                            var v21 bool = true
                            _ = v21
                        case "":
                            v22 := GenS1{F0: -2147483648, F1: 0, F2: uint16(v1)}
                            _ = v22
                        case "a":
                            v2 -= v2
                        default:
                            if (int32(v1) >= -519) {
                                fmt.Println(v2)
                                fmt.Println(fi19, v1, fi18)
                                v2 = (^(v2))
                            } else {
                                var v23 string = "\n\t"
                                v2 = 255
                                _ = v23
                            }
                        }
                        _ = fi19
                        _ = v20
                    }
                    _ = fi18
                }
            default:
                if (!false) {
                    v24 := GenS2{F0: 267}
                    v2 = (^(v2))
                    v2 -= 139
                    _ = v24
                } else {
                    var v25 float32 = 1e10
                    var v26 uint16 = 1
                    var v27 int64 = 42
                    _ = v25
                    _ = v26
                    _ = v27
                }
            }
        } else {
            v28 := GenS3{F0: 345}
            v1 = 42
            for fi29 := 0; fi29 < 15; fi29++ {
                fmt.Println(v2, v1, v2)
                for fi30 := 0; fi30 < 17; fi30++ {
                    _ = len(v3)
                    switch "hello" {
                    case "foo bar":
                        v31 := GenS4{F0: v2, F1: 120, F2: -450}
                        _ = v31
                    case "hello":
                        fmt.Println(fi29, fi29, v2)
                    default:
                        v32 := GenS1{F0: 2147483647, F1: 4294967295, F2: 0}
                        _ = v32
                    }
                    v33 := v3[1]
                    _ = fi30
                    _ = v33
                }
                _ = fi29
            }
            _ = v28
        }
        if false {
            v1 = (v1 % ((v1) | 1))
        }
    } else {
        var v34 int8 = -1
        v35 := GenS5{F0: v1, F1: 2147483647, F2: 32767}
        _ = v34
        _ = v35
    }
    delete(v3, uint32(v1))
    for fi36 := 0; fi36 < 16; fi36++ {
        var v37 float64 = 1.5
        _ = fi36
        _ = v37
    }
    var v38 string = "\n\t"
    if false {
        var v39 int16 = -32768
        fmt.Println(v2)
        _ = v39
    } else {
        for fi40 := 0; fi40 < 11; fi40++ {
            var v41 int = fi40
            var v42 uint16 = 789
            _ = fi40
            _ = v41
            _ = v42
        }
    }
    var v43 uint16 = 1
    if (uint(v43) <= (v2 - 786)) {
        for fi44 := 0; fi44 < 2; fi44++ {
            v43 = v43
            v2 -= (^(v2))
            _ = fi44
        }
    } else {
        if false {
            for fi45 := 0; fi45 < 18; fi45++ {
                var v46 uint64 = v1
                v2 += (v2 >> ((v2) & 63))
                _ = fi45
                _ = v46
            }
        }
    }
    fmt.Println(v43, v1)
    delete(v3, 144)
    v47 := GenS4{F0: uint(v43), F1: uint8(v1), F2: 0}
    switch 1 {
    case 0:
        var v48 int32 = 42
        _ = v48
    case 9223372036854775807:
        v38 = (v38 + (v38 + (v38 + v38)))
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v38
    _ = v43
    _ = v47
}
