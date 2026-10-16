package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 uint16; F1 string; F2 uint32 }
type GenS3 struct { F0 int32; F1 bool }
type GenS4 struct { F0 int8; F1 int; F2 int32 }

func main() {
    var v1 int = 476
    var v2 uint32 = 65535
    var v3 int32 = 2147483647
    v2 = (^(v2))
    if false {
        fmt.Println(v1, v3)
        fmt.Println(v3, v1)
        v4 := []int{734, (v1 >> ((2) & 63)), v1}
        _ = v4
    }
    if false {
        fmt.Println(v3)
        v2 = (v2 / (((v2 % ((v2) | 1))) | 1))
    } else {
        v1 = 128
        v3 = (-(v3))
        fmt.Println(v3, v1, v3)
    }
    var v5 uint16 = 1
    v5 = 42
    fmt.Println(v1, v2)
    if true {
        if (2 <= 128) {
            if false {
                v6 := GenS1{F0: (v1 + (v1 ^ 2147483647)), F1: 42}
                var v7 uint8 = 93
                v8 := []int{4294967295}
                _ = v6
                _ = v7
                _ = v8
            } else {
                v9 := GenS2{F0: uint16(v2), F1: "a", F2: v2}
                if ((127 < 0) && (!false)) {
                    fmt.Println(v2, v3)
                    for fi10 := 0; fi10 < 1; fi10++ {
                        if ((false && true) && true) {
                            for fi11 := 0; fi11 < 8; fi11++ {
                                switch (("foo bar" + "foo bar") + "a") {
                                case "hello":
                                    v2 += (v2 % ((v2) | 1))
                                case "foo bar":
                                    if false {
                                        v12 := []int{0, (v1 << ((18446744073709551615) & 63)), int(v3)}
                                        var v13 uint32 = (v2 << ((18446744073709551615) & 63))
                                        if len(v12) >= 2 {
                                            v12 = v12[1:len(v12)]
                                        }
                                        _ = v12
                                        _ = v13
                                    } else {
                                        var v14 uint8 = 213
                                        v15 := GenS3{F0: v3, F1: true}
                                        _ = v14
                                        _ = v15
                                    }
                                }
                                if (false && (!((1.5 - 1.5) != -2.25))) {
                                    switch -9223372036854775808 {
                                    case 42:
                                        v16 := GenS1{F0: 45, F1: v3}
                                        _ = v16
                                    case -1:
                                        var v17 float64 = 0.1
                                        _ = v17
                                    }
                                    v3 = 0
                                    if (int8(v5) <= 1) {
                                        if ((v2 / ((v2) | 1)) != 4294967295) {
                                            fmt.Println(fi10)
                                        } else {
                                            var v18 int = (fi11 << ((65535) & 63))
                                            switch "hello" {
                                            case "foo bar":
                                                v19 := []int{v1, (^(fi10)), (fi11 + fi10), 1}
                                                _ = v19
                                            default:
                                                if (!(true || (true || ((fi11 / ((127) | 1)) >= (^(v18)))))) {
                                                    fmt.Println(v2)
                                                    v3 += (^(v3))
                                                    switch "a" {
                                                    case "\n\t":
                                                        v1 = (v18 ^ (fi11 - (fi11 / ((v1) | 1))))
                                                        if false {
                                                            if (193 >= 220) {
                                                                v20 := GenS2{F0: (v5 % (((^(v5))) | 1)), F1: "\n\t", F2: (v2 & 2147483648)}
                                                                var v21 string = "foo bar"
                                                                _ = v20
                                                                _ = v21
                                                            }
                                                        } else {
                                                            v22 := GenS2{F0: (^(v5)), F1: "", F2: v2}
                                                            if false {
                                                                var v23 uint64 = 42
                                                                switch fi10 {
                                                                case -1:
                                                                    v5 = v5
                                                                case -38:
                                                                    v1 += (fi11 & (^(v18)))
                                                                }
                                                                switch (fi10 + (fi11 + fi10)) {
                                                                case 127:
                                                                    var v24 uint = 897
                                                                    _ = v24
                                                                default:
                                                                    v25 := []int{460, (fi11 ^ fi11), fi11, (v18 + 2147483647)}
                                                                    _ = v25
                                                                }
                                                                _ = v23
                                                            }
                                                            v26 := []int{}
                                                            _ = v22
                                                            _ = v26
                                                        }
                                                        var v27 uint64 = uint64(v5)
                                                        _ = v27
                                                    case "foo bar":
                                                        var v28 int16 = 32767
                                                        _ = v28
                                                    default:
                                                        var v29 int16 = 1
                                                        _ = v29
                                                    }
                                                }
                                            }
                                            v30 := GenS2{F0: 0, F1: (("" + ("foo bar" + "foo bar")) + ""), F2: (v2 + v2)}
                                            _ = v18
                                            _ = v30
                                        }
                                        v2 = (v2 % ((v2) | 1))
                                    }
                                }
                                var v31 uint16 = v5
                                _ = fi11
                                _ = v31
                            }
                        }
                        for fi32 := 0; fi32 < 11; fi32++ {
                            var v33 uint8 = 0
                            var v34 uint16 = (^(v5))
                            _ = fi32
                            _ = v33
                            _ = v34
                        }
                        _ = fi10
                    }
                    switch -256 {
                    case 32767:
                        v35 := []int{-9223372036854775808, 4294967295, 32767}
                        _ = v35
                    }
                }
                _ = v9
            }
            v2 += 347
            v36 := []int{(v1 * v1), v1, -2147483648}
            _ = v36
        }
    }
    fmt.Println(v5, v3, v2)
    switch ("hello" + (("" + "\n\t") + "\n\t")) {
    case "hello":
        switch v1 {
        case 423:
            switch "" {
            case "foo bar":
                v37 := GenS4{F0: 1, F1: v1, F2: -492}
                _ = v37
            case "hello":
                v38 := []int{v1}
                _ = v38
            case "\n\t":
                v39 := []int{9223372036854775807}
                _ = v39
            default:
                v40 := []int{}
                _ = v40
            }
        case 4294967295:
            v2 = (v2 ^ (^(v2)))
        case 163:
            var v41 int = (v1 >> ((42) & 63))
            _ = v41
        default:
            var v42 uint8 = 215
            _ = v42
        }
    default:
        fmt.Println(v2, v2)
    }
    v43 := []int{v1, int(v5)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v43:", len(v43), v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v43
    _ = v5
}
