package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 string }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int; F1 uint16; F2 uint16 }
func genFn0(p0 bool, p1 int8, p2 int64) int8 {
    var v7 int32 = 1
    var v8 int64 = p2
    p0 = (1.5 > ((1.5 + (float32(0.0) / 1.5)) + 0.1))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return (^(p1))
}

func main() {
    var v1 uint16 = 610
    var v2 int16 = -32768
    var v3 uint64 = 42
    var v4 uint32 = 0
    v5 := make(map[int]int)
    var v6 uint16 = v1
    v9 := []int{434, 42}
    var v10 uint16 = v6
    var v11 int8 = genFn0(false, -58, 0)
    for fi12 := 0; fi12 < 4; fi12++ {
        if ((true && true) && (!false)) {
            v3 = 4294967295
            var v13 string = ""
            _ = v13
        } else {
            var v14 uint16 = (v1 + v6)
            v2 -= v2
            var v15 int32 = -2147483648
            _ = v14
            _ = v15
        }
        for fi16 := 0; fi16 < 13; fi16++ {
            if true {
                if (((("hello" + "hello") + "foo bar") + "a") > "foo bar") {
                    if false {
                        var v17 int = fi16
                        v2 += (v2 & (-(v2)))
                        var v18 uint16 = v6
                        v4 = (^(v4))
                        var v19 string = ((((("foo bar" + (((("" + "foo bar") + "a") + "\n\t") + (("hello" + "") + ""))) + "\n\t") + (("a" + (("hello" + "") + "hello")) + ("\n\t" + "\n\t"))) + "") + ("hello" + "\n\t"))
                        _ = v17
                        _ = v18
                        _ = v19
                    }
                    v10 -= 255
                    var v20 uint8 = 2
                    _ = v20
                }
                v21 := GenS1{F0: 809, F1: 415, F2: "a"}
                if len(v9) > 0 {
                    v9[0] = (fi16 + (fi16 * fi12))
                }
                fmt.Println(v4)
                _ = v21
            } else {
                var v22 int64 = int64(fi16)
                _ = v22
            }
            v23 := GenS1{F0: 106, F1: -7, F2: (("a" + (((("" + ("foo bar" + "foo bar")) + "\n\t") + ("" + "\n\t")) + "hello")) + "")}
            _ = fi16
            _ = v23
        }
        delete(v5, fi12)
        _ = fi12
    }
    fmt.Println(v3)
    for fi24 := 0; fi24 < 9; fi24++ {
        if true {
            v6 -= 32767
        }
        var v25 float64 = float64(0.0)
        _ = fi24
        _ = v25
    }
    var v26 uint32 = (v4 % ((v4) | 1))
    v4 += (^(v4))
    switch "" {
    case "\n\t":
        var v27 int16 = (-(v2))
        _ = v27
    case "foo bar":
        if false {
            switch "a" {
            case "foo bar":
                if false {
                    fmt.Println(v10, v10)
                    v1 += v6
                } else {
                    v4 += v26
                    v5[9223372036854775807] = -587
                    switch "foo bar" {
                    case "\n\t":
                        v9 = append(v9, -570)
                    case "a":
                        _ = len(v9)
                    case "":
                        var v28 int64 = 42
                        _ = v28
                    }
                }
            case "a":
                v9 = append(v9, 42)
            case "\n\t":
                switch "a" {
                case "a":
                    v29 := GenS1{F0: 2, F1: 1, F2: ("\n\t" + "hello")}
                    _ = v29
                case "\n\t":
                    v4 = (v26 >> ((65535) & 63))
                case "":
                    for fi30 := 0; fi30 < 12; fi30++ {
                        v31 := GenS2{F0: 42}
                        v32 := GenS1{F0: 1, F1: 1, F2: ("\n\t" + "hello")}
                        for fi33 := 0; fi33 < 9; fi33++ {
                            switch (fi30 << ((317) & 63)) {
                            case -796:
                                v4 = v26
                            default:
                                var v34 uint64 = v3
                                _ = v34
                            }
                            switch fi33 {
                            case 381:
                                _ = len(v9)
                            default:
                                var v35 string = ("a" + (("" + "a") + ("" + "hello")))
                                _ = v35
                            }
                            _ = fi33
                        }
                        _ = fi30
                        _ = v31
                        _ = v32
                    }
                }
            }
            v6 = 65535
        }
    }
    v26 = 2
    v11 += 100
    for fi36 := 0; fi36 < 9; fi36++ {
        v26 += uint32(v3)
        var v37 int64 = 747
        var v38 int64 = (v37 | v37)
        _ = fi36
        _ = v37
        _ = v38
    }
    v2 -= 0
    if true {
        if (927 > 372) {
            v11 = (v11 & genFn0(false, (^(v11)), -9223372036854775808))
        } else {
            if true {
                for fi39 := 0; fi39 < 3; fi39++ {
                    if false {
                        fmt.Println(v10, v4)
                        v2 = -1
                        if false {
                            switch fi39 {
                            case -600:
                                v40 := GenS3{F0: fi39, F1: v10, F2: (v1 / ((v10) | 1))}
                                _ = v40
                            case 32768:
                                for fi41 := 0; fi41 < 6; fi41++ {
                                    v1 = 42
                                    v26 = v26
                                    _ = fi41
                                }
                            default:
                                v10 += (v10 & v10)
                            }
                            v42 := GenS3{F0: (^(fi39)), F1: v1, F2: (^(v1))}
                            _ = v42
                        }
                    }
                    var v43 float64 = (0.1 / (1e10 * ((1.5 - (((3.14159 / -2.25) * ((0.0 * (float64((0.1 / (1.5 / (-2.25 - (-2.25 + 3.14159))))) + (1.5 / -2.25))) - 1.5)) / 1e10)) / -2.25)))
                    _ = fi39
                    _ = v43
                }
            } else {
                var v44 int32 = -1
                _ = v44
            }
        }
    }
    if (9223372036854775807 > 0) {
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        fmt.Println(v26, v1, v10)
    } else {
        for fi45 := 0; fi45 < 14; fi45++ {
            v2 = v2
            _ = len(v5)
            v10 -= 2
            _ = fi45
        }
    }
    switch "foo bar" {
    case "a":
        v6 -= (^(v6))
    case "hello":
        var v46 uint16 = (^(v10))
        _ = v46
    case "foo bar":
        v47 := GenS1{F0: 18446744073709551615, F1: 986, F2: "\n\t"}
        _ = v47
    }
    for fi48 := 0; fi48 < 13; fi48++ {
        var v49 float64 = ((float64(3.14159) - 0.0) - ((-2.25 * 1e10) / float64(-2.25)))
        v1 -= 32767
        v11 = genFn0(true, (v11 + (v11 * v11)), -9223372036854775808)
        _ = fi48
        _ = v49
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
