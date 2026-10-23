package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string; F1 int16; F2 int16 }
func genFn0(p0 int16, p1 uint32) uint {
    for fi6 := 0; fi6 < 15; fi6++ {
        fmt.Println(p1)
        p0 = p0
        var v7 int8 = 100
        _ = fi6
        _ = v7
    }
    fmt.Println(p1, p1)
    for fi8 := 0; fi8 < 19; fi8++ {
        v9 := GenS1{F0: 1, F1: 145}
        p0 = (^(p0))
        v10 := GenS1{F0: 2, F1: 42}
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 uint32 = 2
    _ = p0
    _ = p1
    _ = v11
    return 4294967295
}
func genFn1(p0 int) string {
    fmt.Println(p0)
    var v12 int32 = 1
    var v13 uint16 = 0
    var v14 string = ""
    _ = p0
    _ = v12
    _ = v13
    _ = v14
    return (v14 + ((v14 + "\n\t") + ("hello" + v14)))
}

func main() {
    var v1 int16 = -936
    v2 := []uint32{2147483648, 502, 2}
    var v3 string = "foo bar"
    v4 := make(map[uint32]uint32)
    var v5 int32 = -1
    for fi15 := 0; fi15 < 19; fi15++ {
        fmt.Println(v3, v5, fi15)
        fmt.Println(v3, v5, v1)
        switch (fi15 ^ -9223372036854775808) {
        case -9223372036854775808:
            for fi16 := 0; fi16 < 3; fi16++ {
                var v17 int16 = v1
                switch -292 {
                case 127:
                    v5 = -1
                case 0:
                    switch -843 {
                    case 159:
                        var v18 bool = false
                        _ = v18
                    case 42:
                        v3 = v3
                    case -92:
                        if false {
                            delete(v4, 368)
                            fmt.Println(v1)
                            if false {
                                _ = len(v2)
                            } else {
                                if (3.14159 == (0.1 * 0.1)) {
                                    v19 := GenS1{F0: 1, F1: 0}
                                    switch v3 {
                                    case "a":
                                        if ((!false) || true) {
                                            v20 := GenS1{F0: 32767, F1: 0}
                                            v17 = v1
                                            if true {
                                                if (0 <= 872) {
                                                    v4[65535] = 42
                                                }
                                                if false {
                                                    var v21 uint32 = 2147483647
                                                    var v22 int = fi16
                                                    var v23 int = fi15
                                                    _ = v21
                                                    _ = v22
                                                    _ = v23
                                                }
                                            }
                                            _ = v20
                                        }
                                    }
                                    var v24 bool = false
                                    _ = v19
                                    _ = v24
                                }
                                v4[2] = 2147483648
                            }
                            var v25 uint = 65535
                            _ = v25
                        } else {
                            var v26 string = v3
                            switch v26 {
                            case "hello":
                                var v27 int = (^(fi16))
                                _ = v27
                            default:
                                v28 := GenS2{F0: 4294967295}
                                _ = v28
                            }
                            _ = v26
                        }
                    }
                case 1:
                    v3 = ((v3 + v3) + ((v3 + "a") + v3))
                }
                v3 = "foo bar"
                _ = fi16
                _ = v17
            }
        case 0:
            v5 += -1
        default:
            switch 811 {
            case -9223372036854775808:
                fmt.Println(v3, v3, fi15)
            case 0:
                v4[65535] = 974
            default:
                fmt.Println(v3)
            }
        }
        _ = fi15
    }
    _ = len(v2)
    switch "foo bar" {
    case "hello":
        var v29 uint16 = 402
        _ = v29
    case "\n\t":
        if (42 > 4294967295) {
            for fi30 := 0; fi30 < 5; fi30++ {
                var v31 int64 = 1
                for fi32 := 0; fi32 < 5; fi32++ {
                    v33 := GenS3{F0: v3, F1: 411, F2: int16(v31)}
                    _ = fi32
                    _ = v33
                }
                var v34 string = "a"
                _ = fi30
                _ = v31
                _ = v34
            }
            for fi35 := 0; fi35 < 7; fi35++ {
                v5 = v5
                _ = len(v4)
                _ = fi35
            }
            for fi36 := 0; fi36 < 16; fi36++ {
                for fi37 := 0; fi37 < 1; fi37++ {
                    v3 = (v3 + (("foo bar" + v3) + (v3 + v3)))
                    v38 := GenS1{F0: 32767, F1: 255}
                    v1 = v1
                    v2 = append(v2, 0)
                    _ = fi37
                    _ = v38
                }
                v39 := v4[2147483648]
                var v40 int16 = v1
                _ = fi36
                _ = v39
                _ = v40
            }
        }
    }
    var v41 int32 = v5
    for fi42 := 0; fi42 < 7; fi42++ {
        v43 := GenS3{F0: v3, F1: -1, F2: v1}
        _ = fi42
        _ = v43
    }
    var v44 float32 = 3.14159
    var v45 float64 = float64(1.5)
    var v46 uint16 = 255
    if ((v1 | (^(v1))) >= 0) {
        var v47 uint32 = 688
        fmt.Println(v3)
        v48 := GenS1{F0: v46, F1: 104}
        _ = v47
        _ = v48
    }
    fmt.Println(v45, v44, v44)
    v44 = v44
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v44
    _ = v45
    _ = v46
    _ = v5
}
