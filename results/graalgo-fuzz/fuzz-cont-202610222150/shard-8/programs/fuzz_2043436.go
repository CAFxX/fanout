package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 uint64; F1 int16; F2 uint32 }
type GenS3 struct { F0 int8; F1 int8; F2 uint32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 bool }

func main() {
    v1 := []uint32{300, 275, 2, 773}
    var v2 uint = 267
    var v3 bool = true
    v2 += (v2 | (^(v2)))
    v2 = (^(v2))
    var v4 uint64 = 2
    v2 = v2
    if v3 {
        switch ("hello" + "") {
        case "\n\t":
            v4 = v4
        case "foo bar":
            var v5 uint32 = 1
            _ = v5
        case "hello":
            switch 32768 {
            case -790:
                v2 -= (v2 >> (((v2 + v2)) & 63))
            case -573:
                v2 = (v2 - v2)
            case 252:
                if (!((255 > 127) || v3)) {
                    switch "foo bar" {
                    case "a":
                        var v6 float32 = 0.0
                        _ = v6
                    case "foo bar":
                        var v7 uint8 = 1
                        _ = v7
                    case "":
                        var v8 string = "a"
                        _ = v8
                    default:
                        if (!v3) {
                            v2 = 4294967295
                            if (960 < -9223372036854775808) {
                                v9 := GenS1{F0: 0, F1: 1}
                                v10 := GenS2{F0: (v4 % ((42) | 1)), F1: 0, F2: 2147483648}
                                _ = v10
                                _ = v9
                            }
                        }
                    }
                    v11 := GenS3{F0: 1, F1: -47, F2: 678}
                    var v12 int16 = 651
                    _ = v11
                    _ = v12
                } else {
                    var v13 uint64 = 2
                    fmt.Println(v13, v3, v3)
                    fmt.Println(v4, v4, v2)
                    _ = v13
                }
            default:
                var v14 int64 = 378
                _ = v14
            }
        }
        for fi15 := 0; fi15 < 16; fi15++ {
            if v3 {
                v16 := GenS4{F0: 4294967295}
                _ = v16
            }
            if (v3 && (32767 > 218)) {
                fmt.Println(v4, v4, v2)
                var v17 uint16 = 0
                _ = v17
            }
            _ = len(v1)
            _ = fi15
        }
    }
    if ((("\n\t" + ("a" + "\n\t")) != "\n\t") || v3) {
        v2 = v2
    } else {
        v18 := GenS5{F0: v3}
        v19 := GenS5{F0: (157 == 30)}
        _ = v18
        _ = v19
    }
    if v3 {
        if true {
            var v20 bool = (255 != 244)
            if (88 != -128) {
                fmt.Println(v3, v2, v2)
                if (true && (v20 || (!true))) {
                    fmt.Println(v3)
                    var v21 float64 = 3.14159
                    _ = v21
                }
                var v22 int16 = 0
                _ = v22
            }
            var v23 int = -701
            _ = v20
            _ = v23
        }
    }
    v3 = v3
    var v24 int = 0
    fmt.Println(v2, v24, v24)
    var v25 string = "foo bar"
    for fi26 := 0; fi26 < 6; fi26++ {
        _ = len(v1)
        if (0.1 >= 0.1) {
            var v27 float32 = 1.5
            switch v25 {
            case "\n\t":
                var v28 int = 4294967295
                _ = v28
            case "hello":
                v3 = true
            }
            v29 := GenS1{F0: (fi26 / ((v24) | 1)), F1: fi26}
            _ = v27
            _ = v29
        }
        for fi30 := 0; fi30 < 14; fi30++ {
            var v31 int64 = 9223372036854775807
            var v32 int64 = (v31 / ((v31) | 1))
            _ = fi30
            _ = v31
            _ = v32
        }
        _ = fi26
    }
    switch (v24 - (v24 & v24)) {
    case 743:
        var v33 uint32 = 467
        _ = v33
    case -878:
        v1 = append(v1, 212)
    case -2147483648:
        var v34 int16 = -32768
        _ = v34
    default:
        for fi35 := 0; fi35 < 19; fi35++ {
            switch v25 {
            case "hello":
                var v36 uint32 = 521
                _ = v36
            }
            var v37 float64 = 1.5
            v3 = (255 > 66)
            _ = fi35
            _ = v37
        }
    }
    fmt.Println(v25)
    v1 = append(v1, 2147483648)
    var v38 float64 = (0.0 + float64(0.1))
    var v39 string = (v25 + v25)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v38
    _ = v39
    _ = v4
}
