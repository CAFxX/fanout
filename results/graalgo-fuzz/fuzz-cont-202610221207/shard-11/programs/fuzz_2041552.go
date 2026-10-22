package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 uint16; F1 int32; F2 int }
type GenS4 struct { F0 uint64; F1 string; F2 int64 }
type GenS5 struct { F0 uint64; F1 int8; F2 int8 }
func genFn0(p0 int) int32 {
    v5 := GenS1{F0: (p0 ^ p0)}
    var v6 int64 = 1
    v6 = 42
    _ = p0
    _ = v5
    _ = v6
    return -12
}

func main() {
    var v1 uint32 = 42
    var v2 int8 = 0
    var v3 int16 = 0
    v4 := []int{127}
    v2 -= 0
    var v7 uint16 = 255
    for fi8 := 0; fi8 < 4; fi8++ {
        v9 := GenS1{F0: (fi8 << ((1) & 63))}
        v10 := GenS2{F0: (fi8 >> ((65535) & 63)), F1: int16(v2)}
        _ = fi8
        _ = v10
        _ = v9
    }
    v4 = append(v4, -343)
    switch "foo bar" {
    case "hello":
        if ((v3 * 567) > int16(v1)) {
            var v11 uint8 = 0
            for fi12 := 0; fi12 < 9; fi12++ {
                switch "foo bar" {
                case "foo bar":
                    v1 = 2147483647
                case "hello":
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                case "":
                    for fi13 := 0; fi13 < 14; fi13++ {
                        v14 := GenS3{F0: (v7 >> ((2) & 63)), F1: -2147483648, F2: (fi13 ^ fi13)}
                        switch (-(fi13)) {
                        case 1:
                            switch int(v11) {
                            case 908:
                                if false {
                                    v15 := GenS1{F0: fi13}
                                    _ = v15
                                }
                            case 148:
                                fmt.Println(v11)
                            default:
                                var v16 int16 = (v3 | 42)
                                _ = v16
                            }
                        default:
                            if (!(!false)) {
                                var v17 string = "a"
                                v7 = (^(v7))
                                _ = v17
                            }
                        }
                        _ = fi13
                        _ = v14
                    }
                default:
                    if false {
                        v3 = 42
                        for fi18 := 0; fi18 < 9; fi18++ {
                            var v19 uint32 = 2
                            v11 -= 204
                            v2 = v2
                            _ = fi18
                            _ = v19
                        }
                    } else {
                        if false {
                            switch "hello" {
                            case "a":
                                v7 -= (v7 % (((v7 - v7)) | 1))
                            }
                            v7 = (v7 - v7)
                            var v20 uint16 = (v7 * v7)
                            _ = v20
                        } else {
                            var v21 int64 = 1
                            v4 = append(v4, (fi12 % (((-(fi12))) | 1)))
                            if (!true) {
                                v3 = v3
                                v22 := GenS2{F0: (fi12 << ((1) & 63)), F1: (v3 + v3)}
                                v2 += 1
                                _ = v22
                            } else {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                v23 := GenS4{F0: 18446744073709551615, F1: ("foo bar" + "\n\t"), F2: (v21 * -9223372036854775808)}
                                var v24 int16 = v3
                                _ = v23
                                _ = v24
                            }
                            _ = v21
                        }
                        for fi25 := 0; fi25 < 8; fi25++ {
                            var v26 uint8 = (v11 / (((v11 ^ (v11 / ((v11) | 1)))) | 1))
                            _ = fi25
                            _ = v26
                        }
                    }
                }
                var v27 int16 = v3
                v1 = (v1 % (((v1 % ((v1) | 1))) | 1))
                _ = fi12
                _ = v27
            }
            for fi28 := 0; fi28 < 4; fi28++ {
                for fi29 := 0; fi29 < 9; fi29++ {
                    v7 += v7
                    _ = fi29
                }
                switch "" {
                case "\n\t":
                    v2 = 100
                }
                _ = fi28
            }
            _ = v11
        } else {
            if (((v2 % (((v2 & (v2 + v2))) | 1)) <= (v2 & v2)) || (42 > 18446744073709551615)) {
                var v30 int64 = 169
                fmt.Println(v3)
                _ = v30
            }
            var v31 uint16 = (v7 - v7)
            _ = v31
        }
    }
    if len(v4) > 0 {
        v4[0] = -9223372036854775808
    }
    var v32 uint = 0
    v1 -= v1
    v32 -= (v32 - (^(v32)))
    for fi33 := 0; fi33 < 14; fi33++ {
        for fi34 := 0; fi34 < 10; fi34++ {
            v35 := GenS5{F0: 368, F1: 0, F2: v2}
            v1 = 0
            var v36 float64 = 1.5
            if true {
                var v37 int16 = 1
                v4 = append(v4, fi33)
                if (("hello" + ("a" + "hello")) != "a") {
                    var v38 float32 = (((0.0 + 0.0) / 1.5) / 1e10)
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    var v39 int8 = v2
                    v39 = -125
                    fmt.Println(v37)
                    _ = v38
                    _ = v39
                }
                _ = v37
            } else {
                var v40 uint64 = 0
                var v41 uint8 = 128
                _ = v40
                _ = v41
            }
            if (2147483647 == 0) {
                v3 += v3
                v36 = v36
                if true {
                    fmt.Println(v3, v7, fi33)
                    fmt.Println(v2, v3, fi34)
                }
            } else {
                v32 = (v32 | v32)
                var v42 uint8 = 216
                if ((!false) || (v2 >= v2)) {
                    v42 -= (^(v42))
                } else {
                    fmt.Println(fi33, v32)
                    var v43 string = "foo bar"
                    v3 = (v3 + v3)
                    _ = v43
                }
                _ = len(v4)
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v7 = uint16(v32)
                fmt.Println(v32, v36, v3)
                _ = v42
            }
            _ = fi34
            _ = v35
            _ = v36
        }
        v1 = v1
        _ = fi33
    }
    var v44 uint32 = (v1 | 65535)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v44:", v44)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v44
    _ = v7
}
