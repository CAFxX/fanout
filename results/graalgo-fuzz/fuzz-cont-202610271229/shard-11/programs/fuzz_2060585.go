package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8; F1 uint16; F2 bool }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int32 }
func genFn0() uint16 {
    var v7 int8 = 127
    v7 -= 100
    v7 -= (v7 ^ v7)
    _ = v7
    return 42
}
func genFn1(p0 int, p1 uint, p2 uint) int32 {
    v8 := []int{p0, (p0 ^ -1)}
    for fi9 := 0; fi9 < 14; fi9++ {
        switch "hello" {
        case "":
            switch "hello" {
            case "a":
                var v10 uint16 = 1
                _ = v10
            case "foo bar":
                if (!true) {
                    var v11 uint8 = 42
                    for fi12 := 0; fi12 < 1; fi12++ {
                        fmt.Println(fi9)
                        _ = fi12
                    }
                    fmt.Println(p0, fi9, p1)
                    _ = v11
                }
            default:
                for fi13 := 0; fi13 < 9; fi13++ {
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                    p0 -= fi13
                    _ = fi13
                }
            }
        case "a":
            if false {
                _ = len(v8)
                if true {
                    p1 = p2
                }
                var v14 bool = false
                _ = v14
            }
        case "foo bar":
            for fi15 := 0; fi15 < 16; fi15++ {
                if len(v8) > 0 {
                    v8[0] = p0
                }
                p0 -= 213
                p1 -= (p1 % ((p1) | 1))
                _ = fi15
            }
        default:
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
        }
        switch (fi9 / ((p0) | 1)) {
        case 0:
            var v16 uint = 739
            _ = v16
        case 299:
            p2 = p2
        case 32768:
            for fi17 := 0; fi17 < 11; fi17++ {
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                if len(v8) > 0 {
                    v8[0] = (p0 + (-(fi17)))
                }
                fmt.Println(p2, fi9, p1)
                _ = fi17
            }
        }
        p2 = p1
        _ = fi9
    }
    v18 := GenS1{F0: 18446744073709551615, F1: p1, F2: p0}
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v8
    return -2147483648
}

func main() {
    var v1 float32 = 0.0
    var v2 string = "a"
    var v3 uint64 = 495
    var v4 uint = 244
    var v5 uint8 = 127
    var v6 int32 = 501
    v19 := []int{}
    var v20 bool = true
    var v21 string = (v2 + ("a" + ""))
    if v20 {
        v2 = ("\n\t" + v21)
        if len(v19) > 0 {
            v19[0] = 9223372036854775807
        }
    } else {
        _ = len(v19)
        v19 = append(v19, -9223372036854775808)
    }
    if (v20 && v20) {
        for fi22 := 0; fi22 < 18; fi22++ {
            var v23 int = 1
            switch fi22 {
            case 32767:
                switch "foo bar" {
                case "hello":
                    if (true && (v20 && (v20 || true))) {
                        _ = len(v19)
                        for fi24 := 0; fi24 < 13; fi24++ {
                            if v20 {
                                v25 := GenS1{F0: (^(v4)), F1: v4, F2: v23}
                                _ = v25
                            } else {
                                if (42 < -128) {
                                    v21 = v2
                                    fmt.Println(v23, fi24)
                                } else {
                                    if false {
                                        if ((v20 && v20) || ((fi22 << ((v4) & 63)) <= (fi22 | (fi24 & (fi24 | (v23 & 127)))))) {
                                            var v26 uint = 581
                                            switch fi22 {
                                            case -2147483648:
                                                if len(v19) >= 2 {
                                                    v19 = v19[1:len(v19)]
                                                }
                                            case -1:
                                                var v27 int = v23
                                                _ = v27
                                            case 1:
                                                var v28 string = "hello"
                                                _ = v28
                                            }
                                            _ = len(v19)
                                            _ = v26
                                        }
                                        v29 := GenS2{F0: 181}
                                        fmt.Println(v4, v21)
                                        _ = v29
                                    }
                                    v21 = "\n\t"
                                    fmt.Println(v3, v23, fi22)
                                }
                            }
                            fmt.Println(fi24)
                            v6 -= (-(v6))
                            _ = fi24
                        }
                    } else {
                        v30 := GenS3{F0: -112, F1: 42, F2: (float64(((3.14159 - v1) + 3.14159)) == 1e10)}
                        _ = v30
                    }
                case "\n\t":
                    var v31 float64 = 1.5
                    _ = v31
                }
            default:
                v32 := GenS4{F0: 690}
                _ = v32
            }
            _ = fi22
            _ = v23
        }
        for fi33 := 0; fi33 < 12; fi33++ {
            v34 := GenS2{F0: (v5 >> ((uint(v3)) & 63))}
            switch fi33 {
            case 2147483647:
                _ = len(v19)
            case -364:
                _ = len(v19)
            default:
                var v35 bool = (v20 || (v5 >= (v5 - v5)))
                _ = v35
            }
            _ = fi33
            _ = v34
        }
    } else {
        v4 = (v4 * (v4 + v4))
    }
    v36 := GenS5{F0: 1}
    for fi37 := 0; fi37 < 15; fi37++ {
        switch "" {
        case "hello":
            if (v20 && v20) {
                var v38 float32 = v1
                if (!(42 < 1000)) {
                    switch (fi37 >> ((v4) & 63)) {
                    case 1:
                        var v39 bool = v20
                        _ = v39
                    case 2147483647:
                        v3 = (^(v3))
                    }
                    var v40 uint16 = uint16(v4)
                    _ = v40
                }
                _ = v38
            }
        case "foo bar":
            var v41 int16 = 42
            _ = v41
        default:
            for fi42 := 0; fi42 < 4; fi42++ {
                v43 := GenS3{F0: 42, F1: 197, F2: (!(338 > 42))}
                _ = fi42
                _ = v43
            }
        }
        switch v2 {
        case "foo bar":
            var v44 uint16 = 42
            _ = v44
        default:
            v6 -= v6
        }
        _ = fi37
    }
    v3 -= (v3 - (v3 * v3))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
