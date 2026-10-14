package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 uint8 }
type GenS2 struct { F0 int16; F1 int16 }

func main() {
    v1 := []string{("" + "foo bar"), "hello"}
    v2 := []int{}
    var v3 int8 = 42
    var v4 int64 = -801
    v5 := GenS1{F0: true, F1: "\n\t", F2: 2}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v6 uint16 = 0
    var v7 int32 = 847
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if true {
        v7 -= v7
    } else {
        var v8 string = ("hello" + "foo bar")
        fmt.Println(v7)
        _ = v8
    }
    for fi9 := 0; fi9 < 17; fi9++ {
        if false {
            v10 := GenS2{F0: 32767, F1: 993}
            _ = v10
        } else {
            switch ("hello" + "hello") {
            case "\n\t":
                _ = len(v1)
            case "hello":
                switch 4294967295 {
                case 4294967295:
                    _ = len(v2)
                default:
                    v3 = v3
                }
            }
        }
        _ = fi9
    }
    var v11 uint = 978
    for fi12 := 0; fi12 < 17; fi12++ {
        switch (^(fi12)) {
        case -675:
            var v13 uint = 256
            _ = v13
        default:
            if (!true) {
                var v14 bool = (42 < 2)
                v4 += (-(v4))
                _ = v14
            }
        }
        if ((0.1 / 1.0) < 0.0) {
            v7 -= (v7 / (((v7 << ((736) & 63))) | 1))
        } else {
            fmt.Println(v3, v3)
        }
        v7 += v7
        _ = fi12
    }
    if true {
        switch 2147483647 {
        case 4294967295:
            if ("foo bar" == "\n\t") {
                switch "foo bar" {
                case "":
                    for fi15 := 0; fi15 < 4; fi15++ {
                        v11 += (v11 | 0)
                        var v16 uint8 = 128
                        var v17 int16 = 32767
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                case "\n\t":
                    var v18 float64 = (3.14159 * 3.14159)
                    _ = v18
                case "foo bar":
                    v4 = v4
                }
            }
        case -1:
            _ = len(v1)
        }
        v4 += (v4 & (v4 % (((v4 & (v4 / ((v4) | 1)))) | 1)))
        v7 += (v7 - v7)
    }
    v11 = 255
    var v19 int16 = -1
    v19 = v19
    var v20 float32 = 1e10
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
