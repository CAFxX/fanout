package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 uint8; F1 string; F2 uint8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 string; F1 uint; F2 uint16 }
func genFn0(p0 bool, p1 int16) uint16 {
    v6 := []int{32767}
    fmt.Println(p0, p1)
    var v7 int8 = -6
    v6 = append(v6, -503)
    fmt.Println(p0, v7)
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 1
}

func main() {
    var v1 int32 = 307
    var v2 uint16 = 255
    v3 := make(map[uint32]uint32)
    var v4 string = "a"
    var v5 float32 = 1.5
    var v8 uint = 4294967295
    v9 := v3[873]
    for fi10 := 0; fi10 < 12; fi10++ {
        if ((v1 <= (^(v1))) || false) {
            var v11 uint = (v8 & (v8 << (((v8 ^ (v8 / ((v8) | 1)))) & 63)))
            _ = v11
        }
        _ = fi10
    }
    v12 := []int{127, 9223372036854775807}
    v13 := GenS1{F0: 42, F1: (-1 == -9223372036854775808)}
    v14 := GenS1{F0: -9223372036854775808, F1: true}
    var v15 uint32 = v9
    if (true || true) {
        v3[2147483648] = (v15 + 1)
        var v16 bool = (false && true)
        _ = v16
    } else {
        v5 -= (1e10 + -2.25)
        fmt.Println(v2)
    }
    v1 = (v1 / ((int32(v9)) | 1))
    if false {
        switch 0 {
        case 4294967295:
            for fi17 := 0; fi17 < 14; fi17++ {
                v8 = v8
                _ = fi17
            }
        case 32768:
            if (-100 >= -128) {
                var v18 uint64 = 331
                _ = v18
            } else {
                v5 = (v5 / ((((((v5 / (v5 / (v5 - 1e10))) / 3.14159) - v5) * (v5 / v5)) / (v5 + ((v5 / (((1e10 / v5) / 1.5) * ((v5 + (1e10 + (v5 - ((v5 - 3.14159) + v5)))) + 1e10))) * v5))) + v5))
            }
        default:
            for fi19 := 0; fi19 < 18; fi19++ {
                for fi20 := 0; fi20 < 2; fi20++ {
                    switch (((v4 + "a") + "hello") + "foo bar") {
                    case "hello":
                        v21 := GenS2{F0: 161, F1: "a", F2: 255}
                        _ = v21
                    }
                    fmt.Println(v2, fi19, v1)
                    _ = fi20
                }
                if false {
                    if (!false) {
                        var v22 uint = v8
                        var v23 string = "hello"
                        _ = v22
                        _ = v23
                    } else {
                        v15 = (v9 * v9)
                        if (1 <= 9223372036854775807) {
                            var v24 float32 = v5
                            _ = v24
                        } else {
                            if (!((0.0 * -2.25) == (3.14159 * (1e10 + (0.1 / -2.25))))) {
                                var v25 uint64 = uint64(v15)
                                _ = v25
                            }
                        }
                        v8 -= v8
                    }
                } else {
                    for fi26 := 0; fi26 < 14; fi26++ {
                        v1 = -230
                        delete(v3, v9)
                        _ = fi26
                    }
                }
                v9 += (v15 * (v9 ^ (v15 % ((v15) | 1))))
                _ = fi19
            }
        }
        var v27 uint8 = 42
        _ = v27
    }
    if (!false) {
        if (uint32(v8) == 65535) {
            var v28 float64 = (0.0 - (((1e10 * 0.0) - 3.14159) / 1.0))
            _ = v28
        }
        for fi29 := 0; fi29 < 19; fi29++ {
            v30 := GenS3{F0: fi29}
            switch 42 {
            case 32767:
                var v31 uint8 = 1
                _ = v31
            case 128:
                var v32 int8 = 0
                _ = v32
            case -323:
                v33 := GenS3{F0: fi29}
                _ = v33
            }
            var v34 int32 = (v1 | (v1 | v1))
            _ = fi29
            _ = v30
            _ = v34
        }
        switch ("" + (((v4 + (v4 + "hello")) + (v4 + (("" + v4) + ((v4 + v4) + v4)))) + v4)) {
        case "foo bar":
            var v35 string = v4
            _ = v35
        default:
            v15 = (v9 + v15)
        }
    }
    if true {
        var v36 int64 = 6
        if true {
            var v37 uint64 = 601
            _ = v37
        } else {
            v8 = (^(v8))
            for fi38 := 0; fi38 < 17; fi38++ {
                for fi39 := 0; fi39 < 19; fi39++ {
                    v1 = -858
                    if (false && ((float32(0.1) > 0.0) || (!(1e10 == 3.14159)))) {
                        v40 := v3[v9]
                        v36 = (v36 & (v36 / (((v36 << ((v8) & 63))) | 1)))
                        var v41 int16 = 0
                        _ = v40
                        _ = v41
                    } else {
                        v15 = v15
                        var v42 string = v4
                        v43 := GenS4{F0: v4, F1: 42, F2: (v2 | (^(v2)))}
                        if (0 == -32768) {
                            if len(v12) >= 2 {
                                v12 = v12[1:len(v12)]
                            }
                            var v44 string = v4
                            switch fi38 {
                            case 279:
                                var v45 uint = v8
                                _ = v45
                            default:
                                var v46 uint8 = 127
                                _ = v46
                            }
                            _ = v44
                        }
                        _ = v42
                        _ = v43
                    }
                    _ = fi39
                }
                _ = len(v12)
                var v47 int = (fi38 - (fi38 << (((v8 + 65535)) & 63)))
                _ = fi38
                _ = v47
            }
            switch v4 {
            case "foo bar":
                var v48 string = v4
                _ = v48
            default:
                var v49 uint64 = 18446744073709551615
                _ = v49
            }
        }
        if (!(!("\n\t" == v4))) {
            _ = len(v3)
            var v50 int8 = 42
            _ = v50
        }
        _ = v36
    }
    var v51 uint32 = uint32(v8)
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    var v52 int = 663
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v52
    _ = v8
    _ = v9
}
