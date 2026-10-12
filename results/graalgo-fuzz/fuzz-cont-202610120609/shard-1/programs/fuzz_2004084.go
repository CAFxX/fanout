package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16; F1 uint32; F2 bool }
type GenS3 struct { F0 bool; F1 uint32 }
type GenS4 struct { F0 int16; F1 int32 }
type GenS5 struct { F0 int64; F1 int16; F2 uint8 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int; F1 uint32; F2 uint16 }
func genFn0() uint64 {
    v5 := GenS1{F0: ("" + "\n\t")}
    for fi6 := 0; fi6 < 19; fi6++ {
        v7 := GenS2{F0: 2, F1: 293, F2: (false && (!(!false)))}
        var v8 int16 = 32767
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 string = "hello"
    switch (("hello" + "") + "foo bar") {
    case "foo bar":
        v9 = (v9 + (v9 + (("a" + v9) + (v9 + v9))))
    case "\n\t":
        var v10 string = v9
        _ = v10
    }
    var v11 uint16 = 1
    _ = v11
    _ = v5
    _ = v9
    return 1
}

func main() {
    var v1 float32 = ((0.0 + 0.1) - -2.25)
    var v2 float32 = v1
    var v3 uint64 = 676
    v4 := []string{("foo bar" + ("\n\t" + "a")), ("a" + ((("a" + "hello") + (("foo bar" + "hello") + (("a" + "foo bar") + ""))) + "a")), ("hello" + "")}
    var v12 int = 128
    if ("a" >= ("hello" + "\n\t")) {
        v13 := GenS3{F0: (!false), F1: 0}
        _ = len(v4)
        var v14 bool = false
        _ = v13
        _ = v14
    }
    if (!false) {
        _ = len(v4)
        switch "a" {
        case "\n\t":
            if (!false) {
                switch "" {
                case "a":
                    v2 = v2
                case "":
                    switch "\n\t" {
                    case "":
                        for fi15 := 0; fi15 < 11; fi15++ {
                            v16 := GenS1{F0: "a"}
                            for fi17 := 0; fi17 < 10; fi17++ {
                                var v18 uint = 0
                                var v19 uint32 = 1
                                _ = fi17
                                _ = v18
                                _ = v19
                            }
                            v20 := GenS2{F0: 65535, F1: 660, F2: true}
                            _ = fi15
                            _ = v16
                            _ = v20
                        }
                    }
                }
                fmt.Println(v12)
            } else {
                if (!(42 >= 1000)) {
                    var v21 int = v12
                    v22 := GenS4{F0: -32768, F1: 773}
                    var v23 float64 = ((0.0 + 1.5) * 1e10)
                    _ = v21
                    _ = v22
                    _ = v23
                }
                v24 := GenS4{F0: 1000, F1: 2147483647}
                _ = v24
            }
        case "a":
            var v25 float64 = -2.25
            _ = v25
        default:
            var v26 int8 = -1
            _ = v26
        }
        if false {
            var v27 int16 = 1
            if false {
                _ = len(v4)
                v28 := GenS5{F0: -9223372036854775808, F1: v27, F2: 128}
                _ = len(v4)
                _ = v28
            }
            for fi29 := 0; fi29 < 14; fi29++ {
                v30 := GenS1{F0: "a"}
                v27 = (v27 * v27)
                fmt.Println(v1)
                _ = fi29
                _ = v30
            }
            _ = v27
        } else {
            if (42 >= (v12 * v12)) {
                v31 := GenS6{F0: "foo bar"}
                v3 += genFn0()
                v32 := GenS7{F0: v12, F1: 745, F2: 42}
                _ = v31
                _ = v32
            } else {
                v33 := GenS2{F0: 1, F1: 517, F2: false}
                _ = v33
            }
            fmt.Println(v12, v12, v2)
        }
    }
    var v34 int64 = 550
    for fi35 := 0; fi35 < 1; fi35++ {
        if false {
            var v36 int16 = 9
            _ = v36
        }
        _ = fi35
    }
    var v37 int8 = 127
    if len(v4) > 0 {
        v4[0] = "hello"
    }
    switch -2147483648 {
    case 661:
        fmt.Println(v12)
    case 0:
        for fi38 := 0; fi38 < 11; fi38++ {
            v3 += (v3 / ((v3) | 1))
            v39 := GenS3{F0: false, F1: 2}
            for fi40 := 0; fi40 < 19; fi40++ {
                if (!false) {
                    v12 += v12
                    switch (^(fi40)) {
                    case 496:
                        var v41 uint32 = 412
                        _ = v41
                    case -1:
                        var v42 uint64 = (^(v3))
                        _ = v42
                    case -9223372036854775808:
                        _ = len(v4)
                    default:
                        switch "a" {
                        case "foo bar":
                            var v43 int = (v12 & fi40)
                            _ = v43
                        }
                        _ = len(v4)
                        v2 = ((v2 / (v1 / v1)) / 1e10)
                    }
                }
                _ = fi40
            }
            _ = fi38
            _ = v39
        }
    default:
        var v44 uint32 = 852
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v34
    _ = v37
    _ = v4
}
