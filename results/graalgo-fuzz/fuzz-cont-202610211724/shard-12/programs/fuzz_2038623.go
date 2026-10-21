package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 uint64 }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 int8; F1 int8; F2 uint64 }
type GenS4 struct { F0 uint }

func main() {
    var v1 uint = 1
    var v2 int8 = 84
    v3 := make(map[string]int)
    var v4 int32 = 2147483647
    v5 := GenS1{F0: 303, F1: true, F2: 42}
    switch "foo bar" {
    case "a":
        var v6 string = ""
        _ = v6
    case "hello":
        if (((!("a" == "foo bar")) || false) && (4294967295 >= 1)) {
            v7 := GenS1{F0: 2, F1: true, F2: 446}
            v8 := []int{32768, -1, -1}
            _ = v7
            _ = v8
        }
    case "\n\t":
        v9 := GenS1{F0: 42, F1: false, F2: 2}
        _ = v9
    }
    for fi10 := 0; fi10 < 4; fi10++ {
        var v11 int8 = v2
        delete(v3, "\n\t")
        switch "" {
        case "\n\t":
            if false {
                var v12 float32 = (1.5 / 1.5)
                v11 += v2
                _ = v12
            } else {
                var v13 bool = (fi10 > fi10)
                switch "foo bar" {
                case "":
                    fmt.Println(v4)
                default:
                    fmt.Println(v2, fi10, v4)
                }
                for fi14 := 0; fi14 < 1; fi14++ {
                    switch "\n\t" {
                    case "hello":
                        switch ("\n\t" + "\n\t") {
                        case "hello":
                            v3["hello"] = fi10
                        default:
                            if (!true) {
                                if true {
                                    var v15 bool = v13
                                    v16 := GenS2{F0: (v4 + v4), F1: 298}
                                    _ = v15
                                    _ = v16
                                }
                                var v17 uint = 426
                                var v18 uint32 = 383
                                _ = v17
                                _ = v18
                            }
                        }
                    case "foo bar":
                        var v19 float32 = 0.0
                        _ = v19
                    }
                    fmt.Println(fi14, v2)
                    _ = fi14
                }
                _ = v13
            }
        case "":
            v20 := GenS1{F0: 65535, F1: (true && ((!true) && false)), F2: 4294967295}
            _ = v20
        case "hello":
            for fi21 := 0; fi21 < 20; fi21++ {
                if ((!false) && (188 < 128)) {
                    delete(v3, "hello")
                } else {
                    if false {
                        if false {
                            var v22 uint16 = 0
                            v23 := GenS1{F0: (v22 >> ((v1) & 63)), F1: (("a" + ("a" + "a")) < ("\n\t" + "")), F2: 2}
                            if (v2 != 1) {
                                var v24 int = (fi21 & -418)
                                switch "a" {
                                case "\n\t":
                                    v25 := []int{v24, (v24 >> (((v1 * 4294967295)) & 63)), (fi21 % ((fi21) | 1))}
                                    _ = v25
                                case "a":
                                    var v26 int16 = 1000
                                    _ = v26
                                default:
                                    var v27 uint16 = (^(v22))
                                    _ = v27
                                }
                                _ = v24
                            } else {
                                v11 -= (v2 & int8(v22))
                            }
                            _ = v22
                            _ = v23
                        }
                        v28 := GenS2{F0: 42, F1: -1}
                        v11 = -1
                        _ = v28
                    } else {
                        v4 = 2147483647
                    }
                    v29 := GenS2{F0: v4, F1: 32767}
                    var v30 string = "\n\t"
                    _ = v29
                    _ = v30
                }
                _ = fi21
            }
        default:
            _ = len(v3)
        }
        _ = fi10
        _ = v11
    }
    switch 169 {
    case 2147483647:
        delete(v3, "foo bar")
    case 1:
        v31 := GenS2{F0: (v4 / (((v4 / (((v4 & (v4 >> (((^(v1))) & 63)))) | 1))) | 1)), F1: -763}
        _ = v31
    case 127:
        v2 = (v2 + (^(v2)))
    }
    v32 := GenS2{F0: 1, F1: -32768}
    if true {
        var v33 uint64 = 18446744073709551615
        _ = v33
    } else {
        if (575 <= 42) {
            var v34 float32 = 0.1
            _ = v34
        }
        if true {
            v35 := GenS2{F0: (v4 + v4), F1: -453}
            _ = v35
        }
        for fi36 := 0; fi36 < 16; fi36++ {
            var v37 uint8 = 173
            if (v4 == (v4 & (v4 ^ 42))) {
                var v38 uint = 4294967295
                var v39 int = (fi36 & 32768)
                _ = v38
                _ = v39
            } else {
                if (!false) {
                    var v40 float32 = ((1.5 + 3.14159) - float32(((0.0 - 3.14159) - 1.5)))
                    v4 = int32(v37)
                    for fi41 := 0; fi41 < 1; fi41++ {
                        switch "\n\t" {
                        case "":
                            v42 := GenS3{F0: -128, F1: (v2 + v2), F2: 1}
                            _ = v42
                        default:
                            v37 = 0
                        }
                        if true {
                            v43 := GenS4{F0: uint(v37)}
                            _ = v43
                        } else {
                            fmt.Println(v40)
                            var v44 uint8 = v37
                            var v45 uint = (^(v1))
                            _ = v44
                            _ = v45
                        }
                        var v46 uint8 = v37
                        _ = fi41
                        _ = v46
                    }
                    _ = v40
                }
                var v47 int = 42
                _ = v47
            }
            delete(v3, ("hello" + ""))
            _ = fi36
            _ = v37
        }
    }
    v48 := v3["\n\t"]
    var v49 int32 = (-(v4))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v48
    _ = v49
    _ = v5
}
