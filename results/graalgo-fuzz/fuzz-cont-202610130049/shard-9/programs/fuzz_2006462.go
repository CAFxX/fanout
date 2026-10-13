package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int16; F1 int64; F2 int8 }
type GenS5 struct { F0 uint32; F1 string }
type GenS6 struct { F0 bool; F1 uint16 }

func main() {
    var v1 bool = (((1 < 42) && true) && true)
    var v2 uint8 = 42
    var v3 uint64 = 1
    v4 := GenS1{F0: -801, F1: 2147483648, F2: -502}
    if ((false || v1) || (42 < 1)) {
        for fi5 := 0; fi5 < 4; fi5++ {
            var v6 int16 = -670
            v7 := GenS2{F0: 255}
            var v8 uint16 = uint16(v2)
            _ = fi5
            _ = v6
            _ = v7
            _ = v8
        }
        v9 := GenS1{F0: 67, F1: 204, F2: -1}
        _ = v9
    } else {
        v10 := GenS3{F0: 1}
        _ = v10
    }
    v2 = (v2 * (v2 | (v2 >> ((42) & 63))))
    v11 := GenS3{F0: 42}
    v12 := GenS4{F0: 42, F1: -9223372036854775808, F2: 100}
    var v13 bool = (-1 < 1000)
    var v14 bool = ((127 <= -9223372036854775808) || (100 > 127))
    v2 -= v2
    var v15 float32 = (1e10 / -2.25)
    for fi16 := 0; fi16 < 3; fi16++ {
        var v17 uint16 = 0
        v18 := GenS5{F0: 1, F1: ("hello" + "")}
        var v19 float64 = 3.14159
        _ = fi16
        _ = v17
        _ = v18
        _ = v19
    }
    v20 := []int{}
    switch -1 {
    case 0:
        if (!v1) {
            for fi21 := 0; fi21 < 15; fi21++ {
                v13 = false
                _ = fi21
            }
            var v22 float64 = 1.5
            _ = v22
        } else {
            switch -2147483648 {
            case -2147483648:
                v23 := GenS5{F0: 888, F1: "foo bar"}
                _ = v23
            case 331:
                for fi24 := 0; fi24 < 3; fi24++ {
                    for fi25 := 0; fi25 < 20; fi25++ {
                        if (0 > 246) {
                            switch "a" {
                            case "hello":
                                fmt.Println(v13, fi25)
                            }
                        }
                        switch "\n\t" {
                        case "":
                            v1 = ((v15 * 1.5) != 1e10)
                        case "\n\t":
                            switch "a" {
                            case "":
                                v3 -= v3
                            case "\n\t":
                                v26 := GenS6{F0: (v1 || v14), F1: uint16(v3)}
                                _ = v26
                            default:
                                switch fi25 {
                                case -9223372036854775808:
                                    v14 = (100 >= -1)
                                case -2147483648:
                                    if len(v20) >= 2 {
                                        v20 = v20[1:len(v20)]
                                    }
                                default:
                                    v15 -= (0.1 * float32(0.1))
                                }
                            }
                        }
                        var v27 uint64 = v3
                        _ = fi25
                        _ = v27
                    }
                    fmt.Println(v14, v14)
                    var v28 float32 = (v15 * (v15 / ((v15 * ((0.1 - (0.1 - (((v15 + 1e10) * (v15 / float32(((((-2.25 / 1.0) - (-2.25 * ((0.1 + 1.5) - (1e10 * 0.1)))) * 1.5) - 1e10)))) * v15))) / v15)) * v15)))
                    _ = fi24
                    _ = v28
                }
            case -9223372036854775808:
                v1 = (-1 > 1)
            }
            fmt.Println(v13)
        }
    case 4294967295:
        v29 := GenS4{F0: 567, F1: -19, F2: 1}
        _ = v29
    default:
        if ((v1 && (!v1)) || v14) {
            if len(v20) >= 2 {
                v20 = v20[1:len(v20)]
            }
        } else {
            var v30 uint64 = 18446744073709551615
            var v31 bool = false
            _ = v30
            _ = v31
        }
    }
    switch "a" {
    case "a":
        for fi32 := 0; fi32 < 4; fi32++ {
            var v33 int8 = 22
            _ = fi32
            _ = v33
        }
    default:
        for fi34 := 0; fi34 < 14; fi34++ {
            var v35 float64 = 1.5
            _ = fi34
            _ = v35
        }
    }
    for fi36 := 0; fi36 < 15; fi36++ {
        fmt.Println(v13, v2)
        var v37 int8 = -128
        v13 = v14
        _ = fi36
        _ = v37
    }
    if v14 {
        fmt.Println(v1)
        for fi38 := 0; fi38 < 14; fi38++ {
            v3 += v3
            for fi39 := 0; fi39 < 10; fi39++ {
                v40 := GenS4{F0: 32767, F1: 9223372036854775807, F2: -5}
                _ = fi39
                _ = v40
            }
            switch (("foo bar" + ("foo bar" + ("\n\t" + (("hello" + ("\n\t" + "")) + "a")))) + "") {
            case "":
                v3 += 0
            case "hello":
                if true {
                    var v41 bool = (3.14159 == (1e10 + ((float64(v15) - (float64((v15 * v15)) + 1e10)) * (0.1 * 1e10))))
                    if (!v14) {
                        v3 -= 4294967295
                    }
                    var v42 uint64 = uint64(v2)
                    _ = v41
                    _ = v42
                }
            case "\n\t":
                var v43 float64 = 3.14159
                _ = v43
            }
            _ = fi38
        }
        fmt.Println(v15, v1)
    }
    if len(v20) >= 2 {
        v20 = v20[1:len(v20)]
    }
    v44 := GenS2{F0: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v44
}
