package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 int8; F1 int8 }
type GenS4 struct { F0 uint32 }
func genFn0(p0 int64, p1 uint64, p2 uint64) int32 {
    v4 := []int{-9223372036854775808}
    switch ("foo bar" + "hello") {
    case "foo bar":
        if true {
            if ((255 <= 563) && false) {
                if false {
                    for fi5 := 0; fi5 < 13; fi5++ {
                        p0 -= 42
                        if (!(("" + (("" + ("foo bar" + "foo bar")) + (("hello" + "hello") + "foo bar"))) < "foo bar")) {
                            var v6 int32 = -1
                            _ = v6
                        }
                        v7 := GenS1{F0: (fi5 * 661), F1: 0}
                        _ = fi5
                        _ = v7
                    }
                }
                _ = len(v4)
            }
            var v8 uint = 537
            if ((("hello" != ("\n\t" + "a")) && (false && true)) && ((((-2.25 + 3.14159) * ((0.0 - 1e10) * (0.1 - 0.1))) == -2.25) && (true && (v8 != v8)))) {
                v8 = (v8 >> ((v8) & 63))
                switch 158 {
                case 47:
                    for fi9 := 0; fi9 < 18; fi9++ {
                        v10 := GenS1{F0: (fi9 - (fi9 | fi9)), F1: -74}
                        v11 := GenS2{F0: p0, F1: -1}
                        switch "foo bar" {
                        case "hello":
                            var v12 float32 = ((0.0 + 3.14159) * (1.5 - 1.5))
                            _ = v12
                        case "foo bar":
                            if len(v4) > 0 {
                                v4[0] = fi9
                            }
                        case "a":
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                        _ = fi9
                        _ = v10
                        _ = v11
                    }
                default:
                    p0 -= (p0 >> ((v8) & 63))
                }
                fmt.Println(p1, p0)
            }
            _ = v8
        } else {
            var v13 float64 = 1.5
            v14 := GenS1{F0: -2147483648, F1: -1}
            v15 := GenS3{F0: -85, F1: 127}
            _ = v13
            _ = v14
            _ = v15
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 42
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 float64 = 0.1
    var v3 int64 = 0
    fmt.Println(v2, v3)
    var v16 uint32 = 299
    v17 := v1[849]
    var v18 uint32 = (v16 / (((v16 << ((18446744073709551615) & 63))) | 1))
    var v19 float64 = (float64((1e10 * 1.5)) * v2)
    fmt.Println(v17, v16, v16)
    if (!v17) {
        if v17 {
            fmt.Println(v19)
            for fi20 := 0; fi20 < 13; fi20++ {
                switch "a" {
                case "hello":
                    for fi21 := 0; fi21 < 11; fi21++ {
                        v22 := v1[(v18 - (v18 | 4294967295))]
                        _ = fi21
                        _ = v22
                    }
                case "\n\t":
                    if v17 {
                        var v23 uint = 2
                        fmt.Println(fi20, v3)
                        fmt.Println(v16)
                        _ = v23
                    } else {
                        v24 := GenS4{F0: 672}
                        v25 := []int{fi20}
                        _ = v24
                        _ = v25
                    }
                case "a":
                    v18 += (v18 >> ((4294967295) & 63))
                default:
                    if true {
                        v16 += v18
                    } else {
                        v26 := v1[1]
                        for fi27 := 0; fi27 < 7; fi27++ {
                            v16 -= v18
                            var v28 int16 = -32768
                            _ = fi27
                            _ = v28
                        }
                        _ = v26
                    }
                }
                for fi29 := 0; fi29 < 12; fi29++ {
                    v18 += v16
                    if false {
                        var v30 uint16 = 32767
                        if false {
                            var v31 uint = uint(v16)
                            _ = v31
                        }
                        fmt.Println(v3, fi20, v16)
                        _ = v30
                    }
                    _ = fi29
                }
                _ = fi20
            }
            fmt.Println(v2)
        } else {
            switch ("foo bar" + "a") {
            case "\n\t":
                v32 := GenS3{F0: 100, F1: -128}
                _ = v32
            }
        }
        var v33 float32 = 1e10
        for fi34 := 0; fi34 < 14; fi34++ {
            fmt.Println(v18, v3)
            fmt.Println(v18, v33, v16)
            v33 = float32(v2)
            _ = fi34
        }
        _ = v33
    } else {
        if (false || (496 > 0)) {
            var v35 uint8 = 93
            var v36 uint = 2
            var v37 uint16 = 1
            _ = v35
            _ = v36
            _ = v37
        } else {
            _ = len(v1)
        }
    }
    v3 = v3
    var v38 uint32 = (v18 >> ((4294967295) & 63))
    v1[(v38 + v38)] = false
    for fi39 := 0; fi39 < 7; fi39++ {
        var v40 bool = v17
        _ = fi39
        _ = v40
    }
    v2 = (v2 * float64(0.1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v38
}
