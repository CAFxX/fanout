package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 int64 }
type GenS2 struct { F0 uint64; F1 int8; F2 uint32 }
type GenS3 struct { F0 uint64; F1 int32 }
type GenS4 struct { F0 int64; F1 uint32; F2 int32 }
type GenS5 struct { F0 uint32 }

func main() {
    v1 := make(map[int]bool)
    var v2 int64 = 0
    var v3 uint64 = 18446744073709551615
    var v4 uint64 = 837
    var v5 uint32 = 65535
    var v6 uint = 18446744073709551615
    _ = len(v1)
    v7 := v1[-506]
    v5 = 429
    switch "" {
    case "a":
        var v8 uint64 = uint64(v5)
        _ = v8
    case "\n\t":
        v7 = true
    case "":
        switch 32768 {
        case 105:
            switch "\n\t" {
            case "a":
                var v9 uint32 = (^(v5))
                _ = v9
            }
        case 128:
            v10 := []int{}
            _ = v10
        }
    }
    v11 := GenS1{F0: 100, F1: (!((0 > 1000) || true)), F2: (^(v2))}
    delete(v1, 127)
    for fi12 := 0; fi12 < 12; fi12++ {
        v1[fi12] = v7
        var v13 uint8 = 0
        v14 := GenS1{F0: 100, F1: v7, F2: v2}
        _ = fi12
        _ = v13
        _ = v14
    }
    var v15 float32 = 1.5
    if true {
        if true {
            if v7 {
                var v16 int8 = -128
                _ = v16
            }
        } else {
            var v17 uint16 = 270
            if (-2147483648 == 42) {
                var v18 float32 = v15
                if v7 {
                    v4 = 0
                    if false {
                        v7 = (v7 && v7)
                    } else {
                        switch "\n\t" {
                        case "\n\t":
                            fmt.Println(v6, v7)
                        case "foo bar":
                            v4 -= (^(v4))
                        }
                    }
                    var v19 float32 = 3.14159
                    _ = v19
                } else {
                    v7 = false
                }
                _ = v18
            }
            _ = v17
        }
    } else {
        if v7 {
            v20 := []int{-2147483648, -9223372036854775808}
            v5 += (v5 << ((v6) & 63))
            _ = v20
        }
    }
    var v21 int16 = -1
    switch "\n\t" {
    case "":
        v22 := []int{-66, -1}
        _ = v22
    case "\n\t":
        v3 += v4
    case "hello":
        switch ("foo bar" + ("hello" + ("\n\t" + "hello"))) {
        case "\n\t":
            switch "\n\t" {
            case "foo bar":
                delete(v1, -379)
            case "":
                var v23 uint16 = 166
                _ = v23
            case "hello":
                switch 32768 {
                case 9223372036854775807:
                    for fi24 := 0; fi24 < 9; fi24++ {
                        if (("hello" + "hello") != "a") {
                            v3 = 18446744073709551615
                            for fi25 := 0; fi25 < 5; fi25++ {
                                v26 := []int{(fi25 ^ fi25)}
                                var v27 int32 = 569
                                if (v7 || (216 < 445)) {
                                    var v28 string = "foo bar"
                                    v4 = v3
                                    if len(v26) > 0 {
                                        v26[0] = fi25
                                    }
                                    var v29 uint = v6
                                    _ = v28
                                    _ = v29
                                }
                                _ = fi25
                                _ = v26
                                _ = v27
                            }
                        } else {
                            v30 := GenS2{F0: (v3 / ((v3) | 1)), F1: int8(fi24), F2: 956}
                            _ = v30
                        }
                        v21 += (-(v21))
                        _ = fi24
                    }
                case 306:
                    v31 := GenS3{F0: v3, F1: 0}
                    _ = v31
                }
            default:
                v15 += v15
            }
        case "a":
            v2 = 1
        case "foo bar":
            fmt.Println(v3)
        default:
            for fi32 := 0; fi32 < 9; fi32++ {
                switch "" {
                case "a":
                    v4 = v3
                case "foo bar":
                    var v33 float64 = float64(v15)
                    _ = v33
                case "hello":
                    fmt.Println(v15, v6, v3)
                default:
                    fmt.Println(fi32, v7, v6)
                }
                switch (fi32 | fi32) {
                case -9223372036854775808:
                    v6 += v6
                case -1:
                    v34 := GenS4{F0: (v2 << (((v6 / (((v6 - uint(v3))) | 1))) & 63)), F1: (v5 | v5), F2: 42}
                    _ = v34
                }
                for fi35 := 0; fi35 < 4; fi35++ {
                    var v36 int16 = v21
                    v37 := GenS5{F0: uint32(v6)}
                    fmt.Println(fi32, fi32, fi32)
                    _ = fi35
                    _ = v36
                    _ = v37
                }
                _ = fi32
            }
        }
    }
    var v38 bool = (!v7)
    v2 = int64(v21)
    if true {
        var v39 bool = ((!("\n\t" < "a")) || false)
        _ = v39
    } else {
        if v7 {
            v40 := GenS1{F0: 0, F1: (127 >= uint8(v5)), F2: (v2 % ((-1) | 1))}
            if (580 != 855) {
                var v41 uint = 42
                _ = v41
            }
            v5 += (v5 | uint32(v3))
            _ = v40
        }
    }
    if (2147483647 <= -542) {
        fmt.Println(v38)
        var v42 uint = v6
        _ = v42
    } else {
        v43 := v1[32768]
        v6 = v6
        v44 := []int{}
        _ = v43
        _ = v44
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v15
    _ = v2
    _ = v21
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
