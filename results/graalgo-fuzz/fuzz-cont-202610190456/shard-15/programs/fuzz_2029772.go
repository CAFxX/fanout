package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 int32 }
type GenS2 struct { F0 bool; F1 int16; F2 uint }
type GenS3 struct { F0 uint32; F1 int32; F2 uint8 }
type GenS4 struct { F0 int16 }
func genFn0(p0 uint8, p1 uint8, p2 string) bool {
    fmt.Println(p1, p0)
    v6 := GenS1{F0: 42, F1: 0, F2: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return (true || (!(true && true)))
}
func genFn1() int8 {
    var v7 float64 = -2.25
    var v8 int8 = 127
    _ = v7
    _ = v8
    return v8
}

func main() {
    var v1 float64 = 3.14159
    var v2 uint = 249
    v3 := make(map[uint32]uint32)
    var v4 bool = (2147483647 > 258)
    var v5 int32 = 2147483647
    var v9 int = 1
    v10 := GenS1{F0: (v9 * v9), F1: 32767, F2: -957}
    v11 := GenS1{F0: (v9 / (((v9 + 407)) | 1)), F1: 42, F2: (v5 << ((v2) & 63))}
    switch ("\n\t" + ("a" + "foo bar")) {
    case "":
        for fi12 := 0; fi12 < 10; fi12++ {
            var v13 uint32 = 4294967295
            _ = fi12
            _ = v13
        }
    default:
        fmt.Println(v9, v5)
    }
    switch v9 {
    case -774:
        v4 = v4
    default:
        switch 838 {
        case -9223372036854775808:
            fmt.Println(v9, v9)
        case 32768:
            v14 := GenS2{F0: v4, F1: 922, F2: (v2 ^ v2)}
            _ = v14
        case 845:
            fmt.Println(v5)
        }
    }
    var v15 bool = v4
    switch "\n\t" {
    case "\n\t":
        fmt.Println(v9, v5, v4)
    case "":
        v16 := []int{726, 949}
        _ = v16
    default:
        v17 := GenS3{F0: 0, F1: v5, F2: 1}
        _ = v17
    }
    for fi18 := 0; fi18 < 6; fi18++ {
        switch (("foo bar" + "") + "foo bar") {
        case "\n\t":
            fmt.Println(fi18, v5)
        }
        v19 := GenS2{F0: v15, F1: 0, F2: (v2 << ((412) & 63))}
        v5 = v5
        _ = fi18
        _ = v19
    }
    var v20 uint64 = 2
    fmt.Println(v1)
    if (!false) {
        v21 := GenS3{F0: 42, F1: (-(v5)), F2: 0}
        var v22 uint64 = (v20 % ((2) | 1))
        if (-1 >= int64(v9)) {
            fmt.Println(v5)
            v5 = (v5 | (v5 - int32(v9)))
        } else {
            switch v9 {
            case 745:
                fmt.Println(v9, v20, v20)
            case -791:
                v23 := GenS3{F0: 61, F1: (v5 & (v5 << (((v2 + v2)) & 63))), F2: 2}
                _ = v23
            case -319:
                switch "" {
                case "":
                    var v24 uint16 = 544
                    _ = v24
                default:
                    switch (("\n\t" + "hello") + "foo bar") {
                    case "\n\t":
                        v25 := GenS1{F0: 128, F1: 1, F2: -2147483648}
                        _ = v25
                    case "a":
                        if (v5 != int32(v22)) {
                            fmt.Println(v22)
                        }
                    case "hello":
                        v26 := []int{(v9 % ((v9) | 1)), v9, (v9 * (v9 + (v9 | v9))), 32768}
                        _ = v26
                    default:
                        v27 := GenS4{F0: 0}
                        _ = v27
                    }
                }
            }
            var v28 int32 = (v5 - v5)
            switch "" {
            case "a":
                _ = len(v3)
            case "\n\t":
                for fi29 := 0; fi29 < 14; fi29++ {
                    var v30 int32 = 0
                    _ = fi29
                    _ = v30
                }
            default:
                if (genFn1() <= -18) {
                    v31 := GenS1{F0: (v9 - (v9 << (((v2 % ((4294967295) | 1))) & 63))), F1: -788, F2: 42}
                    switch ("\n\t" + ("foo bar" + ("foo bar" + ""))) {
                    case "a":
                        fmt.Println(v9, v5)
                    case "":
                        v32 := []int{32768, 32767}
                        _ = v32
                    }
                    var v33 int16 = -863
                    _ = v31
                    _ = v33
                } else {
                    v34 := GenS1{F0: v9, F1: -568, F2: 0}
                    if v15 {
                        for fi35 := 0; fi35 < 2; fi35++ {
                            v36 := GenS4{F0: 1}
                            v28 = v28
                            for fi37 := 0; fi37 < 20; fi37++ {
                                fmt.Println(v20)
                                _ = fi37
                            }
                            _ = fi35
                            _ = v36
                        }
                    }
                    _ = v34
                }
            }
            _ = v28
        }
        _ = v21
        _ = v22
    }
    fmt.Println(v15, v9)
    switch ("hello" + ("\n\t" + "hello")) {
    case "foo bar":
        var v38 int16 = -305
        _ = v38
    case "":
        var v39 int32 = (v5 + (v5 ^ (v5 ^ v5)))
        _ = v39
    }
    switch "" {
    case "hello":
        _ = len(v3)
    case "":
        v5 -= (^(v5))
    case "foo bar":
        if (81 >= 959) {
            v40 := []int{(^(v9)), (v9 << ((255) & 63)), (v9 << (((v2 | (^(v2)))) & 63))}
            switch "\n\t" {
            case "foo bar":
                var v41 bool = (18446744073709551615 >= v2)
                _ = v41
            case "a":
                v42 := GenS2{F0: ((uint64(v2) < v20) && v15), F1: 1, F2: v2}
                _ = v42
            case "\n\t":
                v15 = v4
            default:
                v1 -= float64((0.0 * float32(((v1 * float64(-2.25)) / (v1 * v1)))))
            }
            _ = v40
        }
    }
    var v43 float64 = (3.14159 + v1)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v9
}
