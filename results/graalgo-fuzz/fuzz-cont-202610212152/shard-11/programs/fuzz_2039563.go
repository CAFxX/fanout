package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32; F1 uint8 }
func genFn0() int32 {
    v5 := GenS1{F0: -2147483648}
    var v6 string = "hello"
    var v7 bool = true
    if false {
        var v8 int8 = 101
        _ = v8
    } else {
        v9 := []int{496, 0, -423, 128}
        switch 2147483647 {
        case -2147483648:
            v10 := GenS2{F0: 321, F1: 1}
            _ = v10
        case 32768:
            for fi11 := 0; fi11 < 5; fi11++ {
                switch fi11 {
                case -977:
                    fmt.Println(v7, v6)
                case -1:
                    switch (v6 + (v6 + v6)) {
                    case "a":
                        v7 = false
                    case "foo bar":
                        for fi12 := 0; fi12 < 2; fi12++ {
                            v6 = v6
                            _ = fi12
                        }
                    case "\n\t":
                        if v7 {
                            v13 := GenS2{F0: 2, F1: 238}
                            _ = len(v9)
                            for fi14 := 0; fi14 < 15; fi14++ {
                                v15 := GenS1{F0: 717}
                                v9 = append(v9, (fi11 | -556))
                                _ = fi14
                                _ = v15
                            }
                            _ = v13
                        } else {
                            var v16 uint16 = 305
                            _ = v16
                        }
                    default:
                        for fi17 := 0; fi17 < 13; fi17++ {
                            switch v6 {
                            case "\n\t":
                                switch ("foo bar" + "a") {
                                case "a":
                                    var v18 string = v6
                                    _ = v18
                                default:
                                    v9 = append(v9, fi17)
                                }
                            case "foo bar":
                                v6 = ((v6 + "") + v6)
                                var v19 int32 = 165
                                fmt.Println(fi17, v7)
                                v7 = (0 > 396)
                                _ = v19
                            case "a":
                                var v20 bool = ((((!false) && (!false)) || ((v6 + "") >= v6)) && false)
                                _ = v20
                            }
                            _ = fi17
                        }
                    }
                }
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
                for fi21 := 0; fi21 < 11; fi21++ {
                    var v22 int32 = 932
                    var v23 uint32 = 87
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                _ = fi11
            }
        }
        _ = v9
    }
    var v24 string = "hello"
    _ = v24
    _ = v5
    _ = v6
    _ = v7
    return 42
}
func genFn1() bool {
    switch -9223372036854775808 {
    case -1:
        var v25 string = "a"
        _ = v25
    case 399:
        v26 := GenS2{F0: 2, F1: 255}
        _ = v26
    }
    v27 := []int{-47, 2147483647, -652, -2147483648}
    _ = v27
    return false
}

func main() {
    v1 := make(map[string]string)
    var v2 int32 = 0
    v3 := make(map[string]bool)
    var v4 int64 = 0
    v2 = (v2 * (-(v2)))
    var v28 uint32 = 32
    if (-2.25 < -2.25) {
        var v29 int64 = v4
        if ((32767 > 1000) || (!true)) {
            switch 9223372036854775807 {
            case 42:
                v30 := GenS1{F0: (v2 << ((2) & 63))}
                _ = v30
            default:
                var v31 uint32 = v28
                _ = v31
            }
        } else {
            var v32 uint32 = (^(v28))
            _ = v32
        }
        _ = v29
    } else {
        v4 = -46
        var v33 uint8 = 2
        _ = v33
    }
    v2 -= (^(v2))
    fmt.Println(v28, v2, v2)
    v4 = v4
    if true {
        v28 = (^(v28))
        v34 := GenS1{F0: (v2 | (v2 / ((v2) | 1)))}
        _ = v34
    } else {
        if (!false) {
            fmt.Println(v4)
        }
    }
    if false {
        var v35 int16 = -32768
        _ = v35
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
}
