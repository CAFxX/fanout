package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 uint }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint; F1 int }

func main() {
    var v1 uint64 = 42
    var v2 uint16 = 776
    var v3 int32 = -110
    var v4 uint = 65535
    v1 = v1
    fmt.Println(v4, v2)
    switch ("foo bar" + ("" + "")) {
    case "\n\t":
        v5 := []int{42, 0, -1, 42}
        _ = v5
    default:
        for fi6 := 0; fi6 < 6; fi6++ {
            if true {
                v4 = (^(v4))
            } else {
                if ((("\n\t" + ("a" + ("\n\t" + "a"))) + "hello") > "a") {
                    for fi7 := 0; fi7 < 1; fi7++ {
                        v8 := []int{32767}
                        _ = fi7
                        _ = v8
                    }
                    var v9 float32 = 0.1
                    if (2 >= 42) {
                        var v10 float64 = 3.14159
                        v11 := GenS1{F0: -615, F1: -1, F2: v4}
                        fmt.Println(v4)
                        _ = v10
                        _ = v11
                    }
                    _ = v9
                } else {
                    switch "\n\t" {
                    case "":
                        var v12 uint = v4
                        _ = v12
                    case "a":
                        for fi13 := 0; fi13 < 12; fi13++ {
                            v14 := []int{(fi13 + fi6)}
                            if false {
                                switch (fi13 * 128) {
                                case 32767:
                                    var v15 int8 = -98
                                    _ = v15
                                case 4294967295:
                                    if ((v3 + v3) < (v3 & (-(v3)))) {
                                        v1 = 18446744073709551615
                                    }
                                case 127:
                                    v3 += 50
                                }
                                v16 := GenS2{F0: (v3 % (((v3 & v3)) | 1)), F1: 65535}
                                _ = v16
                            }
                            _ = fi13
                            _ = v14
                        }
                    case "\n\t":
                        v17 := GenS3{F0: 1}
                        _ = v17
                    default:
                        fmt.Println(v4, v4)
                    }
                }
            }
            _ = fi6
        }
    }
    var v18 uint64 = v1
    v19 := []int{42, -144, 318, int(v2)}
    fmt.Println(v1, v3, v18)
    v20 := GenS4{F0: (v4 & v4), F1: 32768}
    if false {
        var v21 string = "foo bar"
        _ = v21
    } else {
        if true {
            fmt.Println(v1)
        } else {
            if ((((1e10 * -2.25) + ((0.1 / (1e10 - 3.14159)) * (1.5 - 3.14159))) < float64((((0.0 - -2.25) * 0.0) * 0.1))) && true) {
                for fi22 := 0; fi22 < 20; fi22++ {
                    if (0 >= 127) {
                        switch ("hello" + ("\n\t" + "a")) {
                        case "\n\t":
                            if len(v19) >= 2 {
                                v19 = v19[1:len(v19)]
                            }
                        case "foo bar":
                            v18 -= (v1 * (v1 ^ v1))
                        case "":
                            if len(v19) >= 2 {
                                v19 = v19[1:len(v19)]
                            }
                        }
                        v1 = (v1 / (((^(v1))) | 1))
                        for fi23 := 0; fi23 < 7; fi23++ {
                            if len(v19) > 0 {
                                v19[0] = -368
                            }
                            _ = fi23
                        }
                    }
                    for fi24 := 0; fi24 < 4; fi24++ {
                        fmt.Println(fi24)
                        if len(v19) >= 2 {
                            v19 = v19[1:len(v19)]
                        }
                        var v25 uint = (v4 % ((v4) | 1))
                        v26 := GenS4{F0: (v25 << (((v25 % (((v4 * (v4 + 18446744073709551615))) | 1))) & 63)), F1: (fi24 ^ (fi24 << ((v4) & 63)))}
                        var v27 uint16 = v2
                        _ = fi24
                        _ = v25
                        _ = v26
                        _ = v27
                    }
                    _ = fi22
                }
                v28 := GenS1{F0: -399, F1: -597, F2: 2}
                _ = v28
            } else {
                var v29 float64 = float64((-2.25 - 0.0))
                _ = v29
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
