package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int8; F1 string; F2 uint32 }
type GenS4 struct { F0 int8; F1 string }
func genFn0(p0 uint8) int8 {
    p0 += p0
    p0 = (p0 / ((p0) | 1))
    for fi7 := 0; fi7 < 15; fi7++ {
        v8 := GenS1{F0: fi7}
        v9 := []int{(fi7 % (((fi7 / ((fi7) | 1))) | 1)), -2147483648}
        v9 = append(v9, 829)
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = p0
    return 42
}

func main() {
    v1 := []uint32{2147483647, 2, 2, 548}
    v2 := make(map[string]string)
    var v3 float32 = -2.25
    var v4 uint64 = 1
    v5 := make(map[int]uint32)
    var v6 int = 1
    if true {
        fmt.Println(v3, v3, v4)
        for fi10 := 0; fi10 < 15; fi10++ {
            delete(v2, "hello")
            v4 -= 397
            _ = fi10
        }
        var v11 int8 = -128
        _ = v11
    } else {
        var v12 int = v6
        v12 += v12
        _ = v12
    }
    for fi13 := 0; fi13 < 8; fi13++ {
        v3 = v3
        _ = fi13
    }
    v4 -= (v4 >> ((65535) & 63))
    _ = len(v5)
    v5[v6] = 4294967295
    var v14 int64 = -9223372036854775808
    if (704 != 653) {
        v1 = append(v1, 32)
        if (201 > 0) {
            v1 = append(v1, 1)
            v15 := GenS2{F0: 162}
            var v16 uint = 4294967295
            _ = v15
            _ = v16
        } else {
            for fi17 := 0; fi17 < 9; fi17++ {
                var v18 int64 = v14
                _ = fi17
                _ = v18
            }
        }
    }
    switch (((("" + "\n\t") + "") + "hello") + "a") {
    case "":
        fmt.Println(v3)
    }
    if false {
        for fi19 := 0; fi19 < 18; fi19++ {
            for fi20 := 0; fi20 < 7; fi20++ {
                v4 += v4
                if false {
                    v6 = (v6 << ((2) & 63))
                } else {
                    v1 = append(v1, 42)
                    switch (-(v6)) {
                    case 128:
                        var v21 int8 = 31
                        if false {
                            v6 += v6
                            _ = len(v1)
                            v14 = 48
                        } else {
                            v22 := GenS1{F0: -114}
                            var v23 int64 = (-(v14))
                            if false {
                                switch ("hello" + "a") {
                                case "":
                                    switch "\n\t" {
                                    case "foo bar":
                                        fmt.Println(v6, v6, v14)
                                    case "":
                                        var v24 string = "\n\t"
                                        _ = v24
                                    default:
                                        var v25 uint32 = 128
                                        _ = v25
                                    }
                                case "\n\t":
                                    v21 -= -128
                                case "hello":
                                    v4 -= v4
                                default:
                                    var v26 float32 = (v3 * 0.0)
                                    _ = v26
                                }
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            } else {
                                var v27 int8 = 42
                                _ = v27
                            }
                            switch "" {
                            case "":
                                switch (("hello" + (("\n\t" + ("a" + "hello")) + (("" + "a") + ""))) + "foo bar") {
                                case "hello":
                                    var v28 int8 = (v21 + (v21 * v21))
                                    _ = v28
                                default:
                                    var v29 int8 = v21
                                    _ = v29
                                }
                            default:
                                v3 -= v3
                            }
                            _ = v22
                            _ = v23
                        }
                        fmt.Println(v21)
                        _ = v21
                    case 0:
                        var v30 string = "\n\t"
                        _ = v30
                    }
                }
                _ = fi20
            }
            v3 += (v3 - float32((3.14159 / 1.5)))
            _ = fi19
        }
        v3 -= 3.14159
        if false {
            for fi31 := 0; fi31 < 15; fi31++ {
                var v32 int16 = 1
                fmt.Println(v4, fi31)
                var v33 uint64 = (v4 ^ 18446744073709551615)
                _ = fi31
                _ = v32
                _ = v33
            }
            for fi34 := 0; fi34 < 11; fi34++ {
                if ((!(v14 < int64(v4))) && (-32768 > -32768)) {
                    switch v6 {
                    case 9223372036854775807:
                        _ = len(v1)
                    default:
                        v14 = v14
                    }
                    v35 := GenS3{F0: 108, F1: ("a" + "hello"), F2: 629}
                    _ = v35
                } else {
                    var v36 int64 = v14
                    var v37 uint64 = (v4 >> ((517) & 63))
                    _ = v36
                    _ = v37
                }
                _ = fi34
            }
        } else {
            var v38 uint = 2
            _ = v38
        }
    }
    fmt.Println(v3, v4, v4)
    if (float64(v3) > float64(v3)) {
        var v39 string = "\n\t"
        var v40 float64 = -2.25
        var v41 uint32 = 4294967295
        _ = v39
        _ = v40
        _ = v41
    } else {
        v3 = v3
        var v42 int16 = 177
        _ = v42
    }
    var v43 uint = 42
    var v44 uint16 = 42
    v45 := GenS4{F0: -40, F1: "hello"}
    for fi46 := 0; fi46 < 6; fi46++ {
        v47 := GenS2{F0: 2147483647}
        var v48 uint64 = 747
        _ = fi46
        _ = v47
        _ = v48
    }
    fmt.Println(v4, v14, v6)
    if ((v14 * (v14 | -601)) == v14) {
        for fi49 := 0; fi49 < 2; fi49++ {
            v1 = append(v1, 42)
            var v50 int64 = (v14 + v14)
            v14 = (v50 & 205)
            _ = fi49
            _ = v50
        }
        var v51 uint32 = 2147483648
        _ = len(v1)
        _ = v51
    } else {
        _ = len(v1)
    }
    var v52 int32 = 2147483647
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v52
    _ = v6
}
