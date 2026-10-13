package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 uint16; F1 uint32 }
type GenS4 struct { F0 bool; F1 uint8; F2 int }

func main() {
    var v1 float64 = 1e10
    var v2 string = "a"
    var v3 uint = 42
    var v4 float64 = v1
    v5 := []uint32{4294967295, 1, 718}
    v6 := GenS1{F0: 187}
    fmt.Println(v3, v3)
    var v7 float64 = v1
    var v8 uint16 = 1
    fmt.Println(v2)
    if true {
        v4 -= (v1 - v7)
        for fi9 := 0; fi9 < 15; fi9++ {
            switch (fi9 / (((fi9 >> (((v3 ^ 255)) & 63))) | 1)) {
            case 42:
                v10 := GenS1{F0: 204}
                _ = v10
            case 9223372036854775807:
                if (v2 <= ((v2 + v2) + v2)) {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    var v11 uint8 = 2
                    if len(v5) > 0 {
                        v5[0] = 42
                    }
                    _ = v11
                }
            case 128:
                var v12 float64 = (v4 * (((v4 / (((v7 + 0.1) * 0.1) / (((v4 + (float64(1e10) + float64(((1.5 + 1e10) * -2.25)))) + v7) * ((v1 / ((float64(0.1) + 1.5) * v1)) + ((1e10 / v7) * v4))))) - 0.0) + v7))
                _ = v12
            default:
                if false {
                    var v13 uint16 = v8
                    _ = v13
                }
            }
            _ = fi9
        }
        if true {
            _ = len(v5)
        }
    } else {
        v1 = v4
        if true {
            for fi14 := 0; fi14 < 4; fi14++ {
                fmt.Println(v4, v2, v1)
                _ = fi14
            }
            for fi15 := 0; fi15 < 15; fi15++ {
                for fi16 := 0; fi16 < 7; fi16++ {
                    v17 := GenS1{F0: 33}
                    _ = fi16
                    _ = v17
                }
                fmt.Println(fi15, v2)
                _ = fi15
            }
        } else {
            if ((false && false) || (v8 == v8)) {
                switch v2 {
                case "a":
                    if (42 != -1) {
                        switch v2 {
                        case "a":
                            var v18 int8 = 0
                            _ = v18
                        default:
                            v19 := GenS2{F0: 2147483647, F1: false}
                            _ = v19
                        }
                        v8 = (^(v8))
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    } else {
                        fmt.Println(v8, v2, v8)
                    }
                case "hello":
                    var v20 uint64 = 1
                    _ = v20
                case "foo bar":
                    fmt.Println(v4)
                default:
                    var v21 float32 = 3.14159
                    _ = v21
                }
            }
            for fi22 := 0; fi22 < 2; fi22++ {
                v4 = float64(1e10)
                for fi23 := 0; fi23 < 1; fi23++ {
                    v24 := GenS3{F0: 42, F1: 2147483648}
                    _ = fi23
                    _ = v24
                }
                v2 = v2
                _ = fi22
            }
        }
        fmt.Println(v4, v2, v7)
    }
    if ((((1.5 - 0.1) / v1) + v7) <= 0.0) {
        fmt.Println(v4)
    }
    var v25 int16 = -37
    var v26 float64 = v7
    fmt.Println(v2, v8, v7)
    if ((v25 >> ((v3) & 63)) <= 1) {
        v8 -= (v8 >> ((v3) & 63))
    } else {
        var v27 int = -1
        switch (v2 + "hello") {
        case "hello":
            v5 = append(v5, 2147483647)
        }
        for fi28 := 0; fi28 < 2; fi28++ {
            v29 := GenS4{F0: false, F1: 255, F2: v27}
            _ = fi28
            _ = v29
        }
        _ = v27
    }
    fmt.Println(v7)
    var v30 uint = 4294967295
    if (!true) {
        v1 += ((float64((1e10 * 3.14159)) - v4) + (v1 / (0.1 + float64(1.5))))
    }
    switch ("\n\t" + (v2 + (v2 + v2))) {
    case "hello":
        if (v8 > (v8 >> (((v30 + (v30 % ((255) | 1)))) & 63))) {
            v8 = 0
            v25 = 251
        }
    case "foo bar":
        var v31 int32 = -685
        _ = v31
    default:
        v32 := GenS4{F0: (false || (true || true)), F1: 2, F2: 0}
        _ = v32
    }
    fmt.Println(v1, v8, v25)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
