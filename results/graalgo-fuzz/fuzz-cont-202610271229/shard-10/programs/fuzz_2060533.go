package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32; F1 int64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int16; F1 uint16; F2 bool }
type GenS7 struct { F0 uint64 }

func main() {
    v1 := []int{-619, 212}
    v2 := make(map[uint32]int)
    v3 := []uint32{0}
    var v4 int8 = -1
    v5 := []string{((("hello" + (("foo bar" + (("hello" + "hello") + ("foo bar" + ("foo bar" + "")))) + "a")) + "hello") + "\n\t"), "a", (("\n\t" + "\n\t") + "\n\t"), (("hello" + ("" + "hello")) + "hello")}
    if (false && true) {
        var v6 int = 128
        v4 += (-(v4))
        _ = v6
    } else {
        for fi7 := 0; fi7 < 4; fi7++ {
            v2[885] = (fi7 % ((int(v4)) | 1))
            v8 := GenS1{F0: "\n\t"}
            _ = fi7
            _ = v8
        }
        for fi9 := 0; fi9 < 11; fi9++ {
            switch ("a" + "\n\t") {
            case "foo bar":
                var v10 uint = 2
                _ = v10
            case "":
                fmt.Println(v4, v4)
            case "hello":
                var v11 uint32 = 412
                _ = v11
            default:
                v12 := GenS2{F0: 1}
                _ = v12
            }
            v4 = v4
            _ = fi9
        }
    }
    v4 -= v4
    switch 42 {
    case 4294967295:
        switch "hello" {
        case "a":
            fmt.Println(v4)
        }
    case 451:
        v13 := GenS3{F0: 2147483647, F1: -1}
        _ = v13
    case -116:
        var v14 uint = 255
        _ = v14
    }
    v15 := GenS4{F0: 127}
    v4 = (v4 + (^(v4)))
    v4 -= -128
    fmt.Println(v4)
    for fi16 := 0; fi16 < 7; fi16++ {
        if true {
            if (!true) {
                fmt.Println(fi16, v4)
            }
        } else {
            v4 += v4
        }
        v4 = v4
        _ = fi16
    }
    if ((!true) || true) {
        v4 -= v4
        if ((float64(3.14159) - 1e10) > ((0.1 * (0.1 / 1.0)) - (3.14159 * (0.0 / 1.5)))) {
            v17 := GenS2{F0: 0}
            var v18 uint64 = 0
            fmt.Println(v4, v4)
            _ = v17
            _ = v18
        } else {
            v4 = (^(v4))
            for fi19 := 0; fi19 < 9; fi19++ {
                for fi20 := 0; fi20 < 10; fi20++ {
                    if false {
                        v21 := GenS5{F0: "\n\t"}
                        _ = v21
                    }
                    _ = len(v2)
                    v22 := GenS6{F0: -1, F1: 640, F2: false}
                    _ = fi20
                    _ = v22
                }
                _ = fi19
            }
            var v23 string = "foo bar"
            _ = v23
        }
    } else {
        var v24 int = 32767
        v4 -= v4
        _ = v24
    }
    v25 := v2[904]
    switch "\n\t" {
    case "\n\t":
        v25 -= 9223372036854775807
    case "a":
        v3 = append(v3, 876)
    }
    switch "foo bar" {
    case "\n\t":
        switch (-(v25)) {
        case 127:
            if len(v3) > 0 {
                v3[0] = 4294967295
            }
        case -105:
            v26 := GenS1{F0: (("a" + "hello") + "")}
            _ = v26
        case 1:
            var v27 float32 = 0.0
            _ = v27
        }
    case "foo bar":
        switch "foo bar" {
        case "a":
            v28 := GenS6{F0: -32768, F1: 412, F2: (((1 > 1) && false) && (212 != 1))}
            _ = v28
        case "":
            for fi29 := 0; fi29 < 8; fi29++ {
                v30 := v2[65535]
                v31 := GenS7{F0: 0}
                _ = fi29
                _ = v30
                _ = v31
            }
        case "hello":
            if true {
                var v32 uint8 = 255
                _ = v32
            }
        }
    case "a":
        v4 += (-(v4))
    default:
        if len(v5) > 0 {
            v5[0] = ((((("foo bar" + "\n\t") + "hello") + "hello") + "hello") + "a")
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
