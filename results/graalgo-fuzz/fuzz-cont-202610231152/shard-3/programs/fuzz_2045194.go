package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 uint64; F1 uint; F2 int8 }
func genFn0(p0 int) string {
    var v6 uint16 = 1
    if false {
        v7 := []int{}
        _ = v7
    }
    _ = p0
    _ = v6
    return "foo bar"
}

func main() {
    var v1 float64 = -2.25
    var v2 bool = false
    var v3 int = 2147483647
    var v4 uint8 = 42
    var v5 uint8 = v4
    for fi8 := 0; fi8 < 13; fi8++ {
        switch 9223372036854775807 {
        case 4294967295:
            var v9 uint16 = 1
            _ = v9
        case -9223372036854775808:
            if false {
                switch fi8 {
                case 0:
                    v3 = (^(v3))
                case 551:
                    var v10 float32 = (((((0.0 / 1.0) * 1.5) + 1e10) * (0.1 * float32((-2.25 - v1)))) * 1.5)
                    _ = v10
                default:
                    for fi11 := 0; fi11 < 15; fi11++ {
                        var v12 int32 = 0
                        v4 = 42
                        _ = fi11
                        _ = v12
                    }
                }
                var v13 int8 = -128
                v14 := GenS1{F0: true, F1: (("hello" + "hello") + ("hello" + "foo bar"))}
                _ = v13
                _ = v14
            } else {
                v2 = (!(!true))
            }
        case 42:
            for fi15 := 0; fi15 < 7; fi15++ {
                var v16 uint = 887
                _ = fi15
                _ = v16
            }
        default:
            if (v2 || (v2 || (1 > 2147483648))) {
                v3 = v3
            } else {
                for fi17 := 0; fi17 < 10; fi17++ {
                    var v18 uint16 = 391
                    fmt.Println(v5)
                    _ = fi17
                    _ = v18
                }
                var v19 int8 = -1
                v5 -= v4
                _ = v19
            }
        }
        _ = fi8
    }
    for fi20 := 0; fi20 < 11; fi20++ {
        for fi21 := 0; fi21 < 2; fi21++ {
            v2 = (-97 >= 127)
            _ = fi21
        }
        fmt.Println(v1)
        _ = fi20
    }
    fmt.Println(v4, v3, v4)
    v22 := GenS2{F0: 18446744073709551615, F1: 18446744073709551615, F2: 127}
    v23 := []int{-2147483648, 636, v3}
    var v24 int32 = 0
    for fi25 := 0; fi25 < 13; fi25++ {
        for fi26 := 0; fi26 < 1; fi26++ {
            if len(v23) > 0 {
                v23[0] = fi26
            }
            if v2 {
                v27 := GenS1{F0: (1.5 == float64(-2.25)), F1: "a"}
                v2 = true
                v1 = float64((1.5 + (1.5 + float32(((v1 - -2.25) * ((v1 + 3.14159) / v1))))))
                _ = v27
            } else {
                if len(v23) >= 2 {
                    v23 = v23[1:len(v23)]
                }
                v24 -= (v24 * v24)
            }
            if true {
                if (4294967295 < 4294967295) {
                    if v2 {
                        fmt.Println(fi26)
                        v2 = ((v5 * 127) < 127)
                    }
                    var v28 uint8 = (^(v5))
                    v2 = true
                    _ = v28
                }
                fmt.Println(fi25, v2, fi25)
            }
            v4 = v4
            var v29 uint64 = 42
            var v30 float32 = (3.14159 - ((-2.25 + (-2.25 + 0.1)) - (0.1 + 0.1)))
            _ = fi26
            _ = v29
            _ = v30
        }
        for fi31 := 0; fi31 < 17; fi31++ {
            var v32 float64 = v1
            v32 = ((v32 - float64(0.1)) * v1)
            _ = fi31
            _ = v32
        }
        _ = fi25
    }
    if len(v23) >= 2 {
        v23 = v23[1:len(v23)]
    }
    v33 := GenS2{F0: 1, F1: 2, F2: -124}
    var v34 float32 = (1e10 / (0.1 - 0.0))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
