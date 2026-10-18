package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 string; F1 int8; F2 int16 }
func genFn0() uint8 {
    var v6 int32 = 203
    var v7 float32 = (((0.1 - 1e10) + -2.25) - ((-2.25 / 0.1) - (3.14159 / -2.25)))
    fmt.Println(v7)
    switch -198 {
    case 182:
        var v8 int64 = -9223372036854775808
        _ = v8
    case -2147483648:
        switch 128 {
        case 1:
            fmt.Println(v6, v6)
        case 42:
            if (!(false || (0 <= 2))) {
                var v9 uint16 = 32767
                var v10 uint = 718
                _ = v10
                _ = v9
            } else {
                v6 = -1
                fmt.Println(v6, v6)
                var v11 float64 = 0.0
                _ = v11
            }
        }
    }
    _ = v6
    _ = v7
    return 255
}

func main() {
    var v1 int32 = 1
    var v2 int16 = 487
    v3 := make(map[uint32]int)
    v4 := []string{((("\n\t" + ("hello" + ("foo bar" + "a"))) + "\n\t") + ("a" + (("" + ("" + "\n\t")) + (("a" + "hello") + "")))), "", "hello"}
    var v5 int16 = (v2 % (((v2 << ((1) & 63))) | 1))
    switch "\n\t" {
    case "foo bar":
        v3[2147483648] = 4294967295
    }
    for fi12 := 0; fi12 < 5; fi12++ {
        v5 += (v2 ^ 42)
        v13 := GenS1{F0: true, F1: (^(fi12))}
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 12; fi14++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi14
    }
    switch (("foo bar" + ("" + "hello")) + "hello") {
    case "hello":
        v2 = (v2 << ((65535) & 63))
    case "a":
        var v15 uint16 = 32767
        _ = v15
    case "":
        for fi16 := 0; fi16 < 5; fi16++ {
            for fi17 := 0; fi17 < 9; fi17++ {
                fmt.Println(v1, v5)
                v2 = -32768
                _ = fi17
            }
            _ = fi16
        }
    }
    for fi18 := 0; fi18 < 15; fi18++ {
        v2 += -32768
        if true {
            for fi19 := 0; fi19 < 2; fi19++ {
                v1 = v1
                fmt.Println(v5, v2)
                fmt.Println(v5, v1, fi18)
                _ = fi19
            }
            if (true && (0 > 4294967295)) {
                var v20 int = fi18
                for fi21 := 0; fi21 < 1; fi21++ {
                    v1 = int32(v5)
                    v20 += (v20 ^ (fi18 + (fi21 % (((-(fi21))) | 1))))
                    _ = fi21
                }
                _ = v20
            }
        }
        _ = fi18
    }
    for fi22 := 0; fi22 < 11; fi22++ {
        if (0.1 != -2.25) {
            v5 = v5
            switch "hello" {
            case "hello":
                fmt.Println(v2, v1)
            case "a":
                var v23 int64 = 9223372036854775807
                _ = v23
            case "\n\t":
                _ = len(v3)
            default:
                v1 = v1
            }
            if false {
                var v24 float32 = 1.5
                fmt.Println(v2)
                v25 := GenS1{F0: false, F1: fi22}
                _ = v24
                _ = v25
            } else {
                if true {
                    switch "\n\t" {
                    case "\n\t":
                        var v26 bool = (65535 != 174)
                        _ = v26
                    case "hello":
                        for fi27 := 0; fi27 < 19; fi27++ {
                            if false {
                                if ((true && (0.0 <= 0.1)) && false) {
                                    fmt.Println(v5)
                                    var v28 float32 = ((0.0 / 1e10) * 0.0)
                                    if true {
                                        v29 := GenS2{F0: "hello", F1: -128, F2: (v2 << ((4294967295) & 63))}
                                        _ = v29
                                    }
                                    _ = v28
                                }
                                fmt.Println(v1, v2)
                            } else {
                                v1 += v1
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                            _ = fi27
                        }
                    case "":
                        var v30 int16 = -642
                        _ = v30
                    }
                    v3[42] = fi22
                    v1 = -1
                }
            }
        } else {
            v4 = append(v4, "a")
        }
        delete(v3, 6)
        if len(v4) > 0 {
            v4[0] = "foo bar"
        }
        _ = fi22
    }
    _ = len(v3)
    v31 := GenS1{F0: false, F1: 32767}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
