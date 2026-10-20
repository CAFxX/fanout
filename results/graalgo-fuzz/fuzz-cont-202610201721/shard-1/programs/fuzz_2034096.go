package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint32; F1 string }
type GenS6 struct { F0 int16 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint32 = 4294967295
    var v3 int32 = -725
    var v4 int = -439
    v5 := []string{"a", "foo bar"}
    v6 := GenS1{F0: -1}
    v5 = append(v5, ("foo bar" + "hello"))
    v7 := GenS1{F0: 1}
    v8 := GenS2{F0: false}
    if false {
        for fi9 := 0; fi9 < 18; fi9++ {
            for fi10 := 0; fi10 < 6; fi10++ {
                v4 = fi9
                v1[v2] = (!(true && ((1.5 + 0.1) <= 1.5)))
                var v11 uint = 18446744073709551615
                v11 -= 1
                if (true && (true || ((v2 + v2) < (v2 ^ (v2 | v2))))) {
                    v1[(v2 - v2)] = false
                    var v12 bool = ((42 >= v3) && true)
                    _ = v12
                } else {
                    if (uint16(v11) > 903) {
                        v3 -= v3
                        v3 = v3
                        switch "" {
                        case "foo bar":
                            var v13 uint32 = v2
                            _ = v13
                        case "hello":
                            var v14 uint = (v11 | 720)
                            _ = v14
                        }
                        v11 -= v11
                        switch (fi10 | (v4 & v4)) {
                        case 32767:
                            if (0.0 >= 0.0) {
                                if false {
                                    switch "foo bar" {
                                    case "\n\t":
                                        v15 := v1[(v2 % ((v2) | 1))]
                                        _ = v15
                                    }
                                    _ = len(v1)
                                } else {
                                    var v16 bool = (0 > 127)
                                    _ = v16
                                }
                                var v17 uint32 = (v2 ^ v2)
                                _ = v17
                            }
                        case 267:
                            v4 = 127
                        }
                        switch (v4 % ((-751) | 1)) {
                        case -603:
                            v18 := GenS3{F0: "\n\t"}
                            _ = v18
                        case 2147483647:
                            var v19 int16 = -1
                            _ = v19
                        default:
                            if ((v3 * (v3 - (v3 & (v3 * (v3 & (v3 - (v3 * (v3 % (((v3 - 0)) | 1))))))))) < v3) {
                                if true {
                                    v4 += -438
                                    var v20 uint64 = 42
                                    _ = v20
                                }
                                v21 := GenS4{F0: 5}
                                switch "\n\t" {
                                case "":
                                    if len(v5) > 0 {
                                        v5[0] = "\n\t"
                                    }
                                case "\n\t":
                                    switch v4 {
                                    case -2147483648:
                                        if len(v5) > 0 {
                                            v5[0] = ""
                                        }
                                    default:
                                        fmt.Println(v11, v3)
                                    }
                                case "a":
                                    v22 := GenS1{F0: 1}
                                    _ = v22
                                }
                                if false {
                                    var v23 float32 = 1e10
                                    v2 = 647
                                    _ = len(v5)
                                    _ = v23
                                }
                                _ = v21
                            }
                        }
                        if len(v5) > 0 {
                            v5[0] = "\n\t"
                        }
                        v1[(v2 + v2)] = (2147483647 != v3)
                        v11 = (^(v11))
                    } else {
                        fmt.Println(v2, v2)
                        if true {
                            _ = len(v1)
                            var v24 int = (fi9 ^ -700)
                            _ = len(v5)
                            _ = v24
                        }
                        fmt.Println(v4, v11, v4)
                    }
                    _ = len(v5)
                    switch (fi9 + int(v2)) {
                    case -489:
                        v3 += 1
                    case 4294967295:
                        v25 := GenS5{F0: v2, F1: "\n\t"}
                        _ = v25
                    }
                }
                _ = fi10
                _ = v11
            }
            _ = fi9
        }
        v26 := GenS1{F0: 61}
        _ = v26
    }
    switch -64 {
    case 127:
        switch (("hello" + "\n\t") + "a") {
        case "a":
            for fi27 := 0; fi27 < 4; fi27++ {
                fmt.Println(fi27, v4, v2)
                var v28 int = v4
                var v29 int8 = -1
                _ = fi27
                _ = v28
                _ = v29
            }
        case "foo bar":
            v30 := GenS6{F0: -32768}
            _ = v30
        }
    case 42:
        delete(v1, (^(v2)))
    }
    for fi31 := 0; fi31 < 11; fi31++ {
        var v32 float32 = -2.25
        for fi33 := 0; fi33 < 13; fi33++ {
            switch (-(fi31)) {
            case -2147483648:
                v5 = append(v5, "hello")
            case 710:
                switch -1 {
                case -86:
                    v4 += v4
                case 0:
                    switch "a" {
                    case "foo bar":
                        var v34 int = 9223372036854775807
                        _ = v34
                    case "":
                        _ = len(v1)
                    }
                case 1:
                    var v35 uint8 = 2
                    _ = v35
                }
            case 965:
                var v36 float32 = (float32((-2.25 * 1.5)) / v32)
                _ = v36
            }
            if (!(-128 >= -79)) {
                v32 -= ((v32 - v32) + v32)
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi33
        }
        for fi37 := 0; fi37 < 13; fi37++ {
            if len(v5) > 0 {
                v5[0] = ("foo bar" + "a")
            }
            if true {
                v38 := v1[v2]
                delete(v1, v2)
                _ = v38
            }
            _ = fi37
        }
        _ = fi31
        _ = v32
    }
    v39 := GenS4{F0: 1}
    if (float32((((float64((0.1 + -2.25)) + 1e10) - 1.5) / ((0.1 + ((0.0 + 1e10) / 0.1)) - 0.0))) < (float32((0.1 + 0.1)) + (0.0 / (3.14159 / 3.14159)))) {
        if ((v2 | v2) > (v2 ^ v2)) {
            var v40 uint8 = 42
            _ = v40
        }
    } else {
        v4 += -702
    }
    for fi41 := 0; fi41 < 15; fi41++ {
        if (float32(1e10) > ((0.1 * 0.0) + 0.1)) {
            var v42 uint16 = 2
            var v43 int64 = 428
            _ = v42
            _ = v43
        } else {
            v4 = (v4 - v4)
            fmt.Println(fi41)
            for fi44 := 0; fi44 < 12; fi44++ {
                var v45 uint8 = 128
                var v46 uint8 = v45
                _ = fi44
                _ = v45
                _ = v46
            }
        }
        _ = fi41
    }
    v4 -= v4
    for fi47 := 0; fi47 < 19; fi47++ {
        fmt.Println(v3)
        for fi48 := 0; fi48 < 17; fi48++ {
            var v49 string = "a"
            if ((((0.1 * 3.14159) >= -2.25) || false) && ((^(v2)) < 2147483647)) {
                var v50 int16 = 1000
                if true {
                    var v51 float64 = 1e10
                    _ = v51
                } else {
                    v2 += (v2 << ((65535) & 63))
                    var v52 uint8 = 1
                    _ = v52
                }
                _ = v50
            } else {
                v2 = 65535
            }
            var v53 int16 = int16(v3)
            _ = fi48
            _ = v49
            _ = v53
        }
        for fi54 := 0; fi54 < 5; fi54++ {
            var v55 int8 = 1
            var v56 int = fi47
            v3 += v3
            fmt.Println(fi54, v3, v56)
            _ = len(v5)
            _ = fi54
            _ = v55
            _ = v56
        }
        _ = fi47
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
