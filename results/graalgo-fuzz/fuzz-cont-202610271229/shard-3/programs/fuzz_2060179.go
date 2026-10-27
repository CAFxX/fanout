package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 string; F2 int64 }
func genFn0(p0 int) bool {
    if (0 != 198) {
        v6 := []int{-1, 32767, -2147483648, p0}
        var v7 float32 = -2.25
        _ = v6
        _ = v7
    }
    if (629 < 65535) {
        var v8 int16 = -596
        var v9 uint16 = 1
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: -128}
    v11 := GenS1{F0: 41}
    if (42 <= 1) {
        v12 := []int{p0, -1, p0, p0}
        v13 := GenS2{F0: 42, F1: "foo bar", F2: 9223372036854775807}
        _ = v12
        _ = v13
    }
    _ = p0
    _ = v10
    _ = v11
    return true
}

func main() {
    var v1 int64 = 0
    var v2 uint8 = 1
    v3 := make(map[uint32]bool)
    v4 := make(map[uint32]int)
    var v5 int8 = 22
    switch "foo bar" {
    case "hello":
        var v14 string = "\n\t"
        _ = v14
    default:
        for fi15 := 0; fi15 < 2; fi15++ {
            if true {
                switch "\n\t" {
                case "":
                    if (((v5 & (v5 ^ 47)) >= v5) && false) {
                        fmt.Println(v2)
                        if ((fi15 & (fi15 * -9223372036854775808)) != (fi15 / (((fi15 - fi15)) | 1))) {
                            switch ("\n\t" + ("" + ("\n\t" + "\n\t"))) {
                            case "hello":
                                fmt.Println(v5)
                            case "":
                                for fi16 := 0; fi16 < 19; fi16++ {
                                    fmt.Println(v1)
                                    if false {
                                        v17 := GenS1{F0: (v5 & v5)}
                                        fmt.Println(fi16, fi16, v5)
                                        _ = v17
                                    }
                                    if (!(v2 <= (^(v2)))) {
                                        v3[65535] = (0.1 < (1.5 / 1.5))
                                        v18 := GenS2{F0: -32768, F1: "foo bar", F2: v1}
                                        v3[373] = true
                                        _ = v18
                                    }
                                    _ = fi16
                                }
                            case "a":
                                var v19 bool = false
                                _ = v19
                            }
                            v1 = (v1 * (v1 / (((v1 / ((v1) | 1))) | 1)))
                            fmt.Println(v1, v2, fi15)
                        } else {
                            fmt.Println(v5, v2)
                            switch ("a" + "foo bar") {
                            case "a":
                                var v20 uint32 = 1
                                _ = v20
                            }
                        }
                    } else {
                        v2 += 127
                        var v21 uint8 = 42
                        v2 -= v21
                        _ = v21
                    }
                case "hello":
                    fmt.Println(v2)
                default:
                    var v22 int = (fi15 ^ 128)
                    _ = v22
                }
            }
            fmt.Println(v1, v1, v2)
            _ = fi15
        }
    }
    var v23 uint = 351
    for fi24 := 0; fi24 < 10; fi24++ {
        v25 := GenS2{F0: 0, F1: "hello", F2: 1}
        switch 0 {
        case 739:
            v23 += v23
        default:
            var v26 uint64 = 1
            _ = v26
        }
        _ = fi24
        _ = v25
    }
    var v27 uint32 = 2
    switch -270 {
    case -1:
        v2 += 107
    }
    var v28 string = (("hello" + "hello") + "")
    if ((0.0 * ((-2.25 - 0.1) - float64(float32(0.1)))) == 0.0) {
        _ = len(v4)
        switch (v28 + "a") {
        case "\n\t":
            v2 = (v2 % (((v2 & (v2 << ((v23) & 63)))) | 1))
        case "a":
            fmt.Println(v23)
        case "foo bar":
            for fi29 := 0; fi29 < 12; fi29++ {
                v5 += (^(v5))
                fmt.Println(fi29, v1, v2)
                _ = fi29
            }
        default:
            var v30 float32 = (float32((1e10 - float64(3.14159))) + 0.1)
            _ = v30
        }
    }
    for fi31 := 0; fi31 < 16; fi31++ {
        var v32 int64 = 1
        var v33 float64 = 1.5
        _ = fi31
        _ = v32
        _ = v33
    }
    var v34 int64 = (v1 | 0)
    var v35 int = 32768
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v27
    _ = v28
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
