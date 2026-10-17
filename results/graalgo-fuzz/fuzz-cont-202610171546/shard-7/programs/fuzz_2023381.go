package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 string }
type GenS2 struct { F0 uint64; F1 int32; F2 uint16 }
type GenS3 struct { F0 int16; F1 uint32 }
type GenS4 struct { F0 int }

func main() {
    var v1 bool = true
    var v2 int8 = 100
    var v3 int64 = 42
    v4 := make(map[string]string)
    var v5 int32 = 24
    var v6 string = ("\n\t" + "")
    v7 := GenS1{F0: v5, F1: 53, F2: v6}
    v8 := []int{}
    var v9 float64 = 1e10
    v1 = v1
    var v10 uint8 = 178
    if (("" + "\n\t") > v6) {
        _ = len(v4)
        if v1 {
            fmt.Println(v9, v9)
            v11 := GenS2{F0: 290, F1: v5, F2: 510}
            var v12 int16 = -1
            _ = v11
            _ = v12
        }
    } else {
        switch v6 {
        case "hello":
            var v13 int = -2147483648
            _ = v13
        case "foo bar":
            for fi14 := 0; fi14 < 15; fi14++ {
                v15 := GenS2{F0: 402, F1: v5, F2: 989}
                _ = fi14
                _ = v15
            }
        }
        var v16 uint = 2
        for fi17 := 0; fi17 < 2; fi17++ {
            v1 = true
            if (65535 > 599) {
                switch ((v6 + "foo bar") + v6) {
                case "a":
                    if (42 != -934) {
                        v6 = ("a" + "hello")
                        v10 += 0
                    } else {
                        v1 = (2147483648 == 1)
                    }
                case "foo bar":
                    if len(v8) > 0 {
                        v8[0] = 42
                    }
                case "\n\t":
                    fmt.Println(v3, v2)
                }
                v6 = v6
                var v18 string = v6
                _ = v18
            } else {
                fmt.Println(v1, v5, v16)
            }
            if (v10 >= (v10 / (((v10 - 1)) | 1))) {
                if v1 {
                    var v19 uint16 = 0
                    _ = v19
                } else {
                    v16 += (v16 - v16)
                    switch fi17 {
                    case 128:
                        fmt.Println(v10)
                    case -606:
                        v20 := GenS3{F0: 42, F1: 674}
                        _ = v20
                    default:
                        var v21 int = (fi17 | 127)
                        _ = v21
                    }
                    if (v1 || (!v1)) {
                        var v22 int32 = (^(v5))
                        if ((fi17 >> ((278) & 63)) != (fi17 + (fi17 % ((-911) | 1)))) {
                            v23 := GenS4{F0: 234}
                            fmt.Println(v3, fi17)
                            _ = v23
                        }
                        switch 128 {
                        case 0:
                            var v24 bool = v1
                            _ = v24
                        case 344:
                            v2 += 100
                        case 32768:
                            for fi25 := 0; fi25 < 2; fi25++ {
                                v9 -= float64(3.14159)
                                _ = fi25
                            }
                        }
                        _ = v22
                    } else {
                        for fi26 := 0; fi26 < 9; fi26++ {
                            v16 -= (v16 * 18446744073709551615)
                            var v27 float32 = float32(float64(3.14159))
                            if len(v8) >= 2 {
                                v8 = v8[1:len(v8)]
                            }
                            _ = fi26
                            _ = v27
                        }
                    }
                }
                fmt.Println(v9, v1)
            }
            _ = fi17
        }
        _ = v16
    }
    var v28 int16 = 42
    var v29 int32 = -1
    switch 2147483647 {
    case 760:
        switch "a" {
        case "hello":
            var v30 int64 = 42
            _ = v30
        default:
            fmt.Println(v2, v1, v10)
        }
    case -522:
        fmt.Println(v1, v1, v9)
    case -1:
        v10 = 2
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
