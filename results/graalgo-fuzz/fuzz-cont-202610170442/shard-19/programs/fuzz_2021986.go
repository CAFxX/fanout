package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 int64 }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 bool; F1 uint8 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int16; F1 int8 }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint = 937
    v3 := make(map[uint32]uint32)
    v2 = (v2 ^ (v2 * 0))
    v2 += v2
    var v4 uint8 = 173
    fmt.Println(v2)
    if true {
        for fi5 := 0; fi5 < 13; fi5++ {
            v2 += (v2 | v2)
            _ = fi5
        }
        v2 -= v2
        v4 += v4
    }
    for fi6 := 0; fi6 < 2; fi6++ {
        v4 = 255
        var v7 uint = v2
        _ = fi6
        _ = v7
    }
    var v8 string = "hello"
    fmt.Println(v2, v8)
    var v9 int16 = 32767
    fmt.Println(v2)
    v10 := v1[1]
    var v11 int64 = 0
    var v12 uint32 = 42
    v13 := GenS1{F0: -70, F1: v10, F2: (v11 & -9223372036854775808)}
    switch (v10 + v10) {
    case "foo bar":
        for fi14 := 0; fi14 < 12; fi14++ {
            var v15 uint16 = 65535
            switch "a" {
            case "":
                var v16 int8 = 0
                _ = v16
            default:
                var v17 int16 = v9
                _ = v17
            }
            var v18 float32 = float32((1.5 + float64(1.5)))
            _ = fi14
            _ = v15
            _ = v18
        }
    case "hello":
        for fi19 := 0; fi19 < 13; fi19++ {
            var v20 int8 = 100
            if ((128 <= v4) || ((fi19 - 42) >= int(v4))) {
                switch "a" {
                case "foo bar":
                    v21 := []int{9223372036854775807, int(v11), (fi19 * (-(fi19)))}
                    _ = v21
                case "a":
                    v22 := GenS1{F0: (v20 & 27), F1: ("" + ("foo bar" + (("hello" + (v10 + v10)) + v8))), F2: 9223372036854775807}
                    _ = v22
                default:
                    var v23 string = v10
                    _ = v23
                }
            } else {
                switch 32767 {
                case 127:
                    if false {
                        v20 = 126
                        if true {
                            for fi24 := 0; fi24 < 10; fi24++ {
                                v12 -= 2147483648
                                delete(v1, v12)
                                _ = fi24
                            }
                            v25 := v3[v12]
                            switch "a" {
                            case "a":
                                switch v8 {
                                case "hello":
                                    _ = len(v1)
                                case "a":
                                    v26 := GenS2{F0: 0, F1: 344}
                                    _ = v26
                                default:
                                    if false {
                                        v25 -= v25
                                        v27 := v1[(v12 ^ v25)]
                                        v25 += v25
                                        _ = v27
                                    }
                                }
                            case "":
                                switch fi19 {
                                case 0:
                                    v9 -= v9
                                }
                            case "foo bar":
                                switch -2147483648 {
                                case 9223372036854775807:
                                    var v28 int16 = (v9 - int16(fi19))
                                    _ = v28
                                }
                            default:
                                if ((v11 & (v11 >> ((18446744073709551615) & 63))) < 576) {
                                    var v29 uint64 = 0
                                    for fi30 := 0; fi30 < 15; fi30++ {
                                        var v31 int64 = v11
                                        _ = fi30
                                        _ = v31
                                    }
                                    _ = v29
                                }
                            }
                            _ = v25
                        }
                    }
                case 814:
                    v20 -= v20
                default:
                    if (true || false) {
                        v8 = "hello"
                        v20 -= 1
                        switch fi19 {
                        case 1:
                            var v32 uint8 = (v4 | (v4 >> ((v2) & 63)))
                            _ = v32
                        case 0:
                            v33 := []int{}
                            _ = v33
                        }
                    } else {
                        v9 -= v9
                        v34 := GenS3{F0: 42}
                        v35 := GenS4{F0: false, F1: 0}
                        _ = v34
                        _ = v35
                    }
                }
                if ((fi19 <= fi19) || ((3.14159 < (((-2.25 / 1.0) * (1e10 * -2.25)) + (0.0 / -2.25))) || (("foo bar" + "") < ("hello" + "hello")))) {
                    v8 = "\n\t"
                } else {
                    v36 := GenS2{F0: 1, F1: 2}
                    for fi37 := 0; fi37 < 10; fi37++ {
                        fmt.Println(v11, v8)
                        v20 = (v20 ^ v20)
                        v38 := GenS5{F0: (v10 + (((((((("foo bar" + "foo bar") + "") + ("hello" + (v10 + (v10 + v8)))) + (v8 + "\n\t")) + "a") + v8) + ("\n\t" + (v8 + (((v8 + v8) + v10) + "\n\t")))) + ((((v8 + "hello") + ((("foo bar" + "\n\t") + ((v10 + (v8 + v8)) + v8)) + v10)) + v8) + v8)))}
                        var v39 int = 751
                        _ = fi37
                        _ = v38
                        _ = v39
                    }
                    _ = v36
                }
            }
            fmt.Println(v2)
            _ = fi19
            _ = v20
        }
    case "\n\t":
        if false {
            v8 = v10
        } else {
            v40 := GenS6{F0: v9, F1: -1}
            var v41 int8 = -128
            fmt.Println(v2, v12)
            _ = v40
            _ = v41
        }
    default:
        if (!true) {
            v42 := []int{9223372036854775807, 9223372036854775807, 128}
            _ = v42
        }
    }
    v10 = "foo bar"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
