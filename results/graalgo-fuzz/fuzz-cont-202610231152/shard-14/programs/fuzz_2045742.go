package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 int8 }

func main() {
    var v1 uint = 878
    var v2 string = "foo bar"
    var v3 string = ((((v2 + v2) + ("\n\t" + "hello")) + (("hello" + ((v2 + v2) + v2)) + "")) + ((((v2 + (v2 + v2)) + (v2 + (v2 + (v2 + v2)))) + "foo bar") + "a"))
    v4 := make(map[uint32]string)
    v1 += 42
    for fi5 := 0; fi5 < 8; fi5++ {
        fmt.Println(fi5, v3, v2)
        if false {
            switch v2 {
            case "hello":
                for fi6 := 0; fi6 < 6; fi6++ {
                    if (!(42 > 42)) {
                        var v7 uint16 = 845
                        _ = v7
                    }
                    var v8 uint8 = 43
                    fmt.Println(fi5)
                    v9 := GenS1{F0: -1, F1: v2, F2: 1}
                    _ = fi6
                    _ = v8
                    _ = v9
                }
            case "a":
                switch ("\n\t" + v2) {
                case "hello":
                    v2 = ("\n\t" + "\n\t")
                case "a":
                    v10 := GenS1{F0: -1, F1: (v2 + ""), F2: 69}
                    _ = v10
                default:
                    v11 := []int{4294967295}
                    _ = v11
                }
            default:
                v1 = v1
            }
            if (65535 > (v1 | v1)) {
                for fi12 := 0; fi12 < 19; fi12++ {
                    var v13 uint32 = uint32(v1)
                    _ = fi12
                    _ = v13
                }
                fmt.Println(fi5, v3, v1)
                v1 = 255
            }
            var v14 string = "a"
            _ = v14
        } else {
            fmt.Println(v1, v2)
        }
        _ = fi5
    }
    switch 4294967295 {
    case 32767:
        v15 := []int{}
        _ = v15
    }
    v16 := v4[2147483648]
    fmt.Println(v2, v16)
    for fi17 := 0; fi17 < 4; fi17++ {
        if false {
            for fi18 := 0; fi18 < 8; fi18++ {
                switch "" {
                case "hello":
                    if true {
                        var v19 uint = 505
                        v3 = (v16 + "hello")
                        _ = v19
                    }
                    if false {
                        var v20 int = fi17
                        var v21 bool = (2 != 65535)
                        v22 := GenS1{F0: 42, F1: (("\n\t" + v16) + ((v16 + "hello") + v16)), F2: -1}
                        _ = v20
                        _ = v21
                        _ = v22
                    }
                    v23 := v4[2]
                    _ = v23
                case "\n\t":
                    if ((!false) || false) {
                        v24 := GenS1{F0: -128, F1: v3, F2: 112}
                        _ = v24
                    }
                case "":
                    fmt.Println(v2)
                default:
                    if ((0 < 0) || true) {
                        v3 = "a"
                        v1 = v1
                        v1 = v1
                    }
                }
                fmt.Println(fi18, v16, v3)
                var v25 uint64 = 196
                _ = fi18
                _ = v25
            }
        } else {
            var v26 int = (fi17 | 32767)
            v27 := []int{}
            var v28 uint = (v1 % (((^(v1))) | 1))
            _ = v26
            _ = v27
            _ = v28
        }
        v29 := GenS1{F0: -123, F1: "\n\t", F2: 37}
        var v30 bool = (false || true)
        _ = fi17
        _ = v29
        _ = v30
    }
    switch 42 {
    case -893:
        if (148 <= 255) {
            if (!((3.14159 - 1e10) <= (float64((3.14159 * 0.1)) + 1e10))) {
                fmt.Println(v3, v1, v1)
                for fi31 := 0; fi31 < 1; fi31++ {
                    fmt.Println(v1, v16)
                    _ = fi31
                }
                var v32 int16 = 747
                _ = v32
            } else {
                fmt.Println(v16)
                _ = len(v4)
            }
            var v33 float64 = -2.25
            _ = v33
        } else {
            for fi34 := 0; fi34 < 4; fi34++ {
                if true {
                    for fi35 := 0; fi35 < 17; fi35++ {
                        v3 = "hello"
                        _ = fi35
                    }
                }
                var v36 int = (fi34 & fi34)
                var v37 float32 = 0.0
                _ = fi34
                _ = v36
                _ = v37
            }
            v4[2147483647] = ""
            v38 := []int{864, -590, 2147483647, 127}
            _ = v38
        }
    case 9223372036854775807:
        var v39 int8 = 1
        _ = v39
    case 128:
        fmt.Println(v1, v1)
    default:
        fmt.Println(v3)
    }
    v40 := GenS1{F0: 100, F1: v3, F2: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v40
}
