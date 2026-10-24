package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 string; F1 int8; F2 int }
type GenS3 struct { F0 bool; F1 uint8 }

func main() {
    var v1 uint16 = 315
    var v2 uint16 = (v1 | 32767)
    v3 := make(map[string]uint32)
    v4 := make(map[uint32]string)
    v5 := make(map[string]uint32)
    var v6 float64 = 0.1
    if false {
        v1 = (v2 * (v2 % ((v2) | 1)))
    } else {
        v2 -= v1
        _ = len(v3)
    }
    for fi7 := 0; fi7 < 12; fi7++ {
        var v8 uint8 = 0
        v8 = v8
        _ = fi7
        _ = v8
    }
    switch -12 {
    case -9223372036854775808:
        var v9 float64 = v6
        _ = v9
    }
    fmt.Println(v2, v6)
    v2 = (v1 | 1)
    var v10 int64 = -9223372036854775808
    for fi11 := 0; fi11 < 19; fi11++ {
        v2 = (v1 % ((v1) | 1))
        var v12 int32 = 42
        _ = fi11
        _ = v12
    }
    var v13 uint = 816
    switch ("foo bar" + "hello") {
    case "":
        if false {
            for fi14 := 0; fi14 < 11; fi14++ {
                var v15 int8 = -60
                fmt.Println(fi14)
                _ = fi14
                _ = v15
            }
            fmt.Println(v1, v1, v6)
        }
    case "foo bar":
        for fi16 := 0; fi16 < 16; fi16++ {
            switch ("" + "a") {
            case "hello":
                if false {
                    for fi17 := 0; fi17 < 18; fi17++ {
                        switch ("\n\t" + "a") {
                        case "foo bar":
                            var v18 uint32 = 2147483648
                            var v19 int8 = -128
                            _ = v18
                            _ = v19
                        case "hello":
                            fmt.Println(fi17, v1, v6)
                        default:
                            var v20 uint16 = v2
                            _ = v20
                        }
                        v13 += v13
                        if (1e10 == ((1e10 + -2.25) - 1e10)) {
                            if ((1.5 - float32(v6)) < (0.0 + 0.1)) {
                                v5["hello"] = 938
                            }
                            v2 = v2
                        }
                        _ = fi17
                    }
                    if (!true) {
                        var v21 uint64 = 0
                        _ = v21
                    } else {
                        v22 := GenS1{F0: 401, F1: 0}
                        v23 := v5["\n\t"]
                        _ = v22
                        _ = v23
                    }
                    for fi24 := 0; fi24 < 15; fi24++ {
                        v1 += v1
                        _ = fi24
                    }
                }
            case "\n\t":
                fmt.Println(v2, v6, fi16)
            case "":
                for fi25 := 0; fi25 < 19; fi25++ {
                    var v26 float64 = float64(((1.5 - (1e10 - (0.1 - (1e10 - (3.14159 * 3.14159))))) + 1.5))
                    _ = fi25
                    _ = v26
                }
            }
            _ = fi16
        }
    default:
        switch ("hello" + ("a" + "foo bar")) {
        case "hello":
            v1 = 65535
        case "a":
            for fi27 := 0; fi27 < 14; fi27++ {
                v28 := GenS2{F0: "foo bar", F1: 101, F2: 342}
                v29 := GenS1{F0: 18446744073709551615, F1: 42}
                v1 -= (v2 | (v2 + (v2 * v1)))
                _ = fi27
                _ = v28
                _ = v29
            }
        case "":
            for fi30 := 0; fi30 < 13; fi30++ {
                v31 := GenS3{F0: true, F1: 42}
                _ = fi30
                _ = v31
            }
        }
    }
    v13 += 65535
    if (42 != 0) {
        var v32 uint = (v13 % ((v13) | 1))
        v33 := []int{-425, -89}
        _ = v32
        _ = v33
    } else {
        v34 := GenS2{F0: ("foo bar" + ("\n\t" + (("" + "hello") + ""))), F1: 100, F2: -806}
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
