package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 uint }
func genFn0(p0 int16, p1 uint8, p2 uint) int8 {
    v6 := GenS1{F0: 308, F1: "\n\t"}
    var v7 uint8 = (^(p1))
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return -128
}

func main() {
    var v1 string = ""
    var v2 uint32 = 2147483648
    var v3 uint32 = (v2 + (v2 | 878))
    var v4 uint32 = v3
    v5 := []uint32{v2}
    switch (v1 + v1) {
    case "foo bar":
        fmt.Println(v1, v4, v3)
    case "":
        var v8 uint64 = 2
        _ = v8
    default:
        _ = len(v5)
    }
    for fi9 := 0; fi9 < 4; fi9++ {
        var v10 float64 = 3.14159
        if (!true) {
            v1 = "a"
            for fi11 := 0; fi11 < 20; fi11++ {
                v12 := GenS1{F0: 0, F1: v1}
                _ = fi11
                _ = v12
            }
            fmt.Println(v4, fi9)
        }
        switch fi9 {
        case 9223372036854775807:
            var v13 int = fi9
            _ = v13
        case 809:
            for fi14 := 0; fi14 < 10; fi14++ {
                var v15 float32 = 1.5
                _ = len(v5)
                switch (fi9 % ((fi14) | 1)) {
                case 101:
                    v16 := GenS1{F0: 902, F1: "a"}
                    _ = v16
                case 173:
                    v5 = append(v5, (v4 ^ (v2 << ((1) & 63))))
                case -335:
                    v10 = float64(v15)
                }
                _ = fi14
                _ = v15
            }
        }
        _ = fi9
        _ = v10
    }
    v2 = v2
    for fi17 := 0; fi17 < 15; fi17++ {
        v1 = ("\n\t" + (v1 + ((v1 + (v1 + "foo bar")) + "foo bar")))
        switch (v1 + v1) {
        case "a":
            if ((!(true || ((1.5 >= float64(((1e10 + (0.1 * ((1.5 / (float32(1.5) / -2.25)) - float32((1e10 / 1e10))))) / ((-2.25 / 3.14159) / float32(((float64(0.0) + -2.25) / 1.5)))))) && (!(42 <= -385))))) || false) {
                fmt.Println(v1, v4, fi17)
                fmt.Println(v1, fi17)
                switch fi17 {
                case 32767:
                    switch "foo bar" {
                    case "foo bar":
                        var v18 float32 = 0.1
                        _ = v18
                    }
                case 1:
                    v2 += (v4 * v2)
                }
            }
        case "":
            for fi19 := 0; fi19 < 20; fi19++ {
                var v20 uint64 = 0
                switch v1 {
                case "hello":
                    fmt.Println(v1, v20, v3)
                case "a":
                    v2 += (v2 & 65535)
                case "foo bar":
                    switch fi19 {
                    case 127:
                        var v21 int32 = -1
                        _ = v21
                    case -1:
                        v4 += 2147483647
                    default:
                        var v22 uint8 = 89
                        _ = v22
                    }
                default:
                    switch "hello" {
                    case "a":
                        if len(v5) > 0 {
                            v5[0] = v2
                        }
                    case "":
                        v20 += (^(v20))
                    default:
                        var v23 uint8 = 0
                        _ = v23
                    }
                    if len(v5) > 0 {
                        v5[0] = (v2 / ((v2) | 1))
                    }
                }
                var v24 int = -9223372036854775808
                _ = fi19
                _ = v20
                _ = v24
            }
        case "foo bar":
            for fi25 := 0; fi25 < 15; fi25++ {
                var v26 int8 = genFn0(32767, 42, 0)
                if true {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    v3 -= 2147483647
                    v27 := GenS1{F0: 1, F1: v1}
                    v1 = v1
                    var v28 uint = 978
                    _ = v27
                    _ = v28
                }
                _ = fi25
                _ = v26
            }
        }
        v29 := GenS2{F0: 18446744073709551615}
        _ = fi17
        _ = v29
    }
    for fi30 := 0; fi30 < 8; fi30++ {
        for fi31 := 0; fi31 < 3; fi31++ {
            fmt.Println(fi30, v3, v1)
            switch fi30 {
            case 687:
                fmt.Println(v2)
            case 441:
                switch v1 {
                case "hello":
                    var v32 float32 = (1.5 + (1.5 * 0.0))
                    _ = v32
                }
            default:
                v3 = 0
                var v33 bool = true
                _ = v33
            }
            _ = fi31
        }
        var v34 uint = 18446744073709551615
        for fi35 := 0; fi35 < 7; fi35++ {
            _ = len(v5)
            fmt.Println(v4, v1, v1)
            _ = fi35
        }
        _ = fi30
        _ = v34
    }
    _ = len(v5)
    var v36 float64 = ((1e10 + (3.14159 - 0.0)) - ((1e10 * 1e10) + 1.5))
    var v37 int = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
