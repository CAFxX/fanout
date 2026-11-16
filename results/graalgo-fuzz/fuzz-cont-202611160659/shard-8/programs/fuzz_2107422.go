package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8; F1 uint32 }

func main() {
    var v1 int8 = -128
    v2 := []uint32{0, 2147483648, 954, 2147483648}
    var v3 bool = (false || (!false))
    var v4 float64 = 0.0
    v5 := make(map[int]string)
    if ((588 != 0) || v3) {
        v6 := GenS1{F0: 146}
        _ = v6
    }
    v1 -= 1
    var v7 uint8 = 0
    fmt.Println(v4, v1)
    for fi8 := 0; fi8 < 14; fi8++ {
        if (4294967295 < 2) {
            fmt.Println(fi8)
            switch (((("\n\t" + "a") + ("a" + "foo bar")) + "") + "foo bar") {
            case "":
                var v9 int8 = (v1 & (v1 >> ((18446744073709551615) & 63)))
                _ = v9
            case "foo bar":
                for fi10 := 0; fi10 < 17; fi10++ {
                    var v11 int16 = 0
                    var v12 float32 = 1e10
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            case "\n\t":
                var v13 int64 = 9223372036854775807
                _ = v13
            }
        }
        switch "foo bar" {
        case "hello":
            switch -1 {
            case -879:
                fmt.Println(v7, v1, v4)
            case 1:
                var v14 string = "a"
                _ = v14
            case 0:
                v2 = append(v2, 2147483647)
            default:
                for fi15 := 0; fi15 < 6; fi15++ {
                    fmt.Println(v3, fi15, fi15)
                    fmt.Println(v3, fi8, v4)
                    switch fi15 {
                    case -1:
                        var v16 uint = 42
                        _ = v16
                    case 1:
                        if v3 {
                            switch "hello" {
                            case "hello":
                                var v17 uint = 351
                                _ = v17
                            case "foo bar":
                                v1 = 0
                            }
                        }
                    case -937:
                        if (0.1 > (float32((3.14159 / float64(float32(-2.25)))) / 1.0)) {
                            fmt.Println(fi8, v4)
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            var v18 int64 = 9223372036854775807
                            _ = v18
                        }
                    default:
                        v19 := GenS1{F0: 42}
                        _ = v19
                    }
                    _ = fi15
                }
            }
        case "a":
            var v20 int32 = 2147483647
            _ = v20
        }
        _ = fi8
    }
    var v21 int32 = -1
    if (216 <= 775) {
        var v22 string = "a"
        if (!false) {
            v21 += v21
        }
        _ = v22
    }
    switch ("foo bar" + "a") {
    case "hello":
        var v23 float64 = v4
        _ = v23
    }
    v24 := GenS2{F0: 253}
    v25 := GenS1{F0: 42}
    var v26 int16 = -260
    if false {
        for fi27 := 0; fi27 < 12; fi27++ {
            v1 -= -53
            fmt.Println(v7)
            _ = fi27
        }
        v21 = -1
        v5[726] = ("a" + "hello")
    }
    if (-457 == -1) {
        var v28 float64 = 3.14159
        v21 = int32(v7)
        fmt.Println(v1, v3, v7)
        _ = v28
    } else {
        var v29 int64 = int64(v21)
        v30 := GenS1{F0: 118}
        _ = v29
        _ = v30
    }
    v4 += 0.0
    var v31 int16 = -836
    v32 := GenS2{F0: 2147483647}
    for fi33 := 0; fi33 < 7; fi33++ {
        var v34 bool = v3
        if v34 {
            v35 := GenS3{F0: v7, F1: 0}
            var v36 uint32 = 65535
            _ = v35
            _ = v36
        }
        _ = fi33
        _ = v34
    }
    for fi37 := 0; fi37 < 19; fi37++ {
        var v38 int64 = -114
        if ((true && (v3 || (false && (3.14159 > -2.25)))) || ((!v3) && true)) {
            fmt.Println(v38, v21)
        }
        _ = fi37
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v7
}
