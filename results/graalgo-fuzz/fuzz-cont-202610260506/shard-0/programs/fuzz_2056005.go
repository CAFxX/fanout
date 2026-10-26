package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 bool }
type GenS2 struct { F0 uint16; F1 int8; F2 bool }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int; F1 int16 }
type GenS5 struct { F0 int8; F1 uint32; F2 uint32 }
func genFn0(p0 int32, p1 uint, p2 uint16) bool {
    for fi6 := 0; fi6 < 5; fi6++ {
        p1 -= 97
        _ = fi6
    }
    var v7 int32 = p0
    v7 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return false
}

func main() {
    var v1 uint = 2
    var v2 uint = v1
    v3 := []uint32{}
    var v4 uint16 = 2
    v5 := []string{}
    if false {
        for fi8 := 0; fi8 < 9; fi8++ {
            var v9 uint8 = 1
            _ = fi8
            _ = v9
        }
    }
    if false {
        v10 := GenS1{F0: 0, F1: "hello", F2: (2147483647 < -1)}
        if true {
            for fi11 := 0; fi11 < 7; fi11++ {
                var v12 float32 = 1.5
                _ = fi11
                _ = v12
            }
            v13 := GenS1{F0: 0, F1: "foo bar", F2: true}
            _ = v13
        }
        _ = v10
    } else {
        var v14 int8 = -1
        switch "" {
        case "\n\t":
            for fi15 := 0; fi15 < 1; fi15++ {
                for fi16 := 0; fi16 < 14; fi16++ {
                    switch (-(fi16)) {
                    case 128:
                        var v17 bool = true
                        _ = v17
                    case 32768:
                        v18 := GenS1{F0: 729, F1: "a", F2: ((v14 <= (v14 - v14)) || false)}
                        _ = v18
                    case 32767:
                        fmt.Println(fi16)
                    }
                    var v19 int64 = 569
                    var v20 bool = (-1 != 216)
                    _ = fi16
                    _ = v19
                    _ = v20
                }
                v21 := GenS2{F0: v4, F1: 100, F2: ((((v4 / ((v4) | 1)) != (v4 << ((v2) & 63))) && (((v2 | (v2 - v1)) >= 42) && (v4 > v4))) && true)}
                v4 -= 255
                _ = fi15
                _ = v21
            }
        default:
            var v22 float32 = (0.0 / 3.14159)
            _ = v22
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v14
    }
    if (!false) {
        v23 := GenS3{F0: 0}
        _ = v23
    } else {
        var v24 bool = (1e10 > -2.25)
        v5 = append(v5, "a")
        var v25 int = -9223372036854775808
        _ = v24
        _ = v25
    }
    v1 += 316
    if true {
        fmt.Println(v4, v1)
        var v26 int = -491
        _ = v26
    }
    if true {
        for fi27 := 0; fi27 < 9; fi27++ {
            for fi28 := 0; fi28 < 2; fi28++ {
                v1 = v1
                var v29 int16 = -32768
                _ = fi28
                _ = v29
            }
            _ = fi27
        }
    }
    if false {
        if false {
            fmt.Println(v2)
            v5 = append(v5, ("\n\t" + "\n\t"))
            for fi30 := 0; fi30 < 7; fi30++ {
                var v31 float64 = 1.5
                v31 = v31
                _ = fi30
                _ = v31
            }
        }
    }
    v32 := GenS4{F0: 0, F1: 0}
    _ = len(v5)
    for fi33 := 0; fi33 < 12; fi33++ {
        if (463 > 1) {
            if ((v4 <= (v4 - 32767)) || false) {
                v34 := GenS3{F0: -128}
                switch fi33 {
                case 128:
                    if len(v5) > 0 {
                        v5[0] = ("" + "a")
                    }
                case -9223372036854775808:
                    if len(v3) > 0 {
                        v3[0] = uint32(v4)
                    }
                case 42:
                    switch "a" {
                    case "":
                        var v35 int64 = -1
                        _ = v35
                    case "\n\t":
                        v36 := GenS1{F0: 2147483648, F1: "\n\t", F2: false}
                        _ = v36
                    case "hello":
                        v2 = (^(v2))
                    }
                default:
                    var v37 uint16 = (v4 % ((v4) | 1))
                    _ = v37
                }
                _ = v34
            } else {
                v1 = 255
                v4 -= v4
                v1 = (v1 >> ((v1) & 63))
            }
        } else {
            fmt.Println(fi33, fi33, fi33)
            var v38 float64 = 3.14159
            if false {
                v39 := GenS3{F0: -128}
                _ = v39
            }
            _ = v38
        }
        v4 = (^(v4))
        v4 = 518
        _ = fi33
    }
    v40 := GenS5{F0: 61, F1: 4294967295, F2: 696}
    switch "foo bar" {
    case "hello":
        v1 += v2
    case "\n\t":
        v3 = append(v3, 4294967295)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v40
    _ = v5
}
