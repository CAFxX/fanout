package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 int8; F2 uint32 }
type GenS3 struct { F0 uint32; F1 uint8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint; F1 uint32; F2 uint8 }
type GenS6 struct { F0 uint64; F1 uint64; F2 bool }
type GenS7 struct { F0 uint }
func genFn0(p0 int32, p1 int) int8 {
    for fi6 := 0; fi6 < 16; fi6++ {
        if (32767 <= p1) {
            v7 := GenS1{F0: 107}
            for fi8 := 0; fi8 < 18; fi8++ {
                p1 = p1
                _ = fi8
            }
            _ = v7
        } else {
            var v9 float32 = (0.1 / ((0.1 - 0.0) / 0.1))
            _ = v9
        }
        _ = fi6
    }
    var v10 int16 = -404
    v11 := []int{(p1 - (p1 / ((p1) | 1))), p1}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return -1
}
func genFn1(p0 uint8) uint32 {
    fmt.Println(p0)
    v12 := []int{-9223372036854775808}
    _ = p0
    _ = v12
    return 65535
}

func main() {
    v1 := []int{127, 128}
    var v2 int32 = -686
    var v3 float32 = 1e10
    v4 := make(map[string]uint32)
    v5 := []uint32{2147483648, 4294967295, 118}
    if (genFn0(v2, 128) <= -128) {
        v13 := v4["a"]
        _ = v13
    } else {
        for fi14 := 0; fi14 < 8; fi14++ {
            var v15 int16 = 0
            _ = fi14
            _ = v15
        }
        for fi16 := 0; fi16 < 14; fi16++ {
            if true {
                var v17 bool = ((true || ((!true) && true)) && false)
                _ = v17
            } else {
                fmt.Println(v3, v3)
                switch (("foo bar" + "hello") + "a") {
                case "\n\t":
                    v3 = (float32(3.14159) - v3)
                case "a":
                    var v18 int8 = -128
                    _ = v18
                case "foo bar":
                    var v19 string = "a"
                    _ = v19
                default:
                    if (255 >= 18446744073709551615) {
                        if true {
                            v3 -= 3.14159
                        }
                        for fi20 := 0; fi20 < 10; fi20++ {
                            v21 := GenS1{F0: 0}
                            _ = fi20
                            _ = v21
                        }
                    }
                }
            }
            v2 -= (v2 & v2)
            _ = fi16
        }
    }
    switch 9223372036854775807 {
    case 0:
        v22 := GenS2{F0: 0, F1: 127, F2: 2}
        _ = v22
    }
    if false {
        switch int(v2) {
        case 42:
            if (1 != 32767) {
                v23 := GenS3{F0: 755, F1: 0}
                for fi24 := 0; fi24 < 14; fi24++ {
                    switch "hello" {
                    case "a":
                        var v25 string = ("hello" + ("" + (("a" + "\n\t") + "\n\t")))
                        _ = v25
                    case "":
                        var v26 int64 = 250
                        _ = v26
                    case "hello":
                        switch 658 {
                        case 421:
                            var v27 string = ""
                            _ = v27
                        case 319:
                            var v28 float64 = 0.0
                            _ = v28
                        case 42:
                            v29 := GenS4{F0: 42}
                            _ = v29
                        }
                    default:
                        v3 += (v3 / (v3 * v3))
                    }
                    _ = fi24
                }
                if (2 != 1) {
                    var v30 int16 = int16(v2)
                    if (!true) {
                        _ = len(v4)
                        v30 -= 0
                        fmt.Println(v3)
                    }
                    _ = v30
                }
                _ = v23
            }
        case -227:
            if (true && ("foo bar" < ("hello" + ((("foo bar" + "a") + ("hello" + "a")) + (("" + ("a" + (("hello" + "") + ("foo bar" + ("hello" + "foo bar"))))) + ""))))) {
                switch (("a" + "a") + "\n\t") {
                case "\n\t":
                    var v31 int64 = 0
                    _ = v31
                }
                switch 2147483647 {
                case 4294967295:
                    var v32 uint = 4294967295
                    _ = v32
                case 32767:
                    v5 = append(v5, 2147483648)
                case 127:
                    v3 -= (((v3 * (3.14159 / (v3 + (float32(0.1) / 0.1)))) - 3.14159) / float32(1e10))
                default:
                    v33 := GenS5{F0: 2, F1: 298, F2: 209}
                    _ = v33
                }
                fmt.Println(v3)
            }
        }
        var v34 int8 = -43
        _ = v34
    } else {
        if len(v1) > 0 {
            v1[0] = 42
        }
        if (true && (888 == 32767)) {
            var v35 uint16 = 542
            _ = v35
        } else {
            v36 := GenS2{F0: 42, F1: 69, F2: 2147483647}
            v3 = 1e10
            v2 = (v2 % (((v2 << ((865) & 63))) | 1))
            _ = v36
        }
    }
    v2 = 2147483647
    v37 := GenS6{F0: 42, F1: 1, F2: true}
    v4["hello"] = 560
    v38 := GenS7{F0: 455}
    v2 -= (v2 / (((v2 - 784)) | 1))
    fmt.Println(v2)
    v2 += -330
    switch ("\n\t" + "a") {
    case "a":
        for fi39 := 0; fi39 < 4; fi39++ {
            fmt.Println(v2, fi39, v3)
            _ = fi39
        }
    case "hello":
        var v40 int16 = -166
        _ = v40
    case "":
        var v41 uint8 = 1
        _ = v41
    default:
        var v42 uint16 = 42
        _ = v42
    }
    var v43 int64 = 9223372036854775807
    if (!(-40 == genFn0(-2147483648, 4294967295))) {
        var v44 float64 = (3.14159 + 1.5)
        if (18446744073709551615 > 42) {
            fmt.Println(v3, v2, v44)
            for fi45 := 0; fi45 < 8; fi45++ {
                v46 := GenS2{F0: 9223372036854775807, F1: 0, F2: 42}
                v2 -= 196
                _ = fi45
                _ = v46
            }
            fmt.Println(v44)
        }
        if len(v1) > 0 {
            v1[0] = 127
        }
        _ = v44
    }
    v3 = float32(0.0)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v43
    _ = v5
}
