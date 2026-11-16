package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 int32; F1 uint16; F2 int8 }
type GenS3 struct { F0 int32; F1 uint16; F2 uint8 }
type GenS4 struct { F0 string; F1 int64 }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]int)
    v3 := []int{-1}
    var v4 uint = 18446744073709551615
    var v5 uint = (v4 - 790)
    var v6 bool = ((v4 - (v5 >> ((v5) & 63))) <= (v4 | (^(v4))))
    if v6 {
        fmt.Println(v4, v4, v4)
        v7 := GenS1{F0: 42, F1: 84}
        v4 += v4
        _ = v7
    } else {
        v8 := GenS1{F0: 255, F1: 89}
        v3 = append(v3, 42)
        _ = v8
    }
    v9 := GenS1{F0: 2, F1: 255}
    var v10 int8 = -1
    for fi11 := 0; fi11 < 18; fi11++ {
        if (((("a" + (("\n\t" + "") + "hello")) + "hello") + (("hello" + ("foo bar" + ("a" + ((("hello" + "foo bar") + "foo bar") + "a")))) + "a")) < "hello") {
            v5 -= v5
        }
        fmt.Println(v6, fi11, v5)
        _ = fi11
    }
    switch "" {
    case "hello":
        var v12 float64 = 3.14159
        _ = v12
    }
    if (-1 <= (v10 & (v10 & (v10 & v10)))) {
        var v13 int32 = 2147483647
        v14 := GenS1{F0: 2, F1: 1}
        v5 -= 4294967295
        _ = v13
        _ = v14
    }
    if v6 {
        if v6 {
            var v15 uint16 = 42
            var v16 uint16 = (v15 * v15)
            _ = v15
            _ = v16
        } else {
            if (485 < -1) {
                var v17 int = 4294967295
                _ = v17
            }
            switch -510 {
            case 32768:
                v5 = v4
            }
        }
    }
    switch "" {
    case "a":
        v18 := GenS1{F0: 179, F1: 30}
        _ = v18
    case "\n\t":
        if true {
            v4 = (v4 + v5)
            switch 551 {
            case 2147483647:
                v19 := GenS1{F0: 1, F1: 1}
                _ = v19
            case -344:
                for fi20 := 0; fi20 < 7; fi20++ {
                    fmt.Println(v6)
                    _ = fi20
                }
            case -994:
                v4 += 65535
            }
            v21 := GenS1{F0: 128, F1: 0}
            _ = v21
        }
    case "":
        v4 = (^(v5))
    default:
        fmt.Println(v5)
    }
    if (float32(0.1) <= (float32(3.14159) / 3.14159)) {
        if (!(("foo bar" + ("foo bar" + (("hello" + "foo bar") + "foo bar"))) >= "a")) {
            var v22 uint16 = 157
            var v23 int64 = 983
            fmt.Println(v22, v6, v6)
            _ = v22
            _ = v23
        }
        var v24 int32 = 0
        _ = v24
    }
    var v25 float32 = 3.14159
    v26 := GenS1{F0: 2, F1: 128}
    v27 := GenS2{F0: -2147483648, F1: 238, F2: (v10 / (((v10 - 42)) | 1))}
    fmt.Println(v25)
    var v28 int16 = 1000
    var v29 uint = 4294967295
    for fi30 := 0; fi30 < 4; fi30++ {
        switch (fi30 % ((fi30) | 1)) {
        case 9223372036854775807:
            v31 := GenS3{F0: 0, F1: 1, F2: 161}
            _ = v31
        default:
            for fi32 := 0; fi32 < 17; fi32++ {
                var v33 uint16 = 42
                if (!(0 != 4294967295)) {
                    v28 += -339
                    var v34 int16 = 32767
                    var v35 int = fi32
                    switch "foo bar" {
                    case "\n\t":
                        fmt.Println(v33, v5, v10)
                    default:
                        v36 := GenS4{F0: "", F1: 1}
                        _ = v36
                    }
                    _ = v34
                    _ = v35
                } else {
                    var v37 int32 = 2147483647
                    _ = v37
                }
                v25 = v25
                _ = fi32
                _ = v33
            }
        }
        _ = fi30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
