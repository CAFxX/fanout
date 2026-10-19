package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 uint32 }
func genFn0(p0 int8) int64 {
    switch "hello" {
    case "\n\t":
        if (42 == 0) {
            for fi5 := 0; fi5 < 16; fi5++ {
                switch "\n\t" {
                case "\n\t":
                    for fi6 := 0; fi6 < 11; fi6++ {
                        fmt.Println(fi6, fi5)
                        if true {
                            if (42 == 923) {
                                var v7 int64 = 9223372036854775807
                                _ = v7
                            } else {
                                var v8 bool = (!true)
                                var v9 uint8 = 176
                                v9 = v9
                                _ = v8
                                _ = v9
                            }
                        }
                        var v10 float64 = 3.14159
                        _ = fi6
                        _ = v10
                    }
                }
                _ = fi5
            }
        }
    case "hello":
        p0 -= -128
    default:
        fmt.Println(p0)
    }
    if (false || (false || (false && (0 != 85)))) {
        var v11 int = -9223372036854775808
        v12 := []int{(v11 / ((v11) | 1)), -749}
        var v13 bool = false
        _ = v11
        _ = v12
        _ = v13
    }
    fmt.Println(p0)
    for fi14 := 0; fi14 < 1; fi14++ {
        p0 += (p0 << ((1) & 63))
        var v15 float32 = (3.14159 + 1.5)
        var v16 float32 = v15
        _ = fi14
        _ = v15
        _ = v16
    }
    _ = p0
    return 246
}

func main() {
    var v1 float32 = -2.25
    var v2 float32 = v1
    v3 := []int{2147483647, 0}
    var v4 uint64 = 311
    v17 := GenS1{F0: 961, F1: 32768, F2: 1}
    var v18 int64 = 1
    v19 := GenS1{F0: 2, F1: 222, F2: 2147483647}
    var v20 uint32 = 42
    var v21 int32 = -64
    switch -2147483648 {
    case 32767:
        v22 := GenS1{F0: v20, F1: 42, F2: v20}
        _ = v22
    case -1:
        var v23 int8 = 127
        _ = v23
    default:
        if (596 >= (v4 ^ 18446744073709551615)) {
            v4 = v4
            v1 = -2.25
        }
    }
    for fi24 := 0; fi24 < 15; fi24++ {
        var v25 int64 = v18
        if (!(fi24 > 887)) {
            var v26 string = "a"
            var v27 int16 = 579
            _ = v26
            _ = v27
        } else {
            var v28 uint64 = 4294967295
            _ = len(v3)
            _ = v28
        }
        _ = fi24
        _ = v25
    }
    var v29 bool = (!(((!false) && ("foo bar" <= "a")) || (-128 < 100)))
    if (96 >= 255) {
        var v30 uint8 = 42
        if false {
            var v31 uint = 81
            _ = v31
        }
        _ = v30
    }
    if len(v3) > 0 {
        v3[0] = -9223372036854775808
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v32 := GenS1{F0: 949, F1: 32768, F2: (v20 % ((v20) | 1))}
    fmt.Println(v21, v2)
    for fi33 := 0; fi33 < 2; fi33++ {
        if (((fi33 % ((861) | 1)) >= (^(fi33))) && v29) {
            v34 := GenS1{F0: v20, F1: 2147483647, F2: (v20 % (((^(v20))) | 1))}
            _ = v34
        }
        v1 = v1
        var v35 uint64 = v4
        _ = fi33
        _ = v35
    }
    if false {
        fmt.Println(v1, v29, v1)
    } else {
        switch "\n\t" {
        case "foo bar":
            fmt.Println(v20)
        case "":
            var v36 int = 32768
            _ = v36
        }
        for fi37 := 0; fi37 < 20; fi37++ {
            for fi38 := 0; fi38 < 14; fi38++ {
                v18 = 1
                v39 := GenS1{F0: 574, F1: (fi37 & (-(fi37))), F2: v20}
                _ = fi38
                _ = v39
            }
            for fi40 := 0; fi40 < 8; fi40++ {
                if true {
                    var v41 string = "\n\t"
                    _ = v41
                } else {
                    v29 = v29
                    switch ("\n\t" + "a") {
                    case "":
                        v42 := GenS1{F0: v20, F1: fi37, F2: (v20 / (((v20 >> ((65535) & 63))) | 1))}
                        _ = v42
                    case "\n\t":
                        var v43 uint8 = 128
                        _ = v43
                    case "hello":
                        fmt.Println(v29, fi40)
                    default:
                        var v44 string = (("hello" + "") + "hello")
                        _ = v44
                    }
                }
                _ = fi40
            }
            v3 = append(v3, (fi37 >> ((928) & 63)))
            _ = fi37
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v29
    _ = v3
    _ = v32
    _ = v4
}
