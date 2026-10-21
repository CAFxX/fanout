package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint }

func main() {
    v1 := make(map[int]bool)
    var v2 uint32 = 0
    var v3 uint16 = uint16(v2)
    var v4 int = -548
    var v5 uint32 = v2
    var v6 int64 = 9223372036854775807
    var v7 float32 = (0.1 * -2.25)
    v1[257] = false
    v4 = int(v3)
    v8 := []int{}
    if false {
        for fi9 := 0; fi9 < 15; fi9++ {
            v4 -= (fi9 + fi9)
            switch "foo bar" {
            case "\n\t":
                v10 := GenS1{F0: "", F1: -963}
                _ = v10
            }
            var v11 uint64 = 18446744073709551615
            _ = fi9
            _ = v11
        }
        var v12 uint32 = (v5 * (v2 & (^(v5))))
        for fi13 := 0; fi13 < 18; fi13++ {
            for fi14 := 0; fi14 < 20; fi14++ {
                if (133 == 127) {
                    var v15 float64 = 3.14159
                    _ = v15
                }
                _ = fi14
            }
            _ = fi13
        }
        _ = v12
    }
    if true {
        if false {
            _ = len(v1)
        } else {
            if (0.1 <= 0.1) {
                var v16 int16 = 1
                delete(v1, v4)
                switch v4 {
                case -770:
                    if true {
                        v17 := GenS2{F0: -1}
                        _ = v17
                    }
                case -545:
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                default:
                    v4 = (v4 % (((v4 & (v4 / (((v4 + v4)) | 1)))) | 1))
                }
                _ = v16
            }
        }
        var v18 int64 = (v6 + (v6 / ((42) | 1)))
        _ = v18
    }
    switch "foo bar" {
    case "\n\t":
        if true {
            v7 -= (0.0 + v7)
        } else {
            fmt.Println(v4, v6, v2)
        }
    case "foo bar":
        fmt.Println(v4, v7)
    case "hello":
        v19 := GenS1{F0: "", F1: (v4 << ((1) & 63))}
        _ = v19
    }
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    switch (("\n\t" + "\n\t") + (("foo bar" + "") + "foo bar")) {
    case "foo bar":
        fmt.Println(v2, v3)
    case "":
        var v20 uint64 = 596
        _ = v20
    }
    var v21 uint64 = 4294967295
    if true {
        switch "foo bar" {
        case "\n\t":
            fmt.Println(v5, v5)
        case "hello":
            for fi22 := 0; fi22 < 12; fi22++ {
                fmt.Println(v7, v2, v21)
                _ = fi22
            }
        default:
            var v23 uint64 = v21
            _ = v23
        }
        v5 = (v5 - 2147483648)
    } else {
        for fi24 := 0; fi24 < 2; fi24++ {
            v1[(v4 - (v4 / (((fi24 / ((v4) | 1))) | 1)))] = (-1 < -32768)
            _ = fi24
        }
        v6 = v6
    }
    v25 := GenS3{F0: 4294967295}
    v4 += 32767
    var v26 int8 = 127
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v21
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
