package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8; F1 int8; F2 uint32 }
type GenS4 struct { F0 int16; F1 int32 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint; F1 int64 }
type GenS7 struct { F0 uint8 }
func genFn0(p0 int16) uint32 {
    var v7 uint64 = 1
    v7 -= (v7 << ((2) & 63))
    v8 := []int{4294967295, -1, 653, -2147483648}
    v7 = v7
    _ = p0
    _ = v7
    _ = v8
    return 2147483648
}

func main() {
    var v1 float32 = (0.0 / 1.0)
    v2 := []uint32{2147483648, 65535}
    v3 := make(map[int]bool)
    v4 := make(map[uint32]bool)
    var v5 int = 540
    v6 := make(map[int]bool)
    _ = len(v2)
    v9 := GenS1{F0: 0, F1: (1 == -9223372036854775808), F2: 2}
    fmt.Println(v5)
    v5 = 32768
    fmt.Println(v5, v1)
    _ = len(v2)
    v1 = (v1 * ((float32(-2.25) / v1) + 1.5))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v2 = append(v2, 2147483648)
    if false {
        fmt.Println(v1, v5)
        if true {
            v10 := GenS2{F0: 2}
            delete(v6, -763)
            _ = v10
        } else {
            var v11 int32 = 42
            _ = v11
        }
        v5 = 4294967295
    }
    if (-753 < int64(v5)) {
        v5 = v5
        for fi12 := 0; fi12 < 18; fi12++ {
            fmt.Println(v5, v1)
            if true {
                v5 -= 9223372036854775807
                v13 := GenS3{F0: -1, F1: 0, F2: 376}
                _ = v13
            } else {
                _ = len(v2)
            }
            for fi14 := 0; fi14 < 1; fi14++ {
                fmt.Println(v1)
                var v15 float32 = v1
                _ = fi14
                _ = v15
            }
            _ = fi12
        }
    } else {
        if ((0.1 > (1.5 - 1e10)) && ((0 >= 433) && (false && (0 != 0)))) {
            var v16 int = v5
            switch v5 {
            case 42:
                if len(v2) > 0 {
                    v2[0] = 4294967295
                }
            default:
                v1 += v1
            }
            _ = v16
        }
        for fi17 := 0; fi17 < 2; fi17++ {
            _ = len(v2)
            _ = len(v6)
            v18 := v4[916]
            _ = fi17
            _ = v18
        }
        for fi19 := 0; fi19 < 1; fi19++ {
            fmt.Println(v5)
            if len(v2) > 0 {
                v2[0] = 4294967295
            }
            _ = fi19
        }
    }
    if true {
        for fi20 := 0; fi20 < 2; fi20++ {
            if ("foo bar" >= ("\n\t" + ("hello" + ((("" + "foo bar") + "hello") + "hello")))) {
                var v21 uint8 = 127
                _ = v21
            }
            for fi22 := 0; fi22 < 14; fi22++ {
                var v23 float64 = 0.0
                switch "hello" {
                case "foo bar":
                    v23 -= v23
                    switch ("a" + ("\n\t" + "foo bar")) {
                    case "hello":
                        var v24 int64 = 9223372036854775807
                        _ = v24
                    case "a":
                        var v25 uint64 = 502
                        _ = v25
                    case "foo bar":
                        if ((fi20 << ((65535) & 63)) >= fi22) {
                            if false {
                                var v26 int64 = int64(v5)
                                _ = v26
                            }
                            if (!(!("foo bar" > "foo bar"))) {
                                fmt.Println(fi20, v1, v5)
                            } else {
                                v27 := GenS4{F0: 0, F1: 932}
                                _ = v27
                            }
                            switch "\n\t" {
                            case "":
                                var v28 int = fi20
                                _ = v28
                            case "hello":
                                delete(v6, 128)
                            default:
                                v3[fi22] = false
                                v5 = fi22
                                v23 += v23
                            }
                            fmt.Println(fi22)
                        }
                    }
                case "":
                    v29 := v3[(fi22 - (fi20 / ((-691) | 1)))]
                    _ = v29
                default:
                    v30 := GenS5{F0: -376}
                    var v31 int64 = -884
                    _ = v30
                    _ = v31
                }
                switch "foo bar" {
                case "a":
                    if true {
                        v32 := GenS6{F0: 315, F1: -9223372036854775808}
                        _ = v32
                    } else {
                        v23 += (((v23 / 1.0) * -2.25) - ((float64(v1) * float64(v1)) / (3.14159 - v23)))
                        v5 = fi20
                    }
                    v23 -= v23
                default:
                    if (2 > 42) {
                        var v33 uint8 = 2
                        _ = v33
                    }
                }
                _ = fi22
                _ = v23
            }
            _ = fi20
        }
        var v34 bool = false
        switch v5 {
        case 967:
            fmt.Println(v5, v34)
        }
        _ = v34
    }
    var v35 uint32 = 2147483648
    delete(v4, (v35 ^ 1))
    fmt.Println(v35, v1)
    switch v5 {
    case 32768:
        if true {
            v5 = v5
            for fi36 := 0; fi36 < 12; fi36++ {
                v37 := GenS7{F0: 0}
                var v38 uint = 921
                var v39 uint = v38
                _ = fi36
                _ = v37
                _ = v38
                _ = v39
            }
        } else {
            for fi40 := 0; fi40 < 13; fi40++ {
                if true {
                    var v41 uint16 = 32767
                    var v42 int8 = -40
                    var v43 uint8 = 223
                    _ = v41
                    _ = v42
                    _ = v43
                } else {
                    v2 = append(v2, (v35 - (^(v35))))
                    var v44 uint64 = 18446744073709551615
                    _ = v44
                }
                _ = fi40
            }
            v35 = v35
            v1 += float32(1e10)
        }
    default:
        fmt.Println(v1)
    }
    var v45 int16 = 1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v45
    _ = v5
    _ = v6
    _ = v9
}
