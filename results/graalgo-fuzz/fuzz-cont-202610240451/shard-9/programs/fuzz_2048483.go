package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 bool }

func main() {
    var v1 bool = false
    var v2 float64 = 0.1
    var v3 string = ("\n\t" + "\n\t")
    if false {
        var v4 int = 1
        var v5 int = v4
        _ = v4
        _ = v5
    } else {
        v2 = (3.14159 / 1e10)
        for fi6 := 0; fi6 < 18; fi6++ {
            for fi7 := 0; fi7 < 15; fi7++ {
                v2 = v2
                v8 := GenS1{F0: 193, F1: 0}
                _ = fi7
                _ = v8
            }
            var v9 uint32 = 42
            var v10 uint = 42
            _ = fi6
            _ = v10
            _ = v9
        }
        if v1 {
            v11 := []int{-663, 128, 2147483647, -2147483648}
            _ = v11
        }
    }
    for fi12 := 0; fi12 < 9; fi12++ {
        switch fi12 {
        case 127:
            v1 = ("" >= v3)
        }
        v2 += v2
        var v13 int16 = -1
        _ = fi12
        _ = v13
    }
    if v1 {
        v14 := []int{32768, 9223372036854775807}
        if v1 {
            switch (((((v3 + (v3 + ("foo bar" + (("hello" + "foo bar") + ("foo bar" + "\n\t"))))) + v3) + v3) + v3) + "hello") {
            case "hello":
                v15 := GenS1{F0: 846, F1: -32768}
                _ = v15
            default:
                switch "a" {
                case "\n\t":
                    switch (v3 + "foo bar") {
                    case "a":
                        for fi16 := 0; fi16 < 20; fi16++ {
                            if false {
                                v2 = float64(1e10)
                                if len(v14) > 0 {
                                    v14[0] = -182
                                }
                                v17 := GenS1{F0: 0, F1: -1}
                                _ = v17
                            } else {
                                var v18 uint8 = 127
                                _ = v18
                            }
                            if len(v14) > 0 {
                                v14[0] = (fi16 ^ fi16)
                            }
                            v19 := GenS1{F0: 0, F1: 1000}
                            _ = fi16
                            _ = v19
                        }
                    default:
                        fmt.Println(v2)
                    }
                case "foo bar":
                    v2 = (v2 - 0.1)
                case "a":
                    var v20 bool = (!true)
                    _ = v20
                }
            }
        }
        _ = v14
    } else {
        if (("hello" + v3) != "\n\t") {
            var v21 uint32 = 929
            _ = v21
        } else {
            if (2147483647 <= 2) {
                v2 = v2
                if false {
                    for fi22 := 0; fi22 < 9; fi22++ {
                        var v23 float64 = 1.5
                        _ = fi22
                        _ = v23
                    }
                    v24 := GenS1{F0: 0, F1: 898}
                    _ = v24
                } else {
                    fmt.Println(v2)
                }
            }
        }
    }
    fmt.Println(v3, v1, v3)
    for fi25 := 0; fi25 < 11; fi25++ {
        switch fi25 {
        case -1:
            fmt.Println(fi25)
        case 396:
            for fi26 := 0; fi26 < 15; fi26++ {
                var v27 float64 = (v2 * float64(((1.5 / 1.0) / float32((v2 / v2)))))
                var v28 uint64 = 1
                switch (fi26 + fi25) {
                case 128:
                    v29 := GenS2{F0: 1, F1: 86}
                    _ = v29
                default:
                    v2 = float64((((3.14159 + ((3.14159 + float32(0.1)) + ((-2.25 * 0.0) - 0.0))) / 0.1) - ((float32(0.0) + 0.0) * (3.14159 / (0.1 + float32(1e10))))))
                }
                _ = fi26
                _ = v27
                _ = v28
            }
        case 42:
            v1 = (!(4294967295 <= 2))
        default:
            var v30 float64 = v2
            _ = v30
        }
        var v31 int = fi25
        _ = fi25
        _ = v31
    }
    v32 := []int{-2147483648, -605, -9223372036854775808, 1}
    v33 := GenS3{F0: ((-409 > 0) && v1)}
    _ = len(v32)
    switch ((v3 + (("hello" + v3) + ("\n\t" + ""))) + v3) {
    case "hello":
        var v34 int64 = -226
        _ = v34
    default:
        v35 := GenS3{F0: true}
        _ = v35
    }
    v1 = v1
    if len(v32) >= 2 {
        v32 = v32[1:len(v32)]
    }
    v2 -= float64((float32(1.5) / 1.0))
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
}
