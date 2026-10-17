package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 uint }
type GenS2 struct { F0 uint8 }

func main() {
    v1 := []uint32{}
    var v2 int16 = 0
    v3 := []int{-1, 32768, 32768, 838}
    var v4 int8 = 100
    var v5 uint16 = 42
    v6 := make(map[string]string)
    for fi7 := 0; fi7 < 3; fi7++ {
        switch (fi7 + (fi7 << ((426) & 63))) {
        case 32768:
            var v8 float32 = float32((-2.25 / (1e10 + 0.0)))
            _ = v8
        case -2147483648:
            for fi9 := 0; fi9 < 5; fi9++ {
                var v10 string = ""
                _ = fi9
                _ = v10
            }
        default:
            var v11 int16 = v2
            _ = v11
        }
        delete(v6, "")
        var v12 uint16 = v5
        _ = fi7
        _ = v12
    }
    v4 -= v4
    var v13 int = -9223372036854775808
    _ = len(v6)
    switch "foo bar" {
    case "\n\t":
        _ = len(v3)
    default:
        v4 += v4
    }
    v5 = 281
    switch "hello" {
    case "a":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case "":
        var v14 uint16 = (v5 >> ((uint(v5)) & 63))
        _ = v14
    case "foo bar":
        v2 = int16(v5)
    }
    fmt.Println(v5, v4, v4)
    v5 = (v5 / ((32767) | 1))
    fmt.Println(v5)
    v13 += v13
    switch "a" {
    case "":
        var v15 int64 = 806
        _ = v15
    case "foo bar":
        v16 := GenS1{F0: (!(-9223372036854775808 == 0)), F1: "foo bar", F2: 4}
        _ = v16
    }
    for fi17 := 0; fi17 < 8; fi17++ {
        switch v13 {
        case -232:
            v18 := GenS1{F0: (!(!(618 > 4294967295))), F1: "a", F2: 4294967295}
            _ = v18
        case -2147483648:
            switch "hello" {
            case "a":
                v13 = (v13 % ((v13) | 1))
            default:
                v19 := GenS1{F0: ((0.1 * 0.1) < 0.1), F1: "a", F2: 18446744073709551615}
                _ = v19
            }
        case 32768:
            v2 += (v2 / ((v2) | 1))
        default:
            switch "a" {
            case "":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case "\n\t":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            default:
                switch 1 {
                case 9223372036854775807:
                    var v20 uint8 = 2
                    _ = v20
                case -220:
                    v21 := GenS1{F0: (true && false), F1: "\n\t", F2: 65535}
                    _ = v21
                case 127:
                    var v22 int = (fi17 | fi17)
                    _ = v22
                }
            }
        }
        _ = fi17
    }
    v13 -= 4294967295
    switch v13 {
    case 487:
        v23 := GenS1{F0: (true && (true || false)), F1: "\n\t", F2: 587}
        _ = v23
    case -297:
        if false {
            var v24 uint8 = 2
            if (true && true) {
                fmt.Println(v4)
            } else {
                v2 -= (v2 | int16(v4))
            }
            v25 := GenS2{F0: (^(v24))}
            _ = v24
            _ = v25
        } else {
            if true {
                var v26 int64 = 792
                _ = v26
            } else {
                v13 = (v13 / (((v13 >> ((0) & 63))) | 1))
            }
            fmt.Println(v4, v5, v5)
            for fi27 := 0; fi27 < 8; fi27++ {
                var v28 int16 = -1
                v4 = (v4 >> ((4294967295) & 63))
                _ = fi27
                _ = v28
            }
        }
    case 127:
        switch ("foo bar" + "") {
        case "hello":
            var v29 uint64 = 392
            _ = v29
        }
    default:
        fmt.Println(v13, v13, v4)
    }
    fmt.Println(v5, v13)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
