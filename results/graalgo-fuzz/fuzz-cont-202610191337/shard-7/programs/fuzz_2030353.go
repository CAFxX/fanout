package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 uint }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int16; F1 uint8 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint64 }
func genFn0(p0 int32) int {
    v5 := []int{128, 32767, -201, 32767}
    var v6 bool = false
    _ = p0
    _ = v5
    _ = v6
    return 42
}

func main() {
    var v1 float64 = 1e10
    var v2 float64 = v1
    v3 := []int{924, 2147483647, -2147483648}
    v4 := make(map[uint32]string)
    for fi7 := 0; fi7 < 11; fi7++ {
        v8 := GenS1{F0: 18446744073709551615}
        _ = fi7
        _ = v8
    }
    v9 := GenS2{F0: -711, F1: 25}
    v1 = 3.14159
    switch 42 {
    case 42:
        var v10 int32 = 2147483647
        _ = v10
    case -491:
        v11 := GenS2{F0: -212, F1: 42}
        _ = v11
    default:
        for fi12 := 0; fi12 < 14; fi12++ {
            var v13 int16 = -32768
            if false {
                fmt.Println(v1, v13, v1)
                _ = len(v3)
            } else {
                v13 -= v13
                v14 := GenS3{F0: 0}
                _ = v14
            }
            var v15 int8 = -78
            _ = fi12
            _ = v13
            _ = v15
        }
    }
    if true {
        if false {
            var v16 float32 = (1e10 + -2.25)
            _ = v16
        } else {
            if (!true) {
                if (387 >= 0) {
                    for fi17 := 0; fi17 < 2; fi17++ {
                        v1 = 0.0
                        v18 := GenS4{F0: 72, F1: 67}
                        delete(v4, 1)
                        _ = fi17
                        _ = v18
                    }
                    if false {
                        fmt.Println(v2)
                        if (!true) {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            v2 -= v1
                            v19 := v4[42]
                            _ = v19
                        }
                    } else {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                } else {
                    v20 := GenS1{F0: 5}
                    _ = v20
                }
            } else {
                switch 32768 {
                case 32768:
                    fmt.Println(v1, v2)
                case -994:
                    switch 0 {
                    case 4294967295:
                        var v21 float64 = (v1 - v1)
                        _ = v21
                    case -1:
                        fmt.Println(v2)
                    case 392:
                        v22 := GenS5{F0: 2147483647}
                        _ = v22
                    }
                }
            }
        }
        switch -1 {
        case 4294967295:
            var v23 uint32 = 1
            _ = v23
        case -9223372036854775808:
            v2 += v2
        case 1:
            for fi24 := 0; fi24 < 13; fi24++ {
                switch "a" {
                case "foo bar":
                    for fi25 := 0; fi25 < 13; fi25++ {
                        var v26 int = genFn0(202)
                        _ = fi25
                        _ = v26
                    }
                case "hello":
                    for fi27 := 0; fi27 < 7; fi27++ {
                        switch "foo bar" {
                        case "foo bar":
                            v1 -= v1
                        case "a":
                            var v28 string = "hello"
                            _ = v28
                        default:
                            v1 = v2
                            v2 = ((float64(float32(3.14159)) + (v1 * v1)) + float64(1.5))
                            _ = len(v3)
                        }
                        var v29 uint8 = 1
                        var v30 uint = 245
                        _ = fi27
                        _ = v29
                        _ = v30
                    }
                default:
                    switch (((("" + "a") + "hello") + "foo bar") + "") {
                    case "\n\t":
                        var v31 float64 = (((0.0 * (v2 * (-2.25 + v2))) - v1) - 0.1)
                        _ = v31
                    default:
                        var v32 int = 4294967295
                        _ = v32
                    }
                }
                if len(v3) > 0 {
                    v3[0] = 9223372036854775807
                }
                if (1 <= 65535) {
                    fmt.Println(v1, v2, fi24)
                    if false {
                        v2 = float64(0.1)
                        _ = len(v4)
                    }
                } else {
                    for fi33 := 0; fi33 < 2; fi33++ {
                        _ = len(v3)
                        v34 := GenS5{F0: 2147483647}
                        _ = fi33
                        _ = v34
                    }
                    v2 -= ((float64((3.14159 + (-2.25 + (3.14159 / (0.1 / 1.0))))) - ((0.1 - -2.25) / v1)) * float64((((1.5 / 1e10) - 3.14159) / (1e10 / 1.5))))
                }
                _ = fi24
            }
        default:
            switch ("a" + "a") {
            case "hello":
                var v35 uint8 = 255
                _ = v35
            case "a":
                v36 := GenS5{F0: 840}
                _ = v36
            default:
                var v37 uint32 = 753
                _ = v37
            }
        }
    }
    if (!false) {
        var v38 int16 = 42
        v39 := GenS6{F0: 567}
        _ = v38
        _ = v39
    } else {
        var v40 float64 = ((-2.25 + v2) / v2)
        _ = v40
    }
    var v41 uint = 261
    if ((!true) || false) {
        v41 = 42
    }
    var v42 uint32 = 345
    v1 = (v2 / v1)
    var v43 float32 = 0.0
    var v44 int8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v9
}
