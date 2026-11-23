package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 uint16; F1 uint }

func main() {
    var v1 float64 = -2.25
    v2 := []uint32{2147483648}
    v3 := make(map[uint32]int)
    v4 := []string{"a"}
    var v5 uint8 = 2
    v6 := []int{-406, 646, 9223372036854775807}
    fmt.Println(v5, v5)
    if (0 < 65535) {
        v1 = ((v1 + float64(((1.5 + 3.14159) / 1.5))) / v1)
        v3[1] = 2147483647
    } else {
        if (1.5 >= (1.5 - 1e10)) {
            delete(v3, 431)
            if false {
                for fi7 := 0; fi7 < 17; fi7++ {
                    if ((^(fi7)) != fi7) {
                        var v8 int8 = 100
                        _ = v8
                    } else {
                        v9 := GenS1{F0: 0, F1: 0}
                        v1 = v1
                        _ = v9
                    }
                    fmt.Println(v1, v5)
                    var v10 uint16 = 65535
                    _ = fi7
                    _ = v10
                }
                switch 127 {
                case 32767:
                    var v11 uint64 = 110
                    _ = v11
                case -1:
                    for fi12 := 0; fi12 < 15; fi12++ {
                        var v13 uint16 = 535
                        _ = fi12
                        _ = v13
                    }
                case -210:
                    var v14 uint16 = 42
                    _ = v14
                }
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        var v15 string = ("a" + "")
        v5 += v5
        _ = v15
    }
    fmt.Println(v5)
    v1 = ((v1 - -2.25) * v1)
    if (-9223372036854775808 < 0) {
        for fi16 := 0; fi16 < 6; fi16++ {
            for fi17 := 0; fi17 < 13; fi17++ {
                if len(v4) > 0 {
                    v4[0] = ((("hello" + (("hello" + "a") + "hello")) + "foo bar") + "")
                }
                v5 = v5
                switch (fi17 | fi16) {
                case 1:
                    var v18 int8 = -128
                    _ = v18
                case -511:
                    var v19 uint8 = v5
                    _ = v19
                case 0:
                    switch fi17 {
                    case 128:
                        var v20 int8 = 0
                        _ = v20
                    case 0:
                        if (false && true) {
                            v5 = (v5 | (v5 * v5))
                            if len(v4) > 0 {
                                v4[0] = "a"
                            }
                        } else {
                            var v21 int = (fi16 ^ fi16)
                            fmt.Println(fi17, v21)
                            var v22 uint16 = 0
                            _ = v21
                            _ = v22
                        }
                    case 2147483647:
                        if true {
                            v5 = 1
                        } else {
                            var v23 uint16 = 255
                            if (false && true) {
                                v24 := GenS1{F0: -9223372036854775808, F1: 0}
                                _ = v24
                            } else {
                                var v25 float32 = 3.14159
                                _ = len(v4)
                                fmt.Println(v25)
                                v26 := v3[0]
                                _ = v25
                                _ = v26
                            }
                            _ = v23
                        }
                    }
                }
                _ = fi17
            }
            fmt.Println(v5)
            var v27 int = fi16
            _ = fi16
            _ = v27
        }
    }
    for fi28 := 0; fi28 < 13; fi28++ {
        if (!(!true)) {
            var v29 float32 = 1.5
            for fi30 := 0; fi30 < 4; fi30++ {
                var v31 string = ("\n\t" + "hello")
                fmt.Println(fi30, v5)
                var v32 uint = 172
                _ = fi30
                _ = v31
                _ = v32
            }
            var v33 string = "foo bar"
            _ = v29
            _ = v33
        }
        if (v5 != (v5 + 39)) {
            var v34 float32 = 1.5
            var v35 string = ""
            v1 = 3.14159
            _ = v34
            _ = v35
        } else {
            for fi36 := 0; fi36 < 2; fi36++ {
                if len(v6) > 0 {
                    v6[0] = fi28
                }
                fmt.Println(v5, fi36, v1)
                v37 := v3[1]
                _ = fi36
                _ = v37
            }
            var v38 float64 = ((float64(((((1.5 / (1e10 / 1e10)) - 1e10) - (-2.25 - -2.25)) - float32(v1))) + v1) * v1)
            var v39 uint = 804
            _ = v38
            _ = v39
        }
        v40 := v3[42]
        _ = fi28
        _ = v40
    }
    v5 = 0
    var v41 int = 863
    v5 += (^(v5))
    if (!false) {
        var v42 uint = uint(v5)
        _ = v42
    }
    switch "a" {
    case "":
        v43 := GenS1{F0: -9223372036854775808, F1: 42}
        _ = v43
    default:
        fmt.Println(v1)
    }
    v41 = v41
    v44 := GenS2{F0: 32767, F1: 550}
    var v45 float64 = ((-2.25 * v1) - 3.14159)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
