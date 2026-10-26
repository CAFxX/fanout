package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 uint; F1 uint; F2 uint32 }
type GenS3 struct { F0 string; F1 int64; F2 string }
type GenS4 struct { F0 uint64; F1 uint }
type GenS5 struct { F0 int8 }
func genFn0() int16 {
    var v5 int = 2147483647
    if true {
        var v6 int64 = 42
        v7 := []int{}
        v5 -= v5
        _ = v6
        _ = v7
    }
    _ = v5
    return -32768
}
func genFn1(p0 uint) uint16 {
    v8 := GenS1{F0: 1, F1: 0}
    v9 := []int{0, -957, 93, 128}
    switch -9223372036854775808 {
    case -508:
        var v10 int64 = 42
        _ = v10
    case 128:
        if (((!(0.0 == float32(((0.1 * ((0.1 - (1e10 * 1e10)) / (-2.25 - (3.14159 / ((0.1 / 1e10) - 0.0))))) * -2.25)))) && false) || (((("foo bar" + "") + ("" + "a")) < "foo bar") || (!true))) {
            var v11 int16 = 780
            var v12 string = ""
            for fi13 := 0; fi13 < 13; fi13++ {
                fmt.Println(p0, fi13)
                _ = fi13
            }
            _ = v11
            _ = v12
        }
    default:
        var v14 int16 = -510
        _ = v14
    }
    _ = p0
    _ = v8
    _ = v9
    return 32767
}

func main() {
    var v1 float64 = -2.25
    var v2 uint8 = 1
    var v3 int8 = -19
    var v4 uint32 = 2
    v2 = (v2 * (v2 & (^(v2))))
    for fi15 := 0; fi15 < 11; fi15++ {
        if false {
            var v16 int16 = 1000
            if (true || false) {
                switch (fi15 ^ 2147483647) {
                case 128:
                    fmt.Println(v1)
                }
                v17 := GenS1{F0: 2147483647, F1: 0}
                _ = v17
            } else {
                if true {
                    v18 := GenS1{F0: (v4 - 2), F1: 995}
                    if (!(("\n\t" + ("hello" + "hello")) != "\n\t")) {
                        fmt.Println(v2, v16, v4)
                    }
                    switch fi15 {
                    case 13:
                        if ((65535 >= 18446744073709551615) || true) {
                            var v19 int32 = 1
                            _ = v19
                        } else {
                            switch ("hello" + "a") {
                            case "hello":
                                var v20 uint8 = v2
                                _ = v20
                            case "\n\t":
                                if false {
                                    v21 := []int{(fi15 * 2147483647), (fi15 ^ (fi15 >> ((0) & 63)))}
                                    _ = v21
                                }
                            case "foo bar":
                                var v22 uint = 96
                                _ = v22
                            }
                            v1 = 1.5
                        }
                    default:
                        if true {
                            v23 := GenS1{F0: v4, F1: 790}
                            if false {
                                for fi24 := 0; fi24 < 14; fi24++ {
                                    v16 = genFn0()
                                    v2 += (^(v2))
                                    v1 += 1.5
                                    _ = fi24
                                }
                            } else {
                                v1 = ((float64(-2.25) + (0.0 - (float64(1.5) * ((v1 - (((v1 / (v1 / v1)) - (v1 * v1)) / 0.1)) + (v1 + 3.14159))))) / v1)
                                var v25 uint64 = 583
                                v16 -= (v16 | v16)
                                _ = v25
                            }
                            _ = v23
                        } else {
                            fmt.Println(v2, v2)
                            if (v1 != v1) {
                                if (!(true && true)) {
                                    var v26 int16 = genFn0()
                                    fmt.Println(fi15, v1, v2)
                                    v27 := GenS2{F0: 42, F1: 553, F2: 664}
                                    _ = v26
                                    _ = v27
                                } else {
                                    var v28 int = fi15
                                    fmt.Println(v28, v4, v2)
                                    var v29 string = ""
                                    _ = v28
                                    _ = v29
                                }
                                v30 := GenS1{F0: (v4 / ((v4) | 1)), F1: 2}
                                _ = v30
                            }
                        }
                    }
                    _ = v18
                } else {
                    var v31 int32 = 0
                    _ = v31
                }
                v32 := GenS3{F0: "a", F1: 0, F2: ""}
                if ((-2147483648 <= 733) && true) {
                    for fi33 := 0; fi33 < 15; fi33++ {
                        if (65535 != 193) {
                            var v34 float64 = ((v1 - v1) / v1)
                            var v35 uint64 = 18446744073709551615
                            v16 -= genFn0()
                            _ = v34
                            _ = v35
                        }
                        var v36 int32 = 773
                        v37 := []int{(fi33 << ((255) & 63)), (fi15 * (fi15 << ((4294967295) & 63))), 1, fi15}
                        _ = fi33
                        _ = v36
                        _ = v37
                    }
                } else {
                    v1 = v1
                }
                _ = v32
            }
            _ = v16
        }
        v4 = (v4 << ((18446744073709551615) & 63))
        _ = fi15
    }
    v38 := []int{32768, 42, 127}
    for fi39 := 0; fi39 < 3; fi39++ {
        if ((!(true && (!(!(4294967295 != 2))))) && false) {
            v40 := GenS2{F0: 255, F1: 18446744073709551615, F2: v4}
            _ = v40
        }
        v41 := GenS4{F0: 870, F1: 255}
        _ = fi39
        _ = v41
    }
    v42 := GenS2{F0: 616, F1: 1, F2: (v4 & v4)}
    if false {
        var v43 uint32 = v4
        for fi44 := 0; fi44 < 8; fi44++ {
            v38 = append(v38, fi44)
            _ = fi44
        }
        _ = v43
    }
    v45 := GenS4{F0: 1, F1: uint(v2)}
    if (492 > 250) {
        if len(v38) > 0 {
            v38[0] = 42
        }
    }
    _ = len(v38)
    v1 -= (v1 + ((v1 / 1.0) + v1))
    v46 := GenS3{F0: "hello", F1: -319, F2: "\n\t"}
    v47 := GenS5{F0: v3}
    var v48 float32 = 3.14159
    fmt.Println(v1, v2, v48)
    _ = len(v38)
    v4 += (v4 * 0)
    fmt.Println(v2, v48, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v42
    _ = v45
    _ = v46
    _ = v47
    _ = v48
}
