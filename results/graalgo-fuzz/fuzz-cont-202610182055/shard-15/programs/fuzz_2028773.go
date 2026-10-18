package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 bool }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 string; F1 uint; F2 int64 }

func main() {
    v1 := []int{}
    v2 := make(map[uint32]uint32)
    var v3 uint64 = 840
    var v4 int = 32767
    v4 = -2147483648
    v3 = 430
    v3 = (v3 % ((v3) | 1))
    var v5 uint64 = (^(v3))
    v5 = 0
    v5 -= v3
    v6 := GenS1{F0: true, F1: (v3 % (((v5 % ((0) | 1))) | 1)), F2: false}
    if (true || ((42 == 163) || (!true))) {
        v7 := GenS2{F0: 1}
        var v8 int8 = -128
        _ = v7
        _ = v8
    }
    v4 -= (-(v4))
    for fi9 := 0; fi9 < 3; fi9++ {
        v4 -= (^(v4))
        for fi10 := 0; fi10 < 20; fi10++ {
            v11 := GenS1{F0: (-1 <= -1), F1: 2, F2: true}
            switch "hello" {
            case "a":
                v5 = (^(v3))
            case "foo bar":
                switch (-(fi9)) {
                case -487:
                    var v12 float64 = ((0.0 - -2.25) / 1e10)
                    _ = v12
                default:
                    switch "a" {
                    case "a":
                        if true {
                            if ((65535 != 18) && true) {
                                v3 -= (v5 ^ v3)
                                v13 := v2[42]
                                _ = v13
                            }
                            var v14 int64 = -9223372036854775808
                            v14 = v14
                            _ = v14
                        }
                    case "\n\t":
                        v4 = (v4 + (fi10 - (v4 / ((-9223372036854775808) | 1))))
                    case "foo bar":
                        v15 := GenS3{F0: "\n\t", F1: 1, F2: -9223372036854775808}
                        _ = v15
                    default:
                        switch "a" {
                        case "a":
                            v4 -= -1
                        case "\n\t":
                            v4 += v4
                        case "":
                            var v16 int = fi10
                            fmt.Println(v5)
                            var v17 uint32 = 1
                            _ = v16
                            _ = v17
                        default:
                            var v18 int64 = 0
                            _ = v18
                        }
                        v3 += v3
                        v5 = (v5 - (^(v3)))
                        var v19 int64 = 0
                        v5 += (^(v5))
                        _ = v19
                    }
                }
            default:
                if false {
                    v3 -= 1
                    v5 += v3
                    var v20 int = 784
                    _ = v20
                }
            }
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    var v21 int64 = 0
    switch (-(v4)) {
    case -1:
        if false {
            var v22 int = v4
            for fi23 := 0; fi23 < 2; fi23++ {
                var v24 int16 = 980
                _ = fi23
                _ = v24
            }
            _ = v22
        } else {
            var v25 float64 = (3.14159 - (1.5 + 0.0))
            for fi26 := 0; fi26 < 13; fi26++ {
                fmt.Println(v25)
                var v27 uint16 = 23
                v27 += v27
                _ = fi26
                _ = v27
            }
            if false {
                var v28 int16 = 42
                fmt.Println(v3, v21, v28)
                v21 = v21
                _ = v28
            } else {
                v29 := GenS1{F0: (false || true), F1: 1, F2: false}
                _ = v29
            }
            _ = v25
        }
    case 971:
        v4 += (v4 & v4)
    }
    var v30 uint16 = 688
    var v31 int64 = (v21 >> ((65535) & 63))
    for fi32 := 0; fi32 < 16; fi32++ {
        var v33 int32 = 42
        v34 := GenS3{F0: "foo bar", F1: 18446744073709551615, F2: (v21 % ((556) | 1))}
        _ = fi32
        _ = v33
        _ = v34
    }
    if ("a" == ("foo bar" + "hello")) {
        var v35 bool = true
        fmt.Println(v5, v31)
        var v36 string = ""
        _ = v35
        _ = v36
    } else {
        fmt.Println(v4)
        v37 := GenS2{F0: -2147483648}
        _ = v37
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
