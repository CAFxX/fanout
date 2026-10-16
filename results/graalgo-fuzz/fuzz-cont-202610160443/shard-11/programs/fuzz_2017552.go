package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 uint8; F2 uint8 }
type GenS3 struct { F0 uint16; F1 uint16; F2 uint8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int64; F1 int16 }
func genFn0(p0 uint32, p1 uint64, p2 string) uint64 {
    fmt.Println(p0, p2, p2)
    v7 := []int{9223372036854775807, -9223372036854775808, 9223372036854775807, 9223372036854775807}
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 float64 = -2.25
        fmt.Println(fi8, p2, v9)
        var v10 uint64 = p1
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 int64 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v7
    return (p1 & (p1 * (p1 << ((1) & 63))))
}

func main() {
    var v1 uint64 = 4294967295
    v2 := make(map[int]string)
    var v3 int8 = 0
    v4 := make(map[string]string)
    var v5 uint16 = 42
    v6 := []int{128, -9223372036854775808}
    v12 := v2[-9223372036854775808]
    v4[(v12 + "hello")] = "a"
    var v13 uint = 1
    v14 := GenS1{F0: v12}
    fmt.Println(v3, v1, v1)
    for fi15 := 0; fi15 < 11; fi15++ {
        switch 128 {
        case 128:
            if ((v5 >= (v5 ^ v5)) && false) {
                v16 := GenS2{F0: ((v3 % ((-128) | 1)) < v3), F1: 0, F2: 1}
                for fi17 := 0; fi17 < 17; fi17++ {
                    if (!(true || ((v1 ^ v1) <= (v1 + (v1 % ((uint64(v13)) | 1)))))) {
                        v18 := GenS3{F0: (v5 - 65535), F1: uint16(v1), F2: 2}
                        _ = v18
                    }
                    if true {
                        v4[v12] = v12
                        _ = len(v4)
                        fmt.Println(v5, v13, fi15)
                        var v19 int = 32767
                        var v20 uint8 = 10
                        _ = v19
                        _ = v20
                    } else {
                        v21 := GenS3{F0: 255, F1: (^(v5)), F2: uint8(v13)}
                        _ = v21
                    }
                    _ = fi17
                }
                var v22 uint = v13
                _ = v16
                _ = v22
            } else {
                var v23 int64 = 1
                var v24 int32 = -136
                for fi25 := 0; fi25 < 5; fi25++ {
                    v23 -= 1
                    switch "foo bar" {
                    case "":
                        fmt.Println(v23, v5)
                    default:
                        switch "hello" {
                        case "\n\t":
                            switch v12 {
                            case "hello":
                                if (-32 != 42) {
                                    fmt.Println(v1, fi25)
                                    var v26 uint8 = 127
                                    v23 -= (v23 % ((-998) | 1))
                                    var v27 float64 = -2.25
                                    var v28 uint8 = v26
                                    _ = v26
                                    _ = v27
                                    _ = v28
                                }
                            case "a":
                                var v29 int = (fi25 | (fi15 % ((fi15) | 1)))
                                _ = v29
                            case "foo bar":
                                var v30 int64 = (v23 ^ (^(v23)))
                                _ = v30
                            }
                        }
                    }
                    _ = fi25
                }
                _ = v23
                _ = v24
            }
        }
        for fi31 := 0; fi31 < 16; fi31++ {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            v32 := GenS3{F0: 255, F1: 956, F2: 221}
            _ = fi31
            _ = v32
        }
        _ = fi15
    }
    var v33 uint = uint(v1)
    var v34 uint32 = 0
    var v35 int32 = 1
    if false {
        if (false || true) {
            fmt.Println(v3, v13, v33)
            switch v12 {
            case "hello":
                v36 := GenS2{F0: false, F1: 127, F2: 2}
                _ = v36
            default:
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            }
        }
        v13 += v13
    } else {
        var v37 uint64 = (v1 >> (((v13 + 42)) & 63))
        v35 = (v35 / ((489) | 1))
        v38 := GenS4{F0: true}
        _ = v37
        _ = v38
    }
    v3 = (v3 - v3)
    if ((-2.25 + ((3.14159 - (1.5 + 0.0)) * (1.5 * (3.14159 / 1.0)))) == 0.0) {
        v39 := GenS5{F0: 0, F1: -1}
        switch v12 {
        case "\n\t":
            var v40 uint32 = v34
            _ = v40
        default:
            for fi41 := 0; fi41 < 2; fi41++ {
                var v42 uint64 = v1
                _ = fi41
                _ = v42
            }
        }
        if (!(!(52 < 127))) {
            if (580 > -773) {
                switch 776 {
                case 600:
                    var v43 string = "\n\t"
                    _ = v43
                case -1:
                    _ = len(v6)
                case 127:
                    var v44 uint = v13
                    _ = v44
                }
                var v45 uint64 = 42
                _ = v45
            }
            v46 := GenS3{F0: (v5 | (^(v5))), F1: (v5 * (v5 ^ v5)), F2: 165}
            v47 := GenS3{F0: (v5 + uint16(v1)), F1: (v5 + 0), F2: 2}
            _ = v46
            _ = v47
        }
        _ = v39
    } else {
        v34 = 42
        delete(v4, v12)
    }
    v34 += (v34 << ((v33) & 63))
    if len(v6) > 0 {
        v6[0] = 128
    }
    var v48 uint = 42
    var v49 bool = (false && false)
    fmt.Println(v3)
    switch v12 {
    case "\n\t":
        var v50 uint = (v13 << ((18446744073709551615) & 63))
        _ = v50
    case "foo bar":
        v12 = ""
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v6
}
