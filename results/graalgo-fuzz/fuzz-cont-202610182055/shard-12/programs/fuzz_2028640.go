package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 int }
type GenS2 struct { F0 uint64; F1 uint16; F2 uint8 }
func genFn0(p0 uint8) int {
    fmt.Println(p0)
    var v6 uint = uint(p0)
    var v7 uint16 = 388
    _ = p0
    _ = v6
    _ = v7
    return -9223372036854775808
}

func main() {
    var v1 int16 = 42
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]uint32)
    v4 := make(map[uint32]uint32)
    var v5 float64 = 0.1
    v8 := []int{32767, -86, -850, 809}
    for fi9 := 0; fi9 < 5; fi9++ {
        v10 := GenS1{F0: (v1 + (-(v1))), F1: 42, F2: (fi9 % ((fi9) | 1))}
        var v11 int16 = 1
        for fi12 := 0; fi12 < 11; fi12++ {
            var v13 float64 = v5
            if (-941 >= -19) {
                switch "" {
                case "\n\t":
                    fmt.Println(fi9, v13)
                case "foo bar":
                    if false {
                        if (0 >= 1) {
                            v11 = (-(v11))
                            fmt.Println(fi9)
                            v11 = (^(v11))
                        } else {
                            v5 += v5
                        }
                    }
                }
                if true {
                    var v14 int32 = -620
                    v15 := GenS1{F0: (v1 % (((v1 + (^(v1)))) | 1)), F1: 42, F2: genFn0(2)}
                    _ = v14
                    _ = v15
                }
            }
            v16 := GenS2{F0: 4294967295, F1: 65535, F2: 127}
            _ = fi12
            _ = v13
            _ = v16
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    fmt.Println(v1, v5)
    if true {
        v5 -= 1.5
        v5 = (3.14159 / 0.1)
    }
    fmt.Println(v1)
    for fi17 := 0; fi17 < 4; fi17++ {
        fmt.Println(fi17, fi17, v5)
        switch ("\n\t" + (("foo bar" + ("a" + (("a" + "foo bar") + "\n\t"))) + ("a" + ((("foo bar" + "hello") + ("" + "foo bar")) + "\n\t")))) {
        case "hello":
            _ = len(v8)
        case "foo bar":
            for fi18 := 0; fi18 < 6; fi18++ {
                var v19 float32 = (0.0 * 1e10)
                fmt.Println(v19)
                var v20 uint = 65535
                _ = fi18
                _ = v19
                _ = v20
            }
        case "":
            fmt.Println(fi17, v1)
        default:
            var v21 bool = false
            _ = v21
        }
        fmt.Println(v5)
        _ = fi17
    }
    v22 := GenS1{F0: (v1 - (v1 - v1)), F1: -9223372036854775808, F2: -1}
    v23 := GenS2{F0: 18446744073709551615, F1: 42, F2: 28}
    switch 1 {
    case 9223372036854775807:
        fmt.Println(v1, v1)
    case 899:
        if true {
            v1 += v1
            switch 0 {
            case 0:
                if (("" + "a") != "a") {
                    fmt.Println(v1, v1)
                    v24 := GenS1{F0: (v1 * v1), F1: -1, F2: genFn0(127)}
                    _ = v24
                } else {
                    var v25 uint32 = 176
                    _ = v25
                }
            default:
                if len(v8) > 0 {
                    v8[0] = 4294967295
                }
            }
            if (false || (!(false || ((-(v1)) > v1)))) {
                v5 -= (((v5 * (((v5 + v5) * -2.25) / v5)) / (v5 / v5)) * 0.0)
                switch "hello" {
                case "hello":
                    v1 -= v1
                case "a":
                    v26 := v4[83]
                    _ = v26
                case "\n\t":
                    fmt.Println(v5, v5)
                default:
                    var v27 string = "foo bar"
                    _ = v27
                }
            } else {
                for fi28 := 0; fi28 < 12; fi28++ {
                    var v29 float32 = 1e10
                    for fi30 := 0; fi30 < 9; fi30++ {
                        switch "hello" {
                        case "hello":
                            v8 = append(v8, fi28)
                        case "\n\t":
                            if ((!(0 <= 2)) || (346 != 1)) {
                                fmt.Println(v29, v5)
                                var v31 int16 = 0
                                if len(v8) >= 2 {
                                    v8 = v8[1:len(v8)]
                                }
                                _ = v31
                            }
                        case "a":
                            fmt.Println(fi28, v1, fi30)
                        }
                        switch fi28 {
                        case -1:
                            switch "a" {
                            case "foo bar":
                                delete(v2, 42)
                            }
                        default:
                            v32 := v3[4294967295]
                            _ = v32
                        }
                        switch "hello" {
                        case "":
                            fmt.Println(fi28, fi28)
                        case "\n\t":
                            if false {
                                var v33 float32 = (v29 - (0.0 - v29))
                                fmt.Println(v1, fi30, v1)
                                if (v33 >= (3.14159 - v29)) {
                                    if false {
                                        v1 -= (^(v1))
                                    }
                                    v29 += (0.1 + (v33 / v33))
                                    v1 += v1
                                }
                                var v34 bool = (!(-2147483648 > -793))
                                fmt.Println(fi28, v34, v1)
                                fmt.Println(v33, v34)
                                v5 += float64(float32(float64(v29)))
                                fmt.Println(v1)
                                var v35 uint64 = 42
                                fmt.Println(v34, v29, fi30)
                                _ = v33
                                _ = v34
                                _ = v35
                            }
                        case "a":
                            var v36 float32 = v29
                            _ = v36
                        default:
                            fmt.Println(v1, fi30)
                        }
                        _ = fi30
                    }
                    _ = fi28
                    _ = v29
                }
                var v37 int64 = 763
                switch ("hello" + "\n\t") {
                case "foo bar":
                    fmt.Println(v37, v5)
                }
                _ = v37
            }
        } else {
            if false {
                var v38 int64 = -9223372036854775808
                var v39 int16 = v1
                fmt.Println(v5)
                _ = v38
                _ = v39
            }
        }
    }
    v1 -= v1
    v1 = -1
    switch 1 {
    case 241:
        var v40 uint8 = 42
        _ = v40
    default:
        if true {
            var v41 string = "\n\t"
            for fi42 := 0; fi42 < 8; fi42++ {
                fmt.Println(v5, v5, v41)
                v1 += v1
                _ = fi42
            }
            _ = v41
        }
    }
    for fi43 := 0; fi43 < 14; fi43++ {
        var v44 int = (fi43 | genFn0(108))
        _ = fi43
        _ = v44
    }
    var v45 int = 0
    v1 -= (v1 ^ -324)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v8
}
