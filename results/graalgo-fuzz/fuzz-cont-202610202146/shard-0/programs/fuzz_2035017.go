package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 string }
type GenS2 struct { F0 int16; F1 uint16 }

func main() {
    var v1 string = "foo bar"
    var v2 float32 = 1.5
    var v3 int8 = 42
    v4 := []int{32767, -2147483648, 912}
    switch v1 {
    case "\n\t":
        if true {
            for fi5 := 0; fi5 < 15; fi5++ {
                var v6 uint16 = 812
                var v7 int = fi5
                fmt.Println(v2, fi5, v2)
                _ = fi5
                _ = v6
                _ = v7
            }
            var v8 uint8 = 0
            v2 -= v2
            _ = v8
        } else {
            v4 = append(v4, 9223372036854775807)
            v9 := GenS1{F0: 697, F1: 42, F2: "\n\t"}
            _ = v9
        }
    }
    v2 = v2
    for fi10 := 0; fi10 < 3; fi10++ {
        var v11 uint64 = 42
        var v12 int16 = -1
        _ = fi10
        _ = v11
        _ = v12
    }
    v2 = (1.5 + (v2 / (v2 + v2)))
    v13 := GenS1{F0: 547, F1: 42, F2: (("\n\t" + "\n\t") + (("hello" + v1) + v1))}
    for fi14 := 0; fi14 < 20; fi14++ {
        v3 += v3
        var v15 uint8 = 0
        _ = fi14
        _ = v15
    }
    for fi16 := 0; fi16 < 4; fi16++ {
        for fi17 := 0; fi17 < 7; fi17++ {
            _ = len(v4)
            switch v1 {
            case "a":
                var v18 string = v1
                _ = v18
            case "\n\t":
                var v19 bool = true
                _ = v19
            case "":
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
            fmt.Println(fi16, fi17, v3)
            _ = fi17
        }
        v20 := GenS1{F0: 1, F1: 44, F2: v1}
        _ = fi16
        _ = v20
    }
    switch v1 {
    case "foo bar":
        for fi21 := 0; fi21 < 1; fi21++ {
            switch ((v1 + v1) + v1) {
            case "\n\t":
                v4 = append(v4, fi21)
            case "":
                fmt.Println(fi21)
            }
            _ = fi21
        }
    case "hello":
        fmt.Println(v2)
    default:
        v2 = float32((float64((v2 + (v2 + (v2 + float32(1.5))))) * 3.14159))
    }
    v1 = v1
    _ = len(v4)
    if (!(float64((v2 - v2)) <= ((0.1 + 3.14159) + (1.5 / -2.25)))) {
        switch ((v1 + (v1 + "a")) + "\n\t") {
        case "a":
            var v22 uint16 = 42
            _ = v22
        case "foo bar":
            var v23 float32 = (3.14159 / v2)
            _ = v23
        case "\n\t":
            _ = len(v4)
        }
    } else {
        v2 -= (v2 + v2)
        switch -2147483648 {
        case -952:
            if (-2.25 != -2.25) {
                if len(v4) > 0 {
                    v4[0] = -631
                }
            }
        case -2147483648:
            if (!(42 >= -32768)) {
                fmt.Println(v3)
                switch ((v1 + "\n\t") + v1) {
                case "hello":
                    var v24 int8 = (-(v3))
                    _ = v24
                case "\n\t":
                    if (32768 >= 42) {
                        v4 = append(v4, 0)
                        if (!false) {
                            fmt.Println(v1)
                            fmt.Println(v3, v3)
                        } else {
                            if (true && true) {
                                if (1 < 42) {
                                    var v25 bool = false
                                    var v26 int16 = -1
                                    _ = v25
                                    _ = v26
                                }
                                v2 -= v2
                                if (v3 != (v3 * (v3 + (v3 << ((255) & 63))))) {
                                    v4 = append(v4, 127)
                                }
                            }
                            var v27 uint = 4294967295
                            _ = v27
                        }
                    }
                case "foo bar":
                    if ((32767 < 32767) || false) {
                        var v28 int8 = v3
                        v4 = append(v4, 4294967295)
                        v28 += -128
                        _ = v28
                    }
                default:
                    for fi29 := 0; fi29 < 16; fi29++ {
                        v30 := GenS2{F0: 621, F1: 42}
                        _ = fi29
                        _ = v30
                    }
                }
            }
        default:
            if len(v4) > 0 {
                v4[0] = 127
            }
        }
        for fi31 := 0; fi31 < 15; fi31++ {
            v1 = ("" + (v1 + v1))
            _ = fi31
        }
    }
    var v32 string = ((v1 + "foo bar") + "a")
    var v33 int8 = (v3 * (v3 + -1))
    for fi34 := 0; fi34 < 13; fi34++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi34
    }
    var v35 uint = 18446744073709551615
    _ = len(v4)
    v1 = (v1 + (v32 + (v1 + ("a" + v1))))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v35
    _ = v4
}
