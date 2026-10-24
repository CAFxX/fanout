package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 bool; F1 int8; F2 string }
type GenS3 struct { F0 int; F1 bool }
type GenS4 struct { F0 int32; F1 uint8 }

func main() {
    var v1 int8 = 100
    v2 := make(map[int]int)
    var v3 uint8 = 97
    var v4 uint64 = 0
    var v5 int = 127
    var v6 uint8 = (^(v3))
    v7 := []int{(-(v5)), v5, v5}
    v6 = v6
    for fi8 := 0; fi8 < 14; fi8++ {
        v9 := GenS1{F0: (fi8 ^ (v5 & (fi8 * fi8))), F1: 189}
        var v10 float32 = float32(float64(-2.25))
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 int16 = -932
    if (!true) {
        if (v6 == v3) {
            for fi12 := 0; fi12 < 8; fi12++ {
                var v13 uint16 = 2
                v14 := GenS1{F0: (v5 & (^(fi12))), F1: 2147483647}
                _ = fi12
                _ = v13
                _ = v14
            }
            v4 = 1
            var v15 uint64 = v4
            _ = v15
        }
        switch (("\n\t" + "a") + "hello") {
        case "hello":
            var v16 float32 = (0.0 - 0.0)
            _ = v16
        case "\n\t":
            var v17 uint16 = 42
            _ = v17
        }
    } else {
        v18 := GenS1{F0: v5, F1: 2147483647}
        if false {
            if false {
                if false {
                    switch "a" {
                    case "\n\t":
                        v11 = v11
                    case "foo bar":
                        var v19 bool = true
                        _ = v19
                    case "":
                        var v20 int = v5
                        _ = v20
                    default:
                        var v21 uint64 = v4
                        _ = v21
                    }
                    var v22 uint = 65535
                    switch "" {
                    case "foo bar":
                        v23 := GenS2{F0: false, F1: v1, F2: "hello"}
                        _ = v23
                    case "a":
                        v24 := GenS1{F0: v5, F1: 2147483648}
                        _ = v24
                    }
                    _ = v22
                }
            }
            v4 -= 905
            if false {
                var v25 float32 = 1.5
                var v26 int16 = (v11 / ((v11) | 1))
                _ = v25
                _ = v26
            }
        }
        v27 := GenS2{F0: ((-2.25 + -2.25) >= (1e10 - -2.25)), F1: -1, F2: ("" + "")}
        _ = v18
        _ = v27
    }
    var v28 bool = (v4 != (v4 ^ v4))
    var v29 float32 = 0.0
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    for fi30 := 0; fi30 < 12; fi30++ {
        switch "\n\t" {
        case "\n\t":
            if ((v11 & int16(v5)) == v11) {
                v31 := GenS2{F0: false, F1: (v1 << ((752) & 63)), F2: "\n\t"}
                v32 := v2[v5]
                _ = v31
                _ = v32
            } else {
                _ = len(v2)
                var v33 bool = false
                if v33 {
                    if (0 > 153) {
                        switch ("hello" + "") {
                        case "hello":
                            var v34 string = ("a" + ((((("\n\t" + "") + "") + ("\n\t" + ("hello" + ("a" + "")))) + "foo bar") + ((("" + "a") + "") + (("a" + "") + ("" + ("foo bar" + ((("a" + "hello") + "foo bar") + "\n\t")))))))
                            _ = v34
                        case "a":
                            fmt.Println(v29, v29, v3)
                        }
                    }
                } else {
                    fmt.Println(v1, v33)
                }
                _ = v33
            }
        case "hello":
            v35 := v2[(^(fi30))]
            _ = v35
        }
        if len(v7) > 0 {
            v7[0] = (fi30 << ((4294967295) & 63))
        }
        if v28 {
            for fi36 := 0; fi36 < 10; fi36++ {
                var v37 uint64 = v4
                fmt.Println(v29)
                var v38 int32 = 1
                _ = fi36
                _ = v37
                _ = v38
            }
            for fi39 := 0; fi39 < 20; fi39++ {
                if v28 {
                    if (v28 && true) {
                        v29 = 1e10
                        if ((v4 - 1) < v4) {
                            if len(v7) > 0 {
                                v7[0] = 4294967295
                            }
                        }
                        if (0.1 >= (-2.25 - (1.5 - 1e10))) {
                            v7 = append(v7, (fi30 * (fi30 & 32767)))
                            fmt.Println(v4, fi30, v29)
                        }
                    } else {
                        if false {
                            v2[(-(v5))] = fi30
                            v40 := GenS1{F0: (fi39 % ((fi30) | 1)), F1: 65535}
                            _ = v40
                        }
                        v6 += (v3 | 0)
                    }
                    v4 -= (v4 / ((0) | 1))
                } else {
                    switch ("a" + ("foo bar" + "a")) {
                    case "foo bar":
                        _ = len(v2)
                    case "a":
                        v41 := GenS2{F0: v28, F1: (v1 | 127), F2: "hello"}
                        _ = v41
                    case "":
                        fmt.Println(v6, v4, v5)
                    }
                }
                if false {
                    var v42 int16 = (v11 >> ((497) & 63))
                    fmt.Println(v11, v4)
                    switch "" {
                    case "a":
                        v29 -= -2.25
                    default:
                        switch (("" + ("foo bar" + "\n\t")) + "\n\t") {
                        case "hello":
                            var v43 int64 = 1
                            _ = v43
                        case "foo bar":
                            if (65 < v1) {
                                fmt.Println(v28, v11)
                            }
                        case "":
                            switch (v5 * (fi39 | fi39)) {
                            case -9223372036854775808:
                                fmt.Println(v6, fi39)
                            default:
                                v42 += (v11 * (-(v11)))
                            }
                        default:
                            _ = len(v7)
                            if true {
                                v7 = append(v7, fi30)
                            }
                            v29 -= (v29 - v29)
                            v6 = v6
                        }
                    }
                    _ = v42
                } else {
                    fmt.Println(v29, v29, v11)
                    v29 = 0.0
                    fmt.Println(v11, v4, v1)
                    if true {
                        v3 -= 255
                        v29 += ((((v29 * 1.5) / v29) + float32(0.1)) * (((float32(0.1) / v29) * (v29 / v29)) / 3.14159))
                        fmt.Println(v5, v4, v3)
                    } else {
                        var v44 bool = ((v1 >> ((2) & 63)) != 100)
                        if v44 {
                            fmt.Println(v44)
                        } else {
                            fmt.Println(v1)
                            fmt.Println(v29, v5)
                            var v45 int64 = -369
                            _ = v45
                        }
                        v46 := GenS2{F0: v44, F1: (v1 ^ int8(v11)), F2: (("" + ("\n\t" + "foo bar")) + "a")}
                        _ = v44
                        _ = v46
                    }
                    v1 += v1
                }
                v47 := GenS3{F0: (v5 % ((42) | 1)), F1: false}
                _ = fi39
                _ = v47
            }
        } else {
            for fi48 := 0; fi48 < 5; fi48++ {
                v5 = -9223372036854775808
                var v49 uint64 = v4
                _ = fi48
                _ = v49
            }
        }
        _ = fi30
    }
    var v50 uint16 = uint16(v3)
    for fi51 := 0; fi51 < 19; fi51++ {
        v52 := GenS3{F0: 32768, F1: (!v28)}
        _ = fi51
        _ = v52
    }
    v11 = (^(v11))
    delete(v2, v5)
    if ((v5 ^ -9223372036854775808) <= 621) {
        if (!true) {
            v50 -= uint16(v4)
            var v53 int = 0
            fmt.Println(v28)
            _ = v53
        }
        for fi54 := 0; fi54 < 16; fi54++ {
            v55 := GenS4{F0: -2147483648, F1: (v3 % ((127) | 1))}
            _ = fi54
            _ = v55
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v6
    _ = v7
}
