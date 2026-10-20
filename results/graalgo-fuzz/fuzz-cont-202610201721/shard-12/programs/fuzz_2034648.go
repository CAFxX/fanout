package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32; F1 uint; F2 int16 }
type GenS4 struct { F0 int32; F1 uint64 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int32; F1 uint32 }
type GenS7 struct { F0 string; F1 uint8 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 bool }

func main() {
    var v1 float32 = -2.25
    var v2 uint64 = 2
    var v3 uint64 = 42
    var v4 uint = 0
    var v5 int8 = -1
    var v6 int = 4294967295
    if (("\n\t" + (((("\n\t" + "") + "hello") + ((("hello" + ((("foo bar" + "hello") + ((("hello" + "hello") + "a") + "foo bar")) + "a")) + "hello") + "foo bar")) + "a")) != "hello") {
        v7 := GenS1{F0: -32768, F1: 255, F2: 0}
        _ = v7
    }
    v8 := []int{}
    for fi9 := 0; fi9 < 1; fi9++ {
        switch 4294967295 {
        case 158:
            if true {
                switch ("" + "foo bar") {
                case "a":
                    fmt.Println(v6, v5)
                case "foo bar":
                    for fi10 := 0; fi10 < 15; fi10++ {
                        fmt.Println(v4, v3)
                        _ = fi10
                    }
                case "":
                    var v11 int64 = 1
                    _ = v11
                default:
                    for fi12 := 0; fi12 < 19; fi12++ {
                        if false {
                            var v13 float64 = -2.25
                            var v14 int16 = 870
                            _ = v13
                            _ = v14
                        } else {
                            if (42 < 0) {
                                v15 := GenS2{F0: (v4 ^ v4)}
                                _ = v15
                            }
                        }
                        var v16 uint32 = 0
                        switch "foo bar" {
                        case "\n\t":
                            var v17 int = 58
                            _ = v17
                        case "foo bar":
                            var v18 string = "hello"
                            _ = v18
                        case "hello":
                            if (42 == 2) {
                                switch "" {
                                case "foo bar":
                                    switch v6 {
                                    case -910:
                                        if len(v8) >= 2 {
                                            v8 = v8[1:len(v8)]
                                        }
                                    case 388:
                                        if (-2147483648 != 0) {
                                            var v19 uint64 = (v2 * (^(v3)))
                                            _ = len(v8)
                                            _ = v19
                                        } else {
                                            v20 := GenS3{F0: 986, F1: (^(v4)), F2: -32768}
                                            v21 := GenS4{F0: 205, F1: v2}
                                            if false {
                                                if (!false) {
                                                    var v22 uint16 = 2
                                                    _ = v22
                                                }
                                            }
                                            _ = v20
                                            _ = v21
                                        }
                                    case 127:
                                        var v23 uint64 = v2
                                        _ = v23
                                    default:
                                        v2 = (v2 & (^(v3)))
                                        v24 := GenS3{F0: 0, F1: (^(v4)), F2: 32767}
                                        _ = v24
                                    }
                                    if len(v8) > 0 {
                                        v8[0] = fi12
                                    }
                                    if ((!(true || false)) && false) {
                                        fmt.Println(v4, v2)
                                    }
                                    var v25 uint16 = 32767
                                    v26 := GenS1{F0: -391, F1: 66, F2: uint8(v4)}
                                    v27 := GenS5{F0: "hello"}
                                    _ = v25
                                    _ = v26
                                    _ = v27
                                default:
                                    switch "hello" {
                                    case "a":
                                        v28 := GenS6{F0: -825, F1: v16}
                                        _ = v28
                                    case "\n\t":
                                        switch ("hello" + ("" + "hello")) {
                                        case "a":
                                            switch (("hello" + "\n\t") + "hello") {
                                            case "hello":
                                                var v29 string = ((("\n\t" + "foo bar") + "foo bar") + "")
                                                _ = v29
                                            default:
                                                var v30 float64 = 3.14159
                                                _ = v30
                                            }
                                            v31 := GenS3{F0: 1, F1: (v4 + v4), F2: 32767}
                                            v1 = v1
                                            _ = v31
                                        case "hello":
                                            v3 = uint64(v4)
                                        case "\n\t":
                                            fmt.Println(v3)
                                        default:
                                            v2 -= (v3 / ((v2) | 1))
                                        }
                                    case "hello":
                                        var v32 string = ""
                                        fmt.Println(v5)
                                        _ = v32
                                    }
                                }
                                v16 += (v16 >> ((v4) & 63))
                                var v33 uint64 = 1
                                v34 := GenS7{F0: (((((("foo bar" + "") + "\n\t") + "\n\t") + "a") + "") + (("\n\t" + "foo bar") + (("" + ("a" + "\n\t")) + ("a" + ("hello" + (("foo bar" + "foo bar") + "a")))))), F1: 127}
                                _ = v33
                                _ = v34
                            }
                        default:
                            if len(v8) >= 2 {
                                v8 = v8[1:len(v8)]
                            }
                        }
                        fmt.Println(v4, v3)
                        _ = fi12
                        _ = v16
                    }
                }
                v6 = (-(v6))
            } else {
                for fi35 := 0; fi35 < 5; fi35++ {
                    v36 := GenS8{F0: v4}
                    var v37 int32 = 42
                    _ = fi35
                    _ = v36
                    _ = v37
                }
            }
        default:
            switch 42 {
            case -308:
                v1 -= v1
            case 282:
                if true {
                    var v38 uint32 = uint32(v4)
                    v3 += v2
                    _ = v38
                }
            case 588:
                for fi39 := 0; fi39 < 12; fi39++ {
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                    var v40 uint32 = 1
                    if false {
                        var v41 int16 = 1000
                        if len(v8) >= 2 {
                            v8 = v8[1:len(v8)]
                        }
                        _ = v41
                    }
                    _ = fi39
                    _ = v40
                }
            default:
                v1 -= 1.5
            }
        }
        _ = fi9
    }
    var v42 float32 = 0.0
    switch ("foo bar" + "\n\t") {
    case "\n\t":
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
    case "a":
        if false {
            var v43 int8 = (-(v5))
            _ = v43
        } else {
            fmt.Println(v3, v1, v4)
            v44 := GenS4{F0: -1, F1: (v3 + v3)}
            _ = v44
        }
    case "":
        for fi45 := 0; fi45 < 3; fi45++ {
            v42 -= v1
            v4 += (v4 >> ((430) & 63))
            v4 = v4
            _ = fi45
        }
    default:
        v8 = append(v8, v6)
    }
    v42 = v1
    if ((false && (533 != 339)) || true) {
        fmt.Println(v42, v1)
    } else {
        var v46 uint64 = v2
        v47 := GenS9{F0: false}
        switch 128 {
        case -1:
            v48 := GenS4{F0: -919, F1: 2}
            _ = v48
        case -444:
            fmt.Println(v1, v4, v4)
        case 32767:
            v6 = 127
        }
        _ = v46
        _ = v47
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v6
    _ = v8
}
