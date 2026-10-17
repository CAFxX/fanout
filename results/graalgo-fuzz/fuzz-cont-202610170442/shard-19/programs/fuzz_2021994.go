package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int16; F1 bool; F2 uint16 }
type GenS5 struct { F0 uint; F1 int8; F2 uint16 }
type GenS6 struct { F0 bool }

func main() {
    v1 := make(map[string]string)
    var v2 int32 = 0
    var v3 uint8 = 255
    v4 := make(map[string]bool)
    var v5 uint8 = (^(v3))
    v6 := GenS1{F0: -16}
    switch "foo bar" {
    case "":
        v5 -= v5
    case "foo bar":
        fmt.Println(v3)
    case "\n\t":
        v7 := []int{-177, -9223372036854775808, -2147483648}
        _ = v7
    }
    var v8 uint32 = 886
    if ((42 >= 32767) && (552 >= -32768)) {
        for fi9 := 0; fi9 < 15; fi9++ {
            v10 := GenS1{F0: 1}
            for fi11 := 0; fi11 < 7; fi11++ {
                switch "a" {
                case "a":
                    switch 127 {
                    case -2147483648:
                        fmt.Println(v3, v2, fi9)
                    case 2147483647:
                        switch "hello" {
                        case "":
                            var v12 float64 = (float64(3.14159) * (1e10 / (float64(1.5) + 0.0)))
                            _ = v12
                        case "foo bar":
                            v13 := GenS2{F0: 127}
                            _ = v13
                        default:
                            v2 = (v2 | (v2 >> ((1) & 63)))
                        }
                    default:
                        switch "hello" {
                        case "foo bar":
                            v14 := GenS1{F0: 1}
                            _ = v14
                        case "a":
                            v15 := []int{-523, int(v2), fi11, 128}
                            v16 := GenS3{F0: fi11}
                            _ = v15
                            _ = v16
                        default:
                            switch ("a" + "a") {
                            case "foo bar":
                                fmt.Println(fi9, v8)
                                delete(v1, "foo bar")
                                switch "" {
                                case "":
                                    v17 := GenS4{F0: 32767, F1: false, F2: 32767}
                                    _ = v17
                                case "foo bar":
                                    var v18 float64 = 1e10
                                    _ = v18
                                default:
                                    fmt.Println(v5, v5, v8)
                                }
                            default:
                                var v19 uint8 = 14
                                _ = v19
                            }
                        }
                    }
                case "\n\t":
                    _ = len(v1)
                }
                var v20 int32 = -2147483648
                v5 = (^(v3))
                _ = fi11
                _ = v20
            }
            fmt.Println(v5, v5, v5)
            _ = fi9
            _ = v10
        }
        v21 := GenS1{F0: -9223372036854775808}
        _ = v21
    } else {
        var v22 string = "a"
        v5 += v5
        v3 += v5
        _ = v22
    }
    v23 := []int{4294967295}
    var v24 uint = 18446744073709551615
    v23 = append(v23, 227)
    var v25 int16 = 215
    for fi26 := 0; fi26 < 13; fi26++ {
        for fi27 := 0; fi27 < 18; fi27++ {
            fmt.Println(v8, fi26)
            if false {
                v3 = (v5 << (((^(v24))) & 63))
                var v28 int = (^(fi27))
                var v29 uint64 = 1
                v30 := GenS4{F0: (v25 / ((v25) | 1)), F1: (((1e10 - -2.25) + (3.14159 / 3.14159)) == 0.1), F2: 691}
                _ = v28
                _ = v29
                _ = v30
            } else {
                v5 = (^(v3))
                var v31 float32 = -2.25
                _ = v31
            }
            _ = fi27
        }
        _ = fi26
    }
    switch 346 {
    case 128:
        if ((v8 / (((v8 | 4294967295)) | 1)) < v8) {
            v25 += (v25 >> (((v24 % (((^(v24))) | 1))) & 63))
        } else {
            v5 -= v5
        }
    }
    var v32 int = -9223372036854775808
    var v33 string = "a"
    switch v32 {
    case 665:
        if (false && ((!((1e10 * (float32(float64(1.5)) + -2.25)) != (0.1 * float32(float64(3.14159))))) && true)) {
            v34 := GenS4{F0: v25, F1: false, F2: uint16(v8)}
            v32 = (v32 % ((v32) | 1))
            var v35 int64 = -1
            _ = v34
            _ = v35
        }
    case 391:
        _ = len(v4)
    }
    var v36 float32 = (0.0 - 0.0)
    fmt.Println(v32, v5, v5)
    if true {
        _ = len(v4)
        v37 := GenS1{F0: -1}
        _ = v37
    }
    if (float32((0.0 / 3.14159)) != v36) {
        for fi38 := 0; fi38 < 5; fi38++ {
            for fi39 := 0; fi39 < 15; fi39++ {
                if (!true) {
                    v3 = 42
                    v36 = v36
                    v36 += 0.1
                }
                switch v33 {
                case "":
                    if len(v23) > 0 {
                        v23[0] = v32
                    }
                case "hello":
                    var v40 int8 = 0
                    _ = v40
                }
                _ = fi39
            }
            _ = fi38
        }
        if (!(48 <= 127)) {
            v24 = (v24 << ((65535) & 63))
            v41 := GenS5{F0: 2, F1: 127, F2: 200}
            _ = v41
        }
    } else {
        for fi42 := 0; fi42 < 20; fi42++ {
            for fi43 := 0; fi43 < 2; fi43++ {
                v4[v33] = false
                v44 := GenS6{F0: ((true && ((515 < 18446744073709551615) && (1e10 <= ((0.0 + (((1.5 / -2.25) - 1.5) * 1.5)) + 3.14159)))) || true)}
                _ = fi43
                _ = v44
            }
            _ = fi42
        }
        if len(v23) > 0 {
            v23[0] = (v32 | v32)
        }
        v8 -= v8
    }
    var v45 int8 = 100
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v32
    _ = v33
    _ = v36
    _ = v4
    _ = v45
    _ = v5
    _ = v6
    _ = v8
}
