package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int16; F1 uint8 }

func main() {
    var v1 uint16 = 0
    var v2 uint8 = 127
    v3 := make(map[int]string)
    v4 := []string{}
    v1 += (v1 % (((v1 << ((925) & 63))) | 1))
    if (0 >= v1) {
        v5 := GenS1{F0: -1, F1: 180}
        for fi6 := 0; fi6 < 20; fi6++ {
            if true {
                switch fi6 {
                case -466:
                    var v7 int = fi6
                    _ = v7
                case 32768:
                    var v8 uint = 2
                    _ = v8
                }
                switch (fi6 ^ fi6) {
                case 128:
                    var v9 float64 = 3.14159
                    _ = v9
                case -9223372036854775808:
                    if (false && (0 > 32767)) {
                        fmt.Println(fi6, fi6)
                    } else {
                        v10 := GenS2{F0: ""}
                        if true {
                            if true {
                                v11 := GenS1{F0: 1000, F1: 381}
                                switch (fi6 >> ((2) & 63)) {
                                case 4294967295:
                                    if false {
                                        v12 := GenS3{F0: 255}
                                        _ = v12
                                    }
                                case 778:
                                    if true {
                                        var v13 int8 = -1
                                        v14 := v3[fi6]
                                        _ = v13
                                        _ = v14
                                    } else {
                                        var v15 bool = true
                                        _ = v15
                                    }
                                default:
                                    var v16 float32 = (-2.25 + 1.5)
                                    _ = v16
                                }
                                _ = v11
                            }
                            var v17 bool = (v1 <= 32767)
                            fmt.Println(v1, v17)
                            _ = v17
                        } else {
                            v1 += 2
                            fmt.Println(fi6, v1)
                        }
                        v18 := GenS4{F0: -32768}
                        _ = v10
                        _ = v18
                    }
                default:
                    for fi19 := 0; fi19 < 15; fi19++ {
                        v1 += 0
                        var v20 uint32 = 2147483647
                        var v21 float32 = 3.14159
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                }
            } else {
                if (true || true) {
                    v22 := GenS1{F0: 0, F1: 18446744073709551615}
                    v3[32767] = "foo bar"
                    _ = v22
                } else {
                    v3[fi6] = (("hello" + "\n\t") + "")
                }
            }
            _ = fi6
        }
        _ = v5
    }
    var v23 int8 = -128
    var v24 int64 = 42
    var v25 uint = 65535
    v26 := GenS2{F0: "hello"}
    for fi27 := 0; fi27 < 20; fi27++ {
        v2 = v2
        fmt.Println(v24, v1, v23)
        v28 := GenS5{F0: 0, F1: (^(v2))}
        _ = fi27
        _ = v28
    }
    v23 += (^(v23))
    _ = len(v4)
    var v29 bool = true
    v2 += (v2 & v2)
    for fi30 := 0; fi30 < 6; fi30++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        for fi31 := 0; fi31 < 2; fi31++ {
            _ = len(v4)
            v2 = v2
            _ = fi31
        }
        _ = fi30
    }
    v32 := GenS2{F0: "hello"}
    v3[839] = (((("foo bar" + ("foo bar" + ("hello" + ("hello" + ("hello" + ""))))) + ("" + "hello")) + "foo bar") + ((("" + "\n\t") + "") + "foo bar"))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v32
    _ = v4
}
