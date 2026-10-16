package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint32; F1 int16; F2 int }
type GenS5 struct { F0 uint8 }

func main() {
    var v1 uint16 = 2
    v2 := make(map[uint32]string)
    var v3 int64 = 9223372036854775807
    v4 := make(map[int]string)
    var v5 int32 = -446
    v3 -= -1
    fmt.Println(v1, v3)
    v1 = v1
    if (2 < 2) {
        for fi6 := 0; fi6 < 12; fi6++ {
            v5 -= v5
            var v7 uint32 = 897
            v8 := GenS1{F0: 280, F1: 0}
            _ = fi6
            _ = v7
            _ = v8
        }
        v9 := GenS1{F0: 4294967295, F1: 216}
        v3 += 546
        _ = v9
    } else {
        v10 := []int{}
        _ = v10
    }
    fmt.Println(v3, v3)
    var v11 int16 = 1000
    var v12 uint64 = 956
    var v13 int64 = -571
    if ("" < "foo bar") {
        v12 -= (v12 ^ 42)
    } else {
        v14 := []int{}
        var v15 int16 = int16(v5)
        v12 = 245
        _ = v14
        _ = v15
    }
    v11 = v11
    if true {
        v1 = (v1 / ((0) | 1))
        if (!((v11 * v11) != (v11 + (^(v11))))) {
            v16 := GenS1{F0: uint32(v12), F1: 0}
            v3 = (v3 ^ (v3 / (((v3 - v3)) | 1)))
            v1 = (v1 * 942)
            _ = v16
        } else {
            switch "a" {
            case "a":
                v17 := []int{-1, 916, 2147483647, -1}
                _ = v17
            case "foo bar":
                v18 := []int{int(v11), 2147483647, 1}
                _ = v18
            case "\n\t":
                v13 -= 42
            default:
                for fi19 := 0; fi19 < 3; fi19++ {
                    switch (fi19 >> ((58) & 63)) {
                    case 0:
                        switch (fi19 | int(v1)) {
                        case 42:
                            for fi20 := 0; fi20 < 16; fi20++ {
                                v5 = v5
                                var v21 uint64 = (v12 | (^(v12)))
                                v11 -= 1000
                                _ = fi20
                                _ = v21
                            }
                        default:
                            var v22 string = ("hello" + "a")
                            _ = v22
                        }
                    case -2147483648:
                        v1 = v1
                    case 1:
                        switch (^(fi19)) {
                        case 9223372036854775807:
                            var v23 int16 = -257
                            _ = v23
                        case 32767:
                            var v24 int32 = (v5 >> ((65535) & 63))
                            _ = v24
                        case 308:
                            v25 := []int{}
                            _ = v25
                        }
                    default:
                        v5 -= v5
                    }
                    fmt.Println(v12, v5)
                    _ = fi19
                }
            }
        }
    } else {
        _ = len(v2)
        if (262 >= 1) {
            _ = len(v4)
            var v26 int64 = 9223372036854775807
            fmt.Println(v11)
            _ = v26
        } else {
            v27 := []int{}
            _ = v27
        }
    }
    v28 := []int{int(v11)}
    if false {
        v29 := GenS2{F0: 42, F1: 4294967295}
        var v30 float64 = -2.25
        v12 -= v12
        _ = v29
        _ = v30
    } else {
        if false {
            v31 := GenS3{F0: (^(v1))}
            fmt.Println(v11, v1, v5)
            if len(v28) >= 2 {
                v28 = v28[1:len(v28)]
            }
            _ = v31
        }
        var v32 uint = 65535
        for fi33 := 0; fi33 < 5; fi33++ {
            v34 := GenS4{F0: 42, F1: v11, F2: 127}
            _ = fi33
            _ = v34
        }
        _ = v32
    }
    for fi35 := 0; fi35 < 19; fi35++ {
        switch "" {
        case "":
            fmt.Println(v13, v1)
        case "hello":
            if (true && true) {
                delete(v2, 4294967295)
                v28 = append(v28, (fi35 << ((4294967295) & 63)))
            }
        default:
            switch "hello" {
            case "a":
                var v36 int32 = 42
                _ = v36
            case "foo bar":
                for fi37 := 0; fi37 < 4; fi37++ {
                    v38 := GenS5{F0: 127}
                    fmt.Println(v1, v1, v11)
                    _ = fi37
                    _ = v38
                }
            case "hello":
                fmt.Println(v1, v13, v5)
            default:
                for fi39 := 0; fi39 < 4; fi39++ {
                    v1 += 2
                    v11 = v11
                    v40 := GenS2{F0: 255, F1: 588}
                    v41 := GenS5{F0: 127}
                    var v42 string = "a"
                    _ = fi39
                    _ = v40
                    _ = v41
                    _ = v42
                }
            }
        }
        switch "a" {
        case "hello":
            v28 = append(v28, (-(fi35)))
        }
        var v43 uint16 = (v1 / (((^(v1))) | 1))
        _ = fi35
        _ = v43
    }
    fmt.Println(v1, v11, v1)
    v1 = v1
    for fi44 := 0; fi44 < 4; fi44++ {
        v5 += v5
        var v45 uint32 = 2147483647
        _ = fi44
        _ = v45
    }
    for fi46 := 0; fi46 < 8; fi46++ {
        var v47 string = "foo bar"
        var v48 int16 = v11
        _ = fi46
        _ = v47
        _ = v48
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
