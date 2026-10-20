package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 int16; F1 int32; F2 uint16 }
type GenS3 struct { F0 uint64; F1 bool; F2 int16 }
type GenS4 struct { F0 uint64 }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]string)
    var v3 float32 = (0.1 + ((3.14159 + -2.25) - (((((1.5 * float32(1e10)) / 1.0) - (0.0 / 0.1)) * 3.14159) + 3.14159)))
    v4 := make(map[string]uint32)
    v5 := make(map[uint32]string)
    switch ("foo bar" + ("foo bar" + "")) {
    case "foo bar":
        v6 := GenS1{F0: 42, F1: 4294967295}
        _ = v6
    case "hello":
        v7 := GenS2{F0: 225, F1: -513, F2: 238}
        _ = v7
    default:
        if false {
            switch "\n\t" {
            case "\n\t":
                if (42 != 109) {
                    if (2147483648 != 700) {
                        v2[""] = "a"
                        for fi8 := 0; fi8 < 17; fi8++ {
                            var v9 int64 = -1
                            var v10 string = ("" + ("foo bar" + (("hello" + ("a" + ("a" + "foo bar"))) + ("\n\t" + "hello"))))
                            fmt.Println(fi8)
                            _ = fi8
                            _ = v10
                            _ = v9
                        }
                        v3 = v3
                    }
                    for fi11 := 0; fi11 < 5; fi11++ {
                        for fi12 := 0; fi12 < 6; fi12++ {
                            v13 := GenS1{F0: 0, F1: 0}
                            if true {
                                switch "" {
                                case "a":
                                    v14 := v5[2147483648]
                                    _ = v14
                                case "foo bar":
                                    delete(v1, ("a" + "a"))
                                case "":
                                    delete(v1, "foo bar")
                                default:
                                    v15 := GenS3{F0: 1, F1: false, F2: -965}
                                    _ = v15
                                }
                                var v16 float32 = 3.14159
                                _ = v16
                            } else {
                                v3 += ((v3 + v3) / -2.25)
                                var v17 int32 = 42
                                if (-1 == 380) {
                                    fmt.Println(fi11, fi12, fi11)
                                    switch "hello" {
                                    case "\n\t":
                                        v18 := []int{}
                                        _ = v18
                                    }
                                    v3 = v3
                                }
                                _ = v17
                            }
                            var v19 bool = (!true)
                            if v19 {
                                switch 128 {
                                case 4294967295:
                                    var v20 float32 = v3
                                    _ = v20
                                }
                                v3 -= (float32(3.14159) - (float32(0.1) + ((v3 - v3) / v3)))
                            } else {
                                if (!v19) {
                                    delete(v2, "hello")
                                    var v21 int8 = 122
                                    var v22 uint32 = 1
                                    _ = v21
                                    _ = v22
                                }
                                var v23 bool = v19
                                v19 = v19
                                _ = v23
                            }
                            v5[263] = "hello"
                            var v24 int16 = 42
                            v19 = (4294967295 == 0)
                            _ = fi12
                            _ = v13
                            _ = v19
                            _ = v24
                        }
                        _ = fi11
                    }
                } else {
                    v3 -= v3
                }
            }
            var v25 float64 = (3.14159 - 1e10)
            _ = v25
        }
    }
    switch "" {
    case "foo bar":
        var v26 int = 192
        _ = v26
    case "\n\t":
        for fi27 := 0; fi27 < 7; fi27++ {
            var v28 uint8 = 10
            fmt.Println(v3, fi27)
            _ = fi27
            _ = v28
        }
    }
    v3 -= v3
    var v29 int32 = -1
    for fi30 := 0; fi30 < 12; fi30++ {
        var v31 uint64 = 648
        _ = fi30
        _ = v31
    }
    switch 619 {
    case 2147483647:
        switch 4294967295 {
        case 2147483647:
            _ = len(v2)
        default:
            v32 := GenS4{F0: 1}
            _ = v32
        }
    case 161:
        if (0 > 314) {
            fmt.Println(v29)
        }
    case -704:
        if ((!true) && (42 != 65535)) {
            for fi33 := 0; fi33 < 8; fi33++ {
                if (1.5 > (1e10 + 1e10)) {
                    v34 := v2[""]
                    _ = v34
                }
                switch "" {
                case "foo bar":
                    v29 = (v29 & (v29 * v29))
                case "hello":
                    if (113 <= 65535) {
                        for fi35 := 0; fi35 < 13; fi35++ {
                            v3 = -2.25
                            _ = fi35
                        }
                        v36 := []int{}
                        _ = v36
                    }
                case "a":
                    v37 := GenS4{F0: 681}
                    _ = v37
                }
                _ = fi33
            }
            var v38 int8 = -1
            fmt.Println(v38)
            _ = v38
        } else {
            fmt.Println(v29, v3)
        }
    }
    v29 += (v29 & 372)
    for fi39 := 0; fi39 < 14; fi39++ {
        switch 808 {
        case 0:
            for fi40 := 0; fi40 < 1; fi40++ {
                v2["a"] = ("a" + "\n\t")
                _ = fi40
            }
        case 9223372036854775807:
            v29 = v29
        case 1:
            if false {
                for fi41 := 0; fi41 < 18; fi41++ {
                    v3 += v3
                    v29 -= (^(v29))
                    _ = fi41
                }
                v29 -= v29
            }
        default:
            v42 := GenS1{F0: -154, F1: 20}
            _ = v42
        }
        var v43 uint8 = 255
        v29 -= v29
        _ = fi39
        _ = v43
    }
    switch -2147483648 {
    case 128:
        v3 += -2.25
    case -1:
        v3 += v3
    case 4294967295:
        fmt.Println(v29, v3)
    default:
        if false {
            _ = len(v1)
        }
    }
    var v44 int32 = v29
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v44
    _ = v5
}
