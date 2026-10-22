package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 int8 }
type GenS2 struct { F0 int64 }
func genFn0() int {
    switch "hello" {
    case "\n\t":
        v5 := GenS1{F0: true, F1: 479, F2: -1}
        _ = v5
    }
    v6 := []int{}
    _ = v6
    return 127
}

func main() {
    var v1 uint = 0
    var v2 int16 = 1
    v3 := make(map[int]uint32)
    v4 := make(map[int]string)
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 uint8 = 10
        v9 := v3[fi7]
        var v10 int16 = v2
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    var v11 uint8 = 2
    switch 0 {
    case 127:
        var v12 float64 = ((0.1 / (((float64((1e10 / (1e10 - (-2.25 + ((((float32(float64((float32(float64((0.1 - 0.0))) + 0.1))) + float32(1.5)) * (3.14159 * float32(3.14159))) - ((-2.25 * (1.5 / -2.25)) + 0.0)) + (0.0 * (0.0 / float32(1e10)))))))) + (1.5 / 1e10)) - ((1.5 + 3.14159) / 1.5)) + -2.25)) * 3.14159)
        _ = v12
    }
    var v13 int = 42
    if (false || false) {
        switch ("a" + "hello") {
        case "hello":
            var v14 uint8 = (v11 + 199)
            _ = v14
        case "foo bar":
            for fi15 := 0; fi15 < 18; fi15++ {
                var v16 uint32 = 995
                _ = fi15
                _ = v16
            }
        default:
            fmt.Println(v11, v2, v11)
        }
        for fi17 := 0; fi17 < 9; fi17++ {
            fmt.Println(v13, fi17)
            _ = fi17
        }
        switch "a" {
        case "":
            for fi18 := 0; fi18 < 13; fi18++ {
                var v19 int = fi18
                var v20 string = ""
                for fi21 := 0; fi21 < 1; fi21++ {
                    v22 := GenS2{F0: 1}
                    _ = fi21
                    _ = v22
                }
                _ = fi18
                _ = v19
                _ = v20
            }
        case "foo bar":
            switch "hello" {
            case "hello":
                for fi23 := 0; fi23 < 15; fi23++ {
                    v24 := v3[(v13 / ((fi23) | 1))]
                    for fi25 := 0; fi25 < 12; fi25++ {
                        v24 -= 2147483648
                        _ = fi25
                    }
                    for fi26 := 0; fi26 < 1; fi26++ {
                        switch genFn0() {
                        case -865:
                            var v27 int64 = 9223372036854775807
                            _ = v27
                        case -570:
                            delete(v3, fi26)
                        case -1:
                            v24 = v24
                        }
                        var v28 string = (("foo bar" + "foo bar") + "foo bar")
                        _ = fi26
                        _ = v28
                    }
                    _ = fi23
                    _ = v24
                }
            case "foo bar":
                if ((!(v11 != 42)) && (0 >= 468)) {
                    v11 -= 42
                } else {
                    fmt.Println(v11, v13)
                    v11 -= v11
                }
            case "a":
                v3[genFn0()] = 1
            }
        }
    }
    for fi29 := 0; fi29 < 18; fi29++ {
        var v30 string = "\n\t"
        _ = fi29
        _ = v30
    }
    v13 += -9223372036854775808
    switch 32768 {
    case -1:
        v11 -= v11
    case -56:
        var v31 uint64 = 4294967295
        _ = v31
    case -2147483648:
        if (-9223372036854775808 >= -9223372036854775808) {
            v32 := v4[v13]
            v13 += genFn0()
            _ = v32
        }
    default:
        for fi33 := 0; fi33 < 11; fi33++ {
            var v34 int32 = int32(v1)
            v35 := GenS1{F0: ((-851 != fi33) && true), F1: 919, F2: 0}
            _ = fi33
            _ = v34
            _ = v35
        }
    }
    v2 += (v2 ^ v2)
    var v36 int16 = v2
    var v37 uint8 = 33
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
}
