package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 int8; F1 int16; F2 uint32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint16; F1 int16 }
type GenS5 struct { F0 bool; F1 string }

func main() {
    var v1 uint32 = 1
    v2 := []uint32{(^(v1)), v1}
    v3 := make(map[int]uint32)
    var v4 uint16 = 0
    v1 -= (v1 / ((4294967295) | 1))
    if false {
        var v5 uint8 = 128
        _ = len(v2)
        fmt.Println(v4, v4)
        _ = v5
    }
    switch ("" + "") {
    case "":
        var v6 string = ("\n\t" + "foo bar")
        _ = v6
    }
    v2 = append(v2, v1)
    switch "foo bar" {
    case "a":
        if false {
            for fi7 := 0; fi7 < 15; fi7++ {
                fmt.Println(v1, v1, v1)
                _ = fi7
            }
            v8 := GenS1{F0: -2147483648, F1: 0}
            fmt.Println(v1, v4)
            _ = v8
        } else {
            v9 := GenS2{F0: 100, F1: 32767, F2: (v1 << ((527) & 63))}
            _ = v9
        }
    case "hello":
        _ = len(v3)
    default:
        v4 = (v4 ^ (v4 | (v4 & (v4 << ((1) & 63)))))
    }
    switch "a" {
    case "\n\t":
        if (!(false || true)) {
            switch 2147483647 {
            case 65:
                v10 := GenS3{F0: (255 > 1)}
                _ = v10
            case 2147483647:
                for fi11 := 0; fi11 < 2; fi11++ {
                    var v12 uint64 = 18446744073709551615
                    v13 := GenS3{F0: true}
                    v4 = v4
                    _ = fi11
                    _ = v12
                    _ = v13
                }
            case 127:
                var v14 float32 = -2.25
                _ = v14
            }
            _ = len(v2)
            switch "\n\t" {
            case "hello":
                _ = len(v2)
            default:
                v1 = v1
            }
        } else {
            for fi15 := 0; fi15 < 17; fi15++ {
                v4 += (v4 % ((v4) | 1))
                v16 := GenS3{F0: false}
                _ = fi15
                _ = v16
            }
            if len(v2) > 0 {
                v2[0] = v1
            }
        }
    case "foo bar":
        var v17 int64 = -9223372036854775808
        _ = v17
    default:
        for fi18 := 0; fi18 < 9; fi18++ {
            v19 := GenS1{F0: (fi18 * fi18), F1: 184}
            v20 := GenS4{F0: v4, F1: 0}
            _ = fi18
            _ = v19
            _ = v20
        }
    }
    fmt.Println(v4, v4)
    for fi21 := 0; fi21 < 19; fi21++ {
        switch "" {
        case "a":
            if false {
                v4 = 0
                if (-2147483648 == 1) {
                    v1 += v1
                    fmt.Println(fi21, v1)
                }
                v22 := GenS5{F0: (-9223372036854775808 < 32767), F1: ""}
                _ = v22
            } else {
                for fi23 := 0; fi23 < 13; fi23++ {
                    if true {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        if (1e10 >= 0.1) {
                            fmt.Println(fi21)
                            v4 = v4
                            v24 := GenS4{F0: (v4 - v4), F1: 42}
                            _ = v24
                        }
                        v4 -= v4
                    } else {
                        v25 := GenS4{F0: (v4 ^ 1), F1: 0}
                        _ = v25
                    }
                    v1 += 65535
                    _ = fi23
                }
                switch "foo bar" {
                case "hello":
                    var v26 int8 = 36
                    _ = v26
                case "":
                    var v27 uint64 = 4294967295
                    _ = v27
                default:
                    _ = len(v2)
                }
            }
        }
        _ = fi21
    }
    v3[2147483647] = (^(v1))
    _ = len(v2)
    fmt.Println(v1, v1)
    if true {
        var v28 int32 = -656
        for fi29 := 0; fi29 < 8; fi29++ {
            var v30 uint16 = (v4 % (((v4 ^ (v4 ^ (^(v4))))) | 1))
            _ = fi29
            _ = v30
        }
        var v31 uint64 = 2
        _ = v28
        _ = v31
    } else {
        fmt.Println(v1)
        if true {
            var v32 int = 1
            _ = v32
        } else {
            var v33 string = "hello"
            _ = v33
        }
    }
    if (-382 >= 1000) {
        switch "\n\t" {
        case "foo bar":
            v4 = 255
        default:
            v34 := GenS5{F0: (!((true || false) && ("\n\t" <= "hello"))), F1: ("a" + "")}
            _ = v34
        }
        var v35 int = 1
        _ = v35
    } else {
        v4 = (v4 & v4)
    }
    fmt.Println(v1, v1)
    var v36 uint64 = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
}
