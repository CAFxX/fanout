package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64; F1 uint32; F2 uint32 }
func genFn0(p0 int8) int32 {
    p0 = (p0 + -1)
    if true {
        if (!(true && (false || false))) {
            for fi6 := 0; fi6 < 16; fi6++ {
                for fi7 := 0; fi7 < 11; fi7++ {
                    switch "\n\t" {
                    case "hello":
                        var v8 int32 = -1
                        v8 = (^(v8))
                        _ = v8
                    case "\n\t":
                        p0 = (-(p0))
                    case "foo bar":
                        fmt.Println(fi7)
                    }
                    var v9 uint16 = 32767
                    var v10 uint8 = 0
                    _ = fi7
                    _ = v10
                    _ = v9
                }
                _ = fi6
            }
            p0 -= p0
        }
        switch 128 {
        case -2147483648:
            if (695 != 505) {
                for fi11 := 0; fi11 < 12; fi11++ {
                    fmt.Println(p0)
                    switch fi11 {
                    case 0:
                        switch "\n\t" {
                        case "foo bar":
                            v12 := []int{-9223372036854775808, -9223372036854775808}
                            _ = v12
                        }
                    case 128:
                        for fi13 := 0; fi13 < 12; fi13++ {
                            switch (("" + ("foo bar" + (("" + "foo bar") + "a"))) + ((("hello" + "hello") + "hello") + "\n\t")) {
                            case "hello":
                                fmt.Println(fi11, fi11, fi13)
                            case "a":
                                p0 = (p0 + -128)
                            case "":
                                p0 = p0
                            }
                            if ((true || (1 == (fi11 % ((fi13) | 1)))) || ((p0 == (-(p0))) && false)) {
                                var v14 uint16 = 0
                                _ = v14
                            }
                            _ = fi13
                        }
                    case 42:
                        var v15 int = (-(fi11))
                        _ = v15
                    default:
                        if false {
                            for fi16 := 0; fi16 < 12; fi16++ {
                                if (true && true) {
                                    p0 = 73
                                    if true {
                                        var v17 bool = true
                                        var v18 uint8 = 0
                                        fmt.Println(v17, v18, fi16)
                                        _ = v17
                                        _ = v18
                                    }
                                    var v19 uint = 937
                                    _ = v19
                                }
                                _ = fi16
                            }
                        }
                    }
                    _ = fi11
                }
            }
        case -1:
            if (1 != 127) {
                var v20 float32 = ((0.1 + 0.1) - ((1e10 + ((((-2.25 + (3.14159 - ((-2.25 / (0.0 - (1.5 / 1e10))) / ((float32(3.14159) - (1.5 - 1e10)) * -2.25)))) * 1.5) - (0.1 - 1.5)) + 1.5)) * (0.0 - ((3.14159 + 0.0) + (-2.25 + (((0.1 * 1.5) - (1.5 * 1e10)) * 1.5))))))
                _ = v20
            }
        }
        v21 := GenS1{F0: 140, F1: 1}
        _ = v21
    }
    v22 := []int{32767}
    switch "foo bar" {
    case "hello":
        v23 := GenS2{F0: -9223372036854775808}
        _ = v23
    case "a":
        for fi24 := 0; fi24 < 15; fi24++ {
            p0 = p0
            p0 = p0
            if (!(2 > 131)) {
                for fi25 := 0; fi25 < 16; fi25++ {
                    if (false || (("\n\t" >= "\n\t") && true)) {
                        p0 = p0
                    } else {
                        p0 += p0
                        p0 -= (p0 % ((p0) | 1))
                    }
                    _ = fi25
                }
                var v26 bool = false
                _ = v26
            } else {
                _ = len(v22)
                var v27 int8 = (^(p0))
                var v28 uint16 = 1
                _ = v27
                _ = v28
            }
            _ = fi24
        }
    case "\n\t":
        _ = len(v22)
    }
    _ = p0
    _ = v22
    return -1
}

func main() {
    var v1 int = 1
    v2 := make(map[uint32]bool)
    var v3 uint8 = 0
    var v4 uint64 = 1
    v5 := make(map[string]bool)
    if (0.1 < 0.1) {
        v29 := GenS3{F0: 1, F1: 2, F2: 2147483648}
        for fi30 := 0; fi30 < 8; fi30++ {
            v31 := GenS2{F0: 0}
            _ = fi30
            _ = v31
        }
        switch "foo bar" {
        case "foo bar":
            v32 := GenS3{F0: -690, F1: 2147483647, F2: 2147483647}
            _ = v32
        case "a":
            fmt.Println(v1)
        case "\n\t":
            if false {
                fmt.Println(v3, v1, v4)
            }
        }
        _ = v29
    }
    v4 = 0
    fmt.Println(v3, v4, v1)
    v33 := []int{(v1 >> ((2) & 63)), (v1 - -7)}
    if ((251 >= 65535) && true) {
        v4 = (v4 & (v4 % (((^(v4))) | 1)))
        v1 = (^(v1))
    }
    v1 += (v1 - (v1 / ((v1) | 1)))
    if true {
        var v34 int16 = 1000
        for fi35 := 0; fi35 < 15; fi35++ {
            v36 := GenS2{F0: int64(v34)}
            _ = fi35
            _ = v36
        }
        v37 := GenS3{F0: int64(v34), F1: 263, F2: 1}
        _ = v34
        _ = v37
    }
    switch v1 {
    case -186:
        for fi38 := 0; fi38 < 20; fi38++ {
            fmt.Println(v4)
            _ = fi38
        }
    }
    v1 += (v1 >> ((18446744073709551615) & 63))
    v1 += v1
    var v39 uint32 = 42
    v1 += (v1 - 2147483647)
    v5[(("foo bar" + "") + "foo bar")] = (!((!(!((!(-1 >= 0)) && false))) || (!false)))
    if len(v33) > 0 {
        v33[0] = (v1 / (((v1 % ((-511) | 1))) | 1))
    }
    var v40 uint8 = (^(v3))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}
