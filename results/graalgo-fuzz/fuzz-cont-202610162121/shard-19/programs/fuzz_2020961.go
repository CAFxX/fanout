package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 int32; F1 uint8 }

func main() {
    var v1 float32 = float32(0.1)
    var v2 int = 0
    v3 := []string{("hello" + ("foo bar" + "a")), "a", ("foo bar" + "a")}
    v4 := []int{128}
    v5 := GenS1{F0: 441, F1: (-65 != 126)}
    switch v2 {
    case 127:
        if false {
            if (879 == 1) {
                switch v2 {
                case 0:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                case -696:
                    switch "hello" {
                    case "foo bar":
                        v1 = v1
                    case "\n\t":
                        v1 = 0.0
                    case "":
                        if len(v4) > 0 {
                            v4[0] = v2
                        }
                    default:
                        var v6 int16 = -32768
                        _ = v6
                    }
                }
                v1 = (float32(3.14159) * v1)
                v7 := GenS1{F0: 4294967295, F1: (18446744073709551615 != 42)}
                _ = v7
            } else {
                if true {
                    v8 := GenS1{F0: 533, F1: true}
                    if (1 == 1) {
                        fmt.Println(v1)
                        var v9 bool = (((v1 - v1) < v1) && (false || (0 <= 127)))
                        _ = v9
                    }
                    _ = v8
                }
            }
            v10 := GenS1{F0: 2147483647, F1: (!(4294967295 <= 702))}
            v11 := GenS2{F0: 42, F1: (("hello" + ("foo bar" + "")) + "")}
            _ = v10
            _ = v11
        } else {
            v1 = v1
            for fi12 := 0; fi12 < 18; fi12++ {
                for fi13 := 0; fi13 < 7; fi13++ {
                    switch ("a" + "hello") {
                    case "hello":
                        if (782 != 255) {
                            switch ("\n\t" + "hello") {
                            case "\n\t":
                                v1 -= (((1e10 - (((v1 + v1) * v1) - 0.1)) - float32(-2.25)) * ((v1 / 1.5) / v1))
                            case "foo bar":
                                v2 = fi12
                            case "a":
                                if len(v3) > 0 {
                                    v3[0] = (("a" + "a") + ("hello" + ("a" + ("\n\t" + ("\n\t" + ("hello" + "foo bar"))))))
                                }
                            default:
                                v2 -= 42
                            }
                            v1 = 3.14159
                            fmt.Println(fi13, v2, v2)
                        } else {
                            _ = len(v3)
                        }
                    }
                    var v14 float64 = -2.25
                    v14 -= (-2.25 + 3.14159)
                    _ = fi13
                    _ = v14
                }
                var v15 float32 = float32(((3.14159 / (1e10 / (0.1 + 0.1))) + 1e10))
                for fi16 := 0; fi16 < 10; fi16++ {
                    v17 := GenS2{F0: 1000, F1: (((("a" + "") + "a") + "hello") + "foo bar")}
                    var v18 uint32 = 2
                    v2 += fi12
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                _ = fi12
                _ = v15
            }
        }
    }
    var v19 int = (v2 | v2)
    v20 := GenS3{F0: 1, F1: 145}
    if len(v3) > 0 {
        v3[0] = "\n\t"
    }
    var v21 uint8 = 255
    if false {
        for fi22 := 0; fi22 < 15; fi22++ {
            v3 = append(v3, "hello")
            _ = fi22
        }
        switch (v2 | v19) {
        case 254:
            v21 = 1
        }
        var v23 float32 = v1
        _ = v23
    }
    for fi24 := 0; fi24 < 18; fi24++ {
        for fi25 := 0; fi25 < 18; fi25++ {
            var v26 string = "\n\t"
            v26 = ("" + v26)
            switch (v2 + v2) {
            case -913:
                if true {
                    if true {
                        fmt.Println(fi25)
                    }
                    fmt.Println(v1)
                }
            }
            v2 += v2
            _ = fi25
            _ = v26
        }
        _ = fi24
    }
    if (9223372036854775807 > 684) {
        var v27 float64 = 1e10
        var v28 uint8 = 128
        var v29 int16 = -1
        _ = v27
        _ = v28
        _ = v29
    }
    switch v19 {
    case -806:
        v1 -= v1
    default:
        var v30 uint32 = 42
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
