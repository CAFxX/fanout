package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 string }
type GenS2 struct { F0 uint16; F1 int64; F2 int64 }

func main() {
    var v1 uint64 = 1
    var v2 float32 = 1.5
    var v3 uint64 = (v1 >> ((4294967295) & 63))
    var v4 string = ((("a" + "hello") + (("" + ("hello" + ((("" + "hello") + "foo bar") + ("a" + ("foo bar" + "hello"))))) + "hello")) + ("a" + ("foo bar" + ("" + ""))))
    v5 := make(map[uint32]uint32)
    v6 := make(map[string]int)
    fmt.Println(v1, v2, v2)
    if (128 >= 127) {
        var v7 int = -2147483648
        v7 -= (-(v7))
        var v8 uint = 255
        _ = v7
        _ = v8
    }
    v2 = ((-2.25 / v2) / 0.1)
    for fi9 := 0; fi9 < 15; fi9++ {
        v3 = (^(v1))
        var v10 float32 = (3.14159 / v2)
        fmt.Println(v10)
        _ = fi9
        _ = v10
    }
    if false {
        var v11 uint8 = 42
        for fi12 := 0; fi12 < 10; fi12++ {
            v13 := GenS1{F0: -52, F1: v4, F2: "a"}
            v6[v4] = (fi12 ^ fi12)
            var v14 uint32 = 42
            _ = fi12
            _ = v13
            _ = v14
        }
        if true {
            fmt.Println(v3, v4, v11)
            v3 = v1
        }
        _ = v11
    } else {
        if (false || ((v4 + v4) == v4)) {
            var v15 int = -240
            v16 := []int{v15, (^(v15)), (v15 - (v15 & 128)), (v15 & (v15 * (-(v15))))}
            v2 = float32(-2.25)
            _ = v15
            _ = v16
        } else {
            if false {
                for fi17 := 0; fi17 < 14; fi17++ {
                    v18 := GenS2{F0: 65535, F1: -9223372036854775808, F2: 42}
                    _ = fi17
                    _ = v18
                }
                v1 = v1
            } else {
                switch ("\n\t" + v4) {
                case "foo bar":
                    if (true && (false && true)) {
                        for fi19 := 0; fi19 < 1; fi19++ {
                            switch ((v4 + "") + v4) {
                            case "a":
                                for fi20 := 0; fi20 < 14; fi20++ {
                                    switch fi20 {
                                    case 32767:
                                        switch "a" {
                                        case "a":
                                            switch v4 {
                                            case "":
                                                v21 := []int{}
                                                _ = v21
                                            }
                                        default:
                                            v22 := GenS1{F0: 580, F1: v4, F2: (v4 + ("" + "a"))}
                                            _ = v22
                                        }
                                    case -9223372036854775808:
                                        var v23 int32 = 473
                                        _ = v23
                                    case -822:
                                        v3 = (v1 | (v3 + v3))
                                    default:
                                        if (!(2147483647 > 1)) {
                                            var v24 uint8 = 255
                                            _ = v24
                                        }
                                    }
                                    _ = fi20
                                }
                            case "\n\t":
                                var v25 uint8 = 185
                                _ = v25
                            }
                            _ = fi19
                        }
                    } else {
                        v4 = v4
                    }
                }
                fmt.Println(v3, v4, v1)
                for fi26 := 0; fi26 < 20; fi26++ {
                    if true {
                        var v27 uint32 = 42
                        _ = v27
                    } else {
                        var v28 int64 = 42
                        _ = v28
                    }
                    _ = fi26
                }
            }
            fmt.Println(v3)
            switch "foo bar" {
            case "a":
                var v29 uint64 = 120
                _ = v29
            default:
                v2 = v2
            }
        }
        switch 1 {
        case 32768:
            v30 := GenS2{F0: 42, F1: 42, F2: 1}
            _ = v30
        case 2147483647:
            v31 := GenS2{F0: 65535, F1: 42, F2: 9223372036854775807}
            _ = v31
        case 839:
            v4 = "hello"
        default:
            v3 = (v3 >> ((65535) & 63))
        }
    }
    var v32 uint16 = 65535
    v33 := GenS1{F0: -1, F1: ((v4 + ((((v4 + v4) + (("" + v4) + v4)) + v4) + ((v4 + v4) + v4))) + v4), F2: v4}
    fmt.Println(v32, v1, v32)
    var v34 int64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
