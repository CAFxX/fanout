package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint64; F1 uint }
type GenS6 struct { F0 uint32; F1 int16 }

func main() {
    var v1 string = "foo bar"
    var v2 int64 = -9223372036854775808
    var v3 int = 45
    v4 := []int{}
    var v5 uint8 = 1
    var v6 int8 = 127
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v5, v6)
    for fi7 := 0; fi7 < 8; fi7++ {
        var v8 uint = 793
        for fi9 := 0; fi9 < 18; fi9++ {
            var v10 uint16 = 255
            _ = fi9
            _ = v10
        }
        var v11 uint32 = 2147483647
        _ = fi7
        _ = v11
        _ = v8
    }
    if true {
        v12 := GenS1{F0: 0}
        v3 -= v3
        if (((v2 != v2) || false) && ((v5 & 214) < (^(v5)))) {
            fmt.Println(v3, v6)
        }
        _ = v12
    }
    v2 = v2
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch (("hello" + "\n\t") + "hello") {
    case "":
        if true {
            var v13 uint32 = 2147483647
            v14 := GenS1{F0: 375}
            var v15 uint16 = 255
            _ = v13
            _ = v14
            _ = v15
        } else {
            v6 += int8(v3)
            v16 := GenS2{F0: 2147483647}
            v4 = append(v4, -9223372036854775808)
            _ = v16
        }
    }
    if ((32767 >= 372) || true) {
        if false {
            for fi17 := 0; fi17 < 14; fi17++ {
                if false {
                    if (false && (-88 > 2147483647)) {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        if true {
                            v2 = 9223372036854775807
                        } else {
                            if true {
                                switch 0 {
                                case -9223372036854775808:
                                    v18 := GenS3{F0: 71}
                                    _ = v18
                                case -2147483648:
                                    fmt.Println(v5, v1, v5)
                                }
                            }
                        }
                    }
                    if len(v4) > 0 {
                        v4[0] = (-(fi17))
                    }
                    v1 = "a"
                }
                _ = fi17
            }
            var v19 bool = (false && (!((v2 | (v2 % (((v2 * (v2 | (v2 ^ v2)))) | 1))) <= (v2 * int64(v6)))))
            _ = v19
        }
        var v20 uint32 = 0
        v21 := GenS4{F0: -9223372036854775808}
        _ = v20
        _ = v21
    }
    v4 = append(v4, v3)
    v22 := GenS5{F0: 1, F1: 473}
    v23 := GenS2{F0: -296}
    v24 := GenS3{F0: 42}
    switch "a" {
    case "\n\t":
        if false {
            if (65535 > 42) {
                if true {
                    switch (v1 + "hello") {
                    case "foo bar":
                        var v25 int16 = -172
                        _ = v25
                    default:
                        var v26 int16 = -1
                        _ = v26
                    }
                    switch (v1 + ("\n\t" + v1)) {
                    case "foo bar":
                        for fi27 := 0; fi27 < 9; fi27++ {
                            if len(v4) > 0 {
                                v4[0] = v3
                            }
                            _ = fi27
                        }
                    default:
                        var v28 uint8 = 1
                        _ = v28
                    }
                    v29 := GenS2{F0: -2147483648}
                    _ = v29
                } else {
                    var v30 float64 = 1e10
                    if (v1 <= (v1 + ("foo bar" + "hello"))) {
                        var v31 uint32 = 1
                        v32 := GenS6{F0: uint32(v5), F1: 0}
                        fmt.Println(v5, v1)
                        _ = v31
                        _ = v32
                    }
                    v6 += v6
                    _ = v30
                }
                v2 = (v2 % (((v2 - -798)) | 1))
                for fi33 := 0; fi33 < 1; fi33++ {
                    var v34 int32 = 928
                    _ = fi33
                    _ = v34
                }
            }
        }
    }
    var v35 bool = true
    v36 := GenS5{F0: 42, F1: 255}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
