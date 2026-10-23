package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 string }
type GenS2 struct { F0 uint64; F1 int }
func genFn0(p0 string) int64 {
    if (!(0 != 939)) {
        for fi5 := 0; fi5 < 8; fi5++ {
            var v6 float32 = 1e10
            _ = fi5
            _ = v6
        }
    } else {
        if (false || true) {
            if (!(false || true)) {
                v7 := []int{9223372036854775807}
                var v8 uint16 = 255
                _ = v7
                _ = v8
            }
        }
        var v9 uint64 = 369
        _ = v9
    }
    if (false && true) {
        var v10 int16 = 299
        for fi11 := 0; fi11 < 2; fi11++ {
            if true {
                if (42 >= 2147483647) {
                    for fi12 := 0; fi12 < 12; fi12++ {
                        v13 := []int{128}
                        var v14 float32 = ((float32((0.1 / 0.1)) - float32(float64((3.14159 * (3.14159 / (1.5 - (0.0 - float32(1.5)))))))) * ((0.1 * ((float32(0.0) + ((0.1 * (((3.14159 / 0.1) - float32((-2.25 / float64(float32(1.5))))) - float32(0.1))) * (1e10 + float32(1e10)))) - 1.5)) * (-2.25 / 1.0)))
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                    var v15 float32 = -2.25
                    p0 = ("a" + ("" + p0))
                    _ = v15
                }
            } else {
                v16 := GenS1{F0: 577, F1: 18446744073709551615, F2: p0}
                for fi17 := 0; fi17 < 5; fi17++ {
                    v18 := []int{(fi11 ^ -2147483648), (fi11 * fi17), fi17}
                    if ((fi11 >> ((48) & 63)) != fi11) {
                        _ = len(v18)
                        v10 = v10
                    }
                    switch 32768 {
                    case -1:
                        switch 135 {
                        case 32767:
                            v10 = 0
                        default:
                            fmt.Println(p0, fi17)
                        }
                    case 0:
                        var v19 bool = true
                        _ = v19
                    case 42:
                        v20 := GenS1{F0: 2, F1: 18446744073709551615, F2: (p0 + p0)}
                        _ = v20
                    default:
                        switch fi17 {
                        case 128:
                            var v21 int8 = -128
                            _ = v21
                        case 0:
                            var v22 uint = 255
                            _ = v22
                        case 42:
                            fmt.Println(v10, fi11, v10)
                        default:
                            var v23 float32 = (1e10 + 0.1)
                            _ = v23
                        }
                    }
                    _ = fi17
                    _ = v18
                }
                _ = v16
            }
            _ = fi11
        }
        for fi24 := 0; fi24 < 4; fi24++ {
            v25 := GenS1{F0: 18446744073709551615, F1: 42, F2: "\n\t"}
            var v26 uint8 = 1
            _ = fi24
            _ = v25
            _ = v26
        }
        _ = v10
    } else {
        if true {
            p0 = (p0 + p0)
        }
        switch 32767 {
        case 478:
            switch (p0 + (p0 + p0)) {
            case "a":
                for fi27 := 0; fi27 < 14; fi27++ {
                    var v28 uint16 = 296
                    _ = fi27
                    _ = v28
                }
            case "":
                fmt.Println(p0)
            default:
                var v29 int16 = -577
                _ = v29
            }
        case 1:
            v30 := GenS1{F0: 779, F1: 2, F2: p0}
            _ = v30
        case 42:
            p0 = (p0 + ("\n\t" + (p0 + ((("hello" + "hello") + ("a" + ((("hello" + "a") + ("hello" + p0)) + (((((p0 + p0) + "hello") + (((p0 + "\n\t") + p0) + (p0 + "\n\t"))) + p0) + "")))) + "\n\t"))))
        }
    }
    _ = p0
    return 9223372036854775807
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]string)
    var v3 float32 = (-2.25 - 1.5)
    v4 := make(map[string]bool)
    var v31 bool = true
    v31 = v31
    v31 = v31
    v3 += float32(float64(float32(1.5)))
    fmt.Println(v31, v3)
    if v31 {
        v32 := GenS2{F0: 4294967295, F1: -752}
        if (v31 || (v31 && (0 <= 0))) {
            if (32767 > 65535) {
                switch "\n\t" {
                case "hello":
                    v3 -= v3
                }
            }
            v31 = (!v31)
        }
        _ = v32
    } else {
        v31 = v31
        v33 := GenS2{F0: 2, F1: 32767}
        _ = v33
    }
    v3 -= (1e10 * v3)
    if (0 <= 647) {
        v34 := GenS1{F0: 1, F1: 420, F2: ""}
        v31 = (!true)
        _ = v34
    }
    var v35 bool = ((731 > 2147483648) || v31)
    v35 = false
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v35
    _ = v4
}
