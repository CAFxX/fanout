package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 uint }
type GenS2 struct { F0 int; F1 uint32 }

func main() {
    var v1 int = 1
    var v2 int = v1
    v3 := make(map[string]uint32)
    var v4 uint32 = 4294967295
    v4 -= 1
    for fi5 := 0; fi5 < 11; fi5++ {
        var v6 int = int(v4)
        var v7 uint8 = 40
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 uint16 = 42
    for fi9 := 0; fi9 < 3; fi9++ {
        v8 -= v8
        v2 -= (fi9 >> ((4294967295) & 63))
        v1 = v2
        _ = fi9
    }
    if true {
        if true {
            if (42 != -1) {
                v2 = (v2 / (((-(v2))) | 1))
                for fi10 := 0; fi10 < 15; fi10++ {
                    v11 := []int{-1, (fi10 ^ int(v8))}
                    _ = fi10
                    _ = v11
                }
            }
            var v12 float64 = 0.0
            _ = v12
        }
        v13 := []int{v2, (v2 % ((-916) | 1)), (v2 >> ((42) & 63)), (v1 | 32767)}
        for fi14 := 0; fi14 < 13; fi14++ {
            if ((((3.14159 + 0.0) / (-2.25 + ((0.1 * -2.25) + (3.14159 + 0.0)))) * 3.14159) < ((1.5 - (3.14159 / (0.0 - float64(-2.25)))) - float64(float32(float64(0.0))))) {
                if (false && true) {
                    v15 := GenS1{F0: "foo bar", F1: -170, F2: uint(v8)}
                    _ = v15
                } else {
                    var v16 int8 = -128
                    if false {
                        for fi17 := 0; fi17 < 11; fi17++ {
                            var v18 uint = 2
                            fmt.Println(v2, v18, v1)
                            v19 := GenS2{F0: v1, F1: 2147483648}
                            v2 += fi17
                            fmt.Println(v16)
                            var v20 int = (-(v2))
                            _ = fi17
                            _ = v18
                            _ = v19
                            _ = v20
                        }
                        switch ("\n\t" + "foo bar") {
                        case "foo bar":
                            for fi21 := 0; fi21 < 18; fi21++ {
                                switch "hello" {
                                case "":
                                    switch "a" {
                                    case "\n\t":
                                        v4 += v4
                                    default:
                                        v22 := GenS2{F0: fi14, F1: v4}
                                        _ = v22
                                    }
                                    v2 = fi21
                                case "a":
                                    v23 := GenS1{F0: (("foo bar" + "\n\t") + "a"), F1: 269, F2: 1}
                                    var v24 int = (fi14 << ((255) & 63))
                                    _ = v23
                                    _ = v24
                                default:
                                    fmt.Println(v16)
                                }
                                _ = fi21
                            }
                        case "hello":
                            for fi25 := 0; fi25 < 4; fi25++ {
                                fmt.Println(v1, v8, fi14)
                                v26 := GenS1{F0: ("foo bar" + "\n\t"), F1: 42, F2: 0}
                                if true {
                                    var v27 float64 = float64(3.14159)
                                    _ = v27
                                } else {
                                    v8 = (v8 % (((^(v8))) | 1))
                                    _ = len(v3)
                                    var v28 uint16 = v8
                                    _ = v28
                                }
                                _ = fi25
                                _ = v26
                            }
                        }
                    }
                    _ = v16
                }
                v1 = (v1 / ((fi14) | 1))
                fmt.Println(v1, v4, v4)
            }
            v4 = v4
            _ = fi14
        }
        _ = v13
    }
    var v29 uint32 = 4294967295
    v30 := GenS1{F0: "\n\t", F1: 0, F2: 0}
    var v31 int32 = -2147483648
    var v32 string = "a"
    var v33 bool = (!(!(127 > 42)))
    for fi34 := 0; fi34 < 15; fi34++ {
        if (0 < 654) {
            var v35 bool = v33
            _ = v35
        }
        _ = fi34
    }
    var v36 uint8 = 233
    var v37 int32 = (^(v31))
    fmt.Println(v8)
    v33 = ((false && v33) || v33)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v36
    _ = v37
    _ = v4
    _ = v8
}
