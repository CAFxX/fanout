package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int; F1 int64; F2 int16 }

func main() {
    var v1 uint32 = 2
    v2 := make(map[string]bool)
    var v3 int32 = -138
    var v4 float64 = ((0.1 * ((float64(((float32(1.5) / 3.14159) * 1e10)) / 0.1) + 0.0)) - 3.14159)
    v3 -= (v3 - -2147483648)
    v1 = 786
    for fi5 := 0; fi5 < 19; fi5++ {
        fmt.Println(v3, fi5)
        _ = fi5
    }
    v6 := []int{-429, 128, int(v3)}
    var v7 int = 354
    v8 := GenS1{F0: v7, F1: 1}
    var v9 int32 = v3
    v7 -= v7
    v2["hello"] = false
    for fi10 := 0; fi10 < 8; fi10++ {
        v4 = v4
        var v11 float32 = 1.5
        for fi12 := 0; fi12 < 13; fi12++ {
            var v13 int64 = 42
            _ = fi12
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    if true {
        var v14 int64 = 1
        fmt.Println(v14, v4)
        _ = v14
    }
    if true {
        var v15 int8 = -87
        var v16 int32 = v9
        _ = v15
        _ = v16
    } else {
        v1 -= (v1 - 4294967295)
        for fi17 := 0; fi17 < 8; fi17++ {
            for fi18 := 0; fi18 < 16; fi18++ {
                if false {
                    delete(v2, "")
                    switch fi18 {
                    case 9223372036854775807:
                        v7 = fi17
                    case 32768:
                        v7 += fi18
                    }
                    switch (("\n\t" + "") + ("hello" + "")) {
                    case "a":
                        if (192 < 52) {
                            if false {
                                _ = len(v2)
                                fmt.Println(v9, v3)
                            }
                            v19 := GenS2{F0: (v3 ^ v9)}
                            _ = v19
                        } else {
                            if (18446744073709551615 < 42) {
                                var v20 bool = false
                                if false {
                                    _ = len(v2)
                                }
                                _ = v20
                            } else {
                                switch ("\n\t" + (("foo bar" + ("" + "foo bar")) + "")) {
                                case "foo bar":
                                    if (!false) {
                                        if len(v6) > 0 {
                                            v6[0] = (^(v7))
                                        }
                                    }
                                default:
                                    var v21 uint8 = 42
                                    v4 += (float64(float32(-2.25)) + (((float64(-2.25) * 1e10) - (v4 - (float64(-2.25) - v4))) * 1e10))
                                    _ = v21
                                }
                                v22 := GenS3{F0: 1}
                                _ = v22
                            }
                            if true {
                                v3 += v9
                                var v23 uint8 = 2
                                var v24 int64 = 882
                                _ = v23
                                _ = v24
                            }
                            var v25 float64 = v4
                            fmt.Println(v7, v1, v1)
                            fmt.Println(v3, v9)
                            _ = v25
                        }
                    }
                    fmt.Println(fi17)
                }
                _ = fi18
            }
            var v26 uint64 = 18446744073709551615
            _ = fi17
            _ = v26
        }
        for fi27 := 0; fi27 < 16; fi27++ {
            switch ("\n\t" + "hello") {
            case "":
                var v28 int64 = 838
                _ = v28
            case "\n\t":
                v1 = v1
            }
            delete(v2, "a")
            v29 := GenS4{F0: 559, F1: -9223372036854775808, F2: 42}
            _ = fi27
            _ = v29
        }
    }
    if (!(false || ("hello" <= ("hello" + (("hello" + "") + "hello"))))) {
        v3 = -2147483648
        if true {
            if ((!((false || false) || (-1 <= 0))) || false) {
                v2["\n\t"] = true
            }
        } else {
            v9 += -1
        }
        fmt.Println(v1)
    }
    switch "\n\t" {
    case "\n\t":
        v3 += (^(v3))
    default:
        var v30 uint = 65535
        _ = v30
    }
    var v31 float64 = v4
    var v32 uint64 = 359
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
