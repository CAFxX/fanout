package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 uint64 }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 uint8; F1 uint32 }
type GenS4 struct { F0 uint64; F1 int8 }
type GenS5 struct { F0 int8; F1 bool }

func main() {
    var v1 uint64 = 884
    var v2 string = "hello"
    var v3 float64 = 3.14159
    var v4 int = -1
    var v5 float64 = 1.5
    v6 := []int{v4, 9223372036854775807}
    var v7 int16 = -554
    v4 = (v4 / (((v4 >> ((1) & 63))) | 1))
    var v8 uint64 = 4294967295
    for fi9 := 0; fi9 < 11; fi9++ {
        if false {
            switch (v2 + (v2 + (v2 + (v2 + v2)))) {
            case "":
                if (true && ((!true) || false)) {
                    for fi10 := 0; fi10 < 1; fi10++ {
                        switch fi10 {
                        case 494:
                            v8 = (v8 + v1)
                        case 128:
                            var v11 int32 = 1
                            _ = v11
                        }
                        _ = fi10
                    }
                    var v12 int16 = 371
                    v13 := GenS1{F0: v2, F1: false, F2: (v1 + 182)}
                    _ = v12
                    _ = v13
                } else {
                    if ((!false) || ((((^(v1)) < 650) || true) && (-2147483648 == 42))) {
                        _ = len(v6)
                        v14 := GenS2{F0: v2, F1: -105}
                        var v15 uint16 = 136
                        _ = v14
                        _ = v15
                    } else {
                        v6 = append(v6, (-(v4)))
                        v16 := GenS2{F0: ("\n\t" + v2), F1: -128}
                        _ = v16
                    }
                    var v17 uint8 = 127
                    switch (v4 & v4) {
                    case 32767:
                        v4 = (fi9 - -2147483648)
                    case 2147483647:
                        if false {
                            var v18 string = (v2 + "hello")
                            v19 := GenS3{F0: (v17 << ((699) & 63)), F1: uint32(v8)}
                            _ = v18
                            _ = v19
                        }
                    default:
                        v3 = (v5 + ((1.5 - 0.1) * v3))
                    }
                    _ = v17
                }
            default:
                var v20 int = (v4 - (v4 / (((fi9 + v4)) | 1)))
                _ = v20
            }
            if true {
                _ = len(v6)
            } else {
                switch (v2 + v2) {
                case "":
                    v2 = "foo bar"
                default:
                    v6 = append(v6, (fi9 + fi9))
                }
                if true {
                    if (true || (1.5 == v5)) {
                        if true {
                            _ = len(v6)
                            v21 := GenS4{F0: 2, F1: -77}
                            _ = v21
                        }
                        if (255 < 1) {
                            if false {
                                var v22 bool = (!(true && (1 >= 255)))
                                v23 := GenS3{F0: 169, F1: 2}
                                _ = v22
                                _ = v23
                            } else {
                                var v24 int64 = int64(v1)
                                v24 = v24
                                _ = v24
                            }
                            var v25 int8 = -1
                            _ = v25
                        }
                    }
                    fmt.Println(v2, v7, v4)
                } else {
                    for fi26 := 0; fi26 < 14; fi26++ {
                        var v27 uint = 65535
                        v28 := GenS4{F0: (v1 + (v1 % ((0) | 1))), F1: 1}
                        v1 += 4294967295
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        var v29 uint16 = 251
                        _ = fi26
                        _ = v27
                        _ = v28
                        _ = v29
                    }
                }
                v30 := GenS2{F0: "foo bar", F1: 42}
                _ = v30
            }
            if (2 <= 18446744073709551615) {
                v1 = (^(v1))
                fmt.Println(fi9, v8)
                v31 := GenS5{F0: 127, F1: false}
                _ = v31
            }
        } else {
            if (v5 < float64(1.5)) {
                v32 := GenS2{F0: (v2 + (("\n\t" + (v2 + v2)) + v2)), F1: 100}
                _ = v32
            }
        }
        if true {
            v33 := GenS2{F0: "", F1: -128}
            _ = v33
        }
        var v34 string = v2
        _ = fi9
        _ = v34
    }
    for fi35 := 0; fi35 < 13; fi35++ {
        var v36 float32 = 1.5
        v37 := GenS3{F0: 230, F1: 2147483647}
        _ = fi35
        _ = v36
        _ = v37
    }
    if len(v6) > 0 {
        v6[0] = v4
    }
    fmt.Println(v5)
    fmt.Println(v5)
    var v38 uint64 = 4294967295
    for fi39 := 0; fi39 < 15; fi39++ {
        v8 -= (^(v38))
        _ = fi39
    }
    switch v4 {
    case -808:
        var v40 int16 = 32767
        _ = v40
    default:
        var v41 float64 = v3
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
