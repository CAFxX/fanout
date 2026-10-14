package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 uint32; F1 uint16; F2 int8 }
type GenS3 struct { F0 uint32; F1 uint8; F2 uint32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint; F1 bool; F2 bool }

func main() {
    var v1 bool = (32768 == 32767)
    v2 := []string{(("foo bar" + "foo bar") + "hello"), "\n\t", "foo bar", "a"}
    v3 := []string{"\n\t", ("a" + (("" + ("" + (("\n\t" + "hello") + "hello"))) + "hello"))}
    v4 := make(map[string]bool)
    v5 := make(map[string]int)
    var v6 int16 = 200
    switch 42 {
    case 127:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case -405:
        _ = len(v3)
    case -1:
        v1 = v1
    }
    v7 := GenS1{F0: 0, F1: (v6 % (((^(v6))) | 1))}
    var v8 int64 = 9223372036854775807
    var v9 uint16 = 0
    for fi10 := 0; fi10 < 19; fi10++ {
        if ((321 < 18446744073709551615) || false) {
            if v1 {
                switch -840 {
                case 127:
                    v11 := GenS1{F0: (v8 / ((v8) | 1)), F1: v6}
                    _ = v11
                case 9223372036854775807:
                    var v12 bool = v1
                    _ = v12
                }
                if (v1 || (float64(0.0) <= 3.14159)) {
                    fmt.Println(v9)
                    switch (fi10 + fi10) {
                    case 0:
                        v1 = false
                    default:
                        v13 := GenS2{F0: 33, F1: v9, F2: 100}
                        _ = v13
                    }
                } else {
                    fmt.Println(v9)
                    if v1 {
                        var v14 bool = (!v1)
                        if ((v6 & v6) > (v6 - 32767)) {
                            var v15 uint16 = (^(v9))
                            if v14 {
                                if false {
                                    v16 := GenS2{F0: 42, F1: v15, F2: 127}
                                    fmt.Println(v9)
                                    _ = v16
                                } else {
                                    fmt.Println(v14)
                                    v14 = (1 != 18446744073709551615)
                                }
                            }
                            _ = v15
                        } else {
                            if (!(!false)) {
                                _ = len(v3)
                                switch (fi10 << ((59) & 63)) {
                                case 629:
                                    v8 = (v8 * v8)
                                }
                            }
                            _ = len(v4)
                            for fi17 := 0; fi17 < 15; fi17++ {
                                var v18 string = "foo bar"
                                if len(v3) > 0 {
                                    v3[0] = v18
                                }
                                var v19 int = fi10
                                _ = fi17
                                _ = v18
                                _ = v19
                            }
                        }
                        v20 := GenS1{F0: 42, F1: v6}
                        _ = v14
                        _ = v20
                    }
                }
            }
            var v21 int = fi10
            var v22 uint8 = 42
            _ = v21
            _ = v22
        }
        switch (fi10 | fi10) {
        case 4294967295:
            var v23 float32 = (((-2.25 * 0.1) - 0.1) + 1.5)
            _ = v23
        }
        _ = fi10
    }
    if len(v3) > 0 {
        v3[0] = "a"
    }
    var v24 string = "hello"
    if v1 {
        var v25 int64 = -362
        v8 -= -1
        for fi26 := 0; fi26 < 7; fi26++ {
            if true {
                for fi27 := 0; fi27 < 17; fi27++ {
                    var v28 bool = v1
                    v28 = (v1 && ((!false) && ((-(fi26)) >= (fi27 | fi26))))
                    _ = fi27
                    _ = v28
                }
                v4[((((("hello" + v24) + (("hello" + ((v24 + v24) + v24)) + v24)) + (("a" + v24) + "hello")) + v24) + "")] = false
            } else {
                _ = len(v5)
            }
            v29 := GenS3{F0: 936, F1: 255, F2: 65535}
            for fi30 := 0; fi30 < 11; fi30++ {
                var v31 float32 = (1e10 - (-2.25 / ((0.1 / 1e10) + 0.1)))
                _ = fi30
                _ = v31
            }
            _ = fi26
            _ = v29
        }
        _ = v25
    } else {
        v1 = (!v1)
        v32 := GenS4{F0: (^(v9))}
        _ = v32
    }
    if len(v2) > 0 {
        v2[0] = (v24 + (v24 + "a"))
    }
    v2 = append(v2, v24)
    if true {
        var v33 bool = true
        v34 := GenS5{F0: 65535, F1: (!v1), F2: (!true)}
        _ = v33
        _ = v34
    } else {
        var v35 int64 = v8
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
