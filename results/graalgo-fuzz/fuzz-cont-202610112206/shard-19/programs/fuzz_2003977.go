package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 uint16; F1 uint16 }
type GenS4 struct { F0 int16; F1 int; F2 int16 }
type GenS5 struct { F0 uint8; F1 string }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int }
func genFn0(p0 bool, p1 bool, p2 uint8) int16 {
    if (p2 > p2) {
        switch "" {
        case "hello":
            var v6 int8 = 100
            _ = v6
        case "a":
            var v7 int8 = 127
            _ = v7
        }
        p0 = p0
        if (!p1) {
            for fi8 := 0; fi8 < 18; fi8++ {
                for fi9 := 0; fi9 < 7; fi9++ {
                    var v10 uint64 = 2
                    if true {
                        p2 = (p2 >> ((561) & 63))
                        switch 42 {
                        case 32767:
                            var v11 string = ""
                            _ = v11
                        case -9223372036854775808:
                            var v12 float64 = (-2.25 * -2.25)
                            _ = v12
                        case 2147483647:
                            p0 = ((!false) || p0)
                        }
                    } else {
                        v13 := GenS1{F0: (^(p2))}
                        _ = v13
                    }
                    switch (fi9 % (((^(fi9))) | 1)) {
                    case 2147483647:
                        switch "\n\t" {
                        case "hello":
                            if false {
                                p0 = (false || p1)
                                v14 := GenS1{F0: p2}
                                v10 = v10
                                _ = v14
                            } else {
                                p2 = (p2 ^ 128)
                            }
                        case "\n\t":
                            if true {
                                v15 := GenS2{F0: 65535, F1: 1}
                                v16 := GenS3{F0: 0, F1: 65535}
                                _ = v15
                                _ = v16
                            } else {
                                var v17 int8 = 127
                                fmt.Println(v17, p1)
                                var v18 int32 = 1
                                _ = v17
                                _ = v18
                            }
                        }
                    }
                    _ = fi9
                    _ = v10
                }
                _ = fi8
            }
            fmt.Println(p1, p2, p0)
            v19 := GenS3{F0: 835, F1: 893}
            _ = v19
        } else {
            p2 = p2
        }
    }
    switch (("hello" + ("\n\t" + "\n\t")) + "hello") {
    case "a":
        if p0 {
            fmt.Println(p2, p0, p0)
            if (!(uint32(p2) <= 2147483647)) {
                for fi20 := 0; fi20 < 17; fi20++ {
                    var v21 string = "foo bar"
                    p2 = p2
                    var v22 bool = (false && (45 != 9223372036854775807))
                    _ = fi20
                    _ = v21
                    _ = v22
                }
                v23 := GenS1{F0: 42}
                p2 = 2
                _ = v23
            }
            for fi24 := 0; fi24 < 2; fi24++ {
                var v25 uint32 = 2147483647
                fmt.Println(v25, fi24)
                p2 += (p2 % ((uint8(v25)) | 1))
                _ = fi24
                _ = v25
            }
        }
    default:
        if p0 {
            if p1 {
                p1 = (!((^(p2)) != 1))
                var v26 uint8 = p2
                for fi27 := 0; fi27 < 17; fi27++ {
                    v28 := GenS4{F0: 1, F1: (fi27 | (fi27 ^ (fi27 << ((42) & 63)))), F2: 42}
                    _ = fi27
                    _ = v28
                }
                _ = v26
            } else {
                p1 = ((2 != 1) && p0)
                p0 = p1
                v29 := GenS5{F0: 2, F1: "foo bar"}
                _ = v29
            }
            v30 := []int{4294967295, -2147483648}
            if len(v30) >= 2 {
                v30 = v30[1:len(v30)]
            }
            _ = v30
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return -32768
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint8 = 1
    var v3 string = "\n\t"
    v4 := []uint32{1}
    var v5 float64 = 3.14159
    v31 := GenS6{F0: 1}
    _ = len(v1)
    v2 = (^(v2))
    v32 := GenS7{F0: 0}
    var v33 int8 = 42
    if (824 > 295) {
        fmt.Println(v5, v5, v33)
    }
    v34 := v1[693]
    for fi35 := 0; fi35 < 13; fi35++ {
        if false {
            var v36 float64 = float64(3.14159)
            v33 += (v33 / (((v33 & 11)) | 1))
            _ = v36
        }
        var v37 uint16 = 255
        fmt.Println(v37)
        _ = fi35
        _ = v37
    }
    var v38 float32 = 0.1
    var v39 int8 = v33
    var v40 int64 = 42
    for fi41 := 0; fi41 < 1; fi41++ {
        for fi42 := 0; fi42 < 15; fi42++ {
            if true {
                v43 := GenS1{F0: v2}
                v4 = append(v4, 0)
                var v44 int16 = 1
                _ = v43
                _ = v44
            } else {
                var v45 bool = true
                _ = v45
            }
            v34 = "hello"
            _ = fi42
        }
        _ = fi41
    }
    var v46 int16 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v46
    _ = v5
}
