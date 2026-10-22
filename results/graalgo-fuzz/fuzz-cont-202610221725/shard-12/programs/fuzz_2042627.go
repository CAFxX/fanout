package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8; F1 string; F2 string }

func main() {
    v1 := make(map[int]bool)
    v2 := []int{-656, -1, 42}
    var v3 uint32 = 888
    v4 := make(map[int]string)
    var v5 float32 = ((1.5 * 0.1) - ((1.5 * (-2.25 - (float32(float64(-2.25)) / 0.1))) * 0.0))
    v6 := make(map[int]int)
    for fi7 := 0; fi7 < 19; fi7++ {
        v5 += (v5 - v5)
        _ = fi7
    }
    for fi8 := 0; fi8 < 14; fi8++ {
        fmt.Println(v3, v5, v5)
        _ = fi8
    }
    if false {
        if true {
            fmt.Println(v5, v3)
        }
    }
    for fi9 := 0; fi9 < 2; fi9++ {
        v10 := GenS1{F0: fi9, F1: -58}
        _ = fi9
        _ = v10
    }
    v5 += v5
    for fi11 := 0; fi11 < 5; fi11++ {
        v12 := GenS2{F0: 0}
        if (91 == 100) {
            fmt.Println(v3, fi11)
        } else {
            if len(v2) > 0 {
                v2[0] = 1
            }
            _ = len(v2)
            switch "\n\t" {
            case "a":
                if false {
                    var v13 int16 = 263
                    _ = v13
                }
            case "":
                if (1 == 1000) {
                    fmt.Println(v3, fi11, v5)
                    var v14 uint64 = 42
                    _ = v14
                }
            }
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi11
        _ = v12
    }
    v15 := GenS1{F0: 2147483647, F1: -128}
    if (-2.25 != -2.25) {
        var v16 float32 = (v5 + float32(((0.0 * (-2.25 / float64(v5))) - 0.0)))
        _ = v16
    }
    fmt.Println(v5, v3)
    _ = len(v2)
    v3 = (v3 - 0)
    var v17 int32 = -108
    for fi18 := 0; fi18 < 2; fi18++ {
        switch fi18 {
        case 452:
            if ((8 == 255) || true) {
                v3 += (v3 % ((2) | 1))
            }
        default:
            switch (("\n\t" + "hello") + (((("hello" + "\n\t") + (("hello" + ("foo bar" + ((("hello" + "foo bar") + "") + ("" + ("" + "foo bar"))))) + (("\n\t" + "") + "hello"))) + "hello") + "hello")) {
            case "hello":
                switch (fi18 - -356) {
                case 128:
                    fmt.Println(v17)
                case 294:
                    var v19 int32 = -2147483648
                    _ = v19
                case 42:
                    var v20 uint = 65535
                    _ = v20
                }
            case "\n\t":
                v21 := GenS3{F0: 255, F1: ("hello" + "a"), F2: "\n\t"}
                _ = v21
            case "":
                fmt.Println(v5)
            }
        }
        v22 := GenS3{F0: 127, F1: (("foo bar" + "hello") + "foo bar"), F2: "\n\t"}
        fmt.Println(fi18)
        _ = fi18
        _ = v22
    }
    var v23 int8 = 1
    for fi24 := 0; fi24 < 16; fi24++ {
        if (!true) {
            if true {
                _ = len(v2)
            }
            v1[(fi24 - (fi24 | (fi24 * (fi24 << ((1) & 63)))))] = true
        } else {
            var v25 int = (fi24 | fi24)
            v23 += -128
            _ = v25
        }
        _ = fi24
    }
    v23 = v23
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v17
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
