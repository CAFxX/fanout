package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 string; F1 bool; F2 int8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint16; F1 string }
func genFn0(p0 uint32) string {
    v5 := []int{128, -1}
    v5 = append(v5, 4294967295)
    switch "foo bar" {
    case "\n\t":
        p0 = (p0 & (p0 % ((0) | 1)))
    case "hello":
        p0 += 42
    case "foo bar":
        for fi6 := 0; fi6 < 1; fi6++ {
            if len(v5) > 0 {
                v5[0] = 161
            }
            p0 = (p0 ^ (p0 + 2147483648))
            switch fi6 {
            case 9223372036854775807:
                p0 += (p0 >> ((18446744073709551615) & 63))
            }
            _ = fi6
        }
    default:
        p0 = 2
    }
    _ = p0
    _ = v5
    return ""
}
func genFn1(p0 string, p1 int64) uint {
    var v7 uint16 = 1
    if false {
        if true {
            for fi8 := 0; fi8 < 4; fi8++ {
                for fi9 := 0; fi9 < 10; fi9++ {
                    switch (fi9 + (fi8 * (fi8 % ((fi8) | 1)))) {
                    case 32768:
                        p0 = p0
                    case -946:
                        v10 := []int{(fi9 / ((fi8) | 1)), fi8, fi9, 1}
                        _ = v10
                    default:
                        v11 := GenS1{F0: 324, F1: 0}
                        v12 := []int{(fi8 * fi9), fi9, fi9}
                        p1 += p1
                        p1 -= p1
                        _ = v11
                        _ = v12
                    }
                    _ = fi9
                }
                var v13 int16 = 1000
                _ = fi8
                _ = v13
            }
        } else {
            if false {
                if true {
                    for fi14 := 0; fi14 < 9; fi14++ {
                        for fi15 := 0; fi15 < 15; fi15++ {
                            p0 = "\n\t"
                            _ = fi15
                        }
                        var v16 string = ("hello" + "")
                        v17 := []int{int(p1), int(p1), fi14, -976}
                        _ = fi14
                        _ = v16
                        _ = v17
                    }
                } else {
                    for fi18 := 0; fi18 < 3; fi18++ {
                        for fi19 := 0; fi19 < 17; fi19++ {
                            var v20 uint64 = 2
                            var v21 int8 = 42
                            _ = fi19
                            _ = v20
                            _ = v21
                        }
                        if ("\n\t" < p0) {
                            v22 := []int{-317, fi18}
                            for fi23 := 0; fi23 < 14; fi23++ {
                                fmt.Println(p1, fi23)
                                _ = fi23
                            }
                            _ = v22
                        } else {
                            var v24 uint = 255
                            v25 := GenS2{F0: "hello", F1: false, F2: 100}
                            _ = v24
                            _ = v25
                        }
                        v26 := GenS3{F0: 42}
                        _ = fi18
                        _ = v26
                    }
                }
            }
            var v27 int64 = (p1 % ((p1) | 1))
            _ = v27
        }
    }
    v28 := GenS2{F0: p0, F1: false, F2: -1}
    v29 := []int{}
    _ = p0
    _ = p1
    _ = v28
    _ = v29
    _ = v7
    return 1
}

func main() {
    var v1 int16 = 585
    var v2 float64 = (1e10 - 3.14159)
    var v3 uint64 = 4294967295
    var v4 uint16 = 0
    v4 = v4
    v30 := []int{711, 667, 128}
    var v31 uint64 = (v3 ^ 42)
    v1 = v1
    v31 = 2
    _ = len(v30)
    var v32 int16 = v1
    fmt.Println(v4, v4)
    var v33 int = 32768
    switch v33 {
    case 2147483647:
        fmt.Println(v3, v32, v31)
    case 0:
        for fi34 := 0; fi34 < 20; fi34++ {
            fmt.Println(v33, v3, v2)
            v32 += v1
            _ = fi34
        }
    }
    v32 += -32768
    if len(v30) >= 2 {
        v30 = v30[1:len(v30)]
    }
    var v35 uint8 = 128
    v36 := GenS4{F0: (v4 % ((uint16(v35)) | 1)), F1: "\n\t"}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v35
    _ = v36
    _ = v4
}
