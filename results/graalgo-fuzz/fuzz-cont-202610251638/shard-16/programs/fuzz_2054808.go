package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 int8 }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 int32; F1 uint8; F2 int }
type GenS4 struct { F0 uint64; F1 int64; F2 uint32 }

func main() {
    v1 := []string{("hello" + ("a" + "hello")), ""}
    var v2 int8 = -81
    v3 := []int{}
    var v4 uint8 = 255
    v5 := []int{-1, 4294967295, 530}
    var v6 uint32 = 2147483648
    v7 := GenS1{F0: 151, F1: 0, F2: v2}
    v8 := GenS2{F0: 9223372036854775807, F1: ("" + "a")}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi9 := 0; fi9 < 9; fi9++ {
        v6 -= 1
        v6 += v6
        switch "\n\t" {
        case "hello":
            var v10 uint8 = (v4 % (((v4 & (v4 + 102))) | 1))
            _ = v10
        case "":
            for fi11 := 0; fi11 < 3; fi11++ {
                if (!(255 > 539)) {
                    switch (fi9 % ((fi9) | 1)) {
                    case 717:
                        if ((1e10 != (1.5 * 1e10)) || false) {
                            var v12 bool = true
                            v4 -= v4
                            v1 = append(v1, "a")
                            _ = v12
                        } else {
                            v4 = 55
                            var v13 int8 = -20
                            v3 = append(v3, (fi9 % ((127) | 1)))
                            _ = v13
                        }
                    case 1:
                        v2 += (v2 / ((v2) | 1))
                    default:
                        v14 := GenS1{F0: 0, F1: 65535, F2: v2}
                        _ = v14
                    }
                    switch (fi9 + (fi11 << ((255) & 63))) {
                    case 32768:
                        var v15 int = fi9
                        _ = v15
                    case -164:
                        switch 991 {
                        case 42:
                            var v16 bool = false
                            _ = v16
                        default:
                            var v17 float32 = 0.0
                            v4 -= v4
                            _ = v17
                        }
                    }
                } else {
                    if ((v2 % ((77) | 1)) > v2) {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    } else {
                        var v18 uint32 = 42
                        _ = v18
                    }
                    if true {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    } else {
                        if len(v3) > 0 {
                            v3[0] = 857
                        }
                    }
                }
                _ = fi11
            }
        case "\n\t":
            v2 += v2
        }
        _ = fi9
    }
    var v19 string = "\n\t"
    switch v19 {
    case "":
        v3 = append(v3, 127)
    case "\n\t":
        v3 = append(v3, 0)
    }
    switch ("hello" + ((v19 + (v19 + ("" + "foo bar"))) + v19)) {
    case "foo bar":
        for fi20 := 0; fi20 < 17; fi20++ {
            var v21 uint32 = (v6 - 65535)
            _ = fi20
            _ = v21
        }
    default:
        var v22 uint8 = 240
        _ = v22
    }
    for fi23 := 0; fi23 < 2; fi23++ {
        for fi24 := 0; fi24 < 8; fi24++ {
            if false {
                v19 = ((v19 + (v19 + "foo bar")) + "\n\t")
            } else {
                v25 := GenS1{F0: 0, F1: 4294967295, F2: (v2 - v2)}
                var v26 uint8 = (^(v4))
                var v27 int16 = -1
                switch (fi24 << ((300) & 63)) {
                case 32768:
                    v19 = ""
                default:
                    fmt.Println(v4, v27, fi23)
                }
                _ = v25
                _ = v26
                _ = v27
            }
            fmt.Println(v4)
            v2 += (v2 * int8(fi24))
            var v28 int = (fi24 + fi23)
            if false {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                var v29 uint64 = 42
                _ = v29
            } else {
                var v30 uint32 = 229
                var v31 float32 = 0.1
                var v32 int8 = (v2 | (^(v2)))
                _ = v30
                _ = v31
                _ = v32
            }
            _ = fi24
            _ = v28
        }
        var v33 float64 = -2.25
        _ = fi23
        _ = v33
    }
    v2 -= v2
    var v34 string = ((v19 + v19) + "hello")
    if false {
        v34 = ((v19 + v34) + "\n\t")
        switch int(v2) {
        case -2147483648:
            v4 -= uint8(v6)
        case -431:
            for fi35 := 0; fi35 < 8; fi35++ {
                switch 32768 {
                case -9223372036854775808:
                    switch "hello" {
                    case "a":
                        fmt.Println(v4, v2)
                    default:
                        v36 := GenS1{F0: 0, F1: 42, F2: (v2 >> ((2) & 63))}
                        _ = v36
                    }
                default:
                    fmt.Println(v34, v34)
                }
                var v37 int64 = 74
                if true {
                    v4 -= (v4 ^ v4)
                    v34 = ("a" + (("hello" + v19) + (v34 + (v19 + "foo bar"))))
                } else {
                    var v38 int32 = -176
                    fmt.Println(v4)
                    _ = v38
                }
                _ = fi35
                _ = v37
            }
        case -774:
            fmt.Println(v2, v2, v19)
        default:
            switch -310 {
            case -634:
                var v39 float64 = (0.1 * 0.1)
                _ = v39
            case 741:
                v40 := GenS3{F0: -955, F1: (^(v4)), F2: 9223372036854775807}
                _ = v40
            case 1:
                v41 := GenS4{F0: 858, F1: 42, F2: v6}
                _ = v41
            default:
                switch v19 {
                case "foo bar":
                    var v42 int16 = 1000
                    _ = v42
                default:
                    v4 -= v4
                }
            }
        }
        var v43 int32 = -2147483648
        _ = v43
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
