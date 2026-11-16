package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 bool }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int16; F1 uint8; F2 int8 }

func main() {
    var v1 bool = (-88 < -128)
    var v2 int8 = -1
    var v3 float64 = (3.14159 - 0.1)
    v4 := []uint32{2, 2, 2}
    var v5 uint64 = 388
    fmt.Println(v5)
    v1 = (true || v1)
    v3 -= 3.14159
    v6 := GenS1{F0: 127, F1: 1, F2: ((v2 / (((^(v2))) | 1)) < v2)}
    v5 = 18446744073709551615
    fmt.Println(v3)
    var v7 uint = 84
    if (1 != 483) {
        if ((v2 * v2) >= 1) {
            switch 42 {
            case 1:
                fmt.Println(v1, v1, v2)
            default:
                for fi8 := 0; fi8 < 5; fi8++ {
                    if len(v4) > 0 {
                        v4[0] = 1
                    }
                    for fi9 := 0; fi9 < 18; fi9++ {
                        v10 := GenS1{F0: 252, F1: 900, F2: (v1 && (false && (1 == 0)))}
                        _ = fi9
                        _ = v10
                    }
                    _ = fi8
                }
            }
            fmt.Println(v7)
        } else {
            for fi11 := 0; fi11 < 3; fi11++ {
                var v12 uint64 = v5
                _ = fi11
                _ = v12
            }
            var v13 uint32 = 2147483648
            var v14 int64 = 732
            _ = v13
            _ = v14
        }
        fmt.Println(v3, v3, v5)
        v4 = append(v4, 2)
    }
    v15 := GenS1{F0: 155, F1: 255, F2: false}
    switch "\n\t" {
    case "hello":
        var v16 int64 = -510
        _ = v16
    }
    if v1 {
        v2 = v2
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v17 := GenS2{F0: ("hello" + "hello")}
        _ = v17
    }
    if true {
        var v18 uint8 = 2
        if (!(76 >= v2)) {
            switch "hello" {
            case "foo bar":
                var v19 bool = v1
                _ = v19
            }
        } else {
            if true {
                v5 = (v5 * (v5 / (((v5 + (^(v5)))) | 1)))
                var v20 float32 = (1.5 * 3.14159)
                _ = v20
            } else {
                v21 := GenS2{F0: ""}
                v22 := GenS3{F0: 9223372036854775807}
                for fi23 := 0; fi23 < 17; fi23++ {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    switch ("hello" + "foo bar") {
                    case "hello":
                        var v24 uint32 = uint32(v18)
                        _ = v24
                    default:
                        if v1 {
                            if ((v18 << ((v7) & 63)) != uint8(v7)) {
                                var v25 uint32 = uint32(v18)
                                var v26 float32 = (float32(((((float64(1e10) * 1.5) / (v3 + (v3 / 1.5))) - ((v3 * float64((1e10 * 0.0))) * 1e10)) * -2.25)) * 0.1)
                                var v27 float32 = ((v26 - (v26 - (v26 / -2.25))) - float32((((v3 / (-2.25 / v3)) * v3) / (float64((v26 + 3.14159)) / float64(-2.25)))))
                                _ = v25
                                _ = v26
                                _ = v27
                            }
                            _ = len(v4)
                        }
                    }
                    _ = fi23
                }
                _ = v21
                _ = v22
            }
            if (!true) {
                var v28 uint64 = (v5 ^ 776)
                for fi29 := 0; fi29 < 11; fi29++ {
                    v30 := GenS2{F0: "a"}
                    _ = fi29
                    _ = v30
                }
                var v31 uint16 = 2
                _ = v28
                _ = v31
            }
            var v32 int = 32767
            _ = v32
        }
        v1 = (65535 < 652)
        _ = v18
    } else {
        if v1 {
            var v33 int8 = (v2 % ((-128) | 1))
            _ = v33
        } else {
            fmt.Println(v5, v3, v5)
            var v34 uint = 4294967295
            fmt.Println(v3, v7, v34)
            _ = v34
        }
    }
    v2 += v2
    v35 := GenS1{F0: 0, F1: 1, F2: (v1 || (1 >= 2147483647))}
    if (-32768 == 1) {
        for fi36 := 0; fi36 < 7; fi36++ {
            if len(v4) > 0 {
                v4[0] = 42
            }
            var v37 uint8 = 25
            var v38 float32 = (1.5 * 1.5)
            _ = fi36
            _ = v37
            _ = v38
        }
    } else {
        var v39 uint8 = 128
        _ = v39
    }
    v3 = v3
    v40 := GenS4{F0: 869, F1: 42, F2: (^(v2))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v7
}
