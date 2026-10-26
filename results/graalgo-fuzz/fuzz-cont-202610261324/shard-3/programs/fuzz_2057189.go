package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 uint8 }
type GenS2 struct { F0 uint8; F1 int64; F2 int8 }
type GenS3 struct { F0 uint8; F1 uint16 }

func main() {
    var v1 uint = 832
    var v2 uint16 = 255
    var v3 uint32 = 65535
    v4 := []uint32{(v3 - 4294967295), uint32(v1), (v3 >> ((0) & 63))}
    var v5 uint64 = 161
    var v6 uint64 = v5
    switch "foo bar" {
    case "a":
        if (!true) {
            fmt.Println(v2, v2, v1)
        }
    default:
        var v7 int64 = 0
        _ = v7
    }
    var v8 int8 = -69
    v9 := GenS1{F0: v2, F1: (^(v2)), F2: 42}
    var v10 uint8 = 1
    for fi11 := 0; fi11 < 7; fi11++ {
        for fi12 := 0; fi12 < 5; fi12++ {
            v13 := GenS2{F0: 128, F1: int64(v1), F2: v8}
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    v3 = 882
    for fi14 := 0; fi14 < 18; fi14++ {
        var v15 uint16 = uint16(v10)
        for fi16 := 0; fi16 < 7; fi16++ {
            v10 += v10
            v17 := GenS2{F0: (v10 & (v10 / ((v10) | 1))), F1: -1, F2: (v8 & 66)}
            _ = fi16
            _ = v17
        }
        _ = fi14
        _ = v15
    }
    if false {
        v6 = v6
        if true {
            var v18 int64 = 42
            v3 -= v3
            if false {
                fmt.Println(v18, v6)
            } else {
                switch "" {
                case "":
                    for fi19 := 0; fi19 < 14; fi19++ {
                        fmt.Println(v18)
                        _ = fi19
                    }
                default:
                    if len(v4) > 0 {
                        v4[0] = (v3 | (v3 % (((^(v3))) | 1)))
                    }
                }
                var v20 int64 = (v18 << ((v1) & 63))
                _ = v20
            }
            _ = v18
        } else {
            v4 = append(v4, 2147483648)
            var v21 bool = (true && true)
            v3 = v3
            _ = v21
        }
        if (true && (("hello" + "hello") == "")) {
            _ = len(v4)
            if (!(v6 == (v6 % ((0) | 1)))) {
                fmt.Println(v2, v2, v2)
                var v22 uint8 = 127
                for fi23 := 0; fi23 < 7; fi23++ {
                    var v24 string = "\n\t"
                    _ = fi23
                    _ = v24
                }
                _ = v22
            }
            v1 += v1
        }
    }
    v5 = (v6 >> (((^(v1))) & 63))
    fmt.Println(v2)
    v25 := GenS3{F0: uint8(v1), F1: (v2 ^ (v2 - (v2 / (((v2 | v2)) | 1))))}
    switch -13 {
    case -795:
        if (-2.25 >= ((0.1 * 1e10) * (0.1 + float64(-2.25)))) {
            v26 := GenS3{F0: v10, F1: 36}
            v27 := GenS3{F0: 1, F1: 65535}
            _ = v26
            _ = v27
        } else {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            if ((uint8(v5) < v10) && false) {
                switch "\n\t" {
                case "\n\t":
                    var v28 float32 = 0.1
                    _ = v28
                case "foo bar":
                    var v29 bool = (181 >= 65535)
                    _ = v29
                }
                for fi30 := 0; fi30 < 8; fi30++ {
                    v31 := GenS3{F0: 2, F1: v2}
                    var v32 uint64 = 0
                    v33 := GenS2{F0: v10, F1: -1, F2: (v8 + -98)}
                    _ = fi30
                    _ = v31
                    _ = v32
                    _ = v33
                }
                for fi34 := 0; fi34 < 13; fi34++ {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    var v35 float64 = ((-2.25 / 1e10) - float64(3.14159))
                    _ = fi34
                    _ = v35
                }
            }
        }
    case -847:
        for fi36 := 0; fi36 < 12; fi36++ {
            for fi37 := 0; fi37 < 20; fi37++ {
                fmt.Println(v5, fi37)
                var v38 bool = true
                _ = fi37
                _ = v38
            }
            var v39 int64 = -9223372036854775808
            v3 += 115
            _ = fi36
            _ = v39
        }
    }
    fmt.Println(v2, v1, v2)
    var v40 string = "\n\t"
    v2 -= uint16(v10)
    v41 := GenS2{F0: v10, F1: 854, F2: v8}
    v1 += (v1 + 2)
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
