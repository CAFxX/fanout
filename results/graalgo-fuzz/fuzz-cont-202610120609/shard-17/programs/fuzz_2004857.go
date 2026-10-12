package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 uint16; F1 int32 }

func main() {
    v1 := []string{"\n\t", "\n\t", (("" + ((("\n\t" + (("foo bar" + ("\n\t" + (("foo bar" + "") + ""))) + "\n\t")) + "\n\t") + ("" + ""))) + "hello")}
    v2 := make(map[string]bool)
    v3 := make(map[int]int)
    var v4 uint8 = 2
    var v5 int = 613
    _ = len(v2)
    if false {
        fmt.Println(v4)
    }
    v4 += 2
    if (v5 == v5) {
        v6 := GenS1{F0: false, F1: -444}
        var v7 int8 = -128
        for fi8 := 0; fi8 < 7; fi8++ {
            v2["a"] = (320 != 42)
            var v9 uint32 = 629
            v9 += v9
            _ = fi8
            _ = v9
        }
        _ = v6
        _ = v7
    }
    if len(v1) > 0 {
        v1[0] = (("hello" + ("foo bar" + ("" + "hello"))) + "hello")
    }
    var v10 string = ("hello" + "")
    fmt.Println(v5, v10)
    v11 := GenS1{F0: (498 < 42), F1: 834}
    switch v5 {
    case 2147483647:
        for fi12 := 0; fi12 < 19; fi12++ {
            v10 = v10
            switch v10 {
            case "\n\t":
                var v13 uint16 = 739
                _ = v13
            case "":
                if true {
                    var v14 uint8 = 40
                    if false {
                        v4 = v14
                        v10 = (v10 + ("foo bar" + v10))
                        fmt.Println(v5, v4, v10)
                    } else {
                        v5 += (fi12 + 0)
                    }
                    _ = v14
                } else {
                    for fi15 := 0; fi15 < 6; fi15++ {
                        switch int(v4) {
                        case -2147483648:
                            var v16 int64 = -465
                            v17 := GenS2{F0: 1, F1: -720}
                            if ((0 <= 255) && false) {
                                v4 -= v4
                            }
                            v18 := GenS1{F0: (!true), F1: -1}
                            _ = v16
                            _ = v17
                            _ = v18
                        case 549:
                            if (false && false) {
                                v19 := GenS1{F0: ((^(v4)) > v4), F1: -32768}
                                switch v10 {
                                case "foo bar":
                                    v5 += -705
                                case "":
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                                var v20 int = fi15
                                _ = v19
                                _ = v20
                            } else {
                                if (!false) {
                                    v21 := GenS2{F0: 961, F1: -776}
                                    var v22 int64 = 42
                                    _ = v21
                                    _ = v22
                                }
                                v4 -= 128
                            }
                        }
                        _ = fi15
                    }
                    _ = len(v2)
                }
            case "foo bar":
                fmt.Println(fi12, v5, v4)
            }
            var v23 uint8 = v4
            _ = fi12
            _ = v23
        }
    case -9223372036854775808:
        var v24 float32 = float32(-2.25)
        _ = v24
    default:
        var v25 uint32 = 2
        _ = v25
    }
    v26 := GenS1{F0: true, F1: 42}
    delete(v2, v10)
    _ = len(v1)
    delete(v2, v10)
    for fi27 := 0; fi27 < 16; fi27++ {
        for fi28 := 0; fi28 < 9; fi28++ {
            v4 = v4
            _ = fi28
        }
        _ = fi27
    }
    switch "hello" {
    case "hello":
        fmt.Println(v4, v5)
    case "\n\t":
        var v29 float64 = 0.0
        _ = v29
    case "a":
        if (983 <= 32767) {
            v30 := GenS2{F0: uint16(v4), F1: -1}
            _ = v30
        } else {
            v31 := GenS1{F0: true, F1: 981}
            if false {
                fmt.Println(v10, v10, v10)
                v1 = append(v1, (v10 + v10))
                v4 -= 0
            }
            v10 = ""
            _ = v31
        }
    }
    v5 = v5
    var v32 int = (^(v5))
    var v33 int32 = -1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v26
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
