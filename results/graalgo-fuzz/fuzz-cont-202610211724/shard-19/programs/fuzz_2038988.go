package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint; F1 uint8; F2 uint64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint16; F1 uint8 }
type GenS5 struct { F0 int32; F1 int16; F2 string }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint8 }
func genFn0(p0 uint64) uint32 {
    for fi7 := 0; fi7 < 19; fi7++ {
        if (fi7 != (fi7 / ((fi7) | 1))) {
            p0 += 42
            v8 := GenS1{F0: -1}
            _ = v8
        }
        _ = fi7
    }
    v9 := GenS1{F0: -32768}
    v10 := GenS1{F0: -32768}
    for fi11 := 0; fi11 < 1; fi11++ {
        p0 = p0
        _ = fi11
    }
    p0 -= p0
    _ = p0
    _ = v10
    _ = v9
    return 709
}

func main() {
    var v1 int32 = 2147483647
    var v2 bool = true
    v3 := []int{127}
    v4 := []uint32{}
    var v5 int8 = -45
    var v6 uint = 1
    v12 := GenS1{F0: 984}
    v5 = v5
    fmt.Println(v5, v6, v2)
    v13 := GenS2{F0: (v6 >> (((v6 << ((v6) & 63))) & 63)), F1: 2, F2: 42}
    var v14 int64 = -9223372036854775808
    for fi15 := 0; fi15 < 5; fi15++ {
        if v2 {
            v5 = v5
            v16 := GenS2{F0: (^(v6)), F1: 61, F2: 942}
            _ = v16
        }
        _ = fi15
    }
    switch "foo bar" {
    case "hello":
        for fi17 := 0; fi17 < 1; fi17++ {
            var v18 float64 = (float64((1.5 - 1e10)) - float64(3.14159))
            if false {
                var v19 int64 = 610
                for fi20 := 0; fi20 < 18; fi20++ {
                    switch "hello" {
                    case "hello":
                        if v2 {
                            v21 := GenS1{F0: 42}
                            var v22 float64 = v18
                            _ = v21
                            _ = v22
                        }
                    }
                    v1 = v1
                    _ = len(v4)
                    _ = fi20
                }
                if v2 {
                    var v23 float32 = (((float32(v18) * 1e10) + 3.14159) / float32((3.14159 + (v18 + 0.1))))
                    switch "\n\t" {
                    case "hello":
                        v2 = (v2 && false)
                    case "a":
                        v5 -= int8(v1)
                    }
                    _ = v23
                }
                _ = v19
            } else {
                _ = len(v4)
                fmt.Println(v2, v14)
            }
            switch "a" {
            case "":
                if (4294967295 <= 2) {
                    for fi24 := 0; fi24 < 6; fi24++ {
                        switch (-(fi24)) {
                        case 2147483647:
                            fmt.Println(v14, v18)
                        case 0:
                            if (!false) {
                                switch "foo bar" {
                                case "hello":
                                    fmt.Println(v5)
                                case "":
                                    v25 := GenS1{F0: -1}
                                    if len(v3) > 0 {
                                        v3[0] = 0
                                    }
                                    _ = v25
                                case "a":
                                    if (true && v2) {
                                        v26 := GenS3{F0: 2147483647}
                                        var v27 float32 = 0.0
                                        switch "a" {
                                        case "":
                                            if (v18 != 0.1) {
                                                v28 := GenS4{F0: 65535, F1: 64}
                                                _ = v28
                                            }
                                        default:
                                            v29 := GenS3{F0: genFn0(316)}
                                            _ = v29
                                        }
                                        if v2 {
                                            fmt.Println(v1, v27, v1)
                                            v27 = (float32(0.0) / 1.0)
                                            v30 := GenS4{F0: 919, F1: 127}
                                            _ = v30
                                        }
                                        _ = v26
                                        _ = v27
                                    } else {
                                        v31 := GenS4{F0: 1, F1: 2}
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        _ = v31
                                    }
                                }
                                var v32 uint64 = 1
                                _ = v32
                            } else {
                                if len(v4) > 0 {
                                    v4[0] = 864
                                }
                            }
                        default:
                            var v33 uint32 = 907
                            _ = v33
                        }
                        switch fi17 {
                        case 127:
                            v2 = (!(float32(3.14159) < -2.25))
                        default:
                            v34 := GenS5{F0: 692, F1: 1, F2: "\n\t"}
                            _ = v34
                        }
                        v4 = append(v4, 2147483647)
                        if false {
                            switch (fi24 >> ((18446744073709551615) & 63)) {
                            case 128:
                                v5 += (^(v5))
                                _ = len(v3)
                                fmt.Println(v2, v5, fi24)
                            }
                            _ = len(v4)
                            var v35 float64 = 0.1
                            _ = v35
                        }
                        v36 := GenS6{F0: genFn0(1)}
                        _ = fi24
                        _ = v36
                    }
                } else {
                    fmt.Println(v14, v1, v5)
                    var v37 uint = v6
                    v1 -= v1
                    _ = v37
                }
            case "foo bar":
                var v38 uint32 = 760
                _ = v38
            }
            _ = fi17
            _ = v18
        }
    case "a":
        var v39 float32 = ((1e10 - 0.0) - float32((-2.25 - (1.5 + (1e10 / 0.1)))))
        _ = v39
    default:
        for fi40 := 0; fi40 < 6; fi40++ {
            v1 += (v1 << (((v6 >> (((v6 | (v6 << (((v6 & (v6 | 255))) & 63)))) & 63))) & 63))
            _ = fi40
        }
    }
    v5 = v5
    v1 += (v1 / ((-2147483648) | 1))
    for fi41 := 0; fi41 < 8; fi41++ {
        _ = len(v4)
        _ = fi41
    }
    var v42 float32 = (((1.5 / ((-2.25 - -2.25) - (3.14159 * 1.5))) / 1.5) / 3.14159)
    if v2 {
        v43 := GenS1{F0: 1000}
        _ = v43
    } else {
        v1 += v1
    }
    v44 := GenS7{F0: (v6 ^ v6)}
    v45 := GenS1{F0: 0}
    v46 := GenS8{F0: 137}
    var v47 int = -416
    for fi48 := 0; fi48 < 8; fi48++ {
        fmt.Println(v6, v47, v14)
        v6 += v6
        for fi49 := 0; fi49 < 2; fi49++ {
            _ = len(v3)
            if v2 {
                var v50 uint = v6
                v50 = (v50 * 651)
                if len(v4) > 0 {
                    v4[0] = 65535
                }
                _ = v50
            }
            _ = fi49
        }
        _ = fi48
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v6
}
