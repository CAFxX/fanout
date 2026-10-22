package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 bool }
type GenS2 struct { F0 uint32; F1 uint8; F2 uint64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 bool; F1 bool; F2 int }
type GenS5 struct { F0 uint32; F1 int32 }

func main() {
    var v1 string = "\n\t"
    var v2 float32 = float32(3.14159)
    var v3 int8 = 1
    v4 := make(map[int]string)
    v3 = (v3 / ((v3) | 1))
    delete(v4, -199)
    if true {
        v5 := []int{1}
        _ = v5
    } else {
        var v6 bool = true
        switch 991 {
        case -353:
            if (65535 >= 65535) {
                v2 -= float32(0.0)
                v7 := GenS1{F0: 2, F1: 0, F2: false}
                _ = v7
            }
        default:
            v1 = v1
        }
        v8 := GenS2{F0: 2, F1: 127, F2: 18446744073709551615}
        _ = v6
        _ = v8
    }
    v9 := GenS3{F0: 463}
    for fi10 := 0; fi10 < 4; fi10++ {
        for fi11 := 0; fi11 < 14; fi11++ {
            var v12 float64 = 1e10
            v13 := GenS4{F0: false, F1: true, F2: 128}
            _ = fi11
            _ = v12
            _ = v13
        }
        for fi14 := 0; fi14 < 5; fi14++ {
            switch "\n\t" {
            case "\n\t":
                fmt.Println(fi10)
            }
            switch v1 {
            case "hello":
                if (true && (((!true) || true) && (255 != 65535))) {
                    if true {
                        var v15 int = fi10
                        fmt.Println(fi10)
                        _ = v15
                    }
                    delete(v4, fi10)
                } else {
                    v2 = v2
                }
            case "":
                switch (-(fi10)) {
                case 32768:
                    switch 32767 {
                    case 32768:
                        var v16 int32 = 1
                        _ = v16
                    default:
                        fmt.Println(fi10)
                    }
                case -448:
                    v1 = (("" + "foo bar") + (("foo bar" + ((("" + (v1 + "a")) + v1) + v1)) + v1))
                case -1:
                    switch fi14 {
                    case 115:
                        v3 = (v3 << ((18446744073709551615) & 63))
                    }
                default:
                    var v17 int8 = v3
                    _ = v17
                }
            }
            switch "" {
            case "foo bar":
                v2 -= 0.0
            case "":
                v2 = (v2 + (v2 + (1e10 - (v2 * v2))))
                if true {
                    v2 = (1.5 / v2)
                    if true {
                        fmt.Println(v3, fi14)
                        if (false || false) {
                            v18 := []int{}
                            _ = v18
                        } else {
                            v3 = v3
                            v19 := GenS5{F0: 0, F1: 0}
                            if (110 == 255) {
                                v2 -= v2
                                var v20 int64 = 505
                                _ = v20
                            }
                            delete(v4, (fi14 >> ((205) & 63)))
                            _ = v19
                        }
                        v1 = v1
                        switch ((v1 + v1) + ((v1 + (v1 + v1)) + v1)) {
                        case "":
                            fmt.Println(fi10, fi10)
                        default:
                            switch v1 {
                            case "\n\t":
                                switch (fi10 << ((1) & 63)) {
                                case 9223372036854775807:
                                    var v21 uint16 = 0
                                    v22 := GenS5{F0: 673, F1: 42}
                                    _ = v21
                                    _ = v22
                                default:
                                    v1 = (v1 + v1)
                                }
                            }
                        }
                        if false {
                            var v23 uint32 = 378
                            var v24 int32 = 541
                            var v25 uint64 = 42
                            _ = v23
                            _ = v24
                            _ = v25
                        }
                    }
                }
                if true {
                    v2 = ((3.14159 + float32((0.0 - (3.14159 * ((0.1 / (-2.25 - 0.1)) - -2.25))))) - 0.1)
                    v26 := GenS2{F0: 42, F1: 2, F2: 42}
                    fmt.Println(fi14, v2)
                    delete(v4, fi10)
                    _ = v26
                }
            }
            _ = fi14
        }
        v3 -= (v3 - (^(v3)))
        _ = fi10
    }
    var v27 string = (v1 + ("\n\t" + ((((v1 + (v1 + ("" + (("a" + ((((v1 + ("" + v1)) + v1) + (v1 + "\n\t")) + "\n\t")) + v1)))) + (("hello" + "foo bar") + v1)) + (v1 + (((v1 + (v1 + ((v1 + v1) + ""))) + v1) + (v1 + (v1 + v1))))) + v1)))
    v28 := GenS3{F0: -32768}
    var v29 bool = (2 < 127)
    var v30 float64 = 1e10
    var v31 uint16 = 211
    var v32 float32 = v2
    for fi33 := 0; fi33 < 2; fi33++ {
        fmt.Println(v3, v31, v29)
        var v34 bool = (-208 >= -32768)
        delete(v4, fi33)
        _ = fi33
        _ = v34
    }
    v3 += 32
    v35 := GenS1{F0: 125, F1: v31, F2: (v29 || v29)}
    var v36 float32 = v32
    fmt.Println(v31)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v35
    _ = v36
    _ = v4
    _ = v9
}
