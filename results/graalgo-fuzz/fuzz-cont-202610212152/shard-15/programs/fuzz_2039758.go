package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int16; F1 int32; F2 uint }
type GenS5 struct { F0 string }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []string{("a" + "a"), "\n\t", "foo bar", ""}
    var v3 string = "\n\t"
    v4 := []int{42, 462, -236, 2147483647}
    v5 := make(map[uint32]string)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v3 = (v3 + v3)
    v3 = ((((v3 + (v3 + (v3 + (v3 + v3)))) + v3) + "hello") + "foo bar")
    var v6 uint16 = 88
    if true {
        for fi7 := 0; fi7 < 6; fi7++ {
            fmt.Println(fi7)
            if (!(255 >= 2)) {
                var v8 uint64 = 1
                v6 -= v6
                _ = v8
            } else {
                var v9 uint8 = 127
                fmt.Println(v9, fi7, v9)
                v9 = (v9 ^ v9)
                _ = v9
            }
            switch (v3 + "hello") {
            case "":
                v10 := GenS1{F0: (fi7 * 806), F1: 338}
                _ = v10
            case "\n\t":
                switch ((v3 + ("a" + v3)) + v3) {
                case "hello":
                    v5[42] = ((((v3 + "hello") + ("a" + "")) + v3) + ("hello" + v3))
                case "":
                    fmt.Println(fi7)
                case "a":
                    v2 = append(v2, "foo bar")
                default:
                    var v11 string = (v3 + (v3 + (v3 + v3)))
                    _ = v11
                }
            }
            _ = fi7
        }
        switch -629 {
        case -9223372036854775808:
            v4 = append(v4, 2147483647)
        default:
            fmt.Println(v3)
        }
    } else {
        _ = len(v4)
        for fi12 := 0; fi12 < 14; fi12++ {
            v13 := GenS2{F0: 1}
            v14 := GenS3{F0: (fi12 | fi12)}
            for fi15 := 0; fi15 < 6; fi15++ {
                switch v3 {
                case "\n\t":
                    v16 := GenS4{F0: 42, F1: 2147483647, F2: 108}
                    _ = v16
                case "a":
                    if ((fi15 & (fi15 & fi15)) > (fi15 % ((0) | 1))) {
                        v17 := GenS2{F0: 18446744073709551615}
                        var v18 uint32 = 398
                        var v19 int8 = 127
                        _ = v17
                        _ = v18
                        _ = v19
                    }
                }
                if false {
                    var v20 int64 = -348
                    v21 := GenS5{F0: "\n\t"}
                    _ = v20
                    _ = v21
                }
                var v22 uint32 = 1
                fmt.Println(v22, fi15, v3)
                _ = fi15
                _ = v22
            }
            _ = fi12
            _ = v13
            _ = v14
        }
        v6 = (v6 / ((255) | 1))
    }
    _ = len(v1)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v23 := GenS4{F0: 632, F1: 0, F2: 255}
    for fi24 := 0; fi24 < 13; fi24++ {
        v25 := GenS4{F0: 1, F1: 349, F2: 4294967295}
        var v26 int64 = 248
        switch "\n\t" {
        case "hello":
            if true {
                if ((!false) || true) {
                    var v27 float32 = 0.0
                    _ = v27
                }
            }
        case "":
            for fi28 := 0; fi28 < 12; fi28++ {
                v1[1] = 2147483648
                var v29 float64 = 1e10
                var v30 string = ""
                v2 = append(v2, v3)
                _ = fi28
                _ = v29
                _ = v30
            }
        case "foo bar":
            v26 += (-(v26))
        }
        _ = fi24
        _ = v25
        _ = v26
    }
    if false {
        switch -550 {
        case 128:
            var v31 uint = 2
            _ = v31
        default:
            v3 = (((("hello" + ("hello" + v3)) + v3) + v3) + ("" + v3))
        }
        v3 = (((v3 + v3) + v3) + v3)
        var v32 int = 42
        _ = v32
    } else {
        var v33 int16 = 42
        _ = v33
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
