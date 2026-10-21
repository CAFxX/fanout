package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint16; F1 uint64; F2 int32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint16; F1 int64 }

func main() {
    var v1 uint16 = 3
    v2 := make(map[string]string)
    v3 := []string{}
    var v4 int = -526
    v5 := make(map[int]uint32)
    v1 -= (v1 / ((2) | 1))
    if (!true) {
        v4 = v4
        v6 := GenS1{F0: "foo bar"}
        v4 -= 127
        _ = v6
    } else {
        if false {
            if false {
                v4 = (v4 << ((461) & 63))
                v4 = v4
                v1 += 2
            } else {
                var v7 uint8 = 217
                var v8 int8 = 0
                _ = v7
                _ = v8
            }
        }
        var v9 float64 = (0.1 + float64((0.0 * 0.1)))
        _ = v9
    }
    switch "foo bar" {
    case "hello":
        v10 := GenS1{F0: ""}
        _ = v10
    case "":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case "foo bar":
        v4 += -808
    default:
        var v11 int8 = 100
        _ = v11
    }
    delete(v2, ("" + ""))
    _ = len(v3)
    var v12 uint32 = 2
    if false {
        v12 = v12
        var v13 string = "a"
        v12 += 42
        _ = v13
    }
    for fi14 := 0; fi14 < 9; fi14++ {
        for fi15 := 0; fi15 < 13; fi15++ {
            fmt.Println(v1, v12, v1)
            v16 := GenS2{F0: -133}
            _ = fi15
            _ = v16
        }
        var v17 float32 = (0.0 - -2.25)
        if (0 < (v1 >> ((554) & 63))) {
            if true {
                for fi18 := 0; fi18 < 13; fi18++ {
                    fmt.Println(v12)
                    _ = fi18
                }
                v4 = fi14
                for fi19 := 0; fi19 < 8; fi19++ {
                    if len(v3) > 0 {
                        v3[0] = "a"
                    }
                    _ = fi19
                }
            }
            if (-33 != 0) {
                if true {
                    fmt.Println(v17, v12, v17)
                    v20 := v2[""]
                    var v21 float32 = 0.1
                    _ = v20
                    _ = v21
                } else {
                    if len(v3) > 0 {
                        v3[0] = ("" + "\n\t")
                    }
                    var v22 uint32 = v12
                    var v23 int64 = 9223372036854775807
                    _ = v22
                    _ = v23
                }
                switch "a" {
                case "\n\t":
                    var v24 float32 = 1.5
                    _ = v24
                case "foo bar":
                    v4 = fi14
                case "a":
                    switch fi14 {
                    case 2147483647:
                        v12 -= v12
                    case -437:
                        _ = len(v2)
                    }
                default:
                    v25 := GenS3{F0: (v1 % (((v1 ^ (^(v1)))) | 1)), F1: 18446744073709551615, F2: -148}
                    _ = v25
                }
                v1 -= v1
            } else {
                v26 := GenS1{F0: "foo bar"}
                v5[(v4 | v4)] = (v12 - 2)
                v27 := GenS4{F0: ("" + (("hello" + "\n\t") + "a"))}
                _ = v26
                _ = v27
            }
        } else {
            v1 = uint16(v12)
        }
        _ = fi14
        _ = v17
    }
    v28 := GenS4{F0: "hello"}
    var v29 string = "\n\t"
    switch 42 {
    case -536:
        if len(v3) > 0 {
            v3[0] = "foo bar"
        }
    }
    v4 = (v4 & -9223372036854775808)
    v30 := GenS5{F0: v1, F1: -175}
    if (4294967295 >= 0) {
        fmt.Println(v4, v4, v1)
        var v31 uint32 = 65535
        _ = v31
    } else {
        if (true || true) {
            v1 -= v1
            v29 = "\n\t"
            fmt.Println(v4, v29, v4)
        }
    }
    fmt.Println(v12)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
