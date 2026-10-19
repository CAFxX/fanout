package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 int8; F1 int8 }
type GenS3 struct { F0 uint64; F1 uint64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint32; F1 int16; F2 int16 }
type GenS7 struct { F0 int }
func genFn0(p0 string, p1 string) int32 {
    var v6 uint64 = 135
    for fi7 := 0; fi7 < 13; fi7++ {
        for fi8 := 0; fi8 < 17; fi8++ {
            var v9 int64 = 9223372036854775807
            v9 = v9
            switch p0 {
            case "a":
                v10 := GenS1{F0: 18446744073709551615, F1: 65535}
                _ = v10
            case "foo bar":
                if (true && false) {
                    var v11 float64 = ((((1.5 - (0.0 - 1.5)) * ((1e10 / 0.1) / 0.1)) / (1e10 + (0.1 / 0.1))) - float64(-2.25))
                    _ = v11
                }
            }
            _ = fi8
            _ = v9
        }
        v12 := GenS2{F0: 50, F1: -126}
        _ = fi7
        _ = v12
    }
    var v13 bool = true
    _ = p0
    _ = p1
    _ = v13
    _ = v6
    return -2147483648
}
func genFn1(p0 int32) int64 {
    var v14 uint16 = 2
    v15 := GenS3{F0: 1, F1: 375}
    if (0 > 602) {
        switch "\n\t" {
        case "hello":
            var v16 int64 = -955
            _ = v16
        case "a":
            for fi17 := 0; fi17 < 13; fi17++ {
                for fi18 := 0; fi18 < 11; fi18++ {
                    v14 = v14
                    v19 := GenS2{F0: 123, F1: 42}
                    _ = fi18
                    _ = v19
                }
                fmt.Println(p0, v14)
                _ = fi17
            }
        }
        var v20 float64 = -2.25
        for fi21 := 0; fi21 < 13; fi21++ {
            for fi22 := 0; fi22 < 7; fi22++ {
                fmt.Println(v14)
                _ = fi22
            }
            _ = fi21
        }
        _ = v20
    }
    var v23 int64 = 0
    fmt.Println(v23)
    _ = p0
    _ = v14
    _ = v15
    _ = v23
    return (v23 ^ 535)
}

func main() {
    var v1 int16 = -1
    v2 := []uint32{0, 2147483647, 828, 65535}
    var v3 bool = (true && (v1 == (v1 / ((-921) | 1))))
    var v4 int16 = -1
    var v5 uint = 1
    switch 32768 {
    case 201:
        v1 = v4
    default:
        fmt.Println(v3, v3)
    }
    switch int(v5) {
    case 877:
        v5 -= (v5 - (^(v5)))
    case 127:
        var v24 uint8 = 8
        _ = v24
    }
    v1 -= (v1 & v4)
    var v25 float32 = 3.14159
    v25 = 1e10
    switch 9223372036854775807 {
    case 32767:
        var v26 int64 = -700
        _ = v26
    case -281:
        v5 -= (v5 % (((v5 % ((42) | 1))) | 1))
    default:
        if (false && v3) {
            for fi27 := 0; fi27 < 6; fi27++ {
                switch -155 {
                case 128:
                    v1 = 1
                case -2147483648:
                    v28 := GenS4{F0: fi27}
                    _ = v28
                default:
                    if v3 {
                        fmt.Println(v3, fi27, fi27)
                    }
                }
                v29 := GenS4{F0: (fi27 ^ (fi27 % ((fi27) | 1)))}
                for fi30 := 0; fi30 < 10; fi30++ {
                    fmt.Println(v4, v4, v5)
                    _ = fi30
                }
                _ = fi27
                _ = v29
            }
            var v31 uint = (v5 * v5)
            _ = v31
        }
    }
    v25 += v25
    v1 = v1
    if (v3 || v3) {
        var v32 uint = 2
        switch "" {
        case "foo bar":
            var v33 int = 2147483647
            _ = v33
        default:
            var v34 uint16 = uint16(v5)
            _ = v34
        }
        _ = v32
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v35 uint8 = 12
        v36 := GenS5{F0: 128}
        _ = v35
        _ = v36
    }
    _ = len(v2)
    var v37 bool = ((4294967295 >= 2) || (-15 == 0))
    _ = len(v2)
    if false {
        v38 := GenS6{F0: 0, F1: v4, F2: v1}
        _ = v38
    } else {
        for fi39 := 0; fi39 < 4; fi39++ {
            var v40 string = "foo bar"
            _ = fi39
            _ = v40
        }
        for fi41 := 0; fi41 < 12; fi41++ {
            v42 := GenS7{F0: fi41}
            var v43 int64 = -9223372036854775808
            fmt.Println(v5, v4, v1)
            _ = fi41
            _ = v42
            _ = v43
        }
        v2 = append(v2, 286)
    }
    for fi44 := 0; fi44 < 13; fi44++ {
        v2 = append(v2, 1)
        _ = fi44
    }
    v4 -= (v4 >> ((v5) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v37
    _ = v4
    _ = v5
}
