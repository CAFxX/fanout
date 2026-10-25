package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8; F1 int64 }

func main() {
    v1 := make(map[string]string)
    var v2 string = ""
    v3 := make(map[int]int)
    v4 := make(map[string]bool)
    var v5 int = -660
    v6 := GenS1{F0: 68}
    var v7 bool = true
    fmt.Println(v2, v5, v5)
    v8 := GenS1{F0: 42}
    switch (v5 / (((-(v5))) | 1)) {
    case 32768:
        v9 := GenS1{F0: 65535}
        _ = v9
    default:
        v10 := GenS1{F0: 1}
        _ = v10
    }
    var v11 int16 = -32768
    for fi12 := 0; fi12 < 17; fi12++ {
        if v7 {
            v13 := GenS1{F0: 255}
            v7 = (1.5 < (float64((1e10 + 1.5)) - 1e10))
            fmt.Println(v5, v5)
            _ = v13
        } else {
            if v7 {
                v7 = (1 != 42)
                var v14 int16 = 1
                _ = v14
            } else {
                var v15 uint16 = 219
                var v16 int32 = 1
                v15 = v15
                _ = v15
                _ = v16
            }
            for fi17 := 0; fi17 < 7; fi17++ {
                switch 2147483647 {
                case -1:
                    v18 := []int{}
                    _ = v18
                case -2147483648:
                    v19 := []int{-9223372036854775808, v5, (^(v5)), fi17}
                    _ = v19
                }
                _ = fi17
            }
            v20 := GenS1{F0: 893}
            _ = v20
        }
        v21 := []int{}
        _ = fi12
        _ = v21
    }
    var v22 bool = ("" <= (((v2 + v2) + v2) + ("" + v2)))
    fmt.Println(v5, v5)
    if v7 {
        switch "hello" {
        case "hello":
            v23 := []int{}
            _ = v23
        case "foo bar":
            if (v22 && ((((3.14159 / 0.1) - -2.25) > 1e10) || (((0.1 - (1e10 / -2.25)) * (((((3.14159 / (((((1.5 + 1.5) + float32((((-2.25 + 0.1) + (float64(3.14159) * 0.0)) - 0.1))) + (-2.25 * 0.1)) * (((0.1 + 3.14159) - float32(0.1)) + 0.1)) + (0.1 / 3.14159))) + (1.5 * (3.14159 * 1e10))) / 0.1) * -2.25) + (0.1 / 1.0))) != (1e10 + 3.14159)))) {
                v5 = v5
                var v24 int64 = 42
                fmt.Println(v22, v2, v24)
                _ = v24
            }
        case "a":
            switch v5 {
            case 0:
                fmt.Println(v2, v2)
            case 42:
                var v25 int16 = v11
                _ = v25
            case 9223372036854775807:
                var v26 bool = v22
                _ = v26
            default:
                switch (v5 & v5) {
                case 0:
                    var v27 uint32 = 1
                    _ = v27
                case 32768:
                    switch "" {
                    case "hello":
                        var v28 uint32 = 2147483647
                        _ = v28
                    case "":
                        var v29 bool = true
                        _ = v29
                    case "\n\t":
                        var v30 float64 = float64(float32(-2.25))
                        _ = v30
                    }
                case 179:
                    if true {
                        if (!(!((v2 + v2) >= ((v2 + (v2 + ((((("a" + ("hello" + (("hello" + (v2 + (v2 + v2))) + v2))) + v2) + ((((v2 + v2) + "\n\t") + "foo bar") + v2)) + "a") + ("a" + v2)))) + v2)))) {
                            for fi31 := 0; fi31 < 8; fi31++ {
                                v32 := GenS1{F0: 0}
                                var v33 int32 = 29
                                _ = fi31
                                _ = v32
                                _ = v33
                            }
                            v34 := GenS2{F0: 151, F1: 0}
                            var v35 int16 = (-(v11))
                            _ = v34
                            _ = v35
                        }
                        for fi36 := 0; fi36 < 2; fi36++ {
                            if (123 != 196) {
                                var v37 int8 = -128
                                _ = v37
                            }
                            var v38 bool = false
                            _ = fi36
                            _ = v38
                        }
                    } else {
                        delete(v1, v2)
                        var v39 int8 = 127
                        _ = v39
                    }
                default:
                    v5 += -9223372036854775808
                }
            }
        }
        v22 = (32767 == 149)
        v7 = (431 <= -1)
    }
    if ((v11 << ((4294967295) & 63)) == -775) {
        _ = len(v3)
    }
    v7 = (1e10 < (((1.5 + 0.0) * -2.25) - (1.5 - 0.0)))
    v5 = v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
