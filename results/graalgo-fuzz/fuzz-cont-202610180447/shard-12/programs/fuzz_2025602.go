package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 int; F1 int32; F2 uint32 }
type GenS3 struct { F0 bool; F1 uint64 }

func main() {
    v1 := make(map[int]string)
    v2 := []int{-1, -403, 619, 4294967295}
    var v3 uint16 = 42
    var v4 float64 = 3.14159
    var v5 int64 = 0
    fmt.Println(v3, v5, v3)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v6 := GenS1{F0: 1, F1: -96}
    var v7 int32 = -1
    for fi8 := 0; fi8 < 14; fi8++ {
        v4 = -2.25
        for fi9 := 0; fi9 < 4; fi9++ {
            v4 += 0.1
            if (1 < 4294967295) {
                v1[fi8] = "a"
                v3 = 32767
                switch 940 {
                case 1:
                    if true {
                        v10 := GenS2{F0: (fi8 * fi8), F1: (v7 << ((65535) & 63)), F2: 978}
                        _ = v10
                    }
                default:
                    switch int(v5) {
                    case -912:
                        switch (-(fi9)) {
                        case 4294967295:
                            if ((v5 + v5) > v5) {
                                v7 += v7
                                v4 = (v4 / 1.0)
                                var v11 uint = 0
                                _ = v11
                            }
                        default:
                            _ = len(v2)
                        }
                    case 2147483647:
                        if (!true) {
                            v4 = (((v4 + v4) * float64(1e10)) + v4)
                            v5 = v5
                            v12 := GenS1{F0: 42, F1: 42}
                            _ = v12
                        }
                    case 716:
                        delete(v1, fi9)
                    }
                }
            } else {
                if false {
                    fmt.Println(fi9, v7)
                    v7 -= (-(v7))
                    if ((v3 < 175) || (!(0 >= 99))) {
                        if (!true) {
                            if len(v2) > 0 {
                                v2[0] = 32768
                            }
                            if len(v2) > 0 {
                                v2[0] = 108
                            }
                            v5 = v5
                        } else {
                            if true {
                                fmt.Println(fi8)
                                switch "hello" {
                                case "":
                                    v7 += v7
                                    v3 -= 255
                                    if (4294967295 > 0) {
                                        fmt.Println(v7, fi9)
                                        v4 -= 0.0
                                        v13 := GenS3{F0: true, F1: 2}
                                        if len(v2) > 0 {
                                            v2[0] = int(v3)
                                        }
                                        _ = v13
                                    }
                                case "hello":
                                    var v14 string = ((("\n\t" + "a") + "foo bar") + "foo bar")
                                    _ = v14
                                case "foo bar":
                                    if (!true) {
                                        if true {
                                            v3 = (^(v3))
                                        }
                                        if (((1.5 - (3.14159 * float32((-2.25 / v4)))) / 0.1) > 3.14159) {
                                            v3 = (v3 | v3)
                                        } else {
                                            v15 := GenS1{F0: 42, F1: -36}
                                            _ = v15
                                        }
                                        v16 := GenS2{F0: (fi9 & (^(fi9))), F1: -969, F2: 2147483648}
                                        _ = v16
                                    }
                                }
                            } else {
                                switch ("a" + "") {
                                case "\n\t":
                                    v3 += (v3 % ((32767) | 1))
                                }
                                if (!true) {
                                    if false {
                                        fmt.Println(v4, v3, v7)
                                        fmt.Println(v5, v4, v4)
                                    } else {
                                        var v17 int = fi8
                                        v5 = (v5 + 1)
                                        _ = v17
                                    }
                                    var v18 uint64 = 18446744073709551615
                                    var v19 int8 = 0
                                    _ = v18
                                    _ = v19
                                }
                                v4 -= (((v4 * 0.1) / 1.5) / ((v4 / v4) - 0.1))
                                if (false && (803 > 0)) {
                                    var v20 int = 32768
                                    delete(v1, -9223372036854775808)
                                    _ = v20
                                }
                            }
                            switch ("hello" + "hello") {
                            case "a":
                                v5 += (v5 ^ 9223372036854775807)
                            }
                        }
                        if (float32(v4) > (3.14159 + 1e10)) {
                            var v21 float32 = ((float32(-2.25) + -2.25) / (0.1 - float32((v4 * ((0.1 - 0.1) / float64(3.14159))))))
                            if (100 == 127) {
                                v3 -= 255
                            }
                            _ = v21
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    } else {
                        v5 -= (v5 >> ((480) & 63))
                    }
                    v4 = (0.0 - v4)
                    fmt.Println(fi9)
                }
            }
            v4 = (-2.25 * ((0.0 + v4) - (v4 + v4)))
            _ = fi9
        }
        _ = fi8
    }
    if false {
        v3 = 281
        var v22 uint8 = 255
        for fi23 := 0; fi23 < 2; fi23++ {
            for fi24 := 0; fi24 < 4; fi24++ {
                switch 9223372036854775807 {
                case -419:
                    fmt.Println(fi24, v3, v22)
                    var v25 int32 = (v7 | 266)
                    v26 := GenS3{F0: ("foo bar" < "\n\t"), F1: 4294967295}
                    _ = v25
                    _ = v26
                case 42:
                    v27 := GenS2{F0: 2147483647, F1: v7, F2: uint32(v3)}
                    _ = v27
                case 128:
                    switch "hello" {
                    case "hello":
                        _ = len(v2)
                    }
                default:
                    if false {
                        fmt.Println(v5, v5)
                        var v28 string = "hello"
                        _ = v28
                    }
                }
                v5 += (v5 / ((v5) | 1))
                fmt.Println(v22, v5)
                _ = fi24
            }
            _ = fi23
        }
        _ = v22
    } else {
        var v29 uint = 18446744073709551615
        v5 += v5
        _ = v29
    }
    v7 -= v7
    v4 += v4
    var v30 float64 = -2.25
    fmt.Println(v4, v7)
    v31 := GenS1{F0: -32768, F1: 0}
    for fi32 := 0; fi32 < 14; fi32++ {
        v3 -= v3
        _ = fi32
    }
    if (("\n\t" + "hello") > "a") {
        for fi33 := 0; fi33 < 3; fi33++ {
            var v34 bool = false
            _ = fi33
            _ = v34
        }
        var v35 float64 = (float64(3.14159) * float64(-2.25))
        if (v7 >= -2147483648) {
            var v36 uint64 = 42
            _ = v36
        }
        _ = v35
    } else {
        if len(v2) > 0 {
            v2[0] = 0
        }
        v5 += v5
    }
    var v37 float32 = (0.1 * -2.25)
    for fi38 := 0; fi38 < 20; fi38++ {
        v3 -= 32767
        _ = fi38
    }
    if true {
        fmt.Println(v37, v37)
        v5 += v5
        if true {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v39 int32 = v7
            for fi40 := 0; fi40 < 9; fi40++ {
                v5 = v5
                _ = len(v2)
                var v41 uint = 0
                _ = fi40
                _ = v41
            }
            _ = v39
        } else {
            v7 = 42
            if false {
                var v42 int64 = 9223372036854775807
                if len(v2) > 0 {
                    v2[0] = 127
                }
                v43 := GenS1{F0: -498, F1: 127}
                _ = v42
                _ = v43
            } else {
                var v44 float32 = v37
                _ = v44
            }
            if len(v2) > 0 {
                v2[0] = 835
            }
        }
    } else {
        for fi45 := 0; fi45 < 9; fi45++ {
            _ = len(v2)
            for fi46 := 0; fi46 < 2; fi46++ {
                v37 = v37
                _ = fi46
            }
            v3 -= (v3 + 1)
            _ = fi45
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
